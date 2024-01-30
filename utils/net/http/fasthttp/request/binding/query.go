package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/tiga/utils/net/http/request/binding"
	"github.com/valyala/fasthttp"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	values := req.URI().QueryArgs()
	if err := binding.MapForm(obj, (*argsSource)(values)); err != nil {
		return err
	}
	return binding.Validate(obj)
}

func (queryBinding) FiberBind(ctx fiber.Ctx, obj interface{}) error {
	values := ctx.Request().URI().QueryArgs()
	if err := binding.MapForm(obj, (*argsSource)(values)); err != nil {
		return err
	}
	return binding.Validate(obj)
}
