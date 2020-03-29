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
		r   *protocol.QueryRes
		err error
	)

	result, err := g.Influx.QueryData(p.QueryString, p.Org)

	var record *protocol.Record
	for result.Next() {
		mErr := result.Unmarshal(&record)
		if mErr != nil {
			logrus.Error(mErr)
		}

		r.Record = append(r.Record, record)
	}

	return r, err
}

// InsertData insert data
func (g *GRPCServer) InsertData(c context.Context, p *protocol.InsertDataParams) (*protocol.InsertDataRes, error) {
	var (
		m     []influxdb.Metric
		r     *protocol.InsertDataRes
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

		fields := make(map[string]interface{})
		for k, y := range x.Fields {
			fields[k] = interface{}(y)
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

	r = &protocol.InsertDataRes{
		Count: int64(count),
	}

	return r, err
}
