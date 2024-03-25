package conf_center

import (
	"github.com/hopeio/tiga/initialize/conf_center/http"
	"github.com/hopeio/tiga/initialize/conf_center/local"
	"github.com/hopeio/tiga/initialize/conf_center/nacos"
)

type ConfigType string

const (
	ConfigTypeLocal = "local"
	ConfigTypeNacos = "nacos"
	ConfigTypeHttp  = "http"
)

type ConfigCenter interface {
	HandleConfig(func([]byte)) error
}

type ConfigCenterConfig struct {
	// 配置格式
	Format string `flag:"name:format;default:toml;usage:配置格式" comment:"toml,json,yaml,yml"`
	// 配置类型
	ConfigType string `flag:"name:conf_type;default:local;usage:配置类型"`
	// 是否监听配置文件
	Watch bool `flag:"name:watch;short:w;default:false;usage:是否监听配置文件" env:"WATCH"`
	Nacos *nacos.Nacos
	Local *local.Local
	Http  *http.Config
	/*	Etcd   *etcd.Etcd
		Apollo *apollo.Apollo*/
}

func (c *ConfigCenterConfig) ConfigCenter(debug bool) ConfigCenter {
	if c.Format == "" {
		c.Format = "toml"
	}
	if c.ConfigType == ConfigTypeHttp && c.Http != nil {
		return c.Http
	}

	if c.ConfigType == ConfigTypeNacos && c.Nacos != nil {
		return c.Nacos
	}
	/*	if c.Etcd != nil && ccec.EtcdKey != "" {
		c.Etcd.Key = ccec.EtcdKey
		c.Etcd.Watch = c.Watch
		return c.Etcd
	}*/
	if c.ConfigType == ConfigTypeLocal && c.Local != nil {
		c.Local.Debug = debug
		c.Local.AutoReload = c.Watch
		return c.Local
	}
	return c.Local
}
