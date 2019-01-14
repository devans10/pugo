// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type Smtp struct {
	Password      string `json:"password,omitempty"`
	User_name     string `json:"user_name,omitempty"`
	Relay_host    string `json:"relay_host,omitempty"`
	Sender_domain string `json:"sender_domain,omitempty"`
}
