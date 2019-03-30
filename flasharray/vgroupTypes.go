// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

// Vgroup struct for object returned by array
type Vgroup struct {
	Name    string   `json:"name"`
	Volumes []string `json:"volumes"`

	// Metrics returned with the action=monitor flag
	WritesPerSec      *int   `json:"writes_per_sec,omitempty"`
	ReadsPerSec       *int   `json:"reads_per_sec,omitempty"`
	UsecPerWriteOp    *int   `json:"usec_per_write_op,omitempty"`
	UsecPerReadOp     *int   `json:"usec_per_read_op,omitempty"`
	SanUsecPerReadOp  *int   `json:"san_usec_per_read_op,omitempty"`
	SanUsecPerWriteOp *int   `json:"san_usec_per_write_op,omitempty"`
	OutputPerSec      *int   `json:"output_per_sec,omitempty"`
	InputPerSec       *int   `json:"input_per_sec,omitempty"`
	Time              string `json:"time,omitempty"`

	// Metrics returned with the space=True flag
	Snapshots        *int     `json:"snapshots,omitempty"`
	DataReduction    *float64 `json:"data_reduction,omitempty"`
	Total            *int     `json:"total,omitempty"`
	ThinProvisioning *float64 `json:"thin_provisioning,omitempty"`
	TotalReduction   *float64 `json:"total_reduction,omitempty"`

	// Metrics returned if action=monitor,size=true
	BytesPerRead  *int `json:"bytes_per_read,omitempty"`
	BytesPerWrite *int `json:"bytes_per_write,omitempty"`
	BytesPerOp    *int `json:"bytes_per_op,omitempty"`
}
