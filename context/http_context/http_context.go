package http_context

import (
	"context"
	"encoding/base64"
	"github.com/hopeio/tiga/context"
	httpi "github.com/hopeio/tiga/utils/net/http"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
	"google.golang.org/grpc/metadata"
	"net/http"
	"runtime"
)

type Context = contexti.RequestContext[http.Request]

func ContextFromContext(ctx context.Context) *Context {
	return contexti.ContextFromContext[http.Request](ctx)
}

func ContextFromRequest(r *http.Request, tracing bool) (*Context, *trace.Span) {
	var span *trace.Span
	ctx := context.Background()
	if r != nil {
		ctx = r.Context()
		if tracing {
			// go.opencensus.io/trace 完全包含了golang.org/x/net/trace 的功能
			// grpc内置配合,看了源码并没有启用，根本没调用
			// 系统trace只能追踪单个请求，且只记录时间及是否完成，只能/debug/requests看
			/*			t = gtrace.New(methodFamily(r.RequestURI), r.RequestURI)
						ctx = gtrace.NewContext(ctx, t)
			*/

			// 直接从远程读取Trace信息，Trace是否为空交给propagation包判断
			traceString := r.Header.Get(httpi.GrpcTraceBin)
			if traceString == "" {
				traceString = r.Header.Get(httpi.HeaderTraceBin)
			}
			var traceBin []byte
			if len(traceString)%4 == 0 {
				// Input was padded, or padding was not necessary.
				traceBin, _ = base64.StdEncoding.DecodeString(traceString)
			}
			traceBin, _ = base64.RawStdEncoding.DecodeString(traceString)

			if parent, ok := propagation.FromBinary(traceBin); ok {
				ctx, span = trace.StartSpanWithRemoteParent(ctx, r.RequestURI,
					parent, trace.WithSampler(trace.AlwaysSample()),
					trace.WithSpanKind(trace.SpanKindServer))
			} else {
				ctx, span = trace.StartSpan(ctx, r.RequestURI,
					trace.WithSampler(trace.AlwaysSample()),
					trace.WithSpanKind(trace.SpanKindServer))
			}
		}
	} else {
		if tracing {
			pc, _, _, _ := runtime.Caller(2)
			f := runtime.FuncForPC(pc)

			/*			t = gtrace.New(file, fmt.Sprintf("%s:%d", file, line))
						ctx = gtrace.NewContext(ctx, t)*/

			ctx, span = trace.StartSpan(ctx, f.Name(),
				trace.WithSampler(trace.AlwaysSample()),
				trace.WithSpanKind(trace.SpanKindServer))
		}
	}

	ctxi := contexti.NewContext[http.Request](ctx)
	setWithHttpReq(ctxi, r)
	return ctxi, span
}

func setWithHttpReq(c *contexti.RequestContext[http.Request], r *http.Request) {
	if r == nil {
		return
	}
	c.Request = r
	c.DeviceInfo = DeviceFromHeader(r.Header)
	c.Internal = r.Header.Get(httpi.GrpcInternal)
	c.Token = httpi.GetToken(r)
}

func DeviceFromHeader(r http.Header) *contexti.DeviceInfo {
	return contexti.Device(r.Get(httpi.HeaderDeviceInfo),
		r.Get(httpi.HeaderArea), r.Get(httpi.HeaderLocation),
		r.Get(httpi.HeaderUserAgent), r.Get(httpi.HeaderXForwardedFor))
}

type HttpContext contexti.RequestContext[http.Request]

func (c *HttpContext) SetHeader(md metadata.MD) error {
	for k, v := range md {
		c.Request.Header[k] = v
	}
	if c.ServerTransportStream != nil {
		err := c.ServerTransportStream.SetHeader(md)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *HttpContext) SendHeader(md metadata.MD) error {
	for k, v := range md {
		c.Request.Header[k] = v
	}
	if c.ServerTransportStream != nil {
		err := c.ServerTransportStream.SendHeader(md)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *HttpContext) WriteHeader(k, v string) error {
	c.Request.Header[k] = []string{v}

	if c.ServerTransportStream != nil {
		err := c.ServerTransportStream.SendHeader(metadata.MD{k: []string{v}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *HttpContext) SetCookie(v string) error {
	c.Request.Header[httpi.HeaderSetCookie] = []string{v}

	if c.ServerTransportStream != nil {
		err := c.ServerTransportStream.SendHeader(metadata.MD{httpi.HeaderSetCookie: []string{v}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *HttpContext) SetTrailer(md metadata.MD) error {
	for k, v := range md {
		c.Request.Header[k] = v
	}
	if c.ServerTransportStream != nil {
		err := c.ServerTransportStream.SetTrailer(md)
		if err != nil {
			return err
		}
	}
	return nil
}
