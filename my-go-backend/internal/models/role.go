package models

import (
	"time"
)

// Role مدل نقش‌های سیستم
// این مدل برای تعریف نقش‌های مختلف کاربران در سیستم استفاده می‌شود
type Role struct {
	ID          uint      `json:"id" gorm:"primaryKey"`                               // کلید اصلی
	Name        string    `json:"name" gorm:"type:varchar(100);not null;uniqueIndex"` // نام نقش
	DisplayName string    `json:"display_name" gorm:"type:varchar(150);not null"`     // نام نمایشی نقش
	Description string    `json:"description" gorm:"type:text"`                       // توضیحات نقش
	IsActive    bool      `json:"is_active" gorm:"default:true;index"`                // فعال/غیرفعال بودن
	IsSystem    bool      `json:"is_system" gorm:"default:false;index"`               // آیا نقش سیستمی است (غیرقابل حذف)
	Priority    int       `json:"priority" gorm:"default:0;index"`                    // اولویت نقش (برای نمایش)
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"` // حذف شد

	// فیلدهای محاسبه شده (برای API)
	UsersCount       int `json:"users_count" gorm:"-"`       // تعداد کاربران این نقش
	PermissionsCount int `json:"permissions_count" gorm:"-"` // تعداد دسترسی‌های این نقش

	// روابط
	Users       []User       `json:"users,omitempty" gorm:"foreignKey:RoleID"`
	Permissions []Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions;"`
}

// Permission مدل دسترسی‌های سیستم
// این مدل برای تعریف دسترسی‌های مختلف در سیستم استفاده می‌شود
type Permission struct {
	ID          uint      `json:"id" gorm:"primaryKey"`                               // کلید اصلی
	Name        string    `json:"name" gorm:"type:varchar(100);not null;uniqueIndex"` // نام دسترسی
	DisplayName string    `json:"display_name" gorm:"type:varchar(150);not null"`     // نام نمایشی دسترسی
	Description string    `json:"description" gorm:"type:text"`                       // توضیحات دسترسی
	Module      string    `json:"module" gorm:"type:varchar(50);not null;index"`      // ماژول اصلی (products, users, orders, ...)
	SubModule   string    `json:"sub_module" gorm:"type:varchar(50);index"`           // زیرماژول (categories, attributes, ...)
	Action      string    `json:"action" gorm:"type:varchar(50);not null;index"`      // عملیات (view, create, update, delete, write)
	Resource    string    `json:"resource" gorm:"type:varchar(50);not null;index"`    // منبع (product, user, order, ...)
	IsActive    bool      `json:"is_active" gorm:"default:true;index"`                // فعال/غیرفعال بودن
	Priority    int       `json:"priority" gorm:"default:0;index"`                    // اولویت نمایش
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"` // حذف شد

	// روابط
	Roles []Role `json:"roles,omitempty" gorm:"many2many:role_permissions;"`
}

// RolePermission جدول ارتباطی بین نقش‌ها و دسترسی‌ها
// این جدول برای ذخیره ارتباطات many-to-many بین نقش‌ها و دسترسی‌ها استفاده می‌شود
type RolePermission struct {
	ID           uint       `json:"id" gorm:"primaryKey"`                // کلید اصلی
	RoleID       uint       `json:"role_id" gorm:"not null;index"`       // کلید خارجی نقش
	PermissionID uint       `json:"permission_id" gorm:"not null;index"` // کلید خارجی دسترسی
	GrantedBy    uint       `json:"granted_by" gorm:"index"`             // کاربری که این دسترسی را اعطا کرده
	GrantedAt    time.Time  `json:"granted_at" gorm:"autoCreateTime"`    // زمان اعطای دسترسی
	ExpiresAt    *time.Time `json:"expires_at"`                          // تاریخ انقضای دسترسی (nullable)
	IsActive     bool       `json:"is_active" gorm:"default:true;index"` // فعال/غیرفعال بودن
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`

	// روابط
	Role          Role       `json:"role" gorm:"foreignKey:RoleID"`
	Permission    Permission `json:"permission" gorm:"foreignKey:PermissionID"`
	GrantedByUser User       `json:"granted_by_user" gorm:"foreignKey:GrantedBy"`
}

// TableName نام جدول role_permissions را مشخص می‌کند
func (RolePermission) TableName() string {
	return "role_permissions"
}

// CreateRoleRequest درخواست ایجاد نقش جدید
type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`         // نام نقش
	DisplayName string `json:"display_name" binding:"required"` // نام نمایشی
	Description string `json:"description"`                     // توضیحات
	Permissions []uint `json:"permissions"`                     // لیست ID دسترسی‌ها
}

// UpdateRoleRequest درخواست به‌روزرسانی نقش
type UpdateRoleRequest struct {
	Name        string `json:"name"`         // نام نقش (اختیاری برای تغییر)
	DisplayName string `json:"display_name"` // نام نمایشی
	Description string `json:"description"`  // توضیحات
	IsActive    *bool  `json:"is_active"`    // فعال/غیرفعال بودن
	Permissions []uint `json:"permissions"`  // لیست ID دسترسی‌ها
}

// CreatePermissionRequest درخواست ایجاد دسترسی جدید
type CreatePermissionRequest struct {
	Name        string `json:"name" binding:"required"`         // نام دسترسی
	DisplayName string `json:"display_name" binding:"required"` // نام نمایشی
	Description string `json:"description"`                     // توضیحات
	Module      string `json:"module" binding:"required"`       // ماژول اصلی
	SubModule   string `json:"sub_module"`                      // زیرماژول
	Action      string `json:"action" binding:"required"`       // عملیات
	Resource    string `json:"resource" binding:"required"`     // منبع
	Priority    int    `json:"priority"`                        // اولویت
}

// UpdatePermissionRequest درخواست به‌روزرسانی دسترسی
type UpdatePermissionRequest struct {
	DisplayName string `json:"display_name"` // نام نمایشی
	Description string `json:"description"`  // توضیحات
	IsActive    *bool  `json:"is_active"`    // فعال/غیرفعال بودن
	Priority    *int   `json:"priority"`     // اولویت
}

// PermissionGroup گروه‌بندی دسترسی‌ها برای نمایش در UI
type PermissionGroup struct {
	Module     string      `json:"module"`      // نام ماژول
	SubModules []SubModule `json:"sub_modules"` // زیرماژول‌ها
}

// SubModule زیرماژول با دسترسی‌های مختلف
type SubModule struct {
	Name        string       `json:"name"`         // نام زیرماژول
	DisplayName string       `json:"display_name"` // نام نمایشی
	Permissions []Permission `json:"permissions"`  // دسترسی‌های این زیرماژول
}

// RolePermissionRequest درخواست تخصیص دسترسی‌ها به نقش
type RolePermissionRequest struct {
	RoleID       uint       `json:"role_id" binding:"required"`       // ID نقش
	PermissionID uint       `json:"permission_id" binding:"required"` // ID دسترسی
	GrantedBy    uint       `json:"granted_by"`                       // کاربر اعطاکننده
	ExpiresAt    *time.Time `json:"expires_at"`                       // تاریخ انقضا
}

// BulkRolePermissionRequest درخواست تخصیص دسترسی‌های گروهی
type BulkRolePermissionRequest struct {
	RoleID        uint       `json:"role_id" binding:"required"`        // ID نقش
	PermissionIDs []uint     `json:"permission_ids" binding:"required"` // لیست ID دسترسی‌ها
	GrantedBy     uint       `json:"granted_by"`                        // کاربر اعطاکننده
	ExpiresAt     *time.Time `json:"expires_at"`                        // تاریخ انقضا
}
