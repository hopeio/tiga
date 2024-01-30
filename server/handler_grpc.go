package server

import (
	"context"
	"fmt"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/hopeio/tiga/context/http_context"
	"github.com/hopeio/tiga/protobuf/errorcode"
	"github.com/hopeio/tiga/utils/encoding/json"
	"github.com/hopeio/tiga/utils/log"
	runtimei "github.com/hopeio/tiga/utils/runtime"
	stringsi "github.com/hopeio/tiga/utils/strings"
	"github.com/hopeio/tiga/utils/verification/validator"
	"github.com/modern-go/reflect2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"reflect"
	"runtime/debug"
)

func (s *Server) grpcHandler(conf *ServerConfig) *grpc.Server {
	grpclog.SetLoggerV2(zapgrpc.NewLogger(log.Default.Logger))
	if s.GRPCHandle != nil {
		var stream = []grpc.StreamServerInterceptor{StreamAccess}
		var unary = []grpc.UnaryServerInterceptor{UnaryAccess, Validator}
		if conf.Prometheus {
			stream = append(stream, grpc_prometheus.StreamServerInterceptor)
			unary = append(unary, grpc_prometheus.UnaryServerInterceptor)
		}
		stream = append(stream, grpc_validator.StreamServerInterceptor())
		unary = append(unary, grpc_validator.UnaryServerInterceptor())
		s.GRPCOptions = append([]grpc.ServerOption{
			grpc.ChainStreamInterceptor(stream...),
			grpc.ChainUnaryInterceptor(unary...),
		}, s.GRPCOptions...)
		grpcServer := grpc.NewServer(s.GRPCOptions...)
		if conf.Prometheus {
			grpc_prometheus.Register(grpcServer)
		}
		s.GRPCHandle(grpcServer)
		reflection.Register(grpcServer)
		return grpcServer
	}
	return nil
}

func UnaryAccess(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			frame := debug.Stack()
			log.Default.Errorw(fmt.Sprintf("panic: %v", r), zap.ByteString(log.Stack, frame))
			err = errorcode.SysError.ErrRep()
		}
	}()

	resp, err = handler(ctx, req)
	var code int
	//不能添加错误处理，除非所有返回的结构相同
	if err != nil {
		if v, ok := err.(interface{ GRPCStatus() *status.Status }); !ok {
			err = errorcode.Unknown.Message(err.Error())
			code = int(errorcode.Unknown)
		} else {
			code = int(v.GRPCStatus().Code())
		}
	}
	if err == nil && reflect2.IsNil(resp) {
		resp = reflect.New(reflect.TypeOf(resp).Elem()).Interface()
	}

	body, _ := json.Marshal(req)
	result, _ := json.Marshal(resp)
	ctxi := http_context.ContextFromContext(ctx)
	defaultAccessLog(ctxi, info.FullMethod, "grpc",
		stringsi.BytesToString(body), stringsi.BytesToString(result),
		code)
	return resp, err
}

func StreamAccess(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func() {
		if r := recover(); r != nil {
			frame, _ := runtimei.GetCallerFrame(2)
			log.Default.Errorw(fmt.Sprintf("panic: %v", r), zap.String(log.Stack, fmt.Sprintf("%s:%d (%#x)\n\t%s\n", frame.File, frame.Line, frame.PC, frame.Function)))
			err = errorcode.SysError.ErrRep()
		}
		//不能添加错误处理，除非所有返回的结构相同
		if err != nil {
			if _, ok := err.(interface{ GRPCStatus() *status.Status }); !ok {
				err = errorcode.Unknown.Message(err.Error())
			}
		}
	}()

	return handler(srv, stream)
}

type recvWrapper struct {
	grpc.ServerStream
}

func (s *recvWrapper) SendMsg(m interface{}) error {
	return s.ServerStream.SendMsg(m)
}

func (s *recvWrapper) RecvMsg(m interface{}) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	return nil
}

func Validator(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	if err = validator.Validator.Struct(req); err != nil {
		return nil, errorcode.InvalidArgument.Message(validator.Trans(err))
	}
	return handler(ctx, req)
}
