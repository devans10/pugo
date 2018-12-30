// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

type Protectiongroup struct {
	Name               string                   `json:"name,omitempty"`
	Hgroups            []string                 `json:"hgroups,omitempty"`
	Source             string                   `json:"source,omitempty"`
	Hosts              []string                 `json:"hosts,omitempty"`
	Volumes            []string                 `json:"volumes,omitempty"`
	Targets            []map[string]interface{} `json:"targets,omitempty"`
	Allfor             int                      `json:"all_for,omitempty"`
	Allowed            bool                     `json:"allowed,omitempty"`
	Days               int                      `json:"days,omitempty"`
	Perday             int                      `json:"per_day,omitempty"`
	ReplicateAt        int                      `json:"replicate_at,omitempty"`
	ReplicateBlackout  map[string]int           `json:"replicate_blackout,omitempty"`
	ReplicateEnabled   bool                     `json:"replicate_enabled,omitempty"`
	ReplicateFrequency int                      `json:"replicate_frequency,omitempty"`
	SnapAt             int                      `json:"snap_at,omitempty"`
	SnapEnabled        bool                     `json:"snap_enabled,omitempty"`
	SnapFrequency      int                      `json:"snap_frequency,omitempty"`
	TargetAllfor       int                      `json:"target_all_for,omitempty"`
	TargetDays         int                      `json:"target_days,omitempty"`
	TargetPerDay       int                      `json:"target_per_day,omitempty"`
}

type ProtectiongroupSnapshot struct {
	Source  string `json:"source"`
	Name    string `json:"name"`
	Created string `json:"created"`
}
