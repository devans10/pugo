// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"testing"
)

const testAccProtectiongroupName = "testAccpgroup"

func TestAccProtectiongroups(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	testhost1 := "testaccpgrouphost1"
	testvol := "testaccpgroupvol1"
	testhgroup := "testaccpgrouphgroup1"

	c.Hosts.CreateHost(testhost1, nil)
	c.Volumes.CreateVolume(testvol, 1024000000)
	c.Hostgroups.CreateHostgroup(testhgroup, nil)

	t.Run("CreateProtectiongroup_basic", testAccCreateProtectiongroup_basic(c))
	t.Run("GetProtectiongroup", testAccGetProtectiongroup(c))
	t.Run("GetProtectiongroup_withParams", testAccGetProtectiongroup_withParams(c))
	t.Run("EnablePgroupReplication", testAccEnablePgroupReplication(c))
	t.Run("DisablePgroupReplication", testAccDisablePgroupReplication(c))
	t.Run("EnablePgroupSnapshots", testAccEnablePgroupSnapshots(c))
	t.Run("DisablePgroupSnapshots", testAccDisablePgroupSnapshots(c))
	t.Run("DeleteProtectiongroup", testAccDeleteProtectiongroup(c))
	t.Run("RecoverProtectiongroup", testAccRecoverProtectiongroup(c))
	t.Run("DeleteProtectiongroup", testAccDeleteProtectiongroup(c))
	t.Run("EradicateProtectiongroup", testAccEradicateProtectiongroup(c))

	testhosts := []string{testhost1}
	hostlist := map[string][]string{"hostlist": testhosts}
	t.Run("CreateProtectiongroup_withHosts", testAccCreateProtectiongroup_withHosts(c, hostlist))
	t.Run("ListProtectiongroups", testAccListProtectiongroups(c))
	t.Run("ListProtectiongroups_withParams", testAccListProtectiongroups_withParams(c))
	t.Run("RenameProtectiongroup", testAccRenameProtectiongroup(c, "testaccpgroupnew"))
	c.Protectiongroups.RenameProtectiongroup("testaccpgroupnew", testAccProtectiongroupName)
	t.Run("DeleteProtectiongroup", testAccDeleteProtectiongroup(c))
	t.Run("EradicateProtectiongroup", testAccEradicateProtectiongroup(c))

	testvols := []string{testvol}
	vollist := map[string][]string{"vollist": testvols}
	t.Run("CreateProtectiongroup_withVolumes", testAccCreateProtectiongroup_withVolumes(c, vollist))
	t.Run("DeleteProtectiongroup", testAccDeleteProtectiongroup(c))
	t.Run("EradicateProtectiongroup", testAccEradicateProtectiongroup(c))

	testhgroups := []string{testhgroup}
	hgrouplist := map[string][]string{"hgrouplist": testhgroups}
	t.Run("CreateProtectiongroup_withHostgroups", testAccCreateProtectiongroup_withHostgroups(c, hgrouplist))
	t.Run("DeleteProtectiongroup", testAccDeleteProtectiongroup(c))
	t.Run("EradicateProtectiongroup", testAccEradicateProtectiongroup(c))

	c.Hosts.DeleteHost(testhost1)
	c.Volumes.DeleteVolume(testvol)
	c.Volumes.EradicateVolume(testvol)
	c.Hostgroups.DeleteHostgroup(testpgroup)
}

func testAccCreateProtectiongroup_basic(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Protectiongroups.CreateProtectiongroup(testAccProtectiongroupName, nil)
		if err != nil {
			t.Fatalf("error creating protection group %s: %s", testAccProtectiongroupName, err)
		}

		if h.Name != testAccProtectiongroupName {
			t.Fatalf("expected: %s; got %s", testAccProtectiongroupName, h.Name)
		}
	}
}

func testAccGetProtectiongroup(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Protectiongroups.GetProtectiongroup(testAccProtectiongroupName, nil)
		if err != nil {
			t.Fatalf("error getting protection group %s: %s", testAccProtectiongroupName, err)
		}

		if h.Name != testAccProtectiongroupName {
			t.Fatalf("expected: %s; got %s", testAccProtectiongroupName, h.Name)
		}
	}
}

func testAccGetProtectiongroup_withParams(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		params := map[string]string{"space": "true"}
		h, err := c.Protectiongroups.GetProtectiongroup(testAccProtectiongroupName, params)
		if err != nil {
			t.Fatalf("error getting protection group %s: %s", testAccProtectiongroupName, err)
		}

		if h.Name != testAccProtectiongroupName {
			t.Fatalf("expected: %s; got %s", testAccProtectiongroupName, h.Name)
		}
	}
}

func testAccEnablePgroupReplication(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		p, err := c.Protectiongroups.EnablePgroupReplication(testAccProtectiongroupName)
		if err != nil {
			t.Fatalf("error enabling replication: %s", err)
		}

		if !p.ReplicateEnabled {
			t.Fatalf("replication not enabled.")
		}
	}
}

func testAccDisablePgroupReplication(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		p, err := c.Protectiongroups.DisablePgroupReplication(testAccProtectiongroupName)
		if err != nil {
			t.Fatalf("error disabling replication: %s", err)
		}

		if p.ReplicateEnabled {
			t.Fatalf("replication enabled.")
		}
	}
}

func testAccEnablePgroupSnapshots(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		p, err := c.Protectiongroups.EnablePgroupSnapshots(testAccProtectiongroupName)
		if err != nil {
			t.Fatalf("error enabling snapshots: %s", err)
		}

		if !p.SnapEnabled {
			t.Fatalf("snapshots not enabled.")
		}
	}
}

func testAccDisablePgroupSnapshots(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		p, err := c.Protectiongroups.DisablePgroupSnapshots(testAccProtectiongroupName)
		if err != nil {
			t.Fatalf("error disabling snapshots: %s", err)
		}

		if p.SnapEnabled {
			t.Fatalf("snapshots enabled.")
		}
	}
}

func testAccCreateProtectiongroup_withHosts(c *Client, hostlist map[string][]string) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Protectiongroups.CreateProtectiongroup(testAccProtectiongroupName, hostlist)
		if err != nil {
			t.Fatalf("error creating protection group %s with hosts: %s", testAccProtectiongroupName, err)
		}

		if h.Name != testAccProtectiongroupName {
			t.Fatalf("expected: %s; got %s", testAccProtectiongroupName, h.Name)
		}
	}
}

func testAccCreateProtectiongroup_withVolumes(c *Client, vollist map[string][]string) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Protectiongroups.CreateProtectiongroup(testAccProtectiongroupName, vollist)
		if err != nil {
			t.Fatalf("error creating protection group %s with hosts: %s", testAccProtectiongroupName, err)
		}

		if h.Name != testAccProtectiongroupName {
			t.Fatalf("expected: %s; got %s", testAccProtectiongroupName, h.Name)
		}
	}
}

func testAccCreateProtectiongroup_withHostgroups(c *Client, hgrouplist map[string][]string) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Protectiongroups.CreateProtectiongroup(testAccProtectiongroupName, hgrouplist)
		if err != nil {
			t.Fatalf("error creating protection group %s with hosts: %s", testAccProtectiongroupName, err)
		}

		if h.Name != testAccProtectiongroupName {
			t.Fatalf("expected: %s; got %s", testAccProtectiongroupName, h.Name)
		}
	}
}

func testAccListProtectiongroups(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Protectiongroups.ListProtectiongroups(nil)
		if err != nil {
			t.Fatalf("error listing protection groups: %s", err)
		}
	}
}

func testAccListProtectiongroups_withParams(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		params := map[string]string{"space": "true"}
		_, err := c.Protectiongroups.ListProtectiongroups(params)
		if err != nil {
			t.Fatalf("error listing protection groups: %s", err)
		}
	}
}

func testAccRenameProtectiongroup(c *Client, name string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Protectiongroups.RenameProtectiongroup(testAccProtectiongroupName, name)
		if err != nil {
			t.Fatalf("error renaming protection group %s to %s: %s", testAccProtectiongroupName, name, err)
		}
	}
}

func testAccRemoveHostsFromProtectiongroup(c *Client, hostlist map[string][]string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Protectiongroups.SetProtectiongroup(testAccProtectiongroupName, hostlist)
		if err != nil {
			t.Fatalf("error removing hosts from protection group %s: %s", testAccProtectiongroupName, err)
		}

	}
}

func testAccDeleteProtectiongroup(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Protectiongroups.DestroyProtectiongroup(testAccProtectiongroupName)
		if err != nil {
			t.Fatalf("error deleting protection group: %s", err)
		}
	}
}

func testAccRecoverProtectiongroup(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Protectiongroups.RecoverProtectiongroup(testAccProtectiongroupName)
		if err != nil {
			t.Fatalf("error recovering protection group: %s", err)
		}
	}
}

func testAccEradicateProtectiongroup(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Protectiongroups.EradicateProtectiongroup(testAccProtectiongroupName)
		if err != nil {
			t.Fatalf("error eradicating protection group: %s", err)
		}
	}
}
