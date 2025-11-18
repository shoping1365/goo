package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SystemHandler هندلر عملیات سیستمی
type SystemHandler struct {
	db *gorm.DB
}

// NewSystemHandler ایجاد instance جدید از SystemHandler
func NewSystemHandler(db *gorm.DB) *SystemHandler {
	return &SystemHandler{db: db}
}

// CheckTable بررسی وجود جدول در دیتابیس
func (h *SystemHandler) CheckTable(c *gin.Context) {
	tableName := c.Param("tableName")
	if tableName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "نام جدول الزامی است",
		})
		return
	}

	// بررسی وجود جدول
	var exists bool
	err := h.db.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = ?)", tableName).Scan(&exists).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در بررسی وجود جدول",
			"error":   err.Error(),
		})
		return
	}

	// اگر جدول وجود دارد، تعداد رکوردها را نیز بررسی کن
	var recordCount int64
	if exists {
		err = h.db.Raw("SELECT COUNT(*) FROM " + tableName).Scan(&recordCount).Error
		if err != nil {
			recordCount = -1 // خطا در شمارش
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"exists":      exists,
		"tableName":   tableName,
		"recordCount": recordCount,
		"message":     fmt.Sprintf("جدول %s %s", tableName, map[bool]string{true: "وجود دارد", false: "وجود ندارد"}[exists]),
	})
}
