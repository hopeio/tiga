package initialize

import (
	"fmt"
	"github.com/hopeio/tiga/initialize/conf_center"
	ilocal "github.com/hopeio/tiga/initialize/conf_center/local"
	"github.com/hopeio/tiga/utils/configor/local"
	"github.com/hopeio/tiga/utils/errors/multierr"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/hopeio/tiga/utils/log"
)

// 约定大于配置
var (
	GlobalConfig = &globalConfig{
		BasicConfig: BasicConfig{Module: "tiga-app", Env: DEVELOPMENT, ConfUrl: "./config.toml"},
		confMap:     map[string]interface{}{},
		RWMutex:     sync.RWMutex{},
	}
)

type Env string

const (
	DEVELOPMENT = "dev"
	TEST        = "test"
	PRODUCT     = "prod"
	InitKey     = "initialize"
)

// ConfigCenterConfig
// zh:配置中心的配置
// en:
type ConfigCenterConfig struct {
	NoInject []string
	conf_center.ConfigCenterConfig
}

// BasicConfig
// zh: 基本配置，包含模块名
type BasicConfig struct {
	// 模块名
	Module string `flag:"name:mod;short:m;default:;usage:模块名" env:"name:MODULE"`
	// environment
	Env string `flag:"name:env;short:e;default:dev;usage:环境" env:"name:ENV"`
	// 配置文件路径
	ConfUrl string `flag:"name:conf;short:c;default:config.toml;usage:配置文件路径,默认./config.toml或./config/config.toml" env:"name:CONFIG"`
	// 代理, socks5://localhost:1080
	Proxy string `flag:"name:proxy;short:p;default:'socks5://localhost:1080';usage:代理" env:"HTTP_PROXY"`
}

// FileConfig
// 配置文件映射结构体,每个启动都有一个必要的配置文件,用于初始化基本配置及配置中心配置,
/*```toml
Module = "user"

[dev]
configType = "local"
Watch  = true
NoInject = ["Apollo","Etcd", "Es"]

[dev.local]
Debug = true
ConfigPath = "local.toml"
ReloadType = "fsnotify"

[dev.http]
Interval = 100
Url = "http://localhost:6666/local.toml"

[dev.nacos]
DataId = "pro"
Group = "DEFAULT_GROUP"

[[dev.nacos.ServerConfigs]]
Scheme = "http"
IpAddr = "nacos"
Port = 9000
GrpcPort = 10000

[dev.nacos.ClientConfig]
NamespaceId = ""
username = "nacos"
password = "nacos"
LogLevel = "debug"
```*/
type FileConfig struct {
	// 模块名
	Module          string `flag:"name:mod;short:m;default:;usage:模块名" env:"name:MODULE"`
	Dev, Test, Prod *ConfigCenterConfig
}

// globalConfig
// 全局配置
type globalConfig struct {
	BasicConfig
	ConfigCenterConfig ConfigCenterConfig
	confMap            map[string]any
	conf               NeedInit
	dao                Dao
	//closes     []any
	deferCalls  []func()
	initialized bool
	sync.RWMutex
}

func Start(conf Config, dao Dao, notinit ...string) func(deferCalls ...func()) {
	//逃逸到堆上了
	GlobalConfig.setConfDao(conf, dao)
	GlobalConfig.LoadConfig(notinit...)
	GlobalConfig.initialized = true
	return func(deferCalls ...func()) {
		for _, f := range deferCalls {
			f()
		}
		for _, f := range GlobalConfig.deferCalls {
			f()
		}
	}
}

func (gc *globalConfig) LoadConfig(notinject ...string) {
	onceConfig := FileConfig{}
	if _, err := os.Stat(gc.ConfUrl); os.IsNotExist(err) {
		log.Fatalf("配置错误: 请确保可执行文件和配置文件在同一目录下或在config目录下或指定配置文件")
	}
	err := local.Load(&onceConfig, gc.ConfUrl)
	if err != nil {
		log.Fatalf("配置错误: %v", err)
	}
	fmt.Printf("Load config from: %s\n", gc.ConfUrl)

	gc.Module = onceConfig.Module

	value := reflect.ValueOf(&onceConfig).Elem()
	typ := reflect.TypeOf(&onceConfig).Elem()

	tmpTyp := reflect.TypeOf(&ConfigCenterConfig{})
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Type == tmpTyp && strings.ToUpper(typ.Field(i).Name) == strings.ToUpper(gc.Env) {
			/*tmpConfig = value.Field(i).Interface().(*nacos.Config)
			//真·深度复制
			data,_:=json.Marshal(tmpConfig)
			if err:=json.Unmarshal(data,gc.ConfigCenterConfig);err!=nil{
				log.Fatal(err)
			}*/
			fieldValue := value.Field(i)
			if fieldValue.IsValid() && !fieldValue.IsNil() {
				gc.ConfigCenterConfig = *fieldValue.Interface().(*ConfigCenterConfig)
			} else {
				// 单配置文件
				gc.ConfigCenterConfig = ConfigCenterConfig{
					ConfigCenterConfig: conf_center.ConfigCenterConfig{
						ConfigType: conf_center.ConfigTypeLocal,
						Local: &ilocal.Local{
							Config:     local.Config{},
							ReloadType: ilocal.ReloadTypeFsNotify,
							ConfigPath: gc.ConfUrl,
						},
					},
				}
			}
			break
		}
	}

	for i := range gc.ConfigCenterConfig.NoInject {
		gc.ConfigCenterConfig.NoInject[i] = strings.ToUpper(gc.ConfigCenterConfig.NoInject[i])
	}
	for i := range notinject {
		gc.ConfigCenterConfig.NoInject = append(gc.ConfigCenterConfig.NoInject, strings.ToUpper(notinject[i]))
	}

	gc.applyEnvConfig()
	gc.applyFlagConfig()
	cfgcenter := gc.ConfigCenterConfig.ConfigCenter(gc.Env != PRODUCT)

	err = cfgcenter.HandleConfig(gc.UnmarshalAndSet)
	if err != nil {
		log.Fatalf("配置错误: %v", err)
	}

}

func (gc *globalConfig) setConfDao(conf Config, dao Dao) {
	gc.conf = conf
	gc.dao = dao
	gc.deferCalls = []func(){
		func() { closeDao(dao) },
		func() { log.Sync() },
	}
}

func (gc *globalConfig) RegisterDeferFunc(deferf ...func()) {
	gc.deferCalls = append(gc.deferCalls, deferf...)
}

func (gc *globalConfig) Config() Config {
	return gc.conf
}

func (gc *globalConfig) closeDao() {
	if !gc.initialized {
		return
	}
	err := closeDao(gc.dao)
	if err != nil {
		log.Error(err)
	}
}

func closeDao(dao interface{}) error {
	if dao == nil {
		return nil
	}
	if err, ok := closeDaoHelper(dao); ok {
		return err
	}
	var err multierr.MultiError
	daoValue := reflect.ValueOf(dao).Elem()
	for i := 0; i < daoValue.NumField(); i++ {
		/*	closer := daoValue.Field(i).MethodByName("Close")
			if closer.IsValid() {
				closer.Call(nil)
			}*/
		fieldV := daoValue.Field(i)
		if fieldV.Type().Kind() == reflect.Struct {
			fieldV = daoValue.Field(i).Addr()
		}
		if !fieldV.IsValid() || fieldV.IsNil() {
			continue
		}
		field := fieldV.Interface()
		if err1, ok := closeDaoHelper(field); ok && err1 != nil {
			err.Append(err1)
		}
	}
	if err.HasErrors() {
		return &err
	}
	return nil
}

func closeDaoHelper(dao interface{}) (error, bool) {
	if dao == nil {
		return nil, true
	}
	if closer, ok := dao.(DaoFieldCloserWithError); ok {
		return closer.Close(), true
	}
	if closer, ok := dao.(DaoFieldCloser); ok {
		closer.Close()
		return nil, true
	}
	return nil, false
}

func GetConfig[T any]() *T {
	iconf := GlobalConfig.Config()
	value := reflect.ValueOf(iconf).Elem()
	for i := 0; i < value.NumField(); i++ {
		if conf, ok := value.Field(i).Interface().(T); ok {
			return &conf
		}
	}
	return new(T)
}
