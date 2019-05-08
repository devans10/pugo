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

package pure1

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

	if os.Getenv("PURE1_ACC") == "" {
		t.Skip("set PURE1_ACC to run purestorage acceptance tests (provider connection is required)")
	}

	appID := os.Getenv("PURE1_APPID")
	privateKey := os.Getenv("PURE1_PRIVATEKEY")

	if appID == "" {
		t.Fatalf("PURE1_APPID must be set for acceptance tests")
	}
	if privateKey == "" {
		t.Fatalf("PURE1_PRIVATEKEY must be set for acceptance tests")
	}
}

func testAccGenerateClient(t *testing.T) *Client {
	appid := os.Getenv("PURE1_APPID")
	privateKey := []byte(os.Getenv("PURE1_PRIVATEKEY"))
	restVersion := ""

	c, err := NewClient(appid, privateKey, restVersion)
	if err != nil {
		t.Fatalf("error setting up client: %s", err)
	}
	return c
}
func TestPure1NewClient(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)
	if c == nil {
		t.Fatal("error setting up client")
	}
}

func testGenerateClient(fn RoundTripFunc) *Client {
	restVersion := "1.latest"
	c := &Client{AppID: "pure1:apikey:pafoQkLjHwHORmyF",
		PrivateKey:  []byte("MYPRIVATEKEY"),
		RestVersion: restVersion}

	c.token = &pure1Token{AccessToken: "eyJraWQiOiIwdVA1NWtOVG9XTGNOTWVrQ21PSFVhZW1LZUUyIiwidHlwIjoiSldUIiwiYWxnIjoiUlMyNTYifQ.eyJhdWQiOiJwdWJsaWMtYXBpIiwic3ViIjoiOTEiLCJyb2xlIjoiQ1VTVCIsIm9yZyI6NzMzLCJwaF9vaWQiOjcxODgsInBoX3VpZCI6MCwiZW1haWwiOiJkZXZhbnNAZHVxbGlnaHQuY29tIiwiaXNzIjoiaHR0cHM6Ly90ZXguY2xvdWQtc3VwcG9ydC5wdXJlc3RvcmFnZS5jb20iLCJpYXQiOjE1NTcwNzc5MjEsImV4cCI6MTU1NzExMzkyMSwianRpIjoiNTgxMWFmZTAtY2Q2MC00N2QwLTkwZDUtYmEyYzhmMjAyODNhIn0.Z0Q1x61Aj3QeXAqUi3B7SHpluwrLqcW75s_iRhg6b6VpOq-SEQCdG9oic-QBMPODf2M1wqaWe0zG9opsKpvcokMX7toioD64fQUgubwkE5qJJTdzl4MBwSHThYhvCafQv2LmAcEW_ZWNmdo4NJ4GgICHKxEwz9LWD3CrTuq8llsiFFQgmeBYQUHBGA5VolU5xIOFSeveDcI4yY-8yv3O_j-0VQQ2vLPCvEf87-bT3I3EKdpX5O-hw043OYkeDGQc7k1_dzLNIK40FGNNmUUO_4-dX5K7fDHPMJKMmAdSSfNol8srM2awMZasVSdsK-LrQOi5UMmy2KYgjkZqemAO0Q",
		IssuedTokenType: "urn:ietf:params:oauth:token-type:access_token",
		TokenType:       "Bearer",
		ExpiresIn:       36000000,
	}
	c.client = &http.Client{Transport: RoundTripFunc(fn)}

	c.Arrays = &ArrayService{client: c}
	c.Filesystems = &FilesystemService{client: c}
	c.FilesystemSnapshots = &FilesystemSnapshotService{client: c}
	c.Metrics = &MetricsService{client: c}
	c.NetworkInterfaces = &NetworkInterfacesService{client: c}
	c.Pods = &PodService{client: c}
	c.Volumes = &VolumeService{client: c}
	c.VolumeSnapshots = &VolumeSnapshotService{client: c}

	return c
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
