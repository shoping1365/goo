package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
)

// BrandHandler handles CRUD operations for product brands.
type BrandHandler struct {
	DB *gorm.DB
}

// NewBrandHandler returns a new BrandHandler instance.
func NewBrandHandler(db *gorm.DB) *BrandHandler {
	return &BrandHandler{DB: db}
}

// ListBrands returns all brands ordered by name.
func (h *BrandHandler) ListBrands(c *gin.Context) {
	var brands []models.Brand
	if err := h.DB.Order("name ASC").Find(&brands).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت برندها", err.Error()))
		return
	}

	// Optionally include product_count for each brand (front-end may ignore it)
	var result []gin.H
	for _, b := range brands {
		var productCount int64
		h.DB.Model(&models.Product{}).Where("brand_id = ?", b.ID).Count(&productCount)

		result = append(result, gin.H{
			"id":               b.ID,
			"name":             b.Name,
			"slug":             b.Slug,
			"description":      b.Description,
			"official_name":    b.OfficialName,
			"published":        b.Published,
			"show_on_home":     b.ShowOnHome,
			"show_in_menu":     b.ShowInMenu,
			"meta_title":       b.MetaTitle,
			"meta_keywords":    b.MetaKeywords,
			"meta_description": b.MetaDescription,
			"image_url":        b.ImageURL,
			"video_url":        b.VideoURL,
			"product_count":    productCount,
			"created_at":       b.CreatedAt,
			"updated_at":       b.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, result)
}

// GetBrand returns a single brand by ID.
func (h *BrandHandler) GetBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand
	if err := h.DB.First(&brand, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("BRAND_NOT_FOUND", "برند یافت نشد", nil))
		return
	}

	c.JSON(http.StatusOK, brand)
}

// CreateBrand creates a new brand.
func (h *BrandHandler) CreateBrand(c *gin.Context) {
	type createBrandInput struct {
		Name            string `json:"name" binding:"required"`
		Description     string `json:"description"`
		OfficialName    string `json:"official_name"`
		Published       bool   `json:"published"`
		ShowOnHome      bool   `json:"show_on_home"`
		ShowInMenu      bool   `json:"show_in_menu"`
		MetaTitle       string `json:"meta_title"`
		MetaKeywords    string `json:"meta_keywords"`
		MetaDescription string `json:"meta_description"`
		ImageURL        string `json:"image_url"`
		VideoURL        string `json:"video_url"`
	}

	var input createBrandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "اطلاعات ورودی نامعتبر است", err.Error()))
		return
	}

	newSlug := slug.Make(input.Name)

	// Check slug uniqueness
	var existing models.Brand
	if err := h.DB.Where("slug = ?", newSlug).First(&existing).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_BRAND", "برندی با این نام از قبل وجود دارد", nil))
		return
	}

	brand := models.Brand{
		Name:            input.Name,
		Slug:            newSlug,
		Description:     input.Description,
		OfficialName:    input.OfficialName,
		Published:       input.Published,
		ShowOnHome:      input.ShowOnHome,
		ShowInMenu:      input.ShowInMenu,
		MetaTitle:       input.MetaTitle,
		MetaKeywords:    input.MetaKeywords,
		MetaDescription: input.MetaDescription,
		ImageURL:        input.ImageURL,
		VideoURL:        input.VideoURL,
	}

	if err := h.DB.Create(&brand).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد برند", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, brand)
}

// UpdateBrand updates an existing brand.
func (h *BrandHandler) UpdateBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand
	if err := h.DB.First(&brand, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("BRAND_NOT_FOUND", "برند یافت نشد", nil))
		return
	}

	type updateBrandInput struct {
		Name            string `json:"name"`
		Description     string `json:"description"`
		OfficialName    string `json:"official_name"`
		Published       *bool  `json:"published"`
		ShowOnHome      *bool  `json:"show_on_home"`
		ShowInMenu      *bool  `json:"show_in_menu"`
		MetaTitle       string `json:"meta_title"`
		MetaKeywords    string `json:"meta_keywords"`
		MetaDescription string `json:"meta_description"`
		ImageURL        string `json:"image_url"`
		VideoURL        string `json:"video_url"`
	}

	var input updateBrandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "اطلاعات ورودی نامعتبر است", err.Error()))
		return
	}

	if input.Name != "" {
		newSlug := slug.Make(input.Name)
		brand.Name = input.Name
		brand.Slug = newSlug
	}

	if input.Description != "" {
		brand.Description = input.Description
	}

	if input.OfficialName != "" {
		brand.OfficialName = input.OfficialName
	}

	if input.Published != nil {
		brand.Published = *input.Published
	}

	if input.ShowOnHome != nil {
		brand.ShowOnHome = *input.ShowOnHome
	}

	if input.ShowInMenu != nil {
		brand.ShowInMenu = *input.ShowInMenu
	}

	if input.MetaTitle != "" {
		brand.MetaTitle = input.MetaTitle
	}

	if input.MetaKeywords != "" {
		brand.MetaKeywords = input.MetaKeywords
	}

	if input.MetaDescription != "" {
		brand.MetaDescription = input.MetaDescription
	}

	if input.ImageURL != "" {
		brand.ImageURL = input.ImageURL
	}

	if input.VideoURL != "" {
		brand.VideoURL = input.VideoURL
	}

	if err := h.DB.Save(&brand).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی برند", err.Error()))
		return
	}

	c.JSON(http.StatusOK, brand)
}

// DeleteBrand deletes a brand (soft delete).
func (h *BrandHandler) DeleteBrand(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Brand{}, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف برند", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "برند با موفقیت حذف شد"})
}
