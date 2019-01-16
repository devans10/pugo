// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// NetworkInterface struct for object returned by array
type NetworkInterface struct {
	Name     string   `json:"name,omitempty"`
	Address  string   `json:"address,omitempty"`
	Gateway  string   `json:"gateway,omitempty"`
	Netmask  string   `json:"netmask,omitempty"`
	Enabled  bool     `json:"enabled,omitempty"`
	Subnet   string   `json:"subnet,omitempty"`
	Mtu      int      `json:"mtu,omitempty"`
	Services []string `json:"services,omitempty"`
	Slaves   []string `json:"slaves,omitempty"`
	Hwaddr   string   `json:"hwaddr,omitempty"`
	Speed    int      `json:"speed,omitempty"`
}

// Subnet struct for object returned by array
type Subnet struct {
	Name    string `json:"name"`
	Prefix  string `json:"prefix"`
	Enabled bool   `json:"enabled"`
	Vlan    string `json:"vlan"`
	Gateway string `json:"gateway"`
	Mtu     int    `json:"mtu"`
}

// DNS struct for object returned by array
type DNS struct {
	Nameservers []string `json:"nameservers"`
	Domain      string   `json:"domain"`
}

// Port struct for object returned by array
type Port struct {
	Name     string `json:"name"`
	Portal   string `json:"portal"`
	Failover string `json:"failover"`
	Iqn      string `json:"iqn"`
	Wwn      string `json:"wwn"`
}
