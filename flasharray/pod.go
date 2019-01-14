// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

// PodService struct for pod API endpoints
type PodService struct {
	client *Client
}

// ListPods Lists the attributes or displays the performance monitoring details for pods
func (p *PodService) ListPods(params map[string]string) ([]Pod, error) {

	req, err := p.client.NewRequest("GET", "pod", params, nil)
	if err != nil {
		return nil, err
	}

	m := []Pod{}
	if _, err = p.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// GetPod Lists the attributes or displays the performance monitoring details for the specified pod
func (p *PodService) GetPod(name string, params map[string]string) (*Pod, error) {

	path := fmt.Sprintf("pod/%s", name)
	req, err := p.client.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	m := &Pod{}
	if _, err = p.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// CreatePod Creates a new pod
func (p *PodService) CreatePod(pod string, data interface{}) (*Pod, error) {

	path := fmt.Sprintf("pod/%s", pod)
	req, err := p.client.NewRequest("POST", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Pod{}
	if _, err = p.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// ConnectPod Stretches a pod to a peer array
func (p *PodService) ConnectPod(pod string, array string) (*Pod, error) {

	path := fmt.Sprintf("pod/%s/array/%s", pod, array)
	req, err := p.client.NewRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Pod{}
	if _, err = p.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// SetPod Modifies a pod
func (p *PodService) SetPod(pod string, data interface{}) (*Pod, error) {

	path := fmt.Sprintf("pod/%s", pod)
	req, err := p.client.NewRequest("PUT", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Pod{}
	if _, err = p.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// RenamePod renames a pod
func (p *PodService) RenamePod(pod string, name string) (*Pod, error) {

	data := map[string]string{"name": name}
	m, err := p.SetPod(pod, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// RecoverPod recovers a deleted pod
func (p *PodService) RecoverPod(pod string) (*Pod, error) {

	data := map[string]string{"action": "recover"}
	m, err := p.SetPod(pod, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// DeletePod deletes a pod
func (p *PodService) DeletePod(pod string) (*Pod, error) {

	path := fmt.Sprintf("pod/%s", pod)
	req, err := p.client.NewRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Pod{}
	if _, err = p.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// EradicatePod eradicates a deleted pod
func (p *PodService) EradicatePod(pod string) (*Pod, error) {

	path := fmt.Sprintf("pod/%s", pod)
	data := map[string]bool{"eradicate": true}
	req, err := p.client.NewRequest("DELETE", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Pod{}
	if _, err = p.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// DisconnectPod Disconnects a pod frp, a peer array
func (p *PodService) DisconnectPod(pod string, array string) (*Pod, error) {

	path := fmt.Sprintf("pod/%s/array/%s", pod, array)
	req, err := p.client.NewRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Pod{}
	if _, err = p.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}
