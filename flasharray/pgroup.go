// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

type ProtectiongroupService struct {
	client *Client
}

// Create a Protection group
func (h *ProtectiongroupService) CreateProtectiongroup(name string, data interface{}, params map[string]string) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, err := h.client.NewRequest("POST", path, params, data)
	if err != nil {
		return nil, err
	}

	m := &Protectiongroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Create a Protection Group Snapshot
func (v *ProtectiongroupService) CreatePgroupSnapshot(pgroup string, params map[string]string) (*ProtectiongroupSnapshot, error) {
	pgroups := []string{pgroup}
	m, err := v.CreatePgroupSnapshots(pgroups, nil)
	if err != nil {
		return nil, err
	}

	return &m[0], err
}

// Send the Protection group snapshot to the target
func (v *ProtectiongroupService) SendPgroupSnapshot(pgroup string, params map[string]string) ([]ProtectiongroupSnapshot, error) {
	type data struct {
		Action string   `json:"action"`
		Source []string `json:"source"`
	}

	d := &data{}
	d.Action = "send"
	pgroups := []string{pgroup}
	d.Source = pgroups
	req, err := v.client.NewRequest("POST", "pgroup", params, d)
	if err != nil {
		return nil, err
	}

	m := []ProtectiongroupSnapshot{}
	_, err = v.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Create Protection Group snapshots for multiple Protection groups.
func (v *ProtectiongroupService) CreatePgroupSnapshots(pgroups []string, params map[string]string) ([]ProtectiongroupSnapshot, error) {
	type data struct {
		Snap   bool     `json:"snap"`
		Source []string `json:"source"`
	}

	d := &data{}
	d.Snap = true
	d.Source = pgroups
	req, err := v.client.NewRequest("POST", "pgroup", params, d)
	if err != nil {
		return nil, err
	}

	m := []ProtectiongroupSnapshot{}
	_, err = v.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Destroy Protection group
func (h *ProtectiongroupService) DestroyProtectiongroup(name string, params map[string]string) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, err := h.client.NewRequest("DELETE", path, params, nil)
	if err != nil {
		return nil, err
	}

	m := &Protectiongroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Disable Protection Group Replication
func (h *ProtectiongroupService) DisablePgroupReplication(pgroup string, params map[string]string) (*Protectiongroup, error) {

	data := map[string]bool{"replicate_enabled": false}
	m, err := h.SetProtectiongroup(pgroup, params, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Enable Protection Group Replication
func (h *ProtectiongroupService) EnablePgroupReplication(pgroup string, params map[string]string) (*Protectiongroup, error) {

	data := map[string]bool{"replicate_enabled": true}
	m, err := h.SetProtectiongroup(pgroup, params, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Disable Protection group snapshot schedule
func (h *ProtectiongroupService) DisablePgroupSnapshots(pgroup string, params map[string]string) (*Protectiongroup, error) {

	data := map[string]bool{"snap_enabled": false}
	m, err := h.SetProtectiongroup(pgroup, params, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Enable Protection Group Snapshot schedule
func (h *ProtectiongroupService) EnablePgroupSnapshots(pgroup string, params map[string]string) (*Protectiongroup, error) {

	data := map[string]bool{"snap_enabled": true}
	m, err := h.SetProtectiongroup(pgroup, params, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Eradicate deleted protection group
func (h *ProtectiongroupService) EradicateProtectiongroup(pgroup string, params map[string]string) (*Protectiongroup, error) {

	data := map[string]bool{"eradicate": true}
	path := fmt.Sprintf("pgroup/%s", pgroup)
	req, err := h.client.NewRequest("DELETE", path, params, data)
	if err != nil {
		return nil, err
	}

	m := &Protectiongroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Get protection group attributes
func (h *ProtectiongroupService) GetProtectiongroup(name string, params map[string]string) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, err := h.client.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	m := &Protectiongroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// List Protection groups
func (h *ProtectiongroupService) ListProtectiongroups(params map[string]string) ([]Protectiongroup, error) {

	req, err := h.client.NewRequest("GET", "pgroup", params, nil)
	m := []Protectiongroup{}
	_, err = h.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Recover deleted protection group
func (h *ProtectiongroupService) RecoverProtectiongroup(pgroup string, params map[string]string) (*Protectiongroup, error) {

	data := map[string]string{"action": "recover"}
	m, err := h.SetProtectiongroup(pgroup, params, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Rename Protection group
func (h *ProtectiongroupService) RenameProtectiongroup(pgroup string, name string, params map[string]string) (*Protectiongroup, error) {

	data := map[string]string{"name": name}
	m, err := h.SetProtectiongroup(pgroup, params, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Set protection group attributes
func (h *ProtectiongroupService) SetProtectiongroup(name string, params map[string]string, data interface{}) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, err := h.client.NewRequest("PUT", path, params, data)
	if err != nil {
		return nil, err
	}

	m := &Protectiongroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
