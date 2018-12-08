// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Type string `json:"type"`
}

type ApiToken struct {
	Api_token string `json:"api_token,omitempty"`
	Created   string `json:"created,omitempty"`
	Expires   string `json:"expires,omitempty"`
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
}

type PublicKey struct {
	Publickey string `json:"publickey,omitempty"`
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
}

type GlobalAdmin struct {
}

type LockoutInfo struct {
}
