package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
)

// ProductShippingHandler مدیریت مقادیر حمل‌ونقل و ابعاد محصول را بر عهده دارد.
type ProductShippingHandler struct {
	DB *gorm.DB
}

// NewProductShippingHandler سازنده‌ی هندلر
func NewProductShippingHandler(db *gorm.DB) *ProductShippingHandler {
	return &ProductShippingHandler{DB: db}
}

// GetShipping اطلاعات حمل‌ونقل/ابعاد یک محصول را برمی‌گرداند – GET /product-shipping/:id
func (h *ProductShippingHandler) GetShipping(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := h.DB.First(&product, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "محصول پیدا نشد", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":            product.ID,
		"weight":        product.Weight,
		"length":        product.Length,
		"width":         product.Width,
		"height":        product.Height,
		"shipping_cost": product.ShippingCost,
		"shipping_time": product.ShippingTime,
	})
}

// updateShippingInput ورودی مجاز برای بروزرسانی مقادیر حمل‌ونقل/ابعاد
// از پوینتر استفاده شده تا فیلدهای ارسال‌نشده بدون تغییر باقی بمانند
 type updateShippingInput struct {
	Weight       *float64 `json:"weight"`
	Length       *float64 `json:"length"`
	Width        *float64 `json:"width"`
	Height       *float64 `json:"height"`
	ShippingCost *float64 `json:"shipping_cost"`
	ShippingTime *int     `json:"shipping_time"`
}

// UpdateShipping بروزرسانی مقادیر حمل‌ونقل/ابعاد – PUT /product-shipping/:id
func (h *ProductShippingHandler) UpdateShipping(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var product models.Product
	if err := h.DB.First(&product, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "محصول پیدا نشد", err.Error()))
		return
	}

	var input updateShippingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("BIND_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}

	if input.Weight != nil {
		product.Weight = *input.Weight
	}
	if input.Length != nil {
		product.Length = *input.Length
	}
	if input.Width != nil {
		product.Width = *input.Width
	}
	if input.Height != nil {
		product.Height = *input.Height
	}
	if input.ShippingCost != nil {
		product.ShippingCost = *input.ShippingCost
	}
	if input.ShippingTime != nil {
		product.ShippingTime = *input.ShippingTime
	}

	if err := h.DB.Save(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ذخیره اطلاعات حمل‌ونقل", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "اطلاعات حمل‌ونقل با موفقیت بروزرسانی شد",
		"product": gin.H{
			"id":            product.ID,
			"weight":        product.Weight,
			"length":        product.Length,
			"width":         product.Width,
			"height":        product.Height,
			"shipping_cost": product.ShippingCost,
			"shipping_time": product.ShippingTime,
		},
	})
}