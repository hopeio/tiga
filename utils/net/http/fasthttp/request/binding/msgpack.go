//go:build !nomsgpack
// +build !nomsgpack

package binding

import (
	"bytes"
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/tiga/utils/net/http/request/binding"
	"github.com/valyala/fasthttp"
)

type msgpackBinding struct{}

func (msgpackBinding) Name() string {
	return "msgpack"
}

func (m msgpackBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	return m.BindBody(req.Body(), obj)
}

func (m msgpackBinding) FiberBind(ctx fiber.Ctx, obj interface{}) error {
	return m.BindBody(ctx.Body(), obj)
}

func (msgpackBinding) BindBody(body []byte, obj interface{}) error {
	return binding.DecodeMsgPack(bytes.NewReader(body), obj)
}
