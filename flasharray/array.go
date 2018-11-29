package flasharray

// ArrayService type creates a service to perform functions for administering
// and querying the flash array itself
type ArrayService struct {
        client *Client
}

// set_console_lock is a helper function used to set the console lock
func (v *ArrayService) set_console_lock(b bool) (*console_lock, error) {

	data := map[string]bool{"enabled": b}
	req, err := v.client.NewRequest("PUT", "array/console_lock", nil, data)
        if err != nil {
                return nil, err
        }

        m := &console_lock{}
        _, err = v.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// enable_console_lock enables root lockout from the array at the physical console.
// returns A dictionary mapping "console_lock" to "enabled".
func (v *ArrayService) enable_console_lock() error {

	_, err := v.set_console_lock(true)
	if err != nil {
		return err
	}
	return nil
}

// disable_console_lock disables root lockout from the array at the physical console.
// returns A dictionary mapping "console_lock" to "disabled".
func (v *ArrayService) disable_console_lock() error {

        _, err := v.set_console_lock(false)
        if err != nil {
                return err
        }
        return nil
}

// get_console_lock returns an object giving the console_lock status
func (v *ArrayService) get_console_lock() (*console_lock, error) {

	req, err := v.client.NewRequest("GET", "array/console_lock", nil, nil)
        if err != nil {
                return nil, err
        }

        m := &console_lock{}
        _, err = v.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Get returns and object describing the flash array
func (v *ArrayService) Get(data interface{}) (*Array, error) {

	req, err := v.client.NewRequest("GET", "array", nil, data)
	if err != nil {
		return nil, err
	}

	m := &Array{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Set will change the parameter on the array that is passed in the data map
func (v *ArrayService) Set(params map[string]string, data interface{}) (*Array, error) {

        req, err := v.client.NewRequest("PUT", "array", params, nil)
        if err != nil {
                return nil, err
        }

        m := &Array{}
        _, err = v.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Rename will change the name of the flash array
func (v *ArrayService) Rename(name string) (*Array, error) {

	data := map[string]string{"name": name}
	m, err := v.Set(nil, data)
        return m, err
}

// Set the phonehome service attributes
func (v *ArrayService) setPhoneHome(data interface{}) (*Phonehome, error) {

	req, err := v.client.NewRequest("PUT", "array/phonehome", nil, data)
	if err != nil {
		return nil, err
	}

	m := &Phonehome{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Set the remote assist service attributes
func (v *ArrayService) setRemoteAssist(data interface{}) (*RemoteAssist, error) {

        req, err := v.client.NewRequest("PUT", "array/remoteassist", nil, data)
        if err != nil {
                return nil, err
        }

        m := &RemoteAssist{}
        _, err = v.client.Do(req, m, false)
        if err != nil {
                return nil, err
        }

        return m, err
}

// Disable hourly phonehome
func (v *ArrayService) DisablePhoneHome() (*Phonehome, error) {

	data := map[string]bool{"enabled": false}
	m, err := v.setPhoneHome(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Disable Remote Assist
func (v *ArrayService) DisableRemoteAssist() (*RemoteAssist, error) {

	data := map[string]bool{"enabled": false}
	m, err := v.setRemoteAssist(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Enable hourly phonehome
func (v *ArrayService) EnablePhoneHome() (*Phonehome, error) {

	data := map[string]bool{"enabled": true}
	m, err := v.setPhoneHome(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Enable Remote Assist
func (v *ArrayService) EnableRemoteAssist() (*RemoteAssist, error) {

	data := map[string]bool{"enabled": true}
	m, err := v.setRemoteAssist(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Get manual phone home status
func (v *ArrayService) GetManualPhoneHome() (*Phonehome, error) {

	req, err := v.client.NewRequest("GET", "array/phoneome", nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Phonehome{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Get Phonehome status
func (v *ArrayService) GetPhoneHome() (*Array, error) {

	data := map[string]bool{"phonehome": true}
	m, err := v.Get(data)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Get Remote Status
func (v *ArrayService) GetRemoteAssist() (*RemoteAssist, error) {

	req, err := v.client.NewRequest("GET", "array/remoteassist", nil, nil)
	if err != nil {
		return nil, err
	}

	m := &RemoteAssist{}
	_, err = v.client.Do(req, m, false)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Manually initiate or cancel phonehome
//
// Parameters
// action
// The timeframe of logs to phonehome or cancel the current phonehome
// action must be one of:
// "send_today", "send_yesterday", "send_all", "cancel"
func (v *ArrayService) Phonehome(action string) (*Phonehome, error) {

	data := map[string]string{"action": action}
	m, err := v.setPhoneHome(data)
	if err != nil {
		return nil, err
	}

	return m, err
}
