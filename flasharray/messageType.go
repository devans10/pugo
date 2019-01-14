// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// Message struct for the object returned by the array
type Message struct {
	ComponentName string `json:"component_name,omitempty"`
	ComponentType string `json:"component_type,omitempty"`
	Details       string `json:"details,omitempty"`
	Event         string `json:"event,omitempty"`
	ID            int    `json:"id,omitempty"`
	Opened        string `json:"opened,omitempty"`
	User          string `json:"user,omitempty"`
}
