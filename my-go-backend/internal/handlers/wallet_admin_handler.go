package handlers

import (
	"net/http"
	"strconv"
	"time"

	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// WalletAdminHandler هندلر ادمین برای مدیریت کیف پول
// فقط orchestrate می‌کند و منطق دامنه در Service است.
type WalletAdminHandler struct {
	svc *services.WalletService
}

func NewWalletAdminHandler(svc *services.WalletService) *WalletAdminHandler {
	return &WalletAdminHandler{svc: svc}
}

// POST /api/admin/wallet/credit
func (h *WalletAdminHandler) Credit(c *gin.Context) {
	var req struct {
		UserID  uint    `json:"user_id" binding:"required"`
		Amount  float64 `json:"amount" binding:"required"`
		Method  string  `json:"method"`
		Gateway string  `json:"gateway"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VALIDATION_ERROR", "message": "ورودی نامعتبر", "error": err.Error()})
		return
	}
	if err := h.svc.CreditWallet(c, 0, req.UserID, req.Amount, req.Method, req.Gateway, nil); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "WALLET_CREDIT_FAILED", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// POST /api/admin/wallet/debit
func (h *WalletAdminHandler) Debit(c *gin.Context) {
	var req struct {
		UserID uint    `json:"user_id" binding:"required"`
		Amount float64 `json:"amount" binding:"required"`
		Method string  `json:"method"`
		Reason string  `json:"reason"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VALIDATION_ERROR", "message": "ورودی نامعتبر", "error": err.Error()})
		return
	}
	if err := h.svc.DebitWallet(c, 0, req.UserID, req.Amount, req.Method, req.Reason, nil); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "WALLET_DEBIT_FAILED", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// POST /api/admin/wallet/transfer
func (h *WalletAdminHandler) Transfer(c *gin.Context) {
	var req struct {
		FromUserID uint    `json:"from_user_id" binding:"required"`
		ToUserID   uint    `json:"to_user_id" binding:"required"`
		Amount     float64 `json:"amount" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VALIDATION_ERROR", "message": "ورودی نامعتبر", "error": err.Error()})
		return
	}
	if err := h.svc.Transfer(c, 0, req.FromUserID, req.ToUserID, req.Amount); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "WALLET_TRANSFER_FAILED", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// POST /api/admin/wallet/block/:userId
func (h *WalletAdminHandler) Block(c *gin.Context) {
	id64, _ := strconv.ParseUint(c.Param("userId"), 10, 64)
	if id64 == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VALIDATION_ERROR", "message": "شناسه نامعتبر"})
		return
	}
	if err := h.svc.BlockWallet(c, 0, uint(id64)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "WALLET_BLOCK_FAILED", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// POST /api/admin/wallet/unblock/:userId
func (h *WalletAdminHandler) Unblock(c *gin.Context) {
	id64, _ := strconv.ParseUint(c.Param("userId"), 10, 64)
	if id64 == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VALIDATION_ERROR", "message": "شناسه نامعتبر"})
		return
	}
	if err := h.svc.UnblockWallet(c, 0, uint(id64)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "WALLET_UNBLOCK_FAILED", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// GET /api/admin/wallet/transactions
func (h *WalletAdminHandler) AdminListTransactions(c *gin.Context) {
	page := parseIntDefault(c.DefaultQuery("page", "1"), 1)
	pageSize := parseIntDefault(c.DefaultQuery("pageSize", "20"), 20)
	filter := repository.AdminTransactionFilter{
		Type:     c.Query("type"),
		Status:   c.Query("status"),
		Query:    c.Query("q"),
		FromDate: c.Query("from"),
		ToDate:   c.Query("to"),
	}
	if m := c.Query("method"); m != "" {
		filter.Method = m
	}
	if u := c.Query("user_id"); u != "" {
		if id, err := strconv.Atoi(u); err == nil && id > 0 {
			v := uint(id)
			filter.UserID = &v
		}
	}
	rows, total, err := h.svc.ListTransactionsAdmin(c, page, pageSize, filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": rows, "total": total, "page": page, "pageSize": pageSize, "serverTime": time.Now()})
}

func parseIntDefault(s string, def int) int {
	if v, err := strconv.Atoi(s); err == nil && v > 0 {
		return v
	}
	return def
}

// GET /api/admin/wallet/summary
func (h *WalletAdminHandler) Summary(c *gin.Context) {
	res, err := h.svc.Summary(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GET /api/admin/wallet/trend
func (h *WalletAdminHandler) Trend(c *gin.Context) {
	days := parseIntDefault(c.DefaultQuery("days", "30"), 30)
	res, err := h.svc.TrendLastNDays(c, days)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "message": err.Error()})
		return
	}
	// سازگاری با فرانت در برخی بخش‌ها: کلیدها را lowercase می‌کنیم
	items := make([]gin.H, 0, len(res))
	for _, r := range res {
		items = append(items, gin.H{"day": r.Day, "net": r.Net})
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// POST /api/admin/wallet/transactions/:id/status
func (h *WalletAdminHandler) UpdateTxStatus(c *gin.Context) {
	id64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id64 == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VALIDATION_ERROR", "message": "invalid id"})
		return
	}
	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "VALIDATION_ERROR", "message": "invalid body"})
		return
	}
	if err := h.svc.UpdateTransactionStatus(c, uint(id64), req.Status); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "UPDATE_FAILED", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
