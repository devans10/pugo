// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

// VolumeService struct for volume API endpoints
type VolumeService struct {
	client *Client
}

// SetVolume is a helper function that sets the parameter passed in the data interface
// of the volume in the name argument.
// A Volume object is returned with the new values.
func (v *VolumeService) SetVolume(name string, data interface{}) (*Volume, error) {

	path := fmt.Sprintf("volume/%s", name)
	req, err := v.client.NewRequest("PUT", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Volume{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// CreateSnapshot function creates a volume snapshot of the volume passed in the argument.
func (v *VolumeService) CreateSnapshot(volume string, suffix string) (*Volume, error) {
	volumes := []string{volume}
	m, err := v.CreateSnapshots(volumes, suffix)
	if err != nil {
		return nil, err
	}

	return &m[0], err
}

// CreateSnapshots function will create a snapshot of all the volumes passed in the volumes slice.
// an array of volume objects is returned.
func (v *VolumeService) CreateSnapshots(volumes []string, suffix string) ([]Volume, error) {

	data := make(map[string]interface{})
	data["snap"] = true
	data["source"] = volumes
	data["suffix"] = suffix
	req, err := v.client.NewRequest("POST", "volume", nil, data)
	if err != nil {
		return nil, err
	}

	m := []Volume{}
	_, err = v.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// CreateVolume function will create a volume of the given size.  The size is an integer in bytes
func (v *VolumeService) CreateVolume(name string, size int) (*Volume, error) {

	path := fmt.Sprintf("volume/%s", name)
	data := map[string]int{"size": size}
	req, err := v.client.NewRequest("POST", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Volume{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// CreateConglomerateVolume creates a conglomerate volume.
// This is not a typical volume thus there is no size.  It's main purpose to connect to a
// host/hgroup to create a PE LUN.  Once the conglomerate volume is connected to a
// host/hgroup, it is used as a protocol-endpoint to connect a vvol to a host/hgroup to
// allow traffic.
func (v *VolumeService) CreateConglomerateVolume(name string) (*Volume, error) {

	path := fmt.Sprintf("volume/%s", name)
	data := map[string]bool{"protocol_endpoint": true}
	req, err := v.client.NewRequest("POST", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Volume{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// CopyVolume clones a volume and return a dictionary describing the new volume.
func (v *VolumeService) CopyVolume(dest string, source string, overwrite bool) (*Volume, error) {

	path := fmt.Sprintf("volume/%s", dest)
	data := map[string]interface{}{"source": source, "overwrite": overwrite}
	req, err := v.client.NewRequest("POST", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Volume{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DeleteVolume deletes an existing volume or snapshot
func (v *VolumeService) DeleteVolume(name string) (*Volume, error) {

	path := fmt.Sprintf("volume/%s", name)
	req, err := v.client.NewRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Volume{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// EradicateVolume eradicates a deleted volume or snapshot
func (v *VolumeService) EradicateVolume(name string) (*Volume, error) {

	path := fmt.Sprintf("volume/%s", name)
	data := map[string]bool{"eradicate": true}
	req, err := v.client.NewRequest("DELETE", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Volume{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ExtendVolume increases the size of the volume
func (v *VolumeService) ExtendVolume(name string, size int) (*Volume, error) {

	data := make(map[string]interface{})
	data["size"] = size
	data["truncate"] = false
	m, err := v.SetVolume(name, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetVolume lists attributes of the specified volume
func (v *VolumeService) GetVolume(name string, params map[string]string) (*Volume, error) {

	path := fmt.Sprintf("volume/%s", name)
	if params["action"] == "monitor" {
		m, err := v.MonitorVolume(name, params)
		if err != nil {
			return nil, err
		}
		if len(m) > 0 {
			vol := m[0]
			return &vol, nil
		}
		return nil, nil
	}
	req, err := v.client.NewRequest("GET", path, params, nil)
	m := &Volume{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// MonitorVolume returns metrics of the specified volume.  This is similar to GetVolume when
// params = {"action": "monitor"}, but that endpoint returns an array instead of a single
// Volume struct.
func (v *VolumeService) MonitorVolume(name string, params map[string]string) ([]Volume, error) {

	path := fmt.Sprintf("volume/%s", name)
	p := map[string]string{"action": "monitor"}
	req, err := v.client.NewRequest("GET", path, p, nil)
	m := []Volume{}
	_, err = v.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// AddVolume adds a volume to a protection group
func (v *VolumeService) AddVolume(volume string, pgroup string) (*VolumePgroup, error) {

	path := fmt.Sprintf("volume/%s/pgroup/%s", volume, pgroup)
	req, err := v.client.NewRequest("POST", path, nil, nil)
	m := &VolumePgroup{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// RemoveVolume removes a volume from a protection group
func (v *VolumeService) RemoveVolume(volume string, pgroup string) (*VolumePgroup, error) {

	path := fmt.Sprintf("volume/%s/pgroup/%s", volume, pgroup)
	req, err := v.client.NewRequest("DELETE", path, nil, nil)
	m := &VolumePgroup{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ListVolumeBlockDiff lists Volume Block Differences
func (v *VolumeService) ListVolumeBlockDiff(name string, params map[string]string) ([]Block, error) {
	path := fmt.Sprintf("volume/%s/diff", name)
	req, err := v.client.NewRequest("GET", path, params, nil)
	m := []Block{}
	_, err = v.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ListVolumePrivateConnections lists Volume Private Connections
func (v *VolumeService) ListVolumePrivateConnections(name string) ([]Connection, error) {
	path := fmt.Sprintf("volume/%s/host", name)
	req, err := v.client.NewRequest("GET", path, nil, nil)
	m := []Connection{}
	_, err = v.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ListVolumeSharedConnections lists Volume Shared Connections
func (v *VolumeService) ListVolumeSharedConnections(name string) ([]Connection, error) {
	path := fmt.Sprintf("volume/%s/hgroup", name)
	req, err := v.client.NewRequest("GET", path, nil, nil)
	m := []Connection{}
	_, err = v.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ListVolumes lists attributes for volumes
func (v *VolumeService) ListVolumes(params map[string]string) ([]Volume, error) {

	req, err := v.client.NewRequest("GET", "volume", params, nil)
	m := []Volume{}
	_, err = v.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// RenameVolume renames a volume
func (v *VolumeService) RenameVolume(volume string, name string) (*Volume, error) {

	data := map[string]string{"name": name}
	m, err := v.SetVolume(volume, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// RecoverVolume recovers a deleted volume
func (v *VolumeService) RecoverVolume(volume string) (*Volume, error) {

	data := map[string]string{"action": "recover"}
	m, err := v.SetVolume(volume, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// TruncateVolume decreses the size of a volume
// WARNING!!
// Potential data loss
func (v *VolumeService) TruncateVolume(name string, size int) (*Volume, error) {

	data := make(map[string]interface{})
	data["size"] = size
	data["truncate"] = true
	m, err := v.SetVolume(name, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// MoveVolume moves a volume to a different container
func (v *VolumeService) MoveVolume(name string, container string) (*Volume, error) {

	data := map[string]string{"container": container}
	m, err := v.SetVolume(name, data)
	if err != nil {
		return nil, err
	}

	return m, err
}
