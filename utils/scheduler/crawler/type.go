package crawler

import (
	"context"
	"github.com/hopeio/tiga/utils/scheduler/engine"
)

type Request = engine.Task[string]
type TaskMeta = engine.TaskMeta[string]
type TaskFunc = engine.TaskFunc[string]

func NewRequest(key string, kind engine.Kind, taskFunc TaskFunc) *Request {
	return &Request{
		TaskMeta: TaskMeta{
			Key:  key,
			Kind: kind,
		},
		TaskFunc: taskFunc,
	}
}

type Config = engine.Config[string]
type Engine = engine.Engine[string]

func NewEngine(workerCount uint) *engine.Engine[string] {
	return engine.NewEngine[string](workerCount)
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
