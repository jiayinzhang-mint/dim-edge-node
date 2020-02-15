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
func (g *GRPCServer) StartServer() {
	// Connect to DB
	g.Influx = &store.Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "4oXjSoIuU1F3A1zu-xYp0eJ9q_vsLQmtDPPTNuDnrs7R7H7qGAQ1GNaX4hNtJKx5ZRfnoj_TW5Uwe5NJUBvLOA==",
	}
	connectErr := g.Influx.ConnectToDB()
	if connectErr != nil {
		logrus.Error(connectErr)
	}

	lis, listenErr := net.Listen("tcp", g.Address)
	if listenErr != nil {
		log.Fatalf("dim-edge GRPC server failed to listen: %v", listenErr)
	}
	s := grpc.NewServer()

	// register services
	protocol.RegisterStoreServiceServer(s, g)
	reflection.Register(s)

	// Start serve
	logrus.New().Infof("dim-edge GRPC server listening at %s", g.Address)
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("dim-edge GRPC server failed to serve: %v", err)
	}
}
