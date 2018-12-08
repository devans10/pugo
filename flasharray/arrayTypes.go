// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// type console_lock describes the console_lock status of the array.
type Console_lock struct {
	Console_lock string `json:"console_lock"`
}

// type Array gives information about the array
type Array struct {
	Id         string `json:"id,omitempty"`
	Array_name string `json:"array_name,omitempty"`
	Version    string `json:"version,omitempty"`
	Revision   string `json:"revision,omitempty"`
}

type Phonehome struct {
	Phonehome string `json:"phonehome,omitempty"`
	Status    string `json:"status,omitempty"`
	Action    string `json:"action,omitempty"`
}

type RemoteAssist struct {
	Status string `json:"status,omitempty"`
	Name   string `json:"name,omitempty"`
	Port   string `json:"port,omitempty"`
}

type ArrayConnection struct {
	Throttled           bool     `json:"throttled"`
	Array_name          string   `json:"array_name"`
	Version             string   `json:"version"`
	Connected           bool     `json:"connected"`
	Management_address  string   `json:"management_address"`
	Replication_address string   `json:"replication_address"`
	Type                []string `json:"type"`
	Id                  string   `json:"id"`
}
