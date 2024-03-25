package log

import (
	initialize2 "github.com/hopeio/tiga/initialize"
	"github.com/hopeio/tiga/utils/log"
)

type LogConfig log.Config

func (conf *LogConfig) Init() {
	logConf := (*log.Config)(conf)
	logConf.Development = initialize2.GlobalConfig.Debug
	logConf.ModuleName = initialize2.GlobalConfig.Module
	log.SetDefaultLogger(logConf)
}

func (conf *LogConfig) Build() *log.Logger {
	return (*log.Config)(conf).NewLogger()
}
