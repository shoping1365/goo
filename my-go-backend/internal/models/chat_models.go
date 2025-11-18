package models

import (
	"time"

	"gorm.io/gorm"
)

// ChatOperator مدل اپراتور چت
type ChatOperator struct {
	ID                 uint           `json:"id" gorm:"primaryKey"`
	UserID             uint           `json:"user_id"`
	Status             string         `json:"status" gorm:"default:'offline'"`
	MaxConcurrentChats int            `json:"max_concurrent_chats" gorm:"default:5"`
	CurrentChats       int            `json:"current_chats" gorm:"default:0"`
	WorkStartTime      string         `json:"work_start_time" gorm:"default:'09:00:00'"`
	WorkEndTime        string         `json:"work_end_time" gorm:"default:'18:00:00'"`
	Timezone           string         `json:"timezone" gorm:"default:'Asia/Tehran'"`
	AutoAccept         bool           `json:"auto_accept" gorm:"default:false"`
	IsActive           bool           `json:"is_active" gorm:"default:true"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relations
	User User `json:"user" gorm:"foreignKey:UserID"`
}

// ChatSession مدل جلسه چت
type ChatSession struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	SessionID     string         `json:"session_id" gorm:"uniqueIndex"`
	CustomerName  string         `json:"customer_name"`
	CustomerPhone string         `json:"customer_phone"`
	CustomerEmail string         `json:"customer_email,omitempty"`
	Status        string         `json:"status" gorm:"default:'waiting'"`
	OperatorID    *uint          `json:"operator_id,omitempty"`
	AIEnabled     bool           `json:"ai_enabled" gorm:"default:false"`
	WidgetID      *uint          `json:"widget_id,omitempty"`
	IPAddress     string         `json:"ip_address"`
	UserAgent     string         `json:"user_agent"`
	StartedAt     time.Time      `json:"started_at"`
	EndedAt       *time.Time     `json:"ended_at,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relations
	Operator *ChatOperator       `json:"operator,omitempty" gorm:"foreignKey:OperatorID"`
	Widget   *ChatWidget         `json:"widget,omitempty" gorm:"foreignKey:WidgetID"`
	Messages []OnlineChatMessage `json:"messages,omitempty" gorm:"foreignKey:SessionID;references:ID"`
}

// OnlineChatMessage مدل پیام چت آنلاین
type OnlineChatMessage struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	SessionID   uint           `json:"session_id"`
	SenderType  string         `json:"sender_type"` // customer, operator, ai
	SenderID    *uint          `json:"sender_id,omitempty"`
	Message     string         `json:"message"`
	MessageType string         `json:"message_type" gorm:"default:'text'"` // text, image, video, file
	FileURL     string         `json:"file_url,omitempty"`
	FileSize    int64          `json:"file_size,omitempty"`
	FileType    string         `json:"file_type,omitempty"`
	IsRead      bool           `json:"is_read" gorm:"default:false"`
	ReadAt      *time.Time     `json:"read_at,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relations
	Session  ChatSession   `json:"session" gorm:"foreignKey:SessionID;references:ID"`
	Operator *ChatOperator `json:"operator,omitempty" gorm:"foreignKey:SenderID"`
}

func (OnlineChatMessage) TableName() string { return "chat_messages" }

// ChatWidget مدل ابزارک چت
type ChatWidget struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name"`
	Description    string         `json:"description,omitempty"`
	IsActive       bool           `json:"is_active" gorm:"default:true"`
	WelcomeMessage string         `json:"welcome_message"`
	OfflineMessage string         `json:"offline_message"`
	Theme          string         `json:"theme" gorm:"default:'light'"`
	Position       string         `json:"position" gorm:"default:'bottom-right'"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// ChatAIBot مدل بات هوش مصنوعی
type ChatAIBot struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Description  string         `json:"description,omitempty"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	OpenAIKey    string         `json:"openai_key"`
	Model        string         `json:"model" gorm:"default:'gpt-3.5-turbo'"`
	MaxTokens    int            `json:"max_tokens" gorm:"default:1000"`
	Temperature  float64        `json:"temperature" gorm:"default:0.7"`
	SystemPrompt string         `json:"system_prompt"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// ChatKnowledgeBase مدل پایگاه دانش
type ChatKnowledgeBase struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Category  string         `json:"category"`
	Tags      string         `json:"tags,omitempty"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// ChatStatistics مدل آمار چت
type ChatStatistics struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	Date              time.Time `json:"date"`
	TotalSessions     int       `json:"total_sessions"`
	CompletedSessions int       `json:"completed_sessions"`
	TotalMessages     int       `json:"total_messages"`
	AvgResponseTime   float64   `json:"avg_response_time"`
	OperatorID        *uint     `json:"operator_id,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	// Relations
	Operator *ChatOperator `json:"operator,omitempty" gorm:"foreignKey:OperatorID"`
}
