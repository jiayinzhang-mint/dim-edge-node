package service

import (
	"context"
	"dim-edge-node/protocol"
)

// ListAuthorization list all buckets
func (g *GRPCServer) ListAuthorization(c context.Context, p *protocol.ListAuthorizationParams) (auth *protocol.ListAuthorizationRes, err error) {
	auth.Authorization, err = g.Influx.ListAuthorization(p.UserID, p.User, p.OrgID, p.Org)
	return
}

// CreateAuthorization create authorization
func (g *GRPCServer) CreateAuthorization(c context.Context, p *protocol.CreateAuthorizationParams) (auth *protocol.Authorization, err error) {
	auth, err = g.Influx.CreateAuthorization(p.Status, p.Description, p.OrgID, p.Permissions)
	return
}
