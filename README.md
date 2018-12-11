# Pure Storage REST client for Go
This library is designed to provide a simple interface for issuing commands to a Pure Storage Flash Array using a REST API. 
It communicates with the array using the golang http library, and returns the data into types defined within the library.

This is not designed to be a standalone program.  It is just meant to provide functions and communication within another progrom.

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

# Documentation
https://godoc.org/github.com/devans10/go-purestorage/flasharray

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
