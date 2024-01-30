package crawler

import (
	"context"
	"github.com/hopeio/tiga/utils/scheduler/engine"
)

type Prop struct {
}

type Request = engine.Task[string, Prop]
type TaskMeta = engine.TaskMeta[string]
type TaskFunc = engine.TaskFunc[string, Prop]

func NewRequest(key string, kind engine.Kind, taskFunc TaskFunc) *Request {
	return &Request{
		TaskMeta: TaskMeta{
			Key:  key,
			Kind: kind,
		},
		TaskFunc: taskFunc,
	}
}

type Config = engine.Config[string, Prop, Prop]
type Engine = engine.Engine[string, Prop, Prop]

func NewEngine(workerCount uint) *engine.Engine[string, Prop, Prop] {
	return engine.NewEngine[string, Prop, Prop](workerCount)
}

type HandleFunc func(ctx context.Context, url string) ([]*Request, error)

func NewUrlRequest(url string, handleFunc HandleFunc) *Request {
	if handleFunc == nil {
		return nil
	}
	return &Request{TaskMeta: TaskMeta{Key: url}, TaskFunc: func(ctx context.Context) ([]*Request, error) {
		return handleFunc(ctx, url)
	}}
}

func NewUrlKindRequest(url string, kind engine.Kind, handleFunc HandleFunc) *Request {
	if handleFunc == nil {
		return nil
	}
	req := NewUrlRequest(url, handleFunc)
	req.SetKind(kind)
	return req
}

func NewTaskMeta(key string) TaskMeta {
	return TaskMeta{Key: key}
}
