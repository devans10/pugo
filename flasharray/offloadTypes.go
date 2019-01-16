// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// NFSOffload struct is an object returned by the array
type NFSOffload struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	MountPoint   string `json:"mount_point"`
	MountOptions string `json:"mount_options"`
}
