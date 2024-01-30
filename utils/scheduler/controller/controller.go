package controller

import "github.com/hopeio/tiga/utils/errors/multierr"

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
