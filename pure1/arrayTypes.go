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

// Array type describes the array object returned by the API
type Array struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Model   string `json:"model,omitempty"`
	OS      string `json:"os,omitempty"`
	Version string `json:"version,omitempty"`
	AsOf    int    `json:"_as_of,omitempty"`
}

// Tag type describes the array object returned by the API
type Tag struct {
	Key               string            `json:"key,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	Resource          map[string]string `json:"resource,omitempty"`
	TagOrganizationID int               `json:"tag_organization_id,omitempty"`
	Value             string            `json:"value,omitempty"`
}
