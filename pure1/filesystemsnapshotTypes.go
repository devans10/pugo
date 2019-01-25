// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// FilesystemSnapshot type describes the filesystem snapshot object returned by the API
type FilesystemSnapshot struct {
	ID        string        `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	AsOf      int           `json:"_as_of,omitempty"`
	Arrays    []interface{} `json:"arrays,omitempty"`
	Created   int           `json:"create,omitempty"`
	Destroyed bool          `json:"destroyed,omitempty"`
	On        interface{}   `json:"on,omitempty"`
	Source    interface{}   `json:"source,omitempty"`
	Suffix    string        `json:"suffix,omitempty"`
}
