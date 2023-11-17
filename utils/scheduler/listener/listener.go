package listener

import (
	"context"
	"github.com/hopeio/lemon/utils/scheduler/rate"
	"github.com/hopeio/lemon/utils/scheduler/tiny_engine"
	"time"
)

type TimerTask struct {
	times         uint
	firstExecuted bool
	Do            tiny_engine.TaskFunc
}

func (task *TimerTask) Times() uint {
	return task.times
}

func (task *TimerTask) Timer(ctx context.Context, interval time.Duration) {
	timer := time.NewTicker(interval)
	if !task.firstExecuted {
		task.times = 1
		task.Do(ctx)
		task.firstExecuted = true
	}
	for {
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
			task.times++
			task.Do(ctx)
		}
	}
}

func (task *TimerTask) RandTimer(ctx context.Context, start, stop time.Duration) {
	timer := rate.NewRandSpeedLimiter(start, stop)
	task.times = 1
	task.Do(ctx)
	for {
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
			task.times++
			task.Do(ctx)
			timer.Reset()
		}
	}
}
