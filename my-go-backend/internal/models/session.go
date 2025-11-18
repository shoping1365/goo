package models

import (
	"time"
)

// Session - جلسات کاربری (بروزرسانی شده برای احراز هویت بهینه)
type Session struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"index"`
	Token     string    `json:"-" gorm:"type:varchar(255);uniqueIndex"` // مخفی در JSON
	IPAddress string    `json:"ip_address" gorm:"type:varchar(45)"`
	UserAgent string    `json:"user_agent" gorm:"type:varchar(255)"`
	ExpiresAt time.Time `json:"expires_at" gorm:"index"`
	CreatedAt time.Time `json:"created_at"`
	
	// فیلدهای جدید برای احراز هویت پیشرفته
	Mobile       string     `json:"mobile" gorm:"type:varchar(20)"`           // شماره موبایل
	AuthMethod   string     `json:"auth_method" gorm:"type:varchar(50);default:'password'"` // روش احراز
	IsVerified   bool       `json:"is_verified" gorm:"default:true"`          // تایید شده؟
	VerifiedAt   *time.Time `json:"verified_at"`                              // زمان تایید
	IsActive     bool       `json:"is_active" gorm:"default:true"`            // فعال؟
	LastUsedAt   *time.Time `json:"last_used_at"`                             // آخرین استفاده
	
	// روابط
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName تعیین نام جدول
func (Session) TableName() string {
	return "sessions"
}

// IsValid بررسی اعتبار جلسه
func (s *Session) IsValid() bool {
	return s.IsActive && s.ExpiresAt.After(time.Now())
}

// UpdateLastUsed بروزرسانی زمان آخرین استفاده
func (s *Session) UpdateLastUsed() {
	now := time.Now()
	s.LastUsedAt = &now
}