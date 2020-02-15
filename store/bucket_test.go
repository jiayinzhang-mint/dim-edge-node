package store

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetBucketList(*testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "4oXjSoIuU1F3A1zu-xYp0eJ9q_vsLQmtDPPTNuDnrs7R7H7qGAQ1GNaX4hNtJKx5ZRfnoj_TW5Uwe5NJUBvLOA==",
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
		Token:   "4oXjSoIuU1F3A1zu-xYp0eJ9q_vsLQmtDPPTNuDnrs7R7H7qGAQ1GNaX4hNtJKx5ZRfnoj_TW5Uwe5NJUBvLOA==",
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
