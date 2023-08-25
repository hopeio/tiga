package ctrl

import "github.com/hopeio/lemon/utils/errors/multierr"

type Controller chan func() error

func New() Controller {
	return make(Controller)
}

func (c Controller) AddTask(f func() error) {
	go func() {
		c <- f
	}()
}

func (c Controller) Start() {
	for f := range c {
		err := f()
		if err != nil {
			c.AddTask(f)
		}
	}
}

func ReTry(times int, f func() error) error {
	var errs multierr.MultiError
	for i := 0; i < times; i++ {
		err := f()
		if err == nil {
			return nil
		}
		errs.Append(err)
	}
	if errs.HasErrors() {
		return &errs
	}
	return nil
}
