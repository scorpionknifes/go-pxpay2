package pxpay2

import "testing"

func TestNewClient(t *testing.T) {
	c, err := NewClient("", "", "")
	if err == nil {
		t.Errorf("Expected error for NewClient('','','')")
	}
	if c != nil {
		t.Errorf("Expected nil Client for NewClient('','',''), got %v", c)
	}

	c, err = NewClient("1", "2", "3")
	if err != nil {
		t.Errorf("Not expected error for NewClient(1, 2, 3), got %v", err)
	}
	if c == nil {
		t.Errorf("Expected non-nil Client for NewClient(1, 2, 3)")
	}
}
