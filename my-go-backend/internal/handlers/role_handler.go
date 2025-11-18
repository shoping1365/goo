package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
)

// RoleHandler هندلر مدیریت نقش‌ها
type RoleHandler struct {
	DB *gorm.DB
}

// NewRoleHandler ایجاد نمونه جدید از RoleHandler
func NewRoleHandler(db *gorm.DB) *RoleHandler {
	return &RoleHandler{DB: db}
}

// ListRoles لیست تمام نقش‌ها را برمی‌گرداند
func (h *RoleHandler) ListRoles(c *gin.Context) {
	var roles []models.Role

	query := h.DB.Preload("Permissions").Order("priority ASC, name ASC")

	// فیلتر بر اساس وضعیت فعال/غیرفعال
	if isActive := c.Query("is_active"); isActive != "" {
		if isActive == "true" {
			query = query.Where("is_active = ?", true)
		} else if isActive == "false" {
			query = query.Where("is_active = ?", false)
		}
	}

	// جستجو بر اساس نام
	if search := c.Query("search"); search != "" {
		query = query.Where("name LIKE ? OR display_name LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لیست نقش‌ها", err.Error()))
		return
	}

	// اضافه کردن تعداد کاربران و دسترسی‌ها برای هر نقش
	for i := range roles {
		var userCount int64
		// شمارش کاربران بر اساس role_id
		h.DB.Model(&models.User{}).Where("role_id = ?", roles[i].ID).Count(&userCount)
		roles[i].Users = nil // پاک کردن آرایه کاربران برای جلوگیری از حجم زیاد

		// اضافه کردن فیلدهای محاسبه شده
		roles[i].UsersCount = int(userCount)
		roles[i].PermissionsCount = len(roles[i].Permissions)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    roles,
	})
}

// GetRole نقش خاصی را بر اساس ID برمی‌گرداند
func (h *RoleHandler) GetRole(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه نقش نامعتبر است", nil))
		return
	}

	var role models.Role
	if err := h.DB.First(&role, roleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.New("ROLE_NOT_FOUND", "نقش مورد نظر یافت نشد", nil))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت نقش", err.Error()))
		return
	}

	// فقط مجوزهای فعال را دستی کوئری کن و به role.Permissions اختصاص بده
	var permissions []models.Permission
	if err := h.DB.Table("permissions").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id = ? AND role_permissions.is_active = ?", roleID, true).
		Find(&permissions).Error; err == nil {
		role.Permissions = permissions
	}

	// در GetRole اگر نیاز به شمارش کاربران بود، مشابه بالا اصلاح کن
	var userCount int64
	if err := h.DB.Model(&models.User{}).Where("role_id = ?", roleID).Count(&userCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در شمارش کاربران", err.Error()))
		return
	}
	role.UsersCount = int(userCount)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    role,
	})
}

// validateRoleName نام نقش را برای انگلیسی بودن بررسی می‌کند
func (h *RoleHandler) validateRoleName(name string) error {
	// بررسی اینکه نام فقط شامل حروف انگلیسی، اعداد و underscore باشد
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '_') {
			return fmt.Errorf("نام نقش باید فقط شامل حروف انگلیسی، اعداد و underscore باشد")
		}
	}
	return nil
}

// CreateRole نقش جدید ایجاد می‌کند
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req models.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر است", err.Error()))
		return
	}

	// بررسی نام انگلیسی
	if err := h.validateRoleName(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ROLE_NAME", err.Error(), nil))
		return
	}

	// بررسی تکراری نبودن نام نقش
	var existingRole models.Role
	if err := h.DB.Where("name = ?", req.Name).First(&existingRole).Error; err == nil {
		c.JSON(http.StatusBadRequest, utils.New("DUPLICATE_ROLE", "نقش با این نام قبلاً وجود دارد", nil))
		return
	}

	// ایجاد نقش جدید
	role := models.Role{
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Description: req.Description,
		IsActive:    true,
		IsSystem:    false,
		Priority:    0,
	}

	if err := h.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد نقش", err.Error()))
		return
	}

	// تخصیص دسترسی‌ها به نقش
	if len(req.Permissions) > 0 {
		if err := h.assignPermissionsToRole(role.ID, req.Permissions, c.GetUint("user_id")); err != nil {
			c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در تخصیص دسترسی‌ها", err.Error()))
			return
		}
	}

	// بارگذاری مجدد نقش با دسترسی‌ها
	h.DB.Preload("Permissions").First(&role, role.ID)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    role,
		"message": "نقش با موفقیت ایجاد شد",
	})
}

// UpdateRole نقش موجود را به‌روزرسانی می‌کند
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه نقش نامعتبر است", nil))
		return
	}

	var req models.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر است", err.Error()))
		return
	}

	var role models.Role
	if err := h.DB.First(&role, roleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.New("ROLE_NOT_FOUND", "نقش مورد نظر یافت نشد", nil))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت نقش", err.Error()))
		return
	}

	// حذف شرط عدم ویرایش نقش سیستمی
	// if role.IsSystem {
	// 	c.JSON(http.StatusBadRequest, utils.New("SYSTEM_ROLE_UPDATE_ERROR", "نقش‌های سیستمی قابل ویرایش نیستند", gin.H{
	// 		"role_name":    role.Name,
	// 		"display_name": role.DisplayName,
	// 		"reason":       "این نقش بخشی از سیستم است و برای حفظ امنیت و عملکرد صحیح برنامه قابل تغییر نمی‌باشد",
	// 	}))
	// 	return
	// }

	// اگر نام تغییر کرده، validation اعمال کن
	if req.Name != "" && req.Name != role.Name {
		// بررسی نام انگلیسی
		if err := h.validateRoleName(req.Name); err != nil {
			c.JSON(http.StatusBadRequest, utils.New("INVALID_ROLE_NAME", err.Error(), nil))
			return
		}

		// بررسی تکراری نبودن نام جدید
		var existingRole models.Role
		if err := h.DB.Where("name = ? AND id != ?", req.Name, roleID).First(&existingRole).Error; err == nil {
			c.JSON(http.StatusBadRequest, utils.New("DUPLICATE_ROLE", "نقش با این نام قبلاً وجود دارد", nil))
			return
		}

		role.Name = req.Name
	}

	// به‌روزرسانی فیلدها
	if req.DisplayName != "" {
		role.DisplayName = req.DisplayName
	}
	if req.Description != "" {
		role.Description = req.Description
	}
	if req.IsActive != nil {
		role.IsActive = *req.IsActive
	}

	if err := h.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی نقش", err.Error()))
		return
	}

	// به‌روزرسانی دسترسی‌ها
	if req.Permissions != nil {
		// حذف تمام دسترسی‌های موجود
		if err := h.DB.Where("role_id = ?", role.ID).Delete(&models.RolePermission{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف دسترسی‌های قبلی", err.Error()))
			return
		}

		// تخصیص دسترسی‌های جدید
		if len(req.Permissions) > 0 {
			if err := h.assignPermissionsToRole(role.ID, req.Permissions, c.GetUint("user_id")); err != nil {
				c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در تخصیص دسترسی‌های جدید", err.Error()))
				return
			}
		}
	}

	// بارگذاری مجدد نقش با دسترسی‌ها
	h.DB.Preload("Permissions").First(&role, role.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    role,
		"message": "نقش با موفقیت به‌روزرسانی شد",
	})
}

// DeleteRole نقش را حذف می‌کند
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه نقش نامعتبر است", nil))
		return
	}

	var role models.Role
	if err := h.DB.First(&role, roleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.New("ROLE_NOT_FOUND", "نقش مورد نظر یافت نشد", nil))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت نقش", err.Error()))
		return
	}

	// بررسی اینکه آیا نقش admin یا developer است یا خیر (نقش‌های سیستمی)
	if role.Name == "admin" || role.Name == "developer" {
		c.JSON(http.StatusBadRequest, utils.New("SYSTEM_ROLE_DELETE_ERROR", "نقش‌های سیستمی قابل حذف نیستند", gin.H{
			"role_name":    role.Name,
			"display_name": role.DisplayName,
			"reason":       "این نقش بخشی از سیستم است و برای عملکرد صحیح برنامه ضروری می‌باشد",
		}))
		return
	}

	// بررسی اینکه آیا کاربرانی با این نقش وجود دارند یا خیر
	var userCount int64
	if err := h.DB.Model(&models.User{}).Where("role_id = ?", roleID).Count(&userCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بررسی کاربران", err.Error()))
		return
	}

	// اگر کاربرانی با این نقش وجود دارند، آن‌ها را بدون نقش کن
	if userCount > 0 {
		// به‌روزرسانی تمام کاربران این نقش به بدون نقش (role_id = null)
		if err := h.DB.Model(&models.User{}).Where("role_id = ?", roleID).Update("role_id", nil).Error; err != nil {
			c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف نقش از کاربران", err.Error()))
			return
		}
	}

	// حذف دسترسی‌های نقش
	if err := h.DB.Where("role_id = ?", roleID).Delete(&models.RolePermission{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف دسترسی‌های نقش", err.Error()))
		return
	}

	// حذف نقش
	if err := h.DB.Delete(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف نقش", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "نقش با موفقیت حذف شد",
		"data": gin.H{
			"users_without_role": userCount,
		},
	})
}

// assignPermissionsToRole دسترسی‌ها را به نقش تخصیص می‌دهد
func (h *RoleHandler) assignPermissionsToRole(roleID uint, permissionIDs []uint, grantedBy uint) error {
	for _, permissionID := range permissionIDs {
		rolePermission := models.RolePermission{
			RoleID:       roleID,
			PermissionID: permissionID,
			GrantedBy:    grantedBy,
			IsActive:     true,
		}

		if err := h.DB.Create(&rolePermission).Error; err != nil {
			return err
		}
	}
	return nil
}

// GetRolePermissions دسترسی‌های یک نقش را برمی‌گرداند
func (h *RoleHandler) GetRolePermissions(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه نقش نامعتبر است", nil))
		return
	}

	var rolePermissions []models.RolePermission
	if err := h.DB.Where("role_id = ? AND is_active = ?", roleID, true).
		Preload("Permission").
		Find(&rolePermissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت دسترسی‌های نقش", err.Error()))
		return
	}

	var permissions []models.Permission
	for _, rp := range rolePermissions {
		permissions = append(permissions, rp.Permission)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    permissions,
	})
}

// GetRoleStatistics آمار کلی نقش‌ها را برمی‌گرداند
func (h *RoleHandler) GetRoleStatistics(c *gin.Context) {
	var stats struct {
		TotalRoles       int64 `json:"total_roles"`
		TotalUsers       int64 `json:"total_users"`
		TotalPermissions int64 `json:"total_permissions"`
		ActiveRoles      int64 `json:"active_roles"`
	}

	// تعداد کل نقش‌ها
	if err := h.DB.Model(&models.Role{}).Count(&stats.TotalRoles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت آمار نقش‌ها", err.Error()))
		return
	}

	// تعداد کل کاربران
	if err := h.DB.Model(&models.User{}).Count(&stats.TotalUsers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت آمار کاربران", err.Error()))
		return
	}

	// تعداد کل دسترسی‌ها
	if err := h.DB.Model(&models.Permission{}).Count(&stats.TotalPermissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت آمار دسترسی‌ها", err.Error()))
		return
	}

	// تعداد نقش‌های فعال (نقش‌هایی که کاربر دارند)
	if err := h.DB.Model(&models.Role{}).Joins("JOIN users ON users.role_id = roles.id").Distinct("roles.id").Count(&stats.ActiveRoles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت آمار نقش‌های فعال", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    stats,
	})
}

// UpdateRolePermissions مجوزهای یک نقش را به‌روزرسانی می‌کند
func (h *RoleHandler) UpdateRolePermissions(c *gin.Context) {
	id := c.Param("id")
	roleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه نقش نامعتبر است", nil))
		return
	}

	// بررسی وجود نقش
	var role models.Role
	if err := h.DB.First(&role, roleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.New("ROLE_NOT_FOUND", "نقش مورد نظر یافت نشد", nil))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت نقش", err.Error()))
		return
	}

	// دریافت لیست نام‌های مجوزها از درخواست
	var req struct {
		Permissions []string `json:"permissions" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر است", err.Error()))
		return
	}

	// تبدیل نام‌های مجوزها به ID
	var permissionIDs []uint
	if len(req.Permissions) > 0 {
		var permissions []models.Permission
		if err := h.DB.Where("name IN ?", req.Permissions).Find(&permissions).Error; err != nil {
			c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت مجوزها", err.Error()))
			return
		}

		fmt.Printf("🔍 Permission check: sent=%d, found=%d\n", len(req.Permissions), len(permissions))
		if len(permissions) != len(req.Permissions) {
			fmt.Printf("⚠️  Some permissions not found – creating missing ones...\n")

			foundMap := make(map[string]bool)
			for _, p := range permissions {
				foundMap[p.Name] = true
			}
			var missingNames []string
			for _, name := range req.Permissions {
				if !foundMap[name] {
					missingNames = append(missingNames, name)
				}
			}

			// ایجاد مجوزهای مفقود با گروه "custom"
			for _, name := range missingNames {
				perm := models.Permission{
					Name:        name,
					DisplayName: name,
					Module:      "custom",
					SubModule:   "custom",
					Action:      "custom",
					Resource:    "custom",
					IsActive:    true,
				}
				if err := h.DB.Create(&perm).Error; err != nil {
					c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد مجوزهای جدید", err.Error()))
					return
				}
				permissions = append(permissions, perm)
			}
		}

		fmt.Printf("✅ Permission list ready. Total: %d\n", len(permissions))
		fmt.Printf("✅ All permissions found successfully\n")

		for _, permission := range permissions {
			permissionIDs = append(permissionIDs, permission.ID)
		}
	}

	fmt.Printf("🔄 Starting transaction for role ID: %d\n", roleID)
	// شروع تراکنش
	tx := h.DB.Begin()

	// دریافت مجوزهای فعلی نقش
	var currentRolePermissions []models.RolePermission
	if err := tx.Where("role_id = ? AND is_active = ?", roleID, true).Find(&currentRolePermissions).Error; err != nil {
		fmt.Printf("❌ Error getting current role permissions: %v\n", err)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت مجوزهای فعلی", err.Error()))
		return
	}
	fmt.Printf("📊 Current role permissions count: %d\n", len(currentRolePermissions))

	// ایجاد map از مجوزهای فعلی برای مقایسه سریع
	currentPermissionMap := make(map[uint]bool)
	for _, rp := range currentRolePermissions {
		currentPermissionMap[rp.PermissionID] = true
	}

	// ایجاد map از مجوزهای جدید
	newPermissionMap := make(map[uint]bool)
	for _, permissionID := range permissionIDs {
		newPermissionMap[permissionID] = true
	}

	// پیدا کردن مجوزهایی که باید اضافه شوند (موجود در جدید، موجود در فعلی نیست)
	var permissionsToAdd []uint
	for _, permissionID := range permissionIDs {
		if !currentPermissionMap[permissionID] {
			permissionsToAdd = append(permissionsToAdd, permissionID)
		}
	}

	// پیدا کردن مجوزهایی که باید حذف شوند (موجود در فعلی، موجود در جدید نیست)
	var permissionsToRemove []uint
	for _, rp := range currentRolePermissions {
		if !newPermissionMap[rp.PermissionID] {
			permissionsToRemove = append(permissionsToRemove, rp.PermissionID)
		}
	}

	// حذف مجوزهایی که باید حذف شوند
	if len(permissionsToRemove) > 0 {
		if err := tx.Where("role_id = ? AND permission_id IN ?", roleID, permissionsToRemove).Delete(&models.RolePermission{}).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف مجوزهای قدیمی", err.Error()))
			return
		}
	}

	// اضافه کردن مجوزهای جدید
	var rolePermissions []models.RolePermission
	for _, permissionID := range permissionsToAdd {
		// دریافت user_id از context
		userIDInterface, exists := c.Get("user_id")
		var grantedBy uint
		if exists && userIDInterface != nil {
			if userIDUint, ok := userIDInterface.(uint); ok {
				grantedBy = userIDUint
			} else if userIDInt, ok := userIDInterface.(int); ok {
				grantedBy = uint(userIDInt)
			}
		}

		rolePermission := models.RolePermission{
			RoleID:       uint(roleID),
			PermissionID: permissionID,
			GrantedBy:    grantedBy,
			IsActive:     true,
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	if len(rolePermissions) > 0 {
		if err := tx.Create(&rolePermissions).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در تخصیص مجوزهای جدید", err.Error()))
			return
		}
	}

	// تایید تراکنش
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در تایید تراکنش", err.Error()))
		return
	}

	// بارگذاری مجدد نقش با مجوزهای جدید
	h.DB.Preload("Permissions").First(&role, roleID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    role,
		"message": "مجوزهای نقش با موفقیت به‌روزرسانی شد",
	})
}
