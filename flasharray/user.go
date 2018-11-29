// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

type UserService struct {
	client *Client
}

// listUsers is the private function for returning dictionaries
// which describes remote access
func (n *UserService) listUsers(data interface{}) ([]User, error) {

	req, err := n.client.NewRequest("GET", "admin", nil, data)
        if err != nil {
                return nil, err
        }

        m := []User{}
        _, err = n.client.Do(req, &m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// List Admins
func (n *UserService) ListAdmins() ([]User, error) {

        m, err := n.listUsers(nil)
        if err != nil {
                return nil, err
        }

	return m, err
}

// Create Admin
func (n *UserService) CreateAdmin(name string) (*User, error) {

	path := fmt.Sprintf("admin/%s", name)
        req, err := n.client.NewRequest("POST", path, nil, nil)
        if err != nil {
                return nil, err
        }

        m := &User{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Delete Admin
func (n *UserService) DeleteAdmin(name string) (*User, error) {

	path := fmt.Sprintf("admin/%s", name)
        req, err := n.client.NewRequest("DELETE", path, nil, nil)
        if err != nil {
                return nil, err
        }

        m := &User{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Set Admin Attribute
func (n *UserService) SetAdmin(name string, params map[string]string, data interface{}) (*User, error) {

	path := fmt.Sprintf("admin/%s", name)
        req, err := n.client.NewRequest("PUT", path, params, data)
        if err != nil {
                return nil, err
        }

        m := &User{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}


// Get Admin
func (n *UserService) GetAdmin(name string) (*User, error) {

        path := fmt.Sprintf("admin/%s", name)
        req, err := n.client.NewRequest("GET", path, nil, nil)
        if err != nil {
                return nil, err
        }

        m := &User{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Create API Token
func (n *UserService) CreateApiToken(name string) (*ApiToken, error) {

        path := fmt.Sprintf("admin/%s/apitoken", name)
        req, err := n.client.NewRequest("POST", path, nil, nil)
        if err != nil {
                return nil, err
        }

        m := &ApiToken{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Delete API Token
func (n *UserService) DeleteApiToken(name string) (*ApiToken, error) {

        path := fmt.Sprintf("admin/%s/apitoken", name)
        req, err := n.client.NewRequest("DELETE", path, nil, nil)
        if err != nil {
                return nil, err
        }

        m := &ApiToken{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Return a list of public keys
func (n *UserService) ListPublicKeys() ([]User, error) {

	data := map[string]bool{"publickey": true}
        m, err := n.listUsers(data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Return a list of API Tokens
func (n *UserService) ListApiTokens() ([]User, error) {

        data := map[string]bool{"api_token": true}
        m, err := n.listUsers(data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Refresh the admin permission cache for the specified admin
func (n *UserService) RefreshAdmin(name string, params map[string]string) (*User, error) {

        data := map[string]string{"action": "refresh"}
        m, err := n.SetAdmin(name, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Clear the admin permission cache.
func (n *UserService) RefreshAdmins() (*User, error) {

	type data struct {
		Action	string	`json:"action"`
		Clear	bool	`json:"clear"`
	}
	d := &data{}
	d.Action = "refresh"
	d.Clear = true
        req, err := n.client.NewRequest("PUT", "admin", nil, d)
        if err != nil {
                return nil, err
        }

        m := &User{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Set public key for the specified admin
func (n *UserService) SetPublicKey(name string, key string, params map[string]string) (*User, error) {

        data := map[string]string{"publickey": key}
        m, err := n.SetAdmin(name, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Set the password for the specified admin
func (n *UserService) SetPassword(name string, new_password string, old_password string, params map[string]string) (*User, error) {

	data := map[string]string{"password": new_password, "old_password": old_password}
        m, err := n.SetAdmin(name, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Return a map describing the existing global admin attributes
func (n *UserService) GetGlobalAdminAttr() (*GlobalAdmin, error) {

        req, err := n.client.NewRequest("GET", "admin/settings", nil, nil)
        if err != nil {
                return nil, err
        }

        m := &GlobalAdmin{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Set global admin attributes
func (n *UserService) SetGlobalAdminAttr(data interface{}) (*GlobalAdmin, error) {

        req, err := n.client.NewRequest("PUT", "admin/settings", nil, data)
        if err != nil {
                return nil, err
        }

        m := &GlobalAdmin{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Return a map describing lockout information for locked out admins
func (n *UserService) ListAdminUser() ([]User, error) {

	data := map[string]bool{"lockout": true}
	m, err := n.listUsers(data)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Return a map describing lockout information specified admins
func (n *UserService) GetAdminUser(name string) (*User, error) {

	path := fmt.Sprintf("admin/%s", name)
	data := map[string]bool{"lockout": true}
        req, err := n.client.NewRequest("GET", path, nil, data)
        if err != nil {
                return nil, err
        }

	m := &User{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Unlocks an admin
func (n *UserService) UnlockAdmin(name string) (*User, error) {

        path := fmt.Sprintf("admin/%s/lockout", name)
        req, err := n.client.NewRequest("GET", path, nil, nil)
        if err != nil {
                return nil, err
        }

        m := &User{}
        _, err = n.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

