package engine

import (
	"context"
	"fmt"
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
		e.AddNoPriorityTasks(genTask(id))
		if id == 10 {
			break
		}
	}
}

func genTask(id int) *Task[int, Prop] {
	return &Task[int, Prop]{
		TaskMeta: TaskMeta[int]{Key: id},
		TaskFunc: func(ctx context.Context) ([]*Task[int, Prop], error) {
			fmt.Println(id)
			return []*Task[int, Prop]{genTask2(id * 10)}, nil
		},
		errs:  nil,
		Props: Prop{},
	}
}

func genTask2(id int) *Task[int, Prop] {
	return &Task[int, Prop]{
		TaskMeta: TaskMeta[int]{Key: id},
		TaskFunc: func(ctx context.Context) ([]*Task[int, Prop], error) {
			fmt.Println(id)
			return nil, nil
		},
		errs:  nil,
		Props: Prop{},
	}
}
