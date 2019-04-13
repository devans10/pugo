// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

// SnmpService struct for snmp API endpoints
type SnmpService struct {
	client *Client
}

// ListSnmp Lists the designated SNMP managers and their communication and security attributes
func (s *SnmpService) ListSnmp(params map[string]string) ([]SnmpManager, error) {

	req, _ := s.client.NewRequest("GET", "snmp", params, nil)
	m := []SnmpManager{}
	if _, err := s.client.Do(req, &m, false); err != nil {
		return nil, err
	}

	return m, nil
}

// GetSnmp Lists communication and security attributes for the specified SNMP manager
func (s *SnmpService) GetSnmp(name string) (*SnmpManager, error) {

	path := fmt.Sprintf("snmp/%s", name)
	req, _ := s.client.NewRequest("GET", path, nil, nil)
	m := &SnmpManager{}
	if _, err := s.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, nil
}

// CreateSnmp Creates a Purity SNMP manager object that identifies a host (SNMP manager)
// and specifies the protocol attributes for communicating with it.
// Once a manager object is created, the transmission of SNMP traps is immediately enabled.
func (s *SnmpService) CreateSnmp(name string, data interface{}) (*SnmpManager, error) {

	path := fmt.Sprintf("snmp/%s", name)
	req, _ := s.client.NewRequest("POST", path, nil, data)
	m := &SnmpManager{}
	if _, err := s.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, nil
}

// SetSnmp Modifies a SNMP manager
func (s *SnmpService) SetSnmp(name string, data interface{}) (*SnmpManager, error) {

	path := fmt.Sprintf("snmp/%s", name)
	req, _ := s.client.NewRequest("PUT", path, nil, data)
	m := &SnmpManager{}
	if _, err := s.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, nil
}

// DeleteSnmp deletes a SNMP Manager
func (s *SnmpService) DeleteSnmp(name string) (*SnmpManager, error) {

	path := fmt.Sprintf("snmp/%s", name)
	req, _ := s.client.NewRequest("DELETE", path, nil, nil)
	m := &SnmpManager{}
	if _, err := s.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, nil
}
