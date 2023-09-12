package listener

import (
	"context"
	"github.com/hopeio/lemon/utils/scheduler/rate"
	"github.com/hopeio/lemon/utils/scheduler/tiny_engine"
	"time"
)

type TimerTask struct {
	Times     uint
	FirstExec bool
	Do        tiny_engine.TaskFunc
}

func (task *TimerTask) Timer(ctx context.Context, interval time.Duration) {
	timer := time.NewTicker(interval)
	if task.FirstExec {
		task.Times = 1
		task.Do(ctx)
	}
	for {
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
			task.Times++
			task.Do(ctx)
		}
	}
}

func (task *TimerTask) RandTimer(ctx context.Context, start, stop time.Duration) {
	timer := rate.NewRandSpeedLimiter(start, stop)
	task.Times = 1
	task.Do(ctx)
	for {
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
			task.Times++
			task.Do(ctx)
			timer.Reset()
		}
	}
}
