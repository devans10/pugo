// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

import (
	"testing"
)

func TestPure1Volumes(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("GetVolumes", testPure1GetVolumes(c))
}

func testPure1GetVolumes(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.GetVolumes(nil)
		if err != nil {
			t.Fatalf("error getting Volumes: %s", err)
		}
	}
}
