package pure_tests

import (
	"testing"

	"github.com/devans10/go-purestorage/flasharray"
)

// Test that a NewClient call with no authentication returns an error
func TestNewClientNoAuth(t *testing.T) {

	_, err := flasharray.NewClient("target", "", "", "", "rest_version", false, false, "user_agent", nil)
	if err == nil {
		t.Errorf("An Error was NOT raised when no authentication methods provided")
	}
}

// Test that a NewClient call with all authentication methods provided returns an error
func TestNewClientAllAuth(t *testing.T) {

	_, err := flasharray.NewClient("target", "username", "password", "api_token", "rest_version", false, false, "user_agent", nil)
        if err == nil {
                t.Errorf("An Error was NOT raised when All authentication methods were provided")
        }
}
