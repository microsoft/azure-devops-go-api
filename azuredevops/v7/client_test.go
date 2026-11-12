package azuredevops

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestClient_NewClient(t *testing.T) {
	tlsConfig := &tls.Config{}
	timeout := 20 * time.Millisecond

	conn := &Connection{
		TlsConfig: tlsConfig,
		Timeout:   &timeout,
	}
	baseURL := "localhost"
	client := NewClient(conn, baseURL)
	if client.baseUrl != baseURL {
		t.Errorf("Expected baseURL: %v  Actual baseURL: %v", baseURL, client.baseUrl)
	}
	if actualTLSConfig := client.client.Transport.(*http.Transport).TLSClientConfig; actualTLSConfig != tlsConfig {
		t.Errorf("Expected tlsConfig: %v  Actual tlsConfig: %v", tlsConfig, actualTLSConfig)
	}
}

func TestClient_NewClientWithOptions_WithHTTPClient(t *testing.T) {
	tlsConfig := &tls.Config{}
	httpTimeout := 20 * time.Millisecond
	connTimeout := 40 * time.Millisecond

	conn := &Connection{
		TlsConfig: tlsConfig,
		Timeout:   &connTimeout, // will be ignored in favour of httpTimeout
	}

	httpClient := &http.Client{Timeout: httpTimeout}
	baseURL := "localhost"

	client := newClientWithOptions(conn, baseURL, WithHTTPClient(httpClient))
	if client.baseUrl != baseURL {
		t.Errorf("Expected baseURL: %v  Actual baseURL: %v", baseURL, client.baseUrl)
	}
	if actualHTTPClient := client.client; actualHTTPClient.Timeout != httpClient.Timeout {
		t.Errorf("Expected httpClient.Timeout: %#v  Actual httpClient.Timeout: %#v", httpClient.Timeout, actualHTTPClient.Timeout)
	}
}

func TestClient_NewClientWithOptions_WithRetryOptions(t *testing.T) {
	conn := &Connection{}
	opts := RetryOptions{MaxRetries: 5, Delay: 2 * time.Second}
	client := newClientWithOptions(conn, "localhost", WithRetryOptions(opts))
	if client.retryOptions == nil {
		t.Fatal("Expected retryOptions to be set")
	}
	if client.retryOptions.MaxRetries != 5 {
		t.Errorf("Expected MaxRetries=5, got %d", client.retryOptions.MaxRetries)
	}
	if client.retryOptions.Delay != 2*time.Second {
		t.Errorf("Expected Delay=2s, got %v", client.retryOptions.Delay)
	}
}

// roundTripFunc adapts a function to http.RoundTripper for testing.
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestSendRequest_NoRetry_Success(t *testing.T) {
	client := &Client{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil
			}),
		},
	}
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	resp, err := client.SendRequest(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestSendRequest_NoRetryOptions_TransientError(t *testing.T) {
	// Without retry options, transient errors are returned immediately.
	client := &Client{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				return nil, fmt.Errorf("connection reset by peer")
			}),
		},
	}
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	_, err := client.SendRequest(req)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "connection reset by peer") {
		t.Errorf("expected 'connection reset by peer' error, got: %v", err)
	}
}

func TestSendRequest_RetryOnTransientError(t *testing.T) {
	var attempts int32
	client := &Client{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				n := atomic.AddInt32(&attempts, 1)
				if n <= 2 {
					return nil, fmt.Errorf("read tcp: connection reset by peer")
				}
				return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil
			}),
		},
		retryOptions: &RetryOptions{
			MaxRetries: 3,
			Delay:      time.Millisecond, // fast for tests
		},
	}
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	resp, err := client.SendRequest(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if atomic.LoadInt32(&attempts) != 3 {
		t.Errorf("expected 3 attempts, got %d", atomic.LoadInt32(&attempts))
	}
}

func TestSendRequest_RetryExhausted(t *testing.T) {
	var attempts int32
	client := &Client{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				atomic.AddInt32(&attempts, 1)
				return nil, fmt.Errorf("peer connection closed")
			}),
		},
		retryOptions: &RetryOptions{
			MaxRetries: 2,
			Delay:      time.Millisecond,
		},
	}
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	_, err := client.SendRequest(req)
	if err == nil {
		t.Fatal("expected error after retries exhausted")
	}
	// 1 initial + 2 retries = 3 total attempts
	if atomic.LoadInt32(&attempts) != 3 {
		t.Errorf("expected 3 attempts, got %d", atomic.LoadInt32(&attempts))
	}
}

func TestSendRequest_RetryWithBody(t *testing.T) {
	var bodies []string
	var attempts int32
	client := &Client{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				b, _ := io.ReadAll(req.Body)
				bodies = append(bodies, string(b))
				n := atomic.AddInt32(&attempts, 1)
				if n == 1 {
					return nil, fmt.Errorf("connection reset by peer")
				}
				return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil
			}),
		},
		retryOptions: &RetryOptions{
			MaxRetries: 2,
			Delay:      time.Millisecond,
		},
	}
	body := "request body content"
	req, _ := http.NewRequest("POST", "http://example.com", bytes.NewBufferString(body))
	resp, err := client.SendRequest(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if len(bodies) != 2 {
		t.Fatalf("expected 2 attempts, got %d", len(bodies))
	}
	for i, b := range bodies {
		if b != body {
			t.Errorf("attempt %d: expected body %q, got %q", i, body, b)
		}
	}
}

func TestSendRequest_NoRetryOnContextCanceled(t *testing.T) {
	var attempts int32
	client := &Client{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				atomic.AddInt32(&attempts, 1)
				return nil, context.Canceled
			}),
		},
		retryOptions: &RetryOptions{
			MaxRetries: 3,
			Delay:      time.Millisecond,
		},
	}
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	_, err := client.SendRequest(req)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled, got: %v", err)
	}
	if atomic.LoadInt32(&attempts) != 1 {
		t.Errorf("expected 1 attempt (no retry), got %d", atomic.LoadInt32(&attempts))
	}
}

func TestSendRequest_NoRetryOnNonRetryableError(t *testing.T) {
	var attempts int32
	client := &Client{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				atomic.AddInt32(&attempts, 1)
				return nil, fmt.Errorf("some permanent error")
			}),
		},
		retryOptions: &RetryOptions{
			MaxRetries: 3,
			Delay:      time.Millisecond,
		},
	}
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	_, err := client.SendRequest(req)
	if err == nil {
		t.Fatal("expected error")
	}
	if atomic.LoadInt32(&attempts) != 1 {
		t.Errorf("expected 1 attempt (no retry for non-retryable), got %d", atomic.LoadInt32(&attempts))
	}
}

func TestSendRequest_CustomIsRetryable(t *testing.T) {
	var attempts int32
	client := &Client{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				n := atomic.AddInt32(&attempts, 1)
				if n == 1 {
					return nil, fmt.Errorf("custom transient error")
				}
				return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(""))}, nil
			}),
		},
		retryOptions: &RetryOptions{
			MaxRetries: 2,
			Delay:      time.Millisecond,
			IsRetryable: func(resp *http.Response, err error) bool {
				return err != nil && strings.Contains(err.Error(), "custom transient")
			},
		},
	}
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	resp, err := client.SendRequest(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if atomic.LoadInt32(&attempts) != 2 {
		t.Errorf("expected 2 attempts, got %d", atomic.LoadInt32(&attempts))
	}
}

func TestSendRequest_ContextCancelDuringSleep(t *testing.T) {
	var attempts int32
	client := &Client{
		client: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				atomic.AddInt32(&attempts, 1)
				return nil, fmt.Errorf("connection reset by peer")
			}),
		},
		retryOptions: &RetryOptions{
			MaxRetries: 3,
			Delay:      5 * time.Second, // long delay
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)

	done := make(chan error, 1)
	go func() {
		_, err := client.SendRequest(req)
		done <- err
	}()

	// Cancel context shortly after the first failed attempt triggers a retry sleep.
	time.Sleep(50 * time.Millisecond)
	cancel()

	err := <-done
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled, got: %v", err)
	}
}
