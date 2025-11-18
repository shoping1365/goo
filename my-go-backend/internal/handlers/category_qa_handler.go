package handlers

import (
	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryQAHandler struct {
	DB *gorm.DB
}

func NewCategoryQAHandler(db *gorm.DB) *CategoryQAHandler {
	return &CategoryQAHandler{DB: db}
}

// GetCategories دریافت همه دسته‌بندی‌های Q&A
func (h *CategoryQAHandler) GetCategories(c *gin.Context) {
	var categories []models.CategoryQA
	if err := h.DB.Order("created_at DESC").Find(&categories).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	var result []map[string]interface{}
	for _, cat := range categories {
		var questionCount int64
		h.DB.Model(&models.Question{}).Where("category = ?", cat.Key).Count(&questionCount)
		result = append(result, map[string]interface{}{
			"id":             cat.ID,
			"key":            cat.Key,
			"name":           cat.Name,
			"question_count": questionCount,
			"created_at":     cat.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, result)
}

// CreateCategory ایجاد دسته‌بندی جدید
func (h *CategoryQAHandler) CreateCategory(c *gin.Context) {
	var cat models.CategoryQA
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	if cat.Key == "" || cat.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	if err := h.DB.Create(&cat).Error; err != nil {
		if h.DB.Where("key = ?", cat.Key).First(&models.CategoryQA{}).RowsAffected > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_KEY", utils.GetErrorMessage("DUPLICATE_KEY"), nil))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusCreated, cat)
}

// TransferQuestions انتقال همه سوالات از یک دسته‌بندی به دسته‌بندی دیگر
func (h *CategoryQAHandler) TransferQuestions(c *gin.Context) {
	var req struct {
		FromCategory string `json:"from_category" binding:"required"`
		ToCategory   string `json:"to_category" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	var toCat models.CategoryQA
	if err := h.DB.Where("key = ?", req.ToCategory).First(&toCat).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}
	if req.FromCategory == req.ToCategory {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	var questionCount int64
	h.DB.Model(&models.Question{}).Where("category = ?", req.FromCategory).Count(&questionCount)
	if questionCount == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}
	result := h.DB.Model(&models.Question{}).Where("category = ?", req.FromCategory).Update("category", req.ToCategory)
	c.JSON(http.StatusOK, gin.H{
		"message":           "انتقال سوالات با موفقیت انجام شد",
		"transferred_count": result.RowsAffected,
		"from_category":     req.FromCategory,
		"to_category":       req.ToCategory,
	})
}

// DeleteCategory حذف دسته‌بندی
func (h *CategoryQAHandler) DeleteCategory(c *gin.Context) {
	key := c.Param("key")
	var count int64
	h.DB.Model(&models.Question{}).Where("category = ?", key).Count(&count)
	if count > 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("PAGE_NOT_FOUND"), map[string]interface{}{"question_count": count}))
		return
	}
	result := h.DB.Where("key = ?", key).Delete(&models.CategoryQA{})
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "دسته‌بندی با موفقیت حذف شد"})
}
