// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// Pod type describes the network interface object returned by the API
type Pod struct {
	ID       string        `json:"id,omitempty"`
	Name     string        `json:"name,omitempty"`
	AsOf     int           `json:"_as_of,omitempty"`
	Arrays   []interface{} `json:"arrays,omitempty"`
	Mediator string        `json:"mediator,omitempty"`
	Source   interface{}   `json:"source,omitempty"`
}
