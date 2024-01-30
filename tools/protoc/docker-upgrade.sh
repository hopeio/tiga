cd $(dirname $0) && pwd
tiga_dir=$(cd ../..;pwd)

docker build -t jybl/goprotoc -f $tiga_dir/tools/protoc/Dockerfile-upgrade .