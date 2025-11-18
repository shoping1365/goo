package utils

type ApiError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func New(code, msg string, details interface{}) ApiError {
	return ApiError{Code: code, Message: msg, Details: details}
}
