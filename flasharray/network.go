// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

// NetworkService struct for network API endpoints
type NetworkService struct {
	client *Client
}

// DisableNetworkInterface disables a network interface.
// param: iface: Name of network interface to be disabled.
// Returns an object describing the interface.
func (n *NetworkService) DisableNetworkInterface(iface string) (*NetworkInterface, error) {

	data := map[string]bool{"enabled": false}
	m, err := n.SetNetworkInterface(iface, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// EnableNetworkInterface enables a network interface.
// param: iface: Name of network interface to be enabled.
// Returns an object describing the interface.
func (n *NetworkService) EnableNetworkInterface(iface string) (*NetworkInterface, error) {

	data := map[string]bool{"enabled": true}
	m, err := n.SetNetworkInterface(iface, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetNetworkInterface lists network interface attributes
func (n *NetworkService) GetNetworkInterface(iface string) (*NetworkInterface, error) {

	path := fmt.Sprintf("network/%s", iface)
	req, err := n.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &NetworkInterface{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ListNetworkInterfaces list the attributes of the network interfaces
func (n *NetworkService) ListNetworkInterfaces() ([]NetworkInterface, error) {

	req, err := n.client.NewRequest("GET", "network", nil, nil)
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

// SetNetworkInterface modifies network interface attributes
func (n *NetworkService) SetNetworkInterface(iface string, data interface{}) (*NetworkInterface, error) {

	path := fmt.Sprintf("network/%s", iface)
	req, err := n.client.NewRequest("PUT", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &NetworkInterface{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// CreateSubnet creates a subnet
// param: subnet - Name of subnet to be created
// param: prefix - Routing prefix of subnet to be created
// note:
// prefix should be specified as an IPv4 CIDR address.
// ("xxx.xxx.xxx.xxx/nn", representing prefix and prefix length)
func (n *NetworkService) CreateSubnet(subnet string, prefix string) (*Subnet, error) {

	data := map[string]string{"prefix": prefix}
	path := fmt.Sprintf("subnet/%s", subnet)
	req, err := n.client.NewRequest("POST", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Subnet{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DeleteSubnet deletes a subnet
// param: subnet - Name of subnet to be deleted
func (n *NetworkService) DeleteSubnet(subnet string) (*Subnet, error) {

	path := fmt.Sprintf("subnet/%s", subnet)
	req, err := n.client.NewRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Subnet{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DisableSubnet disables a subnet
// param: subnet: Name of subnet to be disabled.
// Returns an object describing the subnet
func (n *NetworkService) DisableSubnet(subnet string) (*Subnet, error) {

	data := map[string]bool{"enabled": false}
	m, err := n.SetSubnet(subnet, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// EnableSubnet enables a subnet
// param: subnet: Name of subnet to be enabled.
// Returns an object describing the subnet
func (n *NetworkService) EnableSubnet(subnet string) (*Subnet, error) {

	data := map[string]bool{"enabled": true}
	m, err := n.SetSubnet(subnet, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetSubnet lists subnet attributes
func (n *NetworkService) GetSubnet(subnet string) (*Subnet, error) {

	path := fmt.Sprintf("sbunet/%s", subnet)
	req, err := n.client.NewRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Subnet{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ListSubnets lists attributes of subnets
func (n *NetworkService) ListSubnets() ([]Subnet, error) {

	req, err := n.client.NewRequest("GET", "subnet", nil, nil)
	if err != nil {
		return nil, err
	}

	m := []Subnet{}
	_, err = n.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// RenameSubnet renames a subnet
// param: subnet: Name of subnet to be renamed.
// param: name: Name to change the subnet to
// Returns an object describing the subnet
func (n *NetworkService) RenameSubnet(subnet string, name string) (*Subnet, error) {

	data := map[string]string{"name": name}
	m, err := n.SetSubnet(subnet, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// SetSubnet modifies subnet attributes
func (n *NetworkService) SetSubnet(subnet string, data interface{}) (*Subnet, error) {

	path := fmt.Sprintf("subnet/%s", subnet)
	req, err := n.client.NewRequest("PUT", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Subnet{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// CreateVlanInterface creates a VLAN Interface
// param: iface - Name of interface to be created
// param: subnet - Subnet to be associated with the new interface
func (n *NetworkService) CreateVlanInterface(iface string, subnet string) (*NetworkInterface, error) {

	data := map[string]string{"subnet": subnet}
	path := fmt.Sprintf("network/vif/%s", iface)
	req, err := n.client.NewRequest("POST", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &NetworkInterface{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DeleteVlanInterface deletes a VLAN Interface
// param: iface - Name of iface to be deleted
func (n *NetworkService) DeleteVlanInterface(iface string) (*NetworkInterface, error) {

	path := fmt.Sprintf("network/vif/%s", iface)
	req, err := n.client.NewRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &NetworkInterface{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetDNS gets DNS settings
func (n *NetworkService) GetDNS() (*DNS, error) {

	req, err := n.client.NewRequest("GET", "dns", nil, nil)
	if err != nil {
		return nil, err
	}

	m := &DNS{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// SetDNS modifies DNS settings
func (n *NetworkService) SetDNS(data interface{}) (*DNS, error) {

	req, err := n.client.NewRequest("PUT", "dns", nil, data)
	if err != nil {
		return nil, err
	}

	m := &DNS{}
	_, err = n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ListPorts lists attributes of the ports
func (n *NetworkService) ListPorts(data interface{}) ([]Port, error) {

	req, err := n.client.NewRequest("GET", "port", nil, data)
	if err != nil {
		return nil, err
	}

	m := []Port{}
	_, err = n.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
