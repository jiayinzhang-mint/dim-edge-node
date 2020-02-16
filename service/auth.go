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
