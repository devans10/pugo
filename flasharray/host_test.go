// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
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

	t.Run("CreateHost_basic", testAccCreateHost_basic(c))
	t.Run("GetHost", testAccGetHost(c))
	t.Run("GetHost_withParams", testAccGetHost_withParams(c))
	t.Run("DeleteHost", testAccDeleteHost(c))

	wwns := []string{"0000999900009999"}
	wwnlist := map[string][]string{"wwnlist": wwns}
	t.Run("CreateHostWithWWN", testAccCreateHost_withWWN(c, wwnlist))
	t.Run("ConnectVolumeToHost", testAccConnectVolumeToHost(c, testvol))
	t.Run("AddHostToProtectionGroup", testAccAddHost(c, testpgroup))
	t.Run("RemoveHostFromProtectionGroup", testAccRemoveHost(c, testpgroup))
	t.Run("ListHostConnections", testAccListHostConnections(c))
	t.Run("ListHosts", testAccListHosts(c))
	t.Run("ListHosts_withParams", testAccListHosts_withParams(c))
	t.Run("RenameHost", testAccRenameHost(c, "testAcchostnew"))
	c.Hosts.RenameHost("testAcchostnew", testAccHostName)
	t.Run("RemoveVolumeFromHost", testAccDisconnectVolumeFromHost(c, testvol))
	t.Run("DeleteHost", testAccDeleteHost(c))

	c.Volumes.DeleteVolume(testvol)
	c.Volumes.EradicateVolume(testvol)
	c.Protectiongroups.DestroyProtectiongroup(testpgroup)
	c.Protectiongroups.EradicateProtectiongroup(testpgroup)
}

func testAccCreateHost_basic(c *Client) func(*testing.T) {
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

func testAccGetHost_withParams(c *Client) func(*testing.T) {
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
func testAccCreateHost_withWWN(c *Client, wwnlist map[string][]string) func(*testing.T) {
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
		_, err := c.Hosts.ConnectHost(testAccHostName, volume)
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
		_, err := c.Hosts.ListHostConnections(testAccHostName)
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

func testAccListHosts_withParams(c *Client) func(*testing.T) {
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
