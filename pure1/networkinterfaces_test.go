// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

import (
	"testing"
)

func TestPure1NetworkInterfaces(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("GetNetworkInterfaces", testPure1GetNetworkInterfaces(c))
}

func testPure1GetNetworkInterfaces(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.NetworkInterfaces.GetNetworkInterfacess(nil)
		if err != nil {
			t.Fatalf("error getting network interfaces: %s", err)
		}
	}
}
