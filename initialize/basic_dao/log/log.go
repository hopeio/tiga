package log

import (
	logconf "github.com/hopeio/tiga/initialize/basic_conf/log"
	"github.com/hopeio/tiga/utils/log"
)

type Logger struct {
	*log.Logger `init:"entity"`
	Conf        logconf.LogConfig `init:"config"`
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
