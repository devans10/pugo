// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// VolumeSnapshotService type creates a service to expose VolumeSnapshot endpoints
type VolumeSnapshotService struct {
	client *Client
}

// GetVolumeSnapshots returns a list of VolumeSnapshot objects
func (v *VolumeSnapshotService) GetVolumeSnapshots(params map[string]string) ([]VolumeSnapshot, error) {
	req, err := v.client.NewRequest("GET", "volume-snapshots", params, nil)
	if err != nil {
		return nil, err
	}

	m := []VolumeSnapshot{}
	_, err = v.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
