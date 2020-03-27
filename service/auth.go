package service

import (
	"context"
	"dim-edge-node/protocol"

	"github.com/golang/protobuf/ptypes/empty"
)

// ListAuthorization list all authorization
func (g *GRPCServer) ListAuthorization(c context.Context, p *protocol.ListAuthorizationParams) (*protocol.ListAuthorizationRes, error) {
	a := &protocol.ListAuthorizationRes{}
	auth, err := g.Influx.ListAuthorization(p.UserID, p.User, p.OrgID, p.Org)
	if err != nil {
		return &protocol.ListAuthorizationRes{}, err
	}

	a.Authorization = auth
	return a, nil
}

// CreateAuthorization create authorization
func (g *GRPCServer) CreateAuthorization(c context.Context, p *protocol.CreateAuthorizationParams) (*protocol.Authorization, error) {
	auth, err := g.Influx.CreateAuthorization(p.Status, p.Description, p.OrgID, p.Permissions)
	if err != nil {
		return &protocol.Authorization{}, err
	}

	return auth, err
}

// SignIn sign in to influxdb
func (g *GRPCServer) SignIn(c context.Context, p *protocol.SignInParams) (*empty.Empty, error) {
	var (
		err error
	)

	err = g.Influx.SignIn(p.Username, p.Password)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

// SignOut sign out
func (g *GRPCServer) SignOut(c context.Context, p *empty.Empty) (*empty.Empty, error) {
	var (
		err error
	)

	err = g.Influx.SignOut()
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
