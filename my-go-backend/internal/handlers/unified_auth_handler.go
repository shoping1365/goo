package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"my-go-backend/internal/middleware"
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
Handler احراز هویت یکپارچه
-----------------------------------------------------------------
این handler گسترش سیستم JWT موجود است و ویژگی‌های زیر را اضافه می‌کند:
- ورود/ثبت‌نام یکپارچه با موبایل
- انتخاب روش احراز (OTP یا پسورد)
- استفاده از همان کوکی‌های HttpOnly موجود
- سازگاری کامل با سیستم JWT فعلی
*/

type UnifiedAuthHandler struct {
	db                 *gorm.DB
	unifiedAuthService *services.UnifiedAuthService
	authService        *services.AuthService // سرویس JWT موجود
}

// NewUnifiedAuthHandler ایجاد نمونه جدید از handler احراز هویت یکپارچه
func NewUnifiedAuthHandler(db *gorm.DB, unifiedAuthService *services.UnifiedAuthService, authService *services.AuthService) *UnifiedAuthHandler {
	return &UnifiedAuthHandler{
		db:                 db,
		unifiedAuthService: unifiedAuthService,
		authService:        authService,
	}
}

// CheckUser بررسی وجود کاربر
// @Summary بررسی وجود کاربر
// @Description بررسی وجود کاربر با شماره موبایل و وضعیت رمز عبور
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Param request body models.CheckUserRequest true "شماره موبایل"
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/check-user [post]
func (h *UnifiedAuthHandler) CheckUser(c *gin.Context) {
	var req struct {
		Mobile string `json:"mobile" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// بررسی وجود کاربر
	var user models.User
	userExists := h.db.Where("mobile = ? OR username = ?", req.Mobile, req.Mobile).First(&user).Error == nil

	response := gin.H{
		"exists":         userExists,
		"hasPassword":    false,
		"isActive":       true,
		"role_id":        1,     // default customer
		"canUsePassword": false, // آیا می‌تونه با پسورد وارد بشه
		"registered":     userExists,
	}

	if userExists {
		response["hasPassword"] = user.PasswordHash != ""
		response["isActive"] = user.Status == "active"
		response["role_id"] = user.RoleID
		// همه کاربران می‌توانند با پسورد وارد شوند (اگر پسورد تنظیم کرده باشند)
		response["canUsePassword"] = user.PasswordHash != ""
	}

	c.JSON(http.StatusOK, response)
}

// SendOTP ارسال کد تایید
// @Summary ارسال کد تایید
// @Description ارسال کد تایید به شماره موبایل
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Param request body models.SendOTPRequest true "شماره موبایل"
// @Success 200 {object} models.AuthResponse
// @Router /api/auth/send-otp [post]
func (h *UnifiedAuthHandler) SendOTP(c *gin.Context) {
	var req models.SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// دریافت اطلاعات درخواست
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	response, err := h.unifiedAuthService.SendOtp(req.Mobile, ip, userAgent)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("SMS_ERROR", "خطا در ارسال کد تایید", err.Error()))
		return
	}

	if !response.Success {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("AUTH_ERROR", response.Message, nil))
		return
	}

	c.JSON(http.StatusOK, response)
}

// CheckMobileAndSendOTP بررسی موبایل و ارسال کد تایید
// @Summary بررسی موبایل و ارسال کد تایید
// @Description بررسی وجود کاربر با موبایل و ارسال کد تایید. اگر کاربر پسورد داشته باشد، امکان انتخاب روش احراز را می‌دهد
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Param request body models.SendOTPRequest true "شماره موبایل"
// @Success 200 {object} models.AuthResponse
// @Router /api/auth/check-mobile [post]
func (h *UnifiedAuthHandler) CheckMobileAndSendOTP(c *gin.Context) {
	var req models.SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// دریافت اطلاعات درخواست
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	response, err := h.unifiedAuthService.CheckMobileAndSendOTP(req.Mobile, ip, userAgent)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("SMS_ERROR", "خطا در ارسال کد تایید", err.Error()))
		return
	}

	if !response.Success {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("AUTH_ERROR", response.Message, nil))
		return
	}

	c.JSON(http.StatusOK, response)
}

// VerifyOTPAndLogin تایید کد OTP و ورود
// @Summary تایید کد OTP و ورود
// @Description تایید کد OTP و ورود به سیستم. در صورت عدم وجود کاربر، ثبت‌نام خودکار انجام می‌شود
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Param request body models.VerifyOTPRequest true "شماره موبایل و کد تایید"
// @Success 200 {object} models.AuthResponse
// @Router /api/auth/verify-otp [post]
func (h *UnifiedAuthHandler) VerifyOTPAndLogin(c *gin.Context) {
	var req models.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// دریافت اطلاعات درخواست
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	response, err := h.unifiedAuthService.VerifyOTPAndLogin(req.Mobile, req.Code, ip, userAgent)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("AUTH_ERROR", "خطا در احراز هویت", err.Error()))
		return
	}

	if !response.Success {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("AUTH_ERROR", response.Message, nil))
		return
	}

	// تنظیم کوکی‌های JWT (مانند سیستم موجود)
	if response.Token != "" && response.RefreshToken != "" {
		if !h.createSessionAndCookies(c, response, "otp") {
			return
		}
	}

	// حذف توکن‌ها از پاسخ برای امنیت (چون در کوکی ذخیره می‌شوند)
	response.Token = ""
	response.RefreshToken = ""

	c.JSON(http.StatusOK, response)
}

// LoginWithPassword ورود با پسورد
// @Summary ورود با پسورد
// @Description ورود به سیستم با شماره موبایل و پسورد
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Param request body models.LoginPasswordRequest true "شماره موبایل و پسورد"
// @Success 200 {object} models.AuthResponse
// @Router /api/auth/login-password [post]
func (h *UnifiedAuthHandler) LoginWithPassword(c *gin.Context) {
	var req models.LoginPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	if req.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "پسورد الزامی است", nil))
		return
	}

	// دریافت اطلاعات درخواست
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	response, err := h.unifiedAuthService.LoginWithPassword(req.Mobile, req.Password, ip, userAgent)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("AUTH_ERROR", "خطا در احراز هویت", err.Error()))
		return
	}

	if !response.Success {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("LOGIN_FAILED", response.Message, nil))
		return
	}

	// تنظیم کوکی‌های JWT (مانند سیستم موجود)
	if response.Token != "" && response.RefreshToken != "" {
		if !h.createSessionAndCookies(c, response, "password") {
			return
		}
	}

	// حذف توکن‌ها از پاسخ برای امنیت (چون در کوکی ذخیره می‌شوند)
	response.Token = ""
	response.RefreshToken = ""

	c.JSON(http.StatusOK, response)
}

// UnifiedLogin ورود یکپارچه (endpoint اختیاری برای سازگاری)
// @Summary ورود یکپارچه
// @Description ورود با انتخاب روش احراز (OTP یا پسورد)
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Param request body models.UnifiedLoginRequest true "اطلاعات ورود"
// @Success 200 {object} models.AuthResponse
// @Router /api/auth/unified-login [post]
func (h *UnifiedAuthHandler) UnifiedLogin(c *gin.Context) {
	var req models.UnifiedLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// دریافت اطلاعات درخواست
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	var response *models.AuthResponse
	var err error

	switch req.AuthMethod {
	case "otp":
		if req.Code == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "کد تایید الزامی است", nil))
			return
		}
		response, err = h.unifiedAuthService.VerifyOTPAndLogin(req.Mobile, req.Code, ip, userAgent)

	case "password":
		if req.Password == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "پسورد الزامی است", nil))
			return
		}
		response, err = h.unifiedAuthService.LoginWithPassword(req.Mobile, req.Password, ip, userAgent)

	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "روش احراز هویت نامعتبر", nil))
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("AUTH_ERROR", "خطا در احراز هویت", err.Error()))
		return
	}

	if !response.Success {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("LOGIN_FAILED", response.Message, nil))
		return
	}

	// تنظیم کوکی‌های JWT (مانند سیستم موجود)
	if response.Token != "" && response.RefreshToken != "" {
		if !h.createSessionAndCookies(c, response, req.AuthMethod) {
			return
		}
	}

	// حذف توکن‌ها از پاسخ برای امنیت (چون در کوکی ذخیره می‌شوند)
	response.Token = ""
	response.RefreshToken = ""

	c.JSON(http.StatusOK, response)
}

// SetPassword تنظیم پسورد توسط کاربر (نیازمند احراز هویت)
// @Summary تنظیم پسورد
// @Description تنظیم پسورد توسط کاربر احراز هویت شده
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Param request body models.SetPasswordRequest true "پسورد جدید"
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/set-password [post]
// @Security ApiKeyAuth
func (h *UnifiedAuthHandler) SetPassword(c *gin.Context) {
	// دریافت ID کاربر از context (از middleware احراز هویت)
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("AUTH_ERROR", "کاربر احراز هویت نشده", nil))
		return
	}

	var req models.SetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// بررسی تطابق پسوردها
	if req.Password != req.ConfirmPassword {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "پسوردها مطابقت ندارند", nil))
		return
	}

	// هش کردن پسورد
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("INTERNAL_ERROR", "خطا در رمزنگاری پسورد", err.Error()))
		return
	}

	// به‌روزرسانی پسورد کاربر
	if err := h.db.Model(&models.User{}).Where("id = ?", userID).Update("password_hash", hashedPassword).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی پسورد", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "پسورد با موفقیت تنظیم شد",
	})
}

// GetAuthMethods دریافت روش‌های احراز هویت موجود برای یک موبایل
// @Summary دریافت روش‌های احراز هویت
// @Description دریافت روش‌های احراز هویت موجود برای شماره موبایل
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Param mobile query string true "شماره موبایل"
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/methods [get]
func (h *UnifiedAuthHandler) GetAuthMethods(c *gin.Context) {
	mobile := c.Query("mobile")
	if mobile == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شماره موبایل الزامی است", nil))
		return
	}

	// بررسی وجود کاربر
	var user models.User
	userExists := h.db.Where("mobile = ?", mobile).First(&user).Error == nil

	response := gin.H{
		"mobile":      mobile,
		"user_exists": userExists,
		"methods": gin.H{
			"otp":      true, // همیشه فعال
			"password": false,
		},
	}

	if userExists && user.PasswordHash != "" {
		response["methods"].(gin.H)["password"] = true
	}

	c.JSON(http.StatusOK, response)
}

// setAuthCookies تنظیم کوکی‌های احراز هویت (مانند سیستم موجود)
func (h *UnifiedAuthHandler) createSessionAndCookies(c *gin.Context, response *models.AuthResponse, authMethod string) bool {
	if response.User == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("SESSION_ERROR", "اطلاعات کاربر در دسترس نیست", nil))
		return false
	}

	settings, err := h.authService.GetAuthSettings()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("AUTH_ERROR", "خطا در دریافت تنظیمات احراز هویت", err.Error()))
		return false
	}

	if err := h.authService.CreateSessionWithDeviceInfo(response.User, response.Token, response.RefreshToken, settings, c, authMethod); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("SESSION_ERROR", "خطا در ایجاد نشست", err.Error()))
		return false
	}

	h.setAuthCookies(c, response.Token, response.RefreshToken)
	return true
}

func (h *UnifiedAuthHandler) setAuthCookies(c *gin.Context, accessToken, refreshToken string) {
	isOriginalDomain := strings.EqualFold(c.Request.Header.Get("X-Forwarded-Host"), c.Request.Host)

	// دریافت تنظیمات احراز هویت
	settings, err := h.authService.GetAuthSettings()
	if err != nil {
		// استفاده از مقادیر پیش‌فرض در صورت خطا
		settings = &models.AuthSettings{
			JwtExpiryHours:         24,
			RefreshTokenExpiryDays: 7,
		}
	}

	host := c.Request.Host
	host = strings.ToLower(host)
	isProductionHost := strings.Contains(host, ".") && !strings.HasSuffix(host, ".local")
	env := strings.ToLower(os.Getenv("ENV"))
	isProductionEnv := env == "production" || env == "prod"
	isProduction := isProductionEnv || isProductionHost

	// تنظیم cookie domain - در localhost domain خالی باشد تا browser درست set کند
	cookieDomain := ""
	if isProductionHost && isOriginalDomain {
		// production: use domain prefix for subdomain sharing
		parts := strings.Split(host, ".")
		if len(parts) >= 2 {
			cookieDomain = "." + strings.Join(parts[len(parts)-2:], ".")
		}
	}

	// تنظیم SameSite mode برای امنیت و جلوگیری از CSRF
	// SameSiteLaxMode: اجازه می‌دهد cookies در navigation معمولی ارسال شوند
	c.SetSameSite(http.SameSiteLaxMode)

	// تنظیم کوکی Access Token
	c.SetCookie(
		"access_token",
		accessToken,
		int(time.Duration(settings.JwtExpiryHours)*time.Hour/time.Second),
		"/",
		cookieDomain,
		isProduction, // Secure = true در production
		true,         // HttpOnly = true
	)

	// تنظیم کوکی Refresh Token
	c.SetCookie(
		"refresh_token",
		refreshToken,
		int(time.Duration(settings.RefreshTokenExpiryDays)*24*time.Hour/time.Second),
		"/",
		cookieDomain,
		isProduction, // Secure = true در production
		true,         // HttpOnly = true
	)
}

// clearAuthCookies پاک کردن کوکی‌های احراز هویت (برای logout)
func (h *UnifiedAuthHandler) clearAuthCookies(c *gin.Context) {
	// تنظیم SameSite برای consistency
	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
}

func (h *UnifiedAuthHandler) GetAuthStatus(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"authenticated": false,
			"user":          nil,
		})
		return
	}

	// دریافت اطلاعات کاربر
	var user models.User
	if err := h.db.Preload("UserRole").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"authenticated": false,
			"user":          nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"user":          user,
		"role_id":       user.RoleID,
	})
}

// AdminGetLoginAttempts دریافت تلاش‌های ورود (ادمین)
// @Summary دریافت تلاش‌های ورود
// @Description دریافت فهرست تلاش‌های ورود (فقط ادمین)
// @Tags مدیریت احراز هویت
// @Produce json
// @Param mobile query string false "فیلتر بر اساس موبایل"
// @Param page query int false "شماره صفحه"
// @Param limit query int false "تعداد در صفحه"
// @Success 200 {object} map[string]interface{}
// @Router /api/admin/auth/login-attempts [get]
// @Security ApiKeyAuth
func (h *UnifiedAuthHandler) AdminGetLoginAttempts(c *gin.Context) {
	mobile := c.Query("mobile")
	page := utils.ParseInt(c.DefaultQuery("page", "1"))
	limit := utils.ParseInt(c.DefaultQuery("limit", "50"))
	offset := (page - 1) * limit

	query := h.db.Model(&models.LoginAttempt{})

	if mobile != "" {
		query = query.Where("mobile = ?", mobile)
	}

	var attempts []models.LoginAttempt
	var total int64

	query.Count(&total)
	query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&attempts)

	c.JSON(http.StatusOK, gin.H{
		"data":  attempts,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

// AdminUnlockAccount رفع قفل حساب (ادمین)
// @Summary رفع قفل حساب
// @Description رفع قفل حساب کاربری (فقط ادمین)
// @Tags مدیریت احراز هویت
// @Accept json
// @Produce json
// @Param mobile path string true "شماره موبایل"
// @Success 200 {object} map[string]interface{}
// @Router /api/admin/auth/unlock/{mobile} [post]
// @Security ApiKeyAuth
func (h *UnifiedAuthHandler) AdminUnlockAccount(c *gin.Context) {
	mobile := c.Param("mobile")
	if mobile == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شماره موبایل الزامی است", nil))
		return
	}

	// رفع قفل کاربر
	var user models.User
	if err := h.db.Where("mobile = ?", mobile).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "کاربر یافت نشد",
		})
		return
	}

	// رفع قفل موقت و دائمی
	user.LockedUntil = nil
	user.FailedLoginCount = 0
	user.Status = "active" // تغییر وضعیت به فعال

	result := h.db.Save(&user)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در رفع قفل", result.Error.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("قفل حساب %s رفع شد", mobile),
	})
}

// GetCurrentUser دریافت اطلاعات کاربر جاری
// @Summary دریافت اطلاعات کاربر جاری
// @Description دریافت اطلاعات کاربر احراز هویت شده
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/me [get]
func (h *UnifiedAuthHandler) GetCurrentUser(c *gin.Context) {
	// Get user_id from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", "کاربر احراز هویت نشده", nil))
		return
	}

	var user models.User
	if err := h.db.Preload("UserRole").First(&user, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", "کاربر یافت نشد", nil))
		return
	}

	// دریافت نام نقش
	roleName := ""
	if user.UserRole.ID > 0 {
		roleName = user.UserRole.Name
	}

	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"name":     user.Name,
			"email":    user.Email,
			"mobile":   user.Mobile,
			"role_id":  user.RoleID,
			"role":     roleName,
			"status":   user.Status,
		},
		"role_id":   user.RoleID,
		"role":      roleName,
		"role_name": roleName,
	})
}

// GetUserPermissions دریافت دسترسی‌های کاربر
// @Summary دریافت دسترسی‌های کاربر
// @Description دریافت تمام دسترسی‌های کاربر بر اساس نقش‌هایش
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/permissions [get]
func (h *UnifiedAuthHandler) GetUserPermissions(c *gin.Context) {
	// Get user_id from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", "کاربر احراز هویت نشده", nil))
		return
	}

	// دریافت دسترسی‌های کاربر از middleware
	permissions, err := middleware.GetUserPermissions(userID.(uint))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("INTERNAL_ERROR", "خطا در دریافت دسترسی‌ها", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"permissions": permissions,
	})
}

// Logout خروج از حساب کاربری
// @Summary خروج از حساب
// @Description خروج کاربر و باطل کردن توکن‌ها
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/logout [post]
func (h *UnifiedAuthHandler) Logout(c *gin.Context) {
	if sessionToken, hasSession := c.Get("session_token"); hasSession {
		if tokenStr, ok := sessionToken.(string); ok && tokenStr != "" {
			if err := h.authService.DeactivateSessionByToken(tokenStr); err != nil {
				fmt.Printf("failed to deactivate session: %v\n", err)
			}
		}
	} else if userID, exists := c.Get("user_id"); exists {
		if uid, ok := userID.(uint); ok {
			_ = h.authService.RevokeAllUserTokens(uid)
		}
	}

	// پاک کردن کوکی‌های احراز هویت
	h.clearAuthCookies(c)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "خروج با موفقیت انجام شد",
	})
}

// RefreshToken تجدید توکن
// @Summary تجدید توکن
// @Description تجدید توکن دسترسی با استفاده از refresh token
// @Tags احراز هویت یکپارچه
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/refresh [post]
func (h *UnifiedAuthHandler) RefreshToken(c *gin.Context) {
	// دریافت refresh token از کوکی
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("TOKEN_ERROR", "Refresh token یافت نشد", nil))
		return
	}

	// تایید refresh token
	refreshTokenModel, err := h.authService.ValidateRefreshToken(refreshToken)
	if err != nil {
		h.clearAuthCookies(c)
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("TOKEN_ERROR", "Refresh token نامعتبر", err.Error()))
		return
	}

	// دریافت اطلاعات کاربر
	var user models.User
	if err := h.db.First(&user, refreshTokenModel.UserID).Error; err != nil {
		h.clearAuthCookies(c)
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("TOKEN_ERROR", "کاربر یافت نشد", err.Error()))
		return
	}

	// تولید توکن‌های جدید
	settings, settingsErr := h.authService.GetAuthSettings()
	if settingsErr != nil {
		h.clearAuthCookies(c)
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("TOKEN_ERROR", "خطا در تجدید توکن", settingsErr.Error()))
		return
	}

	newAccessToken, newRefreshToken, err := h.authService.GenerateTokens(&user, settings)
	if err != nil {
		h.clearAuthCookies(c)
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("TOKEN_ERROR", "خطا در تجدید توکن", err.Error()))
		return
	}

	if err := h.authService.RotateSessionTokens(refreshTokenModel, newAccessToken, newRefreshToken, settings); err != nil {
		h.clearAuthCookies(c)
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("TOKEN_ERROR", "خطا در بروزرسانی نشست", err.Error()))
		return
	}

	// تنظیم کوکی‌های جدید
	h.setAuthCookies(c, newAccessToken, newRefreshToken)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "توکن با موفقیت تجدید شد",
	})
}

// GetCSRFToken دریافت CSRF Token برای درخواست‌های محافظت شده
// @Summary دریافت CSRF Token
// @Description دریافت CSRF Token برای محافظت در برابر حملات CSRF
// @Tags احراز هویت یکپارچه
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/csrf-token [get]
func (h *UnifiedAuthHandler) GetCSRFToken(c *gin.Context) {
	// تولید یک توکن CSRF یکتا
	csrfToken := h.generateCSRFToken()

	// تنظیم SameSite mode برای امنیت
	c.SetSameSite(http.SameSiteLaxMode)

	// ذخیره توکن در کوکی برای تایید در درخواست‌های بعدی
	c.SetCookie(
		"csrf_token",
		csrfToken,
		3600, // 1 ساعت
		"/",
		"",
		false, // Secure = false برای localhost
		true,  // HttpOnly = true
	)

	c.JSON(http.StatusOK, gin.H{
		"token": csrfToken,
	})
}

// generateCSRFToken تولید یک CSRF Token امن
func (h *UnifiedAuthHandler) generateCSRFToken() string {
	// استفاده از UUID برای تولید توکن یکتا
	return fmt.Sprintf("csrf_%d_%s", time.Now().Unix(), utils.GenerateRandomString(32))
}
