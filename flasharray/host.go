// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

type HostService struct {
	client *Client
}

// Connect a volume to a host
func (h *HostService) ConnectHost(host string, volume string, params map[string]string) (*ConnectedVolume, error) {

	path := fmt.Sprintf("host/%s/volume/%s", host, volume)
	req, err := h.client.NewRequest("POST", path, params, nil)
	m := &ConnectedVolume{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Create a new host
func (h *HostService) CreateHost(name string, params map[string]string) (*Host, error) {

	path := fmt.Sprintf("host/%s", name)
	req, err := h.client.NewRequest("POST", path, params, nil)
	if err != nil {
		return nil, err
	}

	m := &Host{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Delete a host
func (h *HostService) DeleteHost(name string, params map[string]string) (*Host, error) {

	path := fmt.Sprintf("host/%s", name)
	req, err := h.client.NewRequest("DELETE", path, params, nil)
	if err != nil {
		return nil, err
	}

	m := &Host{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Disconnect a volume from a host
func (h *HostService) DisconnectHost(host string, volume string, params map[string]string) (*ConnectedVolume, error) {

	path := fmt.Sprintf("host/%s/volume/%s", host, volume)
	req, err := h.client.NewRequest("DELETE", path, params, nil)
	m := &ConnectedVolume{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Return the attributes of the given host
func (h *HostService) GetHost(name string, params map[string]string) (*Host, error) {

	path := fmt.Sprintf("host/%s", name)
	req, err := h.client.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	m := &Host{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Add a host to a protection group
func (h *HostService) AddHost(host string, pgroup string, params map[string]string) (*HostPgroup, error) {

	path := fmt.Sprintf("host/%s/pgroup/%s", host, pgroup)
	req, err := h.client.NewRequest("POST", path, params, nil)
	m := &HostPgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Remove a host from a protection group
func (h *HostService) RemoveHost(host string, pgroup string, params map[string]string) (*HostPgroup, error) {

	path := fmt.Sprintf("host/%s/pgroup/%s", host, pgroup)
	req, err := h.client.NewRequest("DELETE", path, params, nil)
	m := &HostPgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// The the host connections
func (h *HostService) ListHostConnections(host string, params map[string]string) ([]ConnectedVolume, error) {

	path := fmt.Sprintf("host/%s/volume", host)
	req, err := h.client.NewRequest("GET", path, params, nil)
	m := []ConnectedVolume{}
	_, err = h.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// List hosts
func (h *HostService) ListHosts(host string, params map[string]string) ([]Host, error) {

	path := fmt.Sprintf("host", host)
	req, err := h.client.NewRequest("GET", path, params, nil)
	m := []Host{}
	_, err = h.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Rename a host
func (h *HostService) RenameHost(host string, name string, params map[string]string) (*Host, error) {

	data := map[string]string{"name": name}
	m, err := h.SetHost(host, params, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Set the attribute of a host
func (h *HostService) SetHost(name string, params map[string]string, data interface{}) (*Host, error) {

	path := fmt.Sprintf("host/%s", name)
	req, err := h.client.NewRequest("PUT", path, params, data)
	if err != nil {
		return nil, err
	}

	m := &Host{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
