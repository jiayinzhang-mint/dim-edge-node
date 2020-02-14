package store

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSignIn(t *testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "4oXjSoIuU1F3A1zu-xYp0eJ9q_vsLQmtDPPTNuDnrs7R7H7qGAQ1GNaX4hNtJKx5ZRfnoj_TW5Uwe5NJUBvLOA==",
	}

	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	if authErr := influx.SignIn("mint", "131001250115zHzH"); authErr != nil {
		logrus.Error(authErr)
	}
}

func TestSignOut(t *testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "4oXjSoIuU1F3A1zu-xYp0eJ9q_vsLQmtDPPTNuDnrs7R7H7qGAQ1GNaX4hNtJKx5ZRfnoj_TW5Uwe5NJUBvLOA==",
	}

	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	if authErr := influx.SignOut(); authErr != nil {
		logrus.Error(authErr)
	}
}
