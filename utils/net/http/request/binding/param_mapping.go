package binding

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type paramSource gin.Params

var _ Setter = paramSource(nil)

func (param paramSource) Peek(key string) ([]string, bool) {
	for i := range param {
		if param[i].Key == key {
			return []string{param[i].Value}, true
		}
	}
	return nil, false
}

// TrySet tries to set a value by request's form source (like map[string][]string)
func (param paramSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt SetOptions) (isSetted bool, err error) {
	return SetByKV(value, field, param, tagValue, opt)
}
