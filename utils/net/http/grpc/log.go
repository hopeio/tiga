package grpci

import (
	"github.com/hopeio/tiga/utils/log"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc/grpclog"
)

func init() {
	grpclog.SetLoggerV2(zapgrpc.NewLogger(log.Default.Logger))
}
