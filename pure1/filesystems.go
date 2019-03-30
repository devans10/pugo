// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// FilesystemService type creates a service to expose array endpoints
type FilesystemService struct {
	client *Client
}

// GetFilesystems returns a list of Filesystem objects
func (f *FilesystemService) GetFilesystems(params map[string]string) ([]Filesystem, error) {
	req, err := f.client.NewRequest("GET", "file-systems", params, nil)
	if err != nil {
		return nil, err
	}

	m := []Filesystem{}
	_, err = f.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
