package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// WarehouseHandler هندلر CRUD انبارها
type WarehouseHandler struct{ uowf unitofwork.UnitOfWorkFactory }

func NewWarehouseHandler(uowf unitofwork.UnitOfWorkFactory) *WarehouseHandler {
	return &WarehouseHandler{uowf: uowf}
}

// List لیست انبارها
func (h *WarehouseHandler) List(c *gin.Context) {
	uow := h.uowf.Create()
	warehouses, err := uow.WarehouseRepository().GetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت انبارها"})
		return
	}

	// اگر هیچ انباری وجود ندارد، یک انبار پیشفرض ایجاد کن
	if len(warehouses) == 0 {
		defaultWarehouse := &models.Warehouse{
			Code:      "DEFAULT",
			Name:      "Default Warehouse",
			City:      "تهران",
			IsDefault: true,
			IsActive:  true,
			Priority:  100,
		}
		if err := uow.WarehouseRepository().Create(c.Request.Context(), defaultWarehouse); err == nil {
			// لیست انبارها را دوباره دریافت کن
			warehouses, _ = uow.WarehouseRepository().GetAll(c.Request.Context())
		}
	}

	c.JSON(http.StatusOK, warehouses)
}

// Create ایجاد انبار
func (h *WarehouseHandler) Create(c *gin.Context) {
	var in struct {
		Code      string `json:"code" binding:"required"`
		Name      string `json:"name" binding:"required"`
		City      string `json:"city"`
		Address   string `json:"address"`
		IsDefault bool   `json:"is_default"`
		Priority  int    `json:"priority"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر"})
		return
	}

	uow := h.uowf.Create()
	_ = uow.BeginTx(c.Request.Context())
	defer uow.Rollback()
	w := &models.Warehouse{Code: in.Code, Name: in.Name, City: in.City, Address: in.Address, IsDefault: in.IsDefault, Priority: in.Priority}
	if err := uow.WarehouseRepository().Create(c.Request.Context(), w); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "خطا در ایجاد"})
		return
	}
	if in.IsDefault {
		_ = uow.WarehouseRepository().SetDefault(c.Request.Context(), w.ID)
	}
	if err := uow.Commit(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "خطا در ثبت"})
		return
	}
	c.JSON(http.StatusCreated, w)
}

// Update بروزرسانی انبار
func (h *WarehouseHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var in struct {
		Name      *string `json:"name"`
		City      *string `json:"city"`
		Address   *string `json:"address"`
		IsActive  *bool   `json:"is_active"`
		IsDefault *bool   `json:"is_default"`
		Priority  *int    `json:"priority"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر"})
		return
	}
	uow := h.uowf.Create()
	_ = uow.BeginTx(c.Request.Context())
	defer uow.Rollback()
	w, err := uow.WarehouseRepository().GetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "انبار یافت نشد"})
		return
	}
	if in.Name != nil {
		w.Name = *in.Name
	}
	if in.City != nil {
		w.City = *in.City
	}
	if in.Address != nil {
		w.Address = *in.Address
	}
	if in.IsActive != nil {
		w.IsActive = *in.IsActive
	}
	if in.Priority != nil {
		w.Priority = *in.Priority
	}
	if err := uow.WarehouseRepository().Update(c.Request.Context(), w); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "خطا در بروزرسانی"})
		return
	}
	if in.IsDefault != nil && *in.IsDefault {
		_ = uow.WarehouseRepository().SetDefault(c.Request.Context(), w.ID)
	}
	if err := uow.Commit(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "خطا در ثبت"})
		return
	}
	c.JSON(http.StatusOK, w)
}

// Delete حذف انبار
func (h *WarehouseHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uow := h.uowf.Create()
	_ = uow.BeginTx(c.Request.Context())
	defer uow.Rollback()
	if err := uow.WarehouseRepository().Delete(c.Request.Context(), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "خطا در حذف"})
		return
	}
	if err := uow.Commit(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "خطا در ثبت"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
