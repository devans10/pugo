// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// SMTPService struct for smtp API endpoints
type SMTPService struct {
	client *Client
}

// GetSMTP Get the attributes of the current smtp server configuration
func (s *SMTPService) GetSMTP() (*SMTP, error) {

	req, err := s.client.NewRequest("GET", "smtp", nil, nil)
	if err != nil {
		return nil, err
	}

	m := &SMTP{}
	if _, err = s.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// SetSMTP Set the attributes of the current smtp server configuration
func (s *SMTPService) SetSMTP(data interface{}) (*SMTP, error) {

	req, err := s.client.NewRequest("POST", "smtp", nil, data)
	if err != nil {
		return nil, err
	}

	m := &SMTP{}
	if _, err = s.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}
