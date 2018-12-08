// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type Certificate struct {
	Status     string `json:"status"`
	Issued_to  string `json:"issued_to"`
	Valid_from string `json:"valid_from"`
	Name       string `json:"name"`
	Locality   string `json:"locality"`
	Country    string `json:"country"`
	Issued_by  string `json:"issued_by"`
	Valid_to   string `json:"valid_to"`
	State      string `json:"state"`
	Key_size   int    `json:"key_size"`
	Org_unit   string `json:"organizational_unit"`
	Org        string `json:"organization"`
	Email      string `json:"email"`
}

type CSR struct {
	CSR string `json:"certificate_signing_request"`
}
