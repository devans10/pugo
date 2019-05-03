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

func TestListCert(t *testing.T) {

	restVersion := "1.15"
	testCert := []Certificate{Certificate{Country: "US",
		Email:     "",
		IssuedBy:  "name.purestorage.com",
		IssuedTo:  "name.purestorage.com",
		KeySize:   2048,
		Locality:  "Mountain View",
		Name:      "kmip",
		Org:       "Pure Storage Inc.",
		OrgUnit:   "",
		State:     "CA",
		Status:    "self-signed",
		ValidFrom: "2017-12-16T05:12:47Z",
		ValidTo:   "2027-12-14T05:12:47Z",
	},
		{Country: "US",
			Email:     "",
			IssuedBy:  "db.purestorage.com",
			IssuedTo:  "db.purestorage.com",
			KeySize:   2048,
			Locality:  "Mountain View",
			Name:      "management",
			Org:       "Pure Storage Inc.",
			OrgUnit:   "",
			State:     "CA",
			Status:    "self-signed",
			ValidFrom: "2017-12-16T05:12:46Z",
			ValidTo:   "2027-12-14T05:12:46Z",
		},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetCert(restVersion))),
			Header:     head,
		}
	})

	certs, err := c.Cert.ListCert()
	ok(t, err)
	equals(t, testCert, certs)
}

func TestListCertError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetCert(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Cert.ListCert()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestGetCert(t *testing.T) {

	restVersion := "1.15"
	testCert := Certificate{Country: "US",
		Email:     "",
		IssuedBy:  "db.purestorage.com",
		IssuedTo:  "db.purestorage.com",
		KeySize:   2048,
		Locality:  "Mountain View",
		Name:      "management",
		Org:       "Pure Storage Inc.",
		OrgUnit:   "",
		State:     "CA",
		Status:    "self-signed",
		ValidFrom: "2017-12-16T05:12:46Z",
		ValidTo:   "2027-12-14T05:12:46Z",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert/management", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetCertmanagement(restVersion))),
			Header:     head,
		}
	})

	certs, err := c.Cert.GetCert("management", nil)
	ok(t, err)
	equals(t, &testCert, certs)
}

func TestGetCertError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert/management", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetCert(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Cert.GetCert("management", nil)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestGetCSR(t *testing.T) {

	restVersion := "1.15"
	testCert := Certificate{CSR: "-----BEGIN CERTIFICATE REQUEST-----\nMIICzjCCAbYCAQAwgYgxFjAUBgNVBAcMDU1vdW50YWluIFZpZXcxCzAJBgNVBAYT\nAlVTMQswCQYDVQQIDAJDQTEbMBkGA1UECwwSUHVyZSBTdG9yYWdlLCBJbmMuMRsw\nGQYDVQQDDBJkYi5wdXJlc3RvcmFnZS5jb20xGjAYBgNVBAoMEVB1cmUgU3RvcmFn\nZSBJbmMuMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApF4soHT3Vap8\nR3vhKZUYzoraYgtmcQUQXXe9DbC6joHjGvDpkPm8zGEHu7S6Tm6RqslrQpglAjq7\ndI/ltIhIz03MOdMhGc2+0eB7ZenMRyHirhTHSxQLEhOxIAPex0K7aJigrSIpKCqE\n9g0wl6LT8/C++wKw8LaeaJhdiZLR7B815v6d/o1vCcarSBNKTddY4Q+ScpqRtcub\nrqy/gY1+laNtsu7hdJTzNft83Hf0lKCLdQgFp4Qb4nMfS7j+8Cz52p/lmw+iWH+Q\nhrhFbGZbk5IjJedTKwWls4bbH2OeksUbBiZ6Gi3PZ5o7d3zmevjOp22BMjwMoaKQ\nUibzdnRaxQIDAQABoAAwDQYJKoZIhvcNAQELBQADggEBACOGXZDrr5KUYz0QZG+2\nTtbXNM872iijFtW6pAdYyyQCYddkWnZenbogmmECG/ttvxCiJYhp8W/gwn6AjbBR\nCOpQg9mRcm9A0yk3v5AJFwwX1NLlgciBwN0niex8SDlSUtkeez0Z/34UH7tWjQhG\nnVso5JlfxTa11bbEe4J6vxWPeScwb9xYFhFiPZDFVGiuj4cK121ElHh+FO9RZn1J\np2VBf4Gvo2fEA71BbTIPG9FqcGskwbAWPGOXcEwhiLrHhubs0RzyoKdenyh/+7y1\nJ5jOi0WyJcP7MFCIBDNz/wjxTAlTFKpAUCIBbBFi73mImQgTC7izbcJpk11Q588K\nq5M=\n-----END CERTIFICATE REQUEST-----\n"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert/certificate_signing_request/management", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetCertCSR(restVersion))),
			Header:     head,
		}
	})

	certs, err := c.Cert.GetCSR("management", nil)
	ok(t, err)
	equals(t, &testCert, certs)
}

func TestGetCSRError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert/certificate_signing_request/management", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetCertCSR(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Cert.GetCSR("management", nil)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestCreateCert(t *testing.T) {

	restVersion := "1.15"
	testCert := Certificate{Country: "",
		Email:     "",
		IssuedBy:  "db.example.com",
		IssuedTo:  "db.example.com",
		KeySize:   2048,
		Locality:  "",
		Name:      "new_cert",
		Org:       "Pure Storage, Inc.",
		OrgUnit:   "Pure Storage, Inc.",
		State:     "FL",
		Status:    "self-signed",
		ValidFrom: "2017-12-16T05:13:09Z",
		ValidTo:   "2027-12-14T05:13:09Z",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert/new_cert", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostCertcert(restVersion))),
			Header:     head,
		}
	})

	certs, err := c.Cert.CreateCert("new_cert", nil)
	ok(t, err)
	equals(t, &testCert, certs)
}

func TestCreateCertError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert/new_cert", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetCertCSR(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Cert.CreateCert("new_cert", nil)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestSetCert(t *testing.T) {

	restVersion := "1.15"
	testCert := Certificate{Country: "US",
		Email:     "",
		IssuedBy:  "db.purestorage.com",
		IssuedTo:  "db.purestorage.com",
		KeySize:   2048,
		Locality:  "Mountain View",
		Name:      "management",
		Org:       "Pure Storage Inc.",
		OrgUnit:   "Pure Storage, Inc.",
		State:     "CA",
		Status:    "self-signed",
		ValidFrom: "2017-12-16T05:13:08Z",
		ValidTo:   "2027-12-14T05:13:08Z",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert/management", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutCertcert(restVersion))),
			Header:     head,
		}
	})

	certs, err := c.Cert.SetCert("management", nil)
	ok(t, err)
	equals(t, &testCert, certs)
}

func TestSetCertError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert/management", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutCertcert(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Cert.SetCert("management", nil)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestDeleteCert(t *testing.T) {

	restVersion := "1.15"
	testCert := Certificate{Name: "new_cert"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert/new_cert", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteCertcert(restVersion))),
			Header:     head,
		}
	})

	certs, err := c.Cert.DeleteCert("new_cert")
	ok(t, err)
	equals(t, &testCert, certs)
}

func TestDeleteCertError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/cert/management", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteCertcert(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Cert.DeleteCert("management")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}
func TestAccCert(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)
	cert := "test"
	data := make(map[string]interface{})
	data["common_name"] = "pure.example.com"
	data["self_signed"] = true
	data["state"] = "PA"

	t.Run("ListCert", testAccListCert(c))
	t.Run("CreateCert", testAccCreateCert(cert, data, c))
	t.Run("GetCert", testAccGetCert(cert, c))
	t.Run("GetCSR", testAccGetCSR(cert, c))
	t.Run("DeleteCert", testAccDeleteCert(cert, c))

}

func testAccListCert(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		if _, err := c.Cert.ListCert(); err != nil {
			t.Fatalf("error listing cert: %s", err)
		}
	}
}

func testAccCreateCert(name string, data interface{}, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Cert.CreateCert(name, data)
		if err != nil {
			t.Fatalf("error creating cert: %s", err)
		}
	}
}

func testAccGetCert(name string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Cert.GetCert(name, nil)
		if err != nil {
			t.Fatalf("error getting cert: %s", err)
		}
	}
}

func testAccGetCSR(name string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Cert.GetCSR(name, nil)
		if err != nil {
			t.Fatalf("error getting csr: %s", err)
		}
	}
}

func testAccDeleteCert(name string, c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Cert.DeleteCert(name)
		if err != nil {
			t.Fatalf("error deleting cert: %s", err)
		}
	}
}
