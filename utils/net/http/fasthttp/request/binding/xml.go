package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/tiga/utils/net/http/request/binding"
	"github.com/valyala/fasthttp"
)

type xmlBinding struct{}

func (xmlBinding) Name() string {
	return "xml"
}

func (x xmlBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	return binding.DecodeXml(req.Body(), obj)
}

func (x xmlBinding) FiberBind(ctx fiber.Ctx, obj interface{}) error {
	return binding.DecodeXml(ctx.Body(), obj)
}
