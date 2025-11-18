package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// InventoryHandler handler برای موجودی
type InventoryHandler struct {
	inventoryService *services.InventoryService
}

// NewInventoryHandler ایجاد handler جدید
func NewInventoryHandler(inventoryService *services.InventoryService) *InventoryHandler {
	return &InventoryHandler{
		inventoryService: inventoryService,
	}
}

// GetInventoryStatus دریافت وضعیت موجودی یک محصول
func (h *InventoryHandler) GetInventoryStatus(c *gin.Context) {
	productIDStr := c.Param("product_id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "شناسه محصول نامعتبر است",
		})
		return
	}

	inventory, err := h.inventoryService.GetInventoryStatus(uint(productID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "موجودی محصول یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    inventory,
	})
}

// UpdateInventory به‌روزرسانی موجودی محصول
func (h *InventoryHandler) UpdateInventory(c *gin.Context) {
	productIDStr := c.Param("product_id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "شناسه محصول نامعتبر است",
		})
		return
	}

	var req struct {
		StockQuantity    *int `json:"stock_quantity"`
		ReservedQuantity *int `json:"reserved_quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "داده‌های ورودی نامعتبر است",
		})
		return
	}

	stockQty := 0
	if req.StockQuantity != nil {
		stockQty = *req.StockQuantity
	}

	reservedQty := 0
	if req.ReservedQuantity != nil {
		reservedQty = *req.ReservedQuantity
	}

	err = h.inventoryService.UpdateInventory(uint(productID), stockQty, reservedQty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در به‌روزرسانی موجودی",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "موجودی با موفقیت به‌روزرسانی شد",
	})
}

// CheckAndUpdateCartItems بررسی و به‌روزرسانی آیتم‌های سبد خرید
func (h *InventoryHandler) CheckAndUpdateCartItems(c *gin.Context) {
	var cartItems []models.OrderItem

	if err := c.ShouldBindJSON(&cartItems); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "داده‌های ورودی نامعتبر است",
		})
		return
	}

	updatedItems, unavailableMessages := h.inventoryService.CheckAndUpdateCartItems(cartItems)

	response := gin.H{
		"success": true,
		"data": gin.H{
			"items": updatedItems,
		},
	}

	if len(unavailableMessages) > 0 {
		response["warnings"] = unavailableMessages
	}

	c.JSON(http.StatusOK, response)
}

// GetLowStockProducts دریافت محصولات با موجودی کم
func (h *InventoryHandler) GetLowStockProducts(c *gin.Context) {
	products, err := h.inventoryService.GetLowStockProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در دریافت محصولات با موجودی کم",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    products,
	})
}

// GetOutOfStockProducts دریافت محصولات ناموجود
func (h *InventoryHandler) GetOutOfStockProducts(c *gin.Context) {
	products, err := h.inventoryService.GetOutOfStockProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در دریافت محصولات ناموجود",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    products,
	})
}
