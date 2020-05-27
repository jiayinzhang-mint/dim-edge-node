package store

import (
	"context"
	"dim-edge/node/utils"
	"net/http"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	ot "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// Influx db instance
type Influx struct {
	Address      string `json:"address"`
	Token        string `json:"token"`
	DBClient     influxdb2.Client
	HTTPClient   *http.Client        // http client for operation
	HTTPInstance *utils.HTTPInstance // http instance for session store
}

// GetDB return db instance
func (i *Influx) GetDB() influxdb2.Client {
	return i.DBClient
}

// GetBasicURL return basic rest url
func (i *Influx) GetBasicURL() string {
	return i.Address + "/api/v2"
}

// ConnectToDB connect to influxdb
func (i *Influx) ConnectToDB() (err error) {
	// Create http clients
	i.HTTPClient = &http.Client{
		Timeout: 5 * time.Second,
	}

	// Create http instance
	i.HTTPInstance = &utils.HTTPInstance{}

	// Init tracer
	jcfg := jaegercfg.Configuration{
		ServiceName: "dim-edge-influxdb",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	i.HTTPInstance.Tracer, i.HTTPInstance.TraceCloser, err = jcfg.NewTracer(
		jaegercfg.Logger(jaeger.StdLogger),
	)
	if err != nil {
		return
	}

	ot.SetGlobalTracer(i.HTTPInstance.Tracer)

	// Check setup
	setup, err := i.CheckSetup(context.TODO())
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Info("Influx DB connected")
	logrus.Info("Influx setup status: ", setup)

	return
}

// CreateDBClient create db native client
func (i *Influx) CreateDBClient() (err error) {
	// Create db clients
	i.DBClient = influxdb2.NewClient(i.Address, i.Token)

	if err != nil {
		logrus.Error(err)
		return
	}

	return
}
