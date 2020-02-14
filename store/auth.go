package store

import (
	"encoding/base64"
	"net/http"

	"github.com/sirupsen/logrus"
)

// SignIn sign into db
func (i *Influx) SignIn(username string, password string) (err error) {

	// Generate basic auth
	basicAuth := func() string {
		auth := username + ":" + password
		return base64.StdEncoding.EncodeToString([]byte(auth))
	}()

	// Form request string
	req, err := http.NewRequest("POST", i.GetBasicURL()+"/signin", nil)
	if err != nil {
		return
	}

	// Basic auth
	req.Header.Add("Authorization", "Basic "+basicAuth)

	// Send request
	if _, err = i.HTTPClient.Do(req); err != nil {
		return
	}

	logrus.Info("Influx HTTP Client signed in")

	return
}

// SignOut expire current session
func (i *Influx) SignOut() error {

	// Form request string
	req, err := http.NewRequest("POST", i.GetBasicURL()+"/signout", nil)
	if err != nil {
		return err
	}

	// Send request
	if _, err = i.HTTPClient.Do(req); err != nil {
		return err
	}

	logrus.Info("Influx HTTP Client signed out")

	return err
}
