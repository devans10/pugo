// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

import (
	"testing"
)

func TestPure1Array(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("GetArrays", testPure1GetArrays(c))
	t.Run("GetTags", testPure1GetTags(c))
}

func testPure1GetArrays(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Arrays.GetArrays(nil)
		if err != nil {
			t.Fatalf("error getting arrays: %s", err)
		}
	}
}

func testPure1GetTags(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Arrays.GetTags(nil)
		if err != nil {
			t.Fatalf("error getting tags: %s", err)
		}
	}
}
