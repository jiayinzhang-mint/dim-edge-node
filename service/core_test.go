package service

import "testing"

func TestStartServer(*testing.T) {
	g := GRPCServer{
		Address: ":9090",
	}

	g.StartServer()
}
