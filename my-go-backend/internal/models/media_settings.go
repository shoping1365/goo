package models

import "time"

// MediaSettings represents the configuration for image sizes
type MediaSettings struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`      // نام تنظیمات (مثلاً: محصولات، دسته‌بندی‌ها، پست‌ها)
	Width     int       `json:"width"`     // عرض تصویر
	Height    int       `json:"height"`    // ارتفاع تصویر
	Quality   int       `json:"quality"`   // کیفیت تصویر (0-100)
	IsActive  bool      `json:"is_active"` // آیا فعال است؟
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MediaSettingsResponse represents the API response for media settings
type MediaSettingsResponse struct {
	Settings []MediaSettings `json:"settings"`
	Total    int             `json:"total"`
}
