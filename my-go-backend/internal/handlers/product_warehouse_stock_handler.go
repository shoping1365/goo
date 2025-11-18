package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProductWarehouseStockHandler مدیریت موجودی محصول به تفکیک انبار
// این هندلر فقط نقش هماهنگ‌کننده را دارد و لاجیک در سرویس/ریپازیتوری پیاده شده است.
type ProductWarehouseStockHandler struct {
	uowf unitofwork.UnitOfWorkFactory
	DB   *gorm.DB
}

func NewProductWarehouseStockHandler(uowf unitofwork.UnitOfWorkFactory, db *gorm.DB) *ProductWarehouseStockHandler {
	return &ProductWarehouseStockHandler{uowf: uowf, DB: db}
}

// GetStocksByProduct دریافت لیست موجودی‌های انبار برای یک محصول
// مسیر: GET /product-warehouse-stocks/:productId
func (h *ProductWarehouseStockHandler) GetStocksByProduct(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("productId"))
	if pid <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "شناسه نامعتبر"})
		return
	}

	uow := h.uowf.Create()
	list, err := uow.ProductWarehouseStockRepository().GetByProduct(c.Request.Context(), uint(pid))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت موجودی‌ها"})
		return
	}

	// دریافت لیست انبارها برای اضافه کردن نام انبار
	warehouses, err := uow.WarehouseRepository().GetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت انبارها"})
		return
	}

	// ایجاد map برای دسترسی سریع به نام انبارها
	warehouseMap := make(map[uint]string)
	for _, w := range warehouses {
		warehouseMap[w.ID] = w.Name
	}

	// اضافه کردن نام انبار به هر رکورد موجودی
	var result []gin.H
	for _, stock := range list {
		result = append(result, gin.H{
			"id":             stock.ID,
			"product_id":     stock.ProductID,
			"warehouse_id":   stock.WarehouseID,
			"warehouse_name": warehouseMap[stock.WarehouseID], // اضافه کردن نام انبار
			"quantity":       stock.Quantity,
			"reserved":       stock.Reserved,
			"min_qty":        stock.MinQty,
			"max_qty":        stock.MaxQty,
			"updated_at":     stock.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// PutStocksByProduct ثبت گروهی موجودی انبارهای محصول
// مسیر: PUT /product-warehouse-stocks/:productId
func (h *ProductWarehouseStockHandler) PutStocksByProduct(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("productId"))
	if pid <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "شناسه نامعتبر"})
		return
	}

	var in struct {
		Entries []struct {
			WarehouseID uint `json:"warehouse_id"`
			Quantity    int  `json:"quantity"`
			MinQty      int  `json:"min_qty"`
			MaxQty      int  `json:"max_qty"`
		} `json:"entries"`
		DefaultWarehouseID *uint `json:"default_warehouse_id"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر: " + err.Error()})
		return
	}

	invSvc := services.NewInventoryService(repository.NewInventoryRepository(h.DB), repository.NewOrderRepository(h.DB))
	// تبدیل ساختار
	entries := make([]struct {
		WarehouseID uint
		Quantity    int
		MinQty      int
		MaxQty      int
	}, 0, len(in.Entries))
	for _, e := range in.Entries {
		if e.WarehouseID == 0 {
			continue // نادیده گرفتن شناسه انبار نامعتبر
		}
		entries = append(entries, struct {
			WarehouseID uint
			Quantity    int
			MinQty      int
			MaxQty      int
		}{e.WarehouseID, e.Quantity, e.MinQty, e.MaxQty})
	}
	if len(entries) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "ورودی خالی/نامعتبر (warehouse_id)"})
		return
	}

	// تبدیل ساختار برای متد SetProductStocks
	stockEntries := make([]struct {
		WarehouseID uint
		ProductID   uint
		Quantity    int
		MinQty      int
		MaxQty      int
	}, len(entries))

	for i, entry := range entries {
		stockEntries[i] = struct {
			WarehouseID uint
			ProductID   uint
			Quantity    int
			MinQty      int
			MaxQty      int
		}{
			WarehouseID: entry.WarehouseID,
			ProductID:   uint(pid),
			Quantity:    entry.Quantity,
			MinQty:      entry.MinQty,
			MaxQty:      entry.MaxQty,
		}
	}

	if err := invSvc.SetProductStocks(c.Request.Context(), stockEntries); err != nil {
		// لاگ کردن خطای دقیق برای دیباگ
		fmt.Printf("خطا در SetProductStocks: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "خطا در ثبت موجودی: " + err.Error(),
			"details": err.Error(),
		})
		return
	}

	// تنظیم انبار پیش‌فرض خرید (اختیاری)
	if in.DefaultWarehouseID != nil {
		uow := h.uowf.Create()
		if err := uow.BeginTx(c); err == nil {
			if p, err := uow.ProductRepository().GetByID(c, uint(pid)); err == nil {
				p.DefaultWarehouseID = in.DefaultWarehouseID
				_ = uow.ProductRepository().Update(c, p)
				_ = uow.Commit()
			} else {
				_ = uow.Rollback()
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "ثبت شد"})
}

// AdjustStock افزایش/کاهش موجودی یک انبار برای محصول
// مسیر: POST /product-warehouse-stocks/:productId/adjust
func (h *ProductWarehouseStockHandler) AdjustStock(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("productId"))
	if pid <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "شناسه نامعتبر"})
		return
	}
	var in struct {
		WarehouseID uint   `json:"warehouse_id"`
		Delta       int    `json:"delta"`
		Reason      string `json:"reason"`
	}
	if err := c.ShouldBindJSON(&in); err != nil || in.WarehouseID == 0 || in.Delta == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر"})
		return
	}
	invSvc := services.NewInventoryService(repository.NewInventoryRepository(h.DB), repository.NewOrderRepository(h.DB))
	if err := invSvc.AdjustStock(c, uint(pid), in.WarehouseID, in.Delta, in.Reason, nil); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ثبت شد"})
}
