package fs

import (
	"github.com/hopeio/lemon/utils/errors/multierr"
	"os"
)

// 遍历根目录中的每个文件，为每个文件调用callback,不包括目录,与filepath.WalkDir不同的是回调函数的参数不同,filepath.WalkDir的第一个参数是文件完整路径,RangeFile是文件所在目录的路径
func RangeFile(dir string, callback func(dir string, entry os.DirEntry) error) error {
	entities, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	errs := &multierr.MultiError{}
	for _, e := range entities {
		if e.IsDir() {
			err = RangeFile(dir+PathSeparator+e.Name(), callback)
			if err != nil {
				errs.Append(err)
			}
		}
		err = callback(dir, e)
		if err != nil {
			errs.Append(err)
		}
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

// 遍历根目录中的每个文件夹，为文件夹调用callback
// callback 返回值为需要递归遍历的目录和error
// 几乎每个文件夹下的文件夹都会被循环两次！
func RangeDir(dir string, callback func(dir string, entities []os.DirEntry) ([]os.DirEntry, error)) error {
	entities, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	errs := &multierr.MultiError{}

	dirs, err := callback(dir, entities)
	if err != nil {
		errs.Append(err)
	}
	for _, e := range dirs {
		if e.IsDir() {
			err = RangeDir(dir+PathSeparator+e.Name(), callback)
			if err != nil {
				errs.Append(err)
			}
		}
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}
