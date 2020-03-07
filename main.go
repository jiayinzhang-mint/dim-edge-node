package main

import (
	"dim-edge-node/service"
	"dim-edge-node/store"
	"dim-edge-node/tcp"
	"os"

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

	g.Influx = &store.Influx{
		Address: os.Getenv("INFLUX_ADDRESS"),
	}

	err = g.Influx.ConnectToDB()
	if err != nil {
		logrus.Error(err)
	}

	return
}

var (
	g errgroup.Group
)

func main() {
	logrus.Info("dim-edge node service starting")

	logrus.Info("INFLUX_ADDRESS", os.Getenv("INFLUX_ADDRESS"))

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
