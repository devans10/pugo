// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// Pod struct for object returned by array
type Pod struct {
	Name               string   `json:"name,omitempty"`
	Source             string   `json:"source,omitempty"`
	FailoverPreference []string `json:"failover_preference,omitempty"`
}
