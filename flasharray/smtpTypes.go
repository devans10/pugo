// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type smtp struct {
	Password      string `json:"password"`
	User_name     string `json:"user_name"`
	Relay_host    string `json:"relay_host"`
	Sender_domain string `json:"sender_domain"`
}
