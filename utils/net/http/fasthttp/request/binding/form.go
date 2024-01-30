package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/tiga/utils/net/http/request/binding"
	"github.com/valyala/fasthttp"
)

type formBinding struct{}
type formPostBinding struct{}
type formMultipartBinding struct{}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (formMultipartBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	if err := binding.MappingByPtr(obj, (*multipartFasthttpRequest)(req), binding.Tag); err != nil {
		return err
	}

	return binding.Validate(obj)
}

func (formMultipartBinding) FiberBind(ctx fiber.Ctx, obj interface{}) error {
	if err := binding.MappingByPtr(obj, (*multipartFasthttpRequest)(ctx.Request()), binding.Tag); err != nil {
		return err
	}

	return binding.Validate(obj)
}

func (formBinding) Name() string {
	return "form"
}
func (formBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	if err := binding.MapForm(obj, (*argsSource)(req.PostArgs())); err != nil {
		return err
	}
	return binding.Validate(obj)
}

func (formBinding) FiberBind(ctx fiber.Ctx, obj interface{}) error {
	if err := binding.MapForm(obj, (*argsSource)(ctx.Request().PostArgs())); err != nil {
		return err
	}
	return binding.Validate(obj)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}
func (formPostBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	if err := binding.MapForm(obj, (*argsSource)(req.PostArgs())); err != nil {
		return err
	}
	return binding.Validate(obj)
}

func (formPostBinding) FiberBind(ctx fiber.Ctx, obj interface{}) error {
	if err := binding.MapForm(obj, (*argsSource)(ctx.Request().PostArgs())); err != nil {
		return err
	}
	return binding.Validate(obj)
}
