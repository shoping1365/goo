package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword هش کردن رمز عبور با bcrypt
func HashPassword(password string) (string, error) {
	const passwordCost = 14 // هزینهٔ بالاتر برای افزایش امنیت
	hash, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword مقایسه رمز عبور خام با هش ذخیره شده
func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
