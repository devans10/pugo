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

package pure1

// NetworkInterface type describes the network interface object returned by the API
type NetworkInterface struct {
	ID            string        `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	AsOf          int           `json:"_as_of,omitempty"`
	Arrays        []interface{} `json:"arrays,omitempty"`
	Address       string        `json:"address,omitempty"`
	Enabled       bool          `json:"enabled,omitempty"`
	Gateway       string        `json:"gateway,omitempty"`
	HwAddr        string        `json:"hwaddr,omitempty"`
	MTU           int           `json:"mtu,omitempty"`
	Netmask       string        `json:"netmask,omitempty"`
	Services      []string      `json:"services,omitempty"`
	Speed         int           `json:"speed,omitempty"`
	SubInterfaces []string      `json:"subinterfaces"`
}
