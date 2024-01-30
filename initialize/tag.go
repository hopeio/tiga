package initialize

import (
	reflecti "github.com/hopeio/tiga/utils/reflect"
	"github.com/hopeio/tiga/utils/reflect/structtag"
	"github.com/spf13/pflag"
	"os"
	"reflect"
	"strings"
)

// example:
/*
type Dao struct {
	DB mysql.DB `init:"config:MysqlTest"`
}

type Config struct {
	Env string `init:"tag:env"`
}

*/

const (
	initTagName = "init"
	exprTagName = "expr"

	metaTagName = "meta"
)

type DaoTagSettings struct {
	NotInject  bool   `meta:"not_inject"`
	ConfigName string `meta:"config"`
}

type ConfigTagSettings struct {
	Flag string
	Env  string
}

func ParseDaoTagSettings(str string) *DaoTagSettings {
	if str == "" {
		return &DaoTagSettings{}
	}
	var settings DaoTagSettings
	ParseTagSetting(str, ";", &settings)
	return &settings
}

// Deprecated: 使用独立的tag
func ParseConfigTagSettings(str string) *ConfigTagSettings {
	if str == "" {
		return &ConfigTagSettings{}
	}
	var settings ConfigTagSettings
	ParseTagSetting(str, ";", &settings)
	return &settings
}

// ParseTagSetting default sep ;
func ParseTagSetting(str string, sep string, settings any) {
	tagSettings := structtag.ParseTagSetting(str, sep)
	settingsValue := reflect.ValueOf(settings).Elem()
	settingsType := reflect.TypeOf(settings).Elem()
	for i := 0; i < settingsValue.NumField(); i++ {
		if flagtag, ok := tagSettings[strings.ToUpper(settingsType.Field(i).Tag.Get(metaTagName))]; ok {
			reflecti.SetFieldByString(flagtag, settingsValue.Field(i))
		}
	}
}

// Deprecated: 使用独立的tag
func Unmarshal(v reflect.Value) error {
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Ptr:
			Unmarshal(v.Field(i).Elem())
		case reflect.Struct:
			Unmarshal(v.Field(i))
		}
		tag := typ.Field(i).Tag.Get(initTagName)
		if tag != "" {
			settings := ParseConfigTagSettings(tag)
			switch field.Kind() {
			case reflect.String:
				field.Set(reflect.ValueOf(os.Getenv(settings.Env)))
				pflag.StringVarP(field.Addr().Interface().(*string), "", settings.Flag, field.Interface().(string), "")
			case reflect.Int:
				pflag.IntVarP(field.Addr().Interface().(*int), "", settings.Flag, field.Interface().(int), "")
			case reflect.Bool:
				pflag.BoolVarP(field.Addr().Interface().(*bool), "", settings.Flag, field.Interface().(bool), "")
			}

		}
	}
	return nil
}
