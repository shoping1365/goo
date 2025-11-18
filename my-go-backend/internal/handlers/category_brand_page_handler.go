package handlers

import (
	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type CategoryBrandPageHandler struct {
	DB *gorm.DB
}

func NewCategoryBrandPageHandler(db *gorm.DB) *CategoryBrandPageHandler {
	return &CategoryBrandPageHandler{DB: db}
}

// CheckUnique verifies name or slug uniqueness
// GET /api/category-brand-pages/check?field=name|slug&value=foo
func (h *CategoryBrandPageHandler) CheckUnique(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	if field != "name" && field != "slug" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	var count int64
	h.DB.Model(&models.CategoryBrandPage{}).Where(field+" = ?", value).Count(&count)
	c.JSON(http.StatusOK, gin.H{"exists": count > 0})
}

// CreatePage creates new category-brand page
func (h *CategoryBrandPageHandler) CreatePage(c *gin.Context) {
	var input struct {
		CategoryID  uint   `json:"category_id" binding:"required"`
		BrandID     uint   `json:"brand_id" binding:"required"`
		Name        string `json:"name" binding:"required"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	if input.Slug == "" {
		input.Slug = slug.Make(input.Name)
	}

	// uniqueness checks
	var exists int64
	h.DB.Model(&models.CategoryBrandPage{}).Where("name = ?", input.Name).Count(&exists)
	if exists > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_NAME", utils.GetErrorMessage("DUPLICATE_NAME"), nil))
		return
	}
	h.DB.Model(&models.CategoryBrandPage{}).Where("slug = ?", input.Slug).Count(&exists)
	if exists > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_SLUG", utils.GetErrorMessage("DUPLICATE_SLUG"), nil))
		return
	}

	page := models.CategoryBrandPage{
		CategoryID:  input.CategoryID,
		BrandID:     input.BrandID,
		Title:       input.Name,
		Slug:        input.Slug,
		Description: input.Description,
	}

	if err := h.DB.Create(&page).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusCreated, page)
}
