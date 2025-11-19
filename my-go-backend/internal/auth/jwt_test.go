package auth

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// helper to reset env between tests
func withEnv(key, value string, fn func()) {
	old := os.Getenv(key)
	_ = os.Setenv(key, value)
	defer os.Setenv(key, old)
	fn()
}

func TestGenerateAndValidateToken_DefaultSecret(t *testing.T) {
	// ensure no JWT_SECRET so getSecretKey falls back to default
	withEnv("JWT_SECRET", "", func() {
		token, err := GenerateToken(1, "testuser", "admin")
		if err != nil {
			t.Fatalf("GenerateToken returned error: %v", err)
		}
		if token == "" {
			t.Fatal("expected non-empty token")
		}

		claims, err := ValidateToken(token)
		if err != nil {
			t.Fatalf("ValidateToken returned error: %v", err)
		}
		if claims.UserID != 1 || claims.Username != "testuser" || claims.Role != "admin" {
			t.Fatalf("unexpected claims: %+v", claims)
		}
	})
}

func TestGenerateAndValidateToken_CustomSecret(t *testing.T) {
	withEnv("JWT_SECRET", "custom-secret", func() {
		// generate token with custom secret
		token, err := GenerateToken(2, "user2", "editor")
		if err != nil {
			t.Fatalf("GenerateToken returned error: %v", err)
		}

		claims, err := ValidateToken(token)
		if err != nil {
			t.Fatalf("ValidateToken returned error: %v", err)
		}
		if claims.UserID != 2 || claims.Username != "user2" || claims.Role != "editor" {
			t.Fatalf("unexpected claims: %+v", claims)
		}
	})
}

func TestValidateToken_InvalidToken(t *testing.T) {
	_, err := ValidateToken("not-a-jwt-token")
	if err == nil {
		t.Fatal("expected error for invalid token")
	}
	if err != ErrInvalidToken {
		t.Fatalf("expected ErrInvalidToken, got %v", err)
	}
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	// build an already-expired token manually using the same claims type
	claims := Claims{
		UserID:   3,
		Username: "expired",
		Role:     "user",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString(getSecretKey())
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}

	_, err = ValidateToken(str)
	if err == nil {
		t.Fatal("expected error for expired token")
	}
	if err != ErrExpiredToken {
		t.Fatalf("expected ErrExpiredToken, got %v", err)
	}
}

func TestRefreshToken(t *testing.T) {
	token, err := GenerateToken(10, "refresh-user", "member")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	newToken, err := RefreshToken(token)
	if err != nil {
		t.Fatalf("RefreshToken returned error: %v", err)
	}
	if newToken == "" {
		t.Fatal("expected non-empty refreshed token")
	}

	claims, err := ValidateToken(newToken)
	if err != nil {
		t.Fatalf("ValidateToken returned error for refreshed token: %v", err)
	}
	if claims.UserID != 10 || claims.Username != "refresh-user" || claims.Role != "member" {
		t.Fatalf("unexpected claims after refresh: %+v", claims)
	}
}
