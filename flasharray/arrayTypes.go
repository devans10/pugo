// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// ConsoleLock type console_lock describes the console_lock status of the array.
type ConsoleLock struct {
	ConsoleLock string `json:"console_lock"`
}

// Array gives information about the array
type Array struct {
	ID        string `json:"id,omitempty"`
	ArrayName string `json:"array_name,omitempty"`
	Version   string `json:"version,omitempty"`
	Revision  string `json:"revision,omitempty"`
}

// Phonehome struct is the information returned by array
type Phonehome struct {
	Phonehome string `json:"phonehome,omitempty"`
	Status    string `json:"status,omitempty"`
	Action    string `json:"action,omitempty"`
}

// RemoteAssist struct for information returned by array
type RemoteAssist struct {
	Status string `json:"status,omitempty"`
	Name   string `json:"name,omitempty"`
	Port   string `json:"port,omitempty"`
}

// ArrayConnection struct for information returned by array about
// about the connection to a remote array
type ArrayConnection struct {
	Throttled          bool     `json:"throttled"`
	ArrayName          string   `json:"array_name"`
	Version            string   `json:"version"`
	Connected          bool     `json:"connected"`
	ManagementAddress  string   `json:"management_address"`
	ReplicationAddress string   `json:"replication_address"`
	Type               []string `json:"type"`
	ID                 string   `json:"id"`
}
