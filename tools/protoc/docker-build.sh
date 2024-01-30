gopath=/mnt/d/SDK/gopath
protoc=/mnt/d/tools/protoc-22.3-linux-x86_64

cd $(dirname $0) && pwd
tiga_dir=$(cd ../..;pwd)

goproxy=https://goproxy.io,https://goproxy.cn,direct
goimage=golang:1.20


# install tools
docker run --rm -e GOPROXY=$goproxy -v $gopath:/go -v $tiga_dir:/work -w /work/tools/protoc $goimage bash ./install-tools.sh

dockerTmpDir=$tiga_dir/tools/protoc/_docker
mkdir $dockerTmpDir
cp $gopath/bin/protoc-gen-enum $dockerTmpDir/
cp $gopath/bin/protoc-gen-go $dockerTmpDir/
cp $gopath/bin/protoc-gen-go-grpc $dockerTmpDir/
cp $gopath/bin/protoc-gen-go-patch $dockerTmpDir/
cp $gopath/bin/protoc-gen-govalidators $dockerTmpDir/
cp $gopath/bin/protoc-gen-grpc-gateway $dockerTmpDir/
cp $gopath/bin/protoc-gen-grpc-gin $dockerTmpDir/
cp $gopath/bin/protoc-gen-openapiv2 $dockerTmpDir/
cp $gopath/bin/protoc-gen-gql $dockerTmpDir/
cp $gopath/bin/protoc-gen-gogql $dockerTmpDir/
cp $gopath/bin/gqlgen $dockerTmpDir/
cp $gopath/bin/protogen $dockerTmpDir/
cp -r $tiga_dir/protobuf/_proto $dockerTmpDir/_proto
cp -r $protoc $dockerTmpDir/protoc

cd $dockerTmpDir
docker build -t jybl/goprotoc -f $tiga_dir/tools/protoc/Dockerfile $dockerTmpDir
rm -r $dockerTmpDir
docker push jybl/goprotoc