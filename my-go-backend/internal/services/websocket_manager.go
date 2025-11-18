package services

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"my-go-backend/internal/auth"
	"my-go-backend/internal/models"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

// WebSocketManager مدیریت اتصالات WebSocket
type WebSocketManager struct {
	// اتصالات فعال
	connections map[string]*WebSocketConnection
	// اتاق‌های چت
	rooms map[string]*ChatRoom
	// قفل برای thread safety
	mutex sync.RWMutex
	// دیتابیس
	db *gorm.DB
	// سرویس رمزنگاری
	encryptionSvc *EncryptionService
	// مبداهای مجاز WebSocket (Origin)
	allowedOrigins map[string]bool
	// Rate limit برای پیام‌ها بر اساس UserID
	// کلیددهی سطل‌ها بر اساس کاربر یا اتصال (برای مهمان‌ها)
	messageBuckets map[string]*messageBucket
	// تنظیمات نرخ پیام
	rateLimitConfig RateLimitConfig
}

// messageBucket برای کنترل نرخ پیام
// ساده‎ترین شکل token-bucket در حافظه
type RateLimitConfig struct {
	CustomerLimit int `json:"customer_limit"`
	OperatorLimit int `json:"operator_limit"`
	AdminLimit    int `json:"admin_limit"`
	WindowSeconds int `json:"window_seconds"`
}

const (
	defaultCustomerLimit = 20
	defaultOperatorLimit = 60
	defaultAdminLimit    = 120
	defaultWindowSeconds = 60

	// حداکثر طول پیام
	maxMessageLength = 1000
)

// الگوهای امنیتی برای اعتبارسنجی
var (
	// الگوهای مخرب
	maliciousPatterns = []*regexp.Regexp{
		regexp.MustCompile(`(?is)<script[^>]*>.*?</script>`),
		regexp.MustCompile(`(?i)javascript:`),
		regexp.MustCompile(`(?i)on\w+\s*=`),
		regexp.MustCompile(`(?i)eval\s*\(`),
		regexp.MustCompile(`(?i)<iframe\b[^>]*>`),
		regexp.MustCompile(`(?i)<object\b[^>]*>`),
		regexp.MustCompile(`(?i)<embed\b[^>]*>`),
		regexp.MustCompile(`(?i)\.\.\/|\.\.\\`),
		regexp.MustCompile(`(?i)(SELECT|INSERT|UPDATE|DELETE|DROP|CREATE|ALTER|EXEC|UNION)\b`),
		regexp.MustCompile(`--|\#|\/\*|\*\/`),
	}

	// کلمات ممنوع
	bannedWords = []string{
		"eval", "function", "alert", "prompt", "confirm",
		"document.cookie", "localstorage", "sessionstorage",
		"window.location", "location.href",
	}
)

type messageBucket struct {
	tokens int
	last   time.Time
}

// WebSocketConnection اتصال WebSocket
type WebSocketConnection struct {
	ID           string
	Conn         *websocket.Conn
	UserID       uint
	UserType     string // customer, operator, admin
	SessionID    string
	IP           string
	LastActivity time.Time
	Send         chan []byte
}

// ChatRoom اتاق چت
type ChatRoom struct {
	ID           string
	SessionID    string
	Customers    map[string]*WebSocketConnection
	Operators    map[string]*WebSocketConnection
	Messages     []ChatMessageData
	CreatedAt    time.Time
	LastActivity time.Time
}

// ChatMessageData داده‌های پیام چت
type ChatMessageData struct {
	ID          uint      `json:"id"`
	SessionID   string    `json:"session_id"`
	SenderType  string    `json:"sender_type"`
	SenderID    *uint     `json:"sender_id,omitempty"`
	SenderName  string    `json:"sender_name"`
	Message     string    `json:"message"`
	MessageType string    `json:"message_type"`
	FileURL     string    `json:"file_url,omitempty"`
	FileSize    int64     `json:"file_size,omitempty"`
	FileType    string    `json:"file_type,omitempty"`
	IsRead      bool      `json:"is_read"`
	CreatedAt   time.Time `json:"created_at"`
}

// WebSocketEvent رویداد WebSocket
type WebSocketEvent struct {
	Type    string      `json:"type"`
	Data    interface{} `json:"data"`
	Session string      `json:"session,omitempty"`
}

// NewWebSocketManager ایجاد نمونه جدید از WebSocketManager
func NewWebSocketManager(db *gorm.DB, encryptionSvc *EncryptionService) (*WebSocketManager, error) {
	// بارگذاری Originهای مجاز از متغیر محیطی (جدا شده با کاما)
	allowed := os.Getenv("CHAT_ALLOWED_ORIGINS")
	if allowed == "" {
		allowed = "*" // پیش‌فرض: همه‌ی مبداها (Development)
	}
	originMap := make(map[string]bool)
	for _, o := range strings.Split(allowed, ",") {
		originMap[strings.TrimSpace(o)] = true
	}

	// بارگذاری تنظیمات نرخ پیام از دیتابیس
	rateConfig := RateLimitConfig{
		CustomerLimit: defaultCustomerLimit,
		OperatorLimit: defaultOperatorLimit,
		AdminLimit:    defaultAdminLimit,
		WindowSeconds: defaultWindowSeconds,
	}

	// تلاش برای بارگذاری از دیتابیس
	var setting struct {
		Value string `gorm:"column:value"`
	}
	if err := db.Table("api_settings").Where("key = ?", "chat_rate_limit").First(&setting).Error; err == nil {
		if err := json.Unmarshal([]byte(setting.Value), &rateConfig); err == nil {
			// تنظیمات با موفقیت بارگذاری شد
		}
	}

	return &WebSocketManager{
		connections:     make(map[string]*WebSocketConnection),
		rooms:           make(map[string]*ChatRoom),
		db:              db,
		encryptionSvc:   encryptionSvc,
		allowedOrigins:  originMap,
		messageBuckets:  make(map[string]*messageBucket),
		rateLimitConfig: rateConfig,
	}, nil
}

// HandleWebSocket مدیریت اتصال WebSocket جدید
func (wm *WebSocketManager) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// ----------------- بررسی Origin -----------------
	origin := r.Header.Get("Origin")
	if origin != "" && len(wm.allowedOrigins) > 0 {
		if _, ok := wm.allowedOrigins["*"]; !ok {
			if _, ok := wm.allowedOrigins[origin]; !ok {
				http.Error(w, "origin not allowed", http.StatusForbidden)
				return
			}
		}
	}

	// ----------------- احراز هویت -----------------
	// ---- استخراج توکن ----
	token := r.URL.Query().Get("token")
	if token == "" {
		// بررسی Header: Authorization: Bearer <token>
		authHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			token = strings.TrimSpace(authHeader[7:])
		}
	}
	if token == "" {
		// بررسی کوکی Sidebase: next-auth.session-token
		if cookie, err := r.Cookie("next-auth.session-token"); err == nil {
			token = cookie.Value
		}
	}
	var claims struct {
		UserID uint
		Role   string
	}
	if token == "" {
		// کاربر مهمان بدون توکن
		claims.UserID = 0
		claims.Role = "customer"
	} else {
		parsed, err := auth.ValidateToken(token)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		claims.UserID = parsed.UserID
		claims.Role = parsed.Role
	}

	// تنظیمات WebSocket
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Origin قبلاً بررسی شده است
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// استخراج IP کاربر
	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = r.RemoteAddr
	}

	// ارتقا اتصال HTTP به WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("خطا در ارتقا اتصال WebSocket: %v", err)
		return
	}

	// تعیین نوع کاربر بر اساس نقش
	var userType string
	switch claims.Role {
	case "admin", "developer":
		userType = "admin"
	case "operator":
		userType = "operator"
	default:
		userType = "customer"
	}

	// ایجاد اتصال جدید
	connection := &WebSocketConnection{
		ID:           generateConnectionID(),
		Conn:         conn,
		UserID:       claims.UserID,
		UserType:     userType,
		IP:           clientIP,
		LastActivity: time.Now(),
		Send:         make(chan []byte, 256),
	}

	// اضافه کردن اتصال به manager
	wm.addConnection(connection)

	// شروع goroutine برای خواندن پیام‌ها
	go wm.readPump(connection)
	// شروع goroutine برای نوشتن پیام‌ها
	go wm.writePump(connection)
}

// addConnection اضافه کردن اتصال جدید
func (wm *WebSocketManager) addConnection(conn *WebSocketConnection) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()
	wm.connections[conn.ID] = conn
}

// removeConnection حذف اتصال
func (wm *WebSocketManager) removeConnection(conn *WebSocketConnection) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()

	// حذف از connections
	delete(wm.connections, conn.ID)

	// حذف از room
	if conn.SessionID != "" {
		if room, exists := wm.rooms[conn.SessionID]; exists {
			if conn.UserType == "customer" {
				delete(room.Customers, conn.ID)
			} else if conn.UserType == "operator" {
				delete(room.Operators, conn.ID)
			}

			// اگر اتاق خالی شد، آن را حذف کنید
			if len(room.Customers) == 0 && len(room.Operators) == 0 {
				delete(wm.rooms, conn.SessionID)
			}
		}
	}

	// بستن کانال
	close(conn.Send)
}

// readPump خواندن پیام‌های ورودی
func (wm *WebSocketManager) readPump(conn *WebSocketConnection) {
	defer func() {
		wm.removeConnection(conn)
		conn.Conn.Close()
	}()

	conn.Conn.SetReadLimit(512) // محدودیت اندازه پیام
	conn.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.Conn.SetPongHandler(func(string) error {
		conn.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := conn.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("خطا در خواندن پیام WebSocket: %v", err)
			}
			break
		}

		// پردازش پیام
		wm.handleMessage(conn, message)
		conn.LastActivity = time.Now()
	}
}

// writePump نوشتن پیام‌های خروجی
func (wm *WebSocketManager) writePump(conn *WebSocketConnection) {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		conn.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-conn.Send:
			conn.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				conn.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := conn.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			conn.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := conn.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// handleMessage پردازش پیام دریافتی
func (wm *WebSocketManager) handleMessage(conn *WebSocketConnection, message []byte) {
	var event WebSocketEvent
	if err := json.Unmarshal(message, &event); err != nil {
		log.Printf("خطا در parse کردن پیام: %v", err)
		return
	}

	switch event.Type {
	case "join_session":
		wm.handleJoinSession(conn, event)
	case "send_message":
		wm.handleSendMessage(conn, event)
	case "message": // سازگاری عقب‌رو با کلاینت‌های قدیمی
		wm.handleSendMessage(conn, event)
	case "typing":
		wm.handleTyping(conn, event)
	case "read_message":
		wm.handleReadMessage(conn, event)
	case "leave_session":
		wm.handleLeaveSession(conn, event)
	default:
		log.Printf("نوع پیام ناشناخته: %s", event.Type)
	}
}

// handleJoinSession مدیریت پیوستن به جلسه
func (wm *WebSocketManager) handleJoinSession(conn *WebSocketConnection, event WebSocketEvent) {
	// اینجا باید احراز هویت و مجوزها را بررسی کنید
	// برای سادگی، فعلاً فقط session ID را ذخیره می‌کنیم

	if sessionID, ok := event.Data.(map[string]interface{})["session_id"].(string); ok {
		conn.SessionID = sessionID

		// اضافه کردن به room
		wm.mutex.Lock()
		if room, exists := wm.rooms[sessionID]; exists {
			if conn.UserType == "customer" {
				room.Customers[conn.ID] = conn
			} else if conn.UserType == "operator" {
				room.Operators[conn.ID] = conn
			}
		} else {
			// ایجاد room جدید
			room = &ChatRoom{
				ID:           sessionID,
				SessionID:    sessionID,
				Customers:    make(map[string]*WebSocketConnection),
				Operators:    make(map[string]*WebSocketConnection),
				CreatedAt:    time.Now(),
				LastActivity: time.Now(),
			}

			if conn.UserType == "customer" {
				room.Customers[conn.ID] = conn
			} else if conn.UserType == "operator" {
				room.Operators[conn.ID] = conn
			}

			wm.rooms[sessionID] = room
		}
		wm.mutex.Unlock()

		// ارسال تأیید پیوستن
		response := WebSocketEvent{
			Type: "session_joined",
			Data: map[string]interface{}{
				"session_id": sessionID,
				"status":     "success",
			},
		}
		wm.sendToConnection(conn, response)
	}
}

// handleSendMessage مدیریت ارسال پیام
func (wm *WebSocketManager) handleSendMessage(conn *WebSocketConnection, event WebSocketEvent) {
	// Rate limit check
	if !wm.allowMessage(conn) {
		// ارسال پیام خطا به کلاینت و نادیده گرفتن پیام
		wm.sendToConnection(conn, WebSocketEvent{
			Type: "error",
			Data: gin.H{"message": "rate limit exceeded"},
		})
		return
	}

	if conn.SessionID == "" {
		// اگر جلسه وجود ندارد، یک جلسه جدید ایجاد کن (برای مشتریان مهمان)
		// توضیح: این کار باعث می‌شود مهمان بتواند بدون فراخوانی صریح join_session
		// وارد یک سشن موقت شود. برای اپراتورها توصیه می‌شود از رویداد join_session استفاده گردد.
		newSessionID := generateConnectionID()
		conn.SessionID = newSessionID
		// اطلاع به کلاینت از پیوستن به سشن جدید
		wm.sendToConnection(conn, WebSocketEvent{
			Type: "session_joined",
			Data: map[string]interface{}{
				"session_id": newSessionID,
				"status":     "success",
			},
		})
		// ساخت یا گرفتن room
		wm.mutex.Lock()
		if room, exists := wm.rooms[newSessionID]; exists {
			if conn.UserType == "customer" {
				room.Customers[conn.ID] = conn
			} else if conn.UserType == "operator" {
				room.Operators[conn.ID] = conn
			}
		} else {
			room = &ChatRoom{
				ID:           newSessionID,
				SessionID:    newSessionID,
				Customers:    make(map[string]*WebSocketConnection),
				Operators:    make(map[string]*WebSocketConnection),
				CreatedAt:    time.Now(),
				LastActivity: time.Now(),
			}
			if conn.UserType == "customer" {
				room.Customers[conn.ID] = conn
			} else if conn.UserType == "operator" {
				room.Operators[conn.ID] = conn
			}
			wm.rooms[newSessionID] = room
		}
		wm.mutex.Unlock()
	}

	// استخراج داده‌های پیام
	messageData, ok := event.Data.(map[string]interface{})
	if !ok {
		wm.sendToConnection(conn, WebSocketEvent{
			Type: "error",
			Data: gin.H{"message": "invalid message format"},
		})
		return
	}

	// استخراج محتوای پیام
	messageContent, ok := messageData["content"].(string)
	if !ok {
		messageContent, ok = messageData["message"].(string)
		if !ok {
			wm.sendToConnection(conn, WebSocketEvent{
				Type: "error",
				Data: gin.H{"message": "message content is required"},
			})
			return
		}
	}

	// اعتبارسنجی امنیتی پیام
	validationResult := wm.validateMessage(messageContent)

	// لاگ امنیتی
	wm.logSecurityEvent(conn.UserID, conn.UserType, messageContent, validationResult)

	if !validationResult.IsValid {
		// ارسال پیام خطا به کلاینت
		wm.sendToConnection(conn, WebSocketEvent{
			Type: "error",
			Data: gin.H{"message": validationResult.Error},
		})

		// در صورت تهدید شدید، اتصال را قطع کن
		if validationResult.RiskLevel == "critical" {
			log.Printf("🚨 Closing connection due to critical security threat from user %d", conn.UserID)
			conn.Conn.Close()
			wm.removeConnection(conn)
		}

		return
	}

	// استفاده از محتوای پاکسازی شده
	sanitizedContent := validationResult.SanitizedContent

	// ایجاد پیام جدید
	message := ChatMessageData{
		SessionID:   conn.SessionID,
		SenderType:  conn.UserType,
		SenderID:    &conn.UserID,
		SenderName:  getSenderName(conn),
		Message:     sanitizedContent, // استفاده از محتوای پاکسازی شده
		MessageType: getStringValue(messageData, "message_type", "text"),
		FileURL:     getStringValue(messageData, "file_url", ""),
		FileSize:    getInt64Value(messageData, "file_size", 0),
		FileType:    getStringValue(messageData, "file_type", ""),
		IsRead:      false,
		CreatedAt:   time.Now(),
	}

	// ذخیره در دیتابیس
	if err := wm.saveMessageToDatabase(message, conn); err != nil {
		log.Printf("خطا در ذخیره پیام: %v", err)
		wm.sendToConnection(conn, WebSocketEvent{
			Type: "error",
			Data: gin.H{"message": "failed to save message"},
		})
		return
	}

	// ارسال به همه اعضای room (هر دو فرمت برای سازگاری)
	wm.broadcastToRoom(conn.SessionID, WebSocketEvent{
		Type: "new_message",
		Data: message,
	})
	// رویداد ساده‌تر برای کلاینت‌های قدیمی
	legacy := map[string]interface{}{
		"type":      "message",
		"sender":    map[string]string{"customer": "user", "operator": "agent", "admin": "agent"}[message.SenderType],
		"content":   message.Message,
		"timestamp": message.CreatedAt,
		"status":    "delivered",
	}
	wm.broadcastToRoom(conn.SessionID, WebSocketEvent{
		Type: "message",
		Data: legacy,
	})
}

// handleTyping مدیریت تایپ کردن
func (wm *WebSocketManager) handleTyping(conn *WebSocketConnection, event WebSocketEvent) {
	if conn.SessionID == "" {
		return
	}

	// ارسال وضعیت تایپ به سایر اعضا
	wm.broadcastToRoom(conn.SessionID, WebSocketEvent{
		Type: "typing",
		Data: map[string]interface{}{
			"user_id":   conn.UserID,
			"user_type": conn.UserType,
			"is_typing": true,
		},
	})
}

// handleReadMessage مدیریت خواندن پیام
func (wm *WebSocketManager) handleReadMessage(_ *WebSocketConnection, _ WebSocketEvent) {
	// اینجا باید وضعیت خوانده شدن پیام را در دیتابیس به‌روزرسانی کنید
	// و به فرستنده اطلاع دهید
}

// handleLeaveSession مدیریت ترک جلسه
func (wm *WebSocketManager) handleLeaveSession(conn *WebSocketConnection, _ WebSocketEvent) {
	wm.removeConnection(conn)
}

// broadcastToRoom ارسال پیام به همه اعضای room
func (wm *WebSocketManager) broadcastToRoom(sessionID string, event WebSocketEvent) {
	wm.mutex.RLock()
	room, exists := wm.rooms[sessionID]
	wm.mutex.RUnlock()

	if !exists {
		return
	}

	// ارسال به همه customers
	for _, conn := range room.Customers {
		wm.sendToConnection(conn, event)
	}

	// ارسال به همه operators
	for _, conn := range room.Operators {
		wm.sendToConnection(conn, event)
	}
}

// sendToConnection ارسال پیام به یک اتصال
// allowMessage بررسی می‌کند آیا کاربر مجاز به ارسال پیام جدید هست یا نه
func (wm *WebSocketManager) allowMessage(conn *WebSocketConnection) bool {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()

	key := wm.bucketKeyFor(conn)
	bucket, ok := wm.messageBuckets[key]
	if !ok {
		bucket = &messageBucket{tokens: wm.maxTokensFor(conn), last: time.Now()}
		wm.messageBuckets[key] = bucket
	}

	// refill
	elapsed := time.Since(bucket.last)
	if elapsed >= time.Duration(wm.rateLimitConfig.WindowSeconds)*time.Second {
		bucket.tokens = wm.maxTokensFor(conn)
		bucket.last = time.Now()
	}

	if bucket.tokens <= 0 {
		return false
	}
	bucket.tokens--
	return true
}

// maxTokensFor تعیین سقف توکن بسته به نوع کاربر
func (wm *WebSocketManager) maxTokensFor(conn *WebSocketConnection) int {
	switch conn.UserType {
	case "admin":
		return wm.rateLimitConfig.AdminLimit
	case "operator":
		return wm.rateLimitConfig.OperatorLimit
	default:
		return wm.rateLimitConfig.CustomerLimit
	}
}

// bucketKeyFor تولید کلید مناسب برای سطل نرخ پیام
// ورودی: اتصال وب‌سوکت شامل شناسه کاربر و اتصال
// خروجی: کلید یکتا. برای کاربران لاگین «u:<userId>» و برای مهمان‌ها «c:<connectionId>»
func (wm *WebSocketManager) bucketKeyFor(conn *WebSocketConnection) string {
	if conn.UserID > 0 {
		return "u:" + strings.TrimSpace(fmt.Sprintf("%d", conn.UserID))
	}
	// برای مهمان‌ها از شناسه اتصال استفاده می‌کنیم تا هر تب/کاربر جدا محدود شود
	return "c:" + conn.ID
}

// UpdateRateLimitConfig بروزرسانی تنظیمات نرخ پیام
func (wm *WebSocketManager) UpdateRateLimitConfig(config RateLimitConfig) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()
	wm.rateLimitConfig = config
}

// GetRateLimitConfig دریافت تنظیمات فعلی نرخ پیام
func (wm *WebSocketManager) GetRateLimitConfig() RateLimitConfig {
	wm.mutex.RLock()
	defer wm.mutex.RUnlock()
	return wm.rateLimitConfig
}

// GetStats دریافت آمار WebSocket Manager
func (wm *WebSocketManager) GetStats() map[string]interface{} {
	wm.mutex.RLock()
	defer wm.mutex.RUnlock()

	stats := make(map[string]interface{})
	stats["total_connections"] = len(wm.connections)
	stats["total_rooms"] = len(wm.rooms)
	stats["rate_limit_config"] = wm.rateLimitConfig

	// محاسبه آمار اضافی
	var activeConnections int
	var totalMessages int

	for _, conn := range wm.connections {
		if conn.Conn != nil {
			activeConnections++
		}
	}

	stats["active_connections"] = activeConnections
	stats["total_messages"] = totalMessages

	return stats
}

// GetConnections دریافت لیست اتصالات فعال
func (wm *WebSocketManager) GetConnections() map[string]*WebSocketConnection {
	wm.mutex.RLock()
	defer wm.mutex.RUnlock()

	// کپی از map برای جلوگیری از race condition
	connections := make(map[string]*WebSocketConnection)
	for k, v := range wm.connections {
		connections[k] = v
	}

	return connections
}

// اعتبارسنجی محتوای پیام
type ValidationResult struct {
	IsValid          bool   `json:"is_valid"`
	SanitizedContent string `json:"sanitized_content"`
	Error            string `json:"error"`
	RiskLevel        string `json:"risk_level"` // low, medium, high, critical
}

// validateMessage اعتبارسنجی محتوای پیام
func (wm *WebSocketManager) validateMessage(content string) ValidationResult {
	// بررسی طول پیام
	if len(content) == 0 {
		return ValidationResult{
			IsValid:   false,
			Error:     "پیام نمی‌تواند خالی باشد",
			RiskLevel: "low",
		}
	}

	if len(content) > maxMessageLength {
		return ValidationResult{
			IsValid:   false,
			Error:     "پیام بیش از حد طولانی است",
			RiskLevel: "medium",
		}
	}

	// بررسی الگوهای مخرب
	for _, pattern := range maliciousPatterns {
		if pattern.MatchString(content) {
			log.Printf("🚨 Malicious pattern detected: %s", pattern.String())
			return ValidationResult{
				IsValid:   false,
				Error:     "پیام حاوی محتوای مخرب است",
				RiskLevel: "critical",
			}
		}
	}

	// بررسی کلمات ممنوع
	lowerContent := strings.ToLower(content)
	for _, word := range bannedWords {
		if strings.Contains(lowerContent, word) {
			log.Printf("⚠️ Banned word detected: %s", word)
			return ValidationResult{
				IsValid:   false,
				Error:     "پیام حاوی کلمات ممنوع است",
				RiskLevel: "high",
			}
		}
	}

	// پاکسازی محتوا
	sanitized := sanitizeContent(content)

	return ValidationResult{
		IsValid:          true,
		SanitizedContent: sanitized,
		RiskLevel:        "low",
	}
}

// sanitizeContent پاکسازی محتوای پیام
func sanitizeContent(content string) string {
	// HTML escape
	sanitized := html.EscapeString(content)

	// حذف کاراکترهای کنترلی
	sanitized = regexp.MustCompile(`[\x00-\x1F\x7F]`).ReplaceAllString(sanitized, "")

	// محدود کردن فضاهای اضافی
	sanitized = regexp.MustCompile(`\s+`).ReplaceAllString(sanitized, " ")
	sanitized = strings.TrimSpace(sanitized)

	// محدود کردن کاراکترهای تکراری بدون استفاده از backreference (سازگار با RE2)
	sanitized = limitRepeatedCharacters(sanitized, 4)

	return sanitized
}

// limitRepeatedCharacters تعداد کاراکترهای تکراری متوالی را به حداکثر max محدود می‌کند
func limitRepeatedCharacters(s string, max int) string {
	if max < 1 || s == "" {
		return s
	}
	var b strings.Builder
	b.Grow(len(s))
	var prev rune
	count := 0
	for i, r := range s {
		if i == 0 {
			prev = r
			count = 1
			b.WriteRune(r)
			continue
		}
		if r == prev {
			count++
			if count <= max {
				b.WriteRune(r)
			}
		} else {
			prev = r
			count = 1
			b.WriteRune(r)
		}
	}
	return b.String()
}

// logSecurityEvent ثبت رویداد امنیتی
func (wm *WebSocketManager) logSecurityEvent(userID uint, userType string, content string, result ValidationResult) {
	logData := map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"user_id":   userID,
		"user_type": userType,
		"content":   content,
		"result":    result,
		"ip":        "unknown", // TODO: دریافت IP کاربر
	}

	logJSON, err := json.Marshal(logData)
	if err != nil {
		log.Printf("Failed to marshal log data: %v", err)
		return
	}

	if result.RiskLevel == "critical" || result.RiskLevel == "high" {
		log.Printf("🚨 SECURITY ALERT: %s", string(logJSON))
	} else if result.RiskLevel == "medium" {
		log.Printf("⚠️ Security Warning: %s", string(logJSON))
	}
}

func (wm *WebSocketManager) sendToConnection(conn *WebSocketConnection, event WebSocketEvent) {
	data, err := json.Marshal(event)
	if err != nil {
		log.Printf("خطا در marshal کردن پیام: %v", err)
		return
	}

	select {
	case conn.Send <- data:
	default:
		// اگر کانال پر است، اتصال را ببندید
		close(conn.Send)
		wm.removeConnection(conn)
	}
}

// saveMessageToDatabase ذخیره پیام در دیتابیس
// saveMessageToDatabase ذخیره امن پیام در پایگاه داده
// ورودی: ساختار پیام (با session_id رشته‌ای) و اتصال فعلی (برای درج IP در صورت نیاز)
// رفتار: اگر سشن با session_id داده‌شده موجود نبود، یک رکورد حداقلی ساخته می‌شود.
func (wm *WebSocketManager) saveMessageToDatabase(message ChatMessageData, conn *WebSocketConnection) error {
	// پیدا کردن شناسه عددی سشن از روی کد سشن رشته‌ای (ایجاد در صورت نبود)
	var session models.ChatSession
	if err := wm.db.Where("session_id = ?", message.SessionID).First(&session).Error; err != nil {
		// ایجاد جلسه حداقلی اگر وجود نداشت
		session = models.ChatSession{
			SessionID: message.SessionID,
			Status:    "waiting",
			IPAddress: conn.IP,
			StartedAt: time.Now(),
		}
		if err := wm.db.Create(&session).Error; err != nil {
			return err
		}
	}

	// رمزنگاری پیام قبل از ذخیره
	encryptedMessage, err := wm.encryptionSvc.EncryptMessage(message.Message)
	if err != nil {
		return err
	}

	// ایجاد مدل پیام برای ذخیره در دیتابیس
	chatMessage := &models.OnlineChatMessage{
		SessionID:   session.ID,
		SenderType:  message.SenderType,
		SenderID:    message.SenderID,
		Message:     encryptedMessage,
		MessageType: message.MessageType,
		FileURL:     message.FileURL,
		FileSize:    message.FileSize,
		FileType:    message.FileType,
		IsRead:      message.IsRead,
		CreatedAt:   message.CreatedAt,
	}

	// ذخیره در دیتابیس
	return wm.db.Create(chatMessage).Error
}

// HandleWebSocketGin Gin adapter برای WebSocket
func (wm *WebSocketManager) HandleWebSocketGin(c *gin.Context) {
	wm.HandleWebSocket(c.Writer, c.Request)
}

// Helper functions
func generateConnectionID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

func getSenderName(conn *WebSocketConnection) string {
	// اینجا باید نام فرستنده را از دیتابیس دریافت کنید
	return "کاربر"
}

func getStringValue(data map[string]interface{}, key, defaultValue string) string {
	if value, ok := data[key].(string); ok {
		return value
	}
	return defaultValue
}

func getInt64Value(data map[string]interface{}, key string, defaultValue int64) int64 {
	if value, ok := data[key].(float64); ok {
		return int64(value)
	}
	return defaultValue
}
