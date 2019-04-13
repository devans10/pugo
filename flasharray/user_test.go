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

func TestListAdmins(t *testing.T) {

	restVersion := "1.15"
	testUser := []User{
		User{
			Name: "pureuser",
			Role: "admin",
			Type: "local",
		},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/admin", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetAdmin(restVersion))),
			Header:     head,
		}
	})

	user, err := c.Users.ListAdmins()
	ok(t, err)
	equals(t, testUser, user)
}

func TestGetAdmin(t *testing.T) {

	restVersion := "1.15"
	testUser := User{
		Name: "pureuser",
		Role: "admin",
		Type: "local",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/admin/pureuser", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetAdminadmin(restVersion))),
			Header:     head,
		}
	})

	user, err := c.Users.GetAdmin("pureuser")
	ok(t, err)
	equals(t, &testUser, user)
}

func TestCreateAdmin(t *testing.T) {

	restVersion := "1.15"
	testUser := User{
		Name: "pureuser",
		Role: "admin",
		Type: "local",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/admin/pureuser", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostAdminadmin(restVersion))),
			Header:     head,
		}
	})

	user, err := c.Users.CreateAdmin("pureuser")
	ok(t, err)
	equals(t, &testUser, user)
}

func TestSetAdmin(t *testing.T) {

	restVersion := "1.15"
	testUser := User{
		Name: "pureuser",
		Role: "admin",
		Type: "local",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/admin/pureuser", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutAdminadmin(restVersion))),
			Header:     head,
		}
	})

	user, err := c.Users.SetAdmin("pureuser", testUser)
	ok(t, err)
	equals(t, &testUser, user)
}

func TestDeleteAdmin(t *testing.T) {

	restVersion := "1.15"
	testUser := User{
		Name: "pureuser",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/admin/pureuser", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteAdminadmin(restVersion))),
			Header:     head,
		}
	})

	user, err := c.Users.DeleteAdmin("pureuser")
	ok(t, err)
	equals(t, &testUser, user)
}

func TestCreateAPIToken(t *testing.T) {

	restVersion := "1.15"
	testToken := Token{
		APIToken: "****",
		Created:  "2017-12-16T05:10:11Z",
		Expires:  "",
		Name:     "pureuser",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/admin/pureuser/apitoken", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostAdminadminAPIToken(restVersion))),
			Header:     head,
		}
	})

	token, err := c.Users.CreateAPIToken("pureuser")
	ok(t, err)
	equals(t, &testToken, token)
}

func TestGetAPIToken(t *testing.T) {

	restVersion := "1.15"
	testToken := Token{
		APIToken: "****",
		Created:  "2017-12-16T05:10:11Z",
		Expires:  "",
		Name:     "pureuser",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/admin/pureuser/apitoken", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetAdminadminAPIToken(restVersion))),
			Header:     head,
		}
	})

	token, err := c.Users.GetAPIToken("pureuser")
	ok(t, err)
	equals(t, &testToken, token)
}

func TestDeleteAPIToken(t *testing.T) {

	restVersion := "1.15"
	testToken := Token{
		Name: "pureuser",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/admin/pureuser/apitoken", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteAdminadminAPIToken(restVersion))),
			Header:     head,
		}
	})

	token, err := c.Users.DeleteAPIToken("pureuser")
	ok(t, err)
	equals(t, &testToken, token)
}

func TestListAPITokens(t *testing.T) {

	restVersion := "1.15"
	testToken := []Token{
		Token{
			APIToken: "****",
			Created:  "2017-12-16T05:11:04Z",
			Expires:  "",
			Name:     "os76",
		},
		Token{
			APIToken: "****",
			Created:  "2017-12-16T05:10:11Z",
			Expires:  "",
			Name:     "pureuser",
		},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/admin?api_token=true", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetAdminAPIToken(restVersion))),
			Header:     head,
		}
	})

	token, err := c.Users.ListAPITokens()
	ok(t, err)
	equals(t, testToken, token)
}

func TestRefreshAdmin(t *testing.T) {

	restVersion := "1.15"
	testUser := User{
		Name: "pureuser",
		Role: "admin",
		Type: "local",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/admin/pureuser", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutAdminadmin(restVersion))),
			Header:     head,
		}
	})

	user, err := c.Users.RefreshAdmin("pureuser")
	ok(t, err)
	equals(t, &testUser, user)
}

func TestRefreshAdmins(t *testing.T) {

	restVersion := "1.15"
	testUser := User{
		Name: "pureuser",
		Role: "admin",
		Type: "local",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/admin", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutAdminadmin(restVersion))),
			Header:     head,
		}
	})

	user, err := c.Users.RefreshAdmins()
	ok(t, err)
	equals(t, &testUser, user)
}
