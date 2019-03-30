// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEnableConsoleLock(t *testing.T) {

	testConsoleLock := ConsoleLock{"enabled"}
	body, _ := json.Marshal(testConsoleLock)
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/array/console_lock")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			Header:     head,
		}
	})

	err := c.Array.EnableConsoleLock()
	ok(t, err)
}

func TestDisableConsoleLock(t *testing.T) {

	testConsoleLock := ConsoleLock{"disabled"}
	body, _ := json.Marshal(testConsoleLock)
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/array/console_lock")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			Header:     head,
		}
	})

	err := c.Array.DisableConsoleLock()
	ok(t, err)
}

func TestGetConsoleLock(t *testing.T) {

	testConsoleLock := ConsoleLock{"disabled"}
	body, _ := json.Marshal(testConsoleLock)
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/array/console_lock")
		equals(t, req.Method, "GET")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			Header:     head,
		}
	})

	cl, err := c.Array.GetConsoleLock()
	ok(t, err)
	equals(t, &testConsoleLock, cl)
}

func TestGet(t *testing.T) {

	testArray := Array{ID: "b75f8356-604b-431d-af5c-64c3ca303749",
		ArrayName: "pure01",
		Version:   "5.0.0",
		Revision:  "201712160033+517009f"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/array")
		equals(t, req.Method, "GET")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetArray("1.15"))),
			Header:     head,
		}
	})

	array, err := c.Array.Get(nil)
	ok(t, err)
	equals(t, &testArray, array)
}

func TestAccArrayConsoleLock(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("GetConsoleLock", testAccGetConsoleLock(c))
}

func testAccEnableConsoleLock(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		err := c.Array.EnableConsoleLock()
		if err != nil {
			t.Fatalf("error enabling console lock: %s", err)
		}
	}
}

func testAccGetConsoleLock(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		lock, err := c.Array.GetConsoleLock()
		if err != nil {
			t.Fatalf("error getting console lock: %s", err)
		}
		if lock.ConsoleLock != "enabled" {
			t.Fatalf("console lock disabled")
		}
	}
}

func testAccDisableConsoleLock(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		err := c.Array.DisableConsoleLock()
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
