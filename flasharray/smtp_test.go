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
