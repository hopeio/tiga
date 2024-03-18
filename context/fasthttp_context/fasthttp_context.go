package fasthttp_context

import (
	"context"
	contexti "github.com/hopeio/tiga/context"
	httpi "github.com/hopeio/tiga/utils/net/http"
	fasthttpi "github.com/hopeio/tiga/utils/net/http/fasthttp"
	stringsi "github.com/hopeio/tiga/utils/strings"
	"github.com/valyala/fasthttp"
)

type Context = contexti.RequestContext[fasthttp.Request]

func ContextFromRequest(ctx context.Context, r *fasthttp.Request) *Context {
	ctxi := contexti.NewContext[fasthttp.Request](ctx)
	setWithReq(ctxi, r)
	return ctxi
}

func setWithReq(c *Context, r *fasthttp.Request) {
	c.Request = r
	c.Token = fasthttpi.GetToken(r)
	c.DeviceInfo = Device(&r.Header)
	c.Internal = stringsi.BytesToString(r.Header.Peek(httpi.GrpcInternal))
}

func Device(r *fasthttp.RequestHeader) *contexti.DeviceInfo {
	return contexti.Device(stringsi.BytesToString(r.Peek(httpi.HeaderDeviceInfo)),
		stringsi.BytesToString(r.Peek(httpi.HeaderArea)),
		stringsi.BytesToString(r.Peek(httpi.HeaderLocation)),
		stringsi.BytesToString(r.Peek(httpi.HeaderUserAgent)),
		stringsi.BytesToString(r.Peek(httpi.HeaderXForwardedFor)),
	)
}
