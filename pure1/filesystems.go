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

import "encoding/json"

// FilesystemService type creates a service to expose array endpoints
type FilesystemService struct {
	client *Client
}

// GetFilesystems returns a list of Filesystem objects
func (f *FilesystemService) GetFilesystems(params map[string]string) ([]Filesystem, error) {
	req, err := f.client.NewRequest("GET", "file-systems", params, nil)
	if err != nil {
		return nil, err
	}

	r := &pure1Response{}

	_, err = f.client.Do(req, r, false)
	if err != nil {
		return nil, err
	}

	m := []Filesystem{}
	for len(m) < r.TotalItems {
		for _, item := range r.Items {
			i := Filesystem{}
			s, _ := json.Marshal(item)
			json.Unmarshal([]byte(s), &i)
			m = append(m, i)
		}

		if len(m) < r.TotalItems {
			if r.ContinuationToken != nil {
				if params == nil {
					params = map[string]string{"continuation_token": r.ContinuationToken.(string)}
				} else {
					params["continuation_token"] = r.ContinuationToken.(string)
				}
				req, err := f.client.NewRequest("GET", "file-systems", params, nil)
				if err != nil {
					return nil, err
				}

				_, err = f.client.Do(req, r, false)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return m, err
}
