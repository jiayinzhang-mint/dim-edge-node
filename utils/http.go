package utils

import (
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
