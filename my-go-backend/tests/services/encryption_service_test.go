package services_test

import (
	"os"
	"testing"

	"my-go-backend/internal/services"
)

func TestEncryptionService_NewEncryptionService(t *testing.T) {
	t.Run("DefaultKey", func(t *testing.T) {
		os.Unsetenv("ENCRYPTION_KEY")
		service, err := services.NewEncryptionService()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if service == nil {
			t.Error("expected service, got nil")
		}
	})

	t.Run("CustomKey", func(t *testing.T) {
		os.Setenv("ENCRYPTION_KEY", "my-secret-key-12345678901234567890")
		service, err := services.NewEncryptionService()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if service == nil {
			t.Error("expected service, got nil")
		}
	})
}
