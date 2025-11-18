package models

import (
	"time"
)

// TrafficSession مدل برای ذخیره سشن‌های ترافیک و مانیتورینگ
// متفاوت از UserSession در auth.go که برای احراز هویت استفاده می‌شود
type TrafficSession struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      *uint     `json:"user_id" gorm:"index"` // null برای کاربران مهمان
	SessionID   string    `json:"session_id" gorm:"type:varchar(255);uniqueIndex;not null"`
	IPAddress   string    `json:"ip_address" gorm:"type:varchar(45);not null"`
	UserAgent   string    `json:"user_agent" gorm:"type:text"`
	CurrentPage string    `json:"current_page" gorm:"type:varchar(500)"`
	LoginTime   time.Time `json:"login_time" gorm:"not null"`
	LastSeen    time.Time `json:"last_seen" gorm:"not null"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// رابطه با کاربر (در صورت وجود)
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName override for TrafficSession
func (TrafficSession) TableName() string {
	return "traffic_sessions"
}

// TrafficStats آمار ترافیک
type TrafficStats struct {
	OnlineUsers        int64               `json:"online_users"`
	TodayRequests      int64               `json:"today_requests"`
	SuspiciousRequests int64               `json:"suspicious_requests"`
	BlockedAttacks     int64               `json:"blocked_attacks"`
	HourlyTraffic      []HourlyTrafficData `json:"hourly_traffic"`
}

// HourlyTrafficData داده‌های ترافیک ساعتی
type HourlyTrafficData struct {
	Hour  int   `json:"hour"`
	Count int64 `json:"count"`
}

// DailyTrafficData داده‌های ترافیک روزانه
type DailyTrafficData struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

// WeeklyTrafficData داده‌های ترافیک هفتگی
type WeeklyTrafficData struct {
	Week  string `json:"week"`
	Count int64  `json:"count"`
}

// MonthlyTrafficData داده‌های ترافیک ماهانه
type MonthlyTrafficData struct {
	Month string `json:"month"`
	Count int64  `json:"count"`
}

// PaginationInfo اطلاعات صفحه‌بندی
type PaginationInfo struct {
	CurrentPage  int   `json:"current_page"`
	TotalPages   int   `json:"total_pages"`
	TotalItems   int64 `json:"total_items"`
	ItemsPerPage int   `json:"items_per_page"`
	HasNext      bool  `json:"has_next"`
	HasPrev      bool  `json:"has_prev"`
}

// BlockedIP آدرس‌های IP مسدود شده
type BlockedIP struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	IPAddress string     `json:"ip_address" gorm:"type:varchar(45);uniqueIndex;not null"`
	Reason    string     `json:"reason" gorm:"type:text"`
	BlockedBy string     `json:"blocked_by" gorm:"type:varchar(100)"`
	BlockedAt time.Time  `json:"blocked_at" gorm:"not null"`
	ExpiresAt *time.Time `json:"expires_at"`
	IsActive  bool       `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
