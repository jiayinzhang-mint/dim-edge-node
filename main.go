package main

import (
	"dim-edge-node/service"
	"dim-edge-node/tcp"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func startSocketListener() (err error) {

	ssl := &tcp.SocketListener{
		Address: "localhost",
		Port:    "9000",
	}

	err = ssl.Listen()
	if err != nil {
		return
	}

	return
}

func startGRPCServer() (err error) {

	g := service.GRPCServer{
		Address: ":9090",
	}

	err = g.StartServer()
	if err != nil {
		return
	}

	return
}

var (
	g errgroup.Group
)

func main() {
	logrus.Info("dim-edge node service starting")

	g.Go(func() error {
		return startGRPCServer()
	})

	g.Go(func() error {
		return startSocketListener()
	})

	if err := g.Wait(); err != nil {
		logrus.Fatal(err)
	}
}
