package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/tiga/utils/net/http/request/binding"
	stringsi "github.com/hopeio/tiga/utils/strings"
	"github.com/valyala/fasthttp"
	"reflect"
)

type argsSource fasthttp.Args

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form *argsSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt binding.SetOptions) (isSetted bool, err error) {
	return binding.SetByKV(value, field, form, tagValue, opt)
}

func (form *argsSource) Peek(key string) ([]string, bool) {
	v := stringsi.BytesToString((*fasthttp.Args)(form).Peek(key))
	return []string{v}, v != ""
}

type ctxSource fiber.Ctx

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form *ctxSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt binding.SetOptions) (isSetted bool, err error) {
	return binding.SetByKV(value, field, form, tagValue, opt)
}

func (form *ctxSource) Peek(key string) ([]string, bool) {
	v := (fiber.Ctx)(form).Params(key)
	return []string{v}, v != ""
}

type reqSource fasthttp.RequestHeader

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form *reqSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt binding.SetOptions) (isSetted bool, err error) {
	return binding.SetByKV(value, field, form, tagValue, opt)
}

func (form *reqSource) Peek(key string) ([]string, bool) {
	v := stringsi.BytesToString((*fasthttp.RequestHeader)(form).Peek(key))
	return []string{v}, v != ""
}
