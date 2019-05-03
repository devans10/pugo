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

const testAccHostgroupName = "testAcchgroup"

func TestAccHostgroups(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	testhost1 := "testacchgrouphost1"
	testhost2 := "testacchgrouphost2"
	testvol := "testacchgroupvol1"
	testpgroup := "testacchgrouppgroup1"

	c.Hosts.CreateHost(testhost1, nil)
	c.Hosts.CreateHost(testhost2, nil)
	c.Volumes.CreateVolume(testvol, 1024000000)
	c.Protectiongroups.CreateProtectiongroup(testpgroup, nil)

	t.Run("CreateHostgroup_basic", testAccCreateHostgroupBasic(c))
	t.Run("GetHostgroup", testAccGetHostgroup(c))
	t.Run("GetHostgroup_withParams", testAccGetHostgroupWithParams(c))
	t.Run("GetHostgroup_withAction", testAccGetHostgroupWithParamAction(c))
	t.Run("DeleteHostgroup", testAccDeleteHostgroup(c))

	testhosts := []string{testhost1, testhost2}
	hostlist := map[string][]string{"hostlist": testhosts}
	t.Run("CreateHostgroup_withHosts", testAccCreateHostgroupWithHosts(c, hostlist))
	t.Run("ConnectVolumeToHostgroup", testAccConnectVolumeToHostgroup(c, testvol))
	t.Run("AddHostgroupToPgroup", testAccAddHostgroup(c, testpgroup))
	t.Run("RemoveHostgroupFromPgroup", testAccRemoveHostgroup(c, testpgroup))
	t.Run("ListHostgroupConnections", testAccListHostgroupConnections(c))
	t.Run("ListHostgroups", testAccListHostgroups(c))
	t.Run("ListHostgroups_withParams", testAccListHostgroupsWithParams(c))
	t.Run("RenameHostgroup", testAccRenameHostgroup(c, "testacchgroupnew"))
	c.Hostgroups.RenameHostgroup("testacchgroupnew", testAccHostgroupName)
	t.Run("DisconnectVolumeFromHostgroup", testAccDisconnectVolumeFromHostgroup(c, testvol))
	testhosts = []string{}
	hostlist = map[string][]string{"hostlist": testhosts}
	t.Run("RemoveHostsFromHostgroup", testAccRemoveHostsFromHostgroup(c, hostlist))
	t.Run("DeleteHostgroup", testAccDeleteHostgroup(c))

	c.Hosts.DeleteHost(testhost1)
	c.Hosts.DeleteHost(testhost2)
	c.Volumes.DeleteVolume(testvol)
	c.Volumes.EradicateVolume(testvol)
	c.Protectiongroups.DestroyProtectiongroup(testpgroup)
	c.Protectiongroups.EradicateProtectiongroup(testpgroup)
}

func testAccCreateHostgroupBasic(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Hostgroups.CreateHostgroup(testAccHostgroupName, nil)
		if err != nil {
			t.Fatalf("error creating hostgroup %s: %s", testAccHostgroupName, err)
		}

		if h.Name != testAccHostgroupName {
			t.Fatalf("expected: %s; got %s", testAccHostgroupName, h.Name)
		}
	}
}

func testAccGetHostgroup(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Hostgroups.GetHostgroup(testAccHostgroupName, nil)
		if err != nil {
			t.Fatalf("error getting hostgroup %s: %s", testAccHostgroupName, err)
		}

		if h.Name != testAccHostgroupName {
			t.Fatalf("expected: %s; got %s", testAccHostgroupName, h.Name)
		}
	}
}

func testAccGetHostgroupWithParams(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		params := map[string]string{"space": "true"}
		h, err := c.Hostgroups.GetHostgroup(testAccHostgroupName, params)
		if err != nil {
			t.Fatalf("error getting hostgroup %s: %s", testAccHostgroupName, err)
		}

		if h.Name != testAccHostgroupName {
			t.Fatalf("expected: %s; got %s", testAccHostgroupName, h.Name)
		}
	}
}

func testAccGetHostgroupWithParamAction(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Hostgroups.GetHostgroup(testAccHostgroupName, map[string]string{"action": "monitor"})
		if err != nil {
			t.Fatalf("error getting host %s: %s", testAccHostgroupName, err)
		}

		if h.Name != testAccHostgroupName {
			t.Fatalf("expected: %s; got %s", testAccHostgroupName, h.Name)
		}
		if h.Time == "" {
			t.Fatalf("time property did not exist")
		}
	}
}

func testAccCreateHostgroupWithHosts(c *Client, hostlist map[string][]string) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Hostgroups.CreateHostgroup(testAccHostgroupName, hostlist)
		if err != nil {
			t.Fatalf("error creating hostgroup %s with hosts: %s", testAccHostgroupName, err)
		}

		if h.Name != testAccHostgroupName {
			t.Fatalf("expected: %s; got %s", testAccHostgroupName, h.Name)
		}
	}
}

func testAccConnectVolumeToHostgroup(c *Client, volume string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hostgroups.ConnectHostgroup(testAccHostgroupName, volume, nil)
		if err != nil {
			t.Fatalf("error connecting volume to hostgroup %s: %s", testAccHostgroupName, err)
		}

	}
}

func testAccAddHostgroup(c *Client, pgroup string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hostgroups.AddHostgroup(testAccHostgroupName, pgroup)
		if err != nil {
			t.Fatalf("error adding hostgroup %s to pgroup %s: %s", testAccHostgroupName, pgroup, err)
		}
	}
}

func testAccRemoveHostgroup(c *Client, pgroup string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hostgroups.RemoveHostgroup(testAccHostgroupName, pgroup)
		if err != nil {
			t.Fatalf("error adding hostgroup %s to pgroup %s: %s", testAccHostgroupName, pgroup, err)
		}
	}
}

func testAccListHostgroupConnections(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hostgroups.ListHostgroupConnections(testAccHostgroupName)
		if err != nil {
			t.Fatalf("error listing hostgroup connections for %s: %s", testAccHostgroupName, err)
		}
	}
}

func testAccListHostgroups(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hostgroups.ListHostgroups(nil)
		if err != nil {
			t.Fatalf("error listing hostgroups: %s", err)
		}
	}
}

func testAccListHostgroupsWithParams(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		params := map[string]string{"space": "true"}
		_, err := c.Hostgroups.ListHostgroups(params)
		if err != nil {
			t.Fatalf("error listing hostgroups: %s", err)
		}
	}
}

func testAccRenameHostgroup(c *Client, name string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hostgroups.RenameHostgroup(testAccHostgroupName, name)
		if err != nil {
			t.Fatalf("error renaming hostgroup %s to %s: %s", testAccHostgroupName, name, err)
		}
	}
}

func testAccDisconnectVolumeFromHostgroup(c *Client, volume string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hostgroups.DisconnectHostgroup(testAccHostgroupName, volume)
		if err != nil {
			t.Fatalf("error disconnecting volume to hostgroup %s: %s", testAccHostgroupName, err)
		}

	}
}

func testAccRemoveHostsFromHostgroup(c *Client, hostlist map[string][]string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hostgroups.SetHostgroup(testAccHostgroupName, hostlist)
		if err != nil {
			t.Fatalf("error removing hosts from hostgroup %s: %s", testAccHostgroupName, err)
		}

	}
}
func testAccDeleteHostgroup(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Hostgroups.DeleteHostgroup(testAccHostgroupName)
		if err != nil {
			t.Fatalf("error deleting hostgroup: %s", err)
		}
	}
}

func TestConnectHostgroup(t *testing.T) {
	restVersion := "1.15"
	testConn := ConnectedVolume{Vol: "v3", Name: "hg3", Lun: 254}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg3/volume/v3", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostHgrouphgroupVolumevol(restVersion))),
			Header:     head,
		}
	})

	conn, err := c.Hostgroups.ConnectHostgroup("hg3", "v3", nil)
	ok(t, err)
	equals(t, &testConn, conn)
}

func TestConnectHostgroupError(t *testing.T) {
	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg3/volume/v3", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostHgrouphgroupVolumevol(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hostgroups.ConnectHostgroup("hg3", "v3", nil)
	if err == nil {
		t.Errorf("error not returned on 500 response")
	}
}

func TestCreateHostgroup(t *testing.T) {
	restVersion := "1.15"
	testHostgroup := Hostgroup{Name: "hg3", Hosts: []string{}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/h3", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostHgrouphgroup(restVersion))),
			Header:     head,
		}
	})

	hgroup, err := c.Hostgroups.CreateHostgroup("h3", testHostgroup)
	ok(t, err)
	equals(t, &testHostgroup, hgroup)
}

func TestCreateHostgroupError(t *testing.T) {
	restVersion := "1.15"
	testHostgroup := Hostgroup{Name: "hg3", Hosts: []string{}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/h3", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostHgrouphgroup(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hostgroups.CreateHostgroup("h3", testHostgroup)
	if err == nil {
		t.Errorf("error not returned on 500 response")
	}
}

func TestDeleteHostgroup(t *testing.T) {
	restVersion := "1.15"
	testHostgroup := Hostgroup{Name: "hg4"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/h4", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteHgrouphgroup(restVersion))),
			Header:     head,
		}
	})

	hgroup, err := c.Hostgroups.DeleteHostgroup("h4")
	ok(t, err)
	equals(t, &testHostgroup, hgroup)
}

func TestDeleteHostgroupError(t *testing.T) {
	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/h4", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteHgrouphgroup(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hostgroups.DeleteHostgroup("h4")
	if err == nil {
		t.Errorf("error not returned on 500 response")
	}
}

func TestDisconnectHostgroup(t *testing.T) {
	restVersion := "1.15"
	testConn := ConnectedVolume{Vol: "v3", Name: "hg4"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg4/volume/v3", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteHgrouphgroupVolumevol(restVersion))),
			Header:     head,
		}
	})

	conn, err := c.Hostgroups.DisconnectHostgroup("hg4", "v3")
	ok(t, err)
	equals(t, &testConn, conn)
}

func TestDisconnectHostgroupError(t *testing.T) {
	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg4/volume/v3", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteHgrouphgroupVolumevol(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hostgroups.DisconnectHostgroup("hg4", "v3")
	if err == nil {
		t.Errorf("error not returned on 500 response")
	}
}

func TestGetHostgroup(t *testing.T) {
	restVersion := "1.15"
	testHostgroup := Hostgroup{Name: "hg1", Hosts: []string{"h1", "h2"}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg1", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHgrouphgroup(restVersion))),
			Header:     head,
		}
	})

	hgroup, err := c.Hostgroups.GetHostgroup("hg1", nil)
	ok(t, err)
	equals(t, &testHostgroup, hgroup)
}

func TestGetHostgroupError(t *testing.T) {
	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg1", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHgrouphgroup(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hostgroups.GetHostgroup("hg1", nil)
	if err == nil {
		t.Errorf("error not returned on 500 response")
	}
}

//TODO: get output
//func TestAddHostgroup(t *testing.T) {}

//TODO: get output
//func TestRemoveHostgroup(t *testing.T) {}

func TestGetHostgroupConnections(t *testing.T) {
	restVersion := "1.15"
	testConn := []HostgroupConnection{HostgroupConnection{Name: "hg1", Vol: "v1", Lun: 254}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg1/volume", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHgrouphgroupVolume(restVersion))),
			Header:     head,
		}
	})

	conn, err := c.Hostgroups.ListHostgroupConnections("hg1")
	ok(t, err)
	equals(t, testConn, conn)
}

func TestGetHostgroupConnectionsError(t *testing.T) {
	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg1/volume", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHgrouphgroupVolume(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hostgroups.ListHostgroupConnections("hg1")
	if err == nil {
		t.Errorf("error not returned on 500 response")
	}
}

func TestListHostgroups(t *testing.T) {
	restVersion := "1.15"
	testHostgroup := []Hostgroup{Hostgroup{Name: "hg1", Hosts: []string{"h1", "h2"}},
		Hostgroup{Name: "hg2", Hosts: []string{}},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHgroup(restVersion))),
			Header:     head,
		}
	})

	hgroup, err := c.Hostgroups.ListHostgroups(nil)
	ok(t, err)
	equals(t, testHostgroup, hgroup)
}

func TestListHostgroupsError(t *testing.T) {
	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHgroup(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hostgroups.ListHostgroups(nil)
	if err == nil {
		t.Errorf("error not returned on 500 response")
	}
}

func TestRenameHostgroup(t *testing.T) {
	restVersion := "1.15"
	testHostgroup := Hostgroup{Name: "hg4_renamed", Hosts: []string{"h3"}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg4", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutHgrouphgroupRename(restVersion))),
			Header:     head,
		}
	})

	hgroup, err := c.Hostgroups.RenameHostgroup("hg4", "hg4_renamed")
	ok(t, err)
	equals(t, &testHostgroup, hgroup)
}

func TestRenameHostgroupError(t *testing.T) {
	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg4", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutHgrouphgroupRename(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hostgroups.RenameHostgroup("hg4", "hg4_renamed")
	if err == nil {
		t.Errorf("error not returned on 500 response")
	}
}

func TestSetHostgroup(t *testing.T) {
	restVersion := "1.15"
	testHostgroup := Hostgroup{Name: "hg4", Hosts: []string{"h3"}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg4", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutHgrouphgroup(restVersion))),
			Header:     head,
		}
	})

	hgroup, err := c.Hostgroups.SetHostgroup("hg4", testHostgroup)
	ok(t, err)
	equals(t, &testHostgroup, hgroup)
}

func TestSetHostgroupError(t *testing.T) {
	restVersion := "1.15"
	testHostgroup := Hostgroup{Name: "hg4", Hosts: []string{"h3"}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/hgroup/hg4", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutHgrouphgroup(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Hostgroups.SetHostgroup("hg4", testHostgroup)
	if err == nil {
		t.Errorf("error not returned on 500 response")
	}
}
