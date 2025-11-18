package models

import (
	"time"
)

type AdminSetting struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    *uint     `gorm:"index" json:"user_id"`
	Key       string    `gorm:"type:varchar(100);not null;index:idx_user_key,unique" json:"key"`
	Value     string    `gorm:"type:text" json:"value"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName نام جدول را مشخص می‌کند
func (AdminSetting) TableName() string {
	return "admin_settings"
}
