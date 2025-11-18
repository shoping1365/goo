package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// EncryptionKey کلید رمزگذاری (در محیط واقعی باید از متغیر محیطی خوانده شود)
// برای AES-256-GCM نیاز به دقیقاً 32 بایت داریم
var EncryptionKey = []byte("your-32-byte-encryption-key-here!!")

// generateKeyFromString یک کلید 32 بایتی از رشته تولید می‌کند
func generateKeyFromString(input string) []byte {
	hash := sha256.Sum256([]byte(input))
	return hash[:]
}

// getEncryptionKey کلید رمزگذاری را برمی‌گرداند (32 بایت)
func getEncryptionKey() []byte {
	// اگر کلید موجود 32 بایت نیست، از hash استفاده کن
	if len(EncryptionKey) != 32 {
		return generateKeyFromString(string(EncryptionKey))
	}
	return EncryptionKey
}

// EncryptString رشته را رمزگذاری می‌کند
func EncryptString(text string) (string, error) {
	// تبدیل متن به بایت
	plaintext := []byte(text)

	// دریافت کلید رمزگذاری
	key := getEncryptionKey()

	// ایجاد cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("خطا در ایجاد cipher: %w", err)
	}

	// ایجاد GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("خطا در ایجاد GCM: %w", err)
	}

	// ایجاد nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("خطا در ایجاد nonce: %w", err)
	}

	// رمزگذاری
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// تبدیل به base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptString رشته رمزگذاری شده را رمزگشایی می‌کند
func DecryptString(encryptedText string) (string, error) {
	// تبدیل از base64
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", fmt.Errorf("خطا در decode کردن base64: %w", err)
	}

	// دریافت کلید رمزگذاری
	key := getEncryptionKey()

	// ایجاد cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("خطا در ایجاد cipher: %w", err)
	}

	// ایجاد GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("خطا در ایجاد GCM: %w", err)
	}

	// بررسی طول nonce
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext خیلی کوتاه است")
	}

	// جدا کردن nonce و ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// رمزگشایی
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("خطا در رمزگشایی: %w", err)
	}

	return string(plaintext), nil
}

// IsEncrypted بررسی می‌کند که آیا متن رمزگذاری شده است یا خیر
func IsEncrypted(text string) bool {
	if text == "" {
		return false
	}

	// بررسی اینکه آیا متن با الگوی رمزگذاری شده مطابقت دارد
	decoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return false
	}

	// بررسی طول (حداقل باید nonce + حداقل 1 بایت باشد)
	// AES-GCM nonce size is 12 bytes
	if len(decoded) < 13 {
		return false
	}

	// بررسی اینکه آیا متن با الگوی API Key شروع می‌شود
	// اگر با "sk-" شروع شود، احتمالاً رمزگذاری نشده است
	if len(text) > 3 && text[:3] == "sk-" {
		return false
	}

	return true
}
