package store

import (
	"dim-edge-node/protocol"
	"encoding/json"
	"strconv"
)

// ListAllBucket list all bucket
func (i *Influx) ListAllBucket(page int, size int, org string, orgID string, name string) (bucket []protocol.Bucket, err error) {
	offset := (page - 1) * size

	offsetStr := strconv.Itoa(offset)
	sizeStr := strconv.Itoa(size)

	res, err := i.HTTPInstance.Get(i.HTTPClient, i.GetBasicURL()+"/buckets", map[string]string{
		"offset": offsetStr,
		"limit":  sizeStr,
		"org":    org,
		"orgID":  orgID,
		"name":   name,
	}, nil)
	if err != nil {
		return
	}

	type b struct {
		Bucket []protocol.Bucket `json:"buckets"`
	}
	var resBody b
	if err = json.Unmarshal(res, &resBody); err != nil {
		return
	}

	bucket = resBody.Bucket
	return
}

// RetrieveBucket retrieve a bucket by id
func (i *Influx) RetrieveBucket(bucketID string) (bucket protocol.Bucket, err error) {
	res, err := i.HTTPInstance.Get(i.HTTPClient, i.GetBasicURL()+"/buckets/"+bucketID, nil, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, &bucket); err != nil {
		return
	}

	return
}

// RetrieveBucketLog retrieve bucket log by id
func (i *Influx) RetrieveBucketLog(bucketID string, page int, size int) (log []protocol.RetreiveBucketLogRes_Log, err error) {
	offset := (page - 1) * size

	offsetStr := strconv.Itoa(offset)
	sizeStr := strconv.Itoa(size)

	res, err := i.HTTPInstance.Get(i.HTTPClient, i.GetBasicURL()+"/buckets/"+bucketID+"/logs", map[string]string{
		"offset": offsetStr,
		"limit":  sizeStr,
	}, nil)
	if err != nil {
		return
	}

	type b struct {
		Logs []protocol.RetreiveBucketLogRes_Log `json:"logs"`
	}
	var resBody b
	if err = json.Unmarshal(res, &resBody); err != nil {
		return
	}

	log = resBody.Logs
	return
}
