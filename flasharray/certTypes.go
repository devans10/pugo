// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// Certificate is a struct for the cert endpoint data
// returned by the array
type Certificate struct {
	Status      string `json:"status,omitempty"`
	IssuedTo    string `json:"issued_to,omitempty"`
	ValidFrom   string `json:"valid_from,omitempty"`
	Name        string `json:"name,omitempty"`
	Locality    string `json:"locality,omitempty"`
	Country     string `json:"country,omitempty"`
	IssuedBy    string `json:"issued_by,omitempty"`
	ValidTo     string `json:"valid_to,omitempty"`
	State       string `json:"state,omitempty"`
	KeySize     int    `json:"key_size,omitempty"`
	OrgUnit     string `json:"organizational_unit,omitempty"`
	Org         string `json:"organization,omitempty"`
	Email       string `json:"email,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	CSR         string `json:"certificate_signing_request,omitempty"`
	CommonName  string `json:"common_name,omitempty"`
	SelfSigned  bool   `json:"self_signed,omitempty"`
}
