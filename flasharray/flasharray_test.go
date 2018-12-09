package flasharray

import (
	"testing"
)

// Test that a NewClient call with no authentication returns an error
func TestNewClientNoAuth(t *testing.T) {

	_, err := NewClient("target", "", "", "", "rest_version", false, false, "user_agent", nil)
	if err == nil {
		t.Errorf("An Error was NOT raised when no authentication methods provided")
	}
}

// Test that a NewClient call with all authentication methods provided returns an error
func TestNewClientAllAuth(t *testing.T) {

	_, err := NewClient("target", "username", "password", "api_token", "rest_version", false, false, "user_agent", nil)
	if err == nil {
		t.Errorf("An Error was NOT raised when All authentication methods were provided")
	}
}

// Test that NewRequest returns an http.Request object
func TestNewRequest(t *testing.T) {

	c := &Client{Target: "flasharray.example.com", Username: "", Password: "", Api_token: "apitoken", Rest_version: "1.0", User_agent: "", Request_kwargs: nil}

	req, err := c.NewRequest("GET", "array", nil, nil)

	if err != nil {
		t.Errorf("NewRequest function call returned error: %s", err)
	}

	if req.Method != "GET" {
		t.Errorf("Request Method: %s; Expected: GET", req.Method)
	}

	if req.URL.String() != "https://flasharray.example.com/api/1.0/array" {
		t.Errorf("Malformed URL returned by NewRequest. Expected: https://flasharray.example.com/api/1.0/array; Got: %s", req.URL.String())
	}
}
