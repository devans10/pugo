// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type SmtpService struct {
	client *Client
}

// GetSmtp Get the attributes of the current smtp server configuration
func (s *SmtpService) GetSmtp() (*Smtp, error) {

	req, err := s.client.NewRequest("GET", "smtp", nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Smtp{}
	if _, err = s.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}

// SetSmtp Set the attributes of the current smtp server configuration
func (s *SmtpService) SetSmtp(data interface{}) (*Smtp, error) {

	req, err := s.client.NewRequest("POST", "smtp", nil, data)
	if err != nil {
		return nil, err
	}

	m := &Smtp{}
	if _, err = s.client.Do(req, m, false); err != nil {
		return nil, err
	}

	return m, err
}
