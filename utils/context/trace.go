package contexti

import (
	"context"
)

func TraceId(ctx context.Context) string {
	if traceId, ok := ctx.Value(traceIdKey{}).(string); ok {
		return traceId
	}
	return "unknown"
}

type traceIdKey struct{}

func SetTranceId(traceId string) context.Context {
	return context.WithValue(context.Background(), traceIdKey{}, traceId)
}
