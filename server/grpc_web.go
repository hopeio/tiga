package server

import (
	"github.com/hopeio/lemon/utils/net/http/grpc/web"
	"google.golang.org/grpc"
	"net/http"
)

func NewGrpcWebServer(grpcServer *grpc.Server) *web.WrappedGrpcServer {
	return web.WrapServer(grpcServer, web.WithAllowedRequestHeaders([]string{"*"}), web.WithWebsockets(true), web.WithWebsocketOriginFunc(func(req *http.Request) bool {
		return true
	}), web.WithOriginFunc(func(origin string) bool {
		return true
	}))
}
