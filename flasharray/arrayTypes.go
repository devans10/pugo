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

// ConsoleLock type console_lock describes the console_lock status of the array.
type ConsoleLock struct {
	ConsoleLock string `json:"console_lock"`
}

// Array gives information about the array
type Array struct {
	ID        string `json:"id,omitempty"`
	ArrayName string `json:"array_name,omitempty"`
	Version   string `json:"version,omitempty"`
	Revision  string `json:"revision,omitempty"`
}

// Phonehome struct is the information returned by array
type Phonehome struct {
	Phonehome string `json:"phonehome,omitempty"`
	Status    string `json:"status,omitempty"`
	Action    string `json:"action,omitempty"`
}

// RemoteAssist struct for information returned by array
type RemoteAssist struct {
	Status string `json:"status,omitempty"`
	Name   string `json:"name,omitempty"`
	Port   string `json:"port,omitempty"`
}

// ArrayConnection struct for information returned by array about
// about the connection to a remote array
type ArrayConnection struct {
	Throttled          bool     `json:"throttled"`
	ArrayName          string   `json:"array_name"`
	Version            string   `json:"version"`
	Connected          bool     `json:"connected"`
	ManagementAddress  string   `json:"management_address"`
	ReplicationAddress string   `json:"replication_address"`
	Type               []string `json:"type"`
	ID                 string   `json:"id"`
}
