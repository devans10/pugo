// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type Protectiongroup struct {
	Name    string   `json:"name"`
	Hgroups []string `json:"hgroups"`
	Source  string   `json:"source"`
	Hosts   []string `json:"hosts"`
	Volumes []string `json:"volumes"`
	Targets []string `json:"targets"`
}

type ProtectiongroupSnapshot struct {
	Source  string `json:"source"`
	Name    string `json:"name"`
	Created string `json:"created"`
}
