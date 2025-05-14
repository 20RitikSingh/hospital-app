package config

import "testing"

func TestLoad(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.DSN == "" {
		t.Error("Expected DSN to be set, got empty string")
	}

	if cfg.JWTSecret == "" {
		t.Error("Expected JWTSecret to be set, got empty string")
	}

	if cfg.JWTExpiration <= 0 {
		t.Errorf("Expected JWTExpiration to be greater than 0, got %d", cfg.JWTExpiration)
	}
}

func TestGetDSN(t *testing.T) {
	dsn := GetDSN()
	if dsn == "" {
		t.Error("Expected DSN to be set, got empty string")
	}
}

func TestGetJWTSecret(t *testing.T) {
	jwtSecret := GetJWTSecret()
	if jwtSecret == "" {
		t.Error("Expected JWTSecret to be set, got empty string")
	}
}

func TestGetJWTExpiration(t *testing.T) {
	jwtExpiration := GetJWTExpiration()
	if jwtExpiration <= 0 {
		t.Errorf("Expected JWTExpiration to be greater than 0, got %d", jwtExpiration)
	}
}
