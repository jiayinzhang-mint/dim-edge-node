package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// HTTPInstance http instance
type HTTPInstance struct {
	Cookies []*http.Cookie
}

// Get do GET request
func (h *HTTPInstance) Get(c *http.Client, url string, params map[string]string, header map[string]string) ([]byte, error) {
	// Form request string
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Query params
	q := req.URL.Query()
	for i := range params {
		if params[i] != "" {
			q.Add(i, params[i])
		}

	}
	req.URL.RawQuery = q.Encode()

	// Set header
	for i := range header {
		req.Header.Set(i, header[i])
	}

	// Set cookies
	for c := range h.Cookies {
		req.AddCookie(h.Cookies[c])
	}

	// Send request
	res, resErr := c.Do(req)
	if resErr != nil {
		return nil, resErr
	}
	defer res.Body.Close()

	// Check status code
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted && res.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("HTTP GET %s %d", req.URL.String(), res.StatusCode)
	}

	logrus.Info("HTTP GET ", req.URL.String(), " ", res.StatusCode)

	// Read body
	resJSON, _ := ioutil.ReadAll(res.Body)

	return resJSON, nil
}

// Post do POST request
func (h *HTTPInstance) Post(c *http.Client, url string, body map[string]interface{}, header map[string]string) ([]byte, error) {
	bodyJSON, _ := json.Marshal(body)

	// Form request string
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, err
	}

	// Set header
	for i := range header {
		req.Header.Set(i, header[i])
	}

	// Set cookies
	for c := range h.Cookies {
		req.AddCookie(h.Cookies[c])
	}

	// Send request
	res, resErr := c.Do(req)
	if resErr != nil {
		return nil, resErr
	}
	defer res.Body.Close()

	// Check status code
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted && res.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("HTTP POST %s %d", req.URL.String(), res.StatusCode)
	}

	logrus.Info("HTTP POST ", req.URL.String(), " ", res.StatusCode)

	// Set cookies
	h.Cookies = res.Cookies()

	// Read body
	resJSON, _ := ioutil.ReadAll(res.Body)

	return resJSON, nil
}
