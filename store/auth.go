package store

import (
	"encoding/base64"

	"github.com/sirupsen/logrus"
)

// SignIn sign into db
func (i *Influx) SignIn(username string, password string) (err error) {

	// Generate basic auth
	basicAuth := func() string {
		auth := username + ":" + password
		return base64.StdEncoding.EncodeToString([]byte(auth))
	}()

	if _, err = i.HTTPInstance.Post(
		i.HTTPClient, i.GetBasicURL()+"/signin",
		nil,
		map[string]string{
			"Authorization": "Basic " + basicAuth,
		}); err != nil {
		return
	}

	logrus.Info("Influx HTTP Client signed in")

	return
}

// SignOut expire current session
func (i *Influx) SignOut() (err error) {
	if _, err = i.HTTPInstance.Post(i.HTTPClient, i.GetBasicURL()+"/signout", nil, nil); err != nil {
		return
	}

	logrus.Info("Influx HTTP Client signed out")

	return
}
