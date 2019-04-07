// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetDirectoryService(t *testing.T) {

	restVersion := "1.15"
	testDirsrv := Dirsrv{BaseDn: "DC=ad1,DC=example,DC=com",
		BindPassword: "****",
		BindUser:     "readonlyuser",
		CheckPeer:    false,
		Enabled:      true,
		URI:          []string{"ldaps://ad1.example.com"},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/directoryservice", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetDirectoryservice(restVersion))),
			Header:     head,
		}
	})

	ds, err := c.Dirsrv.GetDirectoryService()
	ok(t, err)
	equals(t, &testDirsrv, ds)
}

func TestGetDirectoryServiceError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/directoryservice", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetDirectoryservice(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Dirsrv.GetDirectoryService()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}
func TestAccGetDirectoryService(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Dirsrv.GetDirectoryService()
	if err != nil {
		t.Fatalf("error getting directory service: %s", err)
	}
}
