package store

import (
	"context"
	"testing"
	"time"

	"github.com/influxdata/influxdb-client-go"
	"github.com/sirupsen/logrus"
)

func TestCreateFluxClient(t *testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "4oXjSoIuU1F3A1zu-xYp0eJ9q_vsLQmtDPPTNuDnrs7R7H7qGAQ1GNaX4hNtJKx5ZRfnoj_TW5Uwe5NJUBvLOA==",
	}

	err := influx.ConnectToDB()
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(influx.DBClient)
}

func TestInsertData(t *testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "4oXjSoIuU1F3A1zu-xYp0eJ9q_vsLQmtDPPTNuDnrs7R7H7qGAQ1GNaX4hNtJKx5ZRfnoj_TW5Uwe5NJUBvLOA==",
	}

	err := influx.ConnectToDB()
	if err != nil {
		logrus.Error(err)
	}
	myMetrics := []influxdb.Metric{
		influxdb.NewRowMetric(
			map[string]interface{}{"memory": 1000, "cpu": 0.93},
			"system-metrics",
			map[string]string{"hostname": "hal9000"},
			time.Now().Local(),
		),
		influxdb.NewRowMetric(
			map[string]interface{}{"memory": 1000, "cpu": 0.93},
			"system-metrics",
			map[string]string{"hostname": "hal9000"},
			time.Now().Local(),
		),
	}

	// The actual write..., this method can be called concurrently.
	n, err := influx.GetDB().Write(context.Background(), "dim-edge", "INSDIM", myMetrics...)
	if err != nil {
		logrus.Fatal(err) // as above use your own error handling here.
	}

	logrus.Info(n)

	influx.GetDB().Close()
}

func TestQueryData(t *testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "4oXjSoIuU1F3A1zu-xYp0eJ9q_vsLQmtDPPTNuDnrs7R7H7qGAQ1GNaX4hNtJKx5ZRfnoj_TW5Uwe5NJUBvLOA==",
	}

	err := influx.ConnectToDB()
	if err != nil {
		logrus.Error(err)
	}

	// Query
	res, queryErr := influx.GetDB().QueryCSV(context.Background(),
		`from(bucket: "dim-edge")
  		|> range(start: -10h)
  		|> filter(fn: (r)=>
				r._field == "cpu" and
				r._measurement == "system-metrics" and
				r.hostname == "hal9000"
			)`,
		"INSDIM")
	if queryErr != nil {
		logrus.Error(err)
	}

	// Marshal data
	type influxRecord struct {
		Zone   ***string   `flux:"name" json:"zone"`
		Stop   time.Time   `flux:"_stop" json:"-"`
		Start  time.Time   `flux:"_start" json:"-"`
		Time   time.Time   `flux:"_time" json:"date"`
		HostIP string      `flux:"host_ip" json:"-"`
		Count  interface{} `flux:"_value" json:"count"`
	}

	var r influxRecord
	for res.Next() {
		mErr := res.Unmarshal(&r)
		logrus.Info(r)
		if mErr != nil {
			logrus.Error(mErr)
		}
	}

	// Close DB
	influx.GetDB().Close()

}
