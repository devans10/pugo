// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
	"testing"
)

const testAccVolumeName = "testAccvolume"
const testvolsnapshot = "testaccvolsnapshot"
const testvolclone = "testaccvolcone"
const testvolsize = 1024000000
const testvolresize = 2048000000
const testpgroup = "testacchostpgroup"

func TestAccVolumes(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	c.Protectiongroups.CreateProtectiongroup(testpgroup, nil, nil)

	t.Run("CreateVolume", testAccCreateVolume(c))
	t.Run("GetVolume", testAccGetVolume(c))
	t.Run("CreateSnapshot", testAccCreateSnapshot(c))
	t.Run("CloneVolume", testAccCloneVolume(c))
	t.Run("AddVolumeToPgroup", testAccAddVolume(c, testpgroup))
	t.Run("RemoveVolumeFromPgroup", testAccRemoveVolume(c, testpgroup))
	t.Run("ExtendVolume", testAccExtendVolume(c))
	t.Run("TruncateVolume", testAccTruncateVolume(c))
	t.Run("ListVolumes", testAccListVolumes(c))
	t.Run("RenameVolume", testAccRenameVolume(c, "testAccVolnew"))
	c.Volumes.RenameVolume("testAccVolnew", testAccVolumeName, nil)
	t.Run("DeleteVolume", testAccDeleteVolume(c))
	t.Run("RecoverVolume", testAccRecoverVolume(c))
	c.Volumes.DeleteVolume(testAccVolumeName, nil)
	t.Run("EradicateVolume", testAccEradicateVolume(c))

	c.Volumes.DeleteVolume(testvolclone, nil)
	c.Volumes.DeleteVolume(fmt.Sprintf("%s.%s", testAccVolumeName, testvolsnapshot), nil)
	c.Volumes.EradicateVolume(testvolclone, nil)
	c.Volumes.EradicateVolume(fmt.Sprintf("%s.%s", testAccVolumeName, testvolsnapshot), nil)
	c.Protectiongroups.DestroyProtectiongroup(testpgroup, nil)
	c.Protectiongroups.EradicateProtectiongroup(testpgroup, nil)
}

func testAccCreateVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.CreateVolume(testAccVolumeName, testvolsize, nil)
		if err != nil {
			t.Fatalf("error creating volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != testAccVolumeName {
			t.Fatalf("expected: %s; got %s", testAccVolumeName, h.Name)
		}
	}
}

func testAccGetVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.GetVolume(testAccVolumeName, nil)
		if err != nil {
			t.Fatalf("error getting volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != testAccVolumeName {
			t.Fatalf("expected: %s; got %s", testAccVolumeName, h.Name)
		}
	}
}

func testAccCreateSnapshot(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.CreateSnapshot(testAccVolumeName, testvolsnapshot, nil)
		if err != nil {
			t.Fatalf("error snapshotting volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != fmt.Sprintf("%s.%s", testAccVolumeName, testvolsnapshot) {
			t.Fatalf("expected: %s; got %s", fmt.Sprintf("%s.%s", testAccVolumeName, testvolsnapshot), h.Name)
		}
	}
}

func testAccCloneVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.CopyVolume(testvolclone, testAccVolumeName, false, nil)
		if err != nil {
			t.Fatalf("error cloning volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != testvolclone {
			t.Fatalf("expected: %s; got %s", testvolclone, h.Name)
		}
		if h.Source != testAccVolumeName {
			t.Fatalf("expected: %s; got %s", testAccVolumeName, h.Source)
		}
	}
}

func testAccAddVolume(c *Client, pgroup string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.AddVolume(testAccVolumeName, pgroup, nil)
		if err != nil {
			t.Fatalf("error adding volume %s to pgroup %s: %s", testAccVolumeName, pgroup, err)
		}
	}
}

func testAccRemoveVolume(c *Client, pgroup string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.RemoveVolume(testAccVolumeName, pgroup, nil)
		if err != nil {
			t.Fatalf("error removing volume %s from pgroup %s: %s", testAccVolumeName, pgroup, err)
		}
	}
}

func testAccListVolumes(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.ListVolumes(nil)
		if err != nil {
			t.Fatalf("error listing volumes: %s", err)
		}
	}
}

func testAccExtendVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.ExtendVolume(testAccVolumeName, testvolresize, nil)
		if err != nil {
			t.Fatalf("error extending volume %s to %d: %s", testAccVolumeName, testvolresize, err)
		}
	}
}

func testAccTruncateVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.TruncateVolume(testAccVolumeName, testvolsize, nil)
		if err != nil {
			t.Fatalf("error truncating volume %s to %d: %s", testAccVolumeName, testvolsize, err)
		}
	}
}

func testAccRenameVolume(c *Client, name string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.RenameVolume(testAccVolumeName, name, nil)
		if err != nil {
			t.Fatalf("error renaming volume %s to %s: %s", testAccVolumeName, name, err)
		}
	}
}

func testAccDeleteVolume(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.DeleteVolume(testAccVolumeName, nil)
		if err != nil {
			t.Fatalf("error deleting volume: %s", err)
		}
	}
}

func testAccRecoverVolume(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.RecoverVolume(testAccVolumeName, nil)
		if err != nil {
			t.Fatalf("error recovering volume %s: %s", testAccVolumeName, err)
		}
	}
}

func testAccEradicateVolume(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.EradicateVolume(testAccVolumeName, nil)
		if err != nil {
			t.Fatalf("error eradicating volume: %s", err)
		}
	}
}
