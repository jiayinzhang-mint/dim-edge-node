package store

import (
	"context"

	influxdbAPI "github.com/influxdata/influxdb-client-go/api"
	"github.com/influxdata/influxdb-client-go/api/write"
	"github.com/sirupsen/logrus"
)

// InsertData insert data
func (i *Influx) InsertData(p []*write.Point, bucket string, org string) (err error) {
	writeAPI := i.GetDB().WriteApiBlocking(org, bucket)
	err = writeAPI.WritePoint(context.Background(), p...)
	if err != nil {
		logrus.Error(err)
		return
	}

	return
}

// QueryData query data
func (i *Influx) QueryData(queryString string, org string) (res *influxdbAPI.QueryTableResult, err error) {
	queryAPI := i.GetDB().QueryApi(org)

	res, err = queryAPI.Query(context.Background(), queryString)
	if err != nil {
		logrus.Error(err)
		return
	}

	return
}
