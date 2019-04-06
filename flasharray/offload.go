// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

// OffloadService struct for offload API endpoints
type OffloadService struct {
	client *Client
}

// ConnectNFSOffload connects array to NFS Offload server
func (o *OffloadService) ConnectNFSOffload(name string, address string, mountPoint string) (*NFSOffload, error) {

	data := map[string]string{"name": name, "address": address, "mount_point": mountPoint}
	path := fmt.Sprintf("nfs_offload/%s", name)
	req, err := o.client.NewRequest("POST", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &NFSOffload{}
	_, err = o.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DisconnectNFSOffload disconnects array from an NFS Offload server
func (o *OffloadService) DisconnectNFSOffload(name string) (*NFSOffload, error) {

	path := fmt.Sprintf("nfs_offload/%s", name)
	req, err := o.client.NewRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &NFSOffload{}
	_, err = o.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetNFSOffload lists NFS offload attributes
func (o *OffloadService) GetNFSOffload(name string) (*NFSOffload, error) {

	path := fmt.Sprintf("nfs_offload/%s", name)
	req, err := o.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &NFSOffload{}
	_, err = o.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
