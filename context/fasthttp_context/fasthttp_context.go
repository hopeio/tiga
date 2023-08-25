package fasthttp_context

import (
	"context"
	contexti "github.com/hopeio/lemon/context"
	contexti2 "github.com/hopeio/lemon/utils/context"
	httpi "github.com/hopeio/lemon/utils/net/http"
	fasthttpi "github.com/hopeio/lemon/utils/net/http/fasthttp"
	stringsi "github.com/hopeio/lemon/utils/strings"
	"github.com/valyala/fasthttp"
)

type Context = contexti.Context[fasthttp.Request]

func ContextWithRequest(ctx context.Context, r *fasthttp.Request) *Context {
	ctxi := contexti2.NewCtx[fasthttp.Request](ctx)
	c := &Context{RequestContext: ctxi}
	setWithReq(c, r)
	return c
}

func setWithReq(c *Context, r *fasthttp.Request) {
	c.Request = r
	c.Token = fasthttpi.GetToken(r)
	c.DeviceInfo = Device(&r.Header)
	c.Internal = stringsi.BytesToString(r.Header.Peek(httpi.GrpcInternal))
}

func Device(r *fasthttp.RequestHeader) *contexti2.DeviceInfo {
	return contexti2.Device(stringsi.BytesToString(r.Peek(httpi.HeaderDeviceInfo)),
		stringsi.BytesToString(r.Peek(httpi.HeaderArea)),
		stringsi.BytesToString(r.Peek(httpi.HeaderLocation)),
		stringsi.BytesToString(r.Peek(httpi.HeaderUserAgent)),
		stringsi.BytesToString(r.Peek(httpi.HeaderXForwardedFor)),
	)
}
