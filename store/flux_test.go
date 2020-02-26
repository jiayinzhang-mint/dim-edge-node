package store

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestCreateFluxClient(t *testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "4oXjSoIuU1F3A1zu-xYp0eJ9q_vsLQmtDPPTNuDnrs7R7H7qGAQ1GNaX4hNtJKx5ZRfnoj_TW5Uwe5NJUBvLOA==",
	}

	err := influx.ConnectToDB()
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(influx.DBClient)
}
