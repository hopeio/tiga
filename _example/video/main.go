package main

import (
	"github.com/hopeio/tiga/_example/protobuf/user"
	"github.com/hopeio/tiga/utils/log"
	grpci "github.com/hopeio/tiga/utils/net/http/grpc"
	"golang.org/x/net/context"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpci.GetDefaultClient("localhost:8080")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := user.NewUserServiceClient(conn)
	ctx := context.Background()
	_, err = client.Signup(ctx, &user.SignupReq{
		Password: "123456",
		Name:     "123",
		Gender:   1,
		Mail:     "",
		Phone:    "",
		VCode:    "1",
	})
	if err != nil {
		log.Error(err)
		return
	}
}
