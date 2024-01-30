package binding

import (
	"net/http"
	"net/textproto"
	"reflect"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req *http.Request, obj interface{}) error {

	if err := mapHeader(obj, req.Header); err != nil {
		return err
	}

	return Validate(obj)
}

func mapHeader(ptr interface{}, h map[string][]string) error {
	return MappingByPtr(ptr, headerSource(h), Tag)
}

type headerSource map[string][]string

var _ Setter = headerSource(nil)

func (hs headerSource) Peek(key string) ([]string, bool) {
	v, ok := hs[key]
	return v, ok
}

func (hs headerSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt SetOptions) (isSetted bool, err error) {
	return SetByKV(value, field, hs, textproto.CanonicalMIMEHeaderKey(tagValue), opt)
}
