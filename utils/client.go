package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

//Client for http requests
type Client struct {
	*http.Client
	BaseURL     string
	ContentType string
}

var emptyJSON, _ = json.Marshal(struct{}{})

func (c Client) CallRequest(method string, path string, reader io.Reader) (*http.Response, error) {
	return c.CallRequestWithHeaders(method, path, reader, make(map[string]string))
}

func (c Client) CallRequestNoBody(method string, path string) (*http.Response, error) {
	reader := bytes.NewReader(emptyJSON)
	return c.CallRequestWithHeaders(method, path, reader, make(map[string]string))
}

func (c Client) CallRequestNoBodytWithHeaders(method string, path string, headers map[string]string) (*http.Response, error) {
	reader := bytes.NewReader(emptyJSON)
	return c.CallRequestWithHeaders(method, path, reader, headers)
}

func (c Client) CallRequestWithHeaders(method string, path string, reader io.Reader, headers map[string]string) (*http.Response, error) {
	req, _ := http.NewRequest(method, c.BaseURL+path, reader)
	req.Header.Set("Content-Type", c.ContentType)
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	return c.Do(req)
}
