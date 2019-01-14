// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

// HostgroupService struct for hgroup API endpoints
type HostgroupService struct {
	client *Client
}

// ConnectHostgroup connects a Volume to a hostgroup
func (h *HostgroupService) ConnectHostgroup(hgroup string, volume string, data interface{}) (*ConnectedVolume, error) {

	path := fmt.Sprintf("hgroup/%s/volume/%s", hgroup, volume)
	req, err := h.client.NewRequest("POST", path, nil, data)
	m := &ConnectedVolume{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// CreateHostgroup creates a new hostgroup
func (h *HostgroupService) CreateHostgroup(name string, data interface{}) (*Hostgroup, error) {

	path := fmt.Sprintf("hgroup/%s", name)
	req, err := h.client.NewRequest("POST", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Hostgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DeleteHostgroup deletes a hostgroup
func (h *HostgroupService) DeleteHostgroup(name string) (*Hostgroup, error) {

	path := fmt.Sprintf("hgroup/%s", name)
	req, err := h.client.NewRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Hostgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DisconnectHostgroup disconnects a volume from a hostgroup
func (h *HostgroupService) DisconnectHostgroup(hgroup string, volume string) (*ConnectedVolume, error) {

	path := fmt.Sprintf("hgroup/%s/volume/%s", hgroup, volume)
	req, err := h.client.NewRequest("DELETE", path, nil, nil)
	m := &ConnectedVolume{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetHostgroup returns a map of the hostgroup attributes
// see API reference on array for list of valid parameters
func (h *HostgroupService) GetHostgroup(name string, params map[string]string) (*Hostgroup, error) {

	path := fmt.Sprintf("hgroup/%s", name)
	req, err := h.client.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	m := &Hostgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// AddHostgroup adds a hostgroup to a Protection Group
func (h *HostgroupService) AddHostgroup(hgroup string, pgroup string) (*HostgroupPgroup, error) {

	path := fmt.Sprintf("hgroup/%s/pgroup/%s", hgroup, pgroup)
	req, err := h.client.NewRequest("POST", path, nil, nil)
	m := &HostgroupPgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// RemoveHostgroup removes a hostgroup from a protection group.
func (h *HostgroupService) RemoveHostgroup(hgroup string, pgroup string) (*HostgroupPgroup, error) {

	path := fmt.Sprintf("hgroup/%s/pgroup/%s", hgroup, pgroup)
	req, err := h.client.NewRequest("DELETE", path, nil, nil)
	m := &HostgroupPgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ListHostgroupConnections lists the hostgroup volume connections
func (h *HostgroupService) ListHostgroupConnections(hgroup string) ([]HostgroupConnection, error) {

	path := fmt.Sprintf("hgroup/%s/volume", hgroup)
	req, err := h.client.NewRequest("GET", path, nil, nil)
	m := []HostgroupConnection{}
	_, err = h.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ListHostgroups lists hostgroups
func (h *HostgroupService) ListHostgroups(params map[string]string) ([]Hostgroup, error) {

	req, err := h.client.NewRequest("GET", "hgroup", params, nil)
	m := []Hostgroup{}
	_, err = h.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// RenameHostgroup renames a hostgroup
func (h *HostgroupService) RenameHostgroup(hgroup string, name string) (*Hostgroup, error) {

	data := map[string]string{"name": name}
	m, err := h.SetHostgroup(hgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// SetHostgroup modifies the specified hostgroup's attributes
func (h *HostgroupService) SetHostgroup(name string, data interface{}) (*Hostgroup, error) {

	path := fmt.Sprintf("hgroup/%s", name)
	req, err := h.client.NewRequest("PUT", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Hostgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
