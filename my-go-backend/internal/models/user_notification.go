package models

import (
	"time"

	"gorm.io/datatypes"
)

// UserNotification هم‌راستا با مایگریشن 53_user_notifications
type UserNotification struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     uint           `gorm:"not null;index" json:"user_id"`
	Channel    string         `gorm:"size:20;not null;index" json:"channel"` // sms | email | push | system
	Source     string         `gorm:"size:20;not null;index" json:"source"`  // system | admin
	TemplateID *uint          `gorm:"index" json:"template_id,omitempty"`
	Title      string         `gorm:"size:200" json:"title"`
	Body       string         `gorm:"type:text" json:"body"`
	Payload    datatypes.JSON `gorm:"type:jsonb;default:'{}'" json:"payload"`
	Status     string         `gorm:"size:20;default:'sent';index" json:"status"` // pending | sent | failed | read
	ErrorCode  string         `gorm:"size:50" json:"error_code"`
	ErrorMsg   string         `gorm:"size:500" json:"error_msg"`
	SentAt     *time.Time     `gorm:"index" json:"sent_at"`
	ReadAt     *time.Time     `gorm:"index" json:"read_at"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
}

func (UserNotification) TableName() string { return "user_notifications" }
