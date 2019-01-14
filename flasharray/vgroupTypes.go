// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// Vgroup struct for object returned by array
type Vgroup struct {
	Name    string   `json:"name"`
	Volumes []string `json:"volumes"`
}
