package viper

import (
	initialize2 "github.com/hopeio/tiga/initialize"
	"github.com/hopeio/tiga/utils/log"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Config struct {
	Remote     bool
	Watch      bool
	Provider   string
	Endpoint   string
	Path       string
	ConfigType string
}

func (conf *Config) Init() {
	conf.build(viper.GetViper())
}

func (conf *Config) Build() *viper.Viper {

	if conf.ConfigType == "" {
		conf.ConfigType = "toml"
	}
	var runtimeViper = viper.New()
	conf.build(runtimeViper)
	return runtimeViper
}

func (conf *Config) build(runtimeViper *viper.Viper) {

	runtimeViper.SetConfigType(conf.ConfigType) // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "Env", "dotenv"
	if conf.Remote {
		err := runtimeViper.AddRemoteProvider(conf.Provider, conf.Endpoint, initialize2.InitKey)
		if err != nil {
			log.Fatal(err)
		}
		// read from remote Config the first time.
		err = runtimeViper.ReadRemoteConfig()
		if err != nil {
			log.Error(err)
		}
		if conf.Watch {
			err = runtimeViper.WatchRemoteConfig()
			if err != nil {
				log.Fatal(err)
			}
		}

	} else {
		runtimeViper.AddConfigPath(conf.Path)
		err := runtimeViper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
		if conf.Watch {
			runtimeViper.WatchConfig()
		}
	}

	// open a goroutine to watch remote changes forever
	//这段实现不够优雅
	/*	go func() {
		for {
			time.Sleep(time.Second * 5) // delay after each request

			// currently, only tested with etcd support
			err := runtime_viper.WatchRemoteConfig()
			if err != nil {
				log.Errorf("unable to read remote Config: %v", err)
				continue
			}
			vconf :=runtime_viper.AllSettings()
			log.Debug(vconf)
			// unmarshal new Config into our runtime Config struct. you can also use channel
			// to implement a signal to notify the system of the changes
			runtime_viper.Unmarshal(cCopy)
			refresh(cCopy, dCopy)
			log.Debug(cCopy)
		}
	}()*/
}

type Viper struct {
	*viper.Viper
	Conf Config
}

func (v *Viper) Config() any {
	return &v.Conf
}

func (v *Viper) SetEntity() {
	v.Viper = v.Conf.Build()
}

func (v *Viper) Close() error {
	return nil
}
