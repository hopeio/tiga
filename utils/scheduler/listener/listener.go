package listener

import (
	"context"
	"github.com/hopeio/tiga/utils/scheduler/rate"
	"github.com/hopeio/tiga/utils/scheduler/tiny_engine"
	"time"
)

type TimerTask struct {
	times         uint
	firstExecuted bool
	do            tiny_engine.TaskFunc
}

func NewTimerTask() *TimerTask {
	return &TimerTask{}
}

func (task *TimerTask) Times() uint {
	return task.times
}

func (task *TimerTask) Run(ctx context.Context, interval time.Duration, do tiny_engine.TaskFunc) {
	task.do = do
	timer := time.NewTicker(interval)
	if !task.firstExecuted {
		task.times = 1
		task.do(ctx)
		task.firstExecuted = true
	}
	for {
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
			task.times++
			task.do(ctx)
		}
	}
}

func (task *TimerTask) RandRun(ctx context.Context, minInterval, maxInterval time.Duration, do tiny_engine.TaskFunc) {
	task.do = do
	timer := rate.NewRandSpeedLimiter(minInterval, maxInterval)
	ch := timer.Channel()
	task.times = 1
	task.do(ctx)
	for {
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-ch:
			task.times++
			task.do(ctx)
			timer.Reset()
		}
	}
}
