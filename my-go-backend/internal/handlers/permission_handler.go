package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
)

// PermissionHandler هندلر مدیریت دسترسی‌ها
type PermissionHandler struct {
	DB *gorm.DB
}

// NewPermissionHandler ایجاد نمونه جدید از PermissionHandler
func NewPermissionHandler(db *gorm.DB) *PermissionHandler {
	return &PermissionHandler{DB: db}
}

// ListPermissions لیست تمام دسترسی‌ها را برمی‌گرداند
func (h *PermissionHandler) ListPermissions(c *gin.Context) {
	var permissions []models.Permission

	query := h.DB.Order("module ASC, action ASC, resource ASC")

	// فیلتر بر اساس وضعیت فعال/غیرفعال
	if isActive := c.Query("is_active"); isActive != "" {
		if isActive == "true" {
			query = query.Where("is_active = ?", true)
		} else if isActive == "false" {
			query = query.Where("is_active = ?", false)
		}
	}

	// فیلتر بر اساس ماژول
	if module := c.Query("module"); module != "" {
		query = query.Where("module = ?", module)
	}

	// فیلتر بر اساس عملیات
	if action := c.Query("action"); action != "" {
		query = query.Where("action = ?", action)
	}

	// فیلتر بر اساس منبع
	if resource := c.Query("resource"); resource != "" {
		query = query.Where("resource = ?", resource)
	}

	// جستجو بر اساس نام
	if search := c.Query("search"); search != "" {
		query = query.Where("name LIKE ? OR display_name LIKE ? OR description LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Find(&permissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لیست دسترسی‌ها", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    permissions,
	})
}

// GetPermission دسترسی خاصی را بر اساس ID برمی‌گرداند
func (h *PermissionHandler) GetPermission(c *gin.Context) {
	id := c.Param("id")
	permissionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه دسترسی نامعتبر است", nil))
		return
	}

	var permission models.Permission
	if err := h.DB.First(&permission, permissionID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.New("PERMISSION_NOT_FOUND", "دسترسی مورد نظر یافت نشد", nil))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت دسترسی", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    permission,
	})
}

// CreatePermission دسترسی جدید ایجاد می‌کند
func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	var req models.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر است", err.Error()))
		return
	}

	// بررسی تکراری نبودن نام دسترسی
	var existingPermission models.Permission
	if err := h.DB.Where("name = ?", req.Name).First(&existingPermission).Error; err == nil {
		c.JSON(http.StatusBadRequest, utils.New("DUPLICATE_PERMISSION", "دسترسی با این نام قبلاً وجود دارد", nil))
		return
	}

	// ایجاد دسترسی جدید
	permission := models.Permission{
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Description: req.Description,
		Module:      req.Module,
		Action:      req.Action,
		Resource:    req.Resource,
		IsActive:    true,
	}

	if err := h.DB.Create(&permission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد دسترسی", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    permission,
		"message": "دسترسی با موفقیت ایجاد شد",
	})
}

// UpdatePermission دسترسی موجود را به‌روزرسانی می‌کند
func (h *PermissionHandler) UpdatePermission(c *gin.Context) {
	id := c.Param("id")
	permissionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه دسترسی نامعتبر است", nil))
		return
	}

	var req models.UpdatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر است", err.Error()))
		return
	}

	var permission models.Permission
	if err := h.DB.First(&permission, permissionID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.New("PERMISSION_NOT_FOUND", "دسترسی مورد نظر یافت نشد", nil))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت دسترسی", err.Error()))
		return
	}

	// به‌روزرسانی فیلدها
	if req.DisplayName != "" {
		permission.DisplayName = req.DisplayName
	}
	if req.Description != "" {
		permission.Description = req.Description
	}
	if req.IsActive != nil {
		permission.IsActive = *req.IsActive
	}

	if err := h.DB.Save(&permission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی دسترسی", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    permission,
		"message": "دسترسی با موفقیت به‌روزرسانی شد",
	})
}

// DeletePermission دسترسی را حذف می‌کند
func (h *PermissionHandler) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	permissionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه دسترسی نامعتبر است", nil))
		return
	}

	var permission models.Permission
	if err := h.DB.First(&permission, permissionID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.New("PERMISSION_NOT_FOUND", "دسترسی مورد نظر یافت نشد", nil))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت دسترسی", err.Error()))
		return
	}

	// بررسی اینکه آیا این دسترسی در حال استفاده است یا خیر
	var rolePermissionCount int64
	if err := h.DB.Model(&models.RolePermission{}).Where("permission_id = ?", permissionID).Count(&rolePermissionCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بررسی استفاده از دسترسی", err.Error()))
		return
	}

	if rolePermissionCount > 0 {
		c.JSON(http.StatusBadRequest, utils.New("PERMISSION_IN_USE", "این دسترسی در حال استفاده است و نمی‌توان آن را حذف کرد", gin.H{"role_count": rolePermissionCount}))
		return
	}

	// حذف دسترسی
	if err := h.DB.Delete(&permission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف دسترسی", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "دسترسی با موفقیت حذف شد",
	})
}

// GetPermissionModules لیست ماژول‌های موجود را برمی‌گرداند
func (h *PermissionHandler) GetPermissionModules(c *gin.Context) {
	var modules []string
	if err := h.DB.Model(&models.Permission{}).Distinct().Pluck("module", &modules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت ماژول‌ها", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    modules,
	})
}

// GetPermissionActions لیست عملیات‌های موجود را برمی‌گرداند
func (h *PermissionHandler) GetPermissionActions(c *gin.Context) {
	var actions []string
	if err := h.DB.Model(&models.Permission{}).Distinct().Pluck("action", &actions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت عملیات‌ها", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    actions,
	})
}

// GetPermissionResources لیست منابع موجود را برمی‌گرداند
func (h *PermissionHandler) GetPermissionResources(c *gin.Context) {
	var resources []string
	if err := h.DB.Model(&models.Permission{}).Distinct().Pluck("resource", &resources).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت منابع", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resources,
	})
}
