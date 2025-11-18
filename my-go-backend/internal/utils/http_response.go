package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON ارسال پاسخ JSON استاندارد
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// RespondWithError ارسال پاسخ خطا در قالب استاندارد پروژه
// code: کد خطا (مثلاً VALIDATION_ERROR) یا هر کلید تجاری دیگر
func RespondWithError(w http.ResponseWriter, status int, code string, message string) {
	resp := map[string]interface{}{
		"success": false,
		"error":   code,
		"message": message,
	}
	RespondWithJSON(w, status, resp)
}
