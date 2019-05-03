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
	"fmt"
	"io/ioutil"
	"net/http"
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

	c.Protectiongroups.CreateProtectiongroup(testpgroup, nil)

	t.Run("CreateVolume", testAccCreateVolume(c))
	t.Run("GetVolume", testAccGetVolume(c))
	t.Run("GetVolume_withParamSpace", testAccGetVolumeWithParamSpace(c))
	t.Run("GetVolume_withParamAction", testAccGetVolumeWithParamAction(c))
	t.Run("CreateSnapshot", testAccCreateSnapshot(c))
	t.Run("CloneVolume", testAccCloneVolume(c))
	t.Run("AddVolumeToPgroup", testAccAddVolume(c, testpgroup))
	t.Run("RemoveVolumeFromPgroup", testAccRemoveVolume(c, testpgroup))
	t.Run("ExtendVolume", testAccExtendVolume(c))
	t.Run("TruncateVolume", testAccTruncateVolume(c))
	t.Run("ListVolumes", testAccListVolumes(c))
	t.Run("ListVolumes_withParams", testAccListVolumesWithParams(c))
	t.Run("RenameVolume", testAccRenameVolume(c, "testAccVolnew"))
	c.Volumes.RenameVolume("testAccVolnew", testAccVolumeName)
	t.Run("DeleteVolume", testAccDeleteVolume(c))
	t.Run("RecoverVolume", testAccRecoverVolume(c))
	c.Volumes.DeleteVolume(testAccVolumeName)
	t.Run("EradicateVolume", testAccEradicateVolume(c))

	c.Volumes.DeleteVolume(testvolclone)
	c.Volumes.DeleteVolume(fmt.Sprintf("%s.%s", testAccVolumeName, testvolsnapshot))
	c.Volumes.EradicateVolume(testvolclone)
	c.Volumes.EradicateVolume(fmt.Sprintf("%s.%s", testAccVolumeName, testvolsnapshot))
	c.Protectiongroups.DestroyProtectiongroup(testpgroup)
	c.Protectiongroups.EradicateProtectiongroup(testpgroup)
}

func testAccCreateVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.CreateVolume(testAccVolumeName, testvolsize)
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

func testAccGetVolumeWithParamSpace(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		params := map[string]string{"space": "true"}
		h, err := c.Volumes.GetVolume(testAccVolumeName, params)
		if err != nil {
			t.Fatalf("error getting volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != testAccVolumeName {
			t.Fatalf("expected: %s; got %s", testAccVolumeName, h.Name)
		}
		if h.Size != testvolsize {
			t.Fatalf("expected: %d; got %d", testvolsize, h.Size)
		}
	}
}

func testAccGetVolumeWithParamAction(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.GetVolume(testAccVolumeName, map[string]string{"action": "monitor"})
		if err != nil {
			t.Fatalf("error getting volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != testAccVolumeName {
			t.Fatalf("expected: %s; got %s", testAccVolumeName, h.Name)
		}
		if h.Time == "" {
			t.Fatalf("time property did not exist; got: %+v", h)
		}
	}
}

func testAccCreateSnapshot(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.CreateSnapshot(testAccVolumeName, testvolsnapshot)
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
		h, err := c.Volumes.CopyVolume(testvolclone, testAccVolumeName, false)
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
		_, err := c.Volumes.AddVolume(testAccVolumeName, pgroup)
		if err != nil {
			t.Fatalf("error adding volume %s to pgroup %s: %s", testAccVolumeName, pgroup, err)
		}
	}
}

func testAccRemoveVolume(c *Client, pgroup string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.RemoveVolume(testAccVolumeName, pgroup)
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

func testAccListVolumesWithParams(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		params := map[string]string{"space": "true"}
		_, err := c.Volumes.ListVolumes(params)
		if err != nil {
			t.Fatalf("error listing volumes: %s", err)
		}
	}
}

func testAccExtendVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.ExtendVolume(testAccVolumeName, testvolresize)
		if err != nil {
			t.Fatalf("error extending volume %s to %d: %s", testAccVolumeName, testvolresize, err)
		}
	}
}

func testAccTruncateVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.TruncateVolume(testAccVolumeName, testvolsize)
		if err != nil {
			t.Fatalf("error truncating volume %s to %d: %s", testAccVolumeName, testvolsize, err)
		}
	}
}

func testAccRenameVolume(c *Client, name string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.RenameVolume(testAccVolumeName, name)
		if err != nil {
			t.Fatalf("error renaming volume %s to %s: %s", testAccVolumeName, name, err)
		}
	}
}

func testAccDeleteVolume(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.DeleteVolume(testAccVolumeName)
		if err != nil {
			t.Fatalf("error deleting volume: %s", err)
		}
	}
}

func testAccRecoverVolume(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.RecoverVolume(testAccVolumeName)
		if err != nil {
			t.Fatalf("error recovering volume %s: %s", testAccVolumeName, err)
		}
	}
}

func testAccEradicateVolume(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.EradicateVolume(testAccVolumeName)
		if err != nil {
			t.Fatalf("error eradicating volume: %s", err)
		}
	}
}

func TestSetVolume(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v5_renamed"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v5", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutVolumevolName(restVersion))),
			Header:     head,
		}
	})

	data := map[string]string{"name": "v5_renamed"}
	vol, err := c.Volumes.SetVolume("v5", data)
	ok(t, err)
	equals(t, &testVolume, vol)
}

func TestCreateSnapshot(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v5.1",
		Source:  "v5",
		Serial:  "B75F8356604B431D00011020",
		Size:    5368709120,
		Created: "2017-12-16T05:12:53Z",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostVolume(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.CreateSnapshot("v5", "1")
	ok(t, err)
	equals(t, &testVolume, vol)
}

func TestCreateSnapshots(t *testing.T) {
	restVersion := "1.15"
	testVolume := []Volume{Volume{Name: "v5.1",
		Source:  "v5",
		Serial:  "B75F8356604B431D00011020",
		Size:    5368709120,
		Created: "2017-12-16T05:12:53Z",
	},
		Volume{Name: "v6.4129",
			Source:  "v6",
			Serial:  "B75F8356604B431D00011021",
			Size:    1073741824,
			Created: "2017-12-16T05:12:53Z",
		},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostVolume(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.CreateSnapshots([]string{"v5", "v6"}, "")
	ok(t, err)
	equals(t, testVolume, vol)
}

func TestCreateVolume(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v5",
		Source:  "",
		Serial:  "B75F8356604B431D0001101D",
		Size:    5368709120,
		Created: "2017-12-16T05:12:52Z",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v5", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostVolumevol(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.CreateVolume("v5", 5368709120)
	ok(t, err)
	equals(t, &testVolume, vol)
}

//TODO
func TestCreateConglomerateVolume(t *testing.T) {}

func TestCopyVolume(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v5",
		Source:  "",
		Serial:  "B75F8356604B431D0001101D",
		Size:    5368709120,
		Created: "2017-12-16T05:12:52Z",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v5", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostVolumevol(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.CopyVolume("v5", "v6", false)
	ok(t, err)
	equals(t, &testVolume, vol)
}

func TestDeleteVolume(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v5"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v5", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteVolumevol(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.DeleteVolume("v5")
	ok(t, err)
	equals(t, &testVolume, vol)
}

func TestEradicateVolume(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v5"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v5", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteVolumevol(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.EradicateVolume("v5")
	ok(t, err)
	equals(t, &testVolume, vol)
}

func TestExtendVolume(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v5", Size: 1073741824}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v5", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutVolumevolSize(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.ExtendVolume("v5", 1073741824)
	ok(t, err)
	equals(t, &testVolume, vol)
}

func TestGetVolume(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v1",
		Created: "2017-12-16T05:12:38Z",
		Serial:  "B75F8356604B431D00011010",
		Size:    1073741824,
		Source:  "",
	}

	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v1", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetVolumevol(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.GetVolume("v1", nil)
	ok(t, err)
	equals(t, &testVolume, vol)
}

//TODO: Figure out how to compare this.
func TestMonitorVolume(t *testing.T) {
	restVersion := "1.15"
	/*testVolume := Volume{Name: "v3",
		InputPerSec:       &int{0},
		OutputPerSec:      &int{0},
		ReadsPerSec:       &int{0},
		SanUsecPerReadOp:  &int{0},
		SanUsecPerWriteOp: &int{0},
		Time:              "2017-12-16T05:12:51Z",
		UsecPerReadOp:     &int{0},
		UsecPerWriteOp:    &int{0},
		WritesPerSec:      &int{0},
	}*/

	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v3?action=monitor", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetVolumevolMonitor(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Volumes.MonitorVolume("v3", nil)
	ok(t, err)
	//equals(t, testVolume, vol)
}

//TODO: Get Output Sample
func TestAddVolume(t *testing.T) {
	restVersion := "1.15"
	//testVolume := Volume
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v3/pgroup/pg1", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostVolumevol(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Volumes.AddVolume("v3", "pg1")
	ok(t, err)
	//equals(t, testVolume, vol)
}

//TODO: Get Output Sample
func TestRemoveVolume(t *testing.T) {
	restVersion := "1.15"
	//testVolume := Volume
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v3/pgroup/pg1", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostVolumevol(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Volumes.RemoveVolume("v3", "pg1")
	ok(t, err)
	//equals(t, testVolume, vol)
}

func TestListVolumeBlockDiff(t *testing.T) {
	restVersion := "1.15"
	testBlock := []Block{Block{Length: 20480, Offset: 0}, Block{Length: 2139095040, Offset: 8388608}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v3/diff", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetVolumevolDiff(restVersion))),
			Header:     head,
		}
	})

	block, err := c.Volumes.ListVolumeBlockDiff("v3", nil)
	ok(t, err)
	equals(t, testBlock, block)
}

func TestListVolumePrivateConnections(t *testing.T) {
	restVersion := "1.15"
	testConnection := []Connection{Connection{
		Host: "h2",
		Lun:  7,
		Name: "v3",
		Size: 3221225472,
	}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v3/host", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetVolumevolHost(restVersion))),
			Header:     head,
		}
	})

	conn, err := c.Volumes.ListVolumePrivateConnections("v3")
	ok(t, err)
	equals(t, testConnection, conn)
}

func TestListVolumeSharedConnections(t *testing.T) {
	restVersion := "1.15"
	testConnection := []Connection{Connection{
		Hgroup: "hg2",
		Host:   "",
		Lun:    15,
		Name:   "v2",
		Size:   2147483648,
	}}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v2/hgroup", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetVolumevolHgroup(restVersion))),
			Header:     head,
		}
	})

	conn, err := c.Volumes.ListVolumeSharedConnections("v2")
	ok(t, err)
	equals(t, testConnection, conn)
}

func TestListVolumes(t *testing.T) {
	restVersion := "1.15"
	testVolume := []Volume{Volume{Name: "v1",
		Source:  "",
		Serial:  "B75F8356604B431D00011010",
		Size:    1073741824,
		Created: "2017-12-16T05:12:38Z",
	},
		Volume{Name: "v2",
			Source:  "",
			Serial:  "B75F8356604B431D00011011",
			Size:    2147483648,
			Created: "2017-12-16T05:12:38Z",
		},
		Volume{Name: "v3",
			Source:  "",
			Serial:  "B75F8356604B431D00011012",
			Size:    3221225472,
			Created: "2017-12-16T05:12:39Z",
		},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetVolume(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.ListVolumes(nil)
	ok(t, err)
	equals(t, testVolume, vol)
}

func TestRenameVolume(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v5_renamed"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v5", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutVolumevolName(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.RenameVolume("v5", "v5_renamed")
	ok(t, err)
	equals(t, &testVolume, vol)
}

func TestRecoverVolume(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v5_renamed"}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v5_renamed?action=recover", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutVolumevolName(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.RecoverVolume("v5_renamed")
	ok(t, err)
	equals(t, &testVolume, vol)
}

func TestTruncateVolume(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v5", Size: 1073741824}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v5", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutVolumevolSize(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.TruncateVolume("v5", 1073741824)
	ok(t, err)
	equals(t, &testVolume, vol)
}

//TODO: Get Sample Output
func TestMoveVolume(t *testing.T) {
	restVersion := "1.15"
	testVolume := Volume{Name: "v5", Size: 1073741824}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/volume/v5", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutVolumevolSize(restVersion))),
			Header:     head,
		}
	})

	vol, err := c.Volumes.MoveVolume("v5", "c1")
	ok(t, err)
	equals(t, &testVolume, vol)
}
