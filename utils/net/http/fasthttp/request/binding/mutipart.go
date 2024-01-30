package binding

import (
	"github.com/hopeio/tiga/utils/net/http/request/binding"
	"github.com/valyala/fasthttp"
	"reflect"
)

type multipartFasthttpRequest fasthttp.Request

var _ binding.Setter = (*multipartFasthttpRequest)(nil)

// TrySet tries to set a value by the multipart request with the binding a form file
func (r *multipartFasthttpRequest) TrySet(value reflect.Value, field reflect.StructField, key string, opt binding.SetOptions) (isSetted bool, err error) {
	req := (*fasthttp.Request)(r)
	form, err := req.MultipartForm()
	if err != nil {
		return false, err
	}
	if files := form.File[key]; len(files) != 0 {
		return binding.SetByMultipartFormFile(value, field, files)
	}

	return binding.SetByKV(value, field, binding.FormSource(form.Value), key, opt)
}
