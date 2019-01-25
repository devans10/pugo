// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

import (
	"strconv"
)

// MetricsService type creates a service to expose metrics endpoints
type MetricsService struct {
	client *Client
}

// GetMetrics returns a list of metric objects
func (s *MetricsService) GetMetrics(params map[string]string) ([]Metric, error) {
	req, err := s.client.NewRequest("GET", "metrics", params, nil)
	if err != nil {
		return nil, err
	}

	m := []Metric{}
	_, err = s.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetMetricHistory returns a list of metric objects
// aggregation: 'avg' or 'max'
// endTime: in milliseconds since epoch
// startTime: in milliseconds since epoch
// resolution: duration between data points, in milliseconds
// params: ids or names, and resource_ids or resource_names are required
func (s *MetricsService) GetMetricHistory(aggregation string, startTime int, endTime int, resolution int, params map[string]string) ([]Metric, error) {
	params["aggregation"] = aggregation
	params["start_time"] = strconv.Itoa(startTime)
	params["end_time"] = strconv.Itoa(endTime)
	params["resolution"] = strconv.Itoa(resolution)
	req, err := s.client.NewRequest("GET", "metrics/history", params, nil)
	if err != nil {
		return nil, err
	}

	m := []Metric{}
	_, err = s.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
