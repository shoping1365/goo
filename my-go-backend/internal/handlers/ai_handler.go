package handlers

import (
	"context"
	"net/http"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// AIHandler هندلر تولید محتوای AI
// تمام توضیحات به فارسی طبق پروتکل پروژه
// این هندلر فقط برای ادمین و با پرمیشن مناسب فعال است

type AIHandler struct {
	AIService *services.AIService
}

func NewAIHandler(aiService *services.AIService) *AIHandler {
	return &AIHandler{AIService: aiService}
}

// GenerateContent تولید محتوای متنی با OpenAI
// @Summary تولید محتوا با OpenAI
// @Tags AI
// @Accept json
// @Produce json
// @Param body body models.AIGenerateContentRequest true "درخواست تولید محتوا"
// @Success 200 {object} models.AIGenerateContentResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/ai/generate-content [post]
func (h *AIHandler) GenerateContent(c *gin.Context) {
	var req models.AIGenerateContentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse := utils.HandleValidationError("درخواست", "داده‌های ارسالی نامعتبر است")
		c.AbortWithStatusJSON(errorResponse.StatusCode, errorResponse)
		return
	}

	// دریافت userId از context (در صورت نیاز)
	userId, _ := c.Get("userId")

	ctx := context.Background()
	resp, err := h.AIService.GenerateContent(ctx, &req, userId)
	if err != nil {
		// خطا از سرویس با پیام کاربرپسند دریافت شده
		errorResponse := utils.HandleBusinessError("تولید محتوا", err.Error())
		c.AbortWithStatusJSON(errorResponse.StatusCode, errorResponse)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// HandleChat چت تعاملی برای تولید محتوا
// @Summary چت تعاملی AI
// @Tags AI
// @Accept json
// @Produce json
// @Param body body models.AIChatRequest true "درخواست چت"
// @Success 200 {object} models.AIChatResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/ai/chat [post]
func (h *AIHandler) HandleChat(c *gin.Context) {
	var req models.AIChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse := utils.HandleValidationError("درخواست چت", "داده‌های ارسالی نامعتبر است")
		c.AbortWithStatusJSON(errorResponse.StatusCode, errorResponse)
		return
	}

	// دریافت userId از context
	userId, _ := c.Get("user_id")

	ctx := context.Background()
	resp, err := h.AIService.HandleChat(ctx, &req, userId)
	if err != nil {
		errorResponse := utils.HandleBusinessError("چت AI", err.Error())
		c.AbortWithStatusJSON(errorResponse.StatusCode, errorResponse)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetChatHistory دریافت تاریخچه چت برای یک جلسه
func (h *AIHandler) GetChatHistory(c *gin.Context) {
	sessionID := c.Param("session_id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id required"})
		return
	}

	// دریافت پیام‌ها از دیتابیس
	var messages []models.AIChatMessage
	if err := h.AIService.DB.Where("session_id = ?", sessionID).Order("created_at ASC").Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت تاریخچه چت"})
		return
	}

	// تبدیل به فرمت مناسب
	var chatHistory []gin.H
	for _, msg := range messages {
		chatHistory = append(chatHistory, gin.H{
			"role":       msg.Role,
			"content":    msg.Content,
			"type":       msg.Type,
			"created_at": msg.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id": sessionID,
		"messages":   chatHistory,
		"count":      len(chatHistory),
	})
}

// GetUserSessions دریافت همه جلسات کاربر
func (h *AIHandler) GetUserSessions(c *gin.Context) {
	userId, _ := c.Get("user_id")

	var sessions []models.AISessionDB
	if err := h.AIService.DB.Where("user_id = ?", userId).Order("updated_at DESC").Find(&sessions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت جلسات"})
		return
	}

	var sessionList []gin.H
	for _, session := range sessions {
		// شمارش پیام‌ها
		var messageCount int64
		h.AIService.DB.Model(&models.AIChatMessage{}).Where("session_id = ?", session.SessionID).Count(&messageCount)

		sessionList = append(sessionList, gin.H{
			"session_id":     session.SessionID,
			"current_state":  session.CurrentState,
			"selected_title": session.SelectedTitle,
			"message_count":  messageCount,
			"created_at":     session.CreatedAt,
			"updated_at":     session.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"sessions": sessionList,
		"count":    len(sessionList),
	})
}
