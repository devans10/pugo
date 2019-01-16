// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// Dirsrv struct for data returned by array
type Dirsrv struct {
	BindUser     string   `json:"bind_user"`
	BindPassword string   `json:"bind_password"`
	BaseDn       string   `json:"base_dn"`
	CheckPeer    bool     `json:"check_peer"`
	Enabled      bool     `json:"enabled"`
	URI          []string `json:"uri"`
}

// DirsrvTest struct for data returned by array
type DirsrvTest struct {
	Output string `json:"output"`
}

// DirsrvRole struct for data returned by array
type DirsrvRole struct {
	Name      string `json:"name,omitempty"`
	Group     string `json:"group,omitempty"`
	GroupBase string `json:"group_base,omitempty"`
}
