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

func TestListMessages(t *testing.T) {

	restVersion := "1.15"
	testMessage := []Message{}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/message", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetMessage(restVersion))),
			Header:     head,
		}
	})

	message, err := c.Messages.ListMessages(nil)
	ok(t, err)
	equals(t, testMessage, message)
}

func TestListMessagesError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/message", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetMessage(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Messages.ListMessages(nil)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

// Acceptance Tests
func TestAccListMessages(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Messages.ListMessages(nil)
	if err != nil {
		t.Fatalf("error listing messages: %s", err)
	}
}
