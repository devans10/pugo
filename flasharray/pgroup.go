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
func (h *ProtectiongroupService) CreateProtectiongroup(name string, data interface{}) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, err := h.client.NewRequest("POST", path, nil, data)
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
func (v *ProtectiongroupService) CreatePgroupSnapshot(pgroup string) (*ProtectiongroupSnapshot, error) {
	pgroups := []string{pgroup}
	m, err := v.CreatePgroupSnapshots(pgroups)
	if err != nil {
		return nil, err
	}

	return &m[0], err
}

// Send the Protection group snapshot to the target
func (v *ProtectiongroupService) SendPgroupSnapshot(pgroup string) ([]ProtectiongroupSnapshot, error) {
	data := make(map[string]interface{})
	data["action"] = "send"
	pgroups := []string{pgroup}
	data["source"] = pgroups
	req, err := v.client.NewRequest("POST", "pgroup", nil, data)
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
func (v *ProtectiongroupService) CreatePgroupSnapshots(pgroups []string) ([]ProtectiongroupSnapshot, error) {
	data := make(map[string]interface{})
	data["snap"] = true
	data["source"] = pgroups
	req, err := v.client.NewRequest("POST", "pgroup", nil, data)
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
func (h *ProtectiongroupService) DestroyProtectiongroup(name string) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, err := h.client.NewRequest("DELETE", path, nil, nil)
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
func (h *ProtectiongroupService) DisablePgroupReplication(pgroup string) (*Protectiongroup, error) {

	data := map[string]bool{"replicate_enabled": false}
	m, err := h.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Enable Protection Group Replication
func (h *ProtectiongroupService) EnablePgroupReplication(pgroup string) (*Protectiongroup, error) {

	data := map[string]bool{"replicate_enabled": true}
	m, err := h.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Disable Protection group snapshot schedule
func (h *ProtectiongroupService) DisablePgroupSnapshots(pgroup string) (*Protectiongroup, error) {

	data := map[string]bool{"snap_enabled": false}
	m, err := h.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Enable Protection Group Snapshot schedule
func (h *ProtectiongroupService) EnablePgroupSnapshots(pgroup string) (*Protectiongroup, error) {

	data := map[string]bool{"snap_enabled": true}
	m, err := h.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Eradicate deleted protection group
func (h *ProtectiongroupService) EradicateProtectiongroup(pgroup string) (*Protectiongroup, error) {

	data := map[string]bool{"eradicate": true}
	path := fmt.Sprintf("pgroup/%s", pgroup)
	req, err := h.client.NewRequest("DELETE", path, nil, data)
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
func (h *ProtectiongroupService) GetProtectiongroup(name string, data interface{}) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, err := h.client.NewRequest("GET", path, nil, data)
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
func (h *ProtectiongroupService) ListProtectiongroups() ([]Protectiongroup, error) {

	req, err := h.client.NewRequest("GET", "pgroup", nil, nil)
	m := []Protectiongroup{}
	_, err = h.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Recover deleted protection group
func (h *ProtectiongroupService) RecoverProtectiongroup(pgroup string) (*Protectiongroup, error) {

	data := map[string]string{"action": "recover"}
	m, err := h.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Rename Protection group
func (h *ProtectiongroupService) RenameProtectiongroup(pgroup string, name string) (*Protectiongroup, error) {

	data := map[string]string{"name": name}
	m, err := h.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Set protection group attributes
func (h *ProtectiongroupService) SetProtectiongroup(name string, data interface{}) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, err := h.client.NewRequest("PUT", path, nil, data)
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
