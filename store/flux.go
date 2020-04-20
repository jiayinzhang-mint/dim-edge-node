package store

import (
	"dim-edge-node/utils"
	"net/http"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/sirupsen/logrus"
)

// Influx db instance
type Influx struct {
	Address      string `json:"address"`
	Token        string `json:"token"`
	DBClient     influxdb2.InfluxDBClient
	HTTPClient   *http.Client        // http client for operation
	HTTPInstance *utils.HTTPInstance // http instance for session store
}

// GetDB return db instance
func (i *Influx) GetDB() influxdb2.InfluxDBClient {
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

	logrus.Info("Influx DB connected")

	// Check setup
	setup, err := i.CheckSetup()
	if err != nil {
		logrus.Error(err)
		return
	}

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
