package utility

import (
	"encoding/base64"
	"io"
	"net/http"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/gojek/heimdall/v7/plugins"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

//Client serves as a wrapper for out client
type Client struct {
	c *httpclient.Client
}

// GetClient provides a generic instance of http client
func GetClient() *Client {
	timeout := 5 * time.Second
	// Create a new client, sets the retry mechanism, and the number of times you would like to retry
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
	)
	requestLogger := plugins.NewRequestLogger(nil, nil)
	client.AddPlugin(requestLogger)

	return &Client{
		c: client,
	}
}

// Get makes a HTTP GET request to provided URL
func (c *Client) Get(path string, headers http.Header) (*http.Response, error) {
	url := getURL(path)
	var response *http.Response
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return response, errors.Wrap(err, "GET - request creation failed")
	}

	c.addDefaultHeaders(request)

	return c.c.Do(request)
}

// Post makes a HTTP POST request to provided URL and requestBody
func (c *Client) Post(path string, body io.Reader) (*http.Response, error) {
	url := getURL(path)
	var response *http.Response
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return response, errors.Wrap(err, "POST - request creation failed")
	}

	c.addDefaultHeaders(request)

	return c.c.Do(request)
}

// Put makes a HTTP PUT request to provided URL and requestBody
func (c *Client) Put(path string, body io.Reader, headers http.Header) (*http.Response, error) {
	url := getURL(path)
	var response *http.Response
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return response, errors.Wrap(err, "PUT - request creation failed")
	}

	request.Header = headers

	return c.c.Do(request)
}

// Patch makes a HTTP PATCH request to provided URL and requestBody
func (c *Client) Patch(path string, body io.Reader, headers http.Header) (*http.Response, error) {
	url := getURL(path)
	var response *http.Response
	request, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		return response, errors.Wrap(err, "PATCH - request creation failed")
	}

	request.Header = headers

	return c.c.Do(request)
}

// Delete makes a HTTP DELETE request with provided URL
func (c *Client) Delete(path string, headers http.Header) (*http.Response, error) {
	url := getURL(path)
	var response *http.Response
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return response, errors.Wrap(err, "DELETE - request creation failed")
	}

	request.Header = headers

	return c.c.Do(request)
}

func getURL(path string) string {
	baseURL := GetURL()
	return baseURL.String() + path
}

func (c *Client) addDefaultHeaders(req *http.Request) {
	mail := viper.GetString("email")
	key := viper.GetString("apikey")

	if len(key) > 0 && len(mail) > 0 {
		bkey := []byte(mail + ":" + key)
		key := base64.StdEncoding.EncodeToString(bkey)
		req.Header.Set("Authorization", "Basic "+key)
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("User-Agent", userAgent())

	// req.Header.Set("X-Trace-Id", c.c)
}

func userAgent() string {
	userAgent := "modulehubClient"

	// if version.Version != "" {
	// 	userAgent += fmt.Sprintf("-%s", version.Version)
	// }

	return userAgent
}
