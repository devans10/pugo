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
