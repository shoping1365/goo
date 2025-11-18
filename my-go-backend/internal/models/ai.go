package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AIGenerateContentRequest struct {
	Model       string        `json:"model"`
	Messages    []interface{} `json:"messages"`
	WordCount   int           `json:"wordCount"`
	Temperature float64       `json:"temperature"`
}

type AIGenerateContentResponse struct {
	Content string `json:"content"`
}

// ChatMessage پیام چت
type ChatMessage struct {
	Role    string `json:"role"`    // "user" یا "assistant"
	Content string `json:"content"` // محتوای پیام
}

// AIChatRequest درخواست چت تعاملی
type AIChatRequest struct {
	Message   string `json:"message" binding:"required"`
	SessionID string `json:"session_id"`
}

// AIChatResponse پاسخ چت تعاملی
type AIChatResponse struct {
	Message              string                 `json:"message"`
	Type                 string                 `json:"type"` // "suggestion", "confirmation", "article_preview", "article_created", "article_edit"
	Suggestions          []string               `json:"suggestions,omitempty"`
	Article              *GeneratedArticle      `json:"article,omitempty"`
	PostID               *uint                  `json:"post_id,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	RequiresConfirmation bool                   `json:"requires_confirmation,omitempty"`
}

// GeneratedArticle مقاله تولید شده
type GeneratedArticle struct {
	Title           string `json:"title"`
	Content         string `json:"content"`
	Excerpt         string `json:"excerpt"`
	MetaDescription string `json:"meta_description"`
	Slug            string `json:"slug"`
	CategoryID      *uint  `json:"category_id,omitempty"`
	Tags            string `json:"tags,omitempty"`
}

// AISession جلسه چت AI (در حافظه)
type AISession struct {
	SessionID     string            `json:"session_id"`
	CurrentState  string            `json:"current_state"` // "initial", "title_selected", "article_preview", "editing"
	SelectedTitle string            `json:"selected_title,omitempty"`
	DraftArticle  *GeneratedArticle `json:"draft_article,omitempty"`
	EditHistory   []string          `json:"edit_history,omitempty"`
	ChatHistory   []ChatMessage     `json:"chat_history,omitempty"` // تاریخچه چت
	CreatedAt     int64             `json:"created_at"`
	LastActivity  int64             `json:"last_activity"`
}

// AISessionDB جلسه چت AI (در دیتابیس)
type AISessionDB struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	SessionID     string         `json:"session_id" gorm:"size:255;uniqueIndex;not null"`
	UserID        uint           `json:"user_id" gorm:"not null;index"`
	CurrentState  string         `json:"current_state" gorm:"size:50;default:'initial'"`
	SelectedTitle string         `json:"selected_title" gorm:"type:text"`
	DraftArticle  datatypes.JSON `json:"draft_article" gorm:"type:jsonb"`
	EditHistory   datatypes.JSON `json:"edit_history" gorm:"type:jsonb"`
	Metadata      datatypes.JSON `json:"metadata" gorm:"type:jsonb"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// روابط
	User     User            `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Messages []AIChatMessage `json:"messages,omitempty" gorm:"foreignKey:SessionID;references:SessionID"`
}

// AIChatMessage پیام چت AI (در دیتابیس)
type AIChatMessage struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	SessionID string         `json:"session_id" gorm:"size:255;index;not null"`
	Role      string         `json:"role" gorm:"size:20;not null"` // "user" یا "assistant"
	Content   string         `json:"content" gorm:"type:text;not null"`
	Type      string         `json:"type" gorm:"size:50"` // نوع پیام
	Metadata  datatypes.JSON `json:"metadata" gorm:"type:jsonb"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// TableName نام جدول‌ها
func (AISessionDB) TableName() string {
	return "ai_sessions"
}

func (AIChatMessage) TableName() string {
	return "ai_chat_messages"
}
