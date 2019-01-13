// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package flasharray is designed to provide a simple interface for
// issuing commands to a Pure Storage Flash Array using a REST API.
// It communicates with the array using the golang http library,
// and returns the data into types defined within the library.
// This is not designed to be a standalone program.
// It is just meant to provide functions and communication within another program
package flasharray

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

// supported_rest_versions is used to negotiate the API version to use
var supported_rest_versions = [...]string{"1.0", "1.1", "1.2", "1.3", "1.4", "1.5", "1.6", "1.7", "1.8", "1.9", "1.10", "1.11", "1.12", "1.13", "1.14", "1.15", "1.16"}

// Type Client represents a Pure Storage FlashArray and exposes administrative APIs.
type Client struct {
	Target         string
	Username       string
	Password       string
	Api_token      string
	Rest_version   string
	User_agent     string
	Request_kwargs map[string]string

	client *http.Client

	Array            *ArrayService
	Volumes          *VolumeService
	Hosts            *HostService
	Hostgroups       *HostgroupService
	Offloads         *OffloadService
	Protectiongroups *ProtectiongroupService
	Vgroups          *VgroupService
	Networks         *NetworkService
	Hardware         *HardwareService
	Users            *UserService
	Dirsrv           *DirsrvService
	Pods             *PodService
	Alerts           *AlertService
	Messages         *MessageService
	Snmp             *SnmpService
	Cert             *CertService
}

// Type supported is used for retrieving the support API versions from the Flash Array
type supported struct {
	Versions []string `json:"version"`
}

// Type auth is used to for the API token used in API authentication
type auth struct {
	Token string `json:"api_token,omitempty"`
}

// NewClient returns a Client struct used to call the administrative functions.
//
// Parameters:
// target
// IP address or domain name of the target array's management interface.
//
// username
// Username to connect to the array
//
// password
// Password used to connect to the array
//
// api_token
// API token used to connect to the array
//
// The API Token is always used to connect to the REST API.  If username and password
// are provided, then they are used to retrieve the API token for that user before
// the HTTP session is started.  Either api_token or username and password are
// required. If neither or both are provided, then an error is returned.
//
// rest_version
// The REST API version to use for the the session.  If not provied,
// the version will be negotiated between the library and the array.
//
// verify_https
// A bool used to set whether SSL host verification should be performed.
//
// ssl_cert
// Path to SSL certificate or CA Bundle file. Ignored if verify_https=False.
//
// user_agent
// String to be used as the HTTP User-Agent for requests.
//
// request_kwargs
// A map of keyword arguments that we will pass into the the call.
func NewClient(target string, username string, password string, api_token string,
	rest_version string, verify_https bool, ssl_cert bool,
	user_agent string, request_kwargs map[string]string) (*Client, error) {

	//log.Printf("[DEBUG] flasharray.NewClient: checking auth paramters")
	if api_token == "" && (username == "" && password == "") {
		err := errors.New("[ERROR] Must specify API token or both username and password.")
		return nil, err
	}

	if api_token != "" && (username != "" && password != "") {
		err := errors.New("Specify only API token or both username and password.")
		return nil, err
	}

	//log.Printf("[DEBUG] flasharray.NewClient: checking request_kwargs")
	if request_kwargs == nil {
		request_kwargs = make(map[string]string)
	}

	_, ok := request_kwargs["verify"]
	if !ok {
		if ssl_cert && verify_https {
			request_kwargs["verify"] = "false"
		} else {
			request_kwargs["verify"] = "true"
		}
	}

	//log.Printf("[DEBUG] flasharray.NewClient: checking rest_version")
	if rest_version != "" {
		err := checkRestVersion(rest_version, target)
		if err != nil {
			return nil, err
		}
	} else {
		r, err := chooseRestVersion(target)
		if err != nil {
			return nil, err
		}
		rest_version = r
	}

	//log.Printf("[DEBUG] flasharray.NewClient: Rest Vesrion: %s", rest_version)

	//log.Printf("[DEBUG] flasharray.NewClient: creating client")
	cookieJar, _ := cookiejar.New(nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &Client{Target: target, Username: username, Password: password, Api_token: api_token, Rest_version: rest_version, Request_kwargs: request_kwargs}
	c.client = &http.Client{Transport: tr, Jar: cookieJar}

	//log.Printf("[DEBUG] flasharray.NewClient: Authenticating REST session")
	if api_token == "" {
		c.getApiToken()
	}

	authUrl := c.formatPath("auth/session")
	data := map[string]string{"api_token": c.Api_token}
	jsonValue, _ := json.Marshal(data)
	_, err := c.client.Post(authUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	//log.Printf("[DEBUG] flasharray.NewClient: REST session created.")

	c.Array = &ArrayService{client: c}
	c.Volumes = &VolumeService{client: c}
	c.Hosts = &HostService{client: c}
	c.Hostgroups = &HostgroupService{client: c}
	c.Offloads = &OffloadService{client: c}
	c.Protectiongroups = &ProtectiongroupService{client: c}
	c.Vgroups = &VgroupService{client: c}
	c.Networks = &NetworkService{client: c}
	c.Hardware = &HardwareService{client: c}
	c.Users = &UserService{client: c}
	c.Dirsrv = &DirsrvService{client: c}
	c.Pods = &PodService{client: c}
	c.Alerts = &AlertService{client: c}
	c.Messages = &MessageService{client: c}
	c.Snmp = &SnmpService{client: c}
	c.Cert = &CertService{client: c}

	return c, err
}

// NewRequest builds and returns a new HTTP request object.
//
// Parameters:
// method
// This is the HTTP method to be used, i.e. GET, PUT, POST, or DELETE
//
// path
// String of the API URI path to be called.
//
// params
// A map of key value pairs that will be added to the query string of the URL
//
// data
// The data body to be passed in the HTTP request. This will be converted to JSON,
// then added to the request as bytes.
//
func (c *Client) NewRequest(method string, path string, params map[string]string, data interface{}) (*http.Request, error) {

	var fpath string
	if strings.HasPrefix(path, "http") {
		fpath = path
	} else {
		fpath = c.formatPath(path)
	}

	baseUrl, err := url.Parse(fpath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		ps := url.Values{}
		for k, v := range params {
			//log.Printf("[DEBUG] key: %s, value: %s \n", v, k)
			ps.Set(k, v)
		}
		baseUrl.RawQuery = ps.Encode()
	}
	req, err := http.NewRequest(method, baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	if data != nil {
		jsonString, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, baseUrl.String(), bytes.NewBuffer(jsonString))
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if c.User_agent != "" {
		req.Header.Add("User-Agent", c.User_agent)
	}

	return req, err
}

// Do is the client function that performs the HTTP request.
// req	The HTTP request object to be executed.
// v	The data object that will be populated and returned. i.e. Volume struct
// reestablish_session	A bool that states if the session should be reestablished prior to execution.
//			This functionality is NOT implemented yet.  By default the Go HTTP library
//			does not set a timeout, I need to set this implicitly.
//			However, the array will timeout the session after 30 minutes.
func (pc *Client) Do(req *http.Request, v interface{}, reestablish_session bool) (*http.Response, error) {
	resp, err := pc.client.Do(req)
	if err != nil {
		fmt.Println("Do request failed")
		return nil, err
	}
	defer resp.Body.Close()

	//log.Printf("[INFO] Response code: %v", resp.Status)

	if err := validateResponse(resp); err != nil {
		return resp, err
	}

	err = decodeResponse(resp, v)
	return resp, err

}

// decodeResponse function reads the http response body into an interface.
func decodeResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return fmt.Errorf("nil interface provided to decodeResponse")
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	err := json.Unmarshal([]byte(bodyString), &v)
	return err
}

// validateResponse checks that the http response is within the 200 range.
// Some functionality needs to be added here to check for some specific errors,
// and probably add the equivlents to PureError and PureHTTPError from the Python
// REST client.
func validateResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	return fmt.Errorf("Response code: %d, ResponeBody: %s", r.StatusCode, bodyString)
}

// checkRestVersion will check that the specified rest_version is supported
// by the Flash Array, and the library.
func checkRestVersion(v string, t string) error {

	checkUrl, err := url.Parse("https://" + t + "/api/api_version")
	if err != nil {
		return err
	}
	s := &supported{}
	err = getJson(checkUrl.String(), s)

	var array_supported bool
	for _, n := range s.Versions {
		if v == n {
			array_supported = true
		}
	}
	if !array_supported {
		err := errors.New("[ERROR] Array is incompatible with REST API version " + v)
		return err
	}

	var library_supported bool
	for _, n := range supported_rest_versions {
		if v == n {
			library_supported = true
		}
	}
	if !library_supported {
		err := errors.New("[ERROR] Library is incompatible with REST API version " + v)
		return err
	}
	return nil
}

// chooseRestVersion will negotiate the highest REST API version supported by
// the library and the flash array
func chooseRestVersion(t string) (string, error) {

	checkUrl, err := url.Parse("https://" + t + "/api/api_version")
	if err != nil {
		return "", err
	}
	s := &supported{}
	err = getJson(checkUrl.String(), s)
	if err != nil {
		return "", err
	}

	for i := len(supported_rest_versions) - 1; i >= 0; i-- {
		for n := len(s.Versions) - 1; n >= 0; n-- {
			if supported_rest_versions[i] == s.Versions[n] {
				return s.Versions[n], nil
			}
		}
	}
	err = errors.New("[ERROR] Array is incompatible with all supported REST API versions")
	return "", err
}

// getApiToken retrieved the API token for the given user.  The API token
// is then used for all http authentication.
func (c *Client) getApiToken() error {

	authUrl, err := url.Parse(c.formatPath("auth/apitoken"))
	if err != nil {
		return err
	}

	data := map[string]string{"username": c.Username, "password": c.Password}
	jsonValue, _ := json.Marshal(data)
	fmt.Println(bytes.NewBuffer(jsonValue))
	req, err := http.NewRequest("POST", authUrl.String(), bytes.NewBuffer(jsonValue))
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	r, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	t := &auth{}
	err = json.NewDecoder(r.Body).Decode(t)
	fmt.Println(t)
	c.Api_token = t.Token

	return err
}

// formatPath returns the formated string to be used for the base URL in
// all API calls
func (c *Client) formatPath(path string) string {
	return fmt.Sprintf("https://%s/api/%s/%s", c.Target, c.Rest_version, path)
}

// getJson is just a helper function that creates and retrieves information
// from the flash array before the actual session is established.
// Right now, its just grabbing the supported API versions.  I should
// probably find a more graceful way to accomplish this.
func getJson(uri string, target interface{}) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var c = &http.Client{Timeout: 10 * time.Second, Transport: tr}
	r, err := c.Get(uri)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
