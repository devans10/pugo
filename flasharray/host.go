package flasharray

import (
	"fmt"
)

type HostService struct {
	client *Client
}

func (h *HostService) ConnectHost(host string, volume string, params map[string]string) (*ConnectedVolume, error) {

        path := fmt.Sprintf("host/%s/volume/%s", host, volume)
        req, err := h.client.NewRequest("POST", path, params, nil)
        m := &ConnectedVolume{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostService) CreateHost(name string, params map[string]string) (*Host, error) {

        path := fmt.Sprintf("host/%s", name)
        req, err := h.client.NewRequest("POST", path, params, nil)
        if err != nil {
                return nil, err
        }

        m := &Host{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostService) DeleteHost(name string, params map[string]string) (*Host, error) {

        path := fmt.Sprintf("host/%s", name)
        req, err := h.client.NewRequest("DELETE", path, params, nil)
        if err != nil {
                return nil, err
        }

        m := &Host{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostService) DisonnectHost(host string, volume string, params map[string]string) (*ConnectedVolume, error) {

        path := fmt.Sprintf("host/%s/volume/%s", host, volume)
        req, err := h.client.NewRequest("DELETE", path, params, nil)
        m := &ConnectedVolume{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostService) GetHost(name string, params map[string]string) (*Host, error) {

        path := fmt.Sprintf("host/%s", name)
        req, err := h.client.NewRequest("GET", path, params, nil)
        if err != nil {
                return nil, err
        }

        m := &Host{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostService) AddHost(host string, pgroup string, params map[string]string) (*HostPgroup, error) {

        path := fmt.Sprintf("host/%s/pgroup/%s", host, volume)
        req, err := h.client.NewRequest("POST", path, params, nil)
        m := &HostPgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostService) RemoveHost(host string, pgroup string, params map[string]string) (*HostPgroup, error) {

        path := fmt.Sprintf("host/%s/pgroup/%s", host, volume)
        req, err := h.client.NewRequest("DELETE", path, params, nil)
        m := &HostPgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostService) ListHostConnections(host string, params map[string]string) ([]HostConnection, error) {

        path := fmt.Sprintf("host/%s/volume", host)
        req, err := h.client.NewRequest("GET", path, params, nil)
        m := []HostConnection{}
        _, err = h.client.Do(req, &m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostService) ListHosts(host string, params map[string]string) ([]Host, error) {

        path := fmt.Sprintf("host", host)
        req, err := h.client.NewRequest("GET", path, params, nil)
        m := []Host{}
        _, err = h.client.Do(req, &m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostService) RenameHost(host string, name string, params map[string]string) (*Host, error) {

        data := map[string]string{"name": name}
        m, err := h.SetHost(host, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (v *HostService) SetHost(name string, params map[string]string, data interface{}) (*Host, error) {

        path := fmt.Sprintf("host/%s", name)
        req, err := h.client.NewRequest("PUT", path, params, data)
        if err != nil {
                return nil, err
        }

        m := &Host{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}
