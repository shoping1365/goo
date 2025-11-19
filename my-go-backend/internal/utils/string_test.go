package utils

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	length := 10
	str := GenerateRandomString(length)
	if len(str) != length {
		t.Errorf("Expected length %d, got %d", length, len(str))
	}
}

func TestGenerateRandomNumber(t *testing.T) {
	length := 6
	num := GenerateRandomNumber(length)
	if len(num) != length {
		t.Errorf("Expected length %d, got %d", length, len(num))
	}
	for _, c := range num {
		if c < '0' || c > '9' {
			t.Errorf("Expected digit, got %c", c)
		}
	}
}
