package models

import (
	"time"

	"gorm.io/datatypes"
)

// UserBankAccount اطلاعات بانکی کاربر
// این ساختار منطبق با جدول user_bank_accounts است
type UserBankAccount struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UserID          uint           `json:"user_id" gorm:"not null;index"`
	BankName        string         `json:"bank_name" gorm:"size:100"`
	CardNumberLast4 string         `json:"card_number_last4" gorm:"size:4"`
	CardToken       string         `json:"card_token" gorm:"size:255;index"`
	IBAN            string         `json:"iban" gorm:"size:26;index"`
	AccountNumber   string         `json:"account_number" gorm:"size:50"`
	IsDefault       bool           `json:"is_default" gorm:"default:false;index"`
	VerifiedAt      *time.Time     `json:"verified_at" gorm:"index"`
	Status          string         `json:"status" gorm:"size:20;default:'active';index"`
	Metadata        datatypes.JSON `json:"metadata" gorm:"type:jsonb;default:'{}'"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

func (UserBankAccount) TableName() string { return "user_bank_accounts" }
