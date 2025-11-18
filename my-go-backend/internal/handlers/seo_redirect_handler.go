package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
)

// SEORedirectHandler هندلر مدیریت ریدایرکت‌های سئو
type SEORedirectHandler struct {
	DB *gorm.DB
}

func NewSEORedirectHandler(db *gorm.DB) *SEORedirectHandler {
	return &SEORedirectHandler{DB: db}
}

// ListRedirects GET /api/admin/seo/redirects
// پارامترهای اختیاری: group, q (جستجو)، limit، offset
func (h *SEORedirectHandler) ListRedirects(c *gin.Context) {
	group := strings.TrimSpace(c.Query("group"))
	q := strings.TrimSpace(c.Query("q"))
	limit := 200
	offset := 0
	if v := strings.TrimSpace(c.Query("limit")); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 1000 {
			limit = n
		}
	}
	if v := strings.TrimSpace(c.Query("offset")); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			offset = n
		}
	}

	var rows []models.SEORedirect
	query := h.DB.Model(&models.SEORedirect{})
	if group != "" {
		query = query.Where("group_name = ?", group)
	}
	if q != "" {
		like := "%" + q + "%"
		query = query.Where("source_path ILIKE ? OR target_path ILIKE ?", like, like)
	}
	if err := query.Order("id DESC").Limit(limit).Offset(offset).Find(&rows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	// آمار گروه‌ها برای نمایش در UI
	type grp struct {
		GroupName string
		Cnt       int64
	}
	var groups []grp
	_ = h.DB.Model(&models.SEORedirect{}).Select("group_name, COUNT(1) AS cnt").Group("group_name").Order("group_name").Scan(&groups)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rows,
		"groups":  groups,
		"count":   len(rows),
	})
}

// CreateRedirect POST /api/admin/seo/redirects
// ایجاد ریدایرکت جدید
func (h *SEORedirectHandler) CreateRedirect(c *gin.Context) {
	fmt.Printf("🔍 CreateRedirect called\n")

	var req struct {
		SourcePath string `json:"source_path" binding:"required"`
		TargetPath string `json:"target_path" binding:"required"`
		Code       int    `json:"code"`
		GroupName  string `json:"group_name"`
	}

	fmt.Printf("📝 Attempting to bind JSON...\n")
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("❌ JSON binding failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "داده‌های ورودی نامعتبر: " + err.Error(),
		})
		return
	}

	fmt.Printf("✅ JSON bound successfully: %+v\n", req)

	// تنظیم مقادیر پیش‌فرض
	if req.Code == 0 {
		req.Code = 301
	}
	if req.GroupName == "" {
		req.GroupName = "دسته‌بندی جدید"
	}

	// بررسی وجود source_path تکراری
	var existingRedirect models.SEORedirect
	if err := h.DB.Where("source_path = ?", req.SourcePath).First(&existingRedirect).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "ریدایرکت با این مسیر مبدا قبلاً وجود دارد",
		})
		return
	}

	// ایجاد ریدایرکت جدید
	redirect := models.SEORedirect{
		SourcePath:   req.SourcePath,
		TargetPath:   req.TargetPath,
		Code:         req.Code,
		GroupName:    req.GroupName,
		RedirectType: "permanent",
		VisitCount:   0,
	}

	if err := h.DB.Create(&redirect).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در ایجاد ریدایرکت: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "ریدایرکت با موفقیت ایجاد شد",
		"data":    redirect,
	})
}

// DeleteRedirect DELETE /api/admin/seo/redirects/:id
// حذف ریدایرکت با ID مشخص
func (h *SEORedirectHandler) DeleteRedirect(c *gin.Context) {
	fmt.Printf("🗑️ DeleteRedirect called\n")

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID ریدایرکت نامعتبر است",
		})
		return
	}

	// بررسی وجود ریدایرکت
	var redirect models.SEORedirect
	if err := h.DB.First(&redirect, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "ریدایرکت یافت نشد",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در جستجوی ریدایرکت: " + err.Error(),
		})
		return
	}

	// حذف ریدایرکت
	if err := h.DB.Delete(&redirect).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در حذف ریدایرکت: " + err.Error(),
		})
		return
	}

	fmt.Printf("✅ Redirect %d deleted successfully\n", id)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ریدایرکت با موفقیت حذف شد",
	})
}

// BulkDeleteRedirects DELETE /api/admin/seo/redirects/bulk-delete
// حذف دسته‌جمعی ریدایرکت‌ها
func (h *SEORedirectHandler) BulkDeleteRedirects(c *gin.Context) {
	fmt.Printf("🗑️ BulkDeleteRedirects called\n")

	var req struct {
		Ids []uint `json:"ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("❌ JSON binding failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "داده‌های ورودی نامعتبر: " + err.Error(),
		})
		return
	}

	if len(req.Ids) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "لیست ID ها نمی‌تواند خالی باشد",
		})
		return
	}

	fmt.Printf("📝 IDs to delete: %v\n", req.Ids)

	result := h.DB.Delete(&models.SEORedirect{}, req.Ids)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در حذف دسته‌جمعی: " + result.Error.Error(),
		})
		return
	}

	fmt.Printf("✅ %d redirects deleted successfully\n", result.RowsAffected)
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"message":      fmt.Sprintf("%d ریدایرکت با موفقیت حذف شد", result.RowsAffected),
		"deletedCount": result.RowsAffected,
	})
}
