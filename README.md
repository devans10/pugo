[![Build Status](https://travis-ci.com/devans10/go-purestorage.svg?branch=master)](https://travis-ci.com/devans10/go-purestorage)

# Pure Storage REST client for Go
This library is designed to provide a simple interface for issuing commands to a Pure Storage Flash Array using a REST API. 
It communicates with the array using the golang http library, and returns the data into types defined within the library.

This is not designed to be a standalone program.  It is just meant to provide functions and communication within another progrom.

Table of Contents
=================

<!--ts-->
   * [Requirements](#Requirements)
   * [Capabilities](#Capabilities)
   * [Installation](#Installzation)
   * [Testing](#Testing)
   * [Documentation](#Documentation)
      * [FlashArray](#FlashArray)
         * [Client](#Client)
         * [Array](#Array)
         * [Volume](#Volume)

# Requirements
You should have a working Go environment setup.  If not check out the Go [getting started](http://golang.org/doc/install) guide.

# Capabilities
Currently, the library contains all functionality provided by version 1.16 of the REST API for Volumes, Hosts, Host Groups, and
Protection Groups.  Additional functionality will be added as I get time.

Note that different versions of the REST API offer different functionality, and some operations may be unusable except on certain 
versions of the REST API. For example, functionality relating to FlashRecover and protection groups (pgroups) requires the use of 
REST API version 1.2, which is supported only by Purity versions 4.0 and later.

# Installation
```sh
$ go get github.com/devans10/go-purestorage/flasharray
```


# Testing

Run unit tests
```sh
$ make test
```
To Run Acceptance tests
```sh
$ make testacc
```
These tests require a connection to a Pure FlashArray.  They will require environment variables are set for `PURE_TARGET` and `PURE_APITOKEN` or `PURE_USERNAME` and `PURE_PASSWORD`

# Documentation

## FlashArray
https://godoc.org/github.com/devans10/go-purestorage/flasharray

### Client

Create a client to connect to the FlashArray
```go
import (
	"fmt"
	"github.com/devans10/go-purestorage/flasharray"
)

client := flasharray.Client{Target: "flasharray.example.com", Username: "pureuser", Password: "password", Api_token: nil, Rest_version: nil, User_agent: nil, Request_kwargs: nil}
```

### Array

Get the array status
```go
array, _ := client.Array.Get(nil)
fmt.Printf("Array Name: %s", array.Array_name)
fmt.Printf("ID: %s", array.Id)
```

### Volume

Create a new volume
```go
volume, _ := client.Volumes.CreateVolume("testvol", 1024000000)
fmt.Printf("Name: %s, Size: %d", volume.Name, volume.Size)
```

Clone a volume
```go
volume, _ = client.Volumes.CopyVolume("testclone", "testvol")
fmt.Printf("Name: %s, Source: %d", volume.Name, volume.Source)
```

Get a volume
```go
volume, _ = client.Volumes.GetVolume("testclone", nil)
fmt.Printf("Name: %s, Size: %d", volume.Name, volume.Size)
```

Create a snapshot
```go
snapshot, _ := client.Volumes.CreateSnapshot("testvolume", "test")
```

List Volumes
```go
for _, vol := range client.Volumes.ListVolumes(nil) {
	fmt.Printf("Volume Name: %s", vol.Name
}
```



