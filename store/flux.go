package store

import (
	"net/http"

	"github.com/influxdata/influxdb-client-go"
)

// Influx db instance
type Influx struct {
	Address  string `json:"address"`
	Token    string `json:"token"`
	DBClient *influxdb.Client
}

// GetDB return db instance
func (i *Influx) GetDB() *influxdb.Client {
	return i.DBClient
}

// ConnectToDB connect to influxdb
func (i *Influx) ConnectToDB() (err error) {
	myHTTPClient := &http.Client{}

	i.DBClient, err = influxdb.New(i.Address, i.Token, influxdb.WithHTTPClient(myHTTPClient))

	if err != nil {
		return
	}

	return
}
