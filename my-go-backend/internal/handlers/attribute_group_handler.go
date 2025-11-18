package handlers

import (
	"errors"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
)

// AttributeGroupHandler handles CRUD operations for attribute groups.
// An attribute group belongs to exactly one product category and can hold multiple attributes.
//
// Endpoints (to be registered under /attribute-groups):
//   GET    /attribute-groups                list groups (pagination, optional search)
//   GET    /attribute-groups/:id            get single group (with attributes)
//   POST   /attribute-groups                create new group
//   PUT    /attribute-groups/:id            update group
//   DELETE /attribute-groups/:id            soft-delete group
//   POST   /attribute-groups/bulk-delete    bulk delete by IDs (optional)
//
// NOTE: access control (Auth/Admin) should be added via middleware at router level.

type AttributeGroupHandler struct {
	DB *gorm.DB
}

// NewAttributeGroupHandler constructs a handler instance.
func NewAttributeGroupHandler(db *gorm.DB) *AttributeGroupHandler {
	return &AttributeGroupHandler{DB: db}
}

// ListAttributeGroups returns paginated groups with optional search & category filter.
func (h *AttributeGroupHandler) ListAttributeGroups(c *gin.Context) {
	// Pagination params
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}

	query := h.DB.Model(&models.AttributeGroup{})

	// Category filter
	if catStr := c.Query("category_id"); catStr != "" {
		if catID, err := strconv.Atoi(catStr); err == nil {
			query = query.Where("category_id = ?", catID)
		}
	}

	// Search by name (case-insensitive)
	if search := strings.TrimSpace(c.Query("search")); search != "" {
		like := "%" + search + "%"
		query = query.Where("LOWER(name) LIKE LOWER(?)", like)
	}

	// Count total
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	// Pagination
	offset := (page - 1) * perPage

	type GroupResponse struct {
		models.AttributeGroup
		AttributesCount int    `json:"attributes_count"`
		CategoryName    string `json:"category_name"`
		Used            bool   `json:"used"`
	}

	var groups []models.AttributeGroup
	if err := query.Order("id DESC").Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).Limit(perPage).Offset(offset).Find(&groups).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	// Build response with counts
	resp := make([]GroupResponse, len(groups))
	for i, g := range groups {
		var cnt int64
		_ = h.DB.Model(&models.AttributeGroupAttribute{}).Where("attribute_group_id = ?", g.ID).Count(&cnt).Error

		// check if used in any product
		var usedCnt int64
		_ = h.DB.Table("product_attribute_values pav").
			Joins("JOIN attribute_group_attributes aga ON aga.attribute_id = pav.attribute_id").
			Where("aga.attribute_group_id = ?", g.ID).Limit(1).Count(&usedCnt).Error

		resp[i] = GroupResponse{AttributeGroup: g, AttributesCount: int(cnt), CategoryName: g.Category.Name, Used: usedCnt > 0}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
		"meta": gin.H{
			"total":       total,
			"page":        page,
			"per_page":    perPage,
			"total_pages": int(math.Ceil(float64(total) / float64(perPage))),
		},
	})
}

// GetAttributeGroup returns single group by ID with its attributes.
func (h *AttributeGroupHandler) GetAttributeGroup(c *gin.Context) {
	id := c.Param("id")
	var group models.AttributeGroup
	if err := h.DB.Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).Preload("Attributes.Attribute").First(&group, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}
	c.JSON(http.StatusOK, group)
}

// CreateAttributeGroup creates a new attribute group.
func (h *AttributeGroupHandler) CreateAttributeGroup(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required,min=2,max=100"`
		CategoryID  uint   `json:"category_id" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check duplicate in same category (case-insensitive)
	var dup int64
	if err := h.DB.Model(&models.AttributeGroup{}).
		Where("category_id = ? AND LOWER(name) = LOWER(?)", input.CategoryID, input.Name).
		Count(&dup).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	if dup > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_NAME", utils.GetErrorMessage("DUPLICATE_NAME"), nil))
		return
	}

	group := models.AttributeGroup{
		Name:        strings.TrimSpace(input.Name),
		CategoryID:  input.CategoryID,
		Description: input.Description,
	}

	if err := h.DB.Create(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ایجاد گروه ویژگی", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, group)
}

// UpdateAttributeGroup updates name or description of a group.
func (h *AttributeGroupHandler) UpdateAttributeGroup(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Name        string `json:"name" binding:"required,min=2,max=100"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var group models.AttributeGroup
	if err := h.DB.First(&group, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "گروه ویژگی یافت نشد"})
		return
	}

	// Check duplicates for name within same category
	var dup int64
	_ = h.DB.Model(&models.AttributeGroup{}).
		Where("category_id = ? AND LOWER(name) = LOWER(?) AND id <> ?", group.CategoryID, input.Name, group.ID).
		Count(&dup).Error
	if dup > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "نام گروه در این دسته‌بندی تکراری است"})
		return
	}

	group.Name = strings.TrimSpace(input.Name)
	group.Description = input.Description

	if err := h.DB.Save(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در بروزرسانی گروه ویژگی", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

// DeleteAttributeGroup performs a TRUE hard-delete: first removes join-table records
// then permanently deletes the group row (bypassing soft-delete). This guarantees
// عدم باقی ماندن رابطه‌های یتیم در attribute_group_attributes.
func (h *AttributeGroupHandler) DeleteAttributeGroup(c *gin.Context) {
	idStr := c.Param("id")

	tx := h.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "مشکل در آغاز تراکنش حذف"})
		return
	}

	// 1) پاک کردن رکوردهای واسط
	if err := tx.Where("attribute_group_id = ?", idStr).Unscoped().Delete(&models.AttributeGroupAttribute{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در حذف روابط گروه", "details": err.Error()})
		return
	}

	// 2) حذف کامل خود گروه (Unscoped => سخت)
	if err := tx.Unscoped().Delete(&models.AttributeGroup{}, idStr).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در حذف گروه", "details": err.Error()})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در نهایی‌سازی حذف", "details": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *AttributeGroupHandler) UpsertGroupAttributes(c *gin.Context) {
	groupIDStr := c.Param("id")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil || groupID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه گروه نامعتبر است"})
		return
	}

	// Expected payload
	type AttrInput struct {
		AttributeID   uint   `json:"attribute_id" binding:"required"`
		SortOrder     uint   `json:"sort_order"`
		Type          string `json:"type" binding:"required"`
		ControlType   string `json:"control_type" binding:"required"`
		HasFilter     bool   `json:"has_filter"`
		IsKey         bool   `json:"is_key"`
		ShowOnProduct bool   `json:"show_on_product"`
		IsRequired    bool   `json:"is_required"`
	}
	var payload struct {
		Attributes []AttrInput `json:"attributes" binding:"required,dive"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(payload.Attributes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "لیست ویژگی‌ها خالی است"})
		return
	}

	tx := h.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در آغاز تراکنش"})
		return
	}

	// Get old attribute types before removing
	var oldGroupAttrs []models.AttributeGroupAttribute
	if err := tx.Where("attribute_group_id = ?", groupID).Find(&oldGroupAttrs).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت ویژگی‌های قبلی", "details": err.Error()})
		return
	}

	// Create map of old types for comparison
	oldTypes := make(map[uint]string)
	for _, old := range oldGroupAttrs {
		oldTypes[old.AttributeID] = old.Type
	}

	// Remove old records
	if err := tx.Where("attribute_group_id = ?", groupID).Delete(&models.AttributeGroupAttribute{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در حذف ویژگی‌های قبلی", "details": err.Error()})
		return
	}

	// Insert new ones
	for _, in := range payload.Attributes {
		rec := models.AttributeGroupAttribute{
			AttributeGroupID: uint(groupID),
			AttributeID:      in.AttributeID,
			SortOrder:        in.SortOrder,
			Type:             in.Type,
			ControlType:      in.ControlType,
			HasFilter:        in.HasFilter,
			IsKey:            in.IsKey,
			ShowOnProduct:    in.ShowOnProduct,
			IsRequired:       in.IsRequired,
		}
		if err := tx.Create(&rec).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ذخیره ویژگی‌ها", "details": err.Error()})
			return
		}
	}

	// Check for type changes and convert attribute values
	for _, attr := range payload.Attributes {
		oldType := oldTypes[attr.AttributeID]
		newType := attr.Type

		// If changed from non-custom-text to custom-text, force convert to text
		if (oldType != "custom_text" && oldType != "متن سفارشی") &&
			(newType == "custom_text" || newType == "متن سفارشی") {
			if err := h.convertAttributeValuesToText(tx, attr.AttributeID); err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در تبدیل داده‌ها به متن", "details": err.Error()})
				return
			}
		}

		// تشخیص خودکار انواع داده حذف شده است؛ در هر دو حالت مقدارها به متن تبدیل می‌شوند و نوع attribute ثابت می‌ماند.

		// --- تعیین data_type مناسب بر اساس نوع کنترل ---
		if err := h.applyDataTypeRules(tx, attr.AttributeID, newType, attr.ControlType); err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ثبت تراکنش", "details": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListGroupsByAttribute returns all attribute groups that include the specified attribute (by attribute_id).
// Route: GET /attribute-groups/by-attribute/:attrId
func (h *AttributeGroupHandler) ListGroupsByAttribute(c *gin.Context) {
	attrIDStr := c.Param("attrId")
	attrID, err := strconv.Atoi(attrIDStr)
	if err != nil || attrID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه ویژگی نامعتبر است"})
		return
	}

	// Anchor the query to the AttributeGroup model to ensure the correct base table
	var groups []models.AttributeGroup
	if err := h.DB.Model(&models.AttributeGroup{}).
		Joins("JOIN attribute_group_attributes aga ON aga.attribute_group_id = attribute_groups.id").
		Where("aga.attribute_id = ?", attrID).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Order("attribute_groups.id DESC").
		Distinct(). // ensure unique groups in case of duplicate join rows
		Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت گروه‌ها", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

func (h *AttributeGroupHandler) RemoveAttributeFromGroup(c *gin.Context) {
	groupIDStr := c.Param("id")
	attrIDStr := c.Param("attrId")

	groupID, err1 := strconv.Atoi(groupIDStr)
	attrID, err2 := strconv.Atoi(attrIDStr)
	if err1 != nil || err2 != nil || groupID <= 0 || attrID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه‌ها نامعتبر هستند"})
		return
	}

	if err := h.DB.Where("attribute_group_id = ? AND attribute_id = ?", groupID, attrID).Delete(&models.AttributeGroupAttribute{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در حذف رابطه", "details": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// convertAttributeValuesToText converts all values of an attribute to text type
func (h *AttributeGroupHandler) convertAttributeValuesToText(tx *gorm.DB, attributeID uint) error {
	var values []models.AttributeValue
	if err := tx.Where("attribute_id = ?", attributeID).Find(&values).Error; err != nil {
		return err
	}

	for _, val := range values {
		// Set text_value to the original value string and clear others
		val.TextValue = &val.Value
		val.NumericValue = nil
		val.BoolValue = nil
		val.DateValue = nil

		if err := tx.Save(&val).Error; err != nil {
			return err
		}
	}

	// Update attribute data_type to text
	if err := tx.Model(&models.Attribute{}).Where("id = ?", attributeID).Update("data_type", utils.DataTypeText).Error; err != nil {
		return err
	}

	return nil
}

// applyDataTypeRules sets correct attributes.data_type for given attributeID based on group Type & ControlType.
func (h *AttributeGroupHandler) applyDataTypeRules(tx *gorm.DB, attributeID uint, typ string, control string) error {
	typLower := strings.ToLower(strings.TrimSpace(typ))
	ctrlLower := strings.ToLower(strings.TrimSpace(control))

	var newType string

	if typLower == "custom_text" || typLower == "متن سفارشی" {
		if strings.Contains(ctrlLower, "multi") || strings.Contains(ctrlLower, "textarea") {
			newType = "string"
		} else {
			newType = "text"
		}
	} else if strings.Contains(typLower, "select") || strings.Contains(typLower, "انتخاب") {
		// انتخابی یا چندانتخابی
		// Fetch values
		var vals []models.AttributeValue
		if err := tx.Where("attribute_id = ?", attributeID).Find(&vals).Error; err != nil {
			return err
		}
		allNumeric := true
		numRegex := regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]+$`)
		for _, v := range vals {
			s := strings.TrimSpace(v.Value)
			s = strings.ReplaceAll(s, ",", "")
			if s == "" || !numRegex.MatchString(s) {
				allNumeric = false
				break
			}
		}
		if allNumeric && len(vals) > 0 {
			newType = "number"
		} else {
			newType = "string"
		}
	} else {
		// fallback
		return nil
	}

	var attr models.Attribute
	if err := tx.First(&attr, attributeID).Error; err != nil {
		return err
	}

	if attr.DataType == newType {
		return nil // already correct
	}

	// if converting to number ensure all numeric
	if newType == "number" {
		var countNonNumeric int64
		tx.Model(&models.AttributeValue{}).Where("attribute_id = ?", attributeID).
			Where("NOT (value ~ '^[+-]?[0-9]*\\.?[0-9]+$')").Count(&countNonNumeric)
		if countNonNumeric > 0 {
			return errors.New("برخی مقادیر عددی نیستند، نمی‌توان نوع داده را به عدد تبدیل کرد")
		}
	}

	return tx.Model(&models.Attribute{}).Where("id = ?", attributeID).Update("data_type", newType).Error
}
