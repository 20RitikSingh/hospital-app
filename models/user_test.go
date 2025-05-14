package models

import (
	"testing"
)

func TestUserRole_String(t *testing.T) {
	cases := []struct {
		role     UserRole
		expected string
	}{
		{Doctor, "doctor"},
		{Receptionist, "receptionist"},
		{UserRole(99), "unknown"},
	}
	for _, c := range cases {
		if got := c.role.String(); got != c.expected {
			t.Errorf("UserRole.String() = %q, want %q", got, c.expected)
		}
	}
}

func TestParseRole(t *testing.T) {
	cases := []struct {
		input    string
		expected UserRole
		wantErr  bool
	}{
		{"doctor", Doctor, false},
		{"receptionist", Receptionist, false},
		{"invalid", -1, true},
	}
	for _, c := range cases {
		got, err := ParseRole(c.input)
		if (err != nil) != c.wantErr {
			t.Errorf("ParseRole(%q) error = %v, wantErr %v", c.input, err, c.wantErr)
		}
		if got != c.expected {
			t.Errorf("ParseRole(%q) = %v, want %v", c.input, got, c.expected)
		}
	}
}

func TestUserRole_Scan_Value(t *testing.T) {
	var r UserRole
	if err := r.Scan(int64(1)); err != nil || r != Receptionist {
		t.Errorf("Scan(int64) failed: %v, got %v", err, r)
	}
	if err := r.Scan(int(0)); err != nil || r != Doctor {
		t.Errorf("Scan(int) failed: %v, got %v", err, r)
	}
	if err := r.Scan("bad"); err == nil {
		t.Error("Scan(string) should fail")
	}
	v, err := r.Value()
	if err != nil || v != int(r) {
		t.Errorf("Value() = %v, %v; want %v, nil", v, err, int(r))
	}
}

func TestUser_IDString(t *testing.T) {
	u := &User{ID: 42}
	if got := u.IDString(); got != "42" {
		t.Errorf("IDString() = %q, want %q", got, "42")
	}
}
