package utils

import (
	"io"
	"net/http"
)

//Client for http requests
type Client struct {
	*http.Client
	BaseURL     string
	ContentType string
}

func (c Client) CallRequest(method string, path string, reader io.Reader) (*http.Response, error) {
	return c.CallRequestWithHeaders(method, path, reader, make(map[string]string))
}

func (c Client) CallRequestWithHeaders(method string, path string, reader io.Reader, headers map[string]string) (*http.Response, error) {
	req, _ := http.NewRequest(method, c.BaseURL+path, reader)
	req.Header.Set("Content-Type", c.ContentType)
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	return c.Do(req)
}
