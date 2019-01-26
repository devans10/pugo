// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

import (
	"testing"
)

func TestPure1Filesystems(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("GetFilesystems", testPure1GetFilesystems(c))
}

func testPure1GetFilesystems(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Filesystems.GetFilesystems(nil)
		if err != nil {
			t.Fatalf("error getting filesystems: %s", err)
		}
	}
}
