package fixer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client ...
type Client struct {
	APIKey    string
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

// APIError ...
type APIError struct {
	Code int    `json:"code"`
	Type string `json:"type"`
	Info string `json:"info"`
}

// LatestResponse ...
type LatestResponse struct {
	Success   bool               `json:"success"`
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float32 `json:"rates"`
	Error     APIError           `json:"error"`
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path, RawQuery: fmt.Sprintf("access_key=%s", c.APIKey)}
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(v)
	return res, err
}

// Latest ...
func (c *Client) Latest() (*LatestResponse, error) {
	req, err := c.newRequest("GET", "/latest", nil)
	if err != nil {
		return nil, err
	}

	var latestResponse LatestResponse
	_, err = c.do(req, &latestResponse)
	return &latestResponse, err
}

// NewClient ...
func NewClient(apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse("http://data.fixer.io/api")

	c := &Client{
		httpClient: httpClient,
		BaseURL:    baseURL,
		APIKey:     apiKey,
		UserAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
	}

	return c
}
