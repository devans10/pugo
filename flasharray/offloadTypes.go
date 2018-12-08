// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type NFSOffload struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	Mount_point   string `json:"mount_point"`
	Mount_options string `json:"mount_options"`
}
