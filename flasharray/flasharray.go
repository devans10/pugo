package flasharray

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var library_supported_versions = [...]string {"1.0","1.1","1.2","1.3","1.4","1.5","1.6","1.7","1.8","1.9","1.10","1.11","1.12","1.13","1.14","1.15","1.16"}

type Client struct {
	Target		string
	Username	string
	Password	string
	Api_token	string
	Rest_version	string
	User_agent	string
	Request_kwargs	map[string]string

	client		*http.Client

	Array		*ArrayService
	//Vols            *VolService
        //Hosts           *HostService
        //Hostgroups      *HostgroupService
}

type supported_versions struct {
        versions        []string
}

type auth struct {
        token           string  `json:"api_token,omitempty"`
}

func NewClient(target string, username string, password string, api_token string,
               rest_version string, verify_https bool, ssl_cert string,
               user_agent string, request_kwargs map[string]string) (*Client, error) {

	if api_token == "" && (username == "" && password == "") {
		err := errors.New("[ERROR] Must specify API token or both username and password.")
		return nil, err
	}

	if api_token != "" && (username != "" && password != "") {
		err := errors.New("Specify only API token or both username and password.")
	}

	if request_kwargs == nil {
		request_kwargs = make(map[string]string)
	}

	elem, ok := request_kwargs["verify"]
	if !ok {
		if ssl_cert && verify_https {
			request_kwargs["verify"] = ssl_cert
		} else {
			request_kwargs["verify"] = verify_https
		}
	}

	if rest_version != "" {
		err := checkRestVersion(rest_version, target)
		if err != nil {
			return err
		}
	} else {
		rest_verion, err := chooseRestVersion(target)
		if err != nil {
			return err
		}
	}

	cookieJar, _ := cookiejar.New(nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: verify_https},
	}
	c := &Client{Target: target, Username: username, Password: password, Api_token: api_token, Rest_version: rest_version, Request_kwargs}
	c.client = &http.Client{Transport: tr, Jar: cookieJar}

	if api_token == "" {
                t := &auth{}
                tokenUrl := c.formatPath("auth/apitoken")
                c.getApiToken(tokenUrl, t)
                c.SetToken(t.token)
        }

        authUrl := c.formatPath("auth/session")
        data := map[string]string{"api_token": c.Api_token}
        jsonValue, _ := json.Marshal(data)
        _, err := c.client.Post(authUrl, "application/json", bytes.NewBuffer(jsonValue))
        if err != nil {
                return nil, err
        }

	c.Array = &ArrayService{client: c}
	//c.Vols = &VolService{client: c}
        //c.Hosts = &HostService{client: c}
        //c.Hostgroups = &HostgroupService{client: c}

        return c, err
}

func (pc *Client) NewRequest(method string, path string, params map[string]string, data interface{}) (*http.Request, error) {

	if strings.HasPrefix(path, "http") {
                baseUrl, err := url.Parse(path)
                if err != nil {
                        return nil, err
                }
        } else {
                baseUrl, err := url.Parse(c.formatPath(path))
                if err != nil {
                        return nil, err
                }
        }

        if params != nil {
                ps := url.Values{}
                for k, v := range params {
                        log.Printf("[DEBUG] key: %s, value: %s \n", v, k)
                        ps.Set(k, v)
                }
                baseUrl.RawQuery = ps.Encode()
        }
        req, err := http.NewRequest(method, baseUrl.String(), nil)
        if data != nil {
                jsonString, _ := json.Marshal(data)
                req, err = http.NewRequest(method, baseUrl.String(), bytes.NewBuffer(jsonString))
        }

        req.Header.Add("content-type", "application/json; charset=utf-8")
        req.Header.Add("Accept", "application/json")
        req.Header.Set("Content-Type", "application/json")
	if c.User_agent != "" {
                req.Header.Add("User-Agent", c.User_agent)
        }

        return req, err
}

func (pc *Client) Do(req *http.Request, v interface{}, reestablish_session bool) (*http.Response, error) {
        resp, err := pc.client.Do(req)
        if err != nil {
                fmt.Println("Do request failed")
                return nil, err
        }
        defer resp.Body.Close()

        log.Printf("[INFO] Response code: %v", resp.Status)

        if err := validateResponse(resp); err != nil {
                return resp, err
        }

        err = decodeResponse(resp, v)
        return resp, err

}
func decodeResponse(r *http.Response, v interface{}) error {
        if v == nil {
                return fmt.Errorf("nil interface provided to decodeResponse")
        }

        bodyBytes, _ := ioutil.ReadAll(r.Body)
        bodyString := string(bodyBytes)
        err := json.Unmarshal([]byte(bodyString), &v)
        return err
}

func validateResponse(r *http.Response) error {
        if c := r.StatusCode; 200 <= c && c <= 299 {
                return nil
        }

        bodyBytes, _ := ioutil.ReadAll(r.Body)
        bodyString := string(bodyBytes)
        return fmt.Errorf("Response code: %d, ResponeBody: %s", r.StatusCode, bodyString)
}


func checkRestVersion(v string, t string) (bool, error) {

	checkUrl, err := url.Parse("https://" + t + "/api/api_version")
	if err != nil {
		return err
	}
	r, err := http.Get(checkUrl.String())
	if err != nil {
		return err
	}

	s := supported_versions{}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(s)

	var array_supported bool
	for _, n := range s.versions {
		if v == n {
			array_supported := true
		}
	}
	if !array_supported {
		err := errors.New("[ERROR] Array is incompatible with REST API version " + v)
		return err
	}

	var library_supported bool
        for _, n := range library_supported_versions {
                if v == n {
                        library_supported := true
                }
        }
        if !library_supported {
                err := errors.New("[ERROR] Library is incompatible with REST API version " + v)
                return err
        }
	return nil
}

func chooseRestVersion(t string) (string, error) {

	checkUrl, err := url.Parse("https://" + t + "/api/api_version")
        if err != nil {
                return err
        }
        r, err := http.Get(checkUrl.String())
        if err != nil {
                return err
        }

        s := supported_versions{}
        defer r.Body.Close()

        json.NewDecoder(r.Body).Decode(s)

	for i := len(library_supported_versions)-1; i >= 0; i-- {
		for _, n := range s.versions {
			if library_supported_versions[i] == n {
				return n, nil
			}
		}
        }
        err = errors.New("[ERROR] Array is incompatible with all supported REST API versions")
        return nil, err
}

func (c *Client) getApiToken(uri string, token interface{}) error {
        authUrl, err := url.Parse(uri)
        if err != nil {
                return err
        }
        data := map[string]string{"username": c.User, "password": c.Password}
        jsonValue, _ := json.Marshal(data)
        req, err := http.NewRequest("POST", authUrl.String(), bytes.NewBuffer(jsonValue))
        resp, err := c.client.Do(req)
        if err != nil {
                return err
        }
        defer req.Body.Close()

        return json.NewDecoder(resp.Body).Decode(token)
}

func (c *Client) SetToken(token string) {
        c.ApiToken = token
}

func (c *Client) formatPath(path) string {
	return fmt.Sprintf("https://%s/api/%s/%s", c.Target, c.Rest_version, path)
}
