package store

import (
	"context"
	"dim-edge/node/protocol"
	"dim-edge/node/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

// SignIn sign into db
func (i *Influx) SignIn(username string, password string) (err error) {
	var (
		basicAuth string
		auth      []*protocol.Authorization
	)

	// Generate basic auth
	basicAuth = func() string {
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

	if auth, err = i.ListAuthorization("", username, "", ""); err != nil {
		logrus.Error("Faield to get authorization")
	}

	if len(auth) >= 1 {
		i.Token = auth[0].GetToken()
	}

	if i.DBClient == nil {
		if err = i.CreateDBClient(); err != nil {
			logrus.Error("Failed to create native influx client")
			return
		}
		logrus.Info("Native Influx client created")
	}

	logrus.Info("Influx user token cached")

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

// ListAuthorization list all authorizations
func (i *Influx) ListAuthorization(userID string, user string, orgID string, org string) (auth []*protocol.Authorization, err error) {
	res, err := i.HTTPInstance.Get(context.TODO(), i.HTTPClient, i.GetBasicURL()+"/authorizations", map[string]string{
		"user":   user,
		"userID": userID,
		"orgID":  orgID,
		"org":    org,
	}, nil)
	if err != nil {
		return
	}

	type a struct {
		Authorizations []*protocol.Authorization `json:"authorizations"`
	}
	var resBody a
	err = json.Unmarshal(res, &resBody)

	auth = resBody.Authorizations
	return
}

// RetrieveAuthorization retrive one authorization with authID
func (i *Influx) RetrieveAuthorization(authID string) (auth *protocol.Authorization, err error) {
	res, err := i.HTTPInstance.Get(context.TODO(), i.HTTPClient, i.GetBasicURL()+"/authorizations", map[string]string{
		"authID": authID,
	}, nil)

	err = json.Unmarshal(res, &auth)

	return
}

// CreateAuthorization create authorization
func (i *Influx) CreateAuthorization(status string, description string, orgID string, p []*protocol.Authorization_Permission) (auth *protocol.Authorization, err error) {
	a := protocol.Authorization{
		Status:      status,
		Description: description,
		OrgID:       orgID,
		Permissions: p,
	}
	ai := utils.StructToMap(a)

	reqBody := make(map[string]interface{})
	reqBody["status"] = status
	reqBody["description"] = description
	reqBody["orgID"] = orgID
	reqBody["permissions"] = ai["permissions"]

	res, err := i.HTTPInstance.Post(i.HTTPClient, i.GetBasicURL()+"/authorizations", reqBody, nil)
	if err != nil {
		var b protocol.OpRes
		json.Unmarshal(res, &b)
		err = fmt.Errorf(b.Code, b.Message)
		return
	}

	err = json.Unmarshal(res, &auth)

	return
}

// GetMe get my info after signing in
func (i *Influx) GetMe() (me *protocol.Me, err error) {
	res, err := i.HTTPInstance.Get(context.TODO(), i.HTTPClient, i.GetBasicURL()+"/me", map[string]string{}, nil)

	err = json.Unmarshal(res, &me)

	return
}
