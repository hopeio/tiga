package fasthttpi

import (
	httpi "github.com/hopeio/lemon/utils/net/http"
	stringsi "github.com/hopeio/lemon/utils/strings"
	"net/url"

	"github.com/valyala/fasthttp"
)

func GetToken(req *fasthttp.Request) string {
	if token := stringsi.BytesToString(req.Header.Peek(httpi.HeaderAuthorization)); token != "" {
		return token
	}
	if cookie := stringsi.BytesToString(req.Header.Cookie("token")); len(cookie) > 0 {
		token, _ := url.QueryUnescape(cookie)
		return token
	}
	return ""
}
