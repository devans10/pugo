// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package pure1

// FilesystemSnapshotService type creates a service to expose array endpoints
type FilesystemSnapshotService struct {
	client *Client
}

// GetFilesystemSnapshots returns a list of Filesystem objects
func (f *FilesystemSnapshotService) GetFilesystemSnapshots(params map[string]string) ([]FilesystemSnapshot, error) {
	req, err := f.client.NewRequest("GET", "file-system-snapshots", params, nil)
	if err != nil {
		return nil, err
	}

	m := []FilesystemSnapshot{}
	_, err = f.client.Do(req, &m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}
