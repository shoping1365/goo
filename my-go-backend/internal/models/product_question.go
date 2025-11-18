package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Question struct {
	ID           uint       `json:"id"`
	UserID       uint       `json:"user_id"`
	User         User       `json:"user" gorm:"foreignKey:UserID"`
	ProductID    uint       `json:"product_id"`
	Product      Product    `json:"product" gorm:"foreignKey:ProductID"`
	CategoryID   *uint      `json:"category_id"`
	QuestionText string     `json:"question"`
	IsAnonymous  bool       `json:"is_anonymous"`
	Status       string     `json:"status"`
	RejectReason string     `json:"reject_reason"`
	AdminReply   string     `json:"admin_reply"`
	AdminReplyAt *time.Time `json:"admin_reply_at"`
	IPAddress    string     `json:"ip_address"`
	// UserAgent و DeviceInfo برای نمایش اطلاعات دستگاه کاربر در پنل ادمین
	UserAgent  string         `json:"user_agent" gorm:"type:text"`
	DeviceInfo datatypes.JSON `json:"device_info,omitempty" gorm:"type:jsonb"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	EditedAt   *time.Time     `json:"edited_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	GuestName  string         `json:"guest_name" gorm:"type:varchar(100);"`
	GuestPhone string         `json:"guest_phone" gorm:"type:varchar(20);"`
}

type AdminQuestion = Question

func (Question) TableName() string {
	return "product_questions"
}
