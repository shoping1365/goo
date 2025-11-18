package models

import (
	"time"

	"gorm.io/datatypes"
)

// TrafficLog رکورد هر درخواست برای مانیتورینگ ترافیک را ذخیره می‌کند
// شامل اطلاعات مسیر، وضعیت پاسخ، زمان پاسخ و جزئیات دستگاه/شبکه
type TrafficLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;index"`

	// ارتباط اختیاری با کاربر و سشن
	UserID    *uint  `json:"user_id" gorm:"index"`
	SessionID string `json:"session_id" gorm:"type:varchar(255);index"`

	// شبکه و مسیر درخواست
	IPAddress      string `json:"ip_address" gorm:"type:varchar(45);index"`
	RequestMethod  string `json:"request_method" gorm:"type:varchar(10);not null"`
	RequestPath    string `json:"request_path" gorm:"type:varchar(1000);not null;index"`
	StatusCode     int16  `json:"status_code" gorm:"type:smallint;not null;index"`
	ResponseTimeMs int    `json:"response_time_ms" gorm:"not null;default:0"`

	// مشخصات دستگاه/سیستم عامل/مرورگر
	UserAgent  string `json:"user_agent" gorm:"type:text"`
	Browser    string `json:"browser" gorm:"type:varchar(50)"`
	DeviceType string `json:"device_type" gorm:"type:varchar(20)"`
	OS         string `json:"os" gorm:"type:varchar(50)"`
	Hostname   string `json:"hostname" gorm:"type:varchar(255)"`

	// داده‌های تکمیلی
	AdType       string         `json:"ad_type" gorm:"type:varchar(20)"`
	City         string         `json:"city" gorm:"type:varchar(100)"`
	CountryCode  string         `json:"country_code" gorm:"type:char(2)"`
	IsSuspicious bool           `json:"is_suspicious" gorm:"not null;default:false;index"`
	Meta         datatypes.JSON `json:"meta" gorm:"type:jsonb;default:'{}'"`
}

func (TrafficLog) TableName() string { return "traffic_logs" }
