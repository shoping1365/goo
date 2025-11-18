package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RolePermissionHandler هندلر مدیریت دسترسی‌های نقش‌ها
// این هندلر برای مدیریت تخصیص دسترسی‌ها به نقش‌ها استفاده می‌شود
type RolePermissionHandler struct {
	db *gorm.DB
}

// NewRolePermissionHandler ایجاد نمونه جدید از RolePermissionHandler
func NewRolePermissionHandler(db *gorm.DB) *RolePermissionHandler {
	return &RolePermissionHandler{
		db: db,
	}
}

// GetRolePermissions دریافت دسترسی‌های یک نقش
// این متد تمام دسترسی‌های تخصیص داده شده به یک نقش را برمی‌گرداند
func (h *RolePermissionHandler) GetRolePermissions(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "شناسه نقش نامعتبر است", nil)
		return
	}

	var role models.Role
	if err := h.db.First(&role, roleID).Error; err != nil {
		utils.NotFound(c, "نقش مورد نظر یافت نشد", nil)
		return
	}

	// فقط مجوزهای فعال را دستی کوئری کن و به role.Permissions اختصاص بده
	var permissions []models.Permission
	if err := h.db.Table("permissions").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id = ? AND role_permissions.is_active = ?", roleID, true).
		Find(&permissions).Error; err == nil {
		role.Permissions = permissions
	}

	// گروه‌بندی دسترسی‌ها بر اساس ماژول و زیرماژول
	permissionGroups := h.groupPermissionsByModule(role.Permissions)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"role":              role,
			"permission_groups": permissionGroups,
		},
	})
}

// GetAllPermissions دریافت تمام دسترسی‌های موجود
// این متد تمام دسترسی‌های موجود در سیستم را به صورت گروه‌بندی شده برمی‌گرداند
func (h *RolePermissionHandler) GetAllPermissions(c *gin.Context) {
	var permissions []models.Permission
	if err := h.db.Where("is_active = ?", true).Order("module, sub_module, priority").Find(&permissions).Error; err != nil {
		utils.InternalServerError(c, "خطا در دریافت دسترسی‌ها", nil)
		return
	}

	// گروه‌بندی دسترسی‌ها
	permissionGroups := h.groupPermissionsByModule(permissions)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"permission_groups": permissionGroups,
		},
	})
}

// AssignPermission تخصیص یک دسترسی به نقش
// این متد یک دسترسی خاص را به یک نقش تخصیص می‌دهد
func (h *RolePermissionHandler) AssignPermission(c *gin.Context) {
	var req models.RolePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "داده‌های ورودی نامعتبر است", nil)
		return
	}

	// بررسی وجود نقش
	var role models.Role
	if err := h.db.First(&role, req.RoleID).Error; err != nil {
		utils.NotFound(c, "نقش مورد نظر یافت نشد", nil)
		return
	}

	// بررسی وجود دسترسی
	var permission models.Permission
	if err := h.db.First(&permission, req.PermissionID).Error; err != nil {
		utils.NotFound(c, "دسترسی مورد نظر یافت نشد", nil)
		return
	}

	// بررسی عدم وجود تخصیص قبلی
	var existingRolePermission models.RolePermission
	if err := h.db.Where("role_id = ? AND permission_id = ?", req.RoleID, req.PermissionID).First(&existingRolePermission).Error; err == nil {
		utils.BadRequest(c, "این دسترسی قبلاً به این نقش تخصیص داده شده است", nil)
		return
	}

	// ایجاد تخصیص جدید
	rolePermission := models.RolePermission{
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
		GrantedBy:    req.GrantedBy,
		ExpiresAt:    req.ExpiresAt,
		IsActive:     true,
	}

	if err := h.db.Create(&rolePermission).Error; err != nil {
		utils.InternalServerError(c, "خطا در تخصیص دسترسی", nil)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "دسترسی با موفقیت تخصیص داده شد",
		"data":    rolePermission,
	})
}

// BulkAssignPermissions تخصیص گروهی دسترسی‌ها به نقش
// این متد چندین دسترسی را همزمان به یک نقش تخصیص می‌دهد
// فقط تغییرات را اعمال می‌کند (اضافه کردن دسترسی‌های جدید و حذف دسترسی‌های حذف شده)
func (h *RolePermissionHandler) BulkAssignPermissions(c *gin.Context) {
	var req models.BulkRolePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "داده‌های ورودی نامعتبر است", nil)
		return
	}

	// بررسی وجود نقش
	var role models.Role
	if err := h.db.First(&role, req.RoleID).Error; err != nil {
		utils.NotFound(c, "نقش مورد نظر یافت نشد", nil)
		return
	}

	// بررسی وجود تمام دسترسی‌ها
	var permissions []models.Permission
	if err := h.db.Where("id IN ?", req.PermissionIDs).Find(&permissions).Error; err != nil {
		utils.NotFound(c, "برخی از دسترسی‌ها یافت نشدند", nil)
		return
	}

	if len(permissions) != len(req.PermissionIDs) {
		utils.NotFound(c, "برخی از دسترسی‌ها یافت نشدند", nil)
		return
	}

	// شروع تراکنش
	tx := h.db.Begin()

	// دریافت دسترسی‌های فعلی نقش
	var currentRolePermissions []models.RolePermission
	if err := tx.Where("role_id = ? AND is_active = ?", req.RoleID, true).Find(&currentRolePermissions).Error; err != nil {
		tx.Rollback()
		utils.InternalServerError(c, "خطا در دریافت دسترسی‌های فعلی", nil)
		return
	}

	// ایجاد map از دسترسی‌های فعلی برای مقایسه سریع
	currentPermissionMap := make(map[uint]bool)
	for _, rp := range currentRolePermissions {
		currentPermissionMap[rp.PermissionID] = true
	}

	// ایجاد map از دسترسی‌های جدید
	newPermissionMap := make(map[uint]bool)
	for _, permissionID := range req.PermissionIDs {
		newPermissionMap[permissionID] = true
	}

	// پیدا کردن دسترسی‌هایی که باید اضافه شوند (موجود در جدید، موجود در فعلی نیست)
	var permissionsToAdd []uint
	for _, permissionID := range req.PermissionIDs {
		if !currentPermissionMap[permissionID] {
			permissionsToAdd = append(permissionsToAdd, permissionID)
		}
	}

	// پیدا کردن دسترسی‌هایی که باید حذف شوند (موجود در فعلی، موجود در جدید نیست)
	var permissionsToRemove []uint
	for _, rp := range currentRolePermissions {
		if !newPermissionMap[rp.PermissionID] {
			permissionsToRemove = append(permissionsToRemove, rp.PermissionID)
		}
	}

	// حذف دسترسی‌هایی که باید حذف شوند
	if len(permissionsToRemove) > 0 {
		if err := tx.Where("role_id = ? AND permission_id IN ?", req.RoleID, permissionsToRemove).Delete(&models.RolePermission{}).Error; err != nil {
			tx.Rollback()
			utils.InternalServerError(c, "خطا در حذف دسترسی‌های قدیمی", nil)
			return
		}
	}

	// اضافه کردن دسترسی‌های جدید
	var rolePermissions []models.RolePermission
	for _, permissionID := range permissionsToAdd {
		rolePermission := models.RolePermission{
			RoleID:       req.RoleID,
			PermissionID: permissionID,
			GrantedBy:    req.GrantedBy,
			ExpiresAt:    req.ExpiresAt,
			IsActive:     true,
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	if len(rolePermissions) > 0 {
		if err := tx.Create(&rolePermissions).Error; err != nil {
			tx.Rollback()
			utils.InternalServerError(c, "خطا در تخصیص دسترسی‌های جدید", nil)
			return
		}
	}

	// تایید تراکنش
	if err := tx.Commit().Error; err != nil {
		utils.InternalServerError(c, "خطا در تایید تراکنش", nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "دسترسی‌ها با موفقیت به‌روزرسانی شدند",
		"data": gin.H{
			"added_count":   len(permissionsToAdd),
			"removed_count": len(permissionsToRemove),
			"total_count":   len(req.PermissionIDs),
		},
	})
}

// RevokePermission لغو دسترسی از نقش
// این متد یک دسترسی خاص را از یک نقش لغو می‌کند
func (h *RolePermissionHandler) RevokePermission(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("roleId"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "شناسه نقش نامعتبر است", nil)
		return
	}

	permissionID, err := strconv.ParseUint(c.Param("permissionId"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "شناسه دسترسی نامعتبر است", nil)
		return
	}

	// حذف تخصیص
	if err := h.db.Where("role_id = ? AND permission_id = ?", roleID, permissionID).Delete(&models.RolePermission{}).Error; err != nil {
		utils.InternalServerError(c, "خطا در لغو دسترسی", nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "دسترسی با موفقیت لغو شد",
	})
}

// CheckUserPermission بررسی دسترسی کاربر
// این متد بررسی می‌کند که آیا کاربر مورد نظر دسترسی خاصی دارد یا خیر
func (h *RolePermissionHandler) CheckUserPermission(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		utils.BadRequest(c, "شناسه کاربر نامعتبر است", nil)
		return
	}

	permissionName := c.Query("permission")
	if permissionName == "" {
		utils.BadRequest(c, "نام دسترسی الزامی است", nil)
		return
	}

	// بررسی دسترسی کاربر
	hasPermission, err := h.checkUserPermission(uint(userID), permissionName)
	if err != nil {
		utils.InternalServerError(c, "خطا در بررسی دسترسی", nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"has_permission": hasPermission,
		},
	})
}

// groupPermissionsByModule گروه‌بندی دسترسی‌ها بر اساس ماژول و زیرماژول
// این متد دسترسی‌ها را به صورت سلسله‌مراتبی گروه‌بندی می‌کند
func (h *RolePermissionHandler) groupPermissionsByModule(permissions []models.Permission) []models.PermissionGroup {
	moduleMap := make(map[string]*models.PermissionGroup)

	for _, permission := range permissions {
		// ایجاد یا دریافت گروه ماژول
		if moduleMap[permission.Module] == nil {
			moduleMap[permission.Module] = &models.PermissionGroup{
				Module:     permission.Module,
				SubModules: []models.SubModule{},
			}
		}

		// پیدا کردن یا ایجاد زیرماژول
		subModuleFound := false
		for i, subModule := range moduleMap[permission.Module].SubModules {
			if subModule.Name == permission.SubModule {
				moduleMap[permission.Module].SubModules[i].Permissions = append(subModule.Permissions, permission)
				subModuleFound = true
				break
			}
		}

		if !subModuleFound {
			// ایجاد زیرماژول جدید
			newSubModule := models.SubModule{
				Name:        permission.SubModule,
				DisplayName: permission.SubModule, // می‌توان از جدول جداگانه‌ای برای نام‌های نمایشی استفاده کرد
				Permissions: []models.Permission{permission},
			}
			moduleMap[permission.Module].SubModules = append(moduleMap[permission.Module].SubModules, newSubModule)
		}
	}

	// تبدیل map به slice
	var result []models.PermissionGroup
	for _, group := range moduleMap {
		result = append(result, *group)
	}

	return result
}

// checkUserPermission بررسی دسترسی کاربر به یک دسترسی خاص
// این متد بررسی می‌کند که آیا کاربر مورد نظر دسترسی خاصی دارد یا خیر
func (h *RolePermissionHandler) checkUserPermission(userID uint, permissionName string) (bool, error) {
	var count int64
	err := h.db.Table("role_permissions").
		Joins("JOIN users ON users.role_id = role_permissions.role_id").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Where("users.id = ? AND permissions.name = ? AND role_permissions.is_active = ?", userID, permissionName, true).
		Count(&count).Error

	return count > 0, err
}
