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

func TestEnableConsoleLock(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/console_lock", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayConsoleLock(restVersion))),
			Header:     head,
		}
	})

	err := c.Array.EnableConsoleLock()
	ok(t, err)
}

func TestDisableConsoleLock(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/console_lock", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayConsoleLock(restVersion))),
			Header:     head,
		}
	})

	err := c.Array.DisableConsoleLock()
	ok(t, err)
}

func TestGetConsoleLock(t *testing.T) {

	restVersion := "1.15"
	testConsoleLock := ConsoleLock{"disabled"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/console_lock", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetArrayConsoleLock(restVersion))),
			Header:     head,
		}
	})

	cl, err := c.Array.GetConsoleLock()
	ok(t, err)
	equals(t, &testConsoleLock, cl)
}

func TestGet(t *testing.T) {

	restVersion := "1.15"
	testArray := Array{ID: "b75f8356-604b-431d-af5c-64c3ca303749",
		ArrayName: "pure01",
		Version:   "5.0.0",
		Revision:  "201712160033+517009f"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetArray(restVersion))),
			Header:     head,
		}
	})

	array, err := c.Array.Get(nil)
	ok(t, err)
	equals(t, &testArray, array)
}

func TestGetArray(t *testing.T) {

	restVersion := "1.15"
	testArray := Array{ID: "b75f8356-604b-431d-af5c-64c3ca303749",
		ArrayName: "pure01",
		Version:   "5.0.0",
		Revision:  "201712160033+517009f"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetArray(restVersion))),
			Header:     head,
		}
	})

	array, err := c.Array.GetArray(nil, nil)
	ok(t, err)
	equals(t, &testArray, array)
}

func TestGetArraySpace(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array?space=true", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetArraySpace(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.GetArraySpace(nil)
	ok(t, err)
}

func TestGetMonitor(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array?action=monitor", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetArrayMonitor(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.GetArrayMonitor(nil)
	ok(t, err)
}

func TestDisablePhoneHome(t *testing.T) {

	restVersion := "1.15"
	testPhonehome := Phonehome{"enabled", "", ""}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/phonehome", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayPhonehome(restVersion))),
			Header:     head,
		}
	})

	ph, err := c.Array.DisablePhoneHome()
	ok(t, err)
	equals(t, &testPhonehome, ph)
}

func TestEnablePhoneHome(t *testing.T) {

	restVersion := "1.15"
	testPhonehome := Phonehome{"enabled", "", ""}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/phonehome", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayPhonehome(restVersion))),
			Header:     head,
		}
	})

	ph, err := c.Array.EnablePhoneHome()
	ok(t, err)
	equals(t, &testPhonehome, ph)
}

func TestGetManualPhoneHome(t *testing.T) {

	restVersion := "1.15"
	testPhonehome := Phonehome{"enabled", "", ""}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/phonehome", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayPhonehome(restVersion))),
			Header:     head,
		}
	})

	ph, err := c.Array.GetManualPhoneHome()
	ok(t, err)
	equals(t, &testPhonehome, ph)
}

func TestEnableRemoteAssist(t *testing.T) {

	restVersion := "1.15"
	testRemoteAssist := RemoteAssist{"enabled", "pure01-ct0", "pure01-ct0.example.com-11679"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/remoteassist", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayRemoteassist(restVersion))),
			Header:     head,
		}
	})

	ra, err := c.Array.EnableRemoteAssist()
	ok(t, err)
	equals(t, &testRemoteAssist, ra)
}

func TestDisableRemoteAssist(t *testing.T) {

	restVersion := "1.15"
	testRemoteAssist := RemoteAssist{"enabled", "pure01-ct0", "pure01-ct0.example.com-11679"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/remoteassist", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayRemoteassist(restVersion))),
			Header:     head,
		}
	})

	ra, err := c.Array.DisableRemoteAssist()
	ok(t, err)
	equals(t, &testRemoteAssist, ra)
}

func TestGetRemoteAssist(t *testing.T) {

	restVersion := "1.15"
	testRemoteAssist := RemoteAssist{"enabled", "pure01-ct0", "pure01-ct0.example.com-11679"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/remoteassist", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayRemoteassist(restVersion))),
			Header:     head,
		}
	})

	ra, err := c.Array.GetRemoteAssist()
	ok(t, err)
	equals(t, &testRemoteAssist, ra)
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

func TestAccGet(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Array.Get(nil)
	if err != nil {
		t.Fatalf("error getting array: %s", err)
	}
}

func TestAccGetArray(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Array.GetArray(nil, nil)
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
