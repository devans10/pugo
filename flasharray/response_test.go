// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

func respListAlert(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"enabled": false,
							"name": "flasharray-alerts@purestorage.com"
						}
					]`
	return resp[restVersion]
}

func respGetAlertaddress(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"enabled": false,
						"name": "flasharray-alerts@purestorage.com"
					}`
	return resp[restVersion]
}

func respPostAlertaddress(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"enabled": true,
						"name": "admin@example.com"
					}`
	return resp[restVersion]
}

func respPutAlertaddress(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"enabled": false,
						"name": "admin@example.com"
					}`
	return resp[restVersion]
}

func respDeleteAlertaddress(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "admin@example.com"
					}`
	return resp[restVersion]
}

func respGetArray(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"array_name": "pure01",
						"id": "b75f8356-604b-431d-af5c-64c3ca303749",
						"revision": "201712160033+517009f",
						"version": "5.0.0"
					}`
	return resp[restVersion]
}

func respGetArrayControllers(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
		    			{
							"mode": "primary",
							"model": "VMware",
							"name": "CT0",
							"status": "ready",
							"version": "5.0.0"
						}
					]`
	return resp[restVersion]
}

func respGetArraySpace(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"capacity": 3761854481,
							"data_reduction": 1.0,
							"hostname": "pure01",
							"parity": 1.0,
							"shared_space": 0,
							"snapshots": 0,
							"system": 0,
							"thin_provisioning": 0.0,
							"total": 0,
							"total_reduction": 1.0,
							"volumes": 0
						}
					]`
	return resp[restVersion]
}

func respGetArrayMonitor(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"input_per_sec": 0,
							"output_per_sec": 0,
							"queue_depth": 0,
							"reads_per_sec": 0,
							"san_usec_per_read_op": 0,
							"san_usec_per_write_op": 0,
							"time": "2017-12-16T05:12:47Z",
							"usec_per_read_op": 0,
							"usec_per_write_op": 0,
							"writes_per_sec": 0
						}
					]`
	return resp[restVersion]
}
