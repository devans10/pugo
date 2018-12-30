// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
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

	t.Run("CreateHostgroup_basic", testAccCreateHostgroup_basic(c))
	t.Run("GetHostgroup", testAccGetHostgroup(c))
	t.Run("GetHostgroup_withParams", testAccGetHostgroup_withParams(c))
	t.Run("DeleteHostgroup", testAccDeleteHostgroup(c))

	testhosts := []string{testhost1, testhost2}
	hostlist := map[string][]string{"hostlist": testhosts}
	t.Run("CreateHostgroup_withHosts", testAccCreateHostgroup_withHosts(c, hostlist))
	t.Run("ConnectVolumeToHostgroup", testAccConnectVolumeToHostgroup(c, testvol))
	t.Run("AddHostgroupToPgroup", testAccAddHostgroup(c, testpgroup))
	t.Run("RemoveHostgroupFromPgroup", testAccRemoveHostgroup(c, testpgroup))
	t.Run("ListHostgroupConnections", testAccListHostgroupConnections(c))
	t.Run("ListHostgroups", testAccListHostgroups(c))
	t.Run("ListHostgroups_withParams", testAccListHostgroups_withParams(c))
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

func testAccCreateHostgroup_basic(c *Client) func(*testing.T) {
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

func testAccGetHostgroup_withParams(c *Client) func(*testing.T) {
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

func testAccCreateHostgroup_withHosts(c *Client, hostlist map[string][]string) func(*testing.T) {
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
		_, err := c.Hostgroups.ConnectHostgroup(testAccHostgroupName, volume)
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

func testAccListHostgroups_withParams(c *Client) func(*testing.T) {
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
