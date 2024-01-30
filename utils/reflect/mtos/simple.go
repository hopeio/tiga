package mtos

import (
	reflecti "github.com/hopeio/tiga/utils/reflect"
	"reflect"
)

// TODO: 由string map 设置结构体，一般用于配置
func SetByMap(m map[string]string, dst any) {
	//dstType := reflect.TypeOf(dst).Elem()

}

// SetStructByMap 由map注入struct
func SetStructByMap(dst any, mapData map[string]any) error {
	structData := reflect.ValueOf(dst).Elem()
	for key, value := range mapData {
		if err := reflecti.SetField(structData, key, value); err != nil {
			return err
		}
	}
	return nil
}
