package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/tiga/utils/net/http/request/binding"
	"github.com/valyala/fasthttp"
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}

func (b protobufBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	return binding.DecodeProtobuf(req.Body(), obj)
}

func (b protobufBinding) FiberBind(ctx fiber.Ctx, obj interface{}) error {
	return binding.DecodeProtobuf(ctx.Body(), obj)
}
