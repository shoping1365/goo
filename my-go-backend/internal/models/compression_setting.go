package models

import (
	"time"
)

// CompressionSetting مدل تنظیمات فشرده‌سازی (سراسری، کاربر، رسانه، عملیات)
// این مدل برای ذخیره تنظیمات پیش‌فرض و سفارشی استفاده می‌شود.
type CompressionSetting struct {
	ID        uint      `gorm:"primaryKey" json:"id"`                         // کلید اصلی
	Scope     string    `gorm:"type:varchar(20);not null;index" json:"scope"` // دامنه تنظیمات (global, user, media, job)
	UserID    *uint     `gorm:"index" json:"user_id"`                         // کلید خارجی به users (nullable)
	MediaID   *uint     `gorm:"index" json:"media_id"`                        // کلید خارجی به media_files (nullable)
	JobID     *uint     `gorm:"index" json:"job_id"`                          // کلید خارجی به compression_jobs (nullable)
	Key       string    `gorm:"type:varchar(50);not null;index" json:"key"`   // کلید تنظیمات
	Value     string    `gorm:"type:json;not null" json:"value"`              // مقدار تنظیمات (JSON یا رشته)
	IsDefault bool      `gorm:"default:false;index" json:"is_default"`        // آیا مقدار پیش‌فرض است؟
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
