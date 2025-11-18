package models

import "time"

// MediaDeleteLog ثبت لاگ حذف فایل‌های رسانه به‌همراه کاربر حذف‌کننده.
type MediaDeleteLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	RelPath   string    `gorm:"type:varchar(255);index" json:"rel_path"`
	DeletedBy *uint     `gorm:"index" json:"deleted_by"` // user_id (nullable)
	Note      string    `gorm:"type:text" json:"note"`
	CreatedAt time.Time `json:"created_at"`
}
