package azuredevops

import (
	"net/http"
	"time"
)

// ClientOptionFunc can be used customize a new AzureDevops API client.
type ClientOptionFunc func(*Client)

// WithHTTPClient can be used to configure a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) {
		c.client = httpClient
	}
}

// WithRetryOptions configures retry behavior for transient errors.
// When set, the client will retry failed requests that match the IsRetryable
// predicate, up to MaxRetries times with exponential backoff.
func WithRetryOptions(options RetryOptions) ClientOptionFunc {
	return func(c *Client) {
		c.retryOptions = &options
	}
}

// RetryOptions configures retry behavior for the client.
type RetryOptions struct {
	// MaxRetries is the maximum number of retry attempts.
	// A value of 0 means no retries.
	// Defaults to 3.
	MaxRetries int

	// Delay is the initial delay between retries. Subsequent retries use
	// exponential backoff (delay * 2^attempt). Default: 1 second.
	Delay time.Duration

	// IsRetryable determines whether a failed request should be retried.
	// It receives the HTTP response (may be nil for transport-level errors)
	// and the error. If nil, DefaultIsRetryable is used.
	IsRetryable func(resp *http.Response, err error) bool
}
