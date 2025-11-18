package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type ProductCategoryHandler struct {
	DB *gorm.DB
}

func NewProductCategoryHandler(db *gorm.DB) *ProductCategoryHandler {
	return &ProductCategoryHandler{DB: db}
}

// ListCategories returns all published product categories
func (h *ProductCategoryHandler) ListCategories(c *gin.Context) {
	var categories []models.Category

	// If query param all=1, include both published and unpublished
	showAll := c.Query("all") == "1" || c.Query("all") == "true"

	query := h.DB.Order("\"order\" ASC, created_at DESC")
	if !showAll {
		query = query.Where("published = ?", true)
	}

	if err := query.Find(&categories).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	// Build response with product count for each category
	type categoryWithCount struct {
		models.Category
		ProductCount int64 `json:"product_count"`
	}

	resp := make([]categoryWithCount, 0, len(categories))
	for _, cat := range categories {
		var count int64
		h.DB.Model(&models.Product{}).Where("category_id = ?", cat.ID).Count(&count)
		resp = append(resp, categoryWithCount{Category: cat, ProductCount: count})
	}

	c.JSON(http.StatusOK, resp)
}

// GetCategory returns a single product category by ID
func (h *ProductCategoryHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := h.DB.First(&category, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("CATEGORY_NOT_FOUND", utils.GetErrorMessage("CATEGORY_NOT_FOUND"), nil))
		return
	}

	var productCount int64
	h.DB.Model(&models.Product{}).Where("category_id = ?", category.ID).Count(&productCount)

	c.JSON(http.StatusOK, gin.H{
		"id":               category.ID,
		"name":             category.Name,
		"slug":             category.Slug,
		"parent_id":        category.ParentID,
		"name_en":          category.NameEn,
		"description":      category.Description,
		"image_url":        category.ImageURL,
		"icon":             category.Icon,
		"order":            category.Order,
		"published":        category.Published,
		"show_on_home":     category.ShowOnHome,
		"show_in_menu":     category.ShowInMenu,
		"banner_url":       category.BannerURL,
		"notice_message":   category.NoticeMessage,
		"meta_title":       category.MetaTitle,
		"meta_keywords":    category.MetaKeywords,
		"meta_description": category.MetaDescription,
		"product_count":    productCount,
		"created_at":       category.CreatedAt,
		"updated_at":       category.UpdatedAt,
	})
}

// GetCategoryBySlug returns a single product category by slug
// اگر پارامتر preview=1 باشد، دسته‌های منتشرنشده نیز قابل مشاهده هستند
func (h *ProductCategoryHandler) GetCategoryBySlug(c *gin.Context) {
	slugParam := c.Param("slug")
	isPreview := c.Query("preview") == "1" || c.Query("preview") == "true"

	var category models.Category
	q := h.DB.Where("slug = ?", slugParam)
	if !isPreview {
		q = q.Where("published = ?", true)
	}
	if err := q.First(&category).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("CATEGORY_NOT_FOUND", utils.GetErrorMessage("CATEGORY_NOT_FOUND"), nil))
		return
	}

	var productCount int64
	h.DB.Model(&models.Product{}).Where("category_id = ?", category.ID).Count(&productCount)

	c.JSON(http.StatusOK, gin.H{
		"id":               category.ID,
		"name":             category.Name,
		"slug":             category.Slug,
		"parent_id":        category.ParentID,
		"name_en":          category.NameEn,
		"description":      category.Description,
		"image_url":        category.ImageURL,
		"icon":             category.Icon,
		"order":            category.Order,
		"published":        category.Published,
		"show_on_home":     category.ShowOnHome,
		"show_in_menu":     category.ShowInMenu,
		"banner_url":       category.BannerURL,
		"notice_message":   category.NoticeMessage,
		"meta_title":       category.MetaTitle,
		"meta_keywords":    category.MetaKeywords,
		"meta_description": category.MetaDescription,
		"product_count":    productCount,
		"created_at":       category.CreatedAt,
		"updated_at":       category.UpdatedAt,
	})
}

// CreateCategory creates a new product category
func (h *ProductCategoryHandler) CreateCategory(c *gin.Context) {
	type createCategoryInput struct {
		Name            string `json:"name" binding:"required"`
		NameEn          string `json:"name_en"`
		Slug            string `json:"slug"`
		ParentID        *uint  `json:"parent_id"`
		Description     string `json:"description"`
		ImageURL        string `json:"image_url"`
		Icon            string `json:"icon"`
		Order           uint   `json:"order"`
		Published       bool   `json:"published"`
		ShowOnHome      bool   `json:"show_on_home"`
		ShowInMenu      bool   `json:"show_in_menu"`
		BannerURL       string `json:"banner_url"`
		NoticeMessage   string `json:"notice_message"`
		MetaTitle       string `json:"meta_title"`
		MetaKeywords    string `json:"meta_keywords"`
		MetaDescription string `json:"meta_description"`
	}

	var input createCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// Rule: slug الزامی است؛ اگر خالی بود و name_en موجود بود از آن ساخته می‌شود
	newSlug := strings.TrimSpace(input.Slug)
	if newSlug == "" {
		if strings.TrimSpace(input.NameEn) != "" {
			newSlug = slug.Make(strings.TrimSpace(input.NameEn))
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("CATEGORY_SLUG_REQUIRED", "URL یا نام رسمی (English) الزامی است.", nil))
			return
		}
	}

	// اگر کاربر اصرار بر URL فعلی داشت، قبل از یکتا‌سازی پیام هشدار می‌دهیم (سمت فرانت)
	// Ensure slug uniqueness: if exists, append incremental suffix (-2, -3, ...)
	uniqueSlug := newSlug
	for i := 2; i <= 99; i++ {
		var count int64
		if err := h.DB.Model(&models.Category{}).Where("slug = ?", uniqueSlug).Count(&count).Error; err != nil {
			break
		}
		if count == 0 {
			break
		}
		uniqueSlug = fmt.Sprintf("%s-%d", newSlug, i)
	}
	newSlug = uniqueSlug

	// Set default values if not provided
	if input.Order == 0 {
		input.Order = 1
	}

	category := models.Category{
		Name:            input.Name,
		NameEn:          input.NameEn,
		Slug:            newSlug,
		ParentID:        input.ParentID,
		Description:     input.Description,
		ImageURL:        input.ImageURL,
		Icon:            input.Icon,
		Order:           input.Order,
		Published:       input.Published,
		ShowOnHome:      input.ShowOnHome,
		ShowInMenu:      input.ShowInMenu,
		BannerURL:       input.BannerURL,
		NoticeMessage:   input.NoticeMessage,
		MetaTitle:       input.MetaTitle,
		MetaKeywords:    input.MetaKeywords,
		MetaDescription: input.MetaDescription,
	}

	if err := h.DB.Create(&category).Error; err != nil {
		// بررسی خطای duplicate key
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "UNIQUE constraint") {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_CATEGORY", "نام یا URL دسته‌بندی تکراری است", err.Error()))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusCreated, category)
}

// UpdateCategory updates an existing product category
func (h *ProductCategoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := h.DB.First(&category, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("CATEGORY_NOT_FOUND", utils.GetErrorMessage("CATEGORY_NOT_FOUND"), nil))
		return
	}

	type updateCategoryInput struct {
		Name            string `json:"name"`
		NameEn          string `json:"name_en"`
		Slug            string `json:"slug"`
		ParentID        *uint  `json:"parent_id"`
		Description     string `json:"description"`
		ImageURL        string `json:"image_url"`
		Icon            string `json:"icon"`
		Order           uint   `json:"order"`
		Published       bool   `json:"published"`
		ShowOnHome      bool   `json:"show_on_home"`
		ShowInMenu      bool   `json:"show_in_menu"`
		BannerURL       string `json:"banner_url"`
		NoticeMessage   string `json:"notice_message"`
		MetaTitle       string `json:"meta_title"`
		MetaKeywords    string `json:"meta_keywords"`
		MetaDescription string `json:"meta_description"`
	}

	var input updateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// Update fields if provided
	if input.Name != "" {
		var newSlug string
		if input.Slug != "" {
			newSlug = input.Slug
		} else {
			newSlug = slug.Make(input.Name)
		}
		// Check if another category with this slug exists
		var existingCategory models.Category
		if err := h.DB.Where("slug = ? AND id != ?", newSlug, category.ID).First(&existingCategory).Error; err == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_CATEGORY", utils.GetErrorMessage("DUPLICATE_CATEGORY"), nil))
			return
		}
		category.Name = input.Name
		category.Slug = newSlug
	}

	if input.NameEn != "" {
		category.NameEn = input.NameEn
	}

	// Update parent category even when clearing (0 means no parent)
	if input.ParentID != nil {
		if *input.ParentID == 0 {
			category.ParentID = nil
		} else {
			category.ParentID = input.ParentID
		}
	}

	if input.Description != "" {
		category.Description = input.Description
	}

	if input.ImageURL != "" {
		category.ImageURL = input.ImageURL
	}

	if input.Icon != "" {
		category.Icon = input.Icon
	}

	if input.Order != 0 {
		category.Order = input.Order
	}

	category.Published = input.Published
	category.ShowOnHome = input.ShowOnHome
	category.ShowInMenu = input.ShowInMenu

	if input.BannerURL != "" {
		category.BannerURL = input.BannerURL
	}

	if input.NoticeMessage != "" {
		category.NoticeMessage = input.NoticeMessage
	}

	if input.MetaTitle != "" {
		category.MetaTitle = input.MetaTitle
	}

	if input.MetaKeywords != "" {
		category.MetaKeywords = input.MetaKeywords
	}

	if input.MetaDescription != "" {
		category.MetaDescription = input.MetaDescription
	}

	if err := h.DB.Save(&category).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory deletes a product category
func (h *ProductCategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	categoryID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}

	// واکشی دسته‌بندی جهت بررسی اسلاگ
	var targetCategory models.Category
	if err := h.DB.First(&targetCategory, categoryID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("CATEGORY_NOT_FOUND", utils.GetErrorMessage("CATEGORY_NOT_FOUND"), nil))
		return
	}

	// جلوگیری از حذف دسته‌بندی پیش‌فرض
	if targetCategory.Slug == "default" {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("PAGE_NOT_FOUND"), nil))
		return
	}

	// اگر دسته‌بندی دارای محصول باشد، قبل از حذف محصولات را به دسته‌بندی پیش‌فرض منتقل می‌کنیم
	var productCount int64
	h.DB.Model(&models.Product{}).Where("category_id = ?", categoryID).Count(&productCount)

	// پیدا کردن یا ایجاد دسته‌بندی پیش‌فرض (slug = default)
	var defaultCat models.Category
	if err := h.DB.Where("slug = ?", "default").First(&defaultCat).Error; err != nil {
		// اگر پیدا نشد، بساز
		defaultCat = models.Category{
			Name:      "دسته‌بندی پیش‌فرض",
			Slug:      "default",
			Published: false,
		}
		h.DB.Create(&defaultCat)
	}

	if productCount > 0 {
		// انتقال همه محصولات به دسته‌بندی پیش‌فرض
		h.DB.Model(&models.Product{}).Where("category_id = ?", categoryID).Update("category_id", defaultCat.ID)
	}

	// در حال حاضر محدودیت داشتن زیرمجموعه را حفظ می‌کنیم؛ اگر بخواهید می‌توان این قسمت را هم برداشت
	var childCount int64
	h.DB.Model(&models.Category{}).Where("parent_id = ?", categoryID).Count(&childCount)
	if childCount > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("CATEGORY_HAS_CHILDREN", utils.GetErrorMessage("CATEGORY_HAS_CHILDREN"), nil))
		return
	}

	if err := h.DB.Delete(&models.Category{}, categoryID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "دسته‌بندی با موفقیت حذف شد"})
}

// ReorderCategories updates the order of multiple categories
func (h *ProductCategoryHandler) ReorderCategories(c *gin.Context) {
	type reorderItem struct {
		ID    uint `json:"id"`
		Order uint `json:"order"`
	}
	var items []reorderItem
	if err := c.ShouldBindJSON(&items); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	tx := h.DB.Begin()
	for _, item := range items {
		if err := tx.Model(&models.Category{}).Where("id = ?", item.ID).Update("order", item.Order).Error; err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
			return
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true})
}
