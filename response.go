package httpfake

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response stores the settings defined by the request handler
// of how it will respond the request back
type Response struct {
	StatusCode int
	BodyBuffer []byte
	Header     http.Header
}

// NewResponse creates a new Response
func NewResponse() *Response {
	return &Response{
		Header: make(http.Header),
	}
}

// Status sets the response status
func (r *Response) Status(status int) *Response {
	r.StatusCode = status
	return r
}

// SetHeader sets the a HTTP header to the response
func (r *Response) SetHeader(key, value string) *Response {
	r.Header.Set(key, value)
	return r
}

// AddHeader adds a HTTP header into the response
func (r *Response) AddHeader(key, value string) *Response {
	r.Header.Add(key, value)
	return r
}

// Body sets the response body from a byte array
func (r *Response) Body(body []byte) *Response {
	r.BodyBuffer = body
	return r
}

// BodyString sets the response body from a string
// Example:
//
//	BodyString(`[{"username": "dreamer"}]`)
func (r *Response) BodyString(body string) *Response {
	return r.Body([]byte(body))
}

// BodyStruct sets the response body from a struct.
// The provided struct will be marsheled to json internally.
// Example:
//
//	BodyStruct(&entity.User{UserName: "dreamer"})
func (r *Response) BodyStruct(body interface{}) *Response {
	b, err := json.Marshal(body)
	if err != nil {
		printError(fmt.Sprintf("marshalling body %#v failed with %v", body, err))
	}

	return r.Body(b)
}
