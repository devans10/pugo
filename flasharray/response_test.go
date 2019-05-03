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

package flasharray

func respGetAlert(restVersion string) string {
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

func respGetArrayConnection(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
		  					"throttled": false, 
		  					"array_name": "pure02", 
		  					"version": "5.1.9", 
		  					"connected": true, 
		  					"management_address": null, 
		  					"replication_address": null, 
		  					"type": [
								"async-replication"
		  					], 
		  					"id": "b75f8356-604b-431d-af5c-64c3ca303750"
						}
					  ]`
	return resp[restVersion]
}

func respGetArrayConsoleLock(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"console_lock": "disabled"
					}`
	return resp[restVersion]
}

func respGetArrayPhonehome(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"action": null,
						"status": null
					}`
	return resp[restVersion]
}

func respGetArrayRemoteassist(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pure01-ct0",
						"port": "",
						"status": "disabled"
	  				}`
	return resp[restVersion]
}

func respPutArrayConsoleLock(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"console_lock": "disabled"
					}`
	return resp[restVersion]
}

func respPutArrayPhonehome(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"phonehome": "enabled"
					}`
	return resp[restVersion]
}

func respPutArrayRemoteassist(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pure01-ct0",
						"port": "pure01-ct0.example.com-11679",
						"status": "enabled"
	  				}`
	return resp[restVersion]
}

func respDeleteArrayConnection(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pure02"
					}`
	return resp[restVersion]
}

func respGetVolume(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"created": "2017-12-16T05:12:38Z",
							"name": "v1",
							"serial": "B75F8356604B431D00011010",
							"size": 1073741824,
							"source": null
						},
						{
							"created": "2017-12-16T05:12:38Z",
							"name": "v2",
							"serial": "B75F8356604B431D00011011",
							"size": 2147483648,
							"source": null
						},
						{
							"created": "2017-12-16T05:12:39Z",
							"name": "v3",
							"serial": "B75F8356604B431D00011012",
							"size": 3221225472,
							"source": null
						}
					]`
	return resp[restVersion]
}

func respGetVolumePending(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"created": "2017-12-16T05:12:38Z",
							"name": "v1",
							"serial": "B75F8356604B431D00011010",
							"size": 1073741824,
							"source": null,
							"time_remaining": null
						},
						{
							"created": "2017-12-16T05:12:38Z",
							"name": "v2",
							"serial": "B75F8356604B431D00011011",
							"size": 2147483648,
							"source": null,
							"time_remaining": null
						},
						{
							"created": "2017-12-16T05:12:39Z",
							"name": "v3",
							"serial": "B75F8356604B431D00011012",
							"size": 3221225472,
							"source": null,
							"time_remaining": 86400
						},
					]`
	return resp[restVersion]
}

func respGetVolumePendingOnly(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"created": "2017-12-16T05:12:39Z",
							"name": "v4",
							"serial": "B75F8356604B431D00011013",
							"size": 4294967296,
							"source": null,
							"time_remaining": 86400
						}
					]`
	return resp[restVersion]
}

func respGetVolumeSnap(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"created": "2017-12-16T05:12:40Z",
							"name": "v3.1",
							"serial": "B75F8356604B431D00011014",
							"size": 3221225472,
							"source": "v3"
						},
						{
							"created": "2017-12-16T05:12:40Z",
							"name": "v3.snap1",
							"serial": "B75F8356604B431D00011015",
							"size": 3221225472,
							"source": "v3"
						},
						{
							"created": "2017-12-16T05:12:44Z",
							"name": "pg1.monday.v1",
							"serial": "B75F8356604B431D00011016",
							"size": 1073741824,
							"source": "v1"
						},
					]`
	return resp[restVersion]
}

func respGetVolumeSpace(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"data_reduction": 1.0,
							"name": "v1",
							"shared_space": null,
							"size": 1073741824,
							"snapshots": 0,
							"system": null,
							"thin_provisioning": 1.0,
							"total": 0,
							"total_reduction": 1.0,
							"volumes": 0
						},
						{
							"data_reduction": 1.0,
							"name": "v2",
							"shared_space": null,
							"size": 2147483648,
							"snapshots": 0,
							"system": null,
							"thin_provisioning": 1.0,
							"total": 0,
							"total_reduction": 1.0,
							"volumes": 0
						},
					]`
	return resp[restVersion]
}

func respGetVolumevol(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"created": "2017-12-16T05:12:38Z",
						"name": "v1",
						"serial": "B75F8356604B431D00011010",
						"size": 1073741824,
						"source": null
					}`
	return resp[restVersion]
}

func respGetVolumevolSnap(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"created": "2017-12-16T05:12:40Z",
							"name": "v3.1",
							"serial": "B75F8356604B431D00011014",
							"size": 3221225472,
							"source": "v3"
						},
						{
							"created": "2017-12-16T05:12:40Z",
							"name": "v3.snap1",
							"serial": "B75F8356604B431D00011015",
							"size": 3221225472,
							"source": "v3"
						}
					]`
	return resp[restVersion]
}

func respGetVolumevolMonitor(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
						"input_per_sec": 0,
						"name": "v3",
						"output_per_sec": 0,
						"reads_per_sec": 0,
						"san_usec_per_read_op": 0,
						"san_usec_per_write_op": 0,
						"time": "2017-12-16T05:12:51Z",
						"usec_per_read_op": 0,
						"usec_per_write_op": 0,
						"writes_per_sec": 0
						}
					]`
	return resp[restVersion]
}

func respGetVolumevolDiff(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
						"length": 20480, 
						"offset": 0
						}, 
						{
						"length": 2139095040, 
						"offset": 8388608
						}
					]`
	return resp[restVersion]
}

func respGetVolumevolHgroup(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
						"hgroup": "hg2",
						"host": null,
						"lun": 15,
						"name": "v2",
						"size": 2147483648
						}
					]`
	return resp[restVersion]
}

func respGetVolumevolHost(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
						"host": "h2",
						"lun": 7,
						"name": "v3",
						"size": 3221225472
						}
					]`
	return resp[restVersion]
}

func respPostVolumevol(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"created": "2017-12-16T05:12:52Z",
						"name": "v5",
						"serial": "B75F8356604B431D0001101D",
						"size": 5368709120,
						"source": null
					}`
	return resp[restVersion]
}

func respPostVolume(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
						"created": "2017-12-16T05:12:53Z",
						"name": "v5.1",
						"serial": "B75F8356604B431D00011020",
						"size": 5368709120,
						"source": "v5"
						},
						{
						"created": "2017-12-16T05:12:53Z",
						"name": "v6.4129",
						"serial": "B75F8356604B431D00011021",
						"size": 1073741824,
						"source": "v6"
						}
					]`
	return resp[restVersion]
}

func respPutVolumevolSize(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "v5",
						"size": 1073741824
					}`
	return resp[restVersion]
}

func respPutVolumevolName(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "v5_renamed"
					}`
	return resp[restVersion]
}

func respDeleteVolumevol(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "v5"
					}`
	return resp[restVersion]
}

func respGetHost(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"hgroup": "hg1",
							"iqn": [],
							"name": "h1",
							"wwn": [
								"0000111122223333"
							]
						},
						{
							"hgroup": "hg1",
							"iqn": [],
							"name": "h2",
							"wwn": []
						},
						{
							"hgroup": null,
							"iqn": [],
							"name": "h3",
							"wwn": []
						}
					]`
	return resp[restVersion]
}

func restGetHostMonitor(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"input_per_sec": 0,
							"name": "h1",
							"output_per_sec": 0,
							"reads_per_sec": 0,
							"san_usec_per_read_op": 0,
							"san_usec_per_write_op": 0,
							"time": "2017-12-16T05:12:57Z",
							"usec_per_read_op": 0,
							"usec_per_write_op": 0,
							"writes_per_sec": 0
						},
						{
							"input_per_sec": 0,
							"name": "h2",
							"output_per_sec": 0,
							"reads_per_sec": 0,
							"san_usec_per_read_op": 0,
							"san_usec_per_write_op": 0,
							"time": "2017-12-16T05:12:57Z",
							"usec_per_read_op": 0,
							"usec_per_write_op": 0,
							"writes_per_sec": 0
						}
					]`
	return resp[restVersion]
}

func respGetHosthost(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"hgroup": null,
						"iqn": [],
						"name": "h3",
						"wwn": []
					}`
	return resp[restVersion]
}

func respGetHosthostChap(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"host_password": null,
						"host_user": null,
						"name": "h1",
						"target_password": null,
						"target_user": null
					}`
	return resp[restVersion]
}

func restGetHosthostPersonality(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "h1",
						"personality": null
					}`
	return resp[restVersion]
}

func restGetHosthostMonitor(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"input_per_sec": 0,
						"name": "h1",
						"output_per_sec": 0,
						"reads_per_sec": 0,
						"san_usec_per_read_op": 0,
						"san_usec_per_write_op": 0,
						"time": "2017-12-16T05:12:57Z",
						"usec_per_read_op": 0,
						"usec_per_write_op": 0,
						"writes_per_sec": 0
					}`
	return resp[restVersion]
}

func respGetHosthostVolume(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"hgroup": "hg1",
							"lun": 254,
							"name": "h2",
							"vol": "v1"
						},
						{
							"hgroup": null,
							"lun": 7,
							"name": "h2",
							"vol": "v3"
						}
					]`
	return resp[restVersion]
}

func respGetHosthostVolumePrivate(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"lun": 7,
							"name": "h2",
							"vol": "v3"
						}
					]`
	return resp[restVersion]
}

func respGetHosthostVolumeShared(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"hgroup": "hg1",
							"lun": 254,
							"name": "h2",
							"vol": "v1"
						}
					]`
	return resp[restVersion]
}

func respPostHosthost(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"iqn": [],
						"name": "h4",
						"wwn": [
							"0000999900009999"
						]
					}`
	return resp[restVersion]
}

func respPostHosthostVolumevol(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"lun": 1,
						"name": "h4",
						"vol": "v5_renamed"
					}`
	return resp[restVersion]
}

func respPutHosthost(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"hgroup": null,
						"iqn": [],
						"name": "h4",
						"wwn": [
							"1111222233334444",
							"2222333344445555",
							"4444333322221111",
							"5555444433332222"
						]
					}`
	return resp[restVersion]
}

func respPutHosthostRename(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "h4_renamed"
					}`
	return resp[restVersion]
}

func respDeleteHosthostVolumevol(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "h5",
						"vol": "v5_renamed"
					}`
	return resp[restVersion]
}

func respDeleteHosthost(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "h5"
					}`
	return resp[restVersion]
}

func respGetHgroup(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"hosts": [
								"h1",
								"h2"
							],
							"name": "hg1"
						},
						{
							"hosts": [],
							"name": "hg2"
						}
					]`
	return resp[restVersion]
}

func respGetHgroupSpace(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"data_reduction": 1.0,
							"name": "hg1",
							"size": 1073741824,
							"snapshots": 0,
							"thin_provisioning": 1.0,
							"total": 0,
							"total_reduction": 1.0,
							"volumes": 0
						},
						{
							"data_reduction": 1.0,
							"name": "hg2",
							"size": 2147483648,
							"snapshots": 0,
							"thin_provisioning": 1.0,
							"total": 0,
							"total_reduction": 1.0,
							"volumes": 0
						}
					]`
	return resp[restVersion]
}

func respGetHgroupMonitor(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"input_per_sec": 0,
							"name": "hg2",
							"output_per_sec": 0,
							"reads_per_sec": 0,
							"san_usec_per_read_op": 0,
							"san_usec_per_write_op": 0,
							"time": "2017-12-16T05:13:01Z",
							"usec_per_read_op": 0,
							"usec_per_write_op": 0,
							"writes_per_sec": 0
						},
						{
							"input_per_sec": 0,
							"name": "hg1",
							"output_per_sec": 0,
							"reads_per_sec": 0,
							"san_usec_per_read_op": 0,
							"san_usec_per_write_op": 0,
							"time": "2017-12-16T05:13:01Z",
							"usec_per_read_op": 0,
							"usec_per_write_op": 0,
							"writes_per_sec": 0
						}
					]`
	return resp[restVersion]
}

func respGetHgrouphgroup(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"hosts": [
							"h1",
							"h2"
						],
						"name": "hg1"
					}`
	return resp[restVersion]
}

func respGetHgrouphgroupSpace(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"data_reduction": 1.0,
						"name": "hg1",
						"size": 1073741824,
						"snapshots": 0,
						"thin_provisioning": 1.0,
						"total": 0,
						"total_reduction": 1.0,
						"volumes": 0
					}`
	return resp[restVersion]
}

func respGetHgrouphgroupMonitor(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"input_per_sec": 0,
						"name": "hg1",
						"output_per_sec": 0,
						"reads_per_sec": 0,
						"san_usec_per_read_op": 0,
						"san_usec_per_write_op": 0,
						"time": "2017-12-16T05:13:01Z",
						"usec_per_read_op": 0,
						"usec_per_write_op": 0,
						"writes_per_sec": 0
					}`
	return resp[restVersion]
}

func respGetHgrouphgroupVolume(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"lun": 254,
							"name": "hg1",
							"vol": "v1"
						}
					]`
	return resp[restVersion]
}

func respPostHgrouphgroup(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"hosts": [],
						"name": "hg3"
					}`
	return resp[restVersion]
}

func respPostHgrouphgroupVolumevol(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"lun": 254,
						"name": "hg3",
						"vol": "v3"
					}`
	return resp[restVersion]
}

func respPutHgrouphgroup(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"hosts": [
							"h3"
						],
						"name": "hg4"
					}`
	return resp[restVersion]
}

func respPutHgrouphgroupRename(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"hosts": [
							"h3"
						],
						"name": "hg4_renamed"
					}`
	return resp[restVersion]
}

func respDeleteHgrouphgroupVolumevol(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "hg4",
						"vol": "v3"
					}`
	return resp[restVersion]
}

func respDeleteHgrouphgroup(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "hg4"
					}`
	return resp[restVersion]
}

func respGetPgroup(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"hgroups": null,
							"hosts": null,
							"name": "pg1",
							"source": "pure01",
							"targets": null,
							"volumes": [
								"v1"
							]
						},
						{
							"hgroups": null,
							"hosts": null,
							"name": "pg2",
							"source": "pure01",
							"targets": null,
							"volumes": [
								"v1"
							]
						}
					]`
	return resp[restVersion]
}

func respGetPgroupSchedule(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"name": "pg1",
							"replicate_at": null,
							"replicate_blackout": null,
							"replicate_enabled": false,
							"replicate_frequency": 14400,
							"snap_at": null,
							"snap_enabled": false,
							"snap_frequency": 3600
						},
						{
							"name": "pg2",
							"replicate_at": null,
							"replicate_blackout": null,
							"replicate_enabled": false,
							"replicate_frequency": 14400,
							"snap_at": null,
							"snap_enabled": false,
							"snap_frequency": 3600
						}
					]`
	return resp[restVersion]
}

func respGetPgroupRetention(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"all_for": 86400,
							"days": 7,
							"name": "pg1",
							"per_day": 4,
							"target_all_for": 86400,
							"target_days": 7,
							"target_per_day": 4
						},
						{
							"all_for": 86400,
							"days": 7,
							"name": "pg2",
							"per_day": 4,
							"target_all_for": 86400,
							"target_days": 7,
							"target_per_day": 4
						}
					]`
	return resp[restVersion]
}

func respGetPgroupPending(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"hgroups": null,
							"hosts": null,
							"name": "pg1",
							"source": "pure01",
							"targets": null,
							"time_remaining": null,
							"volumes": [
								"v1"
							]
						},
						{
							"hgroups": null,
							"hosts": null,
							"name": "pg2",
							"source": "pure01",
							"targets": null,
							"time_remaining": 86381,
							"volumes": [
								"v1"
							]
						}
					]`
	return resp[restVersion]
}

func respGetPgroupPendingOnly(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"hgroups": null,
							"hosts": null,
							"name": "pg3",
							"source": "pure01",
							"targets": null,
							"time_remaining": 86381,
							"volumes": [
								"v2",
								"v3"
							]
						}
					]`
	return resp[restVersion]
}

func respGetPgroupSpace(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"name": "pg1",
							"snapshots": 0
						},
						{
							"name": "pg2",
							"snapshots": 0
						}
					]`
	return resp[restVersion]
}

func respGetPgrouppgroup(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"hgroups": null,
						"hosts": null,
						"name": "pg1",
						"source": "pure01",
						"targets": null,
						"volumes": [
							"v1"
						]
					}`
	return resp[restVersion]
}

func respGetPgrouppgroupPending(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"hgroups": null,
						"hosts": null,
						"name": "pg3",
						"source": "pure01",
						"targets": null,
						"time_remaining": 86381,
						"volumes": [
							"v2",
							"v3"
						]
					}`
	return resp[restVersion]
}

func respGetPgrouppgroupSpace(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pg1",
						"snapshots": 0
					}`
	return resp[restVersion]
}

func respGetPgrouppgroupSchedule(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pg1",
						"replicate_at": null,
						"replicate_blackout": null,
						"replicate_enabled": false,
						"replicate_frequency": 14400,
						"snap_at": null,
						"snap_enabled": false,
						"snap_frequency": 3600
					}`
	return resp[restVersion]
}

func respGetPgrouppgroupRetention(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"all_for": 86400,
						"days": 7,
						"name": "pg1",
						"per_day": 4,
						"target_all_for": 86400,
						"target_days": 7,
						"target_per_day": 4
					}`
	return resp[restVersion]
}

func respPostPgroup(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"created": "2017-12-16T05:13:04Z",
							"name": "pg1.3",
							"source": "pg1"
						}
					]`
	return resp[restVersion]
}

func respPostPgrouppgroup(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"hgroups": null,
						"hosts": [
							"h1"
						],
						"name": "pg5",
						"source": "pure01",
						"targets": null,
						"volumes": null
					}`
	return resp[restVersion]
}

func respPutPgrouppgroupRename(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pg6"
					}`
	return resp[restVersion]
}

func respPutPgrouppgroupMembers(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"hgroups": null,
						"hosts": [
							"h3"
						],
						"name": "pg4",
						"source": "pure01",
						"targets": null,
						"volumes": null
					}`
	return resp[restVersion]
}

func respPutPgrouppgroupReplicate(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pg4",
						"replicate_at": null,
						"replicate_blackout": null,
						"replicate_enabled": true,
						"replicate_frequency": 14400,
						"snap_at": null,
						"snap_enabled": false,
						"snap_frequency": 3600
					}`
	return resp[restVersion]
}

func respPutPgrouppgroupRetention(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"all_for": 4320,
						"days": 7,
						"name": "pg4",
						"per_day": 12,
						"target_all_for": 86400,
						"target_days": 7,
						"target_per_day": 4
					}`
	return resp[restVersion]
}

func respDeletePgrouppgroup(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pg5"
	  				}`
	return resp[restVersion]
}

func respGetPort(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"failover": null,
							"iqn": null,
							"name": "CT0.FC0",
							"portal": null,
							"wwn": "5001500150015000"
						},
						{
							"failover": null,
							"iqn": null,
							"name": "CT0.FC1",
							"portal": null,
							"wwn": "5001500150015001"
						},
						{
							"failover": null,
							"iqn": null,
							"name": "CT0.FC2",
							"portal": null,
							"wwn": "5001500150015002"
						},
						{
							"failover": null,
							"iqn": null,
							"name": "CT0.FC3",
							"portal": null,
							"wwn": "5001500150015003"
						}
					]`
	return resp[restVersion]
}

func respGetPortInitiators(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"failover": null,
							"iqn": null,
							"portal": null,
							"target": null,
							"target_iqn": null,
							"target_portal": null,
							"target_wwn": null,
							"wwn": "0000111122223333"
						},
						{
							"failover": null,
							"iqn": null,
							"portal": null,
							"target": "CT0.FC1",
							"target_iqn": null,
							"target_portal": null,
							"target_wwn": "5001500150015001",
							"wwn": "0001000100010001"
						},
					]`
	return resp[restVersion]
}

func respGetMessage(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[]`
	return resp[restVersion]
}

func respGetMessageAudit(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"component_name": null,
							"component_type": "purearray",
							"details": "--senderdomain example.com",
							"event": "setattr",
							"id": 242,
							"opened": "2017-12-16T05:10:10Z",
							"user": "root"
						},
						{
							"component_name": "pureuser",
							"component_type": "pureadmin",
							"details": "--api-token",
							"event": "create",
							"id": 277,
							"opened": "2017-12-16T05:10:11Z",
							"user": "pureuser"
						},
					]`
	return resp[restVersion]
}

func respGetSnmp(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"auth_passphrase": null,
							"auth_protocol": null,
							"community": null,
							"host": "localhost",
							"name": "localhost",
							"notification": null,
							"privacy_passphrase": null,
							"privacy_protocol": null,
							"user": null,
							"version": "v2c"
						}
					]`
	return resp[restVersion]
}

func respGetSnmpsnmp(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"auth_passphrase": null,
						"auth_protocol": null,
						"community": null,
						"host": "localhost",
						"name": "localhost",
						"notification": null,
						"privacy_passphrase": null,
						"privacy_protocol": null,
						"user": null,
						"version": "v2c"
					}`
	return resp[restVersion]
}

func respDeleteSnmpsnmp(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "localhost"
					}`
	return resp[restVersion]
}

func respGetSnmpEngineID(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"engine_id": "80009e2205081114b26f76e863c43f9b"
	  				}`
	return resp[restVersion]
}

func respGetCert(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"country": "US",
							"email": null,
							"issued_by": "name.purestorage.com",
							"issued_to": "name.purestorage.com",
							"key_size": 2048,
							"locality": "Mountain View",
							"name": "kmip",
							"organization": "Pure Storage Inc.",
							"organizational_unit": null,
							"state": "CA",
							"status": "self-signed",
							"valid_from": "2017-12-16T05:12:47Z",
							"valid_to": "2027-12-14T05:12:47Z"
						},
						{
							"country": "US",
							"email": null,
							"issued_by": "db.purestorage.com",
							"issued_to": "db.purestorage.com",
							"key_size": 2048,
							"locality": "Mountain View",
							"name": "management",
							"organization": "Pure Storage Inc.",
							"organizational_unit": null,
							"state": "CA",
							"status": "self-signed",
							"valid_from": "2017-12-16T05:12:46Z",
							"valid_to": "2027-12-14T05:12:46Z"
						}
					]`
	return resp[restVersion]
}

func respGetCertmanagement(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"country": "US",
						"email": null,
						"issued_by": "db.purestorage.com",
						"issued_to": "db.purestorage.com",
						"key_size": 2048,
						"locality": "Mountain View",
						"name": "management",
						"organization": "Pure Storage Inc.",
						"organizational_unit": null,
						"state": "CA",
						"status": "self-signed",
						"valid_from": "2017-12-16T05:12:46Z",
						"valid_to": "2027-12-14T05:12:46Z"
					}`
	return resp[restVersion]
}

func respGetCertkmipCertificate(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"certificate": "-----BEGIN CERTIFICATE-----\nMIIDVjCCAj4CCQD4JW1Ch5/eWjANBgkqhkiG9w0BAQsFADBtMRYwFAYDVQQHDA1N\nb3VudGFpbiBWaWV3MQswCQYDVQQGEwJVUzELMAkGA1UECAwCQ0ExHTAbBgNVBAMM\nFG5hbWUucHVyZXN0b3JhZ2UuY29tMRowGAYDVQQKDBFQdXJlIFN0b3JhZ2UgSW5j\nLjAeFw0xNzEyMTYwNTEyNDdaFw0yNzEyMTQwNTEyNDdaMG0xFjAUBgNVBAcMDU1v\ndW50YWluIFZpZXcxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJDQTEdMBsGA1UEAwwU\nbmFtZS5wdXJlc3RvcmFnZS5jb20xGjAYBgNVBAoMEVB1cmUgU3RvcmFnZSBJbmMu\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3qYqD0ynz1qTJcKm/70Y\nDGVm6zlx92qai1bwd/TRHahC3koACmKteHy0DCdJIIAPGOBAmiC23XRuDY0cHUIM\n89oBcGOIYyTA+QOFoCakInGacp4qKf0Nbcsxcoji3foLGMNuh9/p1M0q7HQrgqfh\nTSKRjEmcyPXr36nv6ONDixZlGlV8zRiaIBancWWp521li8qlLX+4psj5Y2z8c2wR\nKmAs9R8x/FL/CdCnF3NU7CW5pZ3jF3L2VapzophEB5bpJMS8nIuHZpkOl9DX9Lpp\nPY5hxl8K1CArWSNutptYu9K3uuasJQR74OTHHo/MYPWtKQDsMKccrxwqcgmdsN9h\nLwIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQBXdyGS9lO3nyi8ZUYnXgXKTLQ3H6I/\nAxOw6u9vum2BAahKLCdUx9i4iOePxeUwbtMP4Ca3tGwdnBYLGp3+wiMD560m8xP8\ngBn8HrJ1U9uIexJMGql33/5q4BLIITC3vYSl+YhDiHsl94nE6C6wXNKFLIpI2x8K\n3tPl+JWcqeddtfBtkZlDPX2Jow37kiCY2wQ4DiqgpHRtDF7zL8+SN3RUEJYZemvc\nAlAiksAVsKpz7/S4+kkVVtOpMf2mMNpKcDkobpVJQd9DzGp+wlOOY+OlIY9x3QOv\n9ellbqcfgAfr88zbfbYon+4CPrnQ60/AeaaMTJnZf/d7jMF4U3AJtDKG\n-----END CERTIFICATE-----\n"
					}`
	return resp[restVersion]
}

func respGetCertCSR(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"certificate_signing_request": "-----BEGIN CERTIFICATE REQUEST-----\nMIICzjCCAbYCAQAwgYgxFjAUBgNVBAcMDU1vdW50YWluIFZpZXcxCzAJBgNVBAYT\nAlVTMQswCQYDVQQIDAJDQTEbMBkGA1UECwwSUHVyZSBTdG9yYWdlLCBJbmMuMRsw\nGQYDVQQDDBJkYi5wdXJlc3RvcmFnZS5jb20xGjAYBgNVBAoMEVB1cmUgU3RvcmFn\nZSBJbmMuMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApF4soHT3Vap8\nR3vhKZUYzoraYgtmcQUQXXe9DbC6joHjGvDpkPm8zGEHu7S6Tm6RqslrQpglAjq7\ndI/ltIhIz03MOdMhGc2+0eB7ZenMRyHirhTHSxQLEhOxIAPex0K7aJigrSIpKCqE\n9g0wl6LT8/C++wKw8LaeaJhdiZLR7B815v6d/o1vCcarSBNKTddY4Q+ScpqRtcub\nrqy/gY1+laNtsu7hdJTzNft83Hf0lKCLdQgFp4Qb4nMfS7j+8Cz52p/lmw+iWH+Q\nhrhFbGZbk5IjJedTKwWls4bbH2OeksUbBiZ6Gi3PZ5o7d3zmevjOp22BMjwMoaKQ\nUibzdnRaxQIDAQABoAAwDQYJKoZIhvcNAQELBQADggEBACOGXZDrr5KUYz0QZG+2\nTtbXNM872iijFtW6pAdYyyQCYddkWnZenbogmmECG/ttvxCiJYhp8W/gwn6AjbBR\nCOpQg9mRcm9A0yk3v5AJFwwX1NLlgciBwN0niex8SDlSUtkeez0Z/34UH7tWjQhG\nnVso5JlfxTa11bbEe4J6vxWPeScwb9xYFhFiPZDFVGiuj4cK121ElHh+FO9RZn1J\np2VBf4Gvo2fEA71BbTIPG9FqcGskwbAWPGOXcEwhiLrHhubs0RzyoKdenyh/+7y1\nJ5jOi0WyJcP7MFCIBDNz/wjxTAlTFKpAUCIBbBFi73mImQgTC7izbcJpk11Q588K\nq5M=\n-----END CERTIFICATE REQUEST-----\n"
					}`
	return resp[restVersion]
}

func respPutCertcert(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"country": "US",
						"email": null,
						"issued_by": "db.purestorage.com",
						"issued_to": "db.purestorage.com",
						"key_size": 2048,
						"locality": "Mountain View",
						"name": "management",
						"organization": "Pure Storage Inc.",
						"organizational_unit": "Pure Storage, Inc.",
						"state": "CA",
						"status": "self-signed",
						"valid_from": "2017-12-16T05:13:08Z",
						"valid_to": "2027-12-14T05:13:08Z"
					}`
	return resp[restVersion]
}

func respPostCertcert(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"country": null,
						"email": null,
						"issued_by": "db.example.com",
						"issued_to": "db.example.com",
						"key_size": 2048,
						"locality": null,
						"name": "new_cert",
						"organization": "Pure Storage, Inc.",
						"organizational_unit": "Pure Storage, Inc.",
						"state": "FL",
						"status": "self-signed",
						"valid_from": "2017-12-16T05:13:09Z",
						"valid_to": "2027-12-14T05:13:09Z"
					}`
	return resp[restVersion]
}

func respDeleteCertcert(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "new_cert"
					}`
	return resp[restVersion]
}

func respGetDNS(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"domain": "example.com",
						"nameservers": [
							"192.168.0.1",
							"192.168.1.1"
						]
					}`
	return resp[restVersion]
}

func respPutDNS(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"domain": "example.com",
						"nameservers": [
							"192.168.0.1",
							"192.168.1.1"
						]
					}`
	return resp[restVersion]
}

func respGetSubnet(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"enabled": true,
							"gateway": null,
							"interfaces": [],
							"mtu": 1500,
							"name": "subnet100",
							"prefix": "192.168.0.0/24",
							"services": [],
							"vlan": 100
						}
					]`
	return resp[restVersion]
}

func respGetSubnetsubnet(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"enabled": true,
						"gateway": null,
						"interfaces": [],
						"mtu": 1500,
						"name": "subnet100",
						"prefix": "192.168.0.0/24",
						"services": null,
						"vlan": 100
					}`
	return resp[restVersion]
}

func respPostSubnetsubnet(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"enabled": true,
						"gateway": "192.168.100.1",
						"interfaces": null,
						"mtu": 1500,
						"name": "mgmt",
						"prefix": "192.168.100.0/24",
						"services": null,
						"vlan": null
					}`
	return resp[restVersion]
}

func respPutSubnetsubnet(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"enabled": true,
						"gateway": "192.168.100.1",
						"interfaces": null,
						"mtu": 1500,
						"name": "mgmt",
						"prefix": "192.168.100.0/24",
						"services": null,
						"vlan": null
					}`
	return resp[restVersion]
}

func respDeleteSubnetsubnet(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "managementSubnet"
	  				}`
	return resp[restVersion]
}

func respGetNetwork(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"address": "10.14.226.87",
							"enabled": true,
							"gateway": "10.14.224.1",
							"hwaddr": "00:50:56:a5:d9:9a",
							"mtu": 1500,
							"name": "ct0.eth0",
							"netmask": "255.255.240.0",
							"services": [
								"management",
								"replication"
							],
							"slaves": [],
							"speed": 1000000000,
							"subnet": null
						}
					]`
	return resp[restVersion]
}

func respGetNetworkintf(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"address": "10.14.226.87",
						"enabled": true,
						"gateway": "10.14.224.1",
						"hwaddr": "00:50:56:a5:d9:9a",
						"mtu": 1500,
						"name": "ct0.eth0",
						"netmask": "255.255.240.0",
						"services": [
							"management",
							"replication"
						],
						"slaves": [],
						"speed": 1000000000,
						"subnet": null
					}`
	return resp[restVersion]
}

func respPostNetworkVifintf(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"address": "192.168.0.100",
						"enabled": true,
						"gateway": null,
						"hwaddr": "00:50:56:a5:ab:1f",
						"mtu": 1500,
						"name": "ct0.eth1.100",
						"netmask": "255.255.255.0",
						"services": [
							"iscsi",
							"management"
						],
						"slaves": null,
						"speed": 10000000000,
						"subnet": "subnet100"
					}`
	return resp[restVersion]
}

func respPutNetworkintf(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"address": null,
						"enabled": false,
						"gateway": null,
						"hwaddr": "00:50:56:a5:ab:1f",
						"mtu": 1500,
						"name": "ct0.eth1",
						"netmask": null,
						"services": [
							"iscsi",
							"management"
						],
						"slaves": [],
						"speed": 10000000000,
						"subnet": null
					}`
	return resp[restVersion]
}

func respDeleteNetworkintf(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "ct0.eth1.100"
	  				}`
	return resp[restVersion]
}

func respGetHardware(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"details": null,
							"identify": "off",
							"index": 0,
							"model": null,
							"name": "CT0",
							"serial": null,
							"slot": null,
							"speed": null,
							"status": "ok",
							"temperature": null,
							"voltage": null
						}
					]`
	return resp[restVersion]
}

func respGetHardwareComp(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"details": null,
						"identify": "off",
						"index": 0,
						"model": null,
						"name": "SH0.BAY0",
						"serial": null,
						"slot": null,
						"speed": null,
						"status": "ok",
						"temperature": null,
						"voltage": null
					}`
	return resp[restVersion]
}

func respPutHardwareComp(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"identify": "on",
						"index": 0,
						"name": "SH0.BAY0",
						"slot": null
					}`
	return resp[restVersion]
}

func respGetDrive(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"capacity": 494927872,
							"details": null,
							"last_evac_completed": "1970-01-01T00:00:00Z",
							"last_failure": "1970-01-01T00:00:00Z",
							"name": "SH0.BAY0",
							"protocol": "SAS",
							"status": "healthy",
							"type": "SSD"
						},
						{
							"capacity": 494927872,
							"details": null,
							"last_evac_completed": "1970-01-01T00:00:00Z",
							"last_failure": "1970-01-01T00:00:00Z",
							"name": "SH0.BAY1",
							"protocol": "SAS",
							"status": "healthy",
							"type": "SSD"
						}
					]`
	return resp[restVersion]
}

func respGetDrivedrive(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"capacity": 494927872,
						"details": null,
						"last_evac_completed": "1970-01-01T00:00:00Z",
						"last_failure": "1970-01-01T00:00:00Z",
						"name": "SH0.BAY0",
						"protocol": "SAS",
						"status": "healthy",
						"type": "SSD"
					}`
	return resp[restVersion]
}

func respGetAdmin(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"name": "pureuser",
							"role": "admin",
							"type": "local"
						}
					]`
	return resp[restVersion]
}

func respGetAdminadmin(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
							"name": "pureuser",
							"role": "admin",
							"type": "local"
					}`
	return resp[restVersion]
}

func respPutAdminadmin(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
							"name": "pureuser",
							"role": "admin",
							"type": "local"
					}`
	return resp[restVersion]
}

func respPostAdminadmin(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
							"name": "pureuser",
							"role": "admin",
							"type": "local"
					}`
	return resp[restVersion]
}

func respDeleteAdminadmin(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
							"name": "pureuser"
					}`
	return resp[restVersion]
}

func respGetAdminPublicKeys(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[]`
	return resp[restVersion]
}

func respPostAdminadminAPIToken(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"api_token": "****",
						"created": "2017-12-16T05:10:11Z",
						"expires": null,
						"name": "pureuser"
					}`

	return resp[restVersion]
}

func respGetAdminadminAPIToken(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"api_token": "****",
						"created": "2017-12-16T05:10:11Z",
						"expires": null,
						"name": "pureuser"
					}`

	return resp[restVersion]
}

func respDeleteAdminadminAPIToken(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pureuser"
					}`

	return resp[restVersion]
}

func respGetAdminAPIToken(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"api_token": "****",
							"created": "2017-12-16T05:11:04Z",
							"expires": null,
							"name": "os76"
						},
						{
							"api_token": "****",
							"created": "2017-12-16T05:10:11Z",
							"expires": null,
							"name": "pureuser"
						}
					]`
	return resp[restVersion]
}

func respGetDirectoryservice(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"base_dn": "DC=ad1,DC=example,DC=com",
						"bind_password": "****",
						"bind_user": "readonlyuser",
						"check_peer": false,
						"enabled": true,
						"uri": [
							"ldaps://ad1.example.com"
						]
					}`
	return resp[restVersion]
}

func respGetDirectoryserviceCertificate(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"certificate": "-----BEGIN CERTIFICATE-----\nMIIE6jCCBFOgAwIBAgIDEIGKMA0GCSqGSIb3DQEBBQUAME4xCzAJBgNVBAYTAlVT\nMRAwDgYDVQQKEwdFcXVpZmF4MS0wKwYDVQQLEyRFcXVpZmF4IFNlY3VyZSBDZXJ0\naWZpY2F0ZSBBdXRob3JpdHkwHhcNMTAwNDAxMjMwMDE0WhcNMTUwNzAzMDQ1MDAw\nWjCBjzEpMCcGA1UEBRMgMmc4YU81d0kxYktKMlpENTg4VXNMdkRlM2dUYmc4RFUx\nCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRIwEAYDVQQHEwlTdW5u\neXZhbGUxFDASBgNVBAoTC1lhaG9vICBJbmMuMRYwFAYDVQQDEw13d3cueWFob28u\nY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6ZM1jHCkL8rlEKse\n1riTTxyC3WvYQ5m34TlFK7dK4QFI/HPttKGqQm3aVB1Fqi0aiTxe4YQMbd++jnKt\ndjxcpi7sJlFxjMZs4umr1eGo2KgTgSBAJyhxo23k+VpK1SprdPyM3yEfQVdV7JWC\n4Y71CE2nE6+GbsIuhk/to+jJMO7jXx/430jvo8vhNPL6GvWe/D6ObbnxS72ynLSd\nmLtaltykOvZEZiXbbFKgIaYYmCgh89FGVvBkUbGM/Wb5Voiz7ttQLLxKOYRj8Mdk\nTZtzPkM9scIFG1naECPvCxw0NyMyxY3nFOdjUKJ79twanmfCclX2ZO/rk1CpiOuw\nlrrr/QIDAQABo4ICDjCCAgowDgYDVR0PAQH/BAQDAgTwMB0GA1UdDgQWBBSmrfKs\n68m+dDUSf+S7xJrQ/FXAlzA6BgNVHR8EMzAxMC+gLaArhilodHRwOi8vY3JsLmdl\nb3RydXN0LmNvbS9jcmxzL3NlY3VyZWNhLmNybDCCAVsGA1UdEQSCAVIwggFOgg13\nd3cueWFob28uY29tggl5YWhvby5jb22CDHVzLnlhaG9vLmNvbYIMa3IueWFob28u\nY29tggx1ay55YWhvby5jb22CDGllLnlhaG9vLmNvbYIMZnIueWFob28uY29tggxp\nbi55YWhvby5jb22CDGNhLnlhaG9vLmNvbYIMYnIueWFob28uY29tggxkZS55YWhv\nby5jb22CDGVzLnlhaG9vLmNvbYIMbXgueWFob28uY29tggxpdC55YWhvby5jb22C\nDHNnLnlhaG9vLmNvbYIMaWQueWFob28uY29tggxwaC55YWhvby5jb22CDHFjLnlh\naG9vLmNvbYIMdHcueWFob28uY29tggxoay55YWhvby5jb22CDGNuLnlhaG9vLmNv\nbYIMYXUueWFob28uY29tggxhci55YWhvby5jb22CDHZuLnlhaG9vLmNvbTAfBgNV\nHSMEGDAWgBRI5mj5K9KylddH2CMgEE8zmJCf1DAdBgNVHSUEFjAUBggrBgEFBQcD\nAQYIKwYBBQUHAwIwDQYJKoZIhvcNAQEFBQADgYEAp9WOMtcDMM5T0yfPecGv5QhH\nRJZRzgeMPZitLksr1JxxicJrdgv82NWq1bw8aMuRj47ijrtaTEWXaCQCy00yXodD\nzoRJVNoYIvY1arYZf5zv9VZjN5I0HqUc39mNMe9XdZtbkWE+K6yVh6OimKLbizna\ninu9YTrN/4P/w6KzHho=\n-----END CERTIFICATE-----"
	  				}`
	return resp[restVersion]
}

func respGetDirectoryserviceGroups(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"array_admin_group": "pureadmins",
						"group_base": "OU=PureStorage,OU=SAN,OU=IT,OU=US",
						"readonly_group": "purereadonly",
						"storage_admin_group": "pureusers"
					}`
	return resp[restVersion]
}

func respPutDirectoryservice(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"base_dn": "DC=ad1,DC=example,DC=com",
						"bind_password": "****",
						"bind_user": "readonlyuser",
						"check_peer": false,
						"enabled": false,
						"uri": [
							"ldaps://ad1.example.com"
						]
					}`
	return resp[restVersion]
}

func respGetPod(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `[
						{
							"name": "pod1",
							"source": "flasharray1.example.com",
							"failover_preference": [
								"flasharray2.example.com"
							]
						}
					]`
	return resp[restVersion]
}

func respGetPodpod(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pod1",
						"source": "flasharray1.example.com",
						"failover_preference": [
							"flasharray2.example.com"
						]
					}`
	return resp[restVersion]
}

func respPostPodpod(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pod1",
						"source": "flasharray1.example.com",
						"failover_preference": [
							"flasharray2.example.com"
						]
					}`
	return resp[restVersion]
}

func respPutPodpod(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pod1",
						"source": "flasharray1.example.com",
						"failover_preference": [
							"flasharray2.example.com"
						]
					}`
	return resp[restVersion]
}

func respPutPodpodRename(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pod_renamed",
						"source": "flasharray1.example.com",
						"failover_preference": [
							"flasharray2.example.com"
						]
					}`
	return resp[restVersion]
}

func respDeletePodpod(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"name": "pod1"
					}`
	return resp[restVersion]
}

func respGetSMTP(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"user_name": "mailuser",
						"password": "****",
						"relay_host": "mail.example.com",
						"sender_domain": "fa@example.com"
					}`
	return resp[restVersion]
}

func respSetSMTP(restVersion string) string {
	resp := make(map[string]string)
	resp["1.15"] = `{
						"user_name": "mailuser",
						"password": "****",
						"relay_host": "mail.example.com",
						"sender_domain": "fa@example.com"
					}`
	return resp[restVersion]
}
