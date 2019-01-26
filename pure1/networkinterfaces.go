// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// NetworkInterfacesService type creates a service to expose array endpoints
type NetworkInterfacesService struct {
	client *Client
}

// GetNetworkInterfacess returns a list of Filesystem objects
func (n *NetworkInterfacesService) GetNetworkInterfacess(params map[string]string) ([]NetworkInterface, error) {
	req, err := n.client.NewRequest("GET", "network-interfaces", params, nil)
	if err != nil {
		return nil, err
	}

	m := []NetworkInterface{}
	_, err = n.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
