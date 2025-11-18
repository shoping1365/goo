package handlers

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
)

// AttributeValueHandler manages CRUD for attribute_values table.
type AttributeValueHandler struct {
	DB *gorm.DB
}

func NewAttributeValueHandler(db *gorm.DB) *AttributeValueHandler {
	return &AttributeValueHandler{DB: db}
}

// setAttributeTypeBasedOnValues بازخوانی مقادیر یک ویژگی و تعیین نوع بهینه
func (h *AttributeValueHandler) setAttributeTypeBasedOnValues(attributeID uint) {
	// شمارش بر اساس ستون‌های تایپی
	var cntNum, cntBool, cntDate, cntText int64
	_ = h.DB.Model(&models.AttributeValue{}).Where("attribute_id = ? AND numeric_value IS NOT NULL", attributeID).Count(&cntNum).Error
	_ = h.DB.Model(&models.AttributeValue{}).Where("attribute_id = ? AND bool_value IS NOT NULL", attributeID).Count(&cntBool).Error
	_ = h.DB.Model(&models.AttributeValue{}).Where("attribute_id = ? AND date_value IS NOT NULL", attributeID).Count(&cntDate).Error
	_ = h.DB.Model(&models.AttributeValue{}).Where("attribute_id = ? AND text_value IS NOT NULL", attributeID).Count(&cntText).Error

	// انتخاب نوع: اگر فقط یکی از انواع غیرمتنی وجود دارد همان، در غیر این صورت متن/رشته
	newType := "text"
	totalKinds := 0
	if cntNum > 0 {
		newType = "number"
		totalKinds++
	}
	if cntBool > 0 {
		if totalKinds > 0 {
			newType = "string"
		} else {
			newType = "boolean"
		}
		totalKinds++
	}
	if cntDate > 0 {
		if totalKinds > 0 {
			newType = "string"
		} else {
			newType = "date"
		}
		totalKinds++
	}
	if cntText > 0 {
		if totalKinds > 0 {
			newType = "string"
		} else {
			newType = "text"
		}
		totalKinds++
	}

	// به‌روزرسانی فقط در صورت تغییر
	_ = h.DB.Model(&models.Attribute{}).Where("id = ?", attributeID).Update("data_type", newType).Error
}

// normalizeString cleans a value to detect duplicates: trims, removes brackets, collapses spaces, lowercases.
func normalizeString(s string) string {
	// Remove common bracket chars
	re := regexp.MustCompile(`[\(\)\[\]\{\}«»]`)
	cleaned := re.ReplaceAllString(s, " ")
	cleaned = strings.TrimSpace(cleaned)
	cleaned = strings.Join(strings.Fields(cleaned), " ") // collapse spaces
	return strings.ToLower(cleaned)
}

// CreateAttributeValue POST /api/attributes/:attrId/values
func (h *AttributeValueHandler) CreateAttributeValue(c *gin.Context) {
	attrID, err := strconv.Atoi(c.Param("attrId"))
	if err != nil || attrID <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه ویژگی نامعتبر است", nil))
		return
	}

	// ensure attribute exists
	var attr models.Attribute
	if err := h.DB.First(&attr, attrID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "ویژگی یافت نشد", nil))
		return
	}

	// input payload
	var input struct {
		Value     string  `json:"value" binding:"required"`
		SortOrder *uint   `json:"sort_order"`
		Slug      *string `json:"slug"`
		Meta      *string `json:"meta"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}

	// --- Duplicate check ---
	canonical := normalizeString(input.Value)

	var existingVals []models.AttributeValue
	if err := h.DB.Where("attribute_id = ?", attrID).Find(&existingVals).Error; err == nil {
		for _, ev := range existingVals {
			if normalizeString(ev.Value) == canonical {
				c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_VALUE", "این مقدار قبلاً ثبت شده است", nil))
				return
			}
		}
	}

	// Validate numeric value if attribute expects number
	if attr.DataType == "number" {
		cleaned := strings.TrimSpace(input.Value)
		cleaned = strings.ReplaceAll(cleaned, ",", "")
		numRegex := regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]+$`)
		if !numRegex.MatchString(cleaned) {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "برای این ویژگی مقدار باید عددی باشد", nil))
			return
		}
	}

	// از این پس تشخیص نوع داده غیرفعال شده است؛ تمام مقادیر به‌صورت متن ذخیره می‌شوند.

	av := models.AttributeValue{
		AttributeID: uint(attrID),
		Value:       input.Value,
	}
	if input.SortOrder != nil {
		av.SortOrder = *input.SortOrder
	}
	if input.Slug != nil {
		av.Slug = *input.Slug
	}
	if input.Meta != nil {
		av.Meta = *input.Meta
	}

	// Ensure Meta contains valid JSON; empty string → '{}'
	if strings.TrimSpace(av.Meta) == "" {
		av.Meta = "{}"
	}

	// تشخیص خودکار نوع مقدار و پر کردن ستون بهینه
	typ, num, bl, dt := utils.DetectScalarType(input.Value)
	switch typ {
	case "number":
		av.NumericValue = num
	case "boolean":
		av.BoolValue = bl
	case "date":
		av.DateValue = dt
	default:
		txt := av.Value
		av.TextValue = &txt
	}

	if err := h.DB.Create(&av).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد مقدار", err.Error()))
		return
	}

	// همگام‌سازی نوع ویژگی بر اساس داده‌های فعلی
	go h.setAttributeTypeBasedOnValues(uint(attrID))

	c.JSON(http.StatusCreated, av)
}

// UpdateAttributeValue PUT /api/attribute-values/:id
func (h *AttributeValueHandler) UpdateAttributeValue(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه نامعتبر است", nil))
		return
	}
	var av models.AttributeValue
	if err := h.DB.First(&av, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "مقدار یافت نشد", nil))
		return
	}

	var input struct {
		Value     *string `json:"value"`
		SortOrder *uint   `json:"sort_order"`
		Slug      *string `json:"slug"`
		Meta      *string `json:"meta"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}

	if input.Value != nil {
		// Duplicate check (exclude current)
		canonical := normalizeString(*input.Value)
		var others []models.AttributeValue
		if err := h.DB.Where("attribute_id = ? AND id <> ?", av.AttributeID, av.ID).Find(&others).Error; err == nil {
			for _, ov := range others {
				if normalizeString(ov.Value) == canonical {
					c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_VALUE", "این مقدار قبلاً ثبت شده است", nil))
					return
				}
			}
		}

		// Validate numeric if needed
		if av.Attribute.DataType == "number" {
			cleaned := strings.TrimSpace(*input.Value)
			cleaned = strings.ReplaceAll(cleaned, ",", "")
			numRegex := regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]+$`)
			if !numRegex.MatchString(cleaned) {
				c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "برای این ویژگی مقدار باید عددی باشد", nil))
				return
			}
		}

		// reset typed fields
		av.TextValue = nil
		av.NumericValue = nil
		av.BoolValue = nil
		av.DateValue = nil

		av.Value = *input.Value

		// تشخیص خودکار
		typ, num, bl, dt := utils.DetectScalarType(*input.Value)
		switch typ {
		case "number":
			av.NumericValue = num
		case "boolean":
			av.BoolValue = bl
		case "date":
			av.DateValue = dt
		default:
			txt := *input.Value
			av.TextValue = &txt
		}
	}

	if input.SortOrder != nil {
		av.SortOrder = *input.SortOrder
	}
	if input.Slug != nil {
		av.Slug = *input.Slug
	}
	if input.Meta != nil {
		av.Meta = *input.Meta
	}

	// Ensure valid JSON in Meta
	if strings.TrimSpace(av.Meta) == "" {
		av.Meta = "{}"
	}

	if err := h.DB.Save(&av).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بروزرسانی مقدار", err.Error()))
		return
	}

	// همگام‌سازی نوع ویژگی بر اساس داده‌های فعلی
	go h.setAttributeTypeBasedOnValues(av.AttributeID)

	c.JSON(http.StatusOK, av)
}

// DeleteAttributeValue DELETE /api/attribute-values/:id
func (h *AttributeValueHandler) DeleteAttributeValue(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه نامعتبر است", nil))
		return
	}

	// Get attribute ID before deletion
	var av models.AttributeValue
	if err := h.DB.First(&av, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "مقدار یافت نشد", nil))
		return
	}
	attrID := av.AttributeID

	// Prevent deletion if attribute is color
	var attr models.Attribute
	if err := h.DB.First(&attr, attrID).Error; err == nil {
		if attr.DataType == "color" {
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "صفحه مورد نظر پیدا نشد", nil))
			return
		}
	}

	if err := h.DB.Delete(&models.AttributeValue{}, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف مقدار", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// ListAttributeValues GET /api/attributes/:attrId/values
func (h *AttributeValueHandler) ListAttributeValues(c *gin.Context) {
	attrID, err := strconv.Atoi(c.Param("attrId"))
	if err != nil || attrID <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه ویژگی نامعتبر است", nil))
		return
	}
	var vals []models.AttributeValue
	// تلاش اول: مرتب‌سازی بر اساس sort_order (در صورت هم‌تراز بودن اسکیمای DB)
	if err := h.DB.Where("attribute_id = ?", attrID).Order("sort_order asc, id asc").Find(&vals).Error; err != nil {
		// اگر ستون sort_order وجود نداشت (SQLSTATE 42703)، تلاش دوم فقط بر اساس id
		if strings.Contains(strings.ToLower(err.Error()), "sort_order") || strings.Contains(err.Error(), "42703") {
			if err2 := h.DB.Where("attribute_id = ?", attrID).Order("id asc").Find(&vals).Error; err2 != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در لیست مقادیر", err2.Error()))
				return
			}
			c.JSON(http.StatusOK, vals)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در لیست مقادیر", err.Error()))
		return
	}
	c.JSON(http.StatusOK, vals)
}
