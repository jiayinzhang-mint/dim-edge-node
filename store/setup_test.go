package store

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestCheckSteup(t *testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "OmtoG5-MWHplbyT0QS2-HoDyfKAUpbYkkXf_W3nYDqwZe631h-NRGygJoEFyUeVxXknTewpOwa97A-q0BCI3eg==",
	}

	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	if setup := influx.CheckSetup(); setup != nil {
		logrus.Error(setup)
	}
}

func TestSetup(t *testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "OmtoG5-MWHplbyT0QS2-HoDyfKAUpbYkkXf_W3nYDqwZe631h-NRGygJoEFyUeVxXknTewpOwa97A-q0BCI3eg==",
	}

	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	setupErr := influx.Setup("mint", "131001250115zHzH", "INSDIM", "INSDIM", 0)
	if setupErr != nil {
		logrus.Error(setupErr)
	}
}
