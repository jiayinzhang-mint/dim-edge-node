package service

import (
	"context"
	"dim-edge-node/protocol"
)

// ListAllBuckets list all buckets
func (g *GRPCServer) ListAllBuckets(c context.Context, p *protocol.ListAllBucketsParams) (*protocol.ListAllBucketsRes, error) {
	r := &protocol.ListAllBucketsRes{}

	bucket, err := g.Influx.ListAllBucket(int(p.Page), int(p.Size), p.Org, p.OrgID, p.Name)
	if err != nil {
		return r, err
	}

	r.Bucket = bucket
	return r, nil
}

// RetrieveBucket retrieve buckets list
func (g *GRPCServer) RetrieveBucket(c context.Context, p *protocol.RetrieveBucketParams) (*protocol.Bucket, error) {
	bucket, err := g.Influx.RetrieveBucket(p.BucketID)
	if err != nil {
		return &protocol.Bucket{}, err
	}

	return bucket, nil
}

// RetrieveBucketLog retrieve bucket operation logs
func (g *GRPCServer) RetrieveBucketLog(c context.Context, p *protocol.RetreiveBucketLogParams) (*protocol.RetreiveBucketLogRes, error) {
	l, err := g.Influx.RetrieveBucketLog(p.BucketID, int(p.Page), int(p.Size))
	if err != nil {
		return nil, err
	}

	r := &protocol.RetreiveBucketLogRes{}
	r.Log = l
	return r, nil
}

// DeleteBucket delete bucket
func (g *GRPCServer) DeleteBucket(c context.Context, p *protocol.DeleteBucketParams) (*protocol.OpRes, error) {
	err := g.Influx.DeleteBucket(p.BucketID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
