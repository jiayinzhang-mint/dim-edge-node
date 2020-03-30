package service

import (
	"context"
	"dim-edge-node/protocol"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/influxdata/influxdb-client-go"
	"github.com/sirupsen/logrus"
)

// QueryData query data
func (g *GRPCServer) QueryData(c context.Context, p *protocol.QueryParams) (*protocol.QueryRes, error) {
	var (
		r   = &protocol.QueryRes{}
		err error
	)

	result, err := g.Influx.QueryData(p.QueryString, p.Org)
	if err != nil {
		return r, err
	}

	type influxRecord struct {
		Zone   ***string `flux:"name" json:"zone"`
		Stop   time.Time `flux:"_stop" json:"-"`
		Start  time.Time `flux:"_start" json:"-"`
		Time   time.Time `flux:"_time" json:"date"`
		HostIP string    `flux:"host_ip" json:"-"`
		Count  float64   `flux:"_value" json:"count"`
	}

	r.Record = make([]*protocol.Record, 0)

	var rec influxRecord
	for result.Next() {
		mErr := result.Unmarshal(&rec)
		logrus.Info(rec)

		ts, _ := ptypes.TimestampProto(rec.Time)
		r.Record = append(r.Record, &protocol.Record{
			Time:  ts,
			Count: rec.Count,
		})
		if mErr != nil {
			logrus.Error(mErr)
		}
	}

	return r, err
}

// InsertData insert data
func (g *GRPCServer) InsertData(c context.Context, p *protocol.InsertDataParams) (*protocol.InsertDataRes, error) {
	var (
		m     []influxdb.Metric
		r     = &protocol.InsertDataRes{}
		err   error
		ts    time.Time
		count int
	)

	for _, x := range p.Metrics {
		// parse time
		ts, err = ptypes.Timestamp(x.Ts)
		if err != nil {
			return r, err
		}

		// convert map[string]float64 to map[string]interface
		fields := make(map[string]interface{})

		for y, f := range x.Fields {
			fields[y] = interface{}(f)
		}

		// form metric
		m = append(m, influxdb.NewRowMetric(
			fields,
			x.Name,
			x.Tags,
			ts,
		))
	}

	// insert data
	count, err = g.Influx.InsertData(&m, p.Bucket, p.Org)
	if err != nil {
		return r, err
	}

	r.Count = int64(count)

	return r, err
}
