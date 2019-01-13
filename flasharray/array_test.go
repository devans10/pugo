// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"testing"
)

func TestAccArrayConsoleLock(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("EnableConsoleLock", testAccEnableConsoleLock(c))
	t.Run("GetConsoleLock", testAccGetConsoleLock(c))
	t.Run("DisableConsoleLock", testAccGetConsoleLock(c))
}

func testAccEnableConsoleLock(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		err := c.Array.enable_console_lock()
		if err != nil {
			t.Fatalf("error enabling console lock: %s", err)
		}
	}
}

func testAccGetConsoleLock(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		lock, err := c.Array.get_console_lock()
		if err != nil {
			t.Fatalf("error getting console lock: %s", err)
		}
		if lock.Console_lock != "enabled" {
			t.Fatalf("console lock disabled")
		}
	}
}

func testAccDisableConsoleLock(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		err := c.Array.disable_console_lock()
		if err != nil {
			t.Fatalf("error disabling console lock: %s", err)
		}
	}
}

func TestAccGetArray(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Array.Get(nil)
	if err != nil {
		t.Fatalf("error getting array: %s", err)
	}
}

func TestAccArrayRemoteAssist(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("EnableRemoteAssist", testAccEnableRemoteAssist(c))
	t.Run("GetRemoteAssist", testAccGetRemoteAssist(c))
	t.Run("DisableRemoteAssist", testAccGetRemoteAssist(c))
}

func testAccEnableRemoteAssist(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		if _, err := c.Array.EnableRemoteAssist(); err != nil {
			t.Fatalf("error enabling remote assist: %s", err)
		}
	}
}

func testAccGetRemoteAssist(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Array.GetRemoteAssist()
		if err != nil {
			t.Fatalf("error getting remote assist: %s", err)
		}
	}
}

func testAccDisableRemoteAssist(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		if _, err := c.Array.DisableRemoteAssist(); err != nil {
			t.Fatalf("error disabling console lock: %s", err)
		}
	}
}
