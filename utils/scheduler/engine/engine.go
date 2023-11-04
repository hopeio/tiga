package engine

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgraph-io/ristretto"
	"github.com/hopeio/lemon/utils/io/fs"
	"github.com/hopeio/lemon/utils/log"
	rate2 "github.com/hopeio/lemon/utils/scheduler/rate"
	"github.com/hopeio/lemon/utils/slices"
	"github.com/hopeio/lemon/utils/struct/heap"
	"github.com/hopeio/lemon/utils/struct/list/list"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type Config[KEY comparable, T, W any] struct {
	WorkerCount uint
}

func (c *Config[KEY, T, W]) NewEngine() *Engine[KEY, T, W] {
	return NewEngine[KEY, T, W](c.WorkerCount)
}

type Engine[KEY comparable, T, W any] struct {
	limitWorkerCount, currentWorkerCount uint64
	limitWaitTaskCount                   uint
	workerChan                           chan *Worker[KEY, T, W]
	workers                              []*Worker[KEY, T, W]
	workerList                           list.List[*Worker[KEY, T, W]]
	taskChan                             chan *Task[KEY, T]
	taskList                             heap.Heap[*Task[KEY, T], Tasks[KEY, T]]
	ctx                                  context.Context
	cancel                               context.CancelFunc   // 手动停止执行
	wg                                   sync.WaitGroup       // 控制确保所有任务执行完
	fixedWorkers                         []chan *Task[KEY, T] // 固定只执行一种任务的worker,避免并发问题
	rateTimer                            *time.Ticker
	speedLimit                           *rate2.SpeedLimiter
	rateLimiter                          *rate.Limiter
	//TODO
	monitorInterval       time.Duration // 全局检测定时器间隔时间，任务的卡住检测，worker panic recover都可以用这个检测
	isRunning, isFinished bool
	EngineStatistics
	done *ristretto.Cache
	//TasksChan   chan []*Task[KEY, T]
	kindHandler  []*KindHandler[KEY, T]
	errHandler   func(task *Task[KEY, T])
	errChan      chan *Task[KEY, T]
	stopCallBack []func()
}

type KindHandler[KEY comparable, T any] struct {
	Skip        bool
	rateTimer   *time.Ticker
	speedLimit  *rate2.SpeedLimiter
	rateLimiter *rate.Limiter
	// TODO 指定Kind的Handler
	HandleFun TaskFunc[KEY, T]
}

func NewEngine[KEY comparable, T, W any](workerCount uint) *Engine[KEY, T, W] {
	return NewEngineWithContext[KEY, T, W](workerCount, context.Background())
}

func NewEngineWithContext[KEY comparable, T, W any](workerCount uint, ctx context.Context) *Engine[KEY, T, W] {
	ctx, cancel := context.WithCancel(ctx)
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters:        1e4,   // number of keys to track frequency of (10M).
		MaxCost:            1e3,   // maximum cost of cache (MaxCost * 1MB).
		BufferItems:        64,    // number of keys per Get buffer.
		Metrics:            false, // number of keys per Get buffer.
		IgnoreInternalCost: true,
	})
	return &Engine[KEY, T, W]{
		limitWorkerCount:   uint64(workerCount),
		limitWaitTaskCount: workerCount * 10,
		ctx:                ctx,
		cancel:             cancel,
		workerChan:         make(chan *Worker[KEY, T, W]),
		taskChan:           make(chan *Task[KEY, T]),
		workerList:         list.New[*Worker[KEY, T, W]](),
		taskList:           heap.Heap[*Task[KEY, T], Tasks[KEY, T]]{},
		monitorInterval:    time.Second,
		done:               cache,
		errHandler: func(task *Task[KEY, T]) {
			log.Error(task.errs)
		},
		errChan: make(chan *Task[KEY, T]),
	}
}

func (e *Engine[KEY, T, W]) Context() context.Context {
	return e.ctx
}

func (e *Engine[KEY, T, W]) SkipKind(kinds ...Kind) *Engine[KEY, T, W] {
	length := slices.Max(kinds) + 1
	if e.kindHandler == nil {
		e.kindHandler = make([]*KindHandler[KEY, T], length)
	}
	if int(length) > len(e.kindHandler) {
		e.kindHandler = append(e.kindHandler, make([]*KindHandler[KEY, T], int(length)-len(e.kindHandler))...)
	}
	for _, kind := range kinds {
		if e.kindHandler[kind] == nil {
			e.kindHandler[kind] = &KindHandler[KEY, T]{Skip: true}
		} else {
			e.kindHandler[kind].Skip = true
		}

	}
	return e
}

func (e *Engine[KEY, T, W]) MonitorInterval(interval time.Duration) {
	e.monitorInterval = interval
}

func (e *Engine[KEY, T, W]) ErrHandler(errHandler func(task *Task[KEY, T])) *Engine[KEY, T, W] {
	e.errHandler = errHandler
	return e
}

func (e *Engine[KEY, T, W]) ErrHandlerUtilSuccess() *Engine[KEY, T, W] {
	return e.ErrHandler(func(task *Task[KEY, T]) {
		task.errs = task.errs[:0]
		e.AsyncAddTasks(task.Priority, task)
	})
}

func (e *Engine[KEY, T, W]) ErrHandlerRetryTimes(times int) *Engine[KEY, T, W] {
	return e.ErrHandler(func(task *Task[KEY, T]) {
		task.errs = task.errs[:0]
		if task.errTimes < times {
			e.AsyncAddTasks(task.Priority, task)
		} else {
			log.Error(task.errs)
		}

	})
}

func (e *Engine[KEY, T, W]) ErrHandlerWriteToFile(path string) *Engine[KEY, T, W] {
	file, err := fs.Create(path)
	if err != nil {
		panic(err)
	}
	e.StopCallBack(func() {
		file.Close()
	})
	return e.ErrHandler(func(task *Task[KEY, T]) {
		spew.Fdump(file, task)
	})
}

func (e *Engine[KEY, T, W]) StopCallBack(callBack func()) *Engine[KEY, T, W] {
	e.stopCallBack = append(e.stopCallBack, callBack)
	return e
}

func (e *Engine[KEY, T, W]) SpeedLimited(interval time.Duration) *Engine[KEY, T, W] {
	e.speedLimit = rate2.NewSpeedLimiter(interval)
	return e
}

func (e *Engine[KEY, T, W]) RandSpeedLimited(start, stop time.Duration) *Engine[KEY, T, W] {
	e.speedLimit = rate2.NewRandSpeedLimiter(start, stop)
	return e
}

func (e *Engine[KEY, T, W]) RateTimer(interval time.Duration) *Engine[KEY, T, W] {
	e.rateTimer = time.NewTicker(interval)
	return e
}

func (e *Engine[KEY, T, W]) KindRateTimer(kind Kind, interval time.Duration) *Engine[KEY, T, W] {
	ticker := time.NewTicker(interval)
	e.kindRateTimer(kind, ticker)
	return e
}

func (e *Engine[KEY, T, W]) kindRateTimer(kind Kind, ticker *time.Ticker) *Engine[KEY, T, W] {
	if e.kindHandler == nil {
		e.kindHandler = make([]*KindHandler[KEY, T], int(kind)+1)
	}
	if int(kind)+1 > len(e.kindHandler) {
		e.kindHandler = append(e.kindHandler, make([]*KindHandler[KEY, T], int(kind)+1-len(e.kindHandler))...)
	}
	if e.kindHandler[kind] == nil {
		e.kindHandler[kind] = &KindHandler[KEY, T]{rateTimer: ticker}
	} else {
		e.kindHandler[kind].rateTimer = ticker
	}
	return e
}

// 多个kind共用一个timer
func (e *Engine[KEY, T, W]) KindGroupRateTimer(interval time.Duration, kinds ...Kind) *Engine[KEY, T, W] {
	ticker := time.NewTicker(interval)
	for _, kind := range kinds {
		e.kindRateTimer(kind, ticker)
	}
	return e
}

func (e *Engine[KEY, T, W]) Limiter(r rate.Limit, b int) *Engine[KEY, T, W] {
	e.rateLimiter = rate.NewLimiter(r, b)
	return e
}

func (e *Engine[KEY, T, W]) KindLimiter(kind Kind, r rate.Limit, b int) *Engine[KEY, T, W] {
	e.kindLimiter(kind, r, b)
	return e
}

func (e *Engine[KEY, T, W]) kindLimiter(kind Kind, r rate.Limit, b int) {
	if e.kindHandler == nil {
		e.kindHandler = make([]*KindHandler[KEY, T], int(kind)+1)
	}
	if int(kind)+1 > len(e.kindHandler) {
		e.kindHandler = append(e.kindHandler, make([]*KindHandler[KEY, T], int(kind)+1-len(e.kindHandler))...)
	}
	if e.kindHandler[kind] == nil {
		e.kindHandler[kind] = &KindHandler[KEY, T]{rateLimiter: rate.NewLimiter(r, b)}
	} else {
		e.kindHandler[kind].rateLimiter = rate.NewLimiter(r, b)
	}
}

// TaskSourceChannel 任务源,参数是一个channel,channel关闭时，代表任务源停止发送任务
func (e *Engine[KEY, T, W]) TaskSourceChannel(taskSourceChannel <-chan *Task[KEY, T]) {
	e.wg.Add(1)
	go func() {
		for task := range taskSourceChannel {
			if task == nil || task.TaskFunc == nil {
				continue
			}
			e.AddTasks(0, task)
		}
		e.wg.Done()
	}()
}

// TaskSourceFunc,参数为添加任务的函数，直到该函数运行结束，任务引擎才会检测任务是否结束
func (e *Engine[KEY, T, W]) TaskSourceFunc(taskSourceFunc func(*Engine[KEY, T, W])) {
	e.wg.Add(1)
	go func() {
		taskSourceFunc(e)
		e.wg.Done()
	}()
}
