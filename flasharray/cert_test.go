// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"testing"
)

func TestAccCert(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)
	cert := "test"
	data := make(map[string]interface{})
	data["common_name"] = "pure.example.com"
	data["self_signed"] = true
	data["state"] = "PA"

	t.Run("ListCert", testAccListCert(c))
	t.Run("CreateCert", testAccCreateCert(cert, data, c))
	t.Run("GetCert", testAccGetCert(cert, c))
	t.Run("GetCSR", testAccGetCSR(cert, c))
	t.Run("DeleteCert", testAccDeleteCert(cert, c))

}

func testAccListCert(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		if _, err := c.Cert.ListCert(); err != nil {
			t.Fatalf("error listing cert: %s", err)
		}
	}
}

func testAccCreateCert(name string, data interface{}, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Cert.CreateCert(name, data)
		if err != nil {
			t.Fatalf("error creating cert: %s", err)
		}
	}
}

func testAccGetCert(name string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Cert.GetCert(name, nil)
		if err != nil {
			t.Fatalf("error getting cert: %s", err)
		}
	}
}

func testAccGetCSR(name string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Cert.GetCSR(name, nil)
		if err != nil {
			t.Fatalf("error getting csr: %s", err)
		}
	}
}

func testAccDeleteCert(name string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Cert.DeleteCert(name)
		if err != nil {
			t.Fatalf("error deleting cert: %s", err)
		}
	}
}
