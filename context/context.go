package contexti

import (
	"context"
	"github.com/google/uuid"
	"github.com/hopeio/tiga/utils/net/http/request"
	timei "github.com/hopeio/tiga/utils/time"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"strings"
	"sync"
	"time"
)

func GetPool[REQ any]() sync.Pool {
	return sync.Pool{New: func() any {
		return new(RequestContext[REQ])
	}}
}

type RequestContext[REQ any] struct {
	context.Context
	TraceID string

	Token  string
	AuthID string
	AuthInfo
	AuthInfoRaw string

	*DeviceInfo

	request.RequestAt
	Request *REQ
	grpc.ServerTransportStream

	Internal string
	Values   map[string]any
}

func (c *RequestContext[REQ]) StartSpan(name string, o ...trace.StartOption) (*RequestContext[REQ], *trace.Span) {
	ctx, span := trace.StartSpan(c.Context, name, append(o, trace.WithSampler(trace.AlwaysSample()),
		trace.WithSpanKind(trace.SpanKindServer))...)
	c.Context = ctx
	if c.TraceID == "" {
		c.TraceID = span.SpanContext().TraceID.String()
	}
	return c, span
}

func methodFamily(m string) string {
	m = strings.TrimPrefix(m, "/") // remove leading slash
	if i := strings.Index(m, "/"); i >= 0 {
		m = m[:i] // remove everything from second slash
	}
	return m
}

type ctxKey struct{}

func (c *RequestContext[REQ]) ContextWrapper() context.Context {
	return context.WithValue(context.Background(), ctxKey{}, c)
}

func ContextFromContext[REQ any](ctx context.Context) *RequestContext[REQ] {
	ctxi := ctx.Value(ctxKey{})
	c, ok := ctxi.(*RequestContext[REQ])
	if !ok {
		c = NewContext[REQ](ctx)
	}
	if c.ServerTransportStream == nil {
		c.ServerTransportStream = grpc.ServerTransportStreamFromContext(ctx)
	}
	return c
}

func (c *RequestContext[REQ]) WithContext(ctx context.Context) {
	c.Context = ctx
}

func NewContext[REQ any](ctx context.Context) *RequestContext[REQ] {
	span := trace.FromContext(ctx)
	now := time.Now()
	traceId := span.SpanContext().TraceID.String()
	if traceId == "" {
		traceId = uuid.New().String()
	}
	return &RequestContext[REQ]{
		Context: ctx,
		TraceID: traceId,
		RequestAt: request.RequestAt{
			Time:       now,
			TimeStamp:  now.Unix(),
			TimeString: now.Format(timei.TimeFormat),
		},
		ServerTransportStream: grpc.ServerTransportStreamFromContext(ctx),
	}
}

func (c *RequestContext[REQ]) reset(ctx context.Context) *RequestContext[REQ] {
	span := trace.FromContext(ctx)
	now := time.Now()
	traceId := span.SpanContext().TraceID.String()
	if traceId == "" {
		traceId = uuid.New().String()
	}
	c.Context = ctx
	c.RequestAt.Time = now
	c.RequestAt.TimeString = now.Format(timei.TimeFormat)
	c.RequestAt.TimeStamp = now.Unix()
	return c
}

func (c *RequestContext[REQ]) Method() string {
	if c.ServerTransportStream != nil {
		return c.ServerTransportStream.Method()
	}
	return ""
}
