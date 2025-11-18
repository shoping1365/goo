package models

import (
	"time"
)

// UserBankingInfo - مدل اطلاعات بانکی کاربران
// این مدل برای ذخیره اطلاعات بانکی کاربران شامل شماره کارت، نام بانک، شماره حساب و شماره شبا استفاده می‌شود
type UserBankingInfo struct {
	ID uint `json:"id" gorm:"primaryKey"`

	// شناسه کاربر
	UserID uint `json:"user_id" gorm:"not null;index"`

	// اطلاعات بانکی
	BankName          string `json:"bank_name" gorm:"type:varchar(100);not null"`     // نام بانک
	CardNumber        string `json:"card_number" gorm:"type:varchar(20);not null"`    // شماره کارت
	AccountNumber     string `json:"account_number" gorm:"type:varchar(50);not null"` // شماره حساب
	ShebaNumber       string `json:"sheba_number" gorm:"type:varchar(26)"`   // شماره شبا (اختیاری)
	AccountHolderName string `json:"account_holder_name" gorm:"type:varchar(200)"`    // نام صاحب حساب
	AccountType       string `json:"account_type" gorm:"type:varchar(50)"`            // نوع حساب (جاری، قرض الحسنه، پس انداز و...)

	// وضعیت پیشفرض
	IsDefault bool `json:"is_default" gorm:"default:false"` // آیا این حساب پیشفرض است

	// وضعیت تایید
	IsVerified bool `json:"is_verified" gorm:"default:false"` // آیا اطلاعات تایید شده است

	// اطلاعات تایید
	VerifiedAt       *time.Time `json:"verified_at"`                        // زمان تایید
	VerifiedBy       *uint      `json:"verified_by"`                        // شناسه ادمین تایید کننده
	VerificationNote string     `json:"verification_note" gorm:"type:text"` // یادداشت تایید

	// زمان‌ها
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;index"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// روابط
	User           User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	VerifiedByUser *User `json:"verified_by_user,omitempty" gorm:"foreignKey:VerifiedBy"`
}

// TableName نام جدول در دیتابیس
func (UserBankingInfo) TableName() string {
	return "user_banking_infos"
}

// BankingInfoRequest درخواست ایجاد/ویرایش اطلاعات بانکی
type BankingInfoRequest struct {
	BankName          string `json:"bank_name" binding:"required"`
	CardNumber        string `json:"card_number" binding:"required"`
	AccountNumber     string `json:"account_number" binding:"required"`
	ShebaNumber       string `json:"sheba_number"`
	AccountHolderName string `json:"account_holder_name" binding:"required"`
	AccountType       string `json:"account_type"`
	IsDefault         bool   `json:"is_default"`
}

// BankingInfoResponse پاسخ اطلاعات بانکی
type BankingInfoResponse struct {
	ID                uint       `json:"id"`
	BankName          string     `json:"bank_name"`
	CardNumber        string     `json:"card_number"`
	AccountNumber     string     `json:"account_number"`
	ShebaNumber       string     `json:"sheba_number"`
	AccountHolderName string     `json:"account_holder_name"`
	AccountType       string     `json:"account_type"`
	IsDefault         bool       `json:"is_default"`
	IsVerified        bool       `json:"is_verified"`
	VerifiedAt        *time.Time `json:"verified_at"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
