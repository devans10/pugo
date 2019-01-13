// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type Certificate struct {
	Status      string `json:"status,omitempty"`
	Issued_to   string `json:"issued_to,omitempty"`
	Valid_from  string `json:"valid_from,omitempty"`
	Name        string `json:"name,omitempty"`
	Locality    string `json:"locality,omitempty"`
	Country     string `json:"country,omitempty"`
	Issued_by   string `json:"issued_by,omitempty"`
	Valid_to    string `json:"valid_to,omitempty"`
	State       string `json:"state,omitempty"`
	Key_size    int    `json:"key_size,omitempty"`
	Org_unit    string `json:"organizational_unit,omitempty"`
	Org         string `json:"organization,omitempty"`
	Email       string `json:"email,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	CSR         string `json:"certificate_signing_request,omitempty"`
	CommonName  string `json:"common_name,omitempty"`
	SelfSigned  bool   `json:"self_signed,omitempty"`
}
