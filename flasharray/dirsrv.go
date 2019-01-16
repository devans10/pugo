// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// DirsrvService struct for the dirsrv endpoints
type DirsrvService struct {
	client *Client
}

// SetDirectoryService sets attributes for the directory service
func (n *DirsrvService) SetDirectoryService(data interface{}) (*Dirsrv, error) {

	req, err := n.client.NewRequest("PUT", "directoryservice", nil, data)
	if err != nil {
		return nil, err
	}

	m := &Dirsrv{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetDirectoryService lists the attributes for the directory service
func (n *DirsrvService) GetDirectoryService() (*Dirsrv, error) {

	req, err := n.client.NewRequest("GET", "directoryservice", nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Dirsrv{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DisableDirectoryService disables the directory service
// if check_peer is true, enables server authenticity enforcement
func (n *DirsrvService) DisableDirectoryService(checkPeer bool) (*Dirsrv, error) {

	var data map[string]bool
	if checkPeer {
		data = map[string]bool{"check_peer": checkPeer}
	} else {
		data = map[string]bool{"enabled": false}
	}

	m, err := n.SetDirectoryService(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// EnableDirectoryService enables the directory service
// if check_peer is true, enables server authenticity enforcement
func (n *DirsrvService) EnableDirectoryService(checkPeer bool) (*Dirsrv, error) {

	var data map[string]bool
	if checkPeer {
		data = map[string]bool{"check_peer": checkPeer}
	} else {
		data = map[string]bool{"enabled": true}
	}

	m, err := n.SetDirectoryService(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// TestDirectoryService tests the directory service connection
func (n *DirsrvService) TestDirectoryService() (*DirsrvTest, error) {

	data := map[string]string{"action": "test"}
	req, err := n.client.NewRequest("PUT", "directoryservice", nil, data)
	if err != nil {
		return nil, err
	}

	m := &DirsrvTest{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ListDirectoryServiceRoles get directory service groups for roles
func (n *DirsrvService) ListDirectoryServiceRoles() ([]DirsrvRole, error) {

	req, err := n.client.NewRequest("GET", "directoryservice/role", nil, nil)
	if err != nil {
		return nil, err
	}

	m := []DirsrvRole{}
	_, err = n.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// SetDirectoryServiceRoles sets the groups for roles
func (n *DirsrvService) SetDirectoryServiceRoles(data interface{}) (*DirsrvRole, error) {

	req, err := n.client.NewRequest("PUT", "directoryservice/role", nil, data)
	if err != nil {
		return nil, err
	}

	m := &DirsrvRole{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
