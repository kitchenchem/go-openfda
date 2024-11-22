package fda

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
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

func (c *Client) NewRequest(path string) (*http.Request, error) {
	urlBuild := *c.baseUrl
	nonescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	urlBuild.RawPath = c.baseUrl.Path + path
	urlBuild.Path = c.baseUrl.Path + nonescaped

	reqHeader := make(http.Header)
	reqHeader.Set("Accept", "application/json")

	req, err := http.NewRequest("GET", urlBuild.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) { //TODO
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
