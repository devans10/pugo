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

func TestListPods(t *testing.T) {

	restVersion := "1.15"
	testPod := []Pod{
		Pod{
			Name:               "pod1",
			Source:             "flasharray1.example.com",
			FailoverPreference: []string{"flasharray2.example.com"},
		},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/pod", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetPod(restVersion))),
			Header:     head,
		}
	})

	pod, err := c.Pods.ListPods(nil)
	ok(t, err)
	equals(t, testPod, pod)
}

func TestGetPod(t *testing.T) {

	restVersion := "1.15"
	testPod := Pod{
		Name:               "pod1",
		Source:             "flasharray1.example.com",
		FailoverPreference: []string{"flasharray2.example.com"},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/pod/pod1", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetPodpod(restVersion))),
			Header:     head,
		}
	})

	pod, err := c.Pods.GetPod("pod1", nil)
	ok(t, err)
	equals(t, &testPod, pod)
}

func TestCreatePod(t *testing.T) {

	restVersion := "1.15"
	testPod := Pod{
		Name:               "pod1",
		Source:             "flasharray1.example.com",
		FailoverPreference: []string{"flasharray2.example.com"},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/pod/pod1", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostPodpod(restVersion))),
			Header:     head,
		}
	})

	pod, err := c.Pods.CreatePod("pod1", testPod)
	ok(t, err)
	equals(t, &testPod, pod)
}

// TODO ConnectPod

func TestSetPod(t *testing.T) {

	restVersion := "1.15"
	testPod := Pod{
		Name:               "pod1",
		Source:             "flasharray1.example.com",
		FailoverPreference: []string{"flasharray2.example.com"},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/pod/pod1", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutPodpod(restVersion))),
			Header:     head,
		}
	})

	pod, err := c.Pods.SetPod("pod1", testPod)
	ok(t, err)
	equals(t, &testPod, pod)
}

func TestRenamePod(t *testing.T) {

	restVersion := "1.15"
	testPod := Pod{
		Name:               "pod_renamed",
		Source:             "flasharray1.example.com",
		FailoverPreference: []string{"flasharray2.example.com"},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/pod/pod1", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutPodpodRename(restVersion))),
			Header:     head,
		}
	})

	pod, err := c.Pods.RenamePod("pod1", "pod_renamed")
	ok(t, err)
	equals(t, &testPod, pod)
}

// TODO RecoverPod

func TestDeletePod(t *testing.T) {

	restVersion := "1.15"
	testPod := Pod{
		Name: "pod1",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/pod/pod1", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeletePodpod(restVersion))),
			Header:     head,
		}
	})

	pod, err := c.Pods.DeletePod("pod1")
	ok(t, err)
	equals(t, &testPod, pod)
}

// Acceptance Tests

func TestAccListPods(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Pods.ListPods(nil)
	if err != nil {
		t.Fatalf("error listing pods: %s", err)
	}
}
