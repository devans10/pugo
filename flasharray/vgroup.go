package flasharray

import (
	"fmt"
)

type VgroupService struct {
	client *Client
}

func (h *VgroupService) CreateVgroup(name string, params map[string]string) (*Vgroup, error) {

        path := fmt.Sprintf("vgroup/%s", name)
        req, err := h.client.NewRequest("POST", path, params, nil)
        if err != nil {
                return nil, err
        }

        m := &Vgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *VgroupService) DestroyVgroup(name string, params map[string]string) (*Vgroup, error) {

        path := fmt.Sprintf("vgroup/%s", name)
        req, err := h.client.NewRequest("DELETE", path, params, nil)
        if err != nil {
                return nil, err
        }

        m := &Vgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}


func (h *VgroupService) EradicateVgroup(vgroup string, params map[string]string) (*Vgroup, error) {

        data := map[string]bool{"eradicate": true}
	path := fmt.Sprintf("vgroup/%s", vgroup)
	req, err := h.client.NewRequest("DELETE", path, params, data)
        if err != nil {
                return nil, err
        }

        m := &Vgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *VgroupService) GetVgroup(name string, params map[string]string) (*Vgroup, error) {

        path := fmt.Sprintf("vgroup/%s", name)
        req, err := h.client.NewRequest("GET", path, params, nil)
        if err != nil {
                return nil, err
        }

        m := &Vgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}



func (h *VgroupService) ListVgroups(params map[string]string) ([]Vgroup, error) {

        req, err := h.client.NewRequest("GET", "vgroup", params, nil)
        m := []Vgroup{}
        _, err = h.client.Do(req, &m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *VgroupService) RecoverVgroup(vgroup string, params map[string]string) (*Vgroup, error) {

        data := map[string]string{"action": "recover"}
        m, err := h.SetVgroup(vgroup, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *VgroupService) RenameVgroup(vgroup string, name string, params map[string]string) (*Vgroup, error) {

        data := map[string]string{"name": name}
        m, err := h.SetVgroup(vgroup, params, data)
        if err != nil {
                return nil, err
        }

        return m, err
}

func (h *VgroupService) SetVgroup(name string, params map[string]string, data interface{}) (*Vgroup, error) {

        path := fmt.Sprintf("vgroup/%s", name)
        req, err := h.client.NewRequest("PUT", path, params, data)
        if err != nil {
                return nil, err
        }

        m := &Vgroup{}
        _, err = h.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}
