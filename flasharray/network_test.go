/*
   Copyright 2018 David Evans

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package flasharray

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestListNetworkInterface(t *testing.T) {

	restVersion := "1.15"
	testNetIntf := []NetworkInterface{
		NetworkInterface{
			Address:  "10.14.226.87",
			Enabled:  true,
			Gateway:  "10.14.224.1",
			Hwaddr:   "00:50:56:a5:d9:9a",
			Mtu:      1500,
			Name:     "ct0.eth0",
			Netmask:  "255.255.240.0",
			Services: []string{"management", "replication"},
			Slaves:   []string{},
			Speed:    1000000000,
			Subnet:   "",
		},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetNetwork(restVersion))),
			Header:     head,
		}
	})

	intf, err := c.Networks.ListNetworkInterfaces()
	ok(t, err)
	equals(t, testNetIntf, intf)
}

func TestListNetworkInterfaceError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetNetwork(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.ListNetworkInterfaces()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestGetNetworkInterface(t *testing.T) {

	restVersion := "1.15"
	testNetIntf := NetworkInterface{
		Address:  "10.14.226.87",
		Enabled:  true,
		Gateway:  "10.14.224.1",
		Hwaddr:   "00:50:56:a5:d9:9a",
		Mtu:      1500,
		Name:     "ct0.eth0",
		Netmask:  "255.255.240.0",
		Services: []string{"management", "replication"},
		Slaves:   []string{},
		Speed:    1000000000,
		Subnet:   "",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/ct0.eth0", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetNetworkintf(restVersion))),
			Header:     head,
		}
	})

	intf, err := c.Networks.GetNetworkInterface("ct0.eth0")
	ok(t, err)
	equals(t, &testNetIntf, intf)
}

func TestGettNetworkInterfaceError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/ct0.eth0", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetNetworkintf(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.GetNetworkInterface("ct0.eth0")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestEnableNetworkInterface(t *testing.T) {

	restVersion := "1.15"
	testNetIntf := NetworkInterface{
		Address:  "",
		Enabled:  false,
		Gateway:  "",
		Hwaddr:   "00:50:56:a5:ab:1f",
		Mtu:      1500,
		Name:     "ct0.eth1",
		Netmask:  "",
		Services: []string{"iscsi", "management"},
		Slaves:   []string{},
		Speed:    10000000000,
		Subnet:   "",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/ct0.eth0", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutNetworkintf(restVersion))),
			Header:     head,
		}
	})

	intf, err := c.Networks.EnableNetworkInterface("ct0.eth0")
	ok(t, err)
	equals(t, &testNetIntf, intf)
}

func TestEnableNetworkInterfaceError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/ct0.eth0", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutNetworkintf(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.EnableNetworkInterface("ct0.eth0")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestDisableNetworkInterface(t *testing.T) {

	restVersion := "1.15"
	testNetIntf := NetworkInterface{
		Address:  "",
		Enabled:  false,
		Gateway:  "",
		Hwaddr:   "00:50:56:a5:ab:1f",
		Mtu:      1500,
		Name:     "ct0.eth1",
		Netmask:  "",
		Services: []string{"iscsi", "management"},
		Slaves:   []string{},
		Speed:    10000000000,
		Subnet:   "",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/ct0.eth0", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutNetworkintf(restVersion))),
			Header:     head,
		}
	})

	intf, err := c.Networks.DisableNetworkInterface("ct0.eth0")
	ok(t, err)
	equals(t, &testNetIntf, intf)
}

func TestDisableNetworkInterfaceError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/ct0.eth0", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutNetworkintf(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.DisableNetworkInterface("ct0.eth0")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestSetNetworkInterface(t *testing.T) {

	restVersion := "1.15"
	testNetIntf := NetworkInterface{
		Address:  "",
		Enabled:  false,
		Gateway:  "",
		Hwaddr:   "00:50:56:a5:ab:1f",
		Mtu:      1500,
		Name:     "ct0.eth1",
		Netmask:  "",
		Services: []string{"iscsi", "management"},
		Slaves:   []string{},
		Speed:    10000000000,
		Subnet:   "",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/ct0.eth0", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutNetworkintf(restVersion))),
			Header:     head,
		}
	})

	data := map[string]bool{"enabled": false}
	intf, err := c.Networks.SetNetworkInterface("ct0.eth0", data)
	ok(t, err)
	equals(t, &testNetIntf, intf)
}

func TestSetNetworkInterfaceError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/ct0.eth0", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutNetworkintf(restVersion))),
			Header:     head,
		}
	})

	data := map[string]bool{"enabled": false}
	_, err := c.Networks.SetNetworkInterface("ct0.eth0", data)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestCreateSubnet(t *testing.T) {

	restVersion := "1.15"
	testSubnet := Subnet{
		Enabled: true,
		Gateway: "192.168.100.1",
		Mtu:     1500,
		Name:    "mgmt",
		Prefix:  "192.168.100.0/24",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/mgmt", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	subnet, err := c.Networks.CreateSubnet("mgmt", "192.168.100.0/24")
	ok(t, err)
	equals(t, &testSubnet, subnet)
}

func TestCreateSubnetError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/mgmt", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutNetworkintf(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.CreateSubnet("mgmt", "192.168.100.0/24")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestDeleteSubnet(t *testing.T) {

	restVersion := "1.15"
	testSubnet := Subnet{
		Name: "managementSubnet",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/managementSubnet", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	subnet, err := c.Networks.DeleteSubnet("managementSubnet")
	ok(t, err)
	equals(t, &testSubnet, subnet)
}

func TestDeleteSubnetError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/mgmt", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteNetworkintf(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.DeleteSubnet("mgmt")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestGetSubnet(t *testing.T) {

	restVersion := "1.15"
	testSubnet := Subnet{
		Enabled: true,
		Gateway: "",
		Mtu:     1500,
		Name:    "subnet100",
		Prefix:  "192.168.0.0/24",
		Vlan:    100,
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/subnet100", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	subnet, err := c.Networks.GetSubnet("subnet100")
	ok(t, err)
	equals(t, &testSubnet, subnet)
}

func TestGetSubnetError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/subnet100", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.GetSubnet("subnet100")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestListSubnets(t *testing.T) {

	restVersion := "1.15"
	testSubnet := []Subnet{
		Subnet{
			Enabled:  true,
			Gateway:  "",
			Mtu:      1500,
			Name:     "subnet100",
			Prefix:   "192.168.0.0/24",
			Services: []string{},
			Vlan:     100,
		},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetSubnet(restVersion))),
			Header:     head,
		}
	})

	subnet, err := c.Networks.ListSubnets()
	ok(t, err)
	equals(t, testSubnet, subnet)
}

func TestListSubnetError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetSubnet(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.ListSubnets()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestSetSubnet(t *testing.T) {

	restVersion := "1.15"
	testSubnet := Subnet{
		Enabled: true,
		Gateway: "192.168.100.1",
		Mtu:     1500,
		Name:    "mgmt",
		Prefix:  "192.168.100.0/24",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/subnet100", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	data := map[string]string{"name": "mgmt"}
	subnet, err := c.Networks.SetSubnet("subnet100", data)
	ok(t, err)
	equals(t, &testSubnet, subnet)
}

func TestSetSubnetError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/subnet100", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	data := map[string]string{"name": "mgmt"}
	_, err := c.Networks.SetSubnet("subnet100", data)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestRenameSubnet(t *testing.T) {

	restVersion := "1.15"
	testSubnet := Subnet{
		Enabled: true,
		Gateway: "192.168.100.1",
		Mtu:     1500,
		Name:    "mgmt",
		Prefix:  "192.168.100.0/24",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/subnet100", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	subnet, err := c.Networks.RenameSubnet("subnet100", "mgmt")
	ok(t, err)
	equals(t, &testSubnet, subnet)
}

func TestRenameSubnetError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/subnet100", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.RenameSubnet("subnet100", "mgmt")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestEnableSubnet(t *testing.T) {

	restVersion := "1.15"
	testSubnet := Subnet{
		Enabled: true,
		Gateway: "192.168.100.1",
		Mtu:     1500,
		Name:    "mgmt",
		Prefix:  "192.168.100.0/24",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/subnet100", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	subnet, err := c.Networks.EnableSubnet("subnet100")
	ok(t, err)
	equals(t, &testSubnet, subnet)
}

func TestEnableSubnetError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/subnet100", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.EnableSubnet("subnet100")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestDisableSubnet(t *testing.T) {

	restVersion := "1.15"
	testSubnet := Subnet{
		Enabled: true,
		Gateway: "192.168.100.1",
		Mtu:     1500,
		Name:    "mgmt",
		Prefix:  "192.168.100.0/24",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/subnet100", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	subnet, err := c.Networks.DisableSubnet("subnet100")
	ok(t, err)
	equals(t, &testSubnet, subnet)
}

func TestDisableSubnetError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/subnet/subnet100", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutSubnetsubnet(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.DisableSubnet("subnet100")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestCreateVlanInterface(t *testing.T) {

	restVersion := "1.15"
	testNetIntf := NetworkInterface{
		Address:  "192.168.0.100",
		Enabled:  true,
		Gateway:  "",
		Hwaddr:   "00:50:56:a5:ab:1f",
		Mtu:      1500,
		Name:     "ct0.eth1.100",
		Netmask:  "255.255.255.0",
		Services: []string{"iscsi", "management"},
		Speed:    10000000000,
		Subnet:   "subnet100",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/vif/ct0.eth1.100", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostNetworkVifintf(restVersion))),
			Header:     head,
		}
	})

	intf, err := c.Networks.CreateVlanInterface("ct0.eth1.100", "subnet100")
	ok(t, err)
	equals(t, &testNetIntf, intf)
}

func TestCreateVlanInterfaceError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/vif/ct0.eth1.100", req.URL.String())
		equals(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPostNetworkVifintf(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.CreateVlanInterface("ct0.eth1.100", "subnet100")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestDeleteVlanInterface(t *testing.T) {

	restVersion := "1.15"
	testNetIntf := NetworkInterface{
		Name: "ct0.eth1.100",
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/vif/ct0.eth1.100", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteNetworkintf(restVersion))),
			Header:     head,
		}
	})

	intf, err := c.Networks.DeleteVlanInterface("ct0.eth1.100")
	ok(t, err)
	equals(t, &testNetIntf, intf)
}

func TestDeleteVlanInterfaceError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/network/vif/ct0.eth1.100", req.URL.String())
		equals(t, "DELETE", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respDeleteNetworkintf(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.DeleteVlanInterface("ct0.eth1.100")
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestGetDNS(t *testing.T) {

	restVersion := "1.15"
	testDNS := DNS{
		Domain:      "example.com",
		Nameservers: []string{"192.168.0.1", "192.168.1.1"},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/dns", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetDNS(restVersion))),
			Header:     head,
		}
	})

	intf, err := c.Networks.GetDNS()
	ok(t, err)
	equals(t, &testDNS, intf)
}

func TestGetDNSError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/dns", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetDNS(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.GetDNS()
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestSetDNS(t *testing.T) {

	restVersion := "1.15"
	testDNS := DNS{
		Domain:      "example.com",
		Nameservers: []string{"192.168.0.1", "192.168.1.1"},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/dns", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutDNS(restVersion))),
			Header:     head,
		}
	})

	data := map[string][]string{"nameservers": []string{"192.168.0.1", "192.168.1.1"}}
	intf, err := c.Networks.SetDNS(data)
	ok(t, err)
	equals(t, &testDNS, intf)
}

func TestSetDNSError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/dns", req.URL.String())
		equals(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respPutDNS(restVersion))),
			Header:     head,
		}
	})

	data := map[string][]string{"nameservers": []string{"192.168.0.1", "192.168.1.1"}}
	_, err := c.Networks.SetDNS(data)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}

func TestListPorts(t *testing.T) {

	restVersion := "1.15"
	testPort := []Port{
		Port{
			Failover: "",
			Iqn:      "",
			Name:     "CT0.FC0",
			Portal:   "",
			Wwn:      "5001500150015000",
		},
		Port{
			Failover: "",
			Iqn:      "",
			Name:     "CT0.FC1",
			Portal:   "",
			Wwn:      "5001500150015001",
		},
		Port{
			Failover: "",
			Iqn:      "",
			Name:     "CT0.FC2",
			Portal:   "",
			Wwn:      "5001500150015002",
		},
		Port{
			Failover: "",
			Iqn:      "",
			Name:     "CT0.FC3",
			Portal:   "",
			Wwn:      "5001500150015003",
		},
	}
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/port", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetPort(restVersion))),
			Header:     head,
		}
	})

	intf, err := c.Networks.ListPorts(nil)
	ok(t, err)
	equals(t, testPort, intf)
}

func TestListPortsError(t *testing.T) {

	restVersion := "1.15"
	head := make(http.Header)
	head.Add("Content-Type", "application/json")

	c := testGenerateClient(func(req *http.Request) *http.Response {
		equals(t, "https://flasharray.example.com/api/1.15/port", req.URL.String())
		equals(t, "GET", req.Method)
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respGetPort(restVersion))),
			Header:     head,
		}
	})

	_, err := c.Networks.ListPorts(nil)
	if err == nil {
		t.Errorf("error not raised on 500 response")
	}
}
