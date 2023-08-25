package grpci

import (
	"github.com/hopeio/lemon/utils/log"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc/grpclog"
)

func init() {
	grpclog.SetLoggerV2(zapgrpc.NewLogger(log.Default.Logger))
}
