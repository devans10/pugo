// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type Host struct {
	Name           string   `json:"name,omitempty"`
	Wwn            []string `json:"wwn,omitempty"`
	Iqn            []string `json:"iqn,omitempty"`
	Nqn            []string `json:"nqn,omitempty"`
	HostPassword   string   `json:"host_password,omitempty"`
	HostUser       string   `json:"host_user,omitempty"`
	Personality    string   `json:"personality,omitempty"`
	PreferredArray []string `json:"preferred_array,omitempty"`
	TargetPassword string   `json:"target_password,omitempty"`
	TargetUser     string   `json:"target_user,omitempty"`
	Hgroup         string   `json:"hgroup,omitempty"`
}

type ConnectedVolume struct {
	Vol    string `json:"vol,omitempty"`
	Name   string `json:"name,omitempty"`
	Lun    int    `json:"lun,omitempty"`
	Hgroup string `json:"hgroup,omitempty"`
}

type HostPgroup struct {
	Name   string `json:"name,omitempty"`
	Pgroup string `json:"protection_group,omitempty"`
}
