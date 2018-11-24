package flasharray

import (
	"fmt"
)

type OffloadService struct {
	client *Client
}

func (o *OffloadService) ConnectNFSOffload(name string, address string, mount_point string, params map[string]string) (*NFSOffload, error) {

	data := map[string]string{"name": name, "address": address, "mount_point": mount_point}
        path := fmt.Sprintf("nfs_offload/%s", name)
        req, err := o.client.NewRequest("POST", path, params, data)
        m := &NFSOffload{}
        _, err = o.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (o *OffloadService) DisconnectNFSOffload(name string, params map[string]string) (*NFSOffload, error) {

        path := fmt.Sprintf("nfs_offload/%s", name)
        req, err := o.client.NewRequest("DELETE", path, params, data)
        m := &NFSOffload{}
        _, err = o.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (o *OffloadService) GetNFSOffload(name string, params map[string]string) (*NFSOffload, error) {

        path := fmt.Sprintf("nfs_offload/%s", name)
        req, err := o.client.NewRequest("GET", path, params, data)
        m := &NFSOffload{}
        _, err = o.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}
