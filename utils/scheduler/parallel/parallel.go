package parallel

import (
	"github.com/hopeio/lemon/utils/definition/constraints"
	"github.com/hopeio/lemon/utils/errors/multierr"
)

func Run(tasks []constraints.FuncWithErr) error {
	ch := make(chan error)
	for _, task := range tasks {
		go func() {
			ch <- task()
		}()
	}
	var errs multierr.MultiError
	for err := range ch {
		if err != nil {
			errs.Append(err)
		}
	}
	if errs.HasErrors() {
		return &errs
	}
	return nil
}

type Parallel struct {
	taskCh  chan constraints.FuncWithErr
	workNum int
}

func New(workNum int) *Parallel {
	return &Parallel{taskCh: make(chan constraints.FuncWithErr, workNum), workNum: workNum}
}

func (p *Parallel) Run() {
	for i := 0; i < p.workNum; i++ {
		go func() {
			for task := range p.taskCh {
				err := task()
				if err != nil {
					go p.AddTask(task)
				}
			}
		}()
	}
}

func (p *Parallel) AddTask(task constraints.FuncWithErr) {
	p.taskCh <- task
}
