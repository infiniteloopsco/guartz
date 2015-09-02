package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

//Client for http requests
type Client struct {
	*http.Client
	BaseURL     string
	ContentType string
}

type Response struct {
	resp *http.Response
	err  error
}

var emptyJSON, _ = json.Marshal(struct{}{})

func (c Client) CallRequest(method string, path string, reader io.Reader) *Response {
	return c.CallRequestWithHeaders(method, path, reader, make(map[string]string))
}

func (c Client) CallRequestNoBody(method string, path string) *Response {
	reader := bytes.NewReader(emptyJSON)
	return c.CallRequestWithHeaders(method, path, reader, make(map[string]string))
}

func (c Client) CallRequestNoBodytWithHeaders(method string, path string, headers map[string]string) *Response {
	reader := bytes.NewReader(emptyJSON)
	return c.CallRequestWithHeaders(method, path, reader, headers)
}

func (c Client) CallRequestWithHeaders(method string, path string, reader io.Reader, headers map[string]string) *Response {
	req, _ := http.NewRequest(method, c.BaseURL+path, reader)
	req.Header.Set("Content-Type", c.ContentType)
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	resp, err := c.Do(req)
	return &Response{resp, err}
}

func (r *Response) WithResponseJSON(i interface{}, f func(*http.Response) error) error {
	if r.err != nil {
		return r.err
	}
	defer r.resp.Body.Close()
	if i != nil {
		GetBodyJSON(r.resp, i)
	}
	return f(r.resp)
}

func (r *Response) WithResponse(f func(*http.Response) error) error {
	if r.err != nil {
		return r.err
	}
	defer r.resp.Body.Close()
	return f(r.resp)
}

func GetBodyJSON(resp *http.Response, i interface{}) {
	if jsonDataFromHTTP, err := ioutil.ReadAll(resp.Body); err == nil {
		if err := json.Unmarshal([]byte(jsonDataFromHTTP), &i); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}
