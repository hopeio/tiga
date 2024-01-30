package api

import (
	"github.com/hopeio/tiga/_example/protobuf/user"
	userService "github.com/hopeio/tiga/_example/user/service"
	"google.golang.org/grpc"
)

func GrpcRegister(gs *grpc.Server) {
	user.RegisterUserServiceServer(gs, userService.GetUserService())

}
