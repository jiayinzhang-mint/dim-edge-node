package store

import (
	"dim-edge-node/utils"
	"net/http"
	"time"

	"github.com/influxdata/influxdb-client-go"
	"github.com/sirupsen/logrus"
)

// Influx db instance
type Influx struct {
	Address      string `json:"address"`
	Token        string `json:"token"`
	DBClient     *influxdb.Client
	HTTPClient   *http.Client        // http client for operation
	HTTPInstance *utils.HTTPInstance // http instance for session store
}

// GetDB return db instance
func (i *Influx) GetDB() *influxdb.Client {
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

	// Create db clients
	i.DBClient, err = influxdb.New(i.Address, i.Token, influxdb.WithHTTPClient(i.HTTPClient))

	if err != nil {
		return
	}

	logrus.Info("Influx DB connected")

	return
}
