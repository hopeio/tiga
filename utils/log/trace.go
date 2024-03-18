package log

import (
	"context"
	contexti "github.com/hopeio/tiga/utils/context"
	"go.uber.org/zap"
)

func TraceIdField(ctx context.Context) zap.Field {
	return zap.String(TraceId, contexti.TraceId(ctx))
}
