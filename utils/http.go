package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
	"github.com/sirupsen/logrus"
)

// HTTPInstance http instance
type HTTPInstance struct {
	Cookies     []*http.Cookie
	Tracer      opentracing.Tracer
	TraceCloser io.Closer
}

// Get do GET request
func (h *HTTPInstance) Get(ctx context.Context, c *http.Client, url string, params map[string]string, header map[string]string) ([]byte, error) {
	var parentCtx opentracing.SpanContext

	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan != nil {
		parentCtx = parentSpan.Context()
	}

	span := opentracing.StartSpan(
		url,
		opentracing.ChildOf(parentCtx),
	)
	defer span.Finish()

	// Form request string
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Implement tracer context
	req = req.WithContext(opentracing.ContextWithSpan(context.Background(), span))

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
		span.SetTag(string(ext.Error), true)
		span.LogKV(otlog.Error(err))
		return nil, resErr
	}
	defer res.Body.Close()

	logrus.Info("HTTP GET ", req.URL.String(), " ", res.StatusCode)

	// Read body
	resJSON, _ := ioutil.ReadAll(res.Body)

	// Check status code
	if res.StatusCode >= 400 {
		return resJSON, fmt.Errorf("HTTP GET %s %d", req.URL.String(), res.StatusCode)
	}

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

	logrus.Info("HTTP POST ", req.URL.String(), " ", res.StatusCode)

	// Get and store cookies
	h.Cookies = res.Cookies()

	// Read body
	resJSON, _ := ioutil.ReadAll(res.Body)

	// Check status code
	if res.StatusCode >= 400 {
		return resJSON, fmt.Errorf("HTTP POST %s %d", req.URL.String(), res.StatusCode)
	}

	return resJSON, nil
}

// Delete do DELETE request
func (h *HTTPInstance) Delete(c *http.Client, url string, params map[string]string, header map[string]string) ([]byte, error) {
	// Form request string
	req, err := http.NewRequest("DELETE", url, nil)
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

	// Send request
	res, resErr := c.Do(req)
	if resErr != nil {
		return nil, resErr
	}
	defer res.Body.Close()

	logrus.Info("HTTP DELETE ", req.URL.String(), " ", res.StatusCode)

	// Read body
	resJSON, _ := ioutil.ReadAll(res.Body)

	// Check status code
	if res.StatusCode >= 400 {
		return resJSON, fmt.Errorf("HTTP DELETE %s %d", req.URL.String(), res.StatusCode)
	}

	return resJSON, nil
}
