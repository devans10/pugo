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

// Unit Tests

func TestGetDrive(t *testing.T) {

	restVersion := "1.15"
	testDrive := Drive{Capacity: 494927872,
		Details:           "",
		LastEvacCompleted: "1970-01-01T00:00:00Z",
		LastFailure:       "1970-01-01T00:00:00Z",
		Name:              "SH0.BAY0",
		Protocol:          "SAS",
		Status:            "healthy",
		Type:              "SSD",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/drive/SH0.BAY0", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetDrivedrive(restVersion))),
			Header:     head,
		}
	})

	dr, err := c.Hardware.GetDrive("SH0.BAY0")
	ok(t, err)
	equals(t, &testDrive, dr)
}

func TestGetDriveError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/drive/SH0.BAY0", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetDrivedrive(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hardware.GetDrive("SH0.BAY0")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestListDrives(t *testing.T) {

	restVersion := "1.15"
	testDrive := []Drive{Drive{Capacity: 494927872,
		Details:           "",
		LastEvacCompleted: "1970-01-01T00:00:00Z",
		LastFailure:       "1970-01-01T00:00:00Z",
		Name:              "SH0.BAY0",
		Protocol:          "SAS",
		Status:            "healthy",
		Type:              "SSD",
	},
		Drive{Capacity: 494927872,
			Details:           "",
			LastEvacCompleted: "1970-01-01T00:00:00Z",
			LastFailure:       "1970-01-01T00:00:00Z",
			Name:              "SH0.BAY1",
			Protocol:          "SAS",
			Status:            "healthy",
			Type:              "SSD",
		},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/drive", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetDrive(restVersion))),
			Header:     head,
		}
	})

	dr, err := c.Hardware.ListDrives()
	ok(t, err)
	equals(t, testDrive, dr)
}

func TestListDrivesError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/drive", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetDrive(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hardware.ListDrives()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestGetHardware(t *testing.T) {

	restVersion := "1.15"
	testComponent := Component{Details: "",
		Identify:    "off",
		Index:       0,
		Model:       "",
		Name:        "SH0.BAY0",
		Serial:      "",
		Slot:        0,
		Speed:       0,
		Status:      "ok",
		Temperature: 0,
		Voltage:     0,
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hardware/SH0.BAY0", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHardwareComp(restVersion))),
			Header:     head,
		}
	})

	comp, err := c.Hardware.GetHardware("SH0.BAY0")
	ok(t, err)
	equals(t, &testComponent, comp)
}

func TestGetHardwareError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hardware/SH0.BAY0", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHardwareComp(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hardware.GetHardware("SH0.BAY0")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestListHardware(t *testing.T) {

	restVersion := "1.15"
	testComponent := []Component{Component{Details: "",
		Identify:    "off",
		Index:       0,
		Model:       "",
		Name:        "CT0",
		Serial:      "",
		Slot:        0,
		Speed:       0,
		Status:      "ok",
		Temperature: 0,
		Voltage:     0,
	}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hardware", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHardware(restVersion))),
			Header:     head,
		}
	})

	comp, err := c.Hardware.ListHardware()
	ok(t, err)
	equals(t, testComponent, comp)
}

func TestListHardwareError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hardware", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHardware(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hardware.ListHardware()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestSetHardware(t *testing.T) {

	restVersion := "1.15"
	testComponent := Component{
		Identify: "on",
		Index:    0,
		Name:     "SH0.BAY0",
		Slot:     0,
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hardware/SH0.BAY0", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutHardwareComp(restVersion))),
			Header:     head,
		}
	})

	data := map[string]string{"identify": "on"}
	comp, err := c.Hardware.SetHardware("SH0.BAY0", data)
	ok(t, err)
	equals(t, &testComponent, comp)
}

func TestSetHardwareError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hardware/SH0.BAY0", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutHardwareComp(restVersion))),
			Header:     head,
		}
	})

	data := map[string]string{"identify": "on"}
	_, err := c.Hardware.SetHardware("SH0.BAY0", data)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

// Acceptance Tests

func TestAccGetDrive(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	expected := "CH0.BAY0"
	h, err := c.Hardware.GetDrive(expected)
	if err != nil {
		t.Fatalf("error getting drive: %s", err)
	}

	if h.Name != expected {
		t.Fatalf("expected: %s, got: %s", expected, h.Name)
	}
}

func TestAccListDrives(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Hardware.ListDrives()
	if err != nil {
		t.Fatalf("error listing drives: %s", err)
	}
}

func TestAccGetHardware(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	expected := "CT0"
	h, err := c.Hardware.GetHardware(expected)
	if err != nil {
		t.Fatalf("error getting hardware: %s", err)
	}

	if h.Name != expected {
		t.Fatalf("expected: %s, got: %s", expected, h.Name)
	}
}

func TestAccListHardware(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Hardware.ListHardware()
	if err != nil {
		t.Fatalf("error listing hardware: %s", err)
	}
}
