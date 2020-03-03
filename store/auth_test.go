package store

import (
	"dim-edge-node/protocol"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSignIn(t *testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "OmtoG5-MWHplbyT0QS2-HoDyfKAUpbYkkXf_W3nYDqwZe631h-NRGygJoEFyUeVxXknTewpOwa97A-q0BCI3eg==",
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
		Token:   "OmtoG5-MWHplbyT0QS2-HoDyfKAUpbYkkXf_W3nYDqwZe631h-NRGygJoEFyUeVxXknTewpOwa97A-q0BCI3eg==",
	}

	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	if authErr := influx.SignOut(); authErr != nil {
		logrus.Error(authErr)
	}
}

func TestListAuthorization(*testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "OmtoG5-MWHplbyT0QS2-HoDyfKAUpbYkkXf_W3nYDqwZe631h-NRGygJoEFyUeVxXknTewpOwa97A-q0BCI3eg==",
	}

	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	// Sign in first
	lErr := influx.SignIn("mint", "131001250115zHzH")
	if lErr != nil {
		logrus.Error(lErr)
	}

	res, err := influx.ListAuthorization("", "mint", "", "INSDIM")
	logrus.Info(res, err)
}

func TestCreateAuthorization(*testing.T) {
	influx := &Influx{
		Address: "http://127.0.0.1:9999",
		Token:   "OmtoG5-MWHplbyT0QS2-HoDyfKAUpbYkkXf_W3nYDqwZe631h-NRGygJoEFyUeVxXknTewpOwa97A-q0BCI3eg==",
	}

	if err := influx.ConnectToDB(); err != nil {
		logrus.Error(err)
	}

	// Sign in first
	lErr := influx.SignIn("mint", "131001250115zHzH")
	if lErr != nil {
		logrus.Error(lErr)
	}

	var p []*protocol.Authorization_Permission
	p = append(p, &protocol.Authorization_Permission{
		Action: "read",
		Resource: &protocol.Authorization_Permission_Resource{
			Type: "buckets",
		},
	})
	res, err := influx.CreateAuthorization("active", "testbiu", "a9a72b93088f4c8f", p)
	logrus.Info(res, err)
}
