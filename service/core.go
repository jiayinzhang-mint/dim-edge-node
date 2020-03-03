package service

import (
	"dim-edge-node/protocol"
	"dim-edge-node/store"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer instance
type GRPCServer struct {
	Address string `json:"address"`
	Influx  *store.Influx
}

// StartServer start server
func (g *GRPCServer) StartServer() (err error) {

	lis, listenErr := net.Listen("tcp", g.Address)
	if listenErr != nil {
		log.Fatalf("dim-edge node GRPC server failed to listen: %v", listenErr)
	}
	s := grpc.NewServer()

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
