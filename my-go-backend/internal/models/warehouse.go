package models

import "time"

// Warehouse نمایانگر یک انبار فیزیکی برای نگهداری موجودی کالاها است
// این ساختار پایه برای مدیریت چندانباره می‌باشد.
type Warehouse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Code      string    `json:"code" gorm:"type:varchar(64);uniqueIndex;not null"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	City      string    `json:"city" gorm:"type:varchar(100)"`
	Address   string    `json:"address" gorm:"type:varchar(500)"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	IsDefault bool      `json:"is_default" gorm:"default:false"`
	Priority  int       `json:"priority" gorm:"default:100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
