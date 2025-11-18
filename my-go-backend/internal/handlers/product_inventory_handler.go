package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProductInventoryHandler مدیریت موجودی محصولات
type ProductInventoryHandler struct {
	uowf unitofwork.UnitOfWorkFactory
	DB   *gorm.DB
}

func NewProductInventoryHandler(uowf unitofwork.UnitOfWorkFactory, db *gorm.DB) *ProductInventoryHandler {
	return &ProductInventoryHandler{uowf: uowf, DB: db}
}

// GetInventory دریافت اطلاعات موجودی محصول
func (h *ProductInventoryHandler) GetInventory(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("id"))
	if pid <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "شناسه نامعتبر"})
		return
	}

	uow := h.uowf.Create()
	inventoryRepo := repository.NewInventoryRepository(h.DB)
	orderRepo := repository.NewOrderRepository(h.DB)
	invSvc := services.NewInventoryService(inventoryRepo, orderRepo)

	// همگام‌سازی موجودی محصول با انبارها
	_ = invSvc.SyncProductStockFromWarehouses(c.Request.Context(), uint(pid))

	// دریافت اطلاعات محصول
	product, err := uow.ProductRepository().GetByID(c.Request.Context(), uint(pid))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "محصول یافت نشد"})
		return
	}

	// دریافت انبار پیش‌فرض
	var defaultWarehouseID *uint
	if product.DefaultWarehouseID != nil {
		defaultWarehouseID = product.DefaultWarehouseID
	} else {
		// اگر انبار پیش‌فرض تعیین نشده، انبار پیش‌فرض سیستم را ست کن
		warehouses, err := uow.WarehouseRepository().GetAll(c.Request.Context())
		if err == nil && len(warehouses) > 0 {
			for _, w := range warehouses {
				if w.IsDefault {
					defaultWarehouseID = &w.ID
					// به‌روزرسانی محصول با انبار پیش‌فرض
					product.DefaultWarehouseID = &w.ID
					product.UpdatedAt = time.Now()
					_ = uow.ProductRepository().Update(c.Request.Context(), product)
					break
				}
			}
			// اگر هیچ انبار پیش‌فرضی نبود، اولین انبار را انتخاب کن
			if defaultWarehouseID == nil {
				defaultWarehouseID = &warehouses[0].ID
				product.DefaultWarehouseID = defaultWarehouseID
				product.UpdatedAt = time.Now()
				_ = uow.ProductRepository().Update(c.Request.Context(), product)
			}
		}
	}

	// دریافت موجودی کلی از انبارها (نه از جدول products)
	var totalStock int
	var totalReserved int

	// همگام‌سازی موجودی محصول با انبارها (برای اطمینان از به‌روز بودن)
	_ = invSvc.SyncProductStockFromWarehouses(c.Request.Context(), uint(pid))

	// دریافت موجودی به‌روز شده
	inventory, err := invSvc.GetInventoryStatus(uint(pid))
	if err == nil && inventory != nil {
		totalStock = inventory.StockQuantity
		totalReserved = inventory.ReservedQuantity
	} else {
		// اگر موجودی پیدا نشد، از انبارها محاسبه کن
		warehouseStockRepo := repository.NewProductWarehouseStockRepository(h.DB)
		totalStock, totalReserved, _ = warehouseStockRepo.SumByProduct(c.Request.Context(), uint(pid))
	}

	// ساخت پاسخ
	response := gin.H{
		"id":                     product.ID,
		"stock_quantity":         totalStock,
		"reserved_quantity":      totalReserved,
		"available_quantity":     totalStock - totalReserved,
		"min_stock_quantity":     product.MinStockQuantity,
		"max_stock_quantity":     product.MaxStockQuantity,
		"stock_status":           product.StockStatus,
		"track_inventory":        product.TrackInventory,
		"show_stock_to_customer": product.ShowStockToCustomer,
		"allow_reservation":      product.AllowReservation,
		"default_warehouse_id":   defaultWarehouseID,
	}

	c.JSON(http.StatusOK, response)

	// ارسال اعلان کمبود موجودی در پس‌زمینه (اگر لازم باشد)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if product.StockQuantity <= product.MinStockQuantity && product.StockQuantity > 0 {
			_ = uow.StockNotificationRepository().Create(ctx, &models.StockNotification{
				ProductID: uint(pid),
				Status:    "pending",
			})
		}
	}()
}

// UpdateInventory بروزرسانی تنظیمات سطح محصول (نه سطوح انبار)
func (h *ProductInventoryHandler) UpdateInventory(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("id"))
	if pid <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "شناسه نامعتبر"})
		return
	}
	var in struct {
		StockQuantity       *int    `json:"stock_quantity"`
		MinStockQuantity    *int    `json:"min_stock_quantity"`
		MaxStockQuantity    *int    `json:"max_stock_quantity"`
		StockStatus         *string `json:"stock_status"`
		ShowStockToCustomer *bool   `json:"show_stock_to_customer"`
		TrackInventory      *bool   `json:"track_inventory"`
		AllowReservation    *bool   `json:"allow_reservation"`
		// shipping_enabled فقط در فرانت استفاده می‌شود؛ در این نسخه ذخیره نمی‌کنیم
		ShippingEnabled *bool `json:"shipping_enabled"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر"})
		return
	}

	uow := h.uowf.Create()
	if err := uow.BeginTx(c); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "خطا در تراکنش"})
		return
	}
	defer func() { _ = uow.Rollback() }()

	product, err := uow.ProductRepository().GetByID(c, uint(pid))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "محصول پیدا نشد"})
		return
	}
	if in.StockQuantity != nil {
		product.StockQuantity = *in.StockQuantity
	}
	if in.MinStockQuantity != nil {
		product.MinStockQuantity = *in.MinStockQuantity
	}
	if in.MaxStockQuantity != nil {
		product.MaxStockQuantity = *in.MaxStockQuantity
	}
	if in.StockStatus != nil {
		product.StockStatus = *in.StockStatus
	}
	if in.ShowStockToCustomer != nil {
		product.ShowStockToCustomer = *in.ShowStockToCustomer
	}
	if in.TrackInventory != nil {
		product.TrackInventory = *in.TrackInventory
	}
	if in.AllowReservation != nil {
		product.AllowReservation = *in.AllowReservation
	}
	product.UpdatedAt = time.Now()
	if err := uow.ProductRepository().Update(c, product); err != nil {
		// لاگ کردن خطای دقیق برای دیباگ
		fmt.Printf("خطا در به‌روزرسانی محصول: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "خطا در ذخیره",
			"details": err.Error(), // اضافه کردن جزئیات خطا
		})
		return
	}
	if err := uow.Commit(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "خطا در تایید"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "بروزرسانی شد", "product": product})
}
