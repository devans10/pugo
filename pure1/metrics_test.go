// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

import (
	"testing"
)

func TestPure1Metrics(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("GetMetrics", testPure1GetMetrics(c))
}

func testPure1GetMetrics(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Metrics.GetMetrics(nil)
		if err != nil {
			t.Fatalf("error getting metrics: %s", err)
		}
	}
}
