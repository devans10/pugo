// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

// ProtectiongroupService struct for pgroup API endpoints
type ProtectiongroupService struct {
	client *Client
}

// CreateProtectiongroup creates a Protection group
func (p *ProtectiongroupService) CreateProtectiongroup(name string, data interface{}) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, _ := p.client.NewRequest("POST", path, nil, data)
	m := &Protectiongroup{}
	_, err := p.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// CreatePgroupSnapshot creates a Protection Group Snapshot
func (p *ProtectiongroupService) CreatePgroupSnapshot(pgroup string) (*ProtectiongroupSnapshot, error) {
	pgroups := []string{pgroup}
	m, err := p.CreatePgroupSnapshots(pgroups)
	if err != nil {
		return nil, err
	}

	return &m[0], err
}

// SendPgroupSnapshot sends the Protection group snapshot to the target
func (p *ProtectiongroupService) SendPgroupSnapshot(pgroup string) ([]ProtectiongroupSnapshot, error) {
	data := make(map[string]interface{})
	data["action"] = "send"
	pgroups := []string{pgroup}
	data["source"] = pgroups
	req, _ := p.client.NewRequest("POST", "pgroup", nil, data)
	m := []ProtectiongroupSnapshot{}
	_, err := p.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// CreatePgroupSnapshots creates Protection Group snapshots for multiple Protection groups.
func (p *ProtectiongroupService) CreatePgroupSnapshots(pgroups []string) ([]ProtectiongroupSnapshot, error) {
	data := make(map[string]interface{})
	data["snap"] = true
	data["source"] = pgroups
	req, _ := p.client.NewRequest("POST", "pgroup", nil, data)
	m := []ProtectiongroupSnapshot{}
	_, err := p.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// DestroyProtectiongroup destroys a Protection group
func (p *ProtectiongroupService) DestroyProtectiongroup(name string) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, _ := p.client.NewRequest("DELETE", path, nil, nil)
	m := &Protectiongroup{}
	_, err := p.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// DisablePgroupReplication disables Protection Group Replication
func (p *ProtectiongroupService) DisablePgroupReplication(pgroup string) (*Protectiongroup, error) {

	data := map[string]bool{"replicate_enabled": false}
	m, err := p.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// EnablePgroupReplication enables Protection Group Replication
func (p *ProtectiongroupService) EnablePgroupReplication(pgroup string) (*Protectiongroup, error) {

	data := map[string]bool{"replicate_enabled": true}
	m, err := p.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DisablePgroupSnapshots Protection group snapshot schedule
func (p *ProtectiongroupService) DisablePgroupSnapshots(pgroup string) (*Protectiongroup, error) {

	data := map[string]bool{"snap_enabled": false}
	m, err := p.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// EnablePgroupSnapshots enables Protection Group Snapshot schedule
func (p *ProtectiongroupService) EnablePgroupSnapshots(pgroup string) (*Protectiongroup, error) {

	data := map[string]bool{"snap_enabled": true}
	m, err := p.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// EradicateProtectiongroup eradicates deleted protection group
func (p *ProtectiongroupService) EradicateProtectiongroup(pgroup string) (*Protectiongroup, error) {

	data := map[string]bool{"eradicate": true}
	path := fmt.Sprintf("pgroup/%s", pgroup)
	req, _ := p.client.NewRequest("DELETE", path, nil, data)
	m := &Protectiongroup{}
	_, err := p.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetProtectiongroup protection group attributes
func (p *ProtectiongroupService) GetProtectiongroup(name string, params map[string]string) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, _ := p.client.NewRequest("GET", path, params, nil)
	m := &Protectiongroup{}
	_, err := p.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// ListProtectiongroups lists attributes of the Protection groups
func (p *ProtectiongroupService) ListProtectiongroups(params map[string]string) ([]Protectiongroup, error) {

	req, _ := p.client.NewRequest("GET", "pgroup", params, nil)
	m := []Protectiongroup{}
	_, err := p.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// RecoverProtectiongroup recovers deleted protection group
func (p *ProtectiongroupService) RecoverProtectiongroup(pgroup string) (*Protectiongroup, error) {

	data := map[string]string{"action": "recover"}
	m, err := p.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// RenameProtectiongroup renames Protection group
func (p *ProtectiongroupService) RenameProtectiongroup(pgroup string, name string) (*Protectiongroup, error) {

	data := map[string]string{"name": name}
	m, err := p.SetProtectiongroup(pgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// SetProtectiongroup modifies protection group attributes
func (p *ProtectiongroupService) SetProtectiongroup(name string, data interface{}) (*Protectiongroup, error) {

	path := fmt.Sprintf("pgroup/%s", name)
	req, _ := p.client.NewRequest("PUT", path, nil, data)
	m := &Protectiongroup{}
	_, err := p.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}
