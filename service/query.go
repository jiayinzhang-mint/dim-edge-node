package service

import (
	"context"
	"dim-edge/node/protocol"
	"fmt"
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes"
	influxdb2 "github.com/influxdata/influxdb-client-go"
)

var floatType = reflect.TypeOf(float64(0))

func getFloat(unk interface{}) (float64, error) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}

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

	// init a unknown-length array
	r.Record = make([]*protocol.Record, 0)

	// decode res with flux tags
	for result.Next() {

		value, err := getFloat(result.Record().Value())
		if err != nil {
			return r, err
		}

		// convert into proto format
		ts, _ := ptypes.TimestampProto(result.Record().Time())
		r.Record = append(r.Record, &protocol.Record{
			Time:  ts,
			Count: value,
		})

	}

	return r, err
}

// InsertData insert data
func (g *GRPCServer) InsertData(c context.Context, p *protocol.InsertDataParams) (*protocol.InsertDataRes, error) {
	var (
		m     []*influxdb2.Point
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
		m = append(m, influxdb2.NewPoint(
			x.Name,
			x.Tags,
			fields,
			ts,
		))
	}

	// insert data
	err = g.Influx.InsertData(m, p.Bucket, p.Org)
	if err != nil {
		return r, err
	}

	r.Count = int64(count)

	return r, err
}
