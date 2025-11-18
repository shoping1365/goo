package models

import "time"

// UserRole ارتباط کاربر با نقش‌های اضافی (Many-to-Many)
// این جدول امکان تخصیص چند نقش همزمان به کاربر را فراهم می‌کند
// فیلدهای valid_from و valid_to اجازهٔ نقش‌های موقّت را می‌دهند.
type UserRole struct {
	ID        uint       `json:"id" gorm:"primaryKey"`          // کلید اصلی
	UserID    uint       `json:"user_id" gorm:"not null;index"` // کاربر
	RoleID    uint       `json:"role_id" gorm:"not null;index"` // نقش
	ValidFrom *time.Time `json:"valid_from"`                    // شروع اعتبار (nullable)
	ValidTo   *time.Time `json:"valid_to"`                      // پایان اعتبار (nullable)
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// TableName مشخص می‌کند که نام جدول، user_roles باشد
func (UserRole) TableName() string {
	return "user_roles"
}
