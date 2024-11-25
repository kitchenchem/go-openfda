package fda

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseUrl = "https://api.fda.gov/"
	search         = "search"
	sort           = "sort"
	count          = "count"
	limit          = "limit"
	skip           = "skip"
)

type Client struct {
	client  *http.Client
	baseUrl *url.URL
	key     string

	// Devices *DeviceServic TODO
}

type RateLimiter interface {
	Wait(context.Context) error
}

func NewClient(apiKey string) (*Client, error) {
	client := &Client{key: apiKey}

	// Create an http.Client with some sensible wait defaults for various responses
	client.client = &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 10 * time.Second,

			DialContext: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).DialContext,

			IdleConnTimeout:       45 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   10,
		},
	}

	client.setBaseURL(defaultBaseUrl)

	return client, nil
}

func (c *Client) setBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	c.baseUrl = baseURL
	return nil
}

// func (c *Client) NewRequest(path string) (*http.Request, error) {
// 	urlBuild := *c.baseUrl
// 	nonescaped, err := url.PathUnescape(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	urlBuild.RawPath = c.baseUrl.Path + path
// 	fmt.Printf("raw: %s\n", urlBuild.RawPath)
// 	urlBuild.Path = c.baseUrl.Path + nonescaped
// 	fmt.Printf("path: %s\n", urlBuild.Path)
// 	fmt.Printf("str: %s\n", urlBuild.String())

// 	fmt.Printf("url built: %s      --------", &urlBuild)

// 	reqHeaders := make(http.Header)
// 	reqHeaders.Set("Accept", "application/json")

// 	req, err := http.NewRequest("GET", urlBuild.String(), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for k, v := range reqHeaders {
// 		req.Header[k] = v
// 	}

// 	return req, nil
// }

func (c *Client) NewRequest(path string, opt interface{}) (*http.Request, error) {
	urlBuild := *c.baseUrl

	pathAndQuery := strings.SplitN(path, "?", 2)
	basePath := pathAndQuery[0]

	nonescaped, err := url.PathUnescape(basePath)
	if err != nil {
		return nil, err
	}
	urlBuild.Path = c.baseUrl.Path + nonescaped

	if len(pathAndQuery) > 1 {
		urlBuild.RawQuery = pathAndQuery[1]
	}

	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	fmt.Printf("str: %s\n", urlBuild.String())

	// opt represents options available to en endpoint //TODO
	if opt != nil {
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		urlBuild.RawQuery = q.Encode()
		fmt.Printf("url string: %s", urlBuild.String())
	}

	req, err := http.NewRequest("GET", urlBuild.String(), nil)
	if err != nil {
		return nil, err
	}

	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *Client) Do(req *http.Request, w interface{}) (*Response, error) {

	// if w implements io writer interface for the response body to be written to it without trying to first decode it.
	var apiKey string

	if c.key != "" {
		apiKey = c.key
		req.Header.Set("Authorization", "Key "+apiKey) // TODO make to be the correct header text for fda
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	//TODO handling for daily limit breach w/ and w/out token.

	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)

	//TODO make response type, methods.

	response := newResponse(resp)

	if w != nil {
		if x, ok := w.(io.Writer); ok {
			_, err = io.Copy(x, resp.Body) // need original response, not Response type
		} else {
			err = json.NewDecoder(resp.Body).Decode(w) //same here
		}
	}

	return response, err
}

type ErrorResponse struct {
	Body     []byte
	Response *http.Response
	Message  string
}

func (e *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(e.Response.Request.URL.Path)
	url := fmt.Sprintf("%s://%s%s", e.Response.Request.URL.Scheme, e.Response.Request.URL.Host, path)

	if e.Message == "" {
		return fmt.Sprintf("%s %s: %d", e.Response.Request.Method, url, e.Response.StatusCode)
	} else {
		return fmt.Sprintf("%s %s: %d %s", e.Response.Request.Method, url, e.Response.StatusCode, e.Message)
	}
}

func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	case 404:
		return ErrNotFound
	case 429:
		return ErrTooManyRequests
	}

	errorResponse := &ErrorResponse{Response: r}

	data, err := io.ReadAll(r.Body)
	if err == nil && strings.TrimSpace(string(data)) != "" {
		errorResponse.Body = data

		var raw interface{}
		if err := json.Unmarshal(data, &raw); err != nil {
			errorResponse.Message = fmt.Sprintf("failed to parse unknown error format: %s", data)
		} else {

			return fmt.Errorf("failed to parse unexpected error type: %T", raw)
		}
	}

	return errorResponse
}
