/*
   Copyright 2018 David Evans

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package flasharray

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// RoundTripFunc is for returning a test response to the client
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip is the test http Transport
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

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
	apiToken := os.Getenv("PURE_APITOKEN")
	restVersion := ""
	target := os.Getenv("PURE_TARGET")
	verifyHTTPS := false
	sslCert := false
	userAgent := ""

	c, err := NewClient(target, username, password, apiToken, restVersion, verifyHTTPS, sslCert, userAgent, nil)
	if err != nil {
		t.Fatalf("error setting up client: %s", err)
	}
	return c
}

func TestAccClient(t *testing.T) {
	testAccPreChecks(t)
	testAccGenerateClient(t)
}

func testGenerateClient(fn RoundTripFunc) *Client {
	restVersion := "1.15"
	c := &Client{Target: "flasharray.example.com",
		RestVersion: restVersion,
		UserAgent:   ""}

	c.client = &http.Client{Transport: RoundTripFunc(fn)}

	c.Array = &ArrayService{client: c}
	c.Volumes = &VolumeService{client: c}
	c.Hosts = &HostService{client: c}
	c.Hostgroups = &HostgroupService{client: c}
	c.Offloads = &OffloadService{client: c}
	c.Protectiongroups = &ProtectiongroupService{client: c}
	c.Vgroups = &VgroupService{client: c}
	c.Networks = &NetworkService{client: c}
	c.Hardware = &HardwareService{client: c}
	c.Users = &UserService{client: c}
	c.Dirsrv = &DirsrvService{client: c}
	c.Pods = &PodService{client: c}
	c.Alerts = &AlertService{client: c}
	c.Messages = &MessageService{client: c}
	c.Snmp = &SnmpService{client: c}
	c.Cert = &CertService{client: c}
	c.SMTP = &SMTPService{client: c}

	return c
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
func TestNewRequestNoParamNoData(t *testing.T) {

	c := &Client{Target: "flasharray.example.com",
		Username:      "",
		Password:      "",
		APIToken:      "apitoken",
		RestVersion:   "1.0",
		UserAgent:     "",
		RequestKwargs: nil}

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

func TestNewRequestParamNoData(t *testing.T) {

	c := &Client{Target: "flasharray.example.com",
		Username:      "",
		Password:      "",
		APIToken:      "apitoken",
		RestVersion:   "1.0",
		UserAgent:     "",
		RequestKwargs: nil}

	params := map[string]string{"test": "true"}
	req, err := c.NewRequest("GET", "array", params, nil)

	if err != nil {
		t.Errorf("NewRequest function call returned error: %s", err)
	}

	if req.Method != "GET" {
		t.Errorf("Request Method: %s; Expected: GET", req.Method)
	}

	if req.URL.String() != "https://flasharray.example.com/api/1.0/array?test=true" {
		t.Errorf("Malformed URL returned by NewRequest. Expected: https://flasharray.example.com/api/1.0/array?test=true; Got: %s", req.URL.String())
	}
}

func TestNewRequestParamData(t *testing.T) {

	c := &Client{Target: "flasharray.example.com",
		Username:      "user",
		Password:      "secret",
		APIToken:      "",
		RestVersion:   "1.0",
		UserAgent:     "",
		RequestKwargs: nil}

	params := map[string]string{"test": "true"}
	data := map[string]string{"test": "true"}
	req, err := c.NewRequest("GET", "array", params, data)

	if err != nil {
		t.Errorf("NewRequest function call returned error: %s", err)
	}

	if req.Method != "GET" {
		t.Errorf("Request Method: %s; Expected: GET", req.Method)
	}

	if req.URL.String() != "https://flasharray.example.com/api/1.0/array?test=true" {
		t.Errorf("Malformed URL returned by NewRequest. Expected: https://flasharray.example.com/api/1.0/array?test=true; Got: %s", req.URL.String())
	}

	body, _ := req.GetBody()
	if body == nil {
		t.Errorf("Body not generated")
	}
}
func TestNewRequestNoParamData(t *testing.T) {

	c := &Client{Target: "flasharray.example.com",
		Username:      "",
		Password:      "",
		APIToken:      "apitoken",
		RestVersion:   "1.0",
		UserAgent:     "curl",
		RequestKwargs: map[string]string{"verify": "true"},
	}

	data := map[string]string{"test": "true"}
	req, err := c.NewRequest("GET", "array", nil, data)

	if err != nil {
		t.Errorf("NewRequest function call returned error: %s", err)
	}

	if req.Method != "GET" {
		t.Errorf("Request Method: %s; Expected: GET", req.Method)
	}

	if req.URL.String() != "https://flasharray.example.com/api/1.0/array" {
		t.Errorf("Malformed URL returned by NewRequest. Expected: https://flasharray.example.com/api/1.0/array; Got: %s", req.URL.String())
	}

	body, _ := req.GetBody()
	if body == nil {
		t.Errorf("Body not generated")
	}
}

func TestDecodeResponseNilIntf(t *testing.T) {
	r := http.Response{}
	err := decodeResponse(&r, nil)
	if err == nil {
		t.Errorf("error not raised for nil interface")
	}
}
func TestCheckAuthAPIToken(t *testing.T) {
	err := checkAuth("apitoken", "", "")
	ok(t, err)
}

func TestCheckAuthUserPasswd(t *testing.T) {
	err := checkAuth("", "username", "password")
	ok(t, err)
}

func TestCheckAuthErrorAll(t *testing.T) {
	err := checkAuth("apitoken", "username", "password")
	if err == nil {
		t.Errorf("error not raised when all auth options provided")
	}
}

func TestCheckAuthErrorNon(t *testing.T) {
	err := checkAuth("", "", "")
	if err == nil {
		t.Errorf("error not raised when no auth options provided")
	}
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
