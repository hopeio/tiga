package parallel

import (
	"github.com/hopeio/tiga/utils/definition/constraints"
	"github.com/hopeio/tiga/utils/errors/multierr"
	"sync"
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
	wg      sync.WaitGroup
}

func New(workNum int) *Parallel {
	return &Parallel{taskCh: make(chan constraints.FuncWithErr, workNum), workNum: workNum, wg: sync.WaitGroup{}}
}

func (p *Parallel) Run() {
	for i := 0; i < p.workNum; i++ {
		go func() {
			for task := range p.taskCh {
				err := task()
				p.wg.Done()
				if err != nil {
					go p.AddTask(task)
				}
			}
		}()
	}
}

func (p *Parallel) AddTask(task constraints.FuncWithErr) {
	p.wg.Add(1)
	p.taskCh <- task
}

func (p *Parallel) Wait() {
	p.wg.Wait()
}
