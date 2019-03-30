// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// Filesystem type describes the filesystem object returned by the API
type Filesystem struct {
	ID                         string        `json:"id,omitempty"`
	Name                       string        `json:"name,omitempty"`
	AsOf                       int           `json:"_as_of,omitempty"`
	Arrays                     []interface{} `json:"arrays,omitempty"`
	Created                    int           `json:"create,omitempty"`
	Destroyed                  bool          `json:"destroyed,omitempty"`
	FastRemoveDirectoryEnabled bool          `json:"fast_remove_directory_enabled,omitempty"`
	HardLimitEnabled           bool          `json:"hard_limit_enabled,omitempty"`
	HTTP                       interface{}   `json:"http,omitempty"`
	NFS                        interface{}   `json:"nfs,omitempty"`
	Provisioned                int           `json:"provisioned,omitempty"`
	SMB                        interface{}   `json:"smb,omitempty"`
	SnapshotDirectoryEnabled   bool          `json:"snapshot_directory_enabled,omitempty"`
}
