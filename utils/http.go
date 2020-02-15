package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
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
func (*HTTPInstance) Get(url string) (map[string]interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP get error: uri=%v , statusCode=%v", url, response.StatusCode)
	}

	// Ready body
	resJSON, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	mErr := json.Unmarshal(resJSON, &res)
	if mErr != nil {
		return nil, err
	}

	return res, nil
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
