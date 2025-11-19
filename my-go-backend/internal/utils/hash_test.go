package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "secret123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}
	if hash == "" {
		t.Error("HashPassword returned empty string")
	}
	if hash == password {
		t.Error("HashPassword returned the password itself, not hashed")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "secret123"
	hash, _ := HashPassword(password)

	if !CheckPassword(password, hash) {
		t.Error("CheckPassword failed for correct password")
	}

	if CheckPassword("wrongpassword", hash) {
		t.Error("CheckPassword succeeded for wrong password")
	}
}
