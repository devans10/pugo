// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"testing"
)

func TestAccGetConsoleLock(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Array.get_console_lock()
	if err != nil {
		t.Fatalf("error getting console lock: %s", err)
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

func TestAccGetRemoteAssist(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Array.GetRemoteAssist()
	if err != nil {
		t.Fatalf("error getting remote assist: %s", err)
	}
}
