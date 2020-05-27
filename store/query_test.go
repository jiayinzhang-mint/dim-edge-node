package store

import (
	"testing"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api/write"
	"github.com/sirupsen/logrus"
)

func TestInsertData(*testing.T) {
	influx := &Influx{
		Address: "http://192.168.64.16:32565",
	}
	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	influx.SignIn("mint", "131001250115zHzH")

	p := influxdb2.NewPoint(
		"system-metrics",
		map[string]string{"hostname": "hal9000"},
		map[string]interface{}{"memory": 1000.0, "cpu": 0.93},
		time.Now(),
	)

	err := influx.InsertData([]*write.Point{p}, "insdim", "insdim")

	logrus.Info(err)
}

func TestQuery(*testing.T) {
	influx := &Influx{
		Address: "http://192.168.64.18:31048",
	}

	err := influx.ConnectToDB()
	if err != nil {
		logrus.Error(err)
	}

	influx.SignIn("mint", "131001250115zHzH")

	// Query
	res, queryErr := influx.QueryData(
		`from(bucket: "insdim")
  		|> range(start: -10h)
  		|> filter(fn: (r)=>
				r._field == "cpu" and
				r._measurement == "system-metrics" and
				r.hostname == "hal9000"
			)`,
		"insdim")
	if queryErr != nil {
		logrus.Error(err)
	}

	for res.Next() {
		logrus.Info(res.Record().Field(), res.Record().Value())
	}

	// Close DB
	influx.GetDB().Close()
}
