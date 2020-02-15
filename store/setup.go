package store

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// CheckSetup Check if database has default user, org, bucket
func (i *Influx) CheckSetup() (err error) {
	// Form request string
	res, err := i.HTTPInstance.Get(i.HTTPClient, i.GetBasicURL()+"/setup", nil, nil)

	var resBody map[string]interface{}
	json.Unmarshal(res, &resBody)

	if resBody["Allowed"] == "true" {
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

	res, queryErr := i.HTTPInstance.Post(i.HTTPClient, i.GetBasicURL()+"/setup", body, nil)
	if queryErr != nil {
		return
	}

	var resBody map[string]interface{}
	json.Unmarshal(res, &resBody)

	if resBody["code"] == "conflict" {
		logrus.Error("ALREADY setup", res)
		return
	}

	logrus.Info("Setup successfully", res)

	return
}
