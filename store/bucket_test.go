package store

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetBucketList(*testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "OmtoG5-MWHplbyT0QS2-HoDyfKAUpbYkkXf_W3nYDqwZe631h-NRGygJoEFyUeVxXknTewpOwa97A-q0BCI3eg==",
	}

	err := influx.ConnectToDB()
	if err != nil {
		logrus.Error(err)
	}

	// Sign in first
	lErr := influx.SignIn("mint", "131001250115zHzH")
	if lErr != nil {
		logrus.Error(lErr)
	}

	b, qErr := influx.ListAllBucket(1, 10, "INSDIM", "", "")
	if qErr != nil {
		logrus.Error(qErr)
	}

	logrus.Info(b)
}

func TestGetBucket(*testing.T) {

	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "OmtoG5-MWHplbyT0QS2-HoDyfKAUpbYkkXf_W3nYDqwZe631h-NRGygJoEFyUeVxXknTewpOwa97A-q0BCI3eg==",
	}

	err := influx.ConnectToDB()
	if err != nil {
		logrus.Error(err)
	}

	// Sign in first
	lErr := influx.SignIn("mint", "131001250115zHzH")
	if lErr != nil {
		logrus.Error(lErr)
	}

	b, qErr := influx.RetrieveBucket("7b33a28bfd4be452")
	if qErr != nil {
		logrus.Error(qErr)
	}

	l, qErr := influx.RetrieveBucketLog("7b33a28bfd4be452", 1, 20)
	if qErr != nil {
		logrus.Error(qErr)
	}

	logrus.Info(b)
	logrus.Info(l)
}
