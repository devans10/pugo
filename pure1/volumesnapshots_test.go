// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

import (
	"testing"
)

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
