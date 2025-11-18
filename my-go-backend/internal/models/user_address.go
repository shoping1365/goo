package models

import (
	"time"

	"gorm.io/gorm"
)

// UserAddress مدل آدرس‌های کاربر
// این ساختار باید با جدول user_addresses همخوانی داشته باشد
// (بخش زیادی از فیلدها حذف شده تا ساده بماند؛ در صورت نیاز می‌توان تکمیل کرد)
type UserAddress struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UserID          uint           `json:"user_id" gorm:"not null;index"`
	Label           string         `json:"label" gorm:"type:varchar(100)"`
	Street          string         `json:"street" gorm:"column:street;type:text"`
	PostalCode      string         `json:"postal_code"`
	RecipientName   string         `json:"recipient_name"`
	RecipientMobile string         `json:"recipient_mobile"`
	Phone           string         `json:"phone" gorm:"type:varchar(20)"`
	Province        string         `json:"province" gorm:"type:varchar(100)"`
	City            string         `json:"city" gorm:"type:varchar(100)"`
	ProvinceID      *uint          `json:"province_id" gorm:"index"`
	CityID          *uint          `json:"city_id" gorm:"index"`
	IsDefault       bool           `json:"is_default"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
