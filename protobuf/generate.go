package main

import (
	execi "github.com/hopeio/tiga/utils/os/exec"
	_go "github.com/hopeio/tiga/utils/sdk/go"
	"log"
	"os"
	"strings"
)

//go:generate mockgen -destination ../protobuf/user/user.mock.go -package user -source ../protobuf/user/user.service_grpc.pb.go UserServiceServer

var (
	libtigaDir, proto    string
	pwd, gopath, include string
)

func init() {
	gopath = os.Getenv("GOPATH")
	if strings.HasSuffix(gopath, "/") {
		gopath = gopath[:len(gopath)-1]
	}

	pwd, _ = os.Getwd()
	libtigaDir = _go.GetDepDir(Deptiga)
	proto = libtigaDir + "/protobuf/_proto"
	//libGatewayDir := _go.GetDepDir(DepGrpcGateway)
	//libGoogleDir := _go.GetDepDir(DepGoogleapis)

	include = "-I" + proto
}

func main() {
	//single("/content/moment.model.proto")
	generate(proto + "/tiga/protobuf")
	//gengql()
	os.Chdir(pwd)
}

const goOut = "go-patch_out=plugin=go,paths=source_relative"
const grpcOut = "go-patch_out=plugin=go-grpc,paths=source_relative"
const enumOut = "enum_out=paths=source_relative"

const (
	goListDir      = `go list -m -f {{.Dir}} `
	goListDep      = `go list -m -f {{.Path}}@{{.Version}} `
	DepGoogleapis  = "github.com/googleapis/googleapis@v0.0.0-20220520010701-4c6f5836a32f"
	Deptiga        = "github.com/hopeio/tiga"
	DepGrpcGateway = "github.com/grpc-ecosystem/grpc-gateway/v2"
	DepProtopatch  = "github.com/alta/protopatch"
)

var model = []string{goOut, grpcOut}

func generate(dir string) {
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	for i := range fileInfos {
		if fileInfos[i].IsDir() {
			generate(dir + "/" + fileInfos[i].Name())
		}
		if strings.HasSuffix(fileInfos[i].Name(), "enum.proto") {
			arg := "protoc " + include + " " + dir + "/" + fileInfos[i].Name() + " --" + enumOut + ":" + libtigaDir + "/.."
			execi.Run(arg)
		}
	}

	for _, plugin := range model {
		arg := "protoc " + include + " " + dir + "/*.proto --" + plugin + ":" + libtigaDir + "/.."
		execi.Run(arg)
	}

}
