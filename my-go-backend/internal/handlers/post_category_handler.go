package handlers

import (
	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type PostCategoryHandler struct {
	DB *gorm.DB
}

func NewPostCategoryHandler(db *gorm.DB) *PostCategoryHandler {
	return &PostCategoryHandler{DB: db}
}

// ListCategories دریافت همه دسته‌بندی‌های نوشته‌ها
func (h *PostCategoryHandler) ListCategories(c *gin.Context) {
	var categories []models.PostCategory

	// اگر پارامتر all=1 ارسال شده باشد، همه دسته‌بندی‌ها (فعال و غیرفعال) را برگردان
	showAll := c.Query("all") == "1" || c.Query("all") == "true"

	query := h.DB.Order("\"order\" ASC, created_at DESC")
	if !showAll {
		query = query.Where("is_active = ?", true)
	}

	if err := query.Find(&categories).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	// ساخت پاسخ با تعداد نوشته‌ها برای هر دسته‌بندی
	type categoryWithCount struct {
		models.PostCategory
		PostsCount int64 `json:"posts_count"`
	}

	resp := make([]categoryWithCount, 0, len(categories))
	for _, cat := range categories {
		var count int64
		h.DB.Model(&models.Post{}).Where("category_id = ?", cat.ID).Count(&count)
		resp = append(resp, categoryWithCount{PostCategory: cat, PostsCount: count})
	}

	c.JSON(http.StatusOK, resp)
}

// GetCategory دریافت یک دسته‌بندی خاص
func (h *PostCategoryHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.PostCategory

	if err := h.DB.First(&category, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("CATEGORY_NOT_FOUND", utils.GetErrorMessage("CATEGORY_NOT_FOUND"), nil))
		return
	}

	var postsCount int64
	h.DB.Model(&models.Post{}).Where("category_id = ?", category.ID).Count(&postsCount)

	c.JSON(http.StatusOK, gin.H{
		"id":          category.ID,
		"name":        category.Name,
		"slug":        category.Slug,
		"description": category.Description,
		"is_active":   category.IsActive,
		"order":       category.Order,
		"posts_count": postsCount,
		"created_at":  category.CreatedAt,
		"updated_at":  category.UpdatedAt,
	})
}

// CreateCategory ایجاد دسته‌بندی جدید
func (h *PostCategoryHandler) CreateCategory(c *gin.Context) {
	type createCategoryInput struct {
		Name        string `json:"name" binding:"required"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
		IsActive    bool   `json:"is_active"`
		Order       uint   `json:"order"`
	}

	var input createCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// استفاده از slug ارائه شده یا تولید از نام
	newSlug := input.Slug
	if newSlug == "" {
		newSlug = slug.Make(input.Name)
	}

	// بررسی تکراری نبودن slug
	var existingCategory models.PostCategory
	if err := h.DB.Where("slug = ?", newSlug).First(&existingCategory).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_CATEGORY", utils.GetErrorMessage("DUPLICATE_CATEGORY"), nil))
		return
	}

	// تنظیم مقادیر پیش‌فرض
	if input.Order == 0 {
		input.Order = 1
	}

	category := models.PostCategory{
		Name:        input.Name,
		Slug:        newSlug,
		Description: input.Description,
		IsActive:    input.IsActive,
		Order:       input.Order,
	}

	if err := h.DB.Create(&category).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusCreated, category)
}

// UpdateCategory ویرایش دسته‌بندی موجود
func (h *PostCategoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.PostCategory

	if err := h.DB.First(&category, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("CATEGORY_NOT_FOUND", utils.GetErrorMessage("CATEGORY_NOT_FOUND"), nil))
		return
	}

	type updateCategoryInput struct {
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
		IsActive    bool   `json:"is_active"`
		Order       uint   `json:"order"`
	}

	var input updateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// بروزرسانی فیلدها در صورت ارائه
	if input.Name != "" {
		var newSlug string
		if input.Slug != "" {
			newSlug = input.Slug
		} else {
			newSlug = slug.Make(input.Name)
		}
		// بررسی تکراری نبودن slug
		var existingCategory models.PostCategory
		if err := h.DB.Where("slug = ? AND id != ?", newSlug, category.ID).First(&existingCategory).Error; err == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_CATEGORY", utils.GetErrorMessage("DUPLICATE_CATEGORY"), nil))
			return
		}
		category.Name = input.Name
		category.Slug = newSlug
	}

	if input.Description != "" {
		category.Description = input.Description
	}

	category.IsActive = input.IsActive
	category.Order = input.Order

	if err := h.DB.Save(&category).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory حذف دسته‌بندی
func (h *PostCategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.PostCategory

	if err := h.DB.First(&category, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("CATEGORY_NOT_FOUND", utils.GetErrorMessage("CATEGORY_NOT_FOUND"), nil))
		return
	}

	// بررسی وجود نوشته در این دسته‌بندی
	var postsCount int64
	h.DB.Model(&models.Post{}).Where("category_id = ?", category.ID).Count(&postsCount)
	if postsCount > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("CATEGORY_HAS_POSTS", "این دسته‌بندی دارای نوشته است و قابل حذف نیست", nil))
		return
	}

	if err := h.DB.Delete(&category).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "دسته‌بندی با موفقیت حذف شد"})
}

// ToggleCategoryStatus تغییر وضعیت فعال/غیرفعال دسته‌بندی
func (h *PostCategoryHandler) ToggleCategoryStatus(c *gin.Context) {
	id := c.Param("id")
	var category models.PostCategory

	if err := h.DB.First(&category, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("CATEGORY_NOT_FOUND", utils.GetErrorMessage("CATEGORY_NOT_FOUND"), nil))
		return
	}

	category.IsActive = !category.IsActive

	if err := h.DB.Save(&category).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "وضعیت دسته‌بندی با موفقیت تغییر کرد",
		"is_active": category.IsActive,
	})
}
