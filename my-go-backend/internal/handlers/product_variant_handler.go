package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProductVariantHandler مدیریت عملیات مربوط به تنوع‌های محصول را بر عهده دارد
type ProductVariantHandler struct {
	DB *gorm.DB
}

// GetVariants لیست تنوع‌های یک محصول را برمی‌گرداند
func (h *ProductVariantHandler) GetVariants(c *gin.Context) {
	productIDStr := c.Param("productId")
	productID, _ := strconv.Atoi(productIDStr)

	var items []models.ProductVariant
	if err := h.DB.Where("product_id = ?", productID).Order("id ASC").Find(&items).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}

// AddVariant ایجاد یک تنوع جدید برای محصول
func (h *ProductVariantHandler) AddVariant(c *gin.Context) {
	productIDStr := c.Param("productId")
	productID, _ := strconv.Atoi(productIDStr)

	var v models.ProductVariant
	if err := c.ShouldBindJSON(&v); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	v.ProductID = productID
	if err := h.DB.Create(&v).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusCreated, v)
}

// UpdateVariant بروزرسانی اطلاعات یک تنوع
func (h *ProductVariantHandler) UpdateVariant(c *gin.Context) {
	variantIDStr := c.Param("variantId")
	variantID, _ := strconv.Atoi(variantIDStr)

	var v models.ProductVariant
	if err := c.ShouldBindJSON(&v); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	if err := h.DB.Model(&models.ProductVariant{}).Where("id = ?", variantID).Updates(v).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Variant updated"})
}

// DeleteVariant حذف یک تنوع
func (h *ProductVariantHandler) DeleteVariant(c *gin.Context) {
	variantIDStr := c.Param("variantId")
	variantID, _ := strconv.Atoi(variantIDStr)

	if err := h.DB.Delete(&models.ProductVariant{}, variantID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Variant deleted"})
}
