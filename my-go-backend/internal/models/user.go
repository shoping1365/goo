package models

import (
	"time"

	"gorm.io/datatypes"
)

// User - مدل کاربر بهینه شده
type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Username     string `json:"username" gorm:"type:varchar(100);uniqueIndex"`
	Email        string `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	Mobile       string `json:"mobile" gorm:"type:varchar(20);uniqueIndex;not null"`
	Name         string `json:"name" gorm:"type:varchar(100)"`
	PasswordHash string `json:"-" gorm:"column:password_hash;type:varchar(255)"`
	RoleID       uint   `json:"role_id" gorm:"not null;index;default:1"` // default customer role
	Status       string `json:"status" gorm:"type:varchar(20);default:'active';index"`

	// Security fields
	EmailVerified    bool       `json:"email_verified" gorm:"default:false;index"`
	MobileVerified   bool       `json:"mobile_verified" gorm:"default:true;index"`
	FailedLoginCount int        `json:"failed_login_count" gorm:"default:0"`
	LockedUntil      *time.Time `json:"locked_until" gorm:"index"`
	LastLoginAt      *time.Time `json:"last_login_at" gorm:"index"`
	LastLoginIP      string     `json:"last_login_ip" gorm:"type:varchar(45)"`

	// Timestamps
	RegisteredAt time.Time `json:"registered_at" gorm:"autoCreateTime;index"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// JSON fields for flexibility (removing legacy fields)
	Preferences      datatypes.JSON `json:"preferences" gorm:"type:jsonb;default:'{}'"`       // UI preferences, notifications, etc.
	SecuritySettings datatypes.JSON `json:"security_settings" gorm:"type:jsonb;default:'{}'"` // 2FA, trusted devices, etc.
	ProfileData      datatypes.JSON `json:"profile_data" gorm:"type:jsonb;default:'{}'"`      // Additional profile information (address, city, etc.)

	// روابط
	UserRole Role `json:"user_role,omitempty" gorm:"foreignKey:RoleID"`
}

type CreateUserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Username string `json:"username" binding:"required,min=3,max=30"`
	Mobile   string `json:"mobile" binding:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Landline string `json:"landline"`
}
