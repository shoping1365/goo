package models

import (
	"time"

	"gorm.io/datatypes"
)

// UserWallet کیف پول کاربر
// مطابق جدول user_wallets
type UserWallet struct {
	ID                uint       `json:"id" gorm:"primaryKey"`
	UserID            uint       `json:"user_id" gorm:"not null;uniqueIndex"`
	Balance           float64    `json:"balance" gorm:"type:numeric(18,2);default:0"`
	Currency          string     `json:"currency" gorm:"size:10;default:'IRR'"`
	Status            string     `json:"status" gorm:"size:20;default:'active';index"`
	LastTransactionAt *time.Time `json:"last_transaction_at" gorm:"index"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func (UserWallet) TableName() string { return "user_wallets" }

// WalletTransaction تراکنش کیف پول
// مطابق جدول wallet_transactions
type WalletTransaction struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	WalletID        uint           `json:"wallet_id" gorm:"not null;index"`
	Type            string         `json:"type" gorm:"size:10;not null;index"`
	Method          string         `json:"method" gorm:"size:20;index"`
	Amount          float64        `json:"amount" gorm:"type:numeric(18,2);not null"`
	RelatedType     string         `json:"related_type" gorm:"size:50"`
	RelatedID       *uint          `json:"related_id" gorm:"index"`
	CardNumberLast4 string         `json:"card_number_last4" gorm:"size:4"`
	Gateway         string         `json:"gateway" gorm:"size:50"`
	IP              string         `json:"ip" gorm:"size:45;index"`
	UserAgent       string         `json:"user_agent" gorm:"size:500"`
	Platform        string         `json:"platform" gorm:"size:50"`
	Browser         string         `json:"browser" gorm:"size:50"`
	OSName          string         `json:"os_name" gorm:"size:50"`
	SessionID       *string        `json:"session_id" gorm:"size:255;index"`
	Metadata        datatypes.JSON `json:"metadata" gorm:"type:jsonb;default:'{}'"`
	Status          string         `json:"status" gorm:"size:20;default:'success';index"`
	ErrorCode       string         `json:"error_code" gorm:"size:50"`
	ErrorMessage    string         `json:"error_message" gorm:"size:500"`
	CreatedAt       time.Time      `json:"created_at"`
}

func (WalletTransaction) TableName() string { return "wallet_transactions" }
