package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	// Setup test data
	userID := "123"
	role := "doctor"
	secretKey := "test-secret-key"
	expirationMinutes := 60

	// Generate token
	token, err := GenerateToken(userID, role, secretKey, expirationMinutes)

	// Assert no error
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Validate the generated token
	claims, err := ValidateToken(token, secretKey)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, role, claims.Role)
	assert.Equal(t, userID, claims.Subject)
	assert.Equal(t, "hospital-api", claims.Issuer)

	// Check that expiration is set correctly (within acceptable margin)
	expectedExp := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)
	assert.WithinDuration(t, expectedExp, claims.ExpiresAt.Time, 2*time.Second)
}

func TestValidateToken(t *testing.T) {
	// Setup
	userID := "456"
	role := "receptionist"
	secretKey := "another-test-key"

	t.Run("valid token", func(t *testing.T) {
		// Generate a valid token
		token, _ := GenerateToken(userID, role, secretKey, 60)

		// Validate token
		claims, err := ValidateToken(token, secretKey)
		assert.NoError(t, err)
		assert.Equal(t, userID, claims.UserID)
		assert.Equal(t, role, claims.Role)
	})
	t.Run("invalid signing method", func(t *testing.T) {
		// Create token with wrong signing method
		claims := &Claims{
			UserID: userID,
			Role:   role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
		tokenString, _ := token.SignedString(jwt.UnsafeAllowNoneSignatureType)

		// Try to validate token
		_, err := ValidateToken(tokenString, secretKey)
		assert.Error(t, err)
		// Don't check exact error, just verify it failed
		assert.Contains(t, err.Error(), "unverifiable")
	})

	t.Run("wrong secret key", func(t *testing.T) {
		// Generate token with one key
		token, _ := GenerateToken(userID, role, secretKey, 60)

		// Validate with different key
		_, err := ValidateToken(token, "wrong-secret")
		assert.Error(t, err)
	})
	t.Run("expired token", func(t *testing.T) {
		// Generate a token that's already expired
		claims := Claims{
			UserID: userID,
			Role:   role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(-10 * time.Minute)), // expired 10 minutes ago
				IssuedAt:  jwt.NewNumericDate(time.Now().Add(-60 * time.Minute)),
				Issuer:    "hospital-api",
				Subject:   userID,
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(secretKey))

		// Validate expired token
		_, err := ValidateToken(tokenString, secretKey)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "expired") // Just check for "expired" substring
	})

	t.Run("malformed token", func(t *testing.T) {
		// Try to validate a malformed token
		malformedToken := "not.a.token"
		_, err := ValidateToken(malformedToken, secretKey)
		assert.Error(t, err)
	})
}

func TestTokenIntegration(t *testing.T) {
	// This test verifies the full token lifecycle
	userID := "789"
	role := "doctor"
	secretKey := "integration-test-key"
	expirationMinutes := 5

	// Generate token
	token, err := GenerateToken(userID, role, secretKey, expirationMinutes)
	assert.NoError(t, err)

	// Validate token and check claims
	claims, err := ValidateToken(token, secretKey)
	assert.NoError(t, err)

	// Verify all claims are present and correct
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, role, claims.Role)
	assert.Equal(t, userID, claims.Subject)
	assert.Equal(t, "hospital-api", claims.Issuer)

	// Check time claims
	now := time.Now()
	assert.True(t, claims.IssuedAt.Time.Before(now) || claims.IssuedAt.Time.Equal(now),
		"IssuedAt should be now or earlier")
	assert.True(t, claims.ExpiresAt.Time.After(now),
		"ExpiresAt should be after now")
	// Check expiration time
	assert.True(t, claims.ExpiresAt.Time.After(now), "Token should not be expired")
}
