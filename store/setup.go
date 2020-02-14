package store

import (
	"dim-edge-node/utils"

	"github.com/sirupsen/logrus"
)

// CheckSetup Check if database has default user, org, bucket
func (i *Influx) CheckSetup() (err error) {
	// Form request string
	res, err := utils.HTTP().Get(i.GetBasicURL() + "/setup")

	if res["Allowed"] == "true" {
		logrus.Error("Influx has NOT been setup")
		return
	}

	logrus.Error("Influx has ALREADY been setup")
	return
}

// Setup Set up initial user, org and bucket
func (i *Influx) Setup() (err error) {
	return
}
