package models

import "time"

// CategoryQA represents a Q&A category
type CategoryQA struct {
	ID        int       `json:"id"`
	Key       string    `json:"key"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
