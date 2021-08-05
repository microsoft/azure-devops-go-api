package azuredevops

import (
	"crypto/tls"
	"net/http"
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
		t.Errorf("Expected tlsConfig: %v  Actual tlsConfig: %v", tlsConfig, *actualTLSConfig)
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

	client := NewClientWithOptions(conn, baseURL, WithHTTPClient(httpClient))
	if client.baseUrl != baseURL {
		t.Errorf("Expected baseURL: %v  Actual baseURL: %v", baseURL, client.baseUrl)
	}
	if actualHTTPClient := client.client; actualHTTPClient.Timeout != httpClient.Timeout {
		t.Errorf("Expected httpClient.Timeout: %#v  Actual httpClient.Timeout: %#v", httpClient.Timeout, actualHTTPClient.Timeout)
	}
}
