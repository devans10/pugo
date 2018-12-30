// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

type VgroupService struct {
	client *Client
}

// Create Vgroup
func (h *VgroupService) CreateVgroup(name string) (*Vgroup, error) {

	path := fmt.Sprintf("vgroup/%s", name)
	req, err := h.client.NewRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Vgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Destroy Vgroup
func (h *VgroupService) DestroyVgroup(name string) (*Vgroup, error) {

	path := fmt.Sprintf("vgroup/%s", name)
	req, err := h.client.NewRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Vgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Eradicate Vgroup
func (h *VgroupService) EradicateVgroup(vgroup string) (*Vgroup, error) {

	data := map[string]bool{"eradicate": true}
	path := fmt.Sprintf("vgroup/%s", vgroup)
	req, err := h.client.NewRequest("DELETE", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Vgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Get Vgroup attributes
func (h *VgroupService) GetVgroup(name string) (*Vgroup, error) {

	path := fmt.Sprintf("vgroup/%s", name)
	req, err := h.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Vgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// List Vgroups
func (h *VgroupService) ListVgroups() ([]Vgroup, error) {

	req, err := h.client.NewRequest("GET", "vgroup", nil, nil)
	m := []Vgroup{}
	_, err = h.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Recover deleted vgroup
func (h *VgroupService) RecoverVgroup(vgroup string) (*Vgroup, error) {

	data := map[string]string{"action": "recover"}
	m, err := h.SetVgroup(vgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Rename vgroup
func (h *VgroupService) RenameVgroup(vgroup string, name string) (*Vgroup, error) {

	data := map[string]string{"name": name}
	m, err := h.SetVgroup(vgroup, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Set vgroup attribute
func (h *VgroupService) SetVgroup(name string, data interface{}) (*Vgroup, error) {

	path := fmt.Sprintf("vgroup/%s", name)
	req, err := h.client.NewRequest("PUT", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Vgroup{}
	_, err = h.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
