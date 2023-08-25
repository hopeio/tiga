package api

import (
	"github.com/hopeio/lemon/_example/protobuf/user"
	userService "github.com/hopeio/lemon/_example/user/service"
	"google.golang.org/grpc"
)

func GrpcRegister(gs *grpc.Server) {
	user.RegisterUserServiceServer(gs, userService.GetUserService())

}
