package fda

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-cleanhttp"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"golang.org/x/time/rate"
)

// endpoint.json seems to be the only option available.
const (
	defaultBaseUrl       = "https://api.fda.gov/"
	userAgent            = "go-openfda"
	devicePath           = "device/"
	udiRoute             = "udi.json"
	fda510kRoute         = "510k.json"
	classificationRoute  = "classification.json"
	covid19SerologyRoute = "covid19serology.json"
	enforcementRoute     = "enforcement.json"
	eventRoute           = "event.json"
	pmaRoute             = "pma.json"
	recallRoute          = "recall.json"
	reglistRoute         = "reglist.json"

	headerRateLimit = "RateLimit-Limit"
	headerRateReset = "RateLimit-Reset"
)

type Client struct {
	client  *retryablehttp.Client
	baseUrl *url.URL
	// https://open.fda.gov/apis/authentication/
	key                   string
	disableRetries        bool
	configureLimiterOnce  sync.Once
	limiter               RateLimiter
	UserAgent             string
	defaultRequestOptions []RequestOptionFunc

	// Services for different parts of the FDA Api
	Fda510k         *Fda510kService
	Classification  *ClassificationService
	Enforcement     *EnforcementService
	Event           *EventService
	Pma             *PmaService
	Recall          *RecallService
	Reglist         *ReglistService //Registrations and Listings
	Covid19Serology *Covid19SerologyService
	Udi             *UdiService
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

func NewClient(apiKey string, options ...ClientOptionFunc) (*Client, error) {
	c := &Client{
		key:       apiKey,
		UserAgent: userAgent}

	// Create an http.Client with some sensible wait defaults for various responses
	c.client = &retryablehttp.Client{
		Backoff:      c.retryHTTPBackoff,
		CheckRetry:   c.retryHTTPCheck,
		ErrorHandler: retryablehttp.PassthroughErrorHandler,
		HTTPClient:   cleanhttp.DefaultPooledClient(),
		RetryWaitMin: 100 * time.Millisecond,
		RetryWaitMax: 400 * time.Millisecond,
		RetryMax:     5,
	}

	c.setBaseURL(defaultBaseUrl)

	// Apply any given client options.
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(c); err != nil {
			return nil, err
		}
	}

	if c.limiter == nil {
		c.limiter = rate.NewLimiter(rate.Inf, 0)
	}

	// Public Services
	c.Fda510k = &Fda510kService{client: c}
	c.Classification = &ClassificationService{client: c}
	c.Enforcement = &EnforcementService{client: c}
	c.Event = &EventService{client: c}
	c.Pma = &PmaService{client: c}
	c.Recall = &RecallService{client: c}
	c.Reglist = &ReglistService{client: c}
	c.Covid19Serology = &Covid19SerologyService{client: c}
	c.Udi = &UdiService{client: c}

	return c, nil
}

// retryHTTPCheck provides a callback for Client.CheckRetry which
// will retry both rate limit (429) and server (>= 500) errors.
func (c *Client) retryHTTPCheck(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	if err != nil {
		return false, err
	}
	if !c.disableRetries && (resp.StatusCode == 429 || resp.StatusCode >= 500) {
		return true, nil
	}
	return false, nil
}

// retryHTTPBackoff provides a generic callback for Client.Backoff which
// will pass through all calls based on the status code of the response.
func (c *Client) retryHTTPBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	// Use the rate limit backoff function when we are rate limited.
	if resp != nil && resp.StatusCode == 429 {
		return rateLimitBackoff(min, max, attemptNum, resp)
	}

	// Set custom duration's when we experience a service interruption.
	min = 700 * time.Millisecond
	max = 900 * time.Millisecond

	return retryablehttp.LinearJitterBackoff(min, max, attemptNum, resp)
}

// rateLimitBackoff provides a callback for Client.Backoff which will use the
// RateLimit-Reset header to determine the time to wait. We add some jitter
// to prevent a thundering herd.
//
// min and max are mainly used for bounding the jitter that will be added to
// the reset time retrieved from the headers. But if the final wait time is
// less then min, min will be used instead.
func rateLimitBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	// rnd is used to generate pseudo-random numbers.
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// First create some jitter bounded by the min and max durations.
	jitter := time.Duration(rnd.Float64() * float64(max-min))

	if resp != nil {
		if v := resp.Header.Get(headerRateReset); v != "" {
			if reset, _ := strconv.ParseInt(v, 10, 64); reset > 0 {
				// Only update min if the given time to wait is longer.
				if wait := time.Until(time.Unix(reset, 0)); wait > min {
					min = wait
				}
			}
		} else {
			// In case the RateLimit-Reset header is not set, back off an additional
			// 100% exponentially. With the default milliseconds being set to 100 for
			// `min`, this makes the 5th retry wait 3.2 seconds (3,200 ms) by default.
			min = time.Duration(float64(min) * math.Pow(2, float64(attemptNum)))
		}
	}

	return min + jitter
}

// configureLimiter configures the rate limiter.
func (c *Client) configureLimiter(ctx context.Context, headers http.Header) {
	if v := headers.Get(headerRateLimit); v != "" {
		if rateLimit, _ := strconv.ParseFloat(v, 64); rateLimit > 0 {
			// The rate limit is based on requests per minute, so for our limiter to
			// work correctly we divide the limit by 60 to get the limit per second.
			rateLimit /= 60

			// Configure the limit and burst using a split of 2/3 for the limit and
			// 1/3 for the burst. This enables clients to burst 1/3 of the allowed
			// calls before the limiter kicks in. The remaining calls will then be
			// spread out evenly using intervals of time.Second / limit which should
			// prevent hitting the rate limit.
			limit := rate.Limit(rateLimit * 0.66)
			burst := int(rateLimit * 0.33)

			// Need at least one allowed to burst or x/time will throw an error
			if burst == 0 {
				burst = 1
			}

			// Create a new limiter using the calculated values.
			c.limiter = rate.NewLimiter(limit, burst)

			// Call the limiter once as we have already made a request
			// to get the headers and the limiter is not aware of this.
			c.limiter.Wait(ctx)
		}
	}
}

// BaseURL getter.
func (c *Client) BaseURL() *url.URL {
	u := *c.baseUrl
	return &u
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

func (c *Client) BaseUrl() *url.URL {
	u := *c.baseUrl
	return &u
}

func (c *Client) NewRequest(method, path string, opt interface{}, options []RequestOptionFunc) (*retryablehttp.Request, error) {
	u := *c.baseUrl
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}
	u.RawPath = c.baseUrl.Path + path
	u.Path = c.baseUrl.Path + unescaped
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	if c.UserAgent != "" {
		reqHeaders.Set("User-Agent", c.UserAgent)
	}

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
	}

	req, err := retryablehttp.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	for _, fn := range append(c.defaultRequestOptions, options...) {
		if fn == nil {
			continue
		}
		if err := fn(req); err != nil {
			return nil, err
		}
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

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
func (c *Client) Do(req *retryablehttp.Request, w interface{}) (*Response, error) {

	err := c.limiter.Wait(req.Context())
	if err != nil {
		return nil, err
	}

	var apiKey string

	if c.key != "" {
		apiKey = c.key
		req.Header.Set("Authorization", "Basic "+apiKey)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)

	// If not yet configured, try to configure the rate limiter
	// using the response headers we just received. Fail silently
	// so the limiter will remain disabled in case of an error.
	c.configureLimiterOnce.Do(func() { c.configureLimiter(req.Context(), resp.Header) })

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

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
			errorResponse.Message = fmt.Sprintf("unknown error format: %s", data)
		} else {

			errorResponse.Message = fmt.Sprintf("raw error: %T", raw)
		}
	}

	return errorResponse
}
