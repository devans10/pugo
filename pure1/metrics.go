/*
   Copyright 2018 David Evans

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package pure1

import (
	"encoding/json"
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

	r := &pure1Response{}

	_, err = s.client.Do(req, r, false)
	if err != nil {
		return nil, err
	}

	m := []Metric{}

	for len(m) < r.TotalItems {
		for _, item := range r.Items {
			i := Metric{}
			str, _ := json.Marshal(item)
			json.Unmarshal([]byte(str), &i)
			m = append(m, i)
		}

		if len(m) < r.TotalItems {
			if r.ContinuationToken != nil {
				if params == nil {
					params = map[string]string{"continuation_token": r.ContinuationToken.(string)}
				} else {
					params["continuation_token"] = r.ContinuationToken.(string)
				}
				req, err := s.client.NewRequest("GET", "metrics", params, nil)
				if err != nil {
					return nil, err
				}

				_, err = s.client.Do(req, r, false)
				if err != nil {
					return nil, err
				}
			}
		}
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
	if params == nil {
		params = map[string]string{"aggregation": aggregation}
	} else {
		params["aggregation"] = aggregation
		params["start_time"] = strconv.Itoa(startTime)
		params["end_time"] = strconv.Itoa(endTime)
		params["resolution"] = strconv.Itoa(resolution)
	}

	req, err := s.client.NewRequest("GET", "metrics/history", params, nil)
	if err != nil {
		return nil, err
	}

	r := &pure1Response{}

	_, err = s.client.Do(req, r, false)
	if err != nil {
		return nil, err
	}

	m := []Metric{}

	for len(m) < r.TotalItems {
		for _, item := range r.Items {
			i := Metric{}
			str, _ := json.Marshal(item)
			json.Unmarshal([]byte(str), &i)
			m = append(m, i)
		}

		if len(m) < r.TotalItems {
			if r.ContinuationToken != nil {
				if params == nil {
					params = map[string]string{"continuation_token": r.ContinuationToken.(string)}
				} else {
					params["continuation_token"] = r.ContinuationToken.(string)
				}
				req, err := s.client.NewRequest("GET", "metrics/history", params, nil)
				if err != nil {
					return nil, err
				}

				_, err = s.client.Do(req, r, false)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return m, err
}
