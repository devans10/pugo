// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type Volume struct {
	Name    string `json:"name,omitempty"`
	Source  string `json:"source,omitempty"`
	Serial  string `json:"serial,omitempty"`
	Size    int    `json:"size,omitempty"`
	Created string `json:"created,omitempty"`
}

type VolumePgroup struct {
	Name   string `json:"name"`
	Pgroup string `json:"protection_group"`
}

type Connection struct {
	Name   string `json:"name,omitempty"`
	Host   string `json:"host,omitempty"`
	Hgroup string `json:"hgroup,omitempty"`
	Lun    int    `json:"lun,omitempty"`
	Size   int    `json:"size,omitempty"`
}

type Block struct {
	Length int `json:"length,omitempty"`
	Offset int `json:offset,omitempty"`
}
