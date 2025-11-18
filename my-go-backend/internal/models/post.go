package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Post مدل نوشته‌ها
type Post struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	Title         string     `json:"title" gorm:"size:255;not null"`
	Slug          string     `json:"slug" gorm:"size:255;not null;uniqueIndex"`
	Excerpt       string     `json:"excerpt" gorm:"type:text"`
	Content       string     `json:"content" gorm:"type:text"`
	Status        string     `json:"status" gorm:"type:varchar(20);default:'draft'"`
	CategoryID    *uint      `json:"category_id" gorm:"index"`
	FeaturedImage string     `json:"featured_image" gorm:"size:255"`
	AuthorID      uint       `json:"author_id" gorm:"not null;index"`
	PublishedAt   *time.Time `json:"published_at"`
	// زمان‌بندی انتشار پست
	ScheduledAt     *time.Time `json:"scheduled_at" gorm:"column:scheduled_at"`
	MetaTitle       string     `json:"meta_title" gorm:"size:255"`
	MetaKeywords    string     `json:"meta_keywords" gorm:"size:255"`
	MetaDescription string     `json:"meta_description" gorm:"size:255"`

	// فیلدهای SEO جدید
	RobotsMeta string `json:"robots_meta" gorm:"type:varchar(50);default:'index,follow'"` // robots meta tag content

	// فیلدهای مخصوص مقالات آموزشی
	EstimatedTime int `json:"estimated_time" gorm:"default:0"` // زمان تخمینی به دقیقه

	// فیلدهای مخصوص مقالات خبری
	Location  string     `json:"location" gorm:"type:varchar(255)"`
	EventDate *time.Time `json:"event_date"`

	// فیلدهای مخصوص مقالات بررسی محصول
	Pros string `json:"pros" gorm:"type:text"` // مزایا
	Cons string `json:"cons" gorm:"type:text"` // معایب
	// ProductPrice و Rating به صورت پویا تولید می‌شوند

	// فیلدهای عمومی
	Tags                datatypes.JSON `json:"tags" gorm:"type:jsonb"` // آرایه JSON از تگ‌ها
	Language            string         `json:"language" gorm:"type:varchar(10);default:'fa'"`
	IsAccessibleForFree bool           `json:"is_accessible_for_free" gorm:"default:true"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// روابط
	Category PostCategory `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Author   User         `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
}

// TableName نام جدول
func (Post) TableName() string {
	return "posts"
}

// AfterCreate صف همگام‌سازی برای ایجاد یا به‌روزرسانی محتوای جستجو را درج می‌کند.
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return enqueueSearchEvent(tx, SearchResourcePost, uint64(p.ID), SearchOpUpsert)
}

// AfterUpdate صف همگام‌سازی برای ایندکس محتوا را درج می‌کند.
func (p *Post) AfterUpdate(tx *gorm.DB) error {
	return enqueueSearchEvent(tx, SearchResourcePost, uint64(p.ID), SearchOpUpsert)
}

// AfterDelete حذف سند از ایندکس جستجو را صف می‌کند.
func (p *Post) AfterDelete(tx *gorm.DB) error {
	return enqueueSearchEvent(tx, SearchResourcePost, uint64(p.ID), SearchOpDelete)
}
