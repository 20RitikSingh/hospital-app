package config

import (
	"log"
	"os"
	"strconv"
)

// Config holds the configuration for the application.
type Config struct {
	DSN           string
	JWTSecret     string
	JWTExpiration int
}

// Load loads the configuration from a file or environment variables.
func Load() (*Config, error) {
	// This function should load the configuration from a file or environment variables.
	// For simplicity, we are returning a hardcoded configuration here.
	return &Config{
		DSN:           GetDSN(),
		JWTSecret:     GetJWTSecret(),
		JWTExpiration: GetJWTExpiration(),
	}, nil
}

func GetDSN() string {
	// This function should return the Data Source Name (DSN) for your database connection.
	// The DSN typically includes the username, password, host, port, and database name.
	// For example: "user:password@tcp(localhost:3306)/dbname"
	dsn := os.Getenv("POSTGRESQL_DSN")
	if dsn == "" {
		log.Println("POSTGRESQL_DSN environment variable not set, using default DSN")
		return "default_dsn" // Default DSN for local development
	}
	return dsn
}

func GetJWTSecret() string {
	// This function should return the secret key used for signing JWT tokens.
	// Make sure to keep this secret and not expose it in your codebase.
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Println("JWT_SECRET environment variable not set, using default secret")
		return "default"
	}
	return jwtSecret
}

func GetJWTExpiration() int {
	// This function should return the expiration time for JWT tokens in minutes.
	// You can set this value based on your application's requirements.
	jwtExpiration := os.Getenv("JWT_EXPIRATION")
	if jwtExpiration == "" {
		log.Println("JWT_EXPIRATION environment variable not set, using default expiration time")
		return 60 // Default to 60 minutes
	}
	// Convert the expiration time from string to int
	t, err := strconv.Atoi(jwtExpiration)
	if err != nil {
		log.Println("JWT_EXPIRATION environment variable is not a valid integer, using default expiration time")
		return 60 // Default to 60 minutes
	}

	return t // Default to 60 minutes
}
