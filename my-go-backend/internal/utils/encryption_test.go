package utils

import (
	"testing"
)

func TestEncryptDecryptString(t *testing.T) {
	originalText := "This is a secret message"
	
	// Test Encryption
	encrypted, err := EncryptString(originalText)
	if err != nil {
		t.Fatalf("EncryptString failed: %v", err)
	}
	if encrypted == "" {
		t.Error("EncryptString returned empty string")
	}
	if encrypted == originalText {
		t.Error("EncryptString returned original text")
	}

	// Test Decryption
	decrypted, err := DecryptString(encrypted)
	if err != nil {
		t.Fatalf("DecryptString failed: %v", err)
	}
	if decrypted != originalText {
		t.Errorf("DecryptString returned %s, expected %s", decrypted, originalText)
	}
}

func TestDecryptInvalidString(t *testing.T) {
	_, err := DecryptString("invalid-base64-string")
	if err == nil {
		t.Error("DecryptString should fail for invalid base64 string")
	}
}
