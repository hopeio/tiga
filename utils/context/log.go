package contexti

import (
	"github.com/hopeio/lemon/protobuf/errorcode"
	"github.com/hopeio/lemon/utils/log"
	"go.uber.org/zap"
)

func (c *RequestContext[REQ]) Error(args ...interface{}) {
	args = append(args, zap.String(log.TraceId, c.TraceID))
	log.Default.Error(args...)
}

func (c *RequestContext[REQ]) HandleError(err error) {
	if err != nil {
		log.Default.Error(err.Error(), zap.String(log.TraceId, c.TraceID))
	}
}

func (c *RequestContext[REQ]) ErrorLog(err, originErr error, funcName string) error {
	// caller 用原始logger skip刚好
	log.GetSkipLogger(1).Error(originErr.Error(), zap.String(log.TraceId, c.TraceID), zap.Int(log.Type, errorcode.Code(err)), zap.String(log.Position, funcName))
	return err
}
