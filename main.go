package main

import (
	"dim-edge/node/service"
	"dim-edge/node/store"
	"dim-edge/node/tcp"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// BuildEnv build mode (dev or prod)
var BuildEnv string

func startSocketListener() (err error) {

	ssl := &tcp.SocketListener{
		Port: "9000",
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

	if BuildEnv == "prod" {
		g.Influx = &store.Influx{
			Address: os.Getenv("INFLUX_ADDRESS"),
		}
	} else {
		g.Influx = &store.Influx{
			Address: "http://192.168.64.24:31830",
		}
	}

	// connect to db
	if err = g.Influx.ConnectToDB(); err != nil {
		logrus.Error(err)
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
	logrus.Info("build env: " + BuildEnv)

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
