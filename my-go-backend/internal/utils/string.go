package utils

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomString تولید رشته تصادفی امن با طول مشخص
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		// استفاده از crypto/rand برای امنیت بیشتر
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}
	return string(b)
}

// GenerateRandomNumber تولید عدد تصادفی با طول مشخص
func GenerateRandomNumber(length int) string {
	const digits = "0123456789"
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		b[i] = digits[num.Int64()]
	}
	return string(b)
}

// GenerateRandomAlphaNumeric تولید رشته الفبایی-عددی تصادفی
func GenerateRandomAlphaNumeric(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		b[i] = chars[num.Int64()]
	}
	return string(b)
}

