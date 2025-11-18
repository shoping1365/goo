package models

import "time"

// مدل پترن‌های پیامک
// این مدل برای مدیریت الگوهای پیامک در سیستم استفاده می‌شود
// هر پترن متعلق به یک درگاه پیامک است

type SMSPattern struct {
	ID              uint      `gorm:"primaryKey" json:"id"`                                       // کلید اصلی
	Type            string    `gorm:"size:20;default:'sms'" json:"type"`                          // نوع کانال (sms/email/push)
	Scope           string    `gorm:"size:20;default:'customer'" json:"scope"`                    // دریافت‌کننده (manager/customer)
	Feature         string    `gorm:"size:50" json:"feature"`                                     // هدف/آیتم (auth_otp, order_confirmation, ...)
	Name            string    `gorm:"size:255;not null" json:"name"`                              // نام نمایشی پترن
	PatternCode     string    `gorm:"size:100;not null;uniqueIndex" json:"pattern_code"`          // کد پترن از درگاه
	FixedID         uint      `gorm:"not null;index" json:"fixed_id"`                             // ID ثابت پترن - اختیاری، اگر ارسال نشود خودکار تولید می‌شود
	Description     string    `gorm:"size:500" json:"description"`                                // توضیحات پترن
	Variables       string    `gorm:"type:text" json:"variables"`                                 // متغیرهای پترن (JSON)
	MessageTemplate string    `gorm:"column:template;type:text;not null" json:"message_template"` // قالب پیام
	GatewayID       uint      `gorm:"not null;index" json:"gateway_id"`                           // کلید خارجی به sms_gateways
	Status          string    `gorm:"size:20;default:'active'" json:"status"`                     // وضعیت (active/inactive)
	UsageCount      int       `gorm:"default:0" json:"usage_count"`                               // تعداد استفاده
	CreatedAt       time.Time `json:"created_at"`                                                 // زمان ایجاد
	UpdatedAt       time.Time `json:"updated_at"`                                                 // زمان آخرین ویرایش

	// ارتباط با درگاه پیامک
	Gateway SMSGateway `gorm:"foreignKey:GatewayID" json:"gateway,omitempty"`
}

// TableName تعیین نام جدول برای مدل SMSPattern
func (SMSPattern) TableName() string {
	return "sms_patterns"
}

// مدل برای ایجاد/ویرایش پترن
type SMSPatternForm struct {
	Type            string `json:"type"`
	Scope           string `json:"scope"`
	Feature         string `json:"feature"`
	Name            string `json:"name" binding:"required"`
	PatternCode     string `json:"pattern_code" binding:"required"`
	FixedID         uint   `json:"fixed_id,omitempty"` // ID ثابت پترن - اختیاری، اگر ارسال نشود خودکار تولید می‌شود
	Description     string `json:"description"`
	Variables       string `json:"variables"`
	MessageTemplate string `json:"message_template"`
	GatewayID       uint   `json:"gateway_id" binding:"required"`
	Status          string `json:"status" binding:"required"`
}

// مدل برای تست پترن
type SMSPatternTest struct {
	PatternID uint              `json:"pattern_id" binding:"required"`
	Phone     string            `json:"phone" binding:"required"`
	TestData  map[string]string `json:"test_data"`
}

// مدل برای ارسال پیامک با پترن
type SMSSendRequest struct {
	PatternCode string            `json:"pattern_code"` // کد پترن (اختیاری برای ارسال مستقیم)
	Mobile      string            `json:"mobile" binding:"required"`
	GatewayID   uint              `json:"gateway_id"`
	Scope       string            `json:"scope"`   // دریافت‌کننده (manager/customer)
	Feature     string            `json:"feature"` // هدف/آیتم (auth_otp, order_confirmation, ...)
	Variables   map[string]string `json:"variables"`
	Message     string            `json:"message"` // متن پیام برای ارسال مستقیم
}

// مدل پاسخ ارسال پیامک
type SMSSendResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		MessageID string `json:"message_id,omitempty"`
		Cost      int    `json:"cost,omitempty"`
		Balance   int    `json:"balance,omitempty"`
	} `json:"data,omitempty"`
}
