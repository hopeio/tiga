package tiny_engine

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/hopeio/lemon/utils/generator"
	"github.com/hopeio/lemon/utils/log"
	rate2 "github.com/hopeio/lemon/utils/scheduler/rate"
	"github.com/hopeio/lemon/utils/struct/heap"
	"github.com/hopeio/lemon/utils/struct/list/list"
	synci "github.com/hopeio/lemon/utils/sync"
	"golang.org/x/time/rate"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type Config[KEY comparable, T, W any] struct {
	WorkerCount uint
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
	fixedWorker                          []chan *Task[KEY, T] // 固定只执行一种任务的worker,避免并发问题
	speedLimit                           rate2.SpeedLimiter
	rateLimiter                          *rate.Limiter
	//TODO
	monitorInterval time.Duration // 全局检测定时器间隔时间，任务的卡住检测，worker panic recover都可以用这个检测
	EngineStatistics
	isRunning, isFinished bool
}

func NewEngine[KEY comparable, T, W any](workerCount uint) *Engine[KEY, T, W] {
	return NewEngineWithContext[KEY, T, W](workerCount, context.Background())
}

func NewEngineWithContext[KEY comparable, T, W any](workerCount uint, ctx context.Context) *Engine[KEY, T, W] {
	ctx, cancel := context.WithCancel(ctx)
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
	}
}

func (e *Engine[KEY, T, W]) Context() context.Context {
	return e.ctx
}

func (e *Engine[KEY, T, W]) SpeedLimited(interval time.Duration) {
	e.speedLimit = rate2.NewSpeedLimiter(interval)
}

func (e *Engine[KEY, T, W]) RandSpeedLimited(start, stop time.Duration) {
	e.speedLimit = rate2.NewRandSpeedLimiter(start, stop)
}

func (e *Engine[KEY, T, W]) MonitorInterval(interval time.Duration) {
	e.monitorInterval = interval
}

func (e *Engine[KEY, T, W]) Cancel() {
	log.Info("任务取消")
	e.cancel()
	synci.WaitGroupStopWait(&e.wg)

}

func (e *Engine[KEY, T, W]) Run(tasks ...*Task[KEY, T]) {
	e.addWorker()
	e.isRunning = true
	go func() {
		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()
		var emptyTimes uint
		var readyWorkerCh chan *Task[KEY, T]
		var readyTask *Task[KEY, T]
	loop:
		for {
			if e.workerList.Len() > 0 && len(e.taskList) > 0 {
				if readyWorkerCh == nil {
					readyWorkerCh = e.workerList.Pop().taskCh
				}
				if readyTask == nil {
					readyTask = e.taskList.Pop()
				}
			}

			if len(e.taskList) >= int(e.limitWaitTaskCount) {
				select {
				case readyWorker := <-e.workerChan:
					e.workerList.Push(readyWorker)
				case readyWorkerCh <- readyTask:
					readyWorkerCh = nil
					readyTask = nil
				case <-e.ctx.Done():
					break loop
				}
			} else {
				select {
				case readyTaskTmp := <-e.taskChan:
					e.taskList.Push(readyTaskTmp)
				case readyWorker := <-e.workerChan:
					e.workerList.Push(readyWorker)
				case readyWorkerCh <- readyTask:
					readyWorkerCh = nil
					readyTask = nil
				case <-timer.C:
					//检测任务是否已空
					if e.workerList.Len() == uint(e.currentWorkerCount) && len(e.taskList) == 0 {
						counter, _ := synci.WaitGroupState(&e.wg)
						if counter == 1 {
							emptyTimes++
							if emptyTimes > 2 {
								log.Info("任务即将结束")
								e.wg.Done()
								break loop
							}
						}

					}
					timer.Reset(e.monitorInterval)
				case <-e.ctx.Done():
					break loop
				}
			}
		}
	}()

	e.taskTotalCount += uint64(len(tasks))
	e.wg.Add(len(tasks) + 1)
	for _, task := range tasks {
		if task == nil {
			task = &Task[KEY, T]{TaskFunc: emptyTaskFunc[KEY, T]}
		}
		task.id = generator.GenOrderID()
		e.taskChan <- task
	}

	e.wg.Wait()
	e.isRunning = false
	e.isFinished = true
	log.Infof("任务结束,total:%d,done:%d,failed:%d", e.taskTotalCount, e.taskDoneCount, e.taskFailedCount)
}

func (e *Engine[KEY, T, W]) newWorker(readyTask *Task[KEY, T]) {
	atomic.AddUint64(&e.currentWorkerCount, 1)
	//id := c.currentWorkerCount
	taskChan := make(chan *Task[KEY, T])
	worker := &Worker[KEY, T, W]{Id: uint(e.currentWorkerCount), taskCh: taskChan}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Error(r)
				log.Error(string(debug.Stack()))
				log.Info(spew.Sdump(readyTask))
				atomic.AddUint64(&e.taskFailedCount, 1)
				e.wg.Done()
				// 创建一个新的
				e.newWorker(nil)
			}
			atomic.AddUint64(&e.currentWorkerCount, ^uint64(0))
		}()
		if readyTask != nil {
			if readyTask.TaskFunc != nil {
				readyTask.TaskFunc(e.ctx)
			}
			atomic.AddUint64(&e.taskDoneCount, 1)
			e.wg.Done()
		}
		for {
			select {
			case e.workerChan <- worker:
				readyTask = <-taskChan
				if readyTask != nil && readyTask.TaskFunc != nil {
					if e.speedLimit != nil {
						e.speedLimit.Wait()
					}
					if readyTask.ctx == nil {
						readyTask.ctx = e.ctx
					}
					readyTask.TaskFunc(readyTask.ctx)
				}
				atomic.AddUint64(&e.taskDoneCount, 1)
				e.wg.Done()
			case <-e.ctx.Done():
				return
			}
		}
	}()
	e.workers = append(e.workers, worker)
}

func (e *Engine[KEY, T, W]) addWorker() {
	if e.currentWorkerCount != 0 {
		return
	}
	e.newWorker(nil)
	go func() {
		for {
			select {
			case readyTask := <-e.taskChan:
				if e.currentWorkerCount < e.limitWorkerCount {
					e.newWorker(readyTask)
				} else {
					log.Info("worker count is full")
					e.taskChan <- readyTask
					return
				}
			case <-e.ctx.Done():
				return
			}
		}
	}()

}

func (e *Engine[KEY, T, W]) AddTasks(tasks ...*Task[KEY, T]) {
	atomic.AddUint64(&e.taskTotalCount, uint64(len(tasks)))
	e.wg.Add(len(tasks))
	for _, task := range tasks {
		if task == nil {
			task = &Task[KEY, T]{TaskFunc: emptyTaskFunc[KEY, T]}
		}
		task.id = generator.GenOrderID()
		e.taskChan <- task
	}
}

func (e *Engine[KEY, T, W]) AddWorker(num int) {
	atomic.AddUint64(&e.limitWorkerCount, uint64(num))
}

func (e *Engine[KEY, T, W]) NewFixedWorker(interval time.Duration) int {
	ch := make(chan *Task[KEY, T])
	e.fixedWorker = append(e.fixedWorker, ch)
	e.newFixedWorker(ch, interval)
	return len(e.fixedWorker) - 1
}

func (e *Engine[KEY, T, W]) newFixedWorker(ch chan *Task[KEY, T], interval time.Duration) {
	go func() {
		var task *Task[KEY, T]
		defer func() {
			if r := recover(); r != nil {
				log.Error(r)
				log.Error(string(debug.Stack()))
				log.Info(spew.Sdump(task))
				atomic.AddUint64(&e.taskFailedCount, 1)
				e.wg.Done()
				// 创建一个新的
				e.newFixedWorker(ch, interval)
			}
			atomic.AddUint64(&e.currentWorkerCount, ^uint64(0))
		}()
		var timer *time.Ticker
		// 如果有任务时间间隔,
		if interval > 0 {
			timer = time.NewTicker(interval)
		}
		for task = range ch {
			if interval > 0 {
				<-timer.C
			}
			if task.ctx == nil {
				task.ctx = e.ctx
			}
			task.TaskFunc(task.ctx)
			atomic.AddUint64(&e.taskDoneCount, 1)
			e.wg.Done()
		}
	}()
}

func (e *Engine[KEY, T, W]) AddFixedTasks(workerId int, tasks ...*Task[KEY, T]) error {
	if workerId > len(e.fixedWorker)-1 {
		return fmt.Errorf("不存在workId为%d的worker,请调用NewFixedWorker添加", workerId)
	}
	ch := e.fixedWorker[workerId]
	l := len(tasks)
	atomic.AddUint64(&e.taskTotalCount, uint64(l))
	e.wg.Add(l)
	for _, task := range tasks {
		if task == nil {
			task = &Task[KEY, T]{TaskFunc: emptyTaskFunc[KEY, T]}
		}
		task.id = generator.GenOrderID()
		ch <- task
	}
	return nil
}

func (e *Engine[KEY, T, W]) SyncRun(tasks ...*Task[KEY, T]) {
	panic("TODO")
}

func (e *Engine[KEY, T, W]) RunSingleWorker(tasks ...*Task[KEY, T]) {
	e.NewFixedWorker(0)
	for _, task := range tasks {
		e.AddFixedTasks(0, task)
	}
}

func (e *Engine[KEY, T, W]) Stop() {
	e.cancel()
	close(e.workerChan)
	close(e.taskChan)
	for _, ch := range e.fixedWorker {
		close(ch)
	}
	if e.speedLimit != nil {
		e.speedLimit.Stop()
	}
}
