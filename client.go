package openfec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Client is the primary struct that this package provides. It represents the
// connection to the OpenFEC API.
type Client struct {
	APIKey string
	logger *log.Logger
}

// NewClient creates a new Client to interact with the OpenFEC API.
func NewClient(APIKey string) (*Client, error) {
	return &Client{
		APIKey: APIKey,
	}, nil
}

// TraceOn turns on API response tracing to the given logger.
func (c *Client) TraceOn(logger *log.Logger) {
	c.logger = logger
}

// TraceOff turns on API response tracing
func (c *Client) TraceOff() {
	c.logger = nil
}

func (c *Client) trace(args ...interface{}) {
	if c.logger != nil {
		c.logger.Println(args)
	}
}

type APIError struct {
	Err struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s - %s", e.Err.Code, e.Err.Message)
}

func (c *Client) do(endpoint string, parameters interface{}, pagination *pagination) (*http.Response, error) {
	u, err := url.Parse(BaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	q, err := ToValues(parameters)
	if err != nil {
		return nil, err
	}
	q.Add("api_key", c.APIKey)
	if pagination != nil {
		q.Add("page", fmt.Sprint(pagination.Page))
		q.Add("per_page", fmt.Sprint(pagination.PerPage))
	}
	u.RawQuery = q.Encode()
	c.trace(u.String())
	resp, err := http.DefaultClient.Get(u.String())
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
	case 201:
	case 400:
		fallthrough
	case 404:
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		c.trace("got 404, body:", string(body))
		return nil, ErrNotFound
	case 401, 403:
		return resp, ErrUnauthorized
	default:
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		c.trace("got 500, body:", string(body))
		if body != nil {
			var apiErr *APIError
			if err := json.Unmarshal(body, &apiErr); err == nil {
				return resp, apiErr
			}
		}
		return resp, fmt.Errorf("got unexpected status code %d: %s", resp.StatusCode)
	}
	return resp, err
}
