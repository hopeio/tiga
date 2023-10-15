package engine

import (
	"context"
	"github.com/hopeio/lemon/utils/definition/constraints"
	"time"
)

type Kind uint8

const (
	KindNormal = iota
)

type TaskMetaNew[T constraints.Key[KEY], KEY comparable] struct{}

type TaskMeta[KEY comparable] struct {
	id          uint64
	Kind        Kind
	Key         KEY
	Priority    int
	Describe    string
	createdAt   time.Time
	execBeginAt time.Time
	execEndAt   time.Time
	TaskStatistics
}

func (t *TaskMeta[KEY]) OrderKey() int {
	return t.Priority
}

func (t *TaskMeta[KEY]) SetPriority(priority int) {
	t.Priority = priority
}

func (r *TaskMeta[KEY]) SetKind(k Kind) {
	r.Kind = k
}

func (r *TaskMeta[KEY]) SetKey(key KEY) {
	r.Key = key
}

func (r *TaskMeta[KEY]) Id() uint64 {
	return r.id
}

type TaskStatistics struct {
	reDoTimes uint
	errTimes  int
}

type Task[KEY comparable, P any] struct {
	TaskMeta[KEY]
	TaskFunc[KEY, P]
	errs  []error
	Props P
}

func (t *Task[KEY, P]) Errs() []error {
	return t.errs
}

type Tasks[KEY comparable, P any] []*Task[KEY, P]

func (tasks Tasks[KEY, T]) Less(i, j int) bool {
	return tasks[i].Priority > tasks[j].Priority
}

// ---------------

type ErrHandle func(context.Context, error)

type TaskFunc[KEY comparable, P any] func(ctx context.Context) ([]*Task[KEY, P], error)

func emptyTaskFunc[KEY comparable, P any](ctx context.Context) ([]*Task[KEY, P], error) {
	return nil, nil
}
