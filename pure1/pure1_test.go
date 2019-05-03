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
	"os"
	"testing"
)

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
