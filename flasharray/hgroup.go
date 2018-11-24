package flasharray

import (
	"fmt"
)

type HostgroupService struct {
	client *Client
}

func (h *HostgroupService) ConnectHostgroup(hgroup string, volume string, params map[string]string) (*ConnectedVolume, error) {

        path := fmt.Sprintf("hgroup/%s/volume/%s", hgroup, volume)
        req, err := h.client.NewRequest("POST", path, params, nil)
        m := &ConnectedVolume{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostgroupService) CreateHostgroup(name string, params map[string]string) (*Hostgroup, error) {

        path := fmt.Sprintf("hgroup/%s", name)
        req, err := h.client.NewRequest("POST", path, params, nil)
        if err != nil {
                return nil, err
        }

        m := &Hostgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostgroupService) DeleteHostgroup(name string, params map[string]string) (*Hostgroup, error) {

        path := fmt.Sprintf("hgroup/%s", name)
        req, err := h.client.NewRequest("DELETE", path, params, nil)
        if err != nil {
                return nil, err
        }

        m := &Hostgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostgroupService) DisonnectHostgroup(hgroup string, volume string, params map[string]string) (*ConnectedVolume, error) {

        path := fmt.Sprintf("hgroup/%s/volume/%s", hgroup, volume)
        req, err := h.client.NewRequest("DELETE", path, params, nil)
        m := &ConnectedVolume{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostgroupService) GetHostgroup(name string, params map[string]string) (*Hostgroup, error) {

        path := fmt.Sprintf("hgroup/%s", name)
        req, err := h.client.NewRequest("GET", path, params, nil)
        if err != nil {
                return nil, err
        }

        m := &Hostgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostgroupService) AddHostgroup(hgroup string, pgroup string, params map[string]string) (*HostgroupPgroup, error) {

        path := fmt.Sprintf("hgroup/%s/pgroup/%s", hgroup, pgroup)
        req, err := h.client.NewRequest("POST", path, params, nil)
        m := &HostgroupPgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostgroupService) RemoveHostgroup(hgroup string, pgroup string, params map[string]string) (*HostgroupPgroup, error) {

        path := fmt.Sprintf("hgroup/%s/pgroup/%s", hgroup, pgroup)
        req, err := h.client.NewRequest("DELETE", path, params, nil)
        m := &HostgroupPgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostgroupService) ListHostgroupConnections(hgroup string, params map[string]string) ([]HostgroupConnection, error) {

        path := fmt.Sprintf("hgroup/%s/volume", hgroup)
        req, err := h.client.NewRequest("GET", path, params, nil)
        m := []HostgroupConnection{}
        _, err = h.client.Do(req, &m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostgroupService) ListHostgroups(hgroup string, params map[string]string) ([]Hostgroup, error) {

        path := fmt.Sprintf("hgroup", hgroup)
        req, err := h.client.NewRequest("GET", path, params, nil)
        m := []Hostgroup{}
        _, err = h.client.Do(req, &m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostgroupService) RenameHostgroup(hgroup string, name string, params map[string]string) (*Hostgroup, error) {

        data := map[string]string{"name": name}
        m, err := h.SetHostgroup(hgroup, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *HostgroupService) SetHostgroup(name string, params map[string]string, data interface{}) (*Hostgroup, error) {

        path := fmt.Sprintf("hgroup/%s", name)
        req, err := h.client.NewRequest("PUT", path, params, data)
        if err != nil {
                return nil, err
        }

        m := &Hostgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}
