// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

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

	m := []Array{}
	_, err = a.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetTags returns a list of tags on arrays
func (a *ArrayService) GetTags(params map[string]string) ([]Tag, error) {
	req, err := a.client.NewRequest("GET", "arrays/tags", params, nil)
	if err != nil {
		return nil, err
	}

	m := []Tag{}
	_, err = a.client.Do(req, &m, false)
	if err != nil {
		return nil, err
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

	m := []Tag{}
	_, err = a.client.Do(req, &m, false)
	if err != nil {
		return err
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

	m := []Tag{}
	_, err = a.client.Do(req, &m, false)
	if err != nil {
		return err
	}

	return err
}
