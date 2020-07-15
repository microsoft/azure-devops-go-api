package azuredevops

import (
	"log"
	"net/http"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

// ClientOptionFunc can be used customize a new AzureDevops API client.
type ClientOptionFunc func(*Client)

// WithCustomRetry can be used to configure a custom retry policy.
func WithCustomRetry(checkRetry retryablehttp.CheckRetry) ClientOptionFunc {
	return func(c *Client) {
		c.client.CheckRetry = checkRetry
	}
}

// WithHTTPClient can be used to configure a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) {
		c.client.HTTPClient = httpClient
	}
}

// WithCustomLogger allows you to specify a custom logger.
func WithCustomLogger(logger *log.Logger) ClientOptionFunc {
	return func(c *Client) {
		c.logger = logger
	}
}
