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

// Filesystem type describes the filesystem object returned by the API
type Filesystem struct {
	ID                         string        `json:"id,omitempty"`
	Name                       string        `json:"name,omitempty"`
	AsOf                       int           `json:"_as_of,omitempty"`
	Arrays                     []interface{} `json:"arrays,omitempty"`
	Created                    int           `json:"create,omitempty"`
	Destroyed                  bool          `json:"destroyed,omitempty"`
	FastRemoveDirectoryEnabled bool          `json:"fast_remove_directory_enabled,omitempty"`
	HardLimitEnabled           bool          `json:"hard_limit_enabled,omitempty"`
	HTTP                       interface{}   `json:"http,omitempty"`
	NFS                        interface{}   `json:"nfs,omitempty"`
	Provisioned                int           `json:"provisioned,omitempty"`
	SMB                        interface{}   `json:"smb,omitempty"`
	SnapshotDirectoryEnabled   bool          `json:"snapshot_directory_enabled,omitempty"`
}
