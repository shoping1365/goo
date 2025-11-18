package models

import "time"

// SEORedirect نگهدارنده ریدایرکت های سئو (مانند RankMath)
// مثال: /product/something-old -> /product/sku-123/title (code=301)
type SEORedirect struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	SourcePath    string     `json:"source_path" gorm:"type:varchar(500);uniqueIndex;not null"`
	TargetPath    string     `json:"target_path" gorm:"type:varchar(500);not null"`
	Code          int        `json:"code" gorm:"default:301"`
	GroupName     string     `json:"group_name" gorm:"type:varchar(100);index"`
	PageTitle     *string    `json:"page_title" gorm:"type:varchar(200)"`
	RedirectType  string     `json:"redirect_type" gorm:"type:varchar(50);default:'permanent'"`
	VisitCount    int64      `json:"visit_count" gorm:"default:0"`
	LastVisitedAt *time.Time `json:"last_visited_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
