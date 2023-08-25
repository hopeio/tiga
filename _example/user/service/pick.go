package service

import (
	"github.com/hopeio/lemon/_example/user/middle"
	"github.com/hopeio/lemon/context/http_context"
	"github.com/hopeio/lemon/pick"
	"github.com/hopeio/lemon/protobuf/response"
	"net/http"
)

func (u *UserService) Service() (describe, prefix string, middleware []http.HandlerFunc) {
	return "用户相关", "/api/user", []http.HandlerFunc{middle.Log}
}

func (*UserService) Add(ctx *http_context.Context, req *response.TinyRep) (*response.TinyRep, error) {
	//对于一个性能强迫症来说，我宁愿它不优雅一些也不能接受每次都调用
	pick.Api(func() {
		pick.Post("/add").
			Title("用户注册").
			Version(1).
			CreateLog("1.0.0", "jyb", "2019/12/16", "创建").
			ChangeLog("1.0.1", "jyb", "2019/12/16", "修改测试").End()
	})
	return req, nil
}
