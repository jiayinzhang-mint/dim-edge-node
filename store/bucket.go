package store

import (
	"dim-edge-node/utils"
)

// ListAllBucket list all bucket
func (i *Influx) ListAllBucket(page int, size int, org string, orgID string, name string) (bucket []interface{}, err error) {
	offset := (page - 1) * size

	_, err = utils.HTTP().Get(i.HTTPClient, i.GetBasicURL()+"/buckets", map[string]string{
		"offset": string(offset),
		"limit":  string(size),
		"org":    org,
		"orgID":  orgID,
		"name":   name,
	}, nil)

	if err != nil {
		return
	}

	return
}
