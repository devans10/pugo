// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

type VolumeService struct {
	client *Client
}

// function SetVolume is a helper function that sets the parameter passed in the data interface
// of the volume in the name argument.
// A Volume object is returned with the new values.
func (v *VolumeService) SetVolume(name string, params map[string]string, data interface{}) (*Volume, error) {

	path := fmt.Sprintf("volume/%s", name)
        req, err := v.client.NewRequest("PUT", path, params, data)
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
func (v *VolumeService) CreateSnapshot(volume string, params map[string]string) (*Volume, error) {
	volumes := []string{volume}
	m, err := v.CreateSnapshots(volumes, nil)
	if err != nil {
		return nil, err
	}

	return &m[0], err
}

// CreateSnapshosts function will create a snapshot of all the volumes passed in the volumes slice.
// an array of volume objects is returned.
func (v *VolumeService) CreateSnapshots(volumes []string, params map[string]string) ([]Volume, error) {
	type data struct {
		Snap	bool		`json:"snap"`
		Source	[]string	`json:"source"`
	}

	d := &data{}
	d.Snap = true
	d.Source = volumes
        req, err := v.client.NewRequest("POST", "volume", params, d)
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

// CreateVolume function will create a volume of the given size.  The size is a string with a number and
// a suffix representing the size.
// Accepted Suffixes
//        ====== ======== ======
//        Suffix Size     Bytes
//        ====== ======== ======
//        S      Sector   (2^9)
//        K      Kilobyte (2^10)
//        M      Megabyte (2^20)
//        G      Gigabyte (2^30)
//        T      Terabyte (2^40)
//        P      Petabyte (2^50)
//        ====== ======== ======
func (v *VolumeService) CreateVolume(name string, size string, params map[string]string) (*Volume, error) {

	path := fmt.Sprintf("volume/%s", name)
	data := map[string]string{"size": size}
        req, err := v.client.NewRequest("POST", path, params, data)
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

// Create a conglomerate volume.
// This is not a typical volume thus there is no size.  It's main purpose to connect to a
// host/hgroup to create a PE LUN.  Once the conglomerate volume is connected to a
// host/hgroup, it is used as a protocol-endpoint to connect a vvol to a host/hgroup to
// allow traffic.
func (v *VolumeService) CreateConglomerateVolume(name string, params map[string]string) (*Volume, error) {

        path := fmt.Sprintf("volume/%s", name)
        data := map[string]bool{"protocol_endpoint": true}
        req, err := v.client.NewRequest("POST", path, params, data)
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

// Clone a volume and return a dictionary describing the new volume.
func (v *VolumeService) CopyVolume(dest string, source string, params map[string]string) (*Volume, error) {

        path := fmt.Sprintf("volume/%s", dest)
        data := map[string]string{"source": source}
        req, err := v.client.NewRequest("POST", path, params, data)
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

// Delete an existing volume or snapshot
func (v *VolumeService) DeleteVolume(name string, params map[string]string) (*Volume, error) {

        path := fmt.Sprintf("volume/%s", name)
        req, err := v.client.NewRequest("DELETE", path, params, nil)
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

// Eradicate a deleted volume or snapshot
func (v *VolumeService) EradicateVolume(name string, params map[string]string) (*Volume, error) {

        path := fmt.Sprintf("volume/%s", name)
        data := map[string]bool{"eradicate": true}
        req, err := v.client.NewRequest("POST", path, params, data)
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

// Extend the size of the volume
func (v *VolumeService) ExtendVolume(name string, size string, params map[string]string) (*Volume, error) {

	type data struct {
                Truncate	bool	`json:"truncate"`
                Size		string	`json:"size"`
        }
        d := &data{}
        d.Size = size
        d.Truncate = false
	m, err := v.SetVolume(name, params, d)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Get volume attributes
func (v *VolumeService) GetVolume(name string, params map[string]string) (*Volume, error) {

	path := fmt.Sprintf("volume/%s", name)
        req, err := v.client.NewRequest("GET", path, params, nil)
	m := &Volume{}
        _, err = v.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Add a volume to a protection group
func (v *VolumeService) AddVolume(volume string, pgroup string, params map[string]string) (*VolumePgroup, error) {

	path := fmt.Sprintf("volume/%s/pgroup/%s", volume, pgroup)
        req, err := v.client.NewRequest("POST", path, params, nil)
        m := &VolumePgroup{}
        _, err = v.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Remove a volume from a protection group
func (v *VolumeService) RemoveVolume(volume string, pgroup string, params map[string]string) (*VolumePgroup, error) {

        path := fmt.Sprintf("volume/%s/pgroup/%s", volume, pgroup)
        req, err := v.client.NewRequest("DELETE", path, params, nil)
        m := &VolumePgroup{}
        _, err = v.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// List Volume Block Differences
// not implemented yet
func (v *VolumeService) ListVolumeBlockDiff() error {
	return nil
}

// List Volume Private Connections
// not implemented yet
func (v *VolumeService) ListVolumePrivateConnections() error {
        return nil
}

// List volume shared connections.
func (v *VolumeService) ListVolumeSharedConnections() error {
        return nil
}

// List volumes
func (v *VolumeService) ListVolumes(params map[string]string) ([]Volume, error) {

        req, err := v.client.NewRequest("GET", "volume", params, nil)
        m := []Volume{}
        _, err = v.client.Do(req, &m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Rename a volume
func (v *VolumeService) RenameVolume(volume string, name string, params map[string]string) (*Volume, error) {

        data := map[string]string{"name": name}
        m, err := v.SetVolume(volume, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Recover a deleted volume
func (v *VolumeService) RecoverVolume(volume string, params map[string]string) (*Volume, error) {

        data := map[string]string{"action": "recover"}
        m, err := v.SetVolume(volume, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Decrese the size of a volume
// WARNING!!
// Potential data loss
func (v *VolumeService) TruncateVolume(name string, size string, params map[string]string) (*Volume, error) {

	type data struct {
                Truncate        bool    `json:"truncate"`
                Size            string  `json:"size"`
        }
        d := &data{}
        d.Size = size
        d.Truncate = true
        m, err := v.SetVolume(name, params, d)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Move a volume to a different container
func (v *VolumeService) MoveVolume(name string, container string, params map[string]string) (*Volume, error) {

        data := map[string]string{"container": container}
        m, err := v.SetVolume(name, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}
