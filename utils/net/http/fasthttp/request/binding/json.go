package binding

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/tiga/utils/net/http/request/binding"
	"github.com/valyala/fasthttp"
)

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (j jsonBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	body := req.Body()
	if req == nil || body == nil {
		return fmt.Errorf("invalid request")
	}
	return binding.DecodeJson(body, obj)
}

func (j jsonBinding) FiberBind(ctx fiber.Ctx, obj interface{}) error {
	body := ctx.Request().Body()
	if body == nil {
		return fmt.Errorf("invalid request")
	}
	return binding.DecodeJson(body, obj)
}
