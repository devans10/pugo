package flasharray

import (
	"os"
	"testing"
)

func testAccPreChecks(t *testing.T) {

	if os.Getenv("PURE_ACC") == "" {
		t.Skip("set PURE_ACC to run purestorage acceptance tests (provider connection is required)")
	}

	target := os.Getenv("PURE_TARGET")
	username := os.Getenv("PURE_USERNAME")
	password := os.Getenv("PURE_PASSWORD")
	apitoken := os.Getenv("PURE_APITOKEN")
	if target == "" {
		t.Fatalf("PURE_TARGET must be set for acceptance tests")
	}
	if (apitoken == "") && (username == "") && (password == "") {
		t.Fatalf("PURE_USERNAME and PURE_PASSWORD or PURE_APITOKEN must be set for acceptance tests")
	}
	if (username != "") && (password == "") {
		t.Fatalf("PURE_PASSWORD must be set if PURE_USERNAME is set for acceptance tests")
	}
}

func testAccGenerateClient(t *testing.T) *Client {

	username := os.Getenv("PURE_USERNAME")
	password := os.Getenv("PURE_PASSWORD")
	api_token := os.Getenv("PURE_APITOKEN")
	rest_version := ""
	target := os.Getenv("PURE_TARGET")
	verify_https := false
	ssl_cert := false
	user_agent := ""

	c, err := NewClient(target, username, password, api_token, rest_version, verify_https, ssl_cert, user_agent, nil)
	if err != nil {
		t.Fatalf("error setting up client: %s", err)
	}
	return c
}

func TestAccClient(t *testing.T) {
	testAccPreChecks(t)
	testAccGenerateClient(t)
}

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
