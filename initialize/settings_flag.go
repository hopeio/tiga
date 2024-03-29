package initialize

import (
	"flag"
	reflecti "github.com/hopeio/tiga/utils/reflect"
	"github.com/spf13/pflag"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

const flagTagName = "flag"

// TODO: 优先级高于其他Config,覆盖环境变量及配置中心的配置
// example
/*type FlagConfig struct {
	// environment
	Env string `flag:"name:env;short:e;default:dev;usage:环境"`
	// 配置文件路径
	ConfUrl string `flag:"name:conf;short:c;default:config.toml;usage:配置文件路径,默认./config.toml或./config/config.toml"`
}*/

type FlagTagSettings struct {
	Name    string `meta:"name"`
	Short   string `meta:"short"`
	Default string `meta:"default"`
	Usage   string `meta:"usage"`
}

func init() {
	envinit()
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	commandLine := newCommandLine()
	injectFlagConfig(commandLine, reflect.ValueOf(&GlobalConfig.BasicConfig).Elem())
	Parse(commandLine)

	if GlobalConfig.Proxy != "" {
		proxyURL, _ := url.Parse(GlobalConfig.Proxy)
		http.DefaultClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}
}

type anyValue reflect.Value

func (a anyValue) String() string {
	return reflecti.String(reflect.Value(a))
}

func (a anyValue) Type() string {
	return reflect.Value(a).Kind().String()
}

func (a anyValue) Set(v string) error {
	return reflecti.SetFieldByString(v, reflect.Value(a))
}

func injectFlagConfig(commandLine *pflag.FlagSet, fcValue reflect.Value) {
	if !fcValue.IsValid() || fcValue.IsZero() {
		return
	}
	fcTyp := fcValue.Type()

	for i := 0; i < fcTyp.NumField(); i++ {
		fieldType := fcTyp.Field(i)
		if !fieldType.IsExported() {
			continue
		}
		flagTag := fieldType.Tag.Get(flagTagName)
		fieldValue := fcValue.Field(i)
		kind := fieldValue.Kind()
		if kind == reflect.Pointer {
			if !fieldValue.IsValid() || fieldValue.IsNil() {
				fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
			}
			injectFlagConfig(commandLine, fieldValue.Elem())
		}
		if kind == reflect.Struct {
			injectFlagConfig(commandLine, fieldValue)
		}
		if flagTag != "" {
			var flagTagSettings FlagTagSettings
			ParseTagSetting(flagTag, ";", &flagTagSettings)
			flag := commandLine.VarPF(anyValue(fieldValue), flagTagSettings.Name, flagTagSettings.Short, flagTagSettings.Usage)
			if kind == reflect.Bool {
				flag.NoOptDefVal = "true"
			}
		}
	}
}

func (gc *globalConfig) applyFlagConfig() {
	commandLine := newCommandLine()
	fcValue := reflect.ValueOf(&gc.BasicConfig).Elem()
	injectFlagConfig(commandLine, fcValue)
	fcValue = reflect.ValueOf(&gc.ConfigCenterConfig).Elem()
	injectFlagConfig(commandLine, fcValue)
	Parse(commandLine)
}

func newCommandLine() *pflag.FlagSet {
	commandLine := pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	commandLine.ParseErrorsWhitelist.UnknownFlags = true
	return commandLine
}

func Parse(commandLine *pflag.FlagSet) {
	commandLine.Parse(os.Args[1:])
}
