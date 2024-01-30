package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/tiga/utils/net/http/request/binding"
	"github.com/valyala/fasthttp"
)

type yamlBinding struct{}

func (yamlBinding) Name() string {
	return "yaml"
}

func (y yamlBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	return binding.DecodeYaml(req.Body(), obj)
}

func (y yamlBinding) FiberBind(ctx fiber.Ctx, obj interface{}) error {
	return binding.DecodeYaml(ctx.Body(), obj)
}
