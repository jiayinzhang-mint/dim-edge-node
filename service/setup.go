package service

import (
	"context"
	"dim-edge-node/protocol"

	"github.com/golang/protobuf/ptypes/empty"
)

// CheckSetup check if db has been setup
func (g *GRPCServer) CheckSetup(ctx context.Context, in *empty.Empty) (r *protocol.CheckSetupRes, err error) {
	setup, err := g.Influx.CheckSetup()
	r.Setup = setup
	return
}

// Setup setup db
func (g *GRPCServer) Setup(ctx context.Context, in *protocol.SetupParams) (r *empty.Empty, err error) {
	err = g.Influx.Setup(
		in.Username,
		in.Password,
		in.Org,
		in.Bucket,
		int(in.RetentionPeriodHrs),
	)
	return
}
