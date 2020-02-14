package store

import (
	"net/http"

	"github.com/influxdata/influxdb-client-go"
	"github.com/sirupsen/logrus"
)

// Influx db instance
type Influx struct {
	Address    string `json:"address"`
	Token      string `json:"token"`
	DBClient   *influxdb.Client
	HTTPClient *http.Client // http client for operation
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
	i.HTTPClient = &http.Client{}

	// Create db client
	i.DBClient, err = influxdb.New(i.Address, i.Token, influxdb.WithHTTPClient(i.HTTPClient))

	if err != nil {
		return
	}

	logrus.Info("Influx DB connected")

	return
}
