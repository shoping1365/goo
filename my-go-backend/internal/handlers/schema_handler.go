package handlers

import (
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SchemaHandler هندلر مدیریت اسکیمای پیش‌فرض
type SchemaHandler struct {
	schemaService *services.SchemaService
}

// NewSchemaHandler ایجاد نمونه جدید از SchemaHandler
func NewSchemaHandler(schemaService *services.SchemaService) *SchemaHandler {
	return &SchemaHandler{
		schemaService: schemaService,
	}
}

// GetAllTemplates دریافت تمام تمپلیت‌های فعال
func (h *SchemaHandler) GetAllTemplates(c *gin.Context) {
	schemas, err := h.schemaService.GetAllTemplates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "خطا در دریافت تمپلیت‌ها: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    schemas,
		"count":   len(schemas),
	})
}

// GetTemplatesByType دریافت تمپلیت‌ها بر اساس نوع
func (h *SchemaHandler) GetTemplatesByType(c *gin.Context) {
	schemaType := c.Param("type")
	if schemaType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "نوع اسکیما الزامی است",
		})
		return
	}

	schemas, err := h.schemaService.GetTemplatesByType(schemaType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "خطا در دریافت تمپلیت‌ها: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    schemas,
		"count":   len(schemas),
		"type":    schemaType,
	})
}

// GetTemplateByID دریافت تمپلیت بر اساس ID
func (h *SchemaHandler) GetTemplateByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "شناسه نامعتبر است",
		})
		return
	}

	schema, err := h.schemaService.GetTemplateByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "تمپلیت یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    schema,
	})
}

// CreateTemplate ایجاد تمپلیت جدید
func (h *SchemaHandler) CreateTemplate(c *gin.Context) {
	var schema models.Schema
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "داده‌های ورودی نامعتبر است: " + err.Error(),
		})
		return
	}

	// اعتبارسنجی تمپلیت
	if err := h.schemaService.ValidateTemplate(&schema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// دریافت شناسه کاربر از context (باید از middleware auth گرفته شود)
	userID := uint(1) // موقت - باید از context گرفته شود

	if err := h.schemaService.CreateTemplate(&schema, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "خطا در ایجاد تمپلیت: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "تمپلیت با موفقیت ایجاد شد",
		"data":    schema,
	})
}

// UpdateTemplate به‌روزرسانی تمپلیت
func (h *SchemaHandler) UpdateTemplate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "شناسه نامعتبر است",
		})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "داده‌های ورودی نامعتبر است: " + err.Error(),
		})
		return
	}

	// دریافت شناسه کاربر از context
	userID := uint(1) // موقت - باید از context گرفته شود

	if err := h.schemaService.UpdateTemplate(uint(id), updates, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "خطا در به‌روزرسانی تمپلیت: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "تمپلیت با موفقیت به‌روزرسانی شد",
	})
}

// DeleteTemplate حذف تمپلیت
func (h *SchemaHandler) DeleteTemplate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "شناسه نامعتبر است",
		})
		return
	}

	if err := h.schemaService.DeleteTemplate(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "خطا در حذف تمپلیت: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "تمپلیت با موفقیت حذف شد",
	})
}

// ToggleTemplateStatus تغییر وضعیت فعال/غیرفعال تمپلیت
func (h *SchemaHandler) ToggleTemplateStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "شناسه نامعتبر است",
		})
		return
	}

	// دریافت شناسه کاربر از context
	userID := uint(1) // موقت - باید از context گرفته شود

	if err := h.schemaService.ToggleTemplateStatus(uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "خطا در تغییر وضعیت تمپلیت: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "وضعیت تمپلیت با موفقیت تغییر کرد",
	})
}

// GenerateSchemaFromTemplate تولید اسکیما از تمپلیت
func (h *SchemaHandler) GenerateSchemaFromTemplate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "شناسه نامعتبر است",
		})
		return
	}

	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "داده‌های ورودی نامعتبر است: " + err.Error(),
		})
		return
	}

	schema, err := h.schemaService.GenerateSchemaFromTemplate(uint(id), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "خطا در تولید اسکیما: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "اسکیما با موفقیت تولید شد",
		"data":    schema,
	})
}

// GetSchemaTypes دریافت انواع مختلف اسکیما
func (h *SchemaHandler) GetSchemaTypes(c *gin.Context) {
	types := h.schemaService.GetSchemaTypes()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    types,
		"count":   len(types),
	})
}
