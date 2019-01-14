// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// User struct for object returned by array
type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Type string `json:"type"`
}

// Token struct for object returned by array
type Token struct {
	APIToken string `json:"api_token,omitempty"`
	Created  string `json:"created,omitempty"`
	Expires  string `json:"expires,omitempty"`
	Type     string `json:"type,omitempty"`
	Name     string `json:"name,omitempty"`
}

// PublicKey struct for object returned by array
type PublicKey struct {
	Publickey string `json:"publickey,omitempty"`
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
}

// GlobalAdmin struct for object returned by array
type GlobalAdmin struct {
}

// LockoutInfo struct for object returned by array
type LockoutInfo struct {
}
