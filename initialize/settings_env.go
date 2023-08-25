package initialize

import (
	reflecti "github.com/hopeio/lemon/utils/reflect"
	"net/http"
	"os"
	"reflect"
)

const envTagName = "env"

// example
type EnvConfig struct {
	Proxy string `env:"name:HTTP_PROXY"`
}

type EnvTagSettings struct {
	Name  string `meta:"name"`
	Usage string `meta:"usage"`
}

func envinit() {

	if env, ok := os.LookupEnv("ENV"); ok {
		GlobalConfig.Env = env
	}
	if conf, ok := os.LookupEnv("CONFIG"); ok {
		GlobalConfig.ConfUrl = conf
	}

	http.DefaultClient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}
}

func injectEnvConfig(ecValue reflect.Value) {
	if !ecValue.IsValid() || ecValue.IsZero() {
		return
	}
	ecTyp := ecValue.Type()
	for i := 0; i < ecTyp.NumField(); i++ {
		fieldType := ecTyp.Field(i)
		if !fieldType.IsExported() {
			continue
		}
		fieldValue := ecValue.Field(i)

		kind := fieldValue.Kind()
		if kind == reflect.Pointer {
			if !fieldValue.IsValid() || fieldValue.IsNil() {
				fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
			}
			injectEnvConfig(fieldValue.Elem())
		}
		if kind == reflect.Struct {
			injectEnvConfig(fieldValue)
		}
		flagTag := fieldType.Tag.Get(envTagName)
		if flagTag != "" {
			var envTagSettings EnvTagSettings
			ParseTagSetting(flagTag, ";", &envTagSettings)
			if value, ok := os.LookupEnv(envTagSettings.Name); ok {
				reflecti.SetFieldByString(value, ecValue.Field(i))
			}
		}
	}
}

func (gc *globalConfig) applyEnvConfig() {
	fcValue := reflect.ValueOf(&gc.BasicConfig).Elem()
	injectEnvConfig(fcValue)
	fcValue = reflect.ValueOf(&gc.ConfigCenterConfig).Elem()
	injectEnvConfig(fcValue)
}
