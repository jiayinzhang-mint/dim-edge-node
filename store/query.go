package store

import (
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/sirupsen/logrus"
)

// InsertData insert data
func (i *Influx) InsertData(p []*influxdb2.Point, bucket string, org string) (err error) {
	writeAPI := i.GetDB().WriteApiBlocking(org, bucket)
	err = writeAPI.WritePoint(context.Background(), p...)
	if err != nil {
		logrus.Error(err)
		return
	}

	return
}

// QueryData query data
func (i *Influx) QueryData(queryString string, org string) (res *influxdb2.QueryTableResult, err error) {
	queryAPI := i.GetDB().QueryApi(org)

	res, err = queryAPI.Query(context.Background(), queryString)
	if err != nil {
		logrus.Error(err)
		return
	}

	return
}
