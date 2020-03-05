package service

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
)

func TestCheckSetup(*testing.T) {
	g := GRPCServer{
		Address: "http://192.168.64.9:3252",
	}

	g.CheckSetup(context.TODO(), &empty.Empty{})
}
