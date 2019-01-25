// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

import (
	"testing"
)

func TestPure1Pods(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("GetPods", testPure1GetPods(c))
}

func testPure1GetPods(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Pods.GetPods(nil)
		if err != nil {
			t.Fatalf("error getting pods: %s", err)
		}
	}
}
