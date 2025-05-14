package models

import (
	"testing"
)

func TestGender_String(t *testing.T) {
	cases := []struct {
		g        Gender
		expected string
	}{
		{Male, "male"},
		{Female, "female"},
		{Other, "other"},
		{Gender(99), "unknown"},
	}
	for _, c := range cases {
		if got := c.g.String(); got != c.expected {
			t.Errorf("Gender.String() = %q, want %q", got, c.expected)
		}
	}
}

func TestParseGender(t *testing.T) {
	cases := []struct {
		input    string
		expected Gender
		wantErr  bool
	}{
		{"male", Male, false},
		{"female", Female, false},
		{"other", Other, false},
		{"invalid", -1, true},
	}
	for _, c := range cases {
		got, err := ParseGender(c.input)
		if (err != nil) != c.wantErr {
			t.Errorf("ParseGender(%q) error = %v, wantErr %v", c.input, err, c.wantErr)
		}
		if got != c.expected {
			t.Errorf("ParseGender(%q) = %v, want %v", c.input, got, c.expected)
		}
	}
}

func TestGender_Scan_Value(t *testing.T) {
	var g Gender
	if err := g.Scan(int64(1)); err != nil || g != Female {
		t.Errorf("Scan(int64) failed: %v, got %v", err, g)
	}
	if err := g.Scan(int(0)); err != nil || g != Male {
		t.Errorf("Scan(int) failed: %v, got %v", err, g)
	}
	if err := g.Scan("bad"); err == nil {
		t.Error("Scan(string) should fail")
	}
	v, err := g.Value()
	if err != nil || v != int(g) {
		t.Errorf("Value() = %v, %v; want %v, nil", v, err, int(g))
	}
}
