package flasharray

import (
	"fmt"
)

type NetworkService struct {
	client *Client
}

// Disable a network interface.
// param: iface: Name of network interface to be disabled.
// Returns an object describing the interface.
func (n *NetworkService) DisableNetworkInterface(iface string, params map[string]string) (*NetworkInterface, error) {

	data := map[string]bool{"enabled": false}
	m, err := n.SetNetworkInterface(iface, params, data)
	if err != nil {
                return nil, err
        }

        return m, err
}

// Enable a network interface.
// param: iface: Name of network interface to be enabled.
// Returns an object describing the interface.
func (n *NetworkService) EnableNetworkInterface(iface string, params map[string]string) (*NetworkInterface, error) {

        data := map[string]bool{"enabled": true}
        m, err := n.SetNetworkInterface(iface, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Get network interface attributes
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

// List network interfaces
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


// Set network interface attributes
func (n *NetworkService) SetNetworkInterface(iface string, params map[string]string, data interface{}) (*NetworkInterface, error) {

        path := fmt.Sprintf("network/%s", iface)
        req, err := n.client.NewRequest("PUT", path, params, data)
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

// Create a subnet
// param: subnet - Name of subnet to be created
// param: prefix - Routing prefix of subnet to be created
// note:
// prefix should be specified as an IPv4 CIDR address.
// ("xxx.xxx.xxx.xxx/nn", representing prefix and prefix length)
func (n *NetworkService) CreateSubnet(subnet string, prefix string, params map[string]string) (*Subnet, error) {

	data := map[string]string{"prefix": prefix}
	path := fmt.Sprintf("subnet/%s", subnet)
	req, err := n.client.NewRequest("POST", path, params, data)
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

// Delete a subnet
// param: subnet - Name of subnet to be deleted
func (n *NetworkService) DeleteSubnet(subnet string, params map[string]string) (*Subnet, error) {

        path := fmt.Sprintf("subnet/%s", subnet)
        req, err := n.client.NewRequest("DELETE", path, params, nil)
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

// Disable a subnet 
// param: subnet: Name of subnet to be disabled.
// Returns an object describing the subnet
func (n *NetworkService) DisableSubnet(subnet string, params map[string]string) (*Subnet, error) {

        data := map[string]bool{"enabled": false}
        m, err := n.SetSubnet(subnet, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Enable a subnet
// param: subnet: Name of subnet to be enabled.
// Returns an object describing the subnet
func (n *NetworkService) EnableSubnet(subnet string, params map[string]string) (*Subnet, error) {

        data := map[string]bool{"enabled": true}
        m, err := n.SetSubnet(subnet, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Get subnet attributes
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

// List Subnets
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

// Rename a subnet
// param: subnet: Name of subnet to be renamed.
// param: name: Name to change the subnet to
// Returns an object describing the subnet
func (n *NetworkService) RenameSubnet(subnet string, name string, params map[string]string) (*Subnet, error) {

        data := map[string]string{"name": name}
        m, err := n.SetSubnet(subnet, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Set subnet attributes
func (n *NetworkService) SetSubnet(subnet string, params map[string]string, data interface{}) (*Subnet, error) {

        path := fmt.Sprintf("subnet/%s", subnet)
        req, err := n.client.NewRequest("PUT", path, params, data)
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

// Create a VLAN Interface
// param: iface - Name of interface to be created
// param: subnet - Subnet to be associated with the new interface
func (n *NetworkService) CreateVlanInterface(iface string, subnet string, params map[string]string) (*VlanInterface, error) {

        data := map[string]string{"subnet": subnet}
        path := fmt.Sprintf("network/vif/%s", iface)
        req, err := n.client.NewRequest("POST", path, params, data)
        if err != nil {
                return nil, err
        }

        m := &VlanInterface{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Delete a VLAN Interface
// param: iface - Name of iface to be deleted
func (n *NetworkService) DeleteVlanInterface(iface string, params map[string]string) (*VlanInterface, error) {

        path := fmt.Sprintf("network/vif/%s", iface)
        req, err := n.client.NewRequest("DELETE", path, params, nil)
        if err != nil {
                return nil, err
        }

        m := &VlanInterface{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Get DNS settings
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

// Set DNS settings
func (n *NetworkService) SetDNS(params map[string]string) (*DNS, error) {

        req, err := n.client.NewRequest("PUT", "dns", params, nil)
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

// List ports
func (n *NetworkService) ListPorts(params map[string]string) ([]Port, error) {

        req, err := n.client.NewRequest("GET", "port", params, nil)
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


