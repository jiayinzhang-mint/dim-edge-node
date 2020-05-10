package service

import (
	"dim-edge/node/protocol"
	"dim-edge/node/store"
	"log"
	"net"
	"time"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	ot "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer instance
type GRPCServer struct {
	Address string `json:"address"`
	Influx  *store.Influx
	Tracer  ot.Tracer
}

// StartServer start server
func (g *GRPCServer) StartServer() (err error) {
	var options []grpc.ServerOption

	// init a new tracer
	jcfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	g.Tracer, _, err = jcfg.New(
		"dim-edge-node",
		jaegercfg.Logger(jaeger.StdLogger),
	)
	if err != nil {
		return
	}

	// tracer client middleware
	options = append(options, grpc.UnaryInterceptor(
		grpc_opentracing.UnaryServerInterceptor(
			grpc_opentracing.WithTracer(g.Tracer))))

	lis, listenErr := net.Listen("tcp", g.Address)
	if listenErr != nil {
		log.Fatalf("dim-edge node GRPC server failed to listen: %v", listenErr)
	}
	s := grpc.NewServer(options...)

	// register services
	protocol.RegisterStoreServiceServer(s, g)
	reflection.Register(s)

	// Start serve
	logrus.New().Infof("dim-edge node GRPC server listening at %s", g.Address)
	if err = s.Serve(lis); err != nil {
		logrus.Fatalf("dim-edge node GRPC server failed to serve: %v", err)
		return
	}

	return
}
