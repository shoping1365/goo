package handlers

import (
	"fmt"
	"net/http"

	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// BankAccountAdminHandler هندلر ادمین کارت/حساب بانکی
type BankAccountAdminHandler struct {
	svc *services.BankAccountService
}

func NewBankAccountAdminHandler(svc *services.BankAccountService) *BankAccountAdminHandler {
	return &BankAccountAdminHandler{svc: svc}
}

// GET /api/admin/bank-accounts
func (h *BankAccountAdminHandler) List(c *gin.Context) {
	page := parseIntDefault(c.DefaultQuery("page", "1"), 1)
	pageSize := parseIntDefault(c.DefaultQuery("pageSize", "20"), 20)
	status := c.Query("status")
	bank := c.Query("bank")
	items, total, err := h.svc.List(c, page, pageSize, status, bank)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items, "total": total, "page": page, "pageSize": pageSize})
}

// POST /api/admin/bank-accounts/:id/verify
func (h *BankAccountAdminHandler) Verify(c *gin.Context) {
	id := parseUintParam(c, "id")
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VALIDATION_ERROR", "message": "invalid id"})
		return
	}
	if err := h.svc.Verify(c, id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VERIFY_FAILED", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// POST /api/admin/bank-accounts/:id/block
func (h *BankAccountAdminHandler) Block(c *gin.Context) {
	id := parseUintParam(c, "id")
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VALIDATION_ERROR", "message": "invalid id"})
		return
	}
	if err := h.svc.Block(c, id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "BLOCK_FAILED", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// POST /api/admin/bank-accounts/:id/unblock
func (h *BankAccountAdminHandler) Unblock(c *gin.Context) {
	id := parseUintParam(c, "id")
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VALIDATION_ERROR", "message": "invalid id"})
		return
	}
	if err := h.svc.Unblock(c, id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "UNBLOCK_FAILED", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// POST /api/admin/bank-accounts/:id/reject
func (h *BankAccountAdminHandler) Reject(c *gin.Context) {
	id := parseUintParam(c, "id")
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VALIDATION_ERROR", "message": "invalid id"})
		return
	}
	if err := h.svc.Reject(c, id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "REJECT_FAILED", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func parseUintParam(c *gin.Context, name string) uint {
	var id uint
	_, _ = fmt.Sscan(c.Param(name), &id)
	return id
}
