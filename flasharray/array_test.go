/*
   Copyright 2018 David Evans

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

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

func TestEnableConsoleLockError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/console_lock", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayConsoleLock(restVersion))),
			Header:     head,
		}
	})

	err := c.Array.EnableConsoleLock()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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
func TestDisableConsoleLockError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/console_lock", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayConsoleLock(restVersion))),
			Header:     head,
		}
	})

	err := c.Array.DisableConsoleLock()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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

func TestGetConsoleLockError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/console_lock", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetArrayConsoleLock(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.GetConsoleLock()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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

func TestGetError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetArray(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.Get(nil)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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

func TestGetArrayError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetArray(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.GetArray(nil, nil)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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

func TestGetArraySpaceError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array?space=true", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetArraySpace(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.GetArraySpace(nil)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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

func TestGetMonitorError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array?action=monitor", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetArrayMonitor(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.GetArrayMonitor(nil)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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

func TestDisablePhoneHomeError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/phonehome", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayPhonehome(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.DisablePhoneHome()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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

func TestEnablePhoneHomeError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/phonehome", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayPhonehome(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.EnablePhoneHome()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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

func TestGetManualPhoneHomeError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/phonehome", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayPhonehome(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.GetManualPhoneHome()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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

func TestEnableRemoteAssistError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/remoteassist", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayRemoteassist(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.EnableRemoteAssist()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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

func TestDisableRemoteAssistError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/remoteassist", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayRemoteassist(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.DisableRemoteAssist()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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

func TestGetRemoteAssistError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/array/remoteassist", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutArrayRemoteassist(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Array.GetRemoteAssist()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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
