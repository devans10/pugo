// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// Metric type describes the metric object returned by the API
type Metric struct {
	ID             string        `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	AsOf           int           `json:"_as_of,omitempty"`
	Availabilities []interface{} `json:"availabilities,omitempty"`
	Data           []interface{} `json:"data,omitempty"`
	Description    string        `json:"description,omitempty"`
	ResourceTypes  interface{}   `json:"resource_types,omitempty"`
	Resources      []interface{} `json:"resources,omitempty"`
	Resolution     int           `json:"resolution,omitempty"`
	Unit           string        `json:"unit,omitempty"`
}
