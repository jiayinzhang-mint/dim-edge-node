package store

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

// CheckSetup Check if database has default user, org, bucket
func (i *Influx) CheckSetup() (msg bool, err error) {
	// Form request strinsg
	res, err := i.HTTPInstance.Get(i.HTTPClient, i.GetBasicURL()+"/setup", nil, nil)
	if err != nil {
		return
	}

	var resBody map[string]interface{}
	json.Unmarshal(res, &resBody)

	if resBody["allowed"] == true {
		logrus.Info("Influx has NOT been setup")
		msg = true
		return
	}

	msg = false
	logrus.Info("Influx has ALREADY been setup")
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

	res, err := i.HTTPInstance.Post(i.HTTPClient, i.GetBasicURL()+"/setup", body, nil)
	if err != nil {
		return
	}

	var resBody map[string]interface{}
	json.Unmarshal(res, &resBody)

	if resBody["code"] == "conflict" {
		logrus.Error("ALREADY setup", resBody)
		err = fmt.Errorf("ALREADY setup")
		return
	}

	logrus.Info("Setup successfully", resBody)

	return
}
