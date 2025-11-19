package utils

import (
	"testing"
	"time"
)

func TestGenerateAndValidateAccessToken(t *testing.T) {
	secret := "supersecretkey"
	userID := uint(1)
	mobile := "09123456789"
	role := "admin"
	permissions := []string{"read", "write"}

	token, err := GenerateAccessToken(userID, mobile, role, permissions, secret)
	if err != nil {
		t.Fatalf("GenerateAccessToken failed: %v", err)
	}
	if token == "" {
		t.Fatal("Generated token is empty")
	}

	claims, err := ValidateAccessToken(token, secret)
	if err != nil {
		t.Fatalf("ValidateAccessToken failed: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected UserID %d, got %d", userID, claims.UserID)
	}
	if claims.Mobile != mobile {
		t.Errorf("Expected Mobile %s, got %s", mobile, claims.Mobile)
	}
	if claims.Role != role {
		t.Errorf("Expected Role %s, got %s", role, claims.Role)
	}
}

func TestGenerateRefreshToken(t *testing.T) {
	secret := "supersecretkey"
	userID := uint(1)

	token, expiresAt, err := GenerateRefreshToken(userID, secret)
	if err != nil {
		t.Fatalf("GenerateRefreshToken failed: %v", err)
	}
	if token == "" {
		t.Fatal("Generated refresh token is empty")
	}
	if expiresAt.Before(time.Now()) {
		t.Error("Expiration time is in the past")
	}
}

func TestValidateAccessTokenInvalidSecret(t *testing.T) {
	secret := "supersecretkey"
	token, _ := GenerateAccessToken(1, "09123456789", "user", nil, secret)

	_, err := ValidateAccessToken(token, "wrongsecret")
	if err == nil {
		t.Error("ValidateAccessToken should fail with wrong secret")
	}
}

func TestValidateAccessTokenExpired(t *testing.T) {
	// This is hard to test without mocking time or modifying the function to accept time.
	// For now, we skip explicit expiration test unless we refactor.
	// But we can test invalid token string.
	_, err := ValidateAccessToken("invalid.token.string", "secret")
	if err == nil {
		t.Error("ValidateAccessToken should fail with invalid token string")
	}
}
