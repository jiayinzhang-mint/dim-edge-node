package service

import (
	"context"
	"dim-edge-node/protocol"
)

// QueryData query data
func (g *GRPCServer) QueryData(c context.Context, p *protocol.QueryParams) (*protocol.QueryRes, error) {
	var (
		r   *protocol.QueryRes
		err error
	)
	return r, err
}

// InsertData insert data
func (g *GRPCServer) InsertData(c context.Context, p *protocol.InsertDataParams) (*protocol.InsertDataRes, error) {
	var (
		r   *protocol.InsertDataRes
		err error
	)
	return r, err
}
