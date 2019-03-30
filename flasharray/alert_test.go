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

func TestListAlerts(t *testing.T) {

	testAlert := []Alert{Alert{"Test", true}}
	body, _ := json.Marshal(testAlert)
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert")
		equals(t, req.Method, "GET")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			Header:     head,
		}
	})

	alert, err := c.Alerts.ListAlerts(nil)
	ok(t, err)
	equals(t, testAlert, alert)
}

func TestGetAlert(t *testing.T) {

	testAlert := Alert{"Test", true}
	body, _ := json.Marshal(testAlert)
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/Test")
		equals(t, req.Method, "GET")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			Header:     head,
		}
	})

	alert, err := c.Alerts.GetAlert("Test")
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestCreateAlert(t *testing.T) {

	testAlert := Alert{"Test", true}
	body, _ := json.Marshal(testAlert)
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/Test")
		equals(t, req.Method, "POST")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			Header:     head,
		}
	})

	alert, err := c.Alerts.CreateAlert("Test", testAlert)
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestTestAlert(t *testing.T) {

	testAlert := Alert{"Test", true}
	body, _ := json.Marshal(testAlert)
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			Header:     head,
		}
	})

	alert, err := c.Alerts.TestAlert()
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestSetAlert(t *testing.T) {

	testAlert := Alert{"Test", true}
	body, _ := json.Marshal(testAlert)
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/Test")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			Header:     head,
		}
	})

	alert, err := c.Alerts.SetAlert("Test", testAlert)
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestEnableAlert(t *testing.T) {

	testAlert := Alert{"Test", true}
	body, _ := json.Marshal(testAlert)
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/Test")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			Header:     head,
		}
	})

	alert, err := c.Alerts.EnableAlert("Test")
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestDisableAlert(t *testing.T) {

	testAlert := Alert{"Test", false}
	body, _ := json.Marshal(testAlert)
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/Test")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			Header:     head,
		}
	})

	alert, err := c.Alerts.DisableAlert("Test")
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestDeleteAlert(t *testing.T) {

	testAlert := Alert{"Test", true}
	body, _ := json.Marshal(testAlert)
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/Test")
		equals(t, req.Method, "DELETE")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			Header:     head,
		}
	})

	alert, err := c.Alerts.DeleteAlert("Test")
	ok(t, err)
	equals(t, &testAlert, alert)
}

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
