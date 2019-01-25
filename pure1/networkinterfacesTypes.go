// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// NetworkInterface type describes the network interface object returned by the API
type NetworkInterface struct {
	ID            string        `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	AsOf          int           `json:"_as_of,omitempty"`
	Arrays        []interface{} `json:"arrays,omitempty"`
	Address       string        `json:"address,omitempty"`
	Enabled       bool          `json:"enabled,omitempty"`
	Gateway       string        `json:"gateway,omitempty"`
	HwAddr        string        `json:"hwaddr,omitempty"`
	MTU           int           `json:"mtu,omitempty"`
	Netmask       string        `json:"netmask,omitempty"`
	Services      []string      `json:"services,omitempty"`
	Speed         int           `json:"speed,omitempty"`
	SubInterfaces []string      `json:"subinterfaces"`
}
