package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

var Default = 0

//Client for http requests
type Client struct {
	*http.Client
	BaseURL     string
	ContentType string
}

//Response wrapper
type Response struct {
	Resp *http.Response
	Err  error
}

//InfoExec for the MapExec
type InfoExec struct {
	Interface interface{}
	F         func(*http.Response) error
}

//MapExec associates status code with InfoExec
type MapExec map[int]InfoExec

var emptyJSON, _ = json.Marshal(struct{}{})

//CallRequest with body
func (c Client) CallRequest(method string, path string, reader io.Reader) *Response {
	return c.CallRequestWithHeaders(method, path, reader, make(map[string]string))
}

//CallRequestNoBody without body
func (c Client) CallRequestNoBody(method string, path string) *Response {
	reader := bytes.NewReader(emptyJSON)
	return c.CallRequestWithHeaders(method, path, reader, make(map[string]string))
}

//CallRequestNoBodytWithHeaders without body and with headers
func (c Client) CallRequestNoBodytWithHeaders(method string, path string, headers map[string]string) *Response {
	reader := bytes.NewReader(emptyJSON)
	return c.CallRequestWithHeaders(method, path, reader, headers)
}

//CallRequestWithHeaders with headers
func (c Client) CallRequestWithHeaders(method string, path string, reader io.Reader, headers map[string]string) *Response {
	req, _ := http.NewRequest(method, c.BaseURL+path, reader)
	req.Header.Set("Content-Type", c.ContentType)
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	resp, err := c.Do(req)
	return &Response{resp, err}
}

//WithResponse Extracts response
func (r *Response) WithResponse(f func(*http.Response) error) error {
	if r.Err != nil {
		return r.Err
	}
	defer r.Resp.Body.Close()
	return f(r.Resp)
}

//Solve with status codes
func (r *Response) Solve(mapExec MapExec) error {
	if r.Err != nil {
		return r.Err
	}
	if val, ok := mapExec[r.Resp.StatusCode]; ok {
		if val.Interface != nil {
			GetBodyJSON(r.Resp, val.Interface)
		}
		return val.F(r.Resp)
	}
	if val, ok := mapExec[0]; ok {
		return val.F(r.Resp)
	}
	return errors.New("Status key not found")
}

//GetBodyJSON Gets json form body
func GetBodyJSON(resp *http.Response, i interface{}) {
	if jsonDataFromHTTP, err := ioutil.ReadAll(resp.Body); err == nil {
		if err := json.Unmarshal([]byte(jsonDataFromHTTP), &i); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}
