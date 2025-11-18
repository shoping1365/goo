package handlers

import (
	"encoding/json"
	"fmt"
	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type QuestionHandler struct {
	DB *gorm.DB
}

func NewQuestionHandler(db *gorm.DB) *QuestionHandler {
	return &QuestionHandler{DB: db}
}

func (h *QuestionHandler) ListQuestions(c *gin.Context) {
	var questions []models.Question
	if err := h.DB.Find(&questions).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت سوالات", err.Error()))
		return
	}
	c.JSON(http.StatusOK, questions)
}

func (h *QuestionHandler) GetQuestion(c *gin.Context) {
	id := c.Param("id")
	var question models.Question
	if err := h.DB.First(&question, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "سوال یافت نشد", nil))
		return
	}
	c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) CreateQuestion(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}
	if err := h.DB.Create(&question).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد سوال", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, question)
}

func (h *QuestionHandler) UpdateQuestion(c *gin.Context) {
	id := c.Param("id")
	var question models.Question
	if err := h.DB.First(&question, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "سوال یافت نشد", nil))
		return
	}
	if err := c.ShouldBindJSON(&question); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}
	if err := h.DB.Save(&question).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بروزرسانی سوال", err.Error()))
		return
	}
	c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Question{}, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف سوال", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Question deleted"})
}

// GetProductQuestions دریافت پرسش‌های یک محصول خاص
func (h *QuestionHandler) GetProductQuestions(c *gin.Context) {
	productID := c.Param("productId")
	var questions []models.Question
	if err := h.DB.Where("product_id = ? AND status = ?", productID, "approved").Find(&questions).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت سوالات محصول", err.Error()))
		return
	}
	c.JSON(http.StatusOK, questions)
}

// CreateProductQuestion ایجاد پرسش برای محصول خاص
func (h *QuestionHandler) CreateProductQuestion(c *gin.Context) {
	productID := c.Param("productId")

	// Debug: print raw body
	var raw map[string]interface{}
	if err := c.BindJSON(&raw); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ورودی نامعتبر (raw)", err.Error()))
		return
	}
	fmt.Printf("RAW BODY: %+v\n", raw)

	var question models.Question
	_ = c.ShouldBindJSON(&question)
	if q, ok := raw["question"]; ok {
		if qs, ok := q.(string); ok {
			question.QuestionText = qs
		}
	}
	// Guest info
	if name, ok := raw["name"]; ok {
		if ns, ok := name.(string); ok {
			question.GuestName = ns
		}
	}
	if phone, ok := raw["phone"]; ok {
		if ps, ok := phone.(string); ok {
			question.GuestPhone = ps
		}
	}
	// Set IP address & user agent
	question.IPAddress = c.ClientIP()
	question.UserAgent = c.Request.UserAgent()
	// اگر deviceInfo از کلاینت ارسال شد، ذخیره کن
	if di, ok := raw["device_info"]; ok {
		if m, ok := di.(map[string]interface{}); ok {
			b, _ := json.Marshal(m)
			question.DeviceInfo = datatypes.JSON(b)
		}
	}

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه محصول نامعتبر", err.Error()))
		return
	}
	question.ProductID = uint(productIDInt)
	question.Status = "pending"

	if err := h.DB.Create(&question).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد سوال", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, question)
}

// ListAllProductQuestions returns all product questions for admin panel
func (h *QuestionHandler) ListAllProductQuestions(c *gin.Context) {
	// First, let's see the raw data in the database
	var rawQuestions []map[string]interface{}
	h.DB.Raw("SELECT id, user_id, product_id, question_text, is_anonymous, status, created_at FROM product_questions ORDER BY created_at DESC").Scan(&rawQuestions)
	fmt.Printf("Raw database data: %+v\n", rawQuestions)

	var questions []models.Question
	query := h.DB.Model(&models.Question{}).Preload("User").Preload("Product")

	// Optional filters
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
	if search := c.Query("search"); search != "" {
		like := "%" + search + "%"
		query = query.Where("question_text ILIKE ?", like)
	}

	if err := query.Order("created_at DESC").Find(&questions).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "DB error", "details": err.Error()})
		return
	}

	// Map to include customer name, product name, and question text
	mapped := make([]map[string]interface{}, len(questions))
	for i, q := range questions {
		// Debug logging
		fmt.Printf("Question %d: UserID=%d, IsAnonymous=%v, User.Name='%s', User.ID=%d\n",
			q.ID, q.UserID, q.IsAnonymous, q.User.Name, q.User.ID)

		mapped[i] = map[string]interface{}{
			"id": q.ID,
			"customer_name": func() string {
				if q.IsAnonymous {
					return "کاربر مهمان"
				}
				if q.UserID > 0 && q.User.Name != "" {
					return q.User.Name
				}
				if q.GuestName != "" {
					return q.GuestName
				}
				return "کاربر مهمان"
			}(),
			"customer_mobile": func() string {
				if q.UserID > 0 {
					return q.User.Mobile
				}
				if q.GuestPhone != "" {
					return q.GuestPhone
				}
				return "نامشخص"
			}(),
			"product_id":    q.ProductID,
			"product_name":  q.Product.Name,
			"product_price": q.Product.Price,
			"question":      q.QuestionText,
			"category":      q.CategoryID,
			"status":        q.Status,
			"created_at":    q.CreatedAt,
			"answer":        q.AdminReply,
			"answered_at":   q.AdminReplyAt,
			"ip_address":    q.IPAddress,
			"user_agent":    q.UserAgent,
			"device_info":   q.DeviceInfo,
			"priority":      "medium", // If you have a priority field, use it
			"is_anonymous":  q.IsAnonymous,
		}
	}

	c.JSON(200, mapped)
}
