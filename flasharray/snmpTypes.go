// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// SnmpManager struct for object returned by array
type SnmpManager struct {
	Name              string `json:"name"`
	Notification      string `json:"notification"`
	Community         string `json:"community"`
	PrivacyProtocol   string `json:"privacy_protocol"`
	AuthProtocol      string `json:"auth_protocol"`
	Host              string `json:"host"`
	Version           string `json:"version"`
	User              string `json:"user"`
	PrivacyPassphrase string `json:"privacy_passphrase"`
	AuthPassphrase    string `json:"auth_passphrase"`
	EngineID          string `json:"engine_id"`
}
