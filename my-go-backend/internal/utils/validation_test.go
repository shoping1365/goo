package utils

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name@domain.co.in", true},
		{"invalid-email", false},
		{"@domain.com", false},
		{"user@", false},
		{"", false},
	}

	for _, test := range tests {
		if result := IsValidEmail(test.email); result != test.expected {
			t.Errorf("IsValidEmail(%s) = %v; expected %v", test.email, result, test.expected)
		}
	}
}

func TestIsValidMobile(t *testing.T) {
	tests := []struct {
		mobile   string
		expected bool
	}{
		{"09123456789", true},
		{"+989123456789", true},
		{"00989123456789", true},
		{"989123456789", true},
		{"0912-345-6789", true}, // Should handle separators
		{"1234567890", false},
		{"0912345678", false}, // Too short
		{"abc", false},
	}

	for _, test := range tests {
		if result := IsValidMobile(test.mobile); result != test.expected {
			t.Errorf("IsValidMobile(%s) = %v; expected %v", test.mobile, result, test.expected)
		}
	}
}

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		url      string
		expected bool
	}{
		{"https://google.com", true},
		{"http://localhost:8080", true},
		{"ftp://files.com", true},
		{"invalid-url", false},
		{"", false},
	}

	for _, test := range tests {
		if result := IsValidURL(test.url); result != test.expected {
			t.Errorf("IsValidURL(%s) = %v; expected %v", test.url, result, test.expected)
		}
	}
}

func TestIsValidColor(t *testing.T) {
	tests := []struct {
		color    string
		expected bool
	}{
		{"#FFFFFF", true},
		{"#000", true},
		{"red", true},
		{"Blue", true}, // Case insensitive check
		{"invalid-color", false},
		{"#ZZZZZZ", false},
	}

	for _, test := range tests {
		if result := IsValidColor(test.color); result != test.expected {
			t.Errorf("IsValidColor(%s) = %v; expected %v", test.color, result, test.expected)
		}
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"-5", -5},
		{"0", 0},
		{"abc", 0}, // Should return 0 on error
		{"", 0},
	}

	for _, test := range tests {
		if result := ParseInt(test.input); result != test.expected {
			t.Errorf("ParseInt(%s) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestContains(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	
	if !Contains(slice, "banana") {
		t.Error("Contains failed to find existing item")
	}
	
	if Contains(slice, "grape") {
		t.Error("Contains found non-existing item")
	}
}
