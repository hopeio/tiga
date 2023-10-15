package tiny_engine

import (
	"context"
	"time"
)

type TaskFunc func(context.Context)

type Task[KEY comparable, T any] struct {
	TaskMeta[KEY]
	TaskFunc
	Props T
}

type Tasks[KEY comparable, T any] []*Task[KEY, T]

func (tasks Tasks[KEY, T]) Less(i, j int) bool {
	return tasks[i].Priority > tasks[j].Priority
}

type TaskMeta[KEY comparable] struct {
	id          uint64
	Key         KEY
	Describe    string
	Priority    int
	createdAt   time.Time
	execBeginAt time.Time
	execEndAt   time.Time
}

func (t *TaskMeta[KEY]) SetPriority(priority int) {
	t.Priority = priority
}

func emptyTaskFunc[KEY comparable, P any](ctx context.Context) {
}
