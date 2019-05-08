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
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

// Unit Tests
func TestPure1GetVolumeSnapshots(t *testing.T) {
	restVersion := "1.latest"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://api.pure1.purestorage.com/api/1.latest/volume-snapshots", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetVolumeSnapshots(restVersion))),
			Header:     head,
		}
	})

	_, err := c.VolumeSnapshots.GetVolumeSnapshots(nil)
	ok(t, err)
}

// Acceptance Tests
func TestPure1VolumeSnapshots(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("GetVolumeSnapshots", testPure1GetVolumeSnapshots(c))
}

func testPure1GetVolumeSnapshots(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.VolumeSnapshots.GetVolumeSnapshots(nil)
		if err != nil {
			t.Fatalf("error getting VolumeSnapshots: %s", err)
		}
	}
}
