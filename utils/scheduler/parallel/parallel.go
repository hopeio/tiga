package parallel

import "github.com/hopeio/lemon/utils/errors/multierr"

func Run(tasks []func() error) error {
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
