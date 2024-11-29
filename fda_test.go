package fda

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup(t *testing.T) (*http.ServeMux, *Client) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client, err := NewClient("")
	if err != nil {
		t.Fatalf("Error creating client: %v", err)
	}

	client.setBaseURL(server.URL) // Change to a parameter with a mutator function
	fmt.Printf("URL: %s", server.URL)
	return mux, client
}

func testURL(t *testing.T, r *http.Request, want string) {
	if got := r.RequestURI; got != want {
		t.Errorf("Request url: %+v, want %s", got, want)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %s, want %s", got, want)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	if err != nil {
		t.Fatalf("Failed to Read Body: %v", err)
	}

	if got := buffer.String(); got != want {
		t.Errorf("Request body: %s, want %s", got, want)
	}
}

func testParams(t *testing.T, r *http.Request, want string) {
	if got := r.URL.RawQuery; got != want {
		t.Errorf("Request query: %s, want %s", got, want)
	}
}

//TODO:
//TestNewClient
//TestCheckResponse
//Add above test helper tests to all tests
