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

// Package pure1 is a library used to interface with the
// Pure1 Manage REST API
package pure1

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// Client struct represents the Pure1 API and exposes its endpoints
type Client struct {
	AppID       string
	PrivateKey  []byte
	RestVersion string

	client *http.Client
	token  *pure1Token

	Arrays              *ArrayService
	Filesystems         *FilesystemService
	FilesystemSnapshots *FilesystemSnapshotService
	Metrics             *MetricsService
	NetworkInterfaces   *NetworkInterfacesService
	Pods                *PodService
	Volumes             *VolumeService
	VolumeSnapshots     *VolumeSnapshotService
}

type pure1Response struct {
	TotalItems        int           `json:"total_item_count,omitempty"`
	ContinuationToken interface{}   `json:"continuation_token,omitempty"`
	Items             []interface{} `json:"items,omitempty"`
}

type pure1Token struct {
	AccessToken     string `json:"access_token,omitempty"`
	IssuedTokenType string `json:"issued_token_type,omitempty"`
	TokenType       string `json:"token_type,omitempty"`
	ExpiresIn       int    `json:"expires_in,omitempty"`
}

// NewClient creates a Client struct for calling Pure1 API endpoints
func NewClient(appID string, privateKey []byte, restVersion string) (*Client, error) {

	if appID == "" {
		err := errors.New("[error] Must specify an App ID")
		return nil, err
	}

	if privateKey == nil {
		err := errors.New("[error] Must specify a Private Key")
		return nil, err
	}

	if restVersion == "" {
		restVersion = "1.0"
	}

	c := &Client{AppID: appID, PrivateKey: privateKey, RestVersion: restVersion}
	c.client = &http.Client{}
	token, err := getToken(c)
	if err != nil {
		return nil, err
	}

	c.token = token

	c.Arrays = &ArrayService{client: c}
	c.Filesystems = &FilesystemService{client: c}
	c.FilesystemSnapshots = &FilesystemSnapshotService{client: c}
	c.Metrics = &MetricsService{client: c}
	c.NetworkInterfaces = &NetworkInterfacesService{client: c}
	c.Pods = &PodService{client: c}
	c.Volumes = &VolumeService{client: c}
	c.VolumeSnapshots = &VolumeSnapshotService{client: c}

	return c, nil
}

func getToken(c *Client) (*pure1Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": c.AppID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Unix() + 60,
	})

	key, err := jwt.ParseRSAPrivateKeyFromPEM(c.PrivateKey)
	if err != nil {
		return nil, err
	}

	// Generate encoded token and send it as response.
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Add("grant_type", "urn:ietf:params:oauth:grant-type:token-exchange")
	values.Add("subject_token", tokenStr)
	values.Add("subject_token_type", "urn:ietf:params:oauth:token-type:jwt")

	pure1URL := fmt.Sprintf("https://api.pure1.purestorage.com/oauth2/%s/token", "1.0")
	req, err := http.NewRequest(http.MethodPost, pure1URL, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := validateResponse(resp); err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pToken := &pure1Token{}
	if err := json.Unmarshal(responseBody, pToken); err != nil {
		return nil, err
	}

	// log.Printf("[DEBUG] Access Token: %v", pToken)

	return pToken, nil
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

	baseURL, err := url.Parse(fpath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		ps := url.Values{}
		for k, v := range params {
			//log.Printf("[DEBUG] key: %s, value: %s \n", v, k)
			ps.Set(k, v)
		}
		baseURL.RawQuery = ps.Encode()
	}
	req, err := http.NewRequest(method, baseURL.String(), nil)
	if err != nil {
		return nil, err
	}
	if data != nil {
		jsonString, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, baseURL.String(), bytes.NewBuffer(jsonString))
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", `Bearer `+c.token.AccessToken)

	return req, err
}

// Do is the client function that performs the HTTP request.
// req  The HTTP request object to be executed.
// v    The data object that will be populated and returned. i.e. Volume struct
// reestablish_session  A bool that states if the session should be reestablished prior to execution.
//                      This functionality is NOT implemented yet.  By default the Go HTTP library
//                      does not set a timeout, I need to set this implicitly.
//                      However, the array will timeout the session after 30 minutes.
func (c *Client) Do(req *http.Request, v interface{}, reestablishSession bool) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println("Do request failed")
		return nil, err
	}
	defer resp.Body.Close()
	//log.Printf("[debug] URL: %s ", req.URL.String())
	//log.Printf("[debug] Response code: %v", resp.Status)

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

// formatPath returns the formated string to be used for the base URL in
// all API calls
func (c *Client) formatPath(path string) string {
	return fmt.Sprintf("https://api.pure1.purestorage.com/api/%s/%s", c.RestVersion, path)
}
