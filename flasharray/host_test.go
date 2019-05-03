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

const testAccHostName = "testAcchost"

func TestAccHosts(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	testvol := "testacchostvol1"
	testpgroup := "testacchostpgroup"

	c.Volumes.CreateVolume(testvol, 1024000000)
	c.Protectiongroups.CreateProtectiongroup(testpgroup, nil)

	t.Run("CreateHost_basic", testAccCreateHostBasic(c))
	t.Run("GetHost", testAccGetHost(c))
	t.Run("GetHost_withParams", testAccGetHostWithParams(c))
	t.Run("GetHost_withAction", testAccGetHostWithParamAction(c))
	t.Run("DeleteHost", testAccDeleteHost(c))

	wwns := []string{"0000999900009999"}
	wwnlist := map[string][]string{"wwnlist": wwns}
	t.Run("CreateHostWithWWN", testAccCreateHostWithWWN(c, wwnlist))
	t.Run("ConnectVolumeToHost", testAccConnectVolumeToHost(c, testvol))
	t.Run("AddHostToProtectionGroup", testAccAddHost(c, testpgroup))
	t.Run("RemoveHostFromProtectionGroup", testAccRemoveHost(c, testpgroup))
	t.Run("ListHostConnections", testAccListHostConnections(c))
	t.Run("ListHosts", testAccListHosts(c))
	t.Run("ListHosts_withParams", testAccListHostsWithParams(c))
	t.Run("RenameHost", testAccRenameHost(c, "testAcchostnew"))
	c.Hosts.RenameHost("testAcchostnew", testAccHostName)
	t.Run("RemoveVolumeFromHost", testAccDisconnectVolumeFromHost(c, testvol))
	t.Run("DeleteHost", testAccDeleteHost(c))

	c.Volumes.DeleteVolume(testvol)
	c.Volumes.EradicateVolume(testvol)
	c.Protectiongroups.DestroyProtectiongroup(testpgroup)
	c.Protectiongroups.EradicateProtectiongroup(testpgroup)
}

func testAccCreateHostBasic(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Hosts.CreateHost(testAccHostName, nil)
		if err != nil {
			t.Fatalf("error creating hostgroup %s: %s", testAccHostName, err)
		}

		if h.Name != testAccHostName {
			t.Fatalf("expected: %s; got %s", testAccHostName, h.Name)
		}
	}
}

func testAccGetHost(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Hosts.GetHost(testAccHostName, nil)
		if err != nil {
			t.Fatalf("error getting host %s: %s", testAccHostName, err)
		}

		if h.Name != testAccHostName {
			t.Fatalf("expected: %s; got %s", testAccHostName, h.Name)
		}
	}
}

func testAccGetHostWithParams(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Hosts.GetHost(testAccHostName, map[string]string{"personality": "true"})
		if err != nil {
			t.Fatalf("error getting host %s: %s", testAccHostName, err)
		}

		if h.Name != testAccHostName {
			t.Fatalf("expected: %s; got %s", testAccHostName, h.Name)
		}
	}
}

func testAccGetHostWithParamAction(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Hosts.GetHost(testAccHostName, map[string]string{"action": "monitor"})
		if err != nil {
			t.Fatalf("error getting host %s: %s", testAccHostName, err)
		}

		if h.Name != testAccHostName {
			t.Fatalf("expected: %s; got %s", testAccHostName, h.Name)
		}
		if h.Time == "" {
			t.Fatalf("time property did not exist")
		}
	}
}

func testAccCreateHostWithWWN(c *Client, wwnlist map[string][]string) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Hosts.CreateHost(testAccHostName, wwnlist)
		if err != nil {
			t.Fatalf("error creating host %s with wwn: %s", testAccHostName, err)
		}

		if h.Name != testAccHostName {
			t.Fatalf("expected: %s; got %s", testAccHostName, h.Name)
		}
	}
}

func testAccConnectVolumeToHost(c *Client, volume string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hosts.ConnectHost(testAccHostName, volume, nil)
		if err != nil {
			t.Fatalf("error connecting volume to host %s: %s", testAccHostName, err)
		}

	}
}

func testAccAddHost(c *Client, pgroup string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hosts.AddHost(testAccHostName, pgroup)
		if err != nil {
			t.Fatalf("error adding host %s to pgroup %s: %s", testAccHostName, pgroup, err)
		}
	}
}

func testAccRemoveHost(c *Client, pgroup string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hosts.RemoveHost(testAccHostName, pgroup)
		if err != nil {
			t.Fatalf("error adding host %s to pgroup %s: %s", testAccHostName, pgroup, err)
		}
	}
}

func testAccListHostConnections(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hosts.ListHostConnections(testAccHostName, nil)
		if err != nil {
			t.Fatalf("error listing host connections for %s: %s", testAccHostName, err)
		}
	}
}

func testAccListHosts(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hosts.ListHosts(nil)
		if err != nil {
			t.Fatalf("error listing hosts: %s", err)
		}
	}
}

func testAccListHostsWithParams(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hosts.ListHosts(map[string]string{"personality": "true"})
		if err != nil {
			t.Fatalf("error listing hosts: %s", err)
		}
	}
}

func testAccRenameHost(c *Client, name string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hosts.RenameHost(testAccHostName, name)
		if err != nil {
			t.Fatalf("error renaming host %s to %s: %s", testAccHostName, name, err)
		}
	}
}

func testAccDisconnectVolumeFromHost(c *Client, volume string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Hosts.DisconnectHost(testAccHostName, volume)
		if err != nil {
			t.Fatalf("error disconnecting volume from host %s: %s", testAccHostName, err)
		}

	}
}

func testAccDeleteHost(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Hosts.DeleteHost(testAccHostName)
		if err != nil {
			t.Fatalf("error deleting host: %s", err)
		}
	}
}

func TestConnectHost(t *testing.T) {
	restVersion := "1.15"
	testConn := ConnectedVolume{Vol: "v5_renamed", Name: "h4", Lun: 1}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/host/h4/volume/v5_renamed", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostHosthostVolumevol(restVersion))),
			Header:     head,
		}
	})

	conn, err := c.Hosts.ConnectHost("h4", "v5_renamed", nil)
	ok(t, err)
	equals(t, &testConn, conn)
}

func TestCreateHost(t *testing.T) {
	restVersion := "1.15"
	testHost := Host{Name: "h4", Wwn: []string{"0000999900009999"}, Iqn: []string{}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/host/h4", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostHosthost(restVersion))),
			Header:     head,
		}
	})

	host, err := c.Hosts.CreateHost("h4", testHost)
	ok(t, err)
	equals(t, &testHost, host)
}

func TestDeleteHost(t *testing.T) {
	restVersion := "1.15"
	testHost := Host{Name: "h5"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/host/h5", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteHosthost(restVersion))),
			Header:     head,
		}
	})

	host, err := c.Hosts.DeleteHost("h5")
	ok(t, err)
	equals(t, &testHost, host)
}

func TestDisconnectHost(t *testing.T) {
	restVersion := "1.15"
	testConn := ConnectedVolume{Vol: "v5_renamed", Name: "h5"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/host/h5/volume/v5_renamed", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteHosthostVolumevol(restVersion))),
			Header:     head,
		}
	})

	conn, err := c.Hosts.DisconnectHost("h5", "v5_renamed")
	ok(t, err)
	equals(t, &testConn, conn)
}

func TestGetHost(t *testing.T) {
	restVersion := "1.15"
	testHost := Host{Name: "h3", Wwn: []string{}, Iqn: []string{}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/host/h3", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHosthost(restVersion))),
			Header:     head,
		}
	})

	host, err := c.Hosts.GetHost("h3", nil)
	ok(t, err)
	equals(t, &testHost, host)
}

//TODO: get response
/*
func TestAddHost(t *testing.T) {
	restVersion := "1.15"
	testHost := Host{Name: "h3", Wwn: []string{}, Iqn: []string{}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/host/h3", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostHosthostPgrouppgroup(restVersion))),
			Header:     head,
		}
	})

	host, err := c.Hosts.AddHost("h3", "pg1")
	ok(t, err)
	equals(t, &testHost, host)
}*/

//TODO: get response
/*
func TestRemoveHost(t *testing.T) {
	restVersion := "1.15"
	testHost := Host{Name: "h3", Wwn: []string{}, Iqn: []string{}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/host/h3", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteHosthostPgrouppgroup(restVersion))),
			Header:     head,
		}
	})

	host, err := c.Hosts.RemoveHost("h3", "pg1")
	ok(t, err)
	equals(t, &testHost, host)
}*/

func TestListHostConnections(t *testing.T) {
	restVersion := "1.15"
	testConn := []ConnectedVolume{ConnectedVolume{Name: "h2", Vol: "v3", Lun: 7}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/host/h2/volume", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHosthostVolumePrivate(restVersion))),
			Header:     head,
		}
	})

	conn, err := c.Hosts.ListHostConnections("h2", nil)
	ok(t, err)
	equals(t, testConn, conn)
}

func TestListHosts(t *testing.T) {
	restVersion := "1.15"
	testHost := []Host{Host{Hgroup: "hg1", Iqn: []string{}, Name: "h1", Wwn: []string{"0000111122223333"}},
		Host{Hgroup: "hg1", Iqn: []string{}, Name: "h2", Wwn: []string{}},
		Host{Name: "h3", Wwn: []string{}, Iqn: []string{}},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/host", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetHost(restVersion))),
			Header:     head,
		}
	})

	host, err := c.Hosts.ListHosts(nil)
	ok(t, err)
	equals(t, testHost, host)
}

func TestRenameHost(t *testing.T) {
	restVersion := "1.15"
	testHost := Host{Name: "h4_renamed"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/host/h4", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutHosthostRename(restVersion))),
			Header:     head,
		}
	})

	host, err := c.Hosts.RenameHost("h4", "h4_renamed")
	ok(t, err)
	equals(t, &testHost, host)
}

func TestSetHost(t *testing.T) {
	restVersion := "1.15"
	testHost := Host{Name: "h4", Wwn: []string{"1111222233334444", "2222333344445555", "4444333322221111", "5555444433332222"}, Iqn: []string{}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/host/h4", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutHosthost(restVersion))),
			Header:     head,
		}
	})

	host, err := c.Hosts.SetHost("h4", testHost)
	ok(t, err)
	equals(t, &testHost, host)
}
