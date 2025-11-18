package models

import "time"

// مدل درگاه پرداخت
// هر درگاه پرداخت می‌تواند تراکنش‌های مختلفی داشته باشد
// حذف رکوردها به صورت سخت انجام می‌شود

type PaymentGateway struct {
	ID           uint      `gorm:"primaryKey" json:"id"`                       // کلید اصلی
	Name         string    `gorm:"size:100;not null" json:"name"`              // نام فارسی درگاه
	EnglishName  string    `gorm:"size:100;not null" json:"english_name"`      // نام انگلیسی درگاه
	Type         string    `gorm:"size:50;not null" json:"type"`               // نوع درگاه (iranian, international, digital)
	Status       string    `gorm:"size:20;default:'inactive'" json:"status"`   // وضعیت (active, inactive, maintenance)
	Icon         string    `gorm:"size:10" json:"icon"`                        // آیکون درگاه
	Color        string    `gorm:"size:20;default:'bg-blue-500'" json:"color"` // رنگ درگاه
	Fee          float64   `gorm:"default:0" json:"fee"`                       // کارمزد (درصد)
	MinAmount    int64     `gorm:"default:0" json:"min_amount"`                // حداقل مبلغ (تومان)
	MaxAmount    int64     `gorm:"default:0" json:"max_amount"`                // حداکثر مبلغ (تومان)
	Priority     int       `gorm:"default:1" json:"priority"`                  // ترتیب اولویت
	IsTestMode   bool      `gorm:"default:true" json:"is_test_mode"`           // حالت تست
	Description  string    `gorm:"size:500" json:"description"`                // توضیحات
	Instructions string    `gorm:"type:text" json:"instructions"`              // دستورالعمل‌های استفاده
	CreatedAt    time.Time `json:"created_at"`                                 // زمان ایجاد
	UpdatedAt    time.Time `json:"updated_at"`                                 // زمان آخرین ویرایش

	// تنظیمات API
	ApiKeys      PaymentGatewayApiKeys   `gorm:"embedded" json:"api_keys"`      // کلیدهای API
	ApiEndpoints PaymentGatewayEndpoints `gorm:"embedded" json:"api_endpoints"` // آدرس‌های API
	TerminalId   string                  `gorm:"size:100" json:"terminal_id"`   // شماره ترمینال
	MerchantId   string                  `gorm:"size:100" json:"merchant_id"`   // شماره پذیرنده

	// آمار
	TodayTransactions int64 `gorm:"default:0" json:"today_transactions"` // تعداد تراکنش‌های امروز
	TodayRevenue      int64 `gorm:"default:0" json:"today_revenue"`      // درآمد امروز (تومان)
	TotalTransactions int64 `gorm:"default:0" json:"total_transactions"` // کل تراکنش‌ها
	TotalRevenue      int64 `gorm:"default:0" json:"total_revenue"`      // کل درآمد (تومان)

	// تنظیمات اضافی
	Settings PaymentGatewaySettings `gorm:"embedded" json:"settings"` // تنظیمات اضافی
}

// تنظیمات کلیدهای API درگاه پرداخت
type PaymentGatewayApiKeys struct {
	PublicKey  string `gorm:"size:255" json:"public_key"`  // کلید عمومی
	PrivateKey string `gorm:"size:255" json:"private_key"` // کلید خصوصی
	TestKey    string `gorm:"size:255" json:"test_key"`    // کلید تست
	SecretKey  string `gorm:"size:255" json:"secret_key"`  // کلید محرمانه
}

// آدرس‌های API درگاه پرداخت
type PaymentGatewayEndpoints struct {
	Payment      string `gorm:"size:255" json:"payment"`      // آدرس API پرداخت
	Verification string `gorm:"size:255" json:"verification"` // آدرس API تایید
	Refund       string `gorm:"size:255" json:"refund"`       // آدرس API بازگشت وجه
	Balance      string `gorm:"size:255" json:"balance"`      // آدرس API موجودی
	Callback     string `gorm:"size:255" json:"callback"`     // آدرس callback
}

// تنظیمات اضافی درگاه پرداخت
type PaymentGatewaySettings struct {
	AutoRedirect       bool   `gorm:"default:true" json:"auto_redirect"`        // ریدایرکت خودکار
	ShowGatewayName    bool   `gorm:"default:true" json:"show_gateway_name"`    // نمایش نام درگاه
	RequireCardInfo    bool   `gorm:"default:false" json:"require_card_info"`   // نیاز به اطلاعات کارت
	SupportInstallment bool   `gorm:"default:false" json:"support_installment"` // پشتیبانی از اقساط
	MaxInstallments    int    `gorm:"default:1" json:"max_installments"`        // حداکثر تعداد اقساط
	Currency           string `gorm:"size:10;default:'IRR'" json:"currency"`    // واحد پول
	Language           string `gorm:"size:10;default:'fa'" json:"language"`     // زبان
}

// مدل تراکنش‌های درگاه پرداخت
type PaymentTransaction struct {
	ID              uint      `gorm:"primaryKey" json:"id"`                       // کلید اصلی
	GatewayID       uint      `gorm:"not null;index" json:"gateway_id"`           // کلید خارجی به payment_gateways
	OrderID         string    `gorm:"size:100;not null;index" json:"order_id"`    // شماره سفارش
	TransactionID   string    `gorm:"size:100;uniqueIndex" json:"transaction_id"` // شماره تراکنش
	Amount          int64     `gorm:"not null" json:"amount"`                     // مبلغ (تومان)
	Status          string    `gorm:"size:20;not null" json:"status"`             // وضعیت (pending, success, failed, refunded)
	PaymentMethod   string    `gorm:"size:50" json:"payment_method"`              // روش پرداخت
	Description     string    `gorm:"size:500" json:"description"`                // توضیحات تراکنش
	GatewayToken    string    `gorm:"size:255" json:"gateway_token"`              // توکن درگاه
	GatewayType     string    `gorm:"size:50" json:"gateway_type"`                // نوع درگاه
	CardNumber      string    `gorm:"size:20" json:"card_number"`                 // شماره کارت (مخفی)
	CardHolder      string    `gorm:"size:100" json:"card_holder"`                // صاحب کارت
	GatewayResponse string    `gorm:"type:text" json:"gateway_response"`          // پاسخ درگاه
	ErrorMessage    string    `gorm:"size:500" json:"error_message"`              // پیام خطا
	CreatedAt       time.Time `json:"created_at"`                                 // زمان ایجاد
	UpdatedAt       time.Time `json:"updated_at"`                                 // زمان آخرین ویرایش

	// ارتباط با درگاه
	Gateway PaymentGateway `gorm:"foreignKey:GatewayID" json:"gateway"`
}

// مدل لاگ‌های درگاه پرداخت
type PaymentGatewayLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`             // کلید اصلی
	GatewayID uint      `gorm:"not null;index" json:"gateway_id"` // کلید خارجی به payment_gateways
	Level     string    `gorm:"size:20;not null" json:"level"`    // سطح لاگ (info, warning, error)
	Message   string    `gorm:"size:500;not null" json:"message"` // پیام لاگ
	Details   string    `gorm:"type:text" json:"details"`         // جزئیات اضافی
	IPAddress string    `gorm:"size:45" json:"ip_address"`        // آدرس IP
	UserAgent string    `gorm:"size:500" json:"user_agent"`       // User Agent
	CreatedAt time.Time `json:"created_at"`                       // زمان ایجاد

	// ارتباط با درگاه
	Gateway PaymentGateway `gorm:"foreignKey:GatewayID" json:"gateway"`
}
