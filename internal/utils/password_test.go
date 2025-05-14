package utils

import (
	"testing"
)

func TestHashPasswordAndCheck(t *testing.T) {
	password := "securepassword123"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	if hashedPassword == "" {
		t.Error("HashPassword returned empty string")
	}

	match, err := CheckPasswordHash(password, hashedPassword)
	if err != nil {
		t.Errorf("CheckPasswordHash returned error: %v", err)
	}
	if !match {
		t.Error("Password and hash do not match, but they should")
	}
}
