// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// SMTP struct for object returned by array
type SMTP struct {
	Password     string `json:"password,omitempty"`
	Username     string `json:"user_name,omitempty"`
	RelayHost    string `json:"relay_host,omitempty"`
	SenderDomain string `json:"sender_domain,omitempty"`
}
