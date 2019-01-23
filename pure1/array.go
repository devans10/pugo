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
	req, err := a.client.NewRequest("GET", "array", params, nil)
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
