package models

import "time"

// مدل درگاه پیامک
// هر درگاه می‌تواند چندین پترن داشته باشد
// حذف رکوردها به صورت سخت انجام می‌شود

type SMSGateway struct {
	ID           uint                `gorm:"primaryKey" json:"id"`                 // کلید اصلی
	Name         string              `gorm:"size:100" json:"name"`                 // نام سفارشی درگاه
	Type         string              `gorm:"size:50;not null" json:"type"`         // نوع درگاه (مثلاً meli_payamak)
	SenderNumber string              `gorm:"size:30" json:"sender_number"`         // شماره ارسال‌کننده
	ApiURL       string              `gorm:"size:255" json:"api_url"`              // آدرس API
	ApiKey       string              `gorm:"size:255" json:"api_key"`              // کلید API
	Username     string              `gorm:"size:100" json:"username"`             // نام کاربری برای احراز هویت
	Password     string              `gorm:"size:100" json:"password"`             // رمز عبور برای احراز هویت
	PatternBased bool                `gorm:"default:false" json:"pattern_based"`   // ارسال بر اساس پترن
	IsActive     bool                `gorm:"default:true" json:"is_active"`        // فعال/غیرفعال
	Priority     int                 `gorm:"default:1" json:"priority"`            // ترتیب اولویت سوییچ
	CreatedAt    time.Time           `json:"created_at"`                           // زمان ایجاد
	UpdatedAt    time.Time           `json:"updated_at"`                           // زمان آخرین ویرایش
	Patterns     []SMSGatewayPattern `gorm:"foreignKey:GatewayID" json:"patterns"` // لیست پترن‌های این درگاه
}

// مدل پترن‌های هر درگاه پیامک
// هر پترن متعلق به یک درگاه است
// حذف رکوردها به صورت سخت انجام می‌شود

type SMSGatewayPattern struct {
	ID          uint      `gorm:"primaryKey"`        // کلید اصلی
	GatewayID   uint      `gorm:"not null;index"`    // کلید خارجی به sms_gateways
	PatternCode string    `gorm:"size:100;not null"` // کد یا مقدار پترن
	Description string    `gorm:"size:255"`          // توضیح (اختیاری)
	CreatedAt   time.Time // زمان ایجاد
	UpdatedAt   time.Time // زمان آخرین ویرایش
}

func (SMSGatewayPattern) TableName() string { return "sms_gateway_patterns" }

// مدل پیام‌های دریافتی (Inbox) - مطابق با ساختار رسمی IPPanel
type InboxMessage struct {
	To        string    `json:"to"`         // شماره گیرنده
	Message   string    `json:"message"`    // متن پیام
	From      string    `json:"from"`       // شماره فرستنده
	CreatedAt time.Time `json:"created_at"` // زمان ایجاد
	Type      string    `json:"type"`       // نوع پیام
}

// مدل پاسخ API inbox
type InboxResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message,omitempty"`
	Data    InboxData `json:"data,omitempty"`
}

// مدل داده‌های inbox
type InboxData struct {
	Messages   []InboxMessage `json:"messages"`
	Total      int            `json:"total"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"total_pages"`
}
