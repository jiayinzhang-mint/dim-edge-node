package store

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestCheckSteup(t *testing.T) {
	influx := &Influx{
		Address: "http://192.168.64.18:31048",
	}

	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	msg, err := influx.CheckSetup()
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info(msg)
}

func TestSetup(t *testing.T) {
	influx := &Influx{
		Address: "http://192.168.64.18:31048",
	}

	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	setupErr := influx.Setup("mint", "131001250115zHzH", "INSDIM", "INSDIM", 0)
	if setupErr != nil {
		logrus.Error(setupErr)
	}
}
