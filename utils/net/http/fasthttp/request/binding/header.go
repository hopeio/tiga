package binding

import (
	"github.com/hopeio/tiga/utils/net/http/request/binding"
	"github.com/valyala/fasthttp"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {

	if err := binding.MappingByPtr(obj, (*reqSource)(&req.Header), binding.Tag); err != nil {
		return err
	}

	return binding.Validate(obj)
}
