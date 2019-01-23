package pure1

import (
	"os"
	"testing"
)

func testAccPreChecks(t *testing.T) {

	if os.Getenv("PURE1_ACC") == "" {
		t.Skip("set PURE1_ACC to run purestorage acceptance tests (provider connection is required)")
	}

	appID := os.Getenv("PURE1_APPID")
	privateKey := os.Getenv("PURE1_PRIVATEKEY")

	if appID == "" {
		t.Fatalf("PURE1_APPID must be set for acceptance tests")
	}
	if privateKey == "" {
		t.Fatalf("PURE1_PRIVATEKEY must be set for acceptance tests")
	}
}

func testAccGenerateClient(t *testing.T) *Client {
	appid := os.Getenv("PURE1_APPID")
	privateKey := []byte(os.Getenv("PURE1_PRIVATEKEY"))
	restVersion := ""

	c, err := NewClient(appid, privateKey, restVersion)
	if err != nil {
		t.Fatalf("error setting up client: %s", err)
	}
	return c
}
func TestPure1NewClient(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)
	if c == nil {
		t.Fatal("error setting up client")
	}
}
