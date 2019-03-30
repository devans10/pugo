// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// VolumeSnapshot type describes the VolumeSnapshot object returned by the API
type VolumeSnapshot struct {
	ID            string        `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	AsOf          int           `json:"_as_of,omitempty"`
	Arrays        []interface{} `json:"arrays,omitempty"`
	Created       int           `json:"created,omitempty"`
	Destroyed     bool          `json:"destroyed,omitempty"`
	On            interface{}   `json:"on,omitempty"`
	Pod           interface{}   `json:"pod,omitempty"`
	Provisioned   int           `json:"provisioned,omitempty"`
	Serial        string        `json:"serial,omitempty"`
	SnapshotGroup interface{}   `json:"snapshot_group,omitempty"`
	Source        interface{}   `json:"source,omitempty"`
	Suffix        string        `json:"suffix,omitempty"`
}
