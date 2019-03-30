// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// PodService type creates a service to expose pods endpoints
type PodService struct {
	client *Client
}

// GetPods returns a list of pod objects
func (p *PodService) GetPods(params map[string]string) ([]Pod, error) {
	req, err := p.client.NewRequest("GET", "pods", params, nil)
	if err != nil {
		return nil, err
	}

	m := []Pod{}
	_, err = p.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
