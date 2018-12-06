// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

type HostgroupService struct {
	client *Client
}

// Connect a Volume to a hostgroup
func (h *HostgroupService) ConnectHostgroup(hgroup string, volume string, params map[string]string) (*ConnectedVolume, error) {

        path := fmt.Sprintf("hgroup/%s/volume/%s", hgroup, volume)
        req, err := h.client.NewRequest("POST", path, params, nil)
        m := &ConnectedVolume{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Create a new hostgroup
func (h *HostgroupService) CreateHostgroup(name string, hostlist map[string][]string, params map[string]string) (*Hostgroup, error) {

        path := fmt.Sprintf("hgroup/%s", name)
        req, err := h.client.NewRequest("POST", path, params, hostlist)
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

// Delete a hostgroup
func (h *HostgroupService) DeleteHostgroup(name string, params map[string]string) (*Hostgroup, error) {

        path := fmt.Sprintf("hgroup/%s", name)
        req, err := h.client.NewRequest("DELETE", path, params, nil)
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

// Disconnect a volume from a hostgroup
func (h *HostgroupService) DisconnectHostgroup(hgroup string, volume string, params map[string]string) (*ConnectedVolume, error) {

        path := fmt.Sprintf("hgroup/%s/volume/%s", hgroup, volume)
        req, err := h.client.NewRequest("DELETE", path, params, nil)
        m := &ConnectedVolume{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Return a map of the hostgroup attributes
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

// Add a hostgroup to a Protection Group
func (h *HostgroupService) AddHostgroup(hgroup string, pgroup string, params map[string]string) (*HostgroupPgroup, error) {

        path := fmt.Sprintf("hgroup/%s/pgroup/%s", hgroup, pgroup)
        req, err := h.client.NewRequest("POST", path, params, nil)
        m := &HostgroupPgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Remove a hostgroup from a protection group.
func (h *HostgroupService) RemoveHostgroup(hgroup string, pgroup string, params map[string]string) (*HostgroupPgroup, error) {

        path := fmt.Sprintf("hgroup/%s/pgroup/%s", hgroup, pgroup)
        req, err := h.client.NewRequest("DELETE", path, params, nil)
        m := &HostgroupPgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// List the hostgroup connections
func (h *HostgroupService) ListHostgroupConnections(hgroup string, params map[string]string) ([]HostgroupConnection, error) {

        path := fmt.Sprintf("hgroup/%s/volume", hgroup)
        req, err := h.client.NewRequest("GET", path, params, nil)
        m := []HostgroupConnection{}
        _, err = h.client.Do(req, &m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// List hostgroups
func (h *HostgroupService) ListHostgroups(hgroup string, params map[string]string) ([]Hostgroup, error) {

        path := fmt.Sprintf("hgroup", hgroup)
        req, err := h.client.NewRequest("GET", path, params, nil)
        m := []Hostgroup{}
        _, err = h.client.Do(req, &m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Rename a hostgroup
func (h *HostgroupService) RenameHostgroup(hgroup string, name string, params map[string]string) (*Hostgroup, error) {

        data := map[string]string{"name": name}
        m, err := h.SetHostgroup(hgroup, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Set the hostgroup attributes
func (h *HostgroupService) SetHostgroup(name string, params map[string]string, data interface{}) (*Hostgroup, error) {

        path := fmt.Sprintf("hgroup/%s", name)
        req, err := h.client.NewRequest("PUT", path, params, data)
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
