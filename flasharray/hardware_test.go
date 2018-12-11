// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"testing"
)

func TestAccGetDrive(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	expected := "CH0.BAY0"
	h, err := c.Hardware.GetDrive(expected)
	if err != nil {
		t.Fatalf("error getting drive: %s", err)
	}

	if h.Name != expected {
		t.Fatalf("expected: %s, got: %s", expected, h.Name)
	}
}

func TestAccListDrives(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Hardware.ListDrives()
	if err != nil {
		t.Fatalf("error listing drives: %s", err)
	}
}

func TestAccGetHardware(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	expected := "CT0"
	h, err := c.Hardware.GetHardware(expected)
	if err != nil {
		t.Fatalf("error getting hardware: %s", err)
	}

	if h.Name != expected {
		t.Fatalf("expected: %s, got: %s", expected, h.Name)
	}
}

func TestAccListHardware(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	_, err := c.Hardware.ListHardware()
	if err != nil {
		t.Fatalf("error listing hardware: %s", err)
	}
}
