// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

type OffloadService struct {
	client *Client
}

// Connect to NFS Offload
func (o *OffloadService) ConnectNFSOffload(name string, address string, mount_point string) (*NFSOffload, error) {

	data := map[string]string{"name": name, "address": address, "mount_point": mount_point}
	path := fmt.Sprintf("nfs_offload/%s", name)
	req, err := o.client.NewRequest("POST", path, nil, data)
	m := &NFSOffload{}
	_, err = o.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Disconnect an NFS Offload
func (o *OffloadService) DisconnectNFSOffload(name string) (*NFSOffload, error) {

	path := fmt.Sprintf("nfs_offload/%s", name)
	req, err := o.client.NewRequest("DELETE", path, nil, nil)
	m := &NFSOffload{}
	_, err = o.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Get NFS offload attributes
func (o *OffloadService) GetNFSOffload(name string) (*NFSOffload, error) {

	path := fmt.Sprintf("nfs_offload/%s", name)
	req, err := o.client.NewRequest("GET", path, nil, nil)
	m := &NFSOffload{}
	_, err = o.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
