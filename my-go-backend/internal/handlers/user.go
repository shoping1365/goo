package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"my-go-backend/internal/middleware"
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// tableExists بررسی وجود جدول در پایگاه داده (برای fallback های dev)
func tableExists(db *gorm.DB, name string) bool {
	var cnt int64
	if err := db.Raw("SELECT count(*) FROM information_schema.tables WHERE table_name = ?", name).Scan(&cnt).Error; err != nil {
		return false
	}
	return cnt > 0
}

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// RegisterUser handles user registration
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// Check if mobile number already exists
	var count int64
	h.DB.Model(&models.User{}).Where("mobile = ?", req.Mobile).Count(&count)
	if count > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_PHONE", utils.GetErrorMessage("DUPLICATE_PHONE"), nil))
		return
	}

	now := time.Now()
	username := req.Username
	if username == "" {
		username = req.Mobile
	}
	user := models.User{
		Username:     username,
		Name:         req.FullName,
		Mobile:       req.Mobile,
		Email:        fmt.Sprintf("%s@temp.local", req.Mobile), // temporary email
		RoleID:       1,                                        // default customer role
		RegisteredAt: now,
		Status:       "active",
		UpdatedAt:    now,
	}
	if err := h.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUserAddress handles updating user's address information
func (h *UserHandler) UpdateUserAddress(c *gin.Context) {
	userID := c.Param("id")
	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	updates := map[string]interface{}{
		"address":    req.Address,
		"city":       req.City,
		"landline":   req.Landline,
		"updated_at": time.Now(),
	}
	if err := h.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User address updated successfully"})
}

// GetUsers returns all users
func (h *UserHandler) GetUsers(c *gin.Context) {
	// خواندن آستانه از تنظیمات (پیش‌فرض 30 ثانیه)
	onlineThresholdSec := 30
	// تلاش برای خواندن از جدول settings
	var settingVal string
	_ = h.DB.Table("settings").Select("value").Where("key = ?", "realtime.users_online_threshold_sec").Scan(&settingVal).Error
	if settingVal != "" {
		if v, err := strconv.Atoi(settingVal); err == nil && v > 0 {
			onlineThresholdSec = v
		}
	}
	var users []models.User
	// جدول کاربران ستون created_at ندارد؛ از registered_at برای ترتیب استفاده می‌کنیم
	if err := h.DB.Order("registered_at DESC").Find(&users).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	// استخراج نقش‌ها برای نمایش انسانی (با priority)
	type roleInfo struct {
		Name     string
		Priority int
	}
	roleByID := map[uint]roleInfo{}
	{
		// جمع‌آوری role_id های منحصربه‌فرد
		ids := make([]uint, 0, len(users))
		seen := map[uint]struct{}{}
		for _, u := range users {
			if u.RoleID > 0 {
				if _, ok := seen[u.RoleID]; !ok {
					seen[u.RoleID] = struct{}{}
					ids = append(ids, u.RoleID)
				}
			}
		}
		if len(ids) > 0 {
			var rows []struct {
				ID          uint
				Name        string
				DisplayName string
				Priority    int
			}
			_ = h.DB.Table("roles").Select("id, name, display_name, priority").Where("id IN ?", ids).Scan(&rows).Error
			for _, r := range rows {
				nm := r.Name
				if r.DisplayName != "" {
					nm = r.DisplayName
				}
				roleByID[r.ID] = roleInfo{Name: nm, Priority: r.Priority}
			}
		}
	}

	// نقش‌های اضافی (user_roles)
	extraRolesByUser := map[uint][]string{}
	if len(users) > 0 {
		userIDs := make([]uint, 0, len(users))
		for _, u := range users {
			userIDs = append(userIDs, u.ID)
		}
		var extra []struct {
			UserID   uint   `gorm:"column:user_id"`
			RoleName string `gorm:"column:role_name"`
		}
		_ = h.DB.Table("user_roles ur").
			Select("ur.user_id, COALESCE(r.display_name, r.name) AS role_name").
			Joins("JOIN roles r ON r.id = ur.role_id").
			Where("ur.user_id IN ?", userIDs).
			Scan(&extra).Error
		for _, e := range extra {
			extraRolesByUser[e.UserID] = append(extraRolesByUser[e.UserID], e.RoleName)
		}
	}

	// استخراج آخرین فعالیت هر کاربر (user_sessions یا fallback: sessions برای dev)
	type activityRow struct {
		UserID       uint       `gorm:"column:user_id"`
		LastActivity *time.Time `gorm:"column:last_activity"`
	}
	var activities []activityRow
	if tableExists(h.DB, "user_sessions") {
		_ = h.DB.Table("user_sessions").
			Select("user_id, MAX(last_activity) AS last_activity").
			Where("is_active = ?", true).
			Group("user_id").
			Scan(&activities).Error
	} else {
		_ = h.DB.Table("sessions").
			Select("user_id, MAX(last_used_at) AS last_activity").
			Where("is_active = ?", true).
			Group("user_id").
			Scan(&activities).Error
	}

	lastSeenByUser := make(map[uint]*time.Time, len(activities))
	for _, a := range activities {
		lastSeenByUser[a.UserID] = a.LastActivity
	}

	// ساخت خروجی سازگار با فرانت و شامل فیلد lastSeen
	out := make([]gin.H, 0, len(users))
	for _, u := range users {
		var lastSeen *time.Time
		if v, ok := lastSeenByUser[u.ID]; ok {
			lastSeen = v
		} else if u.LastLoginAt != nil {
			// فالو بک: اگر رکوردی در user_sessions/sessions نبود، از آخرین زمان ورود کاربر استفاده کن
			lastSeen = u.LastLoginAt
		}
		// کاربر آنلاین اگر کمتر از onlineThresholdSec از آخرین فعالیت گذشته باشد
		isOnline := false
		if lastSeen != nil && time.Since(*lastSeen) < time.Duration(onlineThresholdSec)*time.Second {
			isOnline = true
		}

		// نام نقش اصلی و لیست تمام نقش‌ها
		primaryRoleName := ""
		if info, ok := roleByID[u.RoleID]; ok {
			primaryRoleName = info.Name
		}
		roleNames := make([]string, 0)
		if primaryRoleName != "" {
			roleNames = append(roleNames, primaryRoleName)
		}
		if extras := extraRolesByUser[u.ID]; len(extras) > 0 {
			roleNames = append(roleNames, extras...)
		}

		// نقش موثر با بالاترین priority بین نقش اصلی و نقش‌های اضافه
		effectiveRoleName := primaryRoleName
		effectivePriority := -1 << 30
		if info, ok := roleByID[u.RoleID]; ok {
			effectiveRoleName = info.Name
			effectivePriority = info.Priority
		}
		// بررسی نقش‌های اضافه کاربر
		var extraRows []struct {
			RoleID      uint
			Name        string
			DisplayName string
			Priority    int
		}
		_ = h.DB.Table("user_roles ur").
			Select("ur.role_id, COALESCE(r.display_name, r.name) AS name, r.priority").
			Joins("JOIN roles r ON r.id = ur.role_id").
			Where("ur.user_id = ?", u.ID).
			Scan(&extraRows).Error
		for _, ex := range extraRows {
			if ex.Priority > effectivePriority {
				effectivePriority = ex.Priority
				effectiveRoleName = ex.Name
			}
		}

		out = append(out, gin.H{
			"id":                  u.ID,
			"username":            u.Username,
			"name":                u.Name,
			"email":               u.Email,
			"mobile":              u.Mobile,
			"role_id":             u.RoleID,
			"role_name":           primaryRoleName,
			"effective_role_name": effectiveRoleName,
			"roles":               roleNames,
			"status":              u.Status,
			"registered_at":       u.RegisteredAt,
			// سازگاری با فرانت: هر دو فرمت را برمی‌گردانیم
			"last_seen":  lastSeen,
			"lastSeenAt": lastSeen,
			"online":     isOnline,
		})
	}

	c.JSON(http.StatusOK, out)
}

// AdminRegisterUser handles admin registration of a new user
func (h *UserHandler) AdminRegisterUser(c *gin.Context) {
	type AdminRegisterRequest struct {
		FullName string `json:"fullName" binding:"required"`
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Mobile   string `json:"mobile" binding:"required"`
		Role     string `json:"role" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var req AdminRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	var count int64
	h.DB.Model(&models.User{}).Where("username = ? OR email = ? OR mobile = ?", req.Username, req.Email, req.Mobile).Count(&count)
	if count > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("DUPLICATE_EMAIL", utils.GetErrorMessage("DUPLICATE_EMAIL"), nil))
		return
	}
	// هش کردن رمز عبور
	hashedPwd, err := utils.HashPassword(req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("INTERNAL_ERROR", utils.GetErrorMessage("INTERNAL_ERROR"), err.Error()))
		return
	}

	if req.Role == "admin" || req.Role == "superadmin" || req.Role == "مدیر" || req.Role == "ادمین" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_ROLE", "ثبت نقش مدیر یا ادمین مجاز نیست.", nil))
		return
	}

	now := time.Now()
	user := models.User{
		Name:         req.FullName,
		Username:     req.Username,
		Email:        req.Email,
		Mobile:       req.Mobile,
		RoleID:       1, // default customer role
		Status:       "active",
		PasswordHash: hashedPwd,
		RegisteredAt: now,
		UpdatedAt:    now,
	}
	if err := h.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UpdateLastSeen(c *gin.Context) {
	// آپدیت session activity به جای last_seen در user
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", utils.GetErrorMessage("UNAUTHORIZED"), nil))
		return
	}
	now := time.Now()

	// آپدیت last_activity (یا fallback last_used_at) برای سیشن‌های فعال
	var err error
	var rows int64 = 0
	if tableExists(h.DB, "user_sessions") {
		tx := h.DB.Model(&models.UserSession{}).
			Where("user_id = ? AND is_active = ?", userID, true).
			Update("last_activity", &now)
		err = tx.Error
		rows = tx.RowsAffected
	} else {
		tx := h.DB.Table("sessions").
			Where("user_id = ? AND is_active = ?", userID, true).
			Update("last_used_at", &now)
		err = tx.Error
		rows = tx.RowsAffected
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	// اگر سطری آپدیت نشد، به‌عنوان فالو‌بک، فیلد last_login_at در جدول users را به‌روزرسانی کن
	if rows == 0 {
		_ = h.DB.Model(&models.User{}).Where("id = ?", userID).Update("last_login_at", now).Error
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// DeleteUser handles user deletion
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if err := h.DB.Delete(&models.User{}, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// AdminSetUserRole تغییر نقش کاربر توسط ادمین
func (h *UserHandler) AdminSetUserRole(c *gin.Context) {
	userID := c.Param("id")
	var req struct {
		RoleID uint `json:"role_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	if err := h.DB.Model(&models.User{}).Where("id = ?", userID).Update("role_id", req.RoleID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// AdminBlockUser مسدودسازی کاربر
func (h *UserHandler) AdminBlockUser(c *gin.Context) {
	userID := c.Param("id")
	var req struct {
		Reason string `json:"reason"`
	}
	_ = c.ShouldBindJSON(&req)
	updates := map[string]interface{}{
		"status":       "blocked",
		"blocked_at":   time.Now(),
		"block_reason": req.Reason,
	}
	if err := h.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// AdminUnblockUser رفع مسدودی کاربر
func (h *UserHandler) AdminUnblockUser(c *gin.Context) {
	userID := c.Param("id")
	updates := map[string]interface{}{
		"status":       "active",
		"blocked_at":   nil,
		"block_reason": nil,
	}
	if err := h.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// ------------------------------
// User Detail Read endpoints
// ------------------------------

// GetUserByID دریافت اطلاعات جزئی یک کاربر (برای صفحهٔ جزئیات)
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}

	// آخرین فعالیت: user_sessions یا fallback: sessions
	var lastSeen sql.NullTime
	if tableExists(h.DB, "user_sessions") {
		_ = h.DB.Table("user_sessions").
			Select("MAX(last_activity) AS last_activity").
			Where("user_id = ? AND is_active = ?", user.ID, true).
			Scan(&lastSeen).Error
	} else {
		_ = h.DB.Table("sessions").
			Select("MAX(last_used_at) AS last_activity").
			Where("user_id = ? AND is_active = ?", user.ID, true).
			Scan(&lastSeen).Error
	}

	isOnline := false
	if lastSeen.Valid && time.Since(lastSeen.Time) < 120*time.Second {
		isOnline = true
	}

	// استخراج فیلدهای اضافی از پروفایل برای نمایش
	var profile map[string]interface{}
	if len(user.ProfileData) > 0 {
		_ = json.Unmarshal(user.ProfileData, &profile)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":            user.ID,
		"username":      user.Username,
		"name":          user.Name,
		"email":         user.Email,
		"mobile":        user.Mobile,
		"role_id":       user.RoleID,
		"status":        user.Status,
		"registered_at": user.RegisteredAt,
		"last_seen": func() *time.Time {
			if lastSeen.Valid {
				return &lastSeen.Time
			}
			return nil
		}(),
		"online":        isOnline,
		"landline":      profile["landline"],
		"national_code": profile["national_code"],
	})
}

// UpdateUserBasicInfo به‌روزرسانی اطلاعات پایه کاربر
func (h *UserHandler) UpdateUserBasicInfo(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Name        *string                `json:"name"`
		Email       *string                `json:"email"`
		Mobile      *string                `json:"mobile"`
		ProfileData map[string]interface{} `json:"profile_data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}

	updates := map[string]interface{}{}
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Email != nil {
		// اگر ایمیل خالی باشد، آن را null قرار بده
		if *req.Email == "" {
			updates["email"] = nil
		} else {
			// فقط اگر ایمیل تغییر کرده باشد، بررسی تکراری بودن
			if user.Email != *req.Email {
				var existingUser models.User
				if err := h.DB.Where("email = ? AND id != ?", *req.Email, id).First(&existingUser).Error; err == nil {
					c.AbortWithStatusJSON(http.StatusConflict, utils.New("EMAIL_EXISTS", "ایمیل قبلاً توسط کاربر دیگری استفاده شده است", nil))
					return
				}
			}
			updates["email"] = *req.Email
		}
	}
	if req.Mobile != nil {
		updates["mobile"] = *req.Mobile
	}

	// Merge profile data JSONB
	if req.ProfileData != nil {
		existing := map[string]interface{}{}
		if len(user.ProfileData) > 0 {
			_ = json.Unmarshal(user.ProfileData, &existing)
		}
		for k, v := range req.ProfileData {
			existing[k] = v
		}
		b, _ := json.Marshal(existing)
		updates["profile_data"] = datatypes.JSON(b)
	}

	if len(updates) == 0 {
		c.JSON(http.StatusOK, gin.H{"updated": 0})
		return
	}

	if err := h.DB.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// AdminForceLogoutUser خروج اجباری کاربر (ابطال سشن‌ها و توکن‌ها)
func (h *UserHandler) AdminForceLogoutUser(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id64 == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), "invalid id"))
		return
	}
	id := uint(id64)
	// غیرفعال کردن سشن‌های فعال
	_ = h.DB.Model(&models.UserSession{}).Where("user_id = ? AND is_active = ?", id, true).Update("is_active", false).Error
	// ابطال توکن‌های رفرش
	authService := services.NewAuthService(h.DB)
	_ = authService.RevokeAllUserTokens(uint(id))
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// AdminResetUserPassword ریست رمز عبور کاربر
func (h *UserHandler) AdminResetUserPassword(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		NewPassword string `json:"new_password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	hashed, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("INTERNAL_ERROR", utils.GetErrorMessage("INTERNAL_ERROR"), err.Error()))
		return
	}
	if err := h.DB.Model(&models.User{}).Where("id = ?", id).Update("password_hash", hashed).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// AdminSendUserWarning ارسال پیام/هشدار سیستمی به کاربر
func (h *UserHandler) AdminSendUserWarning(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id64 == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), "invalid id"))
		return
	}
	id := uint(id64)
	var req struct {
		Title string `json:"title"`
		Body  string `json:"body" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	record := map[string]interface{}{
		"user_id":    id,
		"channel":    "system",
		"source":     "admin",
		"title":      req.Title,
		"body":       req.Body,
		"status":     "sent",
		"sent_at":    time.Now(),
		"created_at": time.Now(),
	}
	if err := h.DB.Table("user_notifications").Create(&record).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// GetUserAddresses لیست آدرس‌های کاربر
func (h *UserHandler) GetUserAddresses(c *gin.Context) {
	userID := c.Param("id")
	var addresses []models.UserAddress
	if err := h.DB.Where("user_id = ?", userID).Order("is_default DESC, updated_at DESC").Find(&addresses).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, addresses)
}

// GetUserBankAccounts لیست حساب‌های بانکی کاربر
func (h *UserHandler) GetUserBankAccounts(c *gin.Context) {
	userID := c.Param("id")
	var accounts []models.UserBankAccount
	if err := h.DB.Where("user_id = ?", userID).Order("is_default DESC, verified_at DESC NULLS LAST, updated_at DESC").Find(&accounts).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, accounts)
}

// GetUserWalletSummary خلاصه کیف پول کاربر
func (h *UserHandler) GetUserWalletSummary(c *gin.Context) {
	userID := c.Param("id")
	var wallet models.UserWallet
	if err := h.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{"exists": false, "balance": 0})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exists":              true,
		"balance":             wallet.Balance,
		"currency":            wallet.Currency,
		"status":              wallet.Status,
		"last_transaction_at": wallet.LastTransactionAt,
	})
}

// GetUserWalletTransactions تراکنش‌های کیف پول کاربر
func (h *UserHandler) GetUserWalletTransactions(c *gin.Context) {
	userID := c.Param("id")
	limit := utils.ParseInt(c.DefaultQuery("limit", "20"))
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	// پیدا کردن wallet_id
	var wallet models.UserWallet
	if err := h.DB.Select("id").Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, []any{})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	var txs []models.WalletTransaction
	if err := h.DB.Where("wallet_id = ?", wallet.ID).Order("created_at DESC").Limit(limit).Find(&txs).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, txs)
}

// GetUserNotifications لیست اعلان/پیام‌های کاربر از user_notifications
func (h *UserHandler) GetUserNotifications(c *gin.Context) {
	userID := c.Param("id")
	limit := utils.ParseInt(c.DefaultQuery("limit", "5"))
	if limit <= 0 || limit > 100 {
		limit = 5
	}

	var items []struct {
		ID        uint       `json:"id"`
		Channel   string     `json:"channel"`
		Source    string     `json:"source"`
		Title     string     `json:"title"`
		Body      string     `json:"body"`
		Status    string     `json:"status"`
		SentAt    *time.Time `json:"sent_at"`
		CreatedAt time.Time  `json:"created_at"`
	}
	err := h.DB.Table("user_notifications").
		Where("user_id = ?", userID).
		Order("COALESCE(sent_at, created_at) DESC").
		Limit(limit).
		Scan(&items).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, items)
}

// -----------------------------------------------
// Login (یکپارچه: username/mobile + password)
// -----------------------------------------------
// @route POST /api/auth/login
// Body: { "identifier": "username or mobile", "password": "string" }
// • اگر identifier فقط شامل اعداد باشد → شماره موبایل در نظر گرفته می‌شود
// • اگر identifier شامل حروف باشد → نام کاربری در نظر گرفته می‌شود
// • همه کاربران (مشتری و ادمین) می‌توانند با پسورد وارد شوند (اگر پسورد تنظیم کرده باشند)
// • اگر کاربر پسورد نداشته باشد، از OTP استفاده کند
// -----------------------------------------------
func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Identifier string `json:"identifier" binding:"required"`
		Password   string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	// تشخیص نوع identifier (موبایل یا نام کاربری)
	isPhone := true
	for _, ch := range req.Identifier {
		if ch < '0' || ch > '9' {
			isPhone = false
			break
		}
	}

	var user models.User
	var err error

	if isPhone {
		// جستجو با شماره موبایل
		err = h.DB.Where("mobile = ?", req.Identifier).First(&user).Error
	} else {
		// جستجو با نام کاربری
		err = h.DB.Where("username = ?", req.Identifier).First(&user).Error
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("LOGIN_FAILED", utils.GetErrorMessage("LOGIN_FAILED"), nil))
		return
	}

	// بررسی وجود پسورد: همه کاربران می‌توانند با پسورد وارد شوند (اگر پسورد داشته باشند)
	if user.PasswordHash == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "برای این کاربر پسورد تنظیم نشده است. لطفاً از کد تایید استفاده کنید", nil))
		return
	}

	// If password not supplied → ask for password phase
	if req.Password == "" {
		c.JSON(http.StatusAccepted, gin.H{"requirePassword": true})
		return
	}

	// validate password
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("LOGIN_FAILED", utils.GetErrorMessage("LOGIN_FAILED"), nil))
		return
	}

	// تولید توکن JWT برای احراز هویت
	authService := services.NewAuthService(h.DB)
	settings, _ := authService.GetAuthSettings()
	token, refreshToken, err := authService.GenerateTokens(&user, settings)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("TOKEN_ERROR", utils.GetErrorMessage("TOKEN_ERROR"), err.Error()))
		return
	}

	if err := authService.CreateSessionWithDeviceInfo(&user, token, refreshToken, settings, c, "password"); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("SESSION_ERROR", "خطا در ایجاد نشست", err.Error()))
		return
	}

	// تنظیم کوکی‌های HttpOnly
	h.setAuthCookies(c, token, refreshToken, settings)

	// ----------------- Merge Guest Cart -----------------
	// اگر کوکی session_id موجود است، احتمالاً کاربر به عنوان مهمان سبدی دارد.
	cookieSession, _ := c.Cookie("session_id")
	if cookieSession != "" {
		var guestCart models.Cart
		if err := h.DB.Preload("Items").Where("session_id = ? AND user_id IS NULL", cookieSession).First(&guestCart).Error; err == nil {
			// بررسی سبد فعال کاربر
			var userCart models.Cart
			userCartErr := h.DB.Preload("Items").Where("user_id = ?", user.ID).First(&userCart).Error

			if userCartErr == nil {
				// ادغام آیتم‌ها
				for _, gi := range guestCart.Items {
					var existing models.CartItem
					errFind := h.DB.Where("cart_id = ? AND product_id = ?", userCart.ID, gi.ProductID).First(&existing).Error
					if errFind == nil {
						// جمع کردن تعداد
						h.DB.Model(&existing).Update("quantity", existing.Quantity+gi.Quantity)
						// حذف آیتم مهمان
						h.DB.Delete(&gi)
					} else {
						// انتقال آیتم مهمان به سبد کاربر
						h.DB.Model(&gi).Update("cart_id", userCart.ID)
					}
				}
				// حذف سبد مهمان
				h.DB.Delete(&guestCart)
			} else {
				// کاربر هنوز سبدی ندارد → سبد مهمان را به او اختصاص می‌دهیم
				h.DB.Model(&guestCart).Update("user_id", user.ID)
			}
		}
	}
	// -----------------------------------------------------

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GetCurrentUser returns the current authenticated user
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	// Get user_id from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", utils.GetErrorMessage("UNAUTHORIZED"), nil))
		return
	}

	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", utils.GetErrorMessage("UNAUTHORIZED"), nil))
		return
	}

	// دریافت پرمیشن‌های کاربر
	permissions, err := middleware.GetUserPermissions(user.ID)
	fmt.Printf("User ID: %d, RoleID: %d, Permissions: %v, Error: %v\n", user.ID, user.RoleID, permissions, err)
	if err != nil {
		// در صورت خطا، آرایه خالی برمی‌گردانیم
		permissions = make([]string, 0)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          user.ID,
		"username":    user.Username,
		"name":        user.Name,
		"email":       user.Email,
		"role_id":     user.RoleID,
		"status":      user.Status,
		"permissions": permissions,
	})
}

// Logout handles user logout
func (h *UserHandler) Logout(c *gin.Context) {
	authService := services.NewAuthService(h.DB)

	if sessionToken, hasSession := c.Get("session_token"); hasSession {
		if tokenStr, ok := sessionToken.(string); ok && tokenStr != "" {
			if err := authService.DeactivateSessionByToken(tokenStr); err != nil {
				fmt.Printf("failed to deactivate session: %v\n", err)
			}
		}
	} else if userID, exists := c.Get("user_id"); exists {
		if uid, ok := userID.(uint); ok {
			authService.RevokeAllUserTokens(uid)
		}
	}

	// پاک کردن کوکی‌ها
	h.clearAuthCookies(c)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// setAuthCookies تنظیم کوکی‌های احراز هویت
func (h *UserHandler) setAuthCookies(c *gin.Context, accessToken, refreshToken string, settings *models.AuthSettings) {
	// تنظیم SameSite mode برای امنیت و جلوگیری از CSRF
	c.SetSameSite(http.SameSiteLaxMode)

	// کوکی Access Token
	c.SetCookie(
		"access_token",
		accessToken,
		int((time.Duration(settings.JwtExpiryHours) * time.Hour).Seconds()),
		"/",
		"",
		false, // Secure = false برای localhost
		true,  // HttpOnly = true
	)

	// کوکی Refresh Token
	c.SetCookie(
		"refresh_token",
		refreshToken,
		int((time.Duration(settings.RefreshTokenExpiryDays) * 24 * time.Hour).Seconds()),
		"/",
		"",
		false, // Secure = false برای localhost
		true,  // HttpOnly = true
	)
}

// clearAuthCookies پاک کردن کوکی‌های احراز هویت
func (h *UserHandler) clearAuthCookies(c *gin.Context) {
	// تنظیم SameSite برای consistency
	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
}
