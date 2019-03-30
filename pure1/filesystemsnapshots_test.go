// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

import (
	"testing"
)

func TestPure1FilesystemSnapshots(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("GetFilesystemSnapshots", testPure1GetFilesystemSnapshots(c))
}

func testPure1GetFilesystemSnapshots(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.FilesystemSnapshots.GetFilesystemSnapshots(nil)
		if err != nil {
			t.Fatalf("error getting filesystem snapshots: %s", err)
		}
	}
}
