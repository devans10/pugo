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

func TestGetSMTP(t *testing.T) {

	restVersion := "1.15"
	testSMTP := SMTP{
		Username:     "mailuser",
		Password:     "****",
		RelayHost:    "mail.example.com",
		SenderDomain: "fa@example.com",
	}

	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/smtp", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetSMTP(restVersion))),
			Header:     head,
		}
	})

	s, err := c.SMTP.GetSMTP()
	ok(t, err)
	equals(t, &testSMTP, s)
}

func TestSetSMTP(t *testing.T) {

	restVersion := "1.15"
	testSMTP := SMTP{
		Username:     "mailuser",
		Password:     "****",
		RelayHost:    "mail.example.com",
		SenderDomain: "fa@example.com",
	}

	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/smtp", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respSetSMTP(restVersion))),
			Header:     head,
		}
	})

	s, err := c.SMTP.SetSMTP(testSMTP)
	ok(t, err)
	equals(t, &testSMTP, s)
}

// Acceptance Tests

func TestAccGetSmtp(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.SMTP.GetSMTP()
	if err != nil {
		t.Fatalf("error getting Smtp: %s", err)
	}
}
