package handlers

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
)

// AttributeHandler handles CRUD operations for attributes (product specifications).
type AttributeHandler struct {
	DB *gorm.DB
}

// NewAttributeHandler constructs a new AttributeHandler.
func NewAttributeHandler(db *gorm.DB) *AttributeHandler {
	return &AttributeHandler{DB: db}
}

// rand.Seed is deprecated in Go 1.20+, the global random generator is automatically seeded
// func init() {
//	rand.Seed(time.Now().UnixNano())
// }

// ListAttributes returns a paginated list of attributes with optional search by name.
// Query params:
//   - search: keyword to search inside name (case-insensitive)
//   - page:   page number (1-based, default 1)
//   - per_page: items per page (default 10)
func (h *AttributeHandler) ListAttributes(c *gin.Context) {
	var (
		page    int
		perPage int
		err     error
	)

	// Pagination params
	pageStr := c.Query("page")
	perPageStr := c.Query("per_page")
	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
	}
	if perPageStr == "" {
		perPage = 10
	} else {
		perPage, err = strconv.Atoi(perPageStr)
		if err != nil || perPage < 1 {
			perPage = 10
		}
	}

	// Base query
	query := h.DB.Model(&models.Attribute{})

	// Search filter
	keyword := strings.TrimSpace(c.Query("search"))
	if keyword != "" {
		like := "%" + keyword + "%"
		query = query.Where("LOWER(name) LIKE LOWER(?)", like)
	}

	// Count total
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در شمارش ویژگی‌ها", err.Error()))
		return
	}

	// Pagination
	offset := (page - 1) * perPage
	var attributes []models.Attribute
	if err := query.Order("sort_order ASC, id DESC").Limit(perPage).Offset(offset).Find(&attributes).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت ویژگی‌ها", err.Error()))
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"data": attributes,
		"meta": gin.H{
			"total":       total,
			"page":        page,
			"per_page":    perPage,
			"total_pages": int(math.Ceil(float64(total) / float64(perPage))),
		},
	})
}

// GetStats returns aggregated statistics for attributes used by the dashboard.
func (h *AttributeHandler) GetStats(c *gin.Context) {
	// Total attributes
	var total int64
	h.DB.Model(&models.Attribute{}).Count(&total)

	// Active attributes
	var active int64
	h.DB.Model(&models.Attribute{}).Where("is_active = ?", true).Count(&active)

	// Total attribute values
	var totalValues int64
	h.DB.Model(&models.AttributeValue{}).Count(&totalValues)

	// Used attributes (those referenced by product_attribute_values)
	var used int64
	h.DB.Table("product_attribute_values").Distinct("attribute_id").Count(&used)

	c.JSON(http.StatusOK, gin.H{
		"total":        total,
		"active":       active,
		"total_values": totalValues,
		"used":         used,
	})
}

// GetAttribute returns single attribute by ID.
func (h *AttributeHandler) GetAttribute(c *gin.Context) {
	id := c.Param("id")
	var attribute models.Attribute
	if err := h.DB.Preload("Values").First(&attribute, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "ویژگی یافت نشد", nil))
		return
	}
	c.JSON(http.StatusOK, attribute)
}

// CreateAttribute creates a new attribute.
func (h *AttributeHandler) CreateAttribute(c *gin.Context) {
	var input struct {
		Name         string `json:"name" binding:"required"`
		DisplayName  string `json:"display_name"`
		Code         string `json:"code"`
		DataType     string `json:"data_type"`
		Unit         string `json:"unit"`
		SortOrder    uint   `json:"sort_order"`
		IsRequired   bool   `json:"is_required"`
		IsFilterable bool   `json:"is_filterable"`
		IsActive     bool   `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", err.Error(), nil))
		return
	}

	// --- Normalize name first ---
	canonicalName := strings.Join(strings.Fields(input.Name), " ")
	input.Name = canonicalName

	// Ignore any code sent by client; always auto-generate
	base := slug.Make(canonicalName)
	if base == "" {
		base = "attr"
	}
	input.Code = fmt.Sprintf("%s-%04d", base, rand.Intn(10000))

	// Ensure uniqueness (very unlikely to collide due to random suffix, but double-check)
	for i := 0; i < 10; i++ {
		var cnt int64
		// اگر ستون code در اسکیمای فعلی وجود نداشت، از بررسی عبور کن
		if err := h.DB.Model(&models.Attribute{}).Where("code = ?", input.Code).Count(&cnt).Error; err != nil {
			cnt = 0
		}
		if cnt == 0 {
			break
		}
		input.Code = fmt.Sprintf("%s-%04d", base, rand.Intn(10000))
	}

	// Check for duplicates (case-insensitive)
	var dupCount int64
	if err := h.DB.Model(&models.Attribute{}).
		Where("LOWER(name) = ?", strings.ToLower(canonicalName)).
		Count(&dupCount).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بررسی تکراری بودن نام ویژگی", err.Error()))
		return
	}
	if dupCount > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_NAME", "نام ویژگی تکراری است یا قبلاً ایجاد شده است", nil))
		return
	}

	// Default data_type to "text" if empty
	if strings.TrimSpace(input.DataType) == "" {
		input.DataType = "text"
	}

	// Prevent multiple color attributes
	if input.DataType == "color" {
		var cnt int64
		_ = h.DB.Model(&models.Attribute{}).Where("data_type = ?", "color").Count(&cnt).Error
		if cnt > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_INPUT", "ویژگی دیگری با نوع داده رنگ وجود دارد", nil))
			return
		}
	}

	attribute := models.Attribute{
		Name:         input.Name,
		DisplayName:  input.DisplayName,
		Code:         input.Code,
		DataType:     input.DataType,
		Unit:         input.Unit,
		SortOrder:    input.SortOrder,
		IsRequired:   input.IsRequired,
		IsFilterable: input.IsFilterable,
		IsActive:     input.IsActive,
	}

	if err := h.DB.Create(&attribute).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد ویژگی", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, attribute)
}

// UpdateAttribute updates an existing attribute by ID.
func (h *AttributeHandler) UpdateAttribute(c *gin.Context) {
	id := c.Param("id")
	var attribute models.Attribute
	if err := h.DB.First(&attribute, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "ویژگی یافت نشد", nil))
		return
	}

	var input struct {
		Name         *string `json:"name"`
		DisplayName  *string `json:"display_name"`
		Code         *string `json:"code"`
		DataType     *string `json:"data_type"`
		Unit         *string `json:"unit"`
		SortOrder    *uint   `json:"sort_order"`
		IsRequired   *bool   `json:"is_required"`
		IsFilterable *bool   `json:"is_filterable"`
		IsActive     *bool   `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", err.Error(), nil))
		return
	}

	// Apply changes
	if input.Name != nil {
		// Normalize name
		newName := strings.Join(strings.Fields(*input.Name), " ")

		// Check duplicate excluding current record
		var cnt int64
		if err := h.DB.Model(&models.Attribute{}).
			Where("LOWER(name) = ? AND id <> ?", strings.ToLower(newName), attribute.ID).
			Count(&cnt).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بررسی تکراری بودن نام ویژگی", err.Error()))
			return
		}
		if cnt > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_NAME", "نام ویژگی تکراری است یا قبلاً ایجاد شده است", nil))
			return
		}

		attribute.Name = newName
	}
	if input.DisplayName != nil {
		attribute.DisplayName = *input.DisplayName
	}
	if input.DataType != nil {
		// If attempting to set to 'color' ensure uniqueness
		if *input.DataType == "color" && attribute.DataType != "color" {
			var cnt int64
			_ = h.DB.Model(&models.Attribute{}).Where("data_type = ? AND id <> ?", "color", attribute.ID).Count(&cnt).Error
			if cnt > 0 {
				c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_INPUT", "ویژگی دیگری با نوع داده رنگ وجود دارد", nil))
				return
			}
		}
		attribute.DataType = *input.DataType
	}
	if input.Unit != nil {
		attribute.Unit = *input.Unit
	}
	if input.SortOrder != nil {
		attribute.SortOrder = *input.SortOrder
	}
	if input.IsRequired != nil {
		attribute.IsRequired = *input.IsRequired
	}
	if input.IsFilterable != nil {
		attribute.IsFilterable = *input.IsFilterable
	}
	if input.IsActive != nil {
		attribute.IsActive = *input.IsActive
	}

	if err := h.DB.Save(&attribute).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی ویژگی", err.Error()))
		return
	}

	c.JSON(http.StatusOK, attribute)
}

// DeleteAttribute performs a hard-delete of the attribute by ID.
func (h *AttributeHandler) DeleteAttribute(c *gin.Context) {
	id := c.Param("id")

	// Prevent deletion of color attribute
	var attr models.Attribute
	if err := h.DB.First(&attr, id).Error; err == nil {
		if attr.DataType == "color" {
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "صفحه مورد نظر پیدا نشد", nil))
			return
		}
	}

	// Hard-delete attribute values first (to respect FK if not ON DELETE CASCADE)
	_ = h.DB.Unscoped().Where("attribute_id = ?", id).Delete(&models.AttributeValue{}).Error

	// Hard-delete attribute itself
	if err := h.DB.Unscoped().Delete(&models.Attribute{}, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف ویژگی", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ویژگی با موفقیت حذف شد"})
}

// BulkDelete deletes multiple attributes based on provided IDs.
// Request body: {"ids": [1,2,3]}
func (h *AttributeHandler) BulkDelete(c *gin.Context) {
	var payload struct {
		IDs []uint `json:"ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", err.Error(), nil))
		return
	}
	if len(payload.IDs) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "هیچ شناسه‌ای ارسال نشده است", nil))
		return
	}

	// Remove color attribute id if exists
	var colorAttr models.Attribute
	if err := h.DB.Where("data_type = ?", "color").First(&colorAttr).Error; err == nil {
		filtered := make([]uint, 0, len(payload.IDs))
		for _, id := range payload.IDs {
			if id != colorAttr.ID {
				filtered = append(filtered, id)
			}
		}
		payload.IDs = filtered
		if len(payload.IDs) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "هیچ آیتمی برای حذف مجاز نیست"})
			return
		}
	}

	// Delete child values
	_ = h.DB.Unscoped().Where("attribute_id IN ?", payload.IDs).Delete(&models.AttributeValue{}).Error

	if err := h.DB.Unscoped().Delete(&models.Attribute{}, payload.IDs).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف گروهی ویژگی‌ها", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ویژگی‌های انتخاب شده حذف شدند"})
}
