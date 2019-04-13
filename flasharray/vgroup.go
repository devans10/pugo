// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

// VgroupService struct for vgroup API endpoints
type VgroupService struct {
	client *Client
}

// CreateVgroup creates a Vgroup
func (v *VgroupService) CreateVgroup(name string) (*Vgroup, error) {

	path := fmt.Sprintf("vgroup/%s", name)
	req, _ := v.client.NewRequest("POST", path, nil, nil)
	m := &Vgroup{}
	_, err := v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// DestroyVgroup destroys a Vgroup
func (v *VgroupService) DestroyVgroup(name string) (*Vgroup, error) {

	path := fmt.Sprintf("vgroup/%s", name)
	req, _ := v.client.NewRequest("DELETE", path, nil, nil)
	m := &Vgroup{}
	_, err := v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// EradicateVgroup eradicates a deleted Vgroup
func (v *VgroupService) EradicateVgroup(vgroup string) (*Vgroup, error) {

	data := map[string]bool{"eradicate": true}
	path := fmt.Sprintf("vgroup/%s", vgroup)
	req, _ := v.client.NewRequest("DELETE", path, nil, data)
	m := &Vgroup{}
	_, err := v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// GetVgroup lists Vgroup attributes
func (v *VgroupService) GetVgroup(name string) (*Vgroup, error) {

	path := fmt.Sprintf("vgroup/%s", name)
	req, _ := v.client.NewRequest("GET", path, nil, nil)
	m := &Vgroup{}
	_, err := v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// ListVgroups lists attributes for Vgroups
func (v *VgroupService) ListVgroups() ([]Vgroup, error) {

	req, _ := v.client.NewRequest("GET", "vgroup", nil, nil)
	m := []Vgroup{}
	_, err := v.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// RecoverVgroup recovers deleted vgroup
func (v *VgroupService) RecoverVgroup(vgroup string) (*Vgroup, error) {

	data := map[string]string{"action": "recover"}
	m, err := v.SetVgroup(vgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// RenameVgroup renames a vgroup
func (v *VgroupService) RenameVgroup(vgroup string, name string) (*Vgroup, error) {

	data := map[string]string{"name": name}
	m, err := v.SetVgroup(vgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// SetVgroup modifies vgroup attribute
func (v *VgroupService) SetVgroup(name string, data interface{}) (*Vgroup, error) {

	path := fmt.Sprintf("vgroup/%s", name)
	req, _ := v.client.NewRequest("PUT", path, nil, data)
	m := &Vgroup{}
	_, err := v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}
