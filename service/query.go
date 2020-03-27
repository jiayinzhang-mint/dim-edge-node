package service

import (
	"context"
	"dim-edge-node/protocol"
	"dim-edge-node/utils"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/influxdata/influxdb-client-go"
)

// QueryData query data
func (g *GRPCServer) QueryData(c context.Context, p *protocol.QueryParams) (*protocol.QueryRes, error) {
	var (
		r   *protocol.QueryRes
		err error
	)

	res, err := g.Influx.QueryData(p.QueryString, p.Org)

	r = &protocol.QueryRes{
		Row:      res.Row,
		ColNames: res.ColNames,
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

		// form metric
		m = append(m, influxdb.NewRowMetric(
			utils.StructToMap(x.Fields),
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
