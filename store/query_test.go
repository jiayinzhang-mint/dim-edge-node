package store

import (
	"testing"
	"time"

	"github.com/influxdata/influxdb-client-go"
	"github.com/sirupsen/logrus"
)

func TestQuery(*testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "4oXjSoIuU1F3A1zu-xYp0eJ9q_vsLQmtDPPTNuDnrs7R7H7qGAQ1GNaX4hNtJKx5ZRfnoj_TW5Uwe5NJUBvLOA==",
	}
	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	c, err := influx.InsertData(&[]influxdb.Metric{
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
		)}, "dim-edge", "INSDIM")

	logrus.Info(c, err)
}
