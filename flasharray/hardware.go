// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

type HardwareService struct {
	client *Client
}

// Get Drive attributes
func (n *HardwareService) GetDrive(name string) (*Drive, error) {

	path := fmt.Sprintf("drive/%s", name)
	req, err := n.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Drive{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// List Drives
func (n *HardwareService) ListDrives() ([]Drive, error) {

	req, err := n.client.NewRequest("GET", "drive", nil, nil)
	if err != nil {
		return nil, err
	}

	m := []Drive{}
	_, err = n.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Get Hardware attributes
func (n *HardwareService) GetHardware(name string) (*Component, error) {

	path := fmt.Sprintf("hardware/%s", name)
	req, err := n.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Component{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// List Hardware
func (n *HardwareService) ListHardware() ([]Component, error) {

	req, err := n.client.NewRequest("GET", "hardware", nil, nil)
	if err != nil {
		return nil, err
	}

	m := []Component{}
	_, err = n.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Set hardware attributes
func (n *HardwareService) SetHardware(name string, params map[string]string, data interface{}) (*Component, error) {

	path := fmt.Sprintf("hardware/%s", name)
	req, err := n.client.NewRequest("PUT", path, params, data)
	if err != nil {
		return nil, err
	}

	m := &Component{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
