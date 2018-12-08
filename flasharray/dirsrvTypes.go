// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type Dirsrv struct {
	Bind_user     string `json:"bind_user"`
	Bind_password string `json:"bind_password"`
	Base_dn       string `json:"base_dn"`
	Check_peer    bool   `json:"check_peer"`
	Enabled       bool   `json:"enabled"`
	Uri           string `json:"uri"`
}

type DirsrvTest struct {
	Output string `json:"output"`
}

type DirsrvRole struct {
	Name       string `json:"name,omitempty"`
	Group      string `json:"group,omitempty"`
	Group_base string `json:"group_base,omitempty"`
}
