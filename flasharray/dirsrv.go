// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type DirsrvService struct {
	client *Client
}

// Set Directory Service Attributes
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


// Get Directory Service Attributes
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

// Disable Directory Service
// if check_peer is true, enables server authenticity enforcement
func (n *DirsrvService) DisableDirectoryService(check_peer bool) (*Dirsrv, error){

	var data map[string]bool
	if check_peer {
		data = map[string]bool{"check_peer": check_peer}
	} else {
		data = map[string]bool{"enabled": false}
	}

	m, err := n.SetDirectoryService(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Enable Directory Service
// if check_peer is true, enables server authenticity enforcement
func (n *DirsrvService) EnableDirectoryService(check_peer bool) (*Dirsrv, error){

	var data map[string]bool
	if check_peer {
		data = map[string]bool{"check_peer": check_peer}
	} else {
		data = map[string]bool{"enabled": true}
	}

	m, err := n.SetDirectoryService(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Test the directory service
func (n *DirsrvService) TestDirectoryService() (*DirsrvTest, error){

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

// Get directory service groups for roles
func (n *DirsrvService) ListDirectoryServiceRoles() ([]DirsrvRole, error){

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

// Set directory service groups for roles
func (n *DirsrvService) SetDirectoryServiceRoles(data interface{}) (*DirsrvRole, error){

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

