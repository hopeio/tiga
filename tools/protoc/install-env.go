package main

import (
	"github.com/hopeio/tiga/utils/log"
	osi "github.com/hopeio/tiga/utils/os"
	"os"
)

// 提供给使用框架的人安装所需环境
func main() {
	libDir, _ := osi.CmdLog("go list -m -f {{.Dir}}  github.com/hopeio/tiga")
	os.Chdir(libDir)
	osi.CmdLog("go install google.golang.org/protobuf/cmd/protoc-gen-go")
	protoccmd := "protoc -I" + libDir + "/protobuf/_proto --go_out=paths=source_relative:" + libDir + "/.. " + libDir + "/protobuf/_proto/tiga/protobuf/utils/"
	osi.CmdLog(protoccmd + "patch/*.proto")
	osi.CmdLog(protoccmd + "apiconfig/*.proto")
	osi.CmdLog(protoccmd + "openapiconfig/*.proto")
	osi.CmdLog(protoccmd + "enum/*.proto")
	osi.CmdLog("go install " + libDir + "/tools/protoc/protoc-gen-grpc-gin")
	osi.CmdLog("go install " + libDir + "/tools/protoc/protoc-gen-enum")
	osi.CmdLog("go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway")
	osi.CmdLog("go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2")
	osi.CmdLog("go install github.com/alta/protopatch/cmd/protoc-gen-go-patch")
	osi.CmdLog("go install google.golang.org/grpc/cmd/protoc-gen-go-grpc")
	osi.CmdLog("go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators")
	osi.CmdLog("go install " + libDir + "/tools/protoc/protoc-gen-go-patch")
	//osi.CmdLog("go install github.com/danielvladco/go-proto-gql/protoc-gen-gql")
	osi.CmdLog("go install " + libDir + "/tools/protoc/protoc-gen-gql")
	//osi.CmdLog("go install github.com/danielvladco/go-proto-gql/protoc-gen-gogql")
	osi.CmdLog("go install " + libDir + "/tools/protoc/protoc-gen-gogql")
	osi.CmdLog("go install " + libDir + "/tools/protoc/protogen")
	log.Info("安装成功")
}
