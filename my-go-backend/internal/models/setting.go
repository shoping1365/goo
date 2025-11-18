package models

import (
	"time"
)

/*
مدل Setting برای مدیریت تنظیمات سیستم

این مدل شامل فیلدهای زیر است:
- ID: شناسه یکتای تنظیم
- Key: کلید یکتای تنظیم (مثلاً max_login_attempts)
- Value: مقدار تنظیم
- Description: توضیحات تنظیم
- Category: دسته‌بندی تنظیم (مثلاً security, user_limits)
- Type: نوع مقدار (string, number, boolean)
- IsPublic: آیا این تنظیم عمومی است یا خیر
- CreatedAt: زمان ایجاد
- UpdatedAt: زمان آخرین بروزرسانی
- CreatedBy: کاربر ایجاد کننده
- UpdatedBy: کاربر بروزرسانی کننده
*/

type Setting struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Key         string     `gorm:"type:varchar(255);unique;not null;index" json:"key"`
	Value       string     `gorm:"type:text" json:"value"`
	Description string     `gorm:"type:text" json:"description"`
	Category    string     `gorm:"type:varchar(50);not null;index" json:"category"`
	Type        string     `gorm:"type:varchar(20)" json:"type"`
	IsPublic    bool       `gorm:"default:false" json:"is_public"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy   string     `gorm:"type:varchar(100)" json:"created_by"`
	UpdatedBy   string     `gorm:"type:varchar(100)" json:"updated_by"`
	DeletedAt   *time.Time `gorm:"index" json:"-"`
}

// ErrorResponse برای پاسخ خطاهای API
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse برای پاسخ موفقیت‌آمیز API
type SuccessResponse struct {
	Message string `json:"message"`
}
