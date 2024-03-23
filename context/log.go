package contexti

import (
	"github.com/hopeio/tiga/protobuf/errorcode"
	"github.com/hopeio/tiga/utils/log"
	"go.uber.org/zap"
)

func (c *RequestContext[REQ, RES]) Error(args ...interface{}) {
	args = append(args, zap.String(log.TraceId, c.TraceID))
	log.Default.Error(args...)
}

func (c *RequestContext[REQ, RES]) HandleError(err error) {
	if err != nil {
		log.Default.Error(err.Error(), zap.String(log.TraceId, c.TraceID))
	}
}

func (c *RequestContext[REQ, RES]) ErrorLog(err, originErr error, funcName string) error {
	// caller 用原始logger skip刚好
	log.GetSkipLogger(1).Error(originErr.Error(), zap.String(log.TraceId, c.TraceID), zap.Int(log.Type, errorcode.Code(err)), zap.String(log.Position, funcName))
	return err
}
