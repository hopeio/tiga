package engine

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/hopeio/lemon/utils/generator"
	"github.com/hopeio/lemon/utils/log"
	synci "github.com/hopeio/lemon/utils/sync"
	"runtime/debug"
	"sync/atomic"
	"time"
)

func (e *Engine[KEY, T, W]) Run(tasks ...*Task[KEY, T]) {
	if !e.ran {
		go func() {
			for task := range e.errChan {
				e.taskErrCount++
				e.errHandler(task)
			}
		}()
		e.addWorker()
		e.ran = true
	}

	e.isRunning = true
	go func() {
		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()
		var emptyTimes uint
		var readyWorkerCh chan *Task[KEY, T]
		var readyTask *Task[KEY, T]
	loop:
		for {
			if e.workerList.Size > 0 && len(e.taskList) > 0 {
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
					if e.workerList.Size == uint(e.currentWorkerCount) && len(e.taskList) == 0 {
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
	if e.stopCallBack != nil {
		for _, callBack := range e.stopCallBack {
			callBack()
		}
	}
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
			e.ExecTask(e.ctx, readyTask)
		}
		for {
			select {
			case e.workerChan <- worker:
				readyTask = <-taskChan
				worker.isExecuting = true
				e.ExecTask(e.ctx, readyTask)
				worker.isExecuting = false
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

func (e *Engine[KEY, T, W]) ReRun(tasks ...*Task[KEY, T]) {
	if e.isRunning {
		e.AddTasks(0, tasks...)
		return
	}
	e.Run(tasks...)
}

func (e *Engine[KEY, T, W]) AddNoPriorityTasks(tasks ...*Task[KEY, T]) {
	e.AddTasks(0, tasks...)
}

func (e *Engine[KEY, T, W]) AddTasks(generation int, tasks ...*Task[KEY, T]) {
	l := len(tasks)
	atomic.AddUint64(&e.taskTotalCount, uint64(l))
	e.wg.Add(l)
	for _, task := range tasks {
		// 如果task为nil,补一个什么都不做的task,为了减少atomic.AddUint64和e.wg.Add的调用次数
		if task == nil {
			task = &Task[KEY, T]{TaskFunc: emptyTaskFunc[KEY, T]}
		}
		task.Priority += generation
		task.id = generator.GenOrderID()
		e.taskChan <- task
	}
}

func (e *Engine[KEY, T, W]) AsyncAddTasks(generation int, tasks ...*Task[KEY, T]) {
	if len(tasks) > 0 {
		go e.AddTasks(generation, tasks...)
	}
}

func (e *Engine[KEY, T, W]) AddWorker(num int) {
	atomic.AddUint64(&e.limitWorkerCount, uint64(num))
}

func (e *Engine[KEY, T, W]) NewFixedWorker(interval time.Duration) int {
	ch := make(chan *Task[KEY, T])
	e.fixedWorkers = append(e.fixedWorkers, ch)
	e.newFixedWorker(ch, interval)
	return len(e.fixedWorkers) - 1
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
			e.ExecTask(e.ctx, task)
		}
	}()
}

func (e *Engine[KEY, T, W]) AddFixedTasks(workerId int, generation int, tasks ...*Task[KEY, T]) error {

	if workerId > len(e.fixedWorkers)-1 {
		return fmt.Errorf("不存在workId为%d的worker,请调用NewFixedWorker添加", workerId)
	}
	ch := e.fixedWorkers[workerId]
	l := len(tasks)
	atomic.AddUint64(&e.taskTotalCount, uint64(l))
	e.wg.Add(l)
	for _, task := range tasks {
		if task == nil {
			task = &Task[KEY, T]{TaskFunc: emptyTaskFunc[KEY, T]}
		}
		task.Priority += generation
		task.id = generator.GenOrderID()
		ch <- task
	}
	return nil
}

func (e *Engine[KEY, T, W]) RunSingleWorker(tasks ...*Task[KEY, T]) {
	e.limitWorkerCount = 1
	e.Run(tasks...)
}

func (e *Engine[KEY, T, W]) Stop() {
	e.cancel()
	close(e.workerChan)
	close(e.taskChan)
	for _, ch := range e.fixedWorkers {
		close(ch)
	}
	if e.speedLimit != nil {
		e.speedLimit.Stop()
	}
	e.done.Close()
	for _, kindHandler := range e.kindHandler {
		if kindHandler != nil {
			if kindHandler.rateTimer != nil {
				kindHandler.rateTimer.Stop()
			}
			if kindHandler.rateLimiter != nil {
				kindHandler.rateLimiter = nil
			}
		}
	}
}

func (e *Engine[KEY, T, W]) ExecTask(ctx context.Context, task *Task[KEY, T]) {

	if task != nil {
		if task.TaskFunc != nil {
			var kindHandler *KindHandler[KEY, T]
			if e.kindHandler != nil && int(task.Kind) < len(e.kindHandler) {
				kindHandler = e.kindHandler[task.Kind]
			}

			if kindHandler != nil && kindHandler.Skip {
				atomic.AddUint64(&e.taskDoneCount, 1)
				e.wg.Done()
				return
			}

			zeroKey := *new(KEY)

			if task.Key != zeroKey {
				if _, ok := e.done.Get(task.Key); ok {
					atomic.AddUint64(&e.taskDoneCount, 1)
					e.wg.Done()
					return
				}
			}
			if kindHandler != nil {
				if kindHandler.rateTimer != nil {
					<-kindHandler.rateTimer.C
				}
				if kindHandler.speedLimit != nil {
					<-kindHandler.speedLimit.C
				}
				if kindHandler.rateLimiter != nil {
					kindHandler.rateLimiter.Wait(ctx)
				}
			}

			if e.rateTimer != nil {
				<-e.rateTimer.C
			}

			if e.speedLimit != nil {
				<-e.speedLimit.C
				e.speedLimit.Reset()
			}

			if e.rateLimiter != nil {
				e.rateLimiter.Wait(ctx)
			}

			tasks, err := task.TaskFunc(ctx)
			if err != nil {
				task.errTimes++
				task.errs = append(task.errs, err)
				if len(task.errs) < 5 {
					task.reDoTimes++
					log.Warnf("%v执行失败:%v,将第%d次执行", task.Key, err, task.reDoTimes+1)
					e.AsyncAddTasks(task.Priority+1, task)
				}
				if len(task.errs) == 5 {
					log.Warn(task.Key, "多次执行失败:", err, ",将执行错误处理")
					e.errChan <- task
				}
				atomic.AddUint64(&e.taskDoneCount, 1)
				e.wg.Done()
				return
			}
			if task.Key != zeroKey {
				e.done.SetWithTTL(task.Key, struct{}{}, 1, time.Hour)
			}
			if len(tasks) > 0 {
				e.AsyncAddTasks(task.Priority+1, tasks...)
			}
		}
	}

	atomic.AddUint64(&e.taskDoneCount, 1)
	e.wg.Done()
}

func (e *Engine[KEY, T, W]) Cancel() {
	log.Info("任务取消")
	e.cancel()
	synci.WaitGroupStopWait(&e.wg)

}

func (e *Engine[KEY, T, W]) StopAfter(interval time.Duration) *Engine[KEY, T, W] {
	time.AfterFunc(interval, e.Cancel)
	return e
}
