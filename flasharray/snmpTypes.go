// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type SnmpManager struct {
	Name               string `json:"name"`
	Notification       string `json:"notification"`
	Community          string `json:"community"`
	Privacy_protocol   string `json:"privacy_protocol"`
	Auth_protocol      string `json:"auth_protocol"`
	Host               string `json:"host"`
	Version            string `json:"version"`
	User               string `json:"user"`
	Privacy_passphrase string `json:"privacy_passphrase"`
	Auth_passphrase    string `json:"auth_passphrase"`
	Engine_id          string `json:"engine_id"`
}
