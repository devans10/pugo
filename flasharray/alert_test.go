// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"testing"
)

func TestAccAlerts(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)
	address := "test@example.com"

	t.Run("ListAlerts", testAccListAlerts(c))
	t.Run("CreateAlert", testAccCreateAlert(address, c))
	t.Run("GetAlert", testAccGetAlert(address, c))
	t.Run("EnableAlert", testAccEnableAlert(address, c))
	t.Run("DisableAlert", testAccDisableAlert(address, c))
	t.Run("DeleteAlert", testAccDeleteAlert(address, c))

}

func testAccListAlerts(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		if _, err := c.Alerts.ListAlerts(nil); err != nil {
			t.Fatalf("error listing alerts: %s", err)
		}
	}
}

func testAccCreateAlert(address string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		alert, err := c.Alerts.CreateAlert(address, nil)
		if err != nil {
			t.Fatalf("error creating alert: %s", err)
		}
		if alert.Name != address {
			t.Fatalf("expecting: %s, got: %s", address, alert.Name)
		}
	}
}

func testAccGetAlert(address string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		alert, err := c.Alerts.GetAlert(address)
		if err != nil {
			t.Fatalf("error getting alert: %s", err)
		}
		if alert.Name != address {
			t.Fatalf("expecting: %s, got: %s", address, alert.Name)
		}
	}
}

func testAccEnableAlert(address string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		alert, err := c.Alerts.EnableAlert(address)
		if err != nil {
			t.Fatalf("error enabling alert: %s", err)
		}
		if alert.Enabled != true {
			t.Fatalf("alert not enabled")
		}
	}
}

func testAccDisableAlert(address string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		alert, err := c.Alerts.DisableAlert(address)
		if err != nil {
			t.Fatalf("error disabling alert: %s", err)
		}
		if alert.Enabled != false {
			t.Fatalf("alert enabled")
		}
	}
}

func testAccDeleteAlert(address string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Alerts.DeleteAlert(address)
		if err != nil {
			t.Fatalf("error deleting alert: %s", err)
		}
	}
}
