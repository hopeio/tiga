package main

import (
	"github.com/hopeio/lemon/_example/user/api"
	"github.com/hopeio/lemon/_example/user/conf"
	"github.com/hopeio/lemon/utils/log"
	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"google.golang.org/grpc"
	"time"

	"github.com/hopeio/lemon/server"
)

func main() {
	//defer initialize.Start(conf.Conf, conf.Dao)()
	view.RegisterExporter(&exporter.PrintExporter{})
	view.SetReportingPeriod(time.Second)
	// GinRegister the view to collect gRPC client stats.
	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		log.Fatal(err)
	}

	// fiber(fasthttp)
	/*	app := fiber.New()
		pick.RegisterFiberService(service.GetUserService())
		pick.FiberWithCtx(app, true, initialize.GlobalConfig.Module)
		go app.Listen(":3000")*/

	server.Start(&server.Server{
		Config: conf.Conf.Server.Origin(),
		//为了可以自定义中间件
		GRPCOptions: []grpc.ServerOption{
			grpc.ChainUnaryInterceptor(),
			grpc.ChainStreamInterceptor(),
			//grpc.StatsHandler(&ocgrpc.ServerHandler{})
		},
		GRPCHandle: api.GrpcRegister,

		GinHandle: api.GinRegister,

		/*		GraphqlResolve: model.NewExecutableSchema(model.Config{
				Resolvers: &model.GQLServer{
					UserService:  service.GetUserService(),
					OauthService: service.GetOauthService(),
				}}),*/
	})
}
