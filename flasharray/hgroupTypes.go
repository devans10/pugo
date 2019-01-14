// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// Hostgroup struct for data returned by array
type Hostgroup struct {
	Name  string   `json:"name,omitempty"`
	Hosts []string `json:"hosts,omitempty"`
}

// HostgroupPgroup struct for data returned by array
type HostgroupPgroup struct {
	Name   string `json:"name,omitempty"`
	Pgroup string `json:"protection_group,omitempty"`
}

// HostgroupConnection struct for data returned by array
type HostgroupConnection struct {
	Name string `json:"name,omitempty"`
	Vol  string `json:"vol,omitempty"`
	Lun  int    `json:"lun,omitempty"`
}
