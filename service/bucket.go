package service

import (
	"context"
	"dim-edge-node/protocol"
)

// ListAllBuckets list all buckets
func (g *GRPCServer) ListAllBuckets(c context.Context, p *protocol.ListAllBucketsParams) (r *protocol.ListAllBucketsRes, err error) {
	bucket, err := g.Influx.ListAllBucket(int(p.Page), int(p.Size), p.Org, p.OrgID, p.Name)
	r.Bucket = bucket
	return
}

// RetrieveBucket retrieve buckets list
func (g *GRPCServer) RetrieveBucket(c context.Context, p *protocol.RetrieveBucketParams) (b *protocol.Bucket, err error) {
	bucket, err := g.Influx.RetrieveBucket(p.BucketID)
	b = bucket
	return
}

// RetrieveBucketLog retrieve bucket operation logs
func (g *GRPCServer) RetrieveBucketLog(c context.Context, p *protocol.RetreiveBucketLogParams) (r *protocol.RetreiveBucketLogRes, err error) {
	l, err := g.Influx.RetrieveBucketLog(p.BucketID, int(p.Page), int(p.Size))
	r.Log = l
	return
}

// DeleteBucket delete bucket
func (g *GRPCServer) DeleteBucket(c context.Context, p *protocol.DeleteBucketParams) (o *protocol.OpRes, err error) {
	err = g.Influx.DeleteBucket(p.BucketID)
	return
}
