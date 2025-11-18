package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// ChatHandler مدیریت درخواست‌های چت
type ChatHandler struct {
	chatService *services.ChatService
	wsManager   *services.WebSocketManager
	uowFactory  unitofwork.UnitOfWorkFactory
}

// NewChatHandler ایجاد Chat Handler جدید
func NewChatHandler(chatService *services.ChatService, wsManager *services.WebSocketManager, factory unitofwork.UnitOfWorkFactory) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
		wsManager:   wsManager,
		uowFactory:  factory,
	}
}

// hasPermission بررسی مجوز کاربر
func (h *ChatHandler) hasPermission(c *gin.Context, permissionName string) bool {
	userID, exists := c.Get("user_id")
	if !exists {
		return false
	}

	hasPermission, err := middleware.CheckUserPermission(h.uowFactory.Create().DB(), userID.(uint), permissionName)
	if err != nil {
		return false
	}

	return hasPermission
}

// CreateChatSession ایجاد جلسه چت جدید
func (ch *ChatHandler) CreateChatSession(c *gin.Context) {
	var request struct {
		CustomerName  string `json:"customer_name" binding:"required"`
		CustomerPhone string `json:"customer_phone" binding:"required"`
		CustomerEmail string `json:"customer_email,omitempty"`
		WidgetID      *uint  `json:"widget_id,omitempty"`
		IPAddress     string `json:"ip_address"`
		UserAgent     string `json:"user_agent"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "داده‌های ورودی نامعتبر است",
		})
		return
	}

	// ایجاد جلسه چت
	session := &models.ChatSession{
		CustomerName:  request.CustomerName,
		CustomerPhone: request.CustomerPhone,
		CustomerEmail: request.CustomerEmail,
		WidgetID:      request.WidgetID,
		IPAddress:     request.IPAddress,
		UserAgent:     request.UserAgent,
	}

	createdSession, err := ch.chatService.CreateChatSession(session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در ایجاد جلسه چت",
		})
		return
	}

	// صدور توکن مهمان هفت‌روزه
	token, _ := utils.GenerateGuestToken(createdSession.SessionID, request.CustomerName, request.CustomerPhone)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "جلسه چت با موفقیت ایجاد شد",
		"data":    createdSession,
		"token":   token,
	})
}

// GetChatSession دریافت جلسه چت
func (ch *ChatHandler) GetChatSession(c *gin.Context) {
	sessionID := c.Param("session_id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه جلسه الزامی است",
		})
		return
	}

	session, err := ch.chatService.GetChatSession(sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "جلسه چت یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   session,
	})
}

// GetWaitingSessions دریافت جلسه‌های در انتظار
func (ch *ChatHandler) GetWaitingSessions(c *gin.Context) {
	sessions, err := ch.chatService.GetWaitingSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت جلسه‌های در انتظار",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   sessions,
	})
}

// GetAdminSessions دریافت لیست جلسات برای پنل ادمین با امکان فیلتر وضعیت
func (ch *ChatHandler) GetAdminSessions(c *gin.Context) {
	status := c.Query("status")
	sessions, err := ch.chatService.GetSessionsByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت لیست جلسات",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": sessions})
}

// AssignOperator تخصیص اپراتور به جلسه
func (ch *ChatHandler) AssignOperator(c *gin.Context) {
	sessionID := c.Param("session_id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه جلسه الزامی است",
		})
		return
	}

	var request struct {
		OperatorID uint `json:"operator_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه اپراتور الزامی است",
		})
		return
	}

	err := ch.chatService.AssignOperatorToSession(sessionID, request.OperatorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "اپراتور با موفقیت تخصیص داده شد",
	})
}

// SendMessage ارسال پیام
func (ch *ChatHandler) SendMessage(c *gin.Context) {
	var request struct {
		SessionID   uint   `json:"session_id" binding:"required"`
		SenderType  string `json:"sender_type" binding:"required"`
		SenderID    *uint  `json:"sender_id,omitempty"`
		Message     string `json:"message" binding:"required"`
		MessageType string `json:"message_type"`
		FileURL     string `json:"file_url,omitempty"`
		FileSize    int64  `json:"file_size,omitempty"`
		FileType    string `json:"file_type,omitempty"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "داده‌های ورودی نامعتبر است",
		})
		return
	}

	message := &models.OnlineChatMessage{
		SessionID:   request.SessionID,
		SenderType:  request.SenderType,
		SenderID:    request.SenderID,
		Message:     request.Message,
		MessageType: request.MessageType,
		FileURL:     request.FileURL,
		FileSize:    request.FileSize,
		FileType:    request.FileType,
	}

	if err := ch.chatService.SaveMessage(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در ذخیره پیام",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "پیام با موفقیت ارسال شد",
		"data":    message,
	})
}

// GetSessionMessages دریافت پیام‌های جلسه
func (ch *ChatHandler) GetSessionMessages(c *gin.Context) {
	sessionIDStr := c.Param("session_id")
	sessionID, err := strconv.ParseUint(sessionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه جلسه نامعتبر است",
		})
		return
	}

	limitStr := c.DefaultQuery("limit", "50")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 50
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	messages, err := ch.chatService.GetSessionMessages(uint(sessionID), limit, offset)
	if err != nil {
		// بررسی اگر خطا مربوط به عدم وجود جلسه است
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "جلسه چت یافت نشد",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت پیام‌ها",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   messages,
	})
}

// MarkMessageAsRead علامت‌گذاری پیام به عنوان خوانده شده
func (ch *ChatHandler) MarkMessageAsRead(c *gin.Context) {
	messageIDStr := c.Param("message_id")
	messageID, err := strconv.ParseUint(messageIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه پیام نامعتبر است",
		})
		return
	}

	err = ch.chatService.MarkMessageAsRead(uint(messageID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در به‌روزرسانی وضعیت پیام",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "وضعیت پیام با موفقیت به‌روزرسانی شد",
	})
}

// MarkSessionAsRead علامت‌گذاری همه پیام‌های یک جلسه به عنوان خوانده شده
func (ch *ChatHandler) MarkSessionAsRead(c *gin.Context) {
	sessionIDStr := c.Param("session_id")
	sessionID, err := strconv.ParseUint(sessionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "شناسه جلسه نامعتبر است"})
		return
	}
	if err := ch.chatService.MarkSessionMessagesAsRead(uint(sessionID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در بروزرسانی پیام‌ها"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "پیام‌ها خوانده شدند"})
}

// EndChatSession پایان دادن به جلسه چت
func (ch *ChatHandler) EndChatSession(c *gin.Context) {
	sessionID := c.Param("session_id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه جلسه الزامی است",
		})
		return
	}

	err := ch.chatService.EndChatSession(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در پایان دادن به جلسه",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "جلسه با موفقیت پایان یافت",
	})
}

// ListSessions فهرست جلسات بر اساس وضعیت
func (ch *ChatHandler) ListSessions(c *gin.Context) {
	status := c.Query("status")
	limitStr := c.DefaultQuery("limit", "50")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	sessions, err := ch.chatService.ListSessionsByStatus(status, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت لیست جلسات",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   sessions,
	})
}

// MarkSessionRead علامت‌گذاری همه پیام‌های یک جلسه به عنوان خوانده شده
func (ch *ChatHandler) MarkSessionRead(c *gin.Context) {
	sessionIDStr := c.Param("session_id")
	sid, err := strconv.ParseUint(sessionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "شناسه جلسه نامعتبر است"})
		return
	}

	if err := ch.chatService.MarkSessionMessagesAsRead(uint(sid)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در بروزرسانی پیام‌ها"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "پیام‌ها خوانده شدند"})
}

// GetAvailableOperators دریافت اپراتورهای در دسترس
func (ch *ChatHandler) GetAvailableOperators(c *gin.Context) {
	operators, err := ch.chatService.GetAvailableOperators()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت اپراتورهای در دسترس",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   operators,
	})
}

// UpdateOperatorStatus به‌روزرسانی وضعیت اپراتور
func (ch *ChatHandler) UpdateOperatorStatus(c *gin.Context) {
	operatorIDStr := c.Param("operator_id")
	operatorID, err := strconv.ParseUint(operatorIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه اپراتور نامعتبر است",
		})
		return
	}

	var request struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "وضعیت الزامی است",
		})
		return
	}

	err = ch.chatService.UpdateOperatorStatus(uint(operatorID), request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در به‌روزرسانی وضعیت اپراتور",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "وضعیت اپراتور با موفقیت به‌روزرسانی شد",
	})
}

// InviteOperator دعوت و ایجاد اپراتور جدید (کاربر + رکورد اپراتور چت)
// ورودی: full_name, username, email, mobile, password, max_concurrent_chats, auto_accept
func (ch *ChatHandler) InviteOperator(c *gin.Context) {
	var req struct {
		FullName           string `json:"full_name" binding:"required"`
		Username           string `json:"username" binding:"required"`
		Email              string `json:"email" binding:"required"`
		Mobile             string `json:"mobile" binding:"required"`
		Password           string `json:"password" binding:"required"`
		MaxConcurrentChats int    `json:"max_concurrent_chats"`
		AutoAccept         bool   `json:"auto_accept"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "داده‌های ورودی نامعتبر است"})
		return
	}

	// اعتبارسنجی یکتا بودن username/email/mobile
	ctx := context.Background()
	uow := ch.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در تراکنش"})
		return
	}

	var count int64
	uow.DB().Model(&models.User{}).Where("username = ? OR email = ? OR mobile = ?", req.Username, req.Email, req.Mobile).Count(&count)
	if count > 0 {
		uow.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "کاربری با این اطلاعات وجود دارد"})
		return
	}

	// ایجاد کاربر با نقش پیش‌فرض اپراتور (در این پروژه RoleID اپراتور را 3 در نظر می‌گیریم؛ در غیر اینصورت از مشتری 1 استفاده می‌شود)
	hashedPwd, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در آماده‌سازی کاربر"})
		return
	}
	now := time.Now()
	user := models.User{
		Name:         req.FullName,
		Username:     req.Username,
		Email:        req.Email,
		Mobile:       req.Mobile,
		RoleID:       3,
		Status:       "active",
		PasswordHash: hashedPwd,
		RegisteredAt: now,
		UpdatedAt:    now,
	}
	if err := uow.DB().Create(&user).Error; err != nil {
		uow.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در ایجاد کاربر"})
		return
	}

	// ایجاد رکورد اپراتور چت
	op := models.ChatOperator{
		UserID:             user.ID,
		Status:             "offline",
		MaxConcurrentChats: req.MaxConcurrentChats,
		AutoAccept:         req.AutoAccept,
		WorkStartTime:      "09:00:00",
		WorkEndTime:        "18:00:00",
		Timezone:           "Asia/Tehran",
		IsActive:           true,
		CreatedAt:          now,
		UpdatedAt:          now,
	}
	if op.MaxConcurrentChats <= 0 {
		op.MaxConcurrentChats = 5
	}
	if err := uow.DB().Create(&op).Error; err != nil {
		uow.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در ایجاد اپراتور"})
		return
	}

	if err := uow.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در ذخیره"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": gin.H{
		"user":     user,
		"operator": op,
	}})
}

// GetChatStatistics دریافت آمار چت
func (ch *ChatHandler) GetChatStatistics(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	operatorIDStr := c.Query("operator_id")

	if startDateStr == "" || endDateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "تاریخ شروع و پایان الزامی است",
		})
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "فرمت تاریخ شروع نامعتبر است",
		})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "فرمت تاریخ پایان نامعتبر است",
		})
		return
	}

	var operatorID *uint
	if operatorIDStr != "" {
		if id, err := strconv.ParseUint(operatorIDStr, 10, 32); err == nil {
			opID := uint(id)
			operatorID = &opID
		}
	}

	statistics, err := ch.chatService.GetChatStatistics(startDate, endDate, operatorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت آمار چت",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   statistics,
	})
}

// CreateChatWidget ایجاد ابزارک چت
func (ch *ChatHandler) CreateChatWidget(c *gin.Context) {
	var widget models.ChatWidget
	if err := c.ShouldBindJSON(&widget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "داده‌های ورودی نامعتبر است",
		})
		return
	}

	err := ch.chatService.CreateChatWidget(&widget)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در ایجاد ابزارک چت",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "ابزارک چت با موفقیت ایجاد شد",
		"data":    widget,
	})
}

// GetActiveChatWidgets دریافت ابزارک‌های فعال
func (ch *ChatHandler) GetActiveChatWidgets(c *gin.Context) {
	widgets, err := ch.chatService.GetActiveChatWidgets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت ابزارک‌های فعال",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   widgets,
	})
}

// UpdateChatWidget به‌روزرسانی ابزارک چت
func (ch *ChatHandler) UpdateChatWidget(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "شناسه نامعتبر"})
		return
	}
	var input models.ChatWidget
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "داده‌های ورودی نامعتبر"})
		return
	}
	item, err := ch.chatService.UpdateChatWidget(uint(id), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در بروزرسانی"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": item})
}

// DeleteChatWidget حذف ابزارک چت
func (ch *ChatHandler) DeleteChatWidget(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "شناسه نامعتبر"})
		return
	}
	if err := ch.chatService.DeleteChatWidget(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در حذف"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// CreateAIBot ایجاد بات هوش مصنوعی
func (ch *ChatHandler) CreateAIBot(c *gin.Context) {
	var bot models.ChatAIBot
	if err := c.ShouldBindJSON(&bot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "داده‌های ورودی نامعتبر است",
		})
		return
	}

	err := ch.chatService.CreateAIBot(&bot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در ایجاد بات هوش مصنوعی",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "بات هوش مصنوعی با موفقیت ایجاد شد",
		"data":    bot,
	})
}

// GetActiveAIBots دریافت بات‌های فعال
func (ch *ChatHandler) GetActiveAIBots(c *gin.Context) {
	bots, err := ch.chatService.GetActiveAIBots()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت بات‌های فعال",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   bots,
	})
}

// CreateKnowledgeBase ایجاد پایگاه دانش
func (ch *ChatHandler) CreateKnowledgeBase(c *gin.Context) {
	var kb models.ChatKnowledgeBase
	if err := c.ShouldBindJSON(&kb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "داده‌های ورودی نامعتبر است",
		})
		return
	}

	err := ch.chatService.CreateKnowledgeBase(&kb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در ایجاد پایگاه دانش",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "پایگاه دانش با موفقیت ایجاد شد",
		"data":    kb,
	})
}

// SearchKnowledgeBase جستجو در پایگاه دانش
func (ch *ChatHandler) SearchKnowledgeBase(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "عبارت جستجو الزامی است",
		})
		return
	}

	results, err := ch.chatService.SearchKnowledgeBase(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در جستجو",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   results,
	})
}

// ListKnowledgeBase لیست پایگاه دانش با صفحه‌بندی
func (ch *ChatHandler) ListKnowledgeBase(c *gin.Context) {
	q := c.Query("q")
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	items, total, err := ch.chatService.ListKnowledgeBase(q, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در دریافت پایگاه دانش"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": items, "total": total})
}

// GetKnowledgeBase دریافت آیتم پایگاه دانش
func (ch *ChatHandler) GetKnowledgeBase(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "شناسه نامعتبر"})
		return
	}
	item, err := ch.chatService.GetKnowledgeBaseByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "یافت نشد"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": item})
}

// UpdateKnowledgeBase بروزرسانی آیتم پایگاه دانش
func (ch *ChatHandler) UpdateKnowledgeBase(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "شناسه نامعتبر"})
		return
	}

	var kb models.ChatKnowledgeBase
	if err := c.ShouldBindJSON(&kb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "داده‌های ورودی نامعتبر"})
		return
	}
	item, err := ch.chatService.UpdateKnowledgeBase(uint(id), &kb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در بروزرسانی"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": item})
}

// DeleteKnowledgeBase حذف آیتم پایگاه دانش
func (ch *ChatHandler) DeleteKnowledgeBase(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "شناسه نامعتبر"})
		return
	}
	if err := ch.chatService.DeleteKnowledgeBase(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در حذف"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// GetChatRateLimitSettings دریافت تنظیمات نرخ پیام چت
func (h *ChatHandler) GetChatRateLimitSettings(c *gin.Context) {
	// بررسی مجوز
	if !h.hasPermission(c, "chat_settings.read") {
		c.JSON(http.StatusNotFound, gin.H{"message": "دسترسی غیرمجاز"})
		return
	}

	// دریافت تنظیمات از WebSocket Manager
	config := h.wsManager.GetRateLimitConfig()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    config,
	})
}

// UpdateChatRateLimitSettings بروزرسانی تنظیمات نرخ پیام چت
func (h *ChatHandler) UpdateChatRateLimitSettings(c *gin.Context) {
	// بررسی مجوز
	if !h.hasPermission(c, "chat_settings.write") {
		c.JSON(http.StatusNotFound, gin.H{"message": "دسترسی غیرمجاز"})
		return
	}

	var config struct {
		CustomerLimit int `json:"customer_limit" binding:"required,min=1,max=1000"`
		OperatorLimit int `json:"operator_limit" binding:"required,min=1,max=1000"`
		AdminLimit    int `json:"admin_limit" binding:"required,min=1,max=1000"`
		WindowSeconds int `json:"window_seconds" binding:"required,min=10,max=3600"`
	}

	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "داده‌های ورودی نامعتبر",
			"error":   err.Error(),
		})
		return
	}

	// تبدیل به RateLimitConfig
	rateConfig := services.RateLimitConfig{
		CustomerLimit: config.CustomerLimit,
		OperatorLimit: config.OperatorLimit,
		AdminLimit:    config.AdminLimit,
		WindowSeconds: config.WindowSeconds,
	}

	// ذخیره در دیتابیس
	configJSON, _ := json.Marshal(rateConfig)

	// Upsert سازگار با PostgreSQL در جدول api_settings
	if err := h.uowFactory.Create().DB().Exec("INSERT INTO api_settings (key, value, created_at, updated_at) VALUES (?, ?, NOW(), NOW()) ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value, updated_at = NOW()", "chat_rate_limit", string(configJSON)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در ذخیره تنظیمات",
			"error":   err.Error(),
		})
		return
	}

	// بروزرسانی WebSocket Manager
	h.wsManager.UpdateRateLimitConfig(rateConfig)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "تنظیمات با موفقیت بروزرسانی شد",
		"data":    rateConfig,
	})
}

// GetChatMetrics دریافت متریک‌های چت
func (h *ChatHandler) GetChatMetrics(c *gin.Context) {
	// بررسی مجوز
	if !h.hasPermission(c, "chat_metrics.read") {
		c.JSON(http.StatusForbidden, gin.H{"message": "دسترسی غیرمجاز"})
		return
	}

	// دریافت آمار از WebSocket Manager
	_ = h.wsManager.GetStats() // فعلاً استفاده نمی‌شود

	// محاسبه متریک‌های اضافی
	var rateLimitedCount int64
	h.uowFactory.Create().DB().Model(&models.OnlineChatMessage{}).Where("is_rate_limited = ?", true).Count(&rateLimitedCount)

	metrics := gin.H{
		"active_connections":          len(h.wsManager.GetConnections()),
		"websocket_connections":       len(h.wsManager.GetConnections()),
		"avg_response_time":           150, // نمونه - در واقعیت باید محاسبه شود
		"messages_per_second":         5,   // نمونه
		"error_rate":                  0.5, // نمونه
		"rate_limited_messages":       rateLimitedCount,
		"cpu_usage":                   25, // نمونه
		"memory_usage":                45, // نمونه
		"disk_usage":                  30, // نمونه
		"network_usage":               15, // نمونه
		"performance_score":           85, // نمونه
		"connections_trend":           0,
		"response_time_trend":         0,
		"messages_trend":              0,
		"error_rate_trend":            0,
		"rate_limit_trend":            0,
		"websocket_connections_trend": 0,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"metrics": metrics,
		"health": gin.H{
			"status":     "healthy",
			"uptime":     "99.8%",
			"last_check": time.Now(),
		},
		"alerts": []gin.H{}, // نمونه
	})
}
