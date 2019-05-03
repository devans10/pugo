## 1.0.0
Changed Project name to Pugo
Changed License to Apache 2.0

## 0.3.0
IMPROVEMENTS:
* Added Library for the Pure1 API

## 0.2.0

IMPROVEMENTS:
* Added PodService
* Added AlertService
* Added MessageService
* Added SnmpService(#4)
* Added CertService(#2)
* Added SmtpService(#3)

FIXES:
* Fixed the Enable/Disable RemoteAssist functions

## 0.1.4

IMPROVEMENTS:
* Added data arg to Hostgroups.ConnectHostgroup to allow setting of LUN

## 0.1.3

IMPROVEMENTS:
* Added CHAP, Personality, and Preferred Array properties to Host Struct.
* Added params to ListHosts and GetHost to be able to query for above properties.
* Added data arg to Hosts.ConnectHost to allow setting of LUN.
* Added params to Host.ListHostConnections to allow returning private or shared connections.

## 0.1.2

Because I messed up my tags

## 0.1.1 (December 30, 2018)

Fixes:
* Changed ProtectionGroup Targets type to []map[sting]interface{}. fixes #1

Notes:
* Added params to the "Get" functions to allow query parameters to be added.

## 0.1.0 (December 29,2018)

* Initial Release
