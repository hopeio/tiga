cd $(dirname $0) && pwd
lemon=$(go list -m -f {{.Dir}}  github.com/hopeio/lemon)
lemon=${lemon//\\/\/}
protoDir=$lemon/protobuf/_proto

# 安装
cd $lemon/tools/protoc
echo "开始安装"
go install google.golang.org/protobuf/cmd/protoc-gen-go
protoc -I$protoDir --go_out=paths=source_relative:$lemon/.. $protoDir/lemon/protobuf/utils/**/*.proto
go install ./protoc-gen-enum
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install github.com/alta/protopatch/cmd/protoc-gen-go-patch
go install ./protoc-gen-grpc-gin
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
go install ./protoc-gen-go-patch
# go install github.com/danielvladco/go-proto-gql/protoc-gen-gql
# go install github.com/danielvladco/go-proto-gql/protoc-gen-gogql
go install ./protoc-gen-gql
go install ./protoc-gen-gogql
go install github.com/99designs/gqlgen
go install ./protogen
echo "安装完成"
