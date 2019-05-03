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

func TestListAlerts(t *testing.T) {

	restVersion := "1.15"
	testAlert := []Alert{Alert{"flasharray-alerts@purestorage.com", false}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert")
		equals(t, req.Method, "GET")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetAlert(restVersion))),
			Header:     head,
		}
	})

	alert, err := c.Alerts.ListAlerts(nil)
	ok(t, err)
	equals(t, testAlert, alert)
}

func TestListAlertsError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert")
		equals(t, req.Method, "GET")
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetAlert(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Alerts.ListAlerts(nil)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestGetAlert(t *testing.T) {

	restVersion := "1.15"
	testAlert := Alert{"flasharray-alerts@purestorage.com", false}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/flasharray-alerts@purestorage.com")
		equals(t, req.Method, "GET")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetAlertaddress(restVersion))),
			Header:     head,
		}
	})

	alert, err := c.Alerts.GetAlert("flasharray-alerts@purestorage.com")
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestGetAlertError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/flasharray-alerts@purestorage.com")
		equals(t, req.Method, "GET")
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetAlertaddress(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Alerts.GetAlert("flasharray-alerts@purestorage.com")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestCreateAlert(t *testing.T) {

	restVersion := "1.15"
	testAlert := Alert{"admin@example.com", true}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "POST")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostAlertaddress(restVersion))),
			Header:     head,
		}
	})

	alert, err := c.Alerts.CreateAlert("admin@example.com", testAlert)
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestCreateAlertError(t *testing.T) {

	restVersion := "1.15"
	testAlert := Alert{"admin@example.com", true}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "POST")
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostAlertaddress(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Alerts.CreateAlert("admin@example.com", testAlert)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestTestAlert(t *testing.T) {

	restVersion := "1.15"
	testAlert := Alert{"admin@example.com", false}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutAlertaddress(restVersion))),
			Header:     head,
		}
	})

	alert, err := c.Alerts.TestAlert("admin@example.com")
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestTestAlertError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutAlertaddress(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Alerts.TestAlert("admin@example.com")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

//TODO
//func TestTestAlerts(t *testing.T) {}

func TestSetAlert(t *testing.T) {

	restVersion := "1.15"
	testAlert := Alert{"admin@example.com", false}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutAlertaddress(restVersion))),
			Header:     head,
		}
	})

	alert, err := c.Alerts.SetAlert("admin@example.com", testAlert)
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestSetAlertError(t *testing.T) {

	restVersion := "1.15"
	testAlert := Alert{"admin@example.com", false}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutAlertaddress(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Alerts.SetAlert("admin@example.com", testAlert)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestEnableAlert(t *testing.T) {

	restVersion := "1.15"
	testAlert := Alert{"admin@example.com", false}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutAlertaddress(restVersion))),
			Header:     head,
		}
	})

	alert, err := c.Alerts.EnableAlert("admin@example.com")
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestEnableAlertError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutAlertaddress(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Alerts.EnableAlert("admin@example.com")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestDisableAlert(t *testing.T) {

	restVersion := "1.15"
	testAlert := Alert{"admin@example.com", false}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutAlertaddress(restVersion))),
			Header:     head,
		}
	})

	alert, err := c.Alerts.DisableAlert("admin@example.com")
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestDisableAlertError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "PUT")
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutAlertaddress(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Alerts.DisableAlert("admin@example.com")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestDeleteAlert(t *testing.T) {

	restVersion := "1.15"
	testAlert := Alert{"admin@example.com", false}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "DELETE")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteAlertaddress(restVersion))),
			Header:     head,
		}
	})

	alert, err := c.Alerts.DeleteAlert("admin@example.com")
	ok(t, err)
	equals(t, &testAlert, alert)
}

func TestDeleteAlertError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), "https://flasharray.example.com/api/1.15/alert/admin@example.com")
		equals(t, req.Method, "DELETE")
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteAlertaddress(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Alerts.DeleteAlert("admin@example.com")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
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
