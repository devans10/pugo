// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"testing"
)

func TestAccSnmp(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)
	snmpmanager := "test"
	data := make(map[string]interface{})
	data["host"] = "snmp.example.com"
	data["version"] = "v2c"
	data["community"] = "test"

	t.Run("ListSnmp", testAccListSnmp(c))
	t.Run("CreateSnmp", testAccCreateSnmp(snmpmanager, data, c))
	t.Run("GetSnmp", testAccGetSnmp(snmpmanager, c))
	t.Run("DeleteSnmp", testAccDeleteSnmp(snmpmanager, c))

}

func testAccListSnmp(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		if _, err := c.Snmp.ListSnmp(nil); err != nil {
			t.Fatalf("error listing snmp: %s", err)
		}
	}
}

func testAccCreateSnmp(name string, data interface{}, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Snmp.CreateSnmp(name, data)
		if err != nil {
			t.Fatalf("error creating snmp: %s", err)
		}
	}
}

func testAccGetSnmp(name string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Snmp.GetSnmp(name)
		if err != nil {
			t.Fatalf("error getting snmp: %s", err)
		}
	}
}

func testAccDeleteSnmp(name string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Snmp.DeleteSnmp(name)
		if err != nil {
			t.Fatalf("error deleting snmp: %s", err)
		}
	}
}
