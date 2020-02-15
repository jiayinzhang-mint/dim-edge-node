package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// HTTPInstance http instance
type HTTPInstance struct {
}

var httpInstance *HTTPInstance

// HTTP return http instance
func HTTP() *HTTPInstance {
	return httpInstance
}

// Get do GET request
func (*HTTPInstance) Get(c *http.Client, url string, params map[string]string, header map[string]string) (map[string]interface{}, error) {
	// Form request string
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	// Query params
	q := req.URL.Query()
	for i := range params {
		q.Add(i, params[i])
	}
	req.URL.Query().Encode()

	// Set header
	for i := range header {
		req.Header.Set(i, header[i])
	}

	// Send request
	res, resErr := c.Do(req)
	if resErr != nil {
		return nil, resErr
	}
	defer res.Body.Close()

	// Read body
	var resBody map[string]interface{}
	resJSON, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(resJSON, &resBody)

	return resBody, nil
}

// Post do POST request
func (*HTTPInstance) Post(c *http.Client, url string, body map[string]interface{}, header map[string]string) (map[string]interface{}, error) {
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

	// Send request
	res, resErr := c.Do(req)
	if resErr != nil {
		return nil, resErr
	}
	defer res.Body.Close()

	// Read body
	var resBody map[string]interface{}
	resJSON, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(resJSON, &resBody)

	return resBody, nil
}
