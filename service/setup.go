package service

import (
	"context"
	"dim-edge/node/protocol"

	"github.com/golang/protobuf/ptypes/empty"
)

// CheckSetup check if db has been setup
func (g *GRPCServer) CheckSetup(ctx context.Context, in *empty.Empty) (*protocol.CheckSetupRes, error) {
	setup, err := g.Influx.CheckSetup()
	if err != nil {
		return nil, err
	}
	r := &protocol.CheckSetupRes{
		Setup: setup,
	}

	return r, nil
}

// Setup setup db
func (g *GRPCServer) Setup(ctx context.Context, in *protocol.SetupParams) (*empty.Empty, error) {
	r := &empty.Empty{}

	err := g.Influx.Setup(
		in.Username,
		in.Password,
		in.Org,
		in.Bucket,
		int(in.RetentionPeriodHrs),
	)
	if err != nil {
		return nil, err
	}

	return r, nil
}
