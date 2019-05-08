package pure1

func respGetArrays(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"continuation_token": "MThkMzJmYWUtZDI3OS00ODEzLWIzODYtMmM3NmFhMTMwM2My",
		"total_item_count": 1,
		"items": [
		  {
			"_as_of": 1502729489760,
			"id": "01c2889a-4124-49ff-8cbd-c33405ede123",
			"name": "example_name",
			"model": "FA-m70r2",
			"os": "Purity//FA",
			"version": "5.0.0"
		  }
		]
	  }`
	return resp[restVersion]
}

func respGetArraysTags(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"continuation_token": "MThkMzJmYWUtZDI3OS00ODEzLWIzODYtMmM3NmFhMTMwM2My",
		"total_item_count": 1,
		"items": [
		  {
			"key": "example_key",
			"namespace": "pure1",
			"resource": {
			  "id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			  "name": "example_name",
			  "resource_type": "example_resource_type"
			},
			"tag_organization_id": 101,
			"value": "value123"
		  }
		]
	  }`
	return resp[restVersion]
}

func respDeleteArraysTags(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{}`
	return resp[restVersion]
}

func respPutArraysTagsBatch(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"items": [
		  {
			"key": "example_key",
			"namespace": "pure1",
			"resource": {
			  "id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			  "name": "example_name",
			  "resource_type": "example_resource_type"
			},
			"tag_organization_id": 101,
			"value": "value123"
		  }
		]
	  }`
	return resp[restVersion]
}

func respGetAudits(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"continuation_token": "MThkMzJmYWUtZDI3OS00ODEzLWIzODYtMmM3NmFhMTMwM2My",
		"total_item_count": 1,
		"items": [
		  {
			"_as_of": 1502729489760,
			"id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			"name": "example_name",
			"arrays": [
			  {
				"id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
				"name": "example_name",
				"resource_type": "example_resource_type"
			  }
			],
			"arguments": "--args",
			"command": "purevol",
			"origin": "array",
			"subcommand": "snap",
			"time": 1502729489760,
			"user": "pureuser"
		  }
		]
	  }`
	return resp[restVersion]
}

func respGetFilesystems(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"continuation_token": "MThkMzJmYWUtZDI3OS00ODEzLWIzODYtMmM3NmFhMTMwM2My",
		"total_item_count": 1,
		"items": [
		  {
			"_as_of": 1502729489760,
			"id": "01c2889a-4124-49ff-8cbd-c33405ede123",
			"name": "example_name",
			"arrays": [
			  {
				"id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
				"name": "example_name",
				"resource_type": "example_resource_type"
			  }
			],
			"created": 1502729489760,
			"destroyed": false,
			"fast_remove_directory_enabled": false,
			"hard_limit_enabled": false,
			"http": {
			  "enabled": false
			},
			"nfs": {
			  "enabled": false,
			  "rules": "10.0.0.1(ro)"
			},
			"provisioned": 1099511627776,
			"smb": {
			  "enabled": true
			},
			"snapshot_directory_enabled": false
		  }
		]
	  }`
	return resp[restVersion]
}

func respGetFilesystemSnapshots(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"continuation_token": "MThkMzJmYWUtZDI3OS00ODEzLWIzODYtMmM3NmFhMTMwM2My",
		"total_item_count": 1,
		"items": [
		  {
			"_as_of": 1502729489760,
			"id": "01c2889a-4124-49ff-8cbd-c33405ede123",
			"name": "example_name",
			"arrays": [
			  {
				"id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
				"name": "example_name",
				"resource_type": "example_resource_type"
			  }
			],
			"created": 1502729489760,
			"destroyed": false,
			"on": {
			  "id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			  "name": "example_name",
			  "resource_type": "example_resource_type"
			},
			"source": {
			  "id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			  "name": "example_name",
			  "resource_type": "example_resource_type"
			},
			"suffix": "checkpoint_2018_01_01"
		  }
		]
	  }`
	return resp[restVersion]
}

func respGetMetrics(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"continuation_token": "MThkMzJmYWUtZDI3OS00ODEzLWIzODYtMmM3NmFhMTMwM2My",
		"total_item_count": 1,
		"items": [
		  {
			"_as_of": 1502729489760,
			"id": "01c2889a-4124-49ff-8cbd-c33405ede123",
			"name": "example_name",
			"availabilities": [
			  {
				"aggregations": [
				  "avg"
				],
				"resolution": 30000,
				"retention": 2592000000
			  }
			],
			"description": "string",
			"resource_types": [
			  "example_resource_type"
			],
			"unit": "B/s"
		  }
		]
	  }`
	return resp[restVersion]
}

func respGetMetricsHistory(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"continuation_token": "MThkMzJmYWUtZDI3OS00ODEzLWIzODYtMmM3NmFhMTMwM2My",
		"total_item_count": 1,
		"items": [
		  {
			"_as_of": 1502729489760,
			"id": "01c2889a-4124-49ff-8cbd-c33405ede123",
			"name": "example_name",
			"aggregation": "avg",
			"data": [
			  [
				1502729489760,
				123
			  ]
			],
			"resolution": 30000,
			"resources": [
			  {
				"id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
				"name": "example_name",
				"resource_type": "example_resource_type"
			  }
			],
			"unit": "B/s"
		  }
		]
	  }`
	return resp[restVersion]
}

func respGetNetworkInterfaces(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"continuation_token": "MThkMzJmYWUtZDI3OS00ODEzLWIzODYtMmM3NmFhMTMwM2My",
		"total_item_count": 1,
		"items": [
		  {
			"_as_of": 1502729489760,
			"id": "01c2889a-4124-49ff-8cbd-c33405ede123",
			"name": "example_name",
			"arrays": [
			  {
				"id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
				"name": "example_name",
				"resource_type": "example_resource_type"
			  }
			],
			"address": "10.11.12.13",
			"enabled": true,
			"gateway": "10.20.30.40",
			"hwaddr": "90:ef:ba:80:70:dc",
			"mtu": 9000,
			"netmask": "255.255.255.0",
			"services": [
			  "iscsi"
			],
			"speed": 10000000000,
			"subinterfaces": [
			  "ct0.eth0"
			]
		  }
		]
	  }`
	return resp[restVersion]
}

func respGetPods(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"continuation_token": "MThkMzJmYWUtZDI3OS00ODEzLWIzODYtMmM3NmFhMTMwM2My",
		"total_item_count": 1,
		"items": [
		  {
			"_as_of": 1502729489760,
			"id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			"name": "example_name",
			"arrays": [
			  {
				"id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
				"name": "example_name",
				"resource_type": "example_resource_type",
				"frozen_at": 1502729489760,
				"mediator_status": "online",
				"status": "resyncing"
			  }
			],
			"mediator": "purestorage",
			"source": {
			  "id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			  "name": "example_name",
			  "resource_type": "example_resource_type"
			}
		  }
		]
	  }`
	return resp[restVersion]
}

func respGetVolumes(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"continuation_token": "MThkMzJmYWUtZDI3OS00ODEzLWIzODYtMmM3NmFhMTMwM2My",
		"total_item_count": 1,
		"items": [
		  {
			"_as_of": 1502729489760,
			"id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			"name": "example_name",
			"arrays": [
			  {
				"id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
				"name": "example_name",
				"resource_type": "example_resource_type"
			  }
			],
			"created": 1502729489760,
			"destroyed": false,
			"eradicated": false,
			"pod": {
			  "id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			  "name": "example_name",
			  "resource_type": "example_resource_type"
			},
			"provisioned": 2048576,
			"serial": "C68B5DCF2C1A4C9400012F92",
			"source": {
			  "id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			  "name": "example_name",
			  "resource_type": "example_resource_type"
			}
		  }
		]
	  }`
	return resp[restVersion]
}

func respGetVolumeSnapshots(restVersion string) string {
	resp := make(map[string]string)
	resp["1.latest"] = `{
		"continuation_token": "MThkMzJmYWUtZDI3OS00ODEzLWIzODYtMmM3NmFhMTMwM2My",
		"total_item_count": 1,
		"items": [
		  {
			"_as_of": 1502729489760,
			"id": "01c2889a-4124-49ff-8cbd-c33405ede123",
			"name": "example_name",
			"arrays": [
			  {
				"id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
				"name": "example_name",
				"resource_type": "example_resource_type"
			  }
			],
			"created": 1502729489760,
			"destroyed": false,
			"on": {
			  "id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			  "name": "example_name",
			  "resource_type": "example_resource_type"
			},
			"pod": {
			  "id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			  "name": "example_name",
			  "resource_type": "example_resource_type"
			},
			"provisioned": 2048576,
			"serial": "C68B5DCF2C1A4C9400012F92",
			"snapshot_group": {
			  "id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			  "name": "example_name",
			  "resource_type": "example_resource_type"
			},
			"source": {
			  "id": "2a989a09-c851-4d52-9ec6-ab728b1c04db",
			  "name": "example_name",
			  "resource_type": "example_resource_type"
			},
			"suffix": "checkpoint_2018_01_01"
		  }
		]
	  }`
	return resp[restVersion]
}
