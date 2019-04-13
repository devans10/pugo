// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

// Unit Tests

func TestGetSnmp(t *testing.T) {

	restVersion := "1.15"
	testSnmp := SnmpManager{
		Name:    "localhost",
		Host:    "localhost",
		Version: "v2c",
	}

	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/snmp/localhost", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetSnmpsnmp(restVersion))),
			Header:     head,
		}
	})

	s, err := c.Snmp.GetSnmp("localhost")
	ok(t, err)
	equals(t, &testSnmp, s)
}

func TestListSnmp(t *testing.T) {

	restVersion := "1.15"
	testSnmp := []SnmpManager{
		SnmpManager{
			Name:    "localhost",
			Host:    "localhost",
			Version: "v2c",
		},
	}

	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/snmp", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetSnmp(restVersion))),
			Header:     head,
		}
	})

	s, err := c.Snmp.ListSnmp(nil)
	ok(t, err)
	equals(t, testSnmp, s)
}

func TestCreateSnmp(t *testing.T) {

	restVersion := "1.15"
	testSnmp := SnmpManager{
		Name:    "localhost",
		Host:    "localhost",
		Version: "v2c",
	}

	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/snmp/localhost", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetSnmpsnmp(restVersion))),
			Header:     head,
		}
	})

	s, err := c.Snmp.CreateSnmp("localhost", testSnmp)
	ok(t, err)
	equals(t, &testSnmp, s)
}

func TestSetSnmp(t *testing.T) {

	restVersion := "1.15"
	testSnmp := SnmpManager{
		Name:    "localhost",
		Host:    "localhost",
		Version: "v2c",
	}

	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/snmp/localhost", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetSnmpsnmp(restVersion))),
			Header:     head,
		}
	})

	s, err := c.Snmp.SetSnmp("localhost", testSnmp)
	ok(t, err)
	equals(t, &testSnmp, s)
}

func TestDeleteSnmp(t *testing.T) {

	restVersion := "1.15"
	testSnmp := SnmpManager{
		Name: "localhost",
	}

	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/snmp/localhost", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteSnmpsnmp(restVersion))),
			Header:     head,
		}
	})

	s, err := c.Snmp.DeleteSnmp("localhost")
	ok(t, err)
	equals(t, &testSnmp, s)
}

// Acceptance Tests

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
