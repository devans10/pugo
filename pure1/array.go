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

import (
	"encoding/json"
)

// ArrayService type creates a service to expose array endpoints
type ArrayService struct {
	client *Client
}

// GetArrays returns a list of FlashArray and FlashBlade objects
func (a *ArrayService) GetArrays(params map[string]string) ([]Array, error) {
	req, err := a.client.NewRequest("GET", "arrays", params, nil)
	if err != nil {
		return nil, err
	}

	r := &pure1Response{}

	_, err = a.client.Do(req, r, false)
	if err != nil {
		return nil, err
	}

	m := []Array{}

	for len(m) < r.TotalItems {
		for _, item := range r.Items {
			i := Array{}
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
				req, err := a.client.NewRequest("GET", "arrays", params, nil)
				if err != nil {
					return nil, err
				}

				_, err = a.client.Do(req, r, false)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return m, err
}

// GetTags returns a list of tags on arrays
func (a *ArrayService) GetTags(params map[string]string) ([]Tag, error) {
	req, err := a.client.NewRequest("GET", "arrays/tags", params, nil)
	if err != nil {
		return nil, err
	}

	r := &pure1Response{}

	_, err = a.client.Do(req, r, false)
	if err != nil {
		return nil, err
	}

	m := []Tag{}
	for len(m) < r.TotalItems {
		for _, item := range r.Items {
			i := Tag{}
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
				req, err := a.client.NewRequest("GET", "arrays/tags", params, nil)
				if err != nil {
					return nil, err
				}

				_, err = a.client.Do(req, r, false)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return m, err
}

// DeleteTags returns a list of tags on arrays
// params resource_id or resource_names is required
// both are comma-separated lists of tags to be deleted
func (a *ArrayService) DeleteTags(params map[string]string) error {
	req, err := a.client.NewRequest("DELETE", "arrays/tags", params, nil)
	if err != nil {
		return err
	}

	r := &pure1Response{}

	_, err = a.client.Do(req, r, false)
	if err != nil {
		return err
	}

	m := []Tag{}
	for len(m) < r.TotalItems {
		for _, item := range r.Items {
			i := Tag{}
			s, _ := json.Marshal(item)
			json.Unmarshal([]byte(s), &i)
			m = append(m, i)
		}

		if r.ContinuationToken != nil {
			params["continuation_token"] = r.ContinuationToken.(string)
			req, err := a.client.NewRequest("GET", "arrays/tags", params, nil)
			if err != nil {
				return err
			}

			_, err = a.client.Do(req, r, false)
			if err != nil {
				return err
			}
		}
	}

	return err
}

// CreateTags creates a list of tags on a array
// params resource_id or resource_names is required
// both are comma-separated lists of tags to be added to
func (a *ArrayService) CreateTags(params map[string]string, data interface{}) error {
	req, err := a.client.NewRequest("PUT", "arrays/tags/batch", params, data)
	if err != nil {
		return err
	}

	r := &pure1Response{}

	_, err = a.client.Do(req, r, false)
	if err != nil {
		return err
	}

	m := []Tag{}
	for len(m) < r.TotalItems {
		for _, item := range r.Items {
			i := Tag{}
			s, _ := json.Marshal(item)
			json.Unmarshal([]byte(s), &i)
			m = append(m, i)
		}

		if r.ContinuationToken != nil {
			params["continuation_token"] = r.ContinuationToken.(string)
			req, err := a.client.NewRequest("GET", "arrays/tags/batch", params, data)
			if err != nil {
				return err
			}

			_, err = a.client.Do(req, r, false)
			if err != nil {
				return err
			}
		}
	}

	return err
}
