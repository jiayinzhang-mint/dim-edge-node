package service

import (
	"context"
	"dim-edge-node/protocol"
)

// ListAuthorization list all buckets
func (g *GRPCServer) ListAuthorization(c context.Context, p *protocol.ListAuthorizationParams) (*protocol.ListAuthorizationRes, error) {
	a := &protocol.ListAuthorizationRes{}
	auth, err := g.Influx.ListAuthorization(p.UserID, p.User, p.OrgID, p.Org)
	if err != nil {
		return nil, err
	}

	a.Authorization = auth
	return a, nil
}

// CreateAuthorization create authorization
func (g *GRPCServer) CreateAuthorization(c context.Context, p *protocol.CreateAuthorizationParams) (*protocol.Authorization, error) {
	auth, err := g.Influx.CreateAuthorization(p.Status, p.Description, p.OrgID, p.Permissions)
	if err != nil {
		return nil, err
	}

	return auth, err
}
