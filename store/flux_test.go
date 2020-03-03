package store

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestCreateFluxClient(t *testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "OmtoG5-MWHplbyT0QS2-HoDyfKAUpbYkkXf_W3nYDqwZe631h-NRGygJoEFyUeVxXknTewpOwa97A-q0BCI3eg==",
	}

	err := influx.ConnectToDB()
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(influx.DBClient)
}
