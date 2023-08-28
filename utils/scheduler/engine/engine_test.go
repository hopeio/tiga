package engine

import (
	"context"
	"errors"
	"strconv"
	"testing"
)

type Prop struct {
}

func TestEngine(t *testing.T) {
	engine := NewEngine[int, Prop, Prop](5)
	engine.ErrHandlerUtilSuccess()
	engine.TaskSourceFunc(taskSourceFunc)
	engine.Run()
}

func taskSourceFunc(e *Engine[int, Prop, Prop]) {
	var id int
	for {
		id++
		e.AddNoPriorityTasks(&Task[int, Prop]{
			TaskMeta: TaskMeta[int]{Key: id},
			TaskFunc: func(ctx context.Context) ([]*Task[int, Prop], error) {
				return nil, errors.New(strconv.Itoa(id))
			},
			errs:  nil,
			Props: Prop{},
		})
		if id == 10 {
			break
		}
	}

}
