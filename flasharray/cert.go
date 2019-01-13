// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

type CertService struct {
	client *Client
}

// ListCert Lists all available certificates
func (c *CertService) ListCert() ([]Certificate, error) {

	req, err := c.client.NewRequest("GET", "cert", nil, nil)
	if err != nil {
		return nil, err
	}

	m := []Certificate{}
	if _, err = c.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// GetCert Lists attributes or exports the specified certificate
func (c *CertService) GetCert(name string, params map[string]string) (*Certificate, error) {

	path := fmt.Sprintf("cert/%s", name)
	req, err := c.client.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	m := &Certificate{}
	if _, err = c.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// GetCSR Constructs a certificate signing request(CSR) for signing by a certificate authority(CA)
func (c *CertService) GetCSR(name string, params map[string]string) (*Certificate, error) {

	path := fmt.Sprintf("cert/certificate_signing_request/%s", name)
	req, err := c.client.NewRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	m := &Certificate{}
	if _, err = c.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// CreateCert Creates a self-signed certificate or imports a certificate signed by a certificate authority(CA)
func (c *CertService) CreateCert(name string, data interface{}) (*Certificate, error) {

	path := fmt.Sprintf("cert/%s", name)
	req, err := c.client.NewRequest("POST", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Certificate{}
	if _, err = c.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// SetCert Creates (and optionally initializes) a new certificate
func (c *CertService) SetCert(name string, data interface{}) (*Certificate, error) {

	path := fmt.Sprintf("cert/%s", name)
	req, err := c.client.NewRequest("PUT", path, nil, data)
	if err != nil {
		return nil, err
	}

	m := &Certificate{}
	if _, err = c.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// DeleteCert deletes a certificate
func (c *CertService) DeleteCert(name string) (*Certificate, error) {

	path := fmt.Sprintf("cert/%s", name)
	req, err := c.client.NewRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Certificate{}
	if _, err = c.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}
