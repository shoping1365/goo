package models

import (
	"time"

	"gorm.io/gorm"
)

// GiftCard مدل اصلی گیفت کارت
type GiftCard struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	Code            string `json:"code" gorm:"uniqueIndex;not null"`
	Amount          int64  `json:"amount" gorm:"not null"` // مبلغ به تومان
	RemainingAmount int64  `json:"remaining_amount" gorm:"not null"`
	Status          string `json:"status" gorm:"not null;default:'active'"` // active, used, expired, inactive
	Type            string `json:"type" gorm:"not null"`                    // digital, physical, hybrid
	Category        string `json:"category" gorm:"not null"`                // birthday, wedding, newyear, corporate, discount, general

	// اطلاعات گیرنده
	RecipientName  string `json:"recipient_name"`
	RecipientEmail string `json:"recipient_email"`
	RecipientPhone string `json:"recipient_phone"`

	// روش تحویل
	DeliveryMethod string    `json:"delivery_method"` // email, sms, both, manual
	DeliveryDate   time.Time `json:"delivery_date"`

	// پیام شخصی
	PersonalMessage string `json:"personal_message"`

	// محدودیت‌های استفاده
	UsageLimit     string `json:"usage_limit"` // once, multiple
	MaxUsageCount  int    `json:"max_usage_count"`
	MinOrderAmount int64  `json:"min_order_amount"`
	MaxOrderAmount int64  `json:"max_order_amount"`

	// آمار استفاده
	UsageCount      int   `json:"usage_count" gorm:"default:0"`
	AvgUsageAmount  int64 `json:"avg_usage_amount"`
	MaxUsageAmount  int64 `json:"max_usage_amount"`
	LastUsageAmount int64 `json:"last_usage_amount"`

	// طراحی
	PrimaryColor   string `json:"primary_color"`
	SecondaryColor string `json:"secondary_color"`
	Font           string `json:"font"`

	// تاریخ‌ها
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	ExpiryDate time.Time      `json:"expiry_date"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// روابط
	CreatedBy uint `json:"created_by" gorm:"not null"`
	User      User `json:"user" gorm:"foreignKey:CreatedBy"`

	// تراکنش‌ها
	Transactions []GiftCardTransaction `json:"transactions,omitempty" gorm:"foreignKey:GiftCardID"`
}

// GiftCardTransaction تراکنش‌های گیفت کارت
type GiftCardTransaction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	GiftCardID  uint      `json:"gift_card_id" gorm:"not null"`
	GiftCard    GiftCard  `json:"gift_card" gorm:"foreignKey:GiftCardID"`
	Amount      int64     `json:"amount" gorm:"not null"`
	Type        string    `json:"type" gorm:"not null"` // usage, refund, activation, deactivation
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   uint      `json:"created_by"`
	User        User      `json:"user" gorm:"foreignKey:CreatedBy"`
}

// GiftCardType انواع گیفت کارت
type GiftCardType struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GiftCardTemplate قالب‌های گیفت کارت
type GiftCardTemplate struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"not null"`
	Description    string    `json:"description"`
	PrimaryColor   string    `json:"primary_color"`
	SecondaryColor string    `json:"secondary_color"`
	Font           string    `json:"font"`
	IsActive       bool      `json:"is_active" gorm:"default:true"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GiftCardStats آمار داشبورد گیفت کارت
type GiftCardStats struct {
	TotalCards       int64   `json:"total_cards"`
	ActiveCards      int64   `json:"active_cards"`
	ExpiringSoon     int64   `json:"expiring_soon"`
	TotalValue       int64   `json:"total_value"`
	NewCardsToday    int64   `json:"new_cards_today"`
	ActivePercentage float64 `json:"active_percentage"`
	ValueGrowth      float64 `json:"value_growth"`
	UsedCards        int64   `json:"used_cards"`
	UsageRate        float64 `json:"usage_rate"`
	AvgUsageAmount   int64   `json:"avg_usage_amount"`
	DeliveredCards   int64   `json:"delivered_cards"`
	DeliveryRate     float64 `json:"delivery_rate"`
	EmailDelivery    float64 `json:"email_delivery"`
	SmsDelivery      float64 `json:"sms_delivery"`
}

// GiftCardActivity فعالیت‌های اخیر
type GiftCardActivity struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	GiftCardID uint      `json:"gift_card_id"`
	GiftCard   GiftCard  `json:"gift_card" gorm:"foreignKey:GiftCardID"`
	Action     string    `json:"action" gorm:"not null"` // created, used, expired, delivered
	Amount     int64     `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  uint      `json:"created_by"`
	User       User      `json:"user" gorm:"foreignKey:CreatedBy"`
}

// TableName نام جدول برای مدل‌ها
func (GiftCard) TableName() string {
	return "gift_cards"
}

func (GiftCardTransaction) TableName() string {
	return "gift_card_transactions"
}

func (GiftCardType) TableName() string {
	return "gift_card_types"
}

func (GiftCardTemplate) TableName() string {
	return "gift_card_templates"
}

func (GiftCardActivity) TableName() string {
	return "gift_card_activities"
}
