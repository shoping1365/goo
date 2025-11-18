package models

import "time"

// NotificationLog ثبت لاگ ارسال اعلان‌ها (موفق یا خطا)
type NotificationLog struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Channel      string    `json:"channel" gorm:"size:20;index"` // sms | email | push
	Type         string    `json:"type" gorm:"size:20;index"`    // stock | discount
	Notification uint      `json:"notification" gorm:"index"`    // reference id
	ProductID    uint      `json:"product_id" gorm:"index"`
	UserID       uint      `json:"user_id" gorm:"index"`
	Status       string    `json:"status" gorm:"size:20;index"` // success | failed
	ErrorCode    string    `json:"error_code" gorm:"size:50"`
	ErrorMessage string    `json:"error_message" gorm:"size:500"`
	CreatedAt    time.Time `json:"created_at"`
}

func (NotificationLog) TableName() string { return "notification_logs" }
