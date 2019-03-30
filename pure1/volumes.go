// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// VolumeService type creates a service to expose volume endpoints
type VolumeService struct {
	client *Client
}

// GetVolumes returns a list of volume objects
func (v *VolumeService) GetVolumes(params map[string]string) ([]Volume, error) {
	req, err := v.client.NewRequest("GET", "volumes", params, nil)
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
