package fda

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseUrl = "https://api.fda.gov/"
	DevicePath     = "device/"
	Udi            = "udi.json"
	F510k          = "510k.json"
)

type Client struct {
	client  *http.Client
	baseUrl *url.URL
	key     string

	FDA510kService *FDA510kService
}

type QueryParameters struct {
	Search string `url:"search,omitempty" json:"search,omitempty"`
	Sort   string `url:"sort,omitempty" json:"sort,omitempty"`
	Count  string `url:"count,omitempty" json:"count,omitempty"`
	Limit  string `url:"limit,omitempty" json:"limit,omitempty"`
	Skip   string `url:"skip,omitempty" json:"skip,omitempty"`
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

	client.FDA510kService = &FDA510kService{client: client}

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

func (c *Client) NewRequest(method, path string, opt interface{}) (*http.Request, error) {
	u := *c.baseUrl
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}
	u.RawPath = c.baseUrl.Path + path
	u.Path = c.baseUrl.Path + unescaped
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	if opt != nil {
		v := reflect.ValueOf(opt)
		t := v.Type()
		if t.Kind() == reflect.Ptr {
			v = v.Elem()
			t = v.Type()
		}

		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}

		searchParts := []string{}
		baseQuery := make(url.Values)

		baseType := reflect.TypeOf(QueryParameters{})
		baseParams := map[string]bool{}
		for i := 0; i < baseType.NumField(); i++ {
			field := baseType.Field(i)
			if tag := field.Tag.Get("url"); tag != "" {
				paramName := strings.Split(tag, ",")[0]
				baseParams[paramName] = true
				if val := q.Get(paramName); val != "" {
					baseQuery.Set(paramName, val)
				}
			}
		}

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if tag := field.Tag.Get("json"); tag != "" && !baseParams[strings.Split(tag, ",")[0]] {
				paramName := strings.Split(tag, ",")[0]
				fieldValue := v.Field(i)

				if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
					fieldValue = fieldValue.Elem()
				}

				if !fieldValue.IsZero() {
					searchParts = append(searchParts, paramName+":"+fieldValue.String())
				}
			}
		}

		queryParts := []string{}

		for key, values := range baseQuery {
			if key != "search" {
				for _, value := range values {
					queryParts = append(queryParts, url.QueryEscape(key)+"="+url.QueryEscape(value))
				}
			}
		}

		if len(searchParts) > 0 {
			searchQuery := strings.Join(searchParts, " AND ")
			if existing := baseQuery.Get("search"); existing != "" {
				searchQuery = existing + " AND " + searchQuery
			}
			escapedParts := []string{}
			for _, part := range strings.Split(searchQuery, ":") {
				escapedParts = append(escapedParts, url.QueryEscape(part))
			}
			searchParam := "search=" + strings.Join(escapedParts, ":")
			queryParts = append(queryParts, searchParam)
		}

		u.RawQuery = strings.Join(queryParts, "&")
		fmt.Printf("final query %s\n", u.String())
	}

	req, err := http.NewRequest(method, u.String(), nil)
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

	var apiKey string

	if c.key != "" {
		apiKey = c.key
		req.Header.Set("Authorization", "Key "+apiKey) // TODO make to be the correct header text for fda
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)

	response := newResponse(resp)

	if w != nil {
		if x, ok := w.(io.Writer); ok {
			_, err = io.Copy(x, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(w)
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
