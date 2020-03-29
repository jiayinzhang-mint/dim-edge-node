package store

import (
	"testing"
	"time"

	"github.com/influxdata/influxdb-client-go"
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
	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	c, err := influx.InsertData(&[]influxdb.Metric{
		influxdb.NewRowMetric(
			map[string]interface{}{"memory": 1000.0, "cpu": 0.93},
			"system-metrics",
			map[string]string{"hostname": "hal9000"},
			time.Now().Local(),
		),
		influxdb.NewRowMetric(
			map[string]interface{}{"memory": 1000.0, "cpu": 0.93},
			"system-metrics",
			map[string]string{"hostname": "hal9000"},
			time.Now().Local(),
		)}, "insdim", "insdim")

	logrus.Info(c, err)
}

func TestQuery(*testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "OmtoG5-MWHplbyT0QS2-HoDyfKAUpbYkkXf_W3nYDqwZe631h-NRGygJoEFyUeVxXknTewpOwa97A-q0BCI3eg==",
	}

	err := influx.ConnectToDB()
	if err != nil {
		logrus.Error(err)
	}

	// Query
	res, queryErr := influx.QueryData(
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
