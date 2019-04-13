// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// Drive struct for data returned by array
type Drive struct {
	Name              string `json:"name"`
	Capacity          int    `json:"capacity,omitempty"`
	Details           string `json:"details,omitempty"`
	LastEvacCompleted string `json:"last_evac_completed,omitempty"`
	LastFailure       string `json:"last_failure,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
	Status            string `json:"status,omitempty"`
	Type              string `json:"type,omitempty"`
}

// Component struct for data returned by array
type Component struct {
	Name        string `json:"name"`
	Details     string `json:"details,omitempty"`
	Identify    string `json:"identify,omitempty"`
	Index       int    `json:"index,omitempty"`
	Model       string `json:"model,omitempty"`
	Serial      string `json:"serial,omitempty"`
	Slot        string `json:"slot,omitempty"`
	Speed       int    `json:"speed,omitempty"`
	Status      string `json:"status,omitempty"`
	Temperature int    `json:"temperature,omitempty"`
	Voltage     int    `json:"voltage,omitempty"`
}
