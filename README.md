[![Build Status](https://travis-ci.com/devans10/pugo.svg?branch=master)](https://travis-ci.com/devans10/pugo) [![Go Report Card](https://goreportcard.com/badge/github.com/devans10/pugo)](https://goreportcard.com/report/github.com/devans10/pugo) [![codecov](https://codecov.io/gh/devans10/pugo/branch/master/graph/badge.svg)](https://codecov.io/gh/devans10/pugo) [![GoDoc](https://godoc.org/github.com/devans10/pugo?status.svg)](https://godoc.org/github.com/devans10/pugo)

# PuGo - A Go REST client for Pure Storage 
This library is designed to provide a simple interface for issuing commands to a Pure Storage FlashArray using a REST API. 
It communicates with the array using the golang http library, and returns the data into types defined within the library.

This is not designed to be a standalone program.  It is just meant to provide functions and communication within another program.

Table of Contents
=================

<!--ts-->
   * [Requirements](#Requirements)
   * [Capabilities](#Capabilities)
   * [Installation](#Installzation)
   * [Testing](#Testing)
   * [Documentation](#Documentation)
      * [FlashArray](#FlashArray)
         * [Client](#flasharray.Client)
         * [Array](#flasharray.Array)
         * [Volume](#flasharray.Volume)
      * [Pure1](#Pure1)
         * [Authentication](#Authentication)
	 * [Client](#pure1.Client)
	 * [Array](#pure1.Array)
	 
# Requirements
You should have a working Go environment setup.  If not check out the Go [getting started](http://golang.org/doc/install) guide.

# Capabilities

### FlashArray
The FlashArray library contains all functionality provided by version 1.16 of the Purity//FA REST API.

Note that different versions of the REST API offer different functionality, and some operations may be unusable except on certain 
versions of the REST API. For example, functionality relating to FlashRecover and protection groups (pgroups) requires the use of 
REST API version 1.2, which is supported only by Purity versions 4.0 and later.

### Pure1
The pure1 library contains all functionality provided by version 1.0 of the Pure1 REST API.

# Installation

### FlashArray
```sh
$ go get github.com/devans10/pugo/flasharray
```
### Pure1
```sh
$ go get github.com/devans10/pugo/pure1
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

The Pure1 tests require a connection to Pure1.  They will require environment variables set for `PURE1_APPID` and `PURE1_PRIVATEKEY`

# Documentation

## FlashArray
https://godoc.org/github.com/devans10/pugo/flasharray

### flasharray.Client

Create a client to connect to the FlashArray
```go
import (
	"fmt"
	"github.com/devans10/pugo/flasharray"
)

client := flasharray.Client{Target: "flasharray.example.com", Username: "pureuser", Password: "password", APIToken: nil, RestVersion: nil, UserAgent: nil, RequestKwargs: nil}
```

### flasharray.Array

Get the array status
```go
array, _ := client.Array.Get(nil)
fmt.Printf("Array Name: %s", array.ArrayName)
fmt.Printf("ID: %s", array.Id)
```

### flasharray.Volume

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

## Pure1
https://godoc.org/github.com/devans10/pugo/pure1

### Authentication
The Pure1 REST API authentication requires a public and private certificate.
A certificate can be created using openssl:
```sh
$ openssl genrsa -aes256 -out private.pem 2048
```

From that cert, you can create a public cert:
```sh
$ openssl rsa -in private.pem -outform PEM -pubout -out public.pem
```

Next, you will need to create a API registration on the Pure1 Manage site.  You will need to be an administrator to complete this.

1. Login to https://pure1.purestorage.com
2. Go to Administration -> API Registration
3. Click "+ Register Application"
4. Give your application a name, and paste the contents of the public.pem created above.
5. the "Application ID" that is generated will be used to authenticate the client.

### pure1.Client

To create a client connection, the Application ID, byte slice of the private key contents and the rest version need to be passed.  In the example below, I have exported the contents of the private.pem and the Application ID to envrionment variables.

Create a client to connect to Pure1
```go
import (
	"fmt"
	"os"
	"github.com/devans10/pugo/pure1"
)

appID := os.Getenv("PURE1_APPID")
privateKey := []byte(os.Getenv("PURE1_PRIVATEKEY"))
restVersion := ""

client := pure1.NewClient(appID, privateKey, restVersion)
```

### pure1.Array
Get a list of FlashArray and FlashBlade objects
```go
array, _ := client.Array.GetArray(nil)
fmt.Println("ID: ", array.ID)
fmt.Println("Name: ", array.Name)
fmt.Println("Model: ", array.Model)
fmt.Println("OS: ", array.OS)
fmt.Println("Version: ", array.Version)
fmt.Println("AsOf: ", array.AsOf)
```



