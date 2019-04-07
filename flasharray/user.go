// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
)

// UserService struct for user API endpoints
type UserService struct {
	client *Client
}

// listUsers is the private function for returning dictionaries
// which describes remote access
func (n *UserService) listUsers(data interface{}) ([]User, error) {

	req, _ := n.client.NewRequest("GET", "admin", nil, data)
	m := []User{}
	_, err := n.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// ListAdmins lists attributes for Admins
func (n *UserService) ListAdmins() ([]User, error) {

	m, err := n.listUsers(nil)
	if err != nil {
		return nil, err
	}

	return m, err
}

// CreateAdmin creates an Admin
func (n *UserService) CreateAdmin(name string) (*User, error) {

	path := fmt.Sprintf("admin/%s", name)
	req, _ := n.client.NewRequest("POST", path, nil, nil)
	m := &User{}
	_, err := n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// DeleteAdmin deletes an Admin
func (n *UserService) DeleteAdmin(name string) (*User, error) {

	path := fmt.Sprintf("admin/%s", name)
	req, _ := n.client.NewRequest("DELETE", path, nil, nil)
	m := &User{}
	_, err := n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// SetAdmin modifies Admin Attributes
func (n *UserService) SetAdmin(name string, data interface{}) (*User, error) {

	path := fmt.Sprintf("admin/%s", name)
	req, _ := n.client.NewRequest("PUT", path, nil, data)
	m := &User{}
	_, err := n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// GetAdmin lists attributes for specified Admin
func (n *UserService) GetAdmin(name string) (*User, error) {

	path := fmt.Sprintf("admin/%s", name)
	req, _ := n.client.NewRequest("GET", path, nil, nil)
	m := &User{}
	_, err := n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// CreateAPIToken creates an API Token
func (n *UserService) CreateAPIToken(name string) (*Token, error) {

	path := fmt.Sprintf("admin/%s/apitoken", name)
	req, _ := n.client.NewRequest("POST", path, nil, nil)
	m := &Token{}
	_, err := n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// DeleteAPIToken deletes an  API Token
func (n *UserService) DeleteAPIToken(name string) (*Token, error) {

	path := fmt.Sprintf("admin/%s/apitoken", name)
	req, _ := n.client.NewRequest("DELETE", path, nil, nil)
	m := &Token{}
	_, err := n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// ListPublicKeys returns a list of public keys
func (n *UserService) ListPublicKeys() ([]User, error) {

	data := map[string]bool{"publickey": true}
	m, err := n.listUsers(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// ListAPITokens returns a list of API Tokens
func (n *UserService) ListAPITokens() ([]User, error) {

	data := map[string]bool{"api_token": true}
	m, err := n.listUsers(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// RefreshAdmin refreshes the admin permission cache for the specified admin
func (n *UserService) RefreshAdmin(name string) (*User, error) {

	data := map[string]string{"action": "refresh"}
	m, err := n.SetAdmin(name, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// RefreshAdmins clear the admin permission cache.
func (n *UserService) RefreshAdmins() (*User, error) {
	data := make(map[string]interface{})
	data["action"] = "refresh"
	data["clear"] = true
	req, _ := n.client.NewRequest("PUT", "admin", nil, data)
	m := &User{}
	_, err := n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// SetPublicKey modifies public key for the specified admin
func (n *UserService) SetPublicKey(name string, key string) (*User, error) {

	data := map[string]string{"publickey": key}
	m, err := n.SetAdmin(name, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// SetPassword sets the password for the specified admin
func (n *UserService) SetPassword(name string, newPassword string, oldPassword string) (*User, error) {

	data := map[string]string{"password": newPassword, "old_password": oldPassword}
	m, err := n.SetAdmin(name, data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetGlobalAdminAttr returns a map describing the existing global admin attributes
func (n *UserService) GetGlobalAdminAttr() (*GlobalAdmin, error) {

	req, _ := n.client.NewRequest("GET", "admin/settings", nil, nil)
	m := &GlobalAdmin{}
	_, err := n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// SetGlobalAdminAttr modifies global admin attributes
func (n *UserService) SetGlobalAdminAttr(data interface{}) (*GlobalAdmin, error) {

	req, _ := n.client.NewRequest("PUT", "admin/settings", nil, data)
	m := &GlobalAdmin{}
	_, err := n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// ListAdminUser return a map describing lockout information for locked out admins
func (n *UserService) ListAdminUser() ([]User, error) {

	data := map[string]bool{"lockout": true}
	m, err := n.listUsers(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// GetAdminUser return a map describing lockout information specified admins
func (n *UserService) GetAdminUser(name string) (*User, error) {

	path := fmt.Sprintf("admin/%s", name)
	data := map[string]bool{"lockout": true}
	req, _ := n.client.NewRequest("GET", path, nil, data)
	m := &User{}
	_, err := n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// UnlockAdmin unlocks an admin
func (n *UserService) UnlockAdmin(name string) (*User, error) {

	path := fmt.Sprintf("admin/%s/lockout", name)
	req, _ := n.client.NewRequest("GET", path, nil, nil)
	m := &User{}
	_, err := n.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, nil
}
