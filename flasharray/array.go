package flasharray

import (
        "log"
)

type ArrayService struct {
        client *Client
}

func (v *ArrayService) set_console_lock(b bool) (*console_lock, error) {

	data := map[string]bool{"enabled": b}
	req, err := v.NewRequest("PUT", "array/console_lock", nil, data)
        if err != nil {
                return nil, err
        }

        m := &console_lock{}
        _, err = v.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return &m, err
}

func (v *ArrayService) enable_console_lock() error {

	_, err := v.set_console_lock(true)
	if err != nil {
		return err
	}
	return nil
}

func (v *ArrayService) disable_console_lock() error {

        _, err := v.set_console_lock(false)
        if err != nil {
                return err
        }
        return nil
}

func (v *ArrayService) get_console_lock() (*console_lock, err) {

	req, err := v.NewRequest("GET", "array/console_lock", nil, nil)
        if err != nil {
                return nil, err
        }

        m := &console_lock{}
        _, err = v.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return &m, err
}

func (v *ArrayService) Get(params map[string]string) (*Array, error) {

	req, err := v.NewRequest("GET", "array", params, nil)
	if err != nil {
		return nil, err
	}

	m := &Array{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return &m, err
}

func (v *ArrayService) Set(params map[string]string) (*Array, error) {

        req, err := v.NewRequest("PUT", "array", params, nil)
        if err != nil {
                return nil, err
        }

        m := &Array{}
        _, err = v.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return &m, err
}

func (v *ArrayService) Rename(name string) (*Array, error) {

	params := map[string]string{"name": name}
	m, err := v.Set(params)
        return &m, err
}
