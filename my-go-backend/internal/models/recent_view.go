package models

import (
	"time"
)

// RecentView represents a user's recent product view
type RecentView struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id" gorm:"not null;index"`
	ProductID       int       `json:"product_id" gorm:"not null;index"`
	ViewedAt        time.Time `json:"viewed_at" gorm:"autoCreateTime;index"`
	DurationSeconds int       `json:"duration_seconds" gorm:"default:0;comment:مدت زمان ماندن در صفحه"`
	ViewCount       int       `json:"view_count" gorm:"default:1;comment:تعداد دفعات بازدید"`
	LastUpdatedAt   time.Time `json:"last_updated_at" gorm:"autoUpdateTime"`
	
	// Device and Browser Information
	Browser        string `json:"browser" gorm:"type:varchar(100);comment:نام مرورگر"`
	BrowserVersion string `json:"browser_version" gorm:"type:varchar(50);comment:نسخه مرورگر"`
	OS             string `json:"os" gorm:"type:varchar(100);comment:سیستم عامل"`
	OSVersion      string `json:"os_version" gorm:"type:varchar(50);comment:نسخه سیستم عامل"`
	DeviceType     string `json:"device_type" gorm:"type:varchar(50);index;comment:نوع دستگاه"`
	DeviceModel    string `json:"device_model" gorm:"type:varchar(100);comment:مدل دستگاه"`
	IPAddress      string `json:"ip_address" gorm:"type:varchar(45);comment:آدرس IP"`
	UserAgent      string `json:"user_agent" gorm:"type:text;comment:User-Agent کامل"`

	// Relations
	User    User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

// RecentViewResponse represents the response structure for recent views
type RecentViewResponse struct {
	ID              uint      `json:"id"`
	ProductID       int       `json:"product_id"`
	ViewedAt        time.Time `json:"viewed_at"`
	DurationSeconds int       `json:"duration_seconds"`
	ViewCount       int       `json:"view_count"`
	LastUpdatedAt   time.Time `json:"last_updated_at"`
	
	// Device and Browser Information
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	OS             string `json:"os"`
	OSVersion      string `json:"os_version"`
	DeviceType     string `json:"device_type"`
	DeviceModel    string `json:"device_model"`
	IPAddress      string `json:"ip_address"`
	UserAgent      string `json:"user_agent"`
	
	Product struct {
		ID       int     `json:"id"`
		UUID     string  `json:"uuid"`
		Slug     string  `json:"slug"`
		Name     string  `json:"name"`
		Image    string  `json:"image"`
		ImageURL string  `json:"image_url"`
		Price    float64 `json:"price"`
		OldPrice float64 `json:"old_price"`
	} `json:"product"`
}

// UpdateDurationRequest represents the request to update view duration
type UpdateDurationRequest struct {
	Duration int `json:"duration" binding:"required,min=1,max=3600"` // حداکثر 1 ساعت
}

// TableName returns the table name for RecentView
func (RecentView) TableName() string {
	return "recent_views"
}
