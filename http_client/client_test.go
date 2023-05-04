package httpclient

import (
	"os"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	clientConfig := &ClientConfig{
		Host: "localhost",
		Port: "8080",
	}

	client := Client(clientConfig)

	// Test that the client's base URL is set correctly
	expectedURL := "localhost:8080"
	if client.HostURL != expectedURL {
		t.Errorf("Client base URL is %v; expected %v", client.HostURL, expectedURL)
	}

	// Test that the client's headers are set correctly
	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"User-Agent":   "flgd-resty-client",
	}
	for k, v := range expectedHeaders {
		if client.Header.Get(k) != v {
			t.Errorf("Client header %v is %v; expected %v", k, client.Header.Get(k), v)
		}
	}

	// Test that the client's retry settings are set correctly
	expectedRetryCount := 2
	if client.RetryCount != expectedRetryCount {
		t.Errorf("Client retry count is %v; expected %v", client.RetryCount, expectedRetryCount)
	}
	expectedRetryWaitTime := 2 * time.Second
	if client.RetryWaitTime != expectedRetryWaitTime {
		t.Errorf("Client retry wait time is %v; expected %v", client.RetryWaitTime, expectedRetryWaitTime)
	}
}

func TestDefault(t *testing.T) {
	err := os.Setenv("FG_INTERNAL_URL", "http://localhost:8080")
	if err != nil {
		return
	}
	err = os.Setenv("PROJECT_NAME", "my-project")
	if err != nil {
		return
	}

	client := Default()

	// Test that the client's base URL is set correctly
	expectedURL := "http://localhost:8080"
	if client.HostURL != expectedURL {
		t.Errorf("Client base URL is %v; expected %v", client.HostURL, expectedURL)
	}

	// Test that the client's headers are set correctly
	expectedHeaders := map[string]string{
		"Content-Type":     "application/json",
		"Accept":           "application/json",
		"User-Agent":       "flgd-resty-client",
		"X-Application-ID": "my-project",
	}
	for k, v := range expectedHeaders {
		if client.Header.Get(k) != v {
			t.Errorf("Client header %v is %v; expected %v", k, client.Header.Get(k), v)
		}
	}

	// Test that the client's retry settings are set correctly
	expectedRetryCount := 3
	if client.RetryCount != expectedRetryCount {
		t.Errorf("Client retry count is %v; expected %v", client.RetryCount, expectedRetryCount)
	}
	expectedRetryWaitTime := 1 * time.Second
	if client.RetryWaitTime != expectedRetryWaitTime {
		t.Errorf("Client retry wait time is %v; expected %v", client.RetryWaitTime, expectedRetryWaitTime)
	}
}
