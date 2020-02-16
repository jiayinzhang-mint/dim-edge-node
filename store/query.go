package store

import (
	"context"

	"github.com/influxdata/influxdb-client-go"
	"github.com/sirupsen/logrus"
)

// InsertData insert data
func (i *Influx) InsertData(metrics *[]influxdb.Metric, bucket string, org string) (count int, err error) {
	count, err = i.GetDB().Write(context.Background(), "dim-edge", "INSDIM", *metrics...)
	if err != nil {
		logrus.Fatal(err) // as above use your own error handling here.
	}

	return
}
