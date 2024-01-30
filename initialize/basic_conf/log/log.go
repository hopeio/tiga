package log

import (
	initialize2 "github.com/hopeio/tiga/initialize"
	"github.com/hopeio/tiga/utils/log"
)

type LogConfig log.Config

func (conf *LogConfig) Init() {
	logConf := (*log.Config)(conf)
	logConf.Development = initialize2.GlobalConfig.Env != initialize2.PRODUCT
	logConf.ModuleName = initialize2.GlobalConfig.Module
	log.SetDefaultLogger(logConf)
}

func (conf *LogConfig) Build() *log.Logger {
	return (*log.Config)(conf).NewLogger()
}

type Logger struct {
	*log.Logger `init:"entity"`
	Conf        LogConfig `init:"config"`
}

func (l *Logger) Config() any {
	return &l.Conf
}

func (l *Logger) SetEntity() {
	l.Logger = l.Conf.Build()
}

func (l *Logger) Close() error {
	if l.Logger == nil {
		return nil
	}
	return l.Logger.Sync()
}
