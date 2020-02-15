package utils

import (
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestPost(*testing.T) {
	c := http.Client{}
	i := &HTTPInstance{}

	res, err := i.Post(&c, "http://127.0.0.1:9999/api/v2/signout", nil, nil)
	logrus.Info(res, err)
}
