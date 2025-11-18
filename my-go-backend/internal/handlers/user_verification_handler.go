package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// UserVerificationHandler - هندلر احراز هویت کاربران
type UserVerificationHandler struct {
	verificationService *services.UserVerificationService
}

// NewUserVerificationHandler - سازنده هندلر
func NewUserVerificationHandler(verificationService *services.UserVerificationService) *UserVerificationHandler {
	return &UserVerificationHandler{
		verificationService: verificationService,
	}
}

// CreateVerificationRequest - ایجاد درخواست احراز هویت جدید
// POST /api/verifications
func (h *UserVerificationHandler) CreateVerificationRequest(c *gin.Context) {
	var req models.CreateUserVerificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر است", err.Error()))
		return
	}

	// ایجاد درخواست
	verification, err := h.verificationService.CreateVerificationRequest(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("CREATE_ERROR", "خطا در ایجاد درخواست احراز هویت", err.Error()))
		return
	}

	// تبدیل به Response
	response := h.verificationService.ConvertToResponse(verification)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "درخواست احراز هویت با موفقیت ایجاد شد",
		"data":    response,
	})
}

// GetVerificationByID - دریافت اطلاعات یک درخواست احراز هویت
// GET /api/verifications/:id
func (h *UserVerificationHandler) GetVerificationByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه نامعتبر است", nil))
		return
	}

	verification, err := h.verificationService.GetVerificationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.New("NOT_FOUND", "درخواست یافت نشد", nil))
		return
	}

	response := h.verificationService.ConvertToResponse(verification)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// GetUserVerifications - دریافت لیست درخواست‌های یک کاربر
// GET /api/users/:id/verifications
func (h *UserVerificationHandler) GetUserVerifications(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه کاربر نامعتبر است", nil))
		return
	}

	verifications, err := h.verificationService.GetUserVerifications(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("FETCH_ERROR", "خطا در دریافت لیست احراز هویت", err.Error()))
		return
	}

	// تبدیل به Response
	responses := make([]models.UserVerificationResponse, 0)
	for _, v := range verifications {
		response := h.verificationService.ConvertToResponse(&v)
		responses = append(responses, *response)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    responses,
	})
}

// GetPendingVerifications - دریافت لیست درخواست‌های در انتظار تایید
// GET /api/admin/verifications/pending
func (h *UserVerificationHandler) GetPendingVerifications(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	verifications, total, err := h.verificationService.GetPendingVerifications(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("FETCH_ERROR", "خطا در دریافت لیست", err.Error()))
		return
	}

	// تبدیل به Response
	responses := make([]models.UserVerificationResponse, 0)
	for _, v := range verifications {
		response := h.verificationService.ConvertToResponse(&v)
		responses = append(responses, *response)
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"data":      responses,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetAllVerifications - دریافت تمام درخواست‌ها
// GET /api/admin/verifications
func (h *UserVerificationHandler) GetAllVerifications(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	status := c.Query("status")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	var verifications []models.UserVerification
	var total int64
	var err error

	if status != "" {
		// جستجو بر اساس وضعیت
		params := map[string]interface{}{
			"status": status,
		}
		verifications, total, err = h.verificationService.SearchVerifications(params, page, pageSize)
	} else {
		// دریافت همه
		verifications, total, err = h.verificationService.GetAllVerifications(page, pageSize)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("FETCH_ERROR", "خطا در دریافت لیست", err.Error()))
		return
	}

	// تبدیل به Response
	responses := make([]models.UserVerificationResponse, 0)
	for _, v := range verifications {
		response := h.verificationService.ConvertToResponse(&v)
		responses = append(responses, *response)
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"data":      responses,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// VerifyRequest - تایید درخواست احراز هویت (فاز 1 - دستی)
// POST /api/admin/verifications/:id/verify
func (h *UserVerificationHandler) VerifyRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه نامعتبر است", nil))
		return
	}

	// دریافت ID ادمین از context (middleware Auth)
	adminID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", "کاربر احراز هویت نشده است", nil))
		return
	}

	// تایید درخواست
	err = h.verificationService.VerifyRequest(uint(id), adminID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("VERIFY_ERROR", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "درخواست با موفقیت تایید شد",
	})
}

// RejectRequest - رد درخواست احراز هویت (فاز 1 - دستی)
// POST /api/admin/verifications/:id/reject
func (h *UserVerificationHandler) RejectRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه نامعتبر است", nil))
		return
	}

	// دریافت دلیل رد
	var req struct {
		Reason string `json:"reason" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "دلیل رد الزامی است", err.Error()))
		return
	}

	// دریافت ID ادمین از context
	adminID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", "کاربر احراز هویت نشده است", nil))
		return
	}

	// رد درخواست
	err = h.verificationService.RejectRequest(uint(id), adminID.(uint), req.Reason)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("REJECT_ERROR", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "درخواست رد شد",
	})
}

// GetVerifiedFields - دریافت فیلدهای تایید شده یک کاربر
// GET /api/users/:id/verified-fields
func (h *UserVerificationHandler) GetVerifiedFields(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه کاربر نامعتبر است", nil))
		return
	}

	verifiedFields, err := h.verificationService.GetVerifiedFields(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("FETCH_ERROR", "خطا در دریافت اطلاعات", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    verifiedFields,
	})
}

// GetPendingCount - دریافت تعداد درخواست‌های در انتظار
// GET /api/admin/verifications/pending/count
func (h *UserVerificationHandler) GetPendingCount(c *gin.Context) {
	count, err := h.verificationService.GetPendingCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("FETCH_ERROR", "خطا در دریافت تعداد", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"count":   count,
	})
}

// DeleteVerification - حذف درخواست احراز هویت
// DELETE /api/admin/verifications/:id
func (h *UserVerificationHandler) DeleteVerification(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه نامعتبر است", nil))
		return
	}

	err = h.verificationService.DeleteVerification(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DELETE_ERROR", "خطا در حذف درخواست", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "درخواست با موفقیت حذف شد",
	})
}
