package store

import (
	"dim-edge-node/utils"

	"github.com/sirupsen/logrus"
)

// CheckSetup Check if database has default user, org, bucket
func (i *Influx) CheckSetup() (err error) {
	// Form request string
	res, err := utils.HTTP().Get(i.HTTPClient, i.GetBasicURL()+"/setup", nil, nil)

	if res["Allowed"] == "true" {
		logrus.Error("Influx has NOT been setup")
		return
	}

	logrus.Error("Influx has ALREADY been setup")
	return
}

// Setup Set up initial user, org and bucket
func (i *Influx) Setup(username string, password string, org string, bucket string, retentionPeriodHrs int) (err error) {
	body := make(map[string]interface{})
	body["username"] = username
	body["password"] = password
	body["org"] = org
	body["bucket"] = bucket
	body["retentionPeriodHrs"] = retentionPeriodHrs

	res, queryErr := utils.HTTP().Post(i.HTTPClient, i.GetBasicURL()+"/setup", body, nil)
	if queryErr != nil {
		return
	}

	if res["code"] == "conflict" {
		logrus.Error("ALREADY setup", res)
		return
	}

	logrus.Info("Setup successfully", res)

	return
}
