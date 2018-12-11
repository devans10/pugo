// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"testing"
)

func TestAccGetDirectoryService(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Dirsrv.GetDirectoryService()
	if err != nil {
		t.Fatalf("error getting directory service: %s", err)
	}
}
