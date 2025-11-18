package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"my-go-backend/internal/database"
	"my-go-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Permission cache for performance optimization
type PermissionCache struct {
	cache sync.Map
	ttl   time.Duration
}

type CacheEntry struct {
	hasPermission bool
	expiry        time.Time
}

var permissionCache = &PermissionCache{
	ttl: 5 * time.Minute, // TTL 5 دقیقه
}

// getCachedPermission بررسی cache برای مجوزهای کاربر
func (pc *PermissionCache) getCachedPermission(userID uint, permission string) (bool, bool) {
	key := fmt.Sprintf("user:%d:perm:%s", userID, permission)

	if value, exists := pc.cache.Load(key); exists {
		entry := value.(CacheEntry)
		if time.Now().Before(entry.expiry) {
			return entry.hasPermission, true // cache hit
		} else {
			// Cache expired
			pc.cache.Delete(key)
		}
	}

	return false, false // cache miss
}

// setCachedPermission ذخیره مجوز در cache
func (pc *PermissionCache) setCachedPermission(userID uint, permission string, hasPermission bool) {
	key := fmt.Sprintf("user:%d:perm:%s", userID, permission)
	entry := CacheEntry{
		hasPermission: hasPermission,
		expiry:        time.Now().Add(pc.ttl),
	}
	pc.cache.Store(key, entry)
}

// invalidateUserPermissions حذف تمام مجوزهای کاربر از cache
func (pc *PermissionCache) invalidateUserPermissions(userID uint) {
	prefix := fmt.Sprintf("user:%d:perm:", userID)

	pc.cache.Range(func(key, value interface{}) bool {
		if keyStr, ok := key.(string); ok {
			if strings.HasPrefix(keyStr, prefix) {
				pc.cache.Delete(key)
			}
		}
		return true
	})
}

// startCacheCleanup تمیزکاری دوره‌ای cache
func (pc *PermissionCache) startCacheCleanup() {
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			now := time.Now()
			pc.cache.Range(func(key, value interface{}) bool {
				if entry, ok := value.(CacheEntry); ok {
					if now.After(entry.expiry) {
						pc.cache.Delete(key)
					}
				}
				return true
			})
		}
	}()
}

func init() {
	// شروع cache cleanup در background
	permissionCache.startCacheCleanup()
}

// InvalidateUserPermissions حذف مجوزهای کاربر از cache (برای استفاده در handlers)
func InvalidateUserPermissions(userID uint) {
	permissionCache.invalidateUserPermissions(userID)
	fmt.Printf("🗑️ Cache permissions برای کاربر %d پاک شد\n", userID)
}

// ClearAllPermissionsCache پاک کردن کل cache (برای maintenance)
func ClearAllPermissionsCache() {
	permissionCache.cache = sync.Map{}
	fmt.Printf("🗑️ تمام permission cache پاک شد\n")
}

// GetCacheStats آمار cache برای monitoring
func GetCacheStats() map[string]int {
	stats := map[string]int{
		"total_entries":   0,
		"expired_entries": 0,
	}

	now := time.Now()
	permissionCache.cache.Range(func(key, value interface{}) bool {
		stats["total_entries"]++
		if entry, ok := value.(CacheEntry); ok {
			if now.After(entry.expiry) {
				stats["expired_entries"]++
			}
		}
		return true
	})

	return stats
}

/*
میدلور بررسی دسترسی‌ها

این میدلور برای بررسی دسترسی‌های کاربر بر اساس نقش‌ها استفاده می‌شود.
*/

// AuthMiddleware میدلور احراز هویت را ایجاد می‌کند
// این middleware صرفاً احراز هویت کاربر را بررسی می‌کند، دسترسی‌ها جداگانه چک می‌شوند
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("🔍 بررسی احراز هویت برای مسیر: %s\n", c.Request.URL.Path)

		userID, exists := c.Get("user_id")
		fmt.Printf("👤 User ID از context: %v\n", userID)

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
			c.Abort()
			return
		}

		idUint := userID.(uint)

		// خواندن کاربر از دیتابیس
		var user models.User
		if err := database.GormDB.Preload("UserRole").First(&user, idUint).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر یافت نشد"})
			c.Abort()
			return
		}

		// بررسی نقش کاربر - RoleID حالا uint هست (نه pointer)
		if user.RoleID > 0 {
			// کاربر نقش جدید دارد (از جدول roles)
			var role models.Role
			if err := database.GormDB.First(&role, user.RoleID).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "نقش کاربر یافت نشد"})
				c.Abort()
				return
			}

			// ست کردن اطلاعات کاربر در context
			c.Set("user_id", userID)
			c.Set("role_id", user.RoleID) // استفاده از role_id بجای role name
			c.Set("user", user)
			c.Next()
			return
		}

		// ست کردن اطلاعات کاربر در context (همه کاربران - حتی customer ها)
		c.Set("user_id", userID)
		c.Set("role_id", user.RoleID)
		c.Set("user", user)

		fmt.Printf("✅ کاربر احراز شده - ID: %d, RoleID: %d\n", user.ID, user.RoleID)
		c.Next()
	}
}

// RequirePermission میدلور بررسی دسترسی خاص را ایجاد می‌کند
// این middleware برای مسیرهای ادمین استفاده می‌شود و customer ها را مسدود می‌کند
func RequirePermission(permissionName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
			c.Abort()
			return
		}

		// دریافت اطلاعات کامل کاربر
		var userObj models.User
		db := database.GormDB
		if err := db.Preload("UserRole").First(&userObj, userID.(uint)).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر یافت نشد"})
			c.Abort()
			return
		}

		fmt.Printf("🔒 RequirePermission '%s' برای کاربر %d (RoleID: %d)\n", permissionName, userObj.ID, userObj.RoleID)

		// شرط ویژه dev/developer - دسترسی کامل
		if userObj.Username == "dev" || userObj.RoleID == 100 || strings.EqualFold(userObj.UserRole.Name, "developer") || strings.EqualFold(userObj.UserRole.Name, "admin") {
			fmt.Printf("✅ دسترسی کامل برای developer (Username=dev یا RoleID=100)\n")
			c.Next()
			return
		}

		// بررسی نقش کاربر - customer ها (RoleID=1) به منابع ادمین دسترسی ندارند
		if userObj.RoleID == 1 {
			fmt.Printf("❌ Customer (RoleID=1) به منابع ادمین دسترسی ندارد\n")
			c.JSON(http.StatusNotFound, gin.H{"error": "صفحه مورد نظر پیدا نشد"})
			c.Abort()
			return
		}

		// بررسی دسترسی کاربران غیر customer با cache
		userIDUint := userID.(uint)

		// بررسی cache ابتدا
		hasPermission, cacheHit := permissionCache.getCachedPermission(userIDUint, permissionName)

		if !cacheHit {
			// Cache miss - کوئری از دیتابیس
			var err error
			hasPermission, err = CheckUserPermission(database.GormDB, userIDUint, permissionName)
			if err != nil {
				fmt.Printf("❌ خطا در بررسی دسترسی: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در بررسی دسترسی"})
				c.Abort()
				return
			}

			// ذخیره در cache
			permissionCache.setCachedPermission(userIDUint, permissionName, hasPermission)
			fmt.Printf("🔄 Permission '%s' برای کاربر %d از DB query شد و در cache ذخیره شد\n", permissionName, userIDUint)
		} else {
			fmt.Printf("⚡ Permission '%s' برای کاربر %d از cache بازیابی شد\n", permissionName, userIDUint)
		}

		if !hasPermission {
			fmt.Printf("❌ دسترسی '%s' برای کاربر %d وجود ندارد\n", permissionName, userObj.ID)
			c.JSON(http.StatusNotFound, gin.H{"error": "صفحه مورد نظر پیدا نشد"})
			c.Abort()
			return
		}

		fmt.Printf("✅ دسترسی '%s' تایید شد\n", permissionName)
		c.Next()
	}
}

// RequireAnyPermission میدلور بررسی حداقل یکی از دسترسی‌ها را ایجاد می‌کند
func RequireAnyPermission(permissionNames ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
			c.Abort()
			return
		}

		// بررسی نقش کاربر
		role, exists := c.Get("role")
		if exists && role == "customer" {
			c.JSON(http.StatusForbidden, gin.H{"error": "دسترسی غیرمجاز"})
			c.Abort()
			return
		}

		// بررسی دسترسی‌های کاربر با cache
		userIDUint := userID.(uint)
		for _, permissionName := range permissionNames {
			// بررسی cache ابتدا
			hasPermission, cacheHit := permissionCache.getCachedPermission(userIDUint, permissionName)

			if !cacheHit {
				// Cache miss - کوئری از دیتابیس
				var err error
				hasPermission, err = CheckUserPermission(database.GormDB, userIDUint, permissionName)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در بررسی دسترسی"})
					c.Abort()
					return
				}

				// ذخیره در cache
				permissionCache.setCachedPermission(userIDUint, permissionName, hasPermission)
			}

			if hasPermission {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "صفحه مورد نظر پیدا نشد"})
		c.Abort()
	}
}

// RequireAllPermissions میدلور بررسی تمام دسترسی‌ها را ایجاد می‌کند
func RequireAllPermissions(permissionNames ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
			c.Abort()
			return
		}

		// بررسی نقش کاربر
		role, exists := c.Get("role")
		if exists && role == "customer" {
			c.JSON(http.StatusForbidden, gin.H{"error": "دسترسی غیرمجاز"})
			c.Abort()
			return
		}

		// بررسی تمام دسترسی‌های کاربر با cache - باید همه مجوزها موجود باشند
		userIDUint := userID.(uint)
		for _, permissionName := range permissionNames {
			// بررسی cache ابتدا
			hasPermission, cacheHit := permissionCache.getCachedPermission(userIDUint, permissionName)

			if !cacheHit {
				// Cache miss - کوئری از دیتابیس
				var err error
				hasPermission, err = CheckUserPermission(database.GormDB, userIDUint, permissionName)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در بررسی دسترسی"})
					c.Abort()
					return
				}

				// ذخیره در cache
				permissionCache.setCachedPermission(userIDUint, permissionName, hasPermission)
			}

			if !hasPermission {
				c.JSON(http.StatusNotFound, gin.H{"error": "صفحه مورد نظر پیدا نشد"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// CheckUserPermission بررسی می‌کند که آیا کاربر دسترسی مورد نظر را دارد یا خیر
func CheckUserPermission(db *gorm.DB, userID uint, permissionName string) (bool, error) {
	// دریافت اطلاعات کاربر
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return false, err
	}

	// شرط ویژه: اگر کاربر dev باشد یا RoleID=100 داشته باشد، دسترسی کامل دارد
	if user.Username == "dev" || user.RoleID == 100 || strings.EqualFold(user.UserRole.Name, "developer") || strings.EqualFold(user.UserRole.Name, "admin") {
		return true, nil
	}

	// اگر کاربر customer است (RoleID=1) و هیچ نقش اضافی ندارد، دسترسی ادمین ندارد
	if user.RoleID == 1 {
		var extraRoles []uint
		if err := db.Table("user_roles").Select("role_id").Where("user_id = ? AND (valid_from IS NULL OR valid_from <= NOW()) AND (valid_to IS NULL OR valid_to >= NOW())", user.ID).Scan(&extraRoles).Error; err != nil {
			return false, err
		}
		if len(extraRoles) == 0 {
			return false, nil
		}
	}

	// استخراج تمامی role_id های کاربر (ستون users.role_id + نقش‌های اضافی)
	roleIDs := []uint{}
	if user.RoleID > 0 {
		roleIDs = append(roleIDs, user.RoleID)
	}
	var extraRoles []uint
	if err := db.Table("user_roles").Select("role_id").Where("user_id = ? AND (valid_from IS NULL OR valid_from <= NOW()) AND (valid_to IS NULL OR valid_to >= NOW())", user.ID).Scan(&extraRoles).Error; err != nil {
		return false, err
	}
	roleIDs = append(roleIDs, extraRoles...)

	// حذف نقش‌های تکراری
	roleSet := map[uint]struct{}{}
	uniqueRoleIDs := []uint{}
	for _, r := range roleIDs {
		if _, ok := roleSet[r]; !ok {
			roleSet[r] = struct{}{}
			uniqueRoleIDs = append(uniqueRoleIDs, r)
		}
	}

	if len(uniqueRoleIDs) == 0 {
		return false, nil
	}

	// بررسی دسترسی در جدول role_permissions برای همه نقش‌های کاربر
	var permissionCount int64
	err := db.Table("role_permissions").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id IN ? AND permissions.name = ? AND permissions.is_active = ?",
			uniqueRoleIDs, permissionName, true).
		Count(&permissionCount).Error

	if err != nil {
		return false, err
	}

	return permissionCount > 0, nil
}

// AdminMiddleware میدلور بررسی دسترسی ادمین
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
			c.Abort()
			return
		}

		// بررسی نقش ادمین - شامل تمام نقش‌های مدیریتی
		roleStr := role.(string)
		adminRoles := []string{"admin", "developer", "super_admin", "manager", "operator"}

		isAdmin := false
		for _, adminRole := range adminRoles {
			if roleStr == adminRole {
				isAdmin = true
				break
			}
		}

		if !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "دسترسی ادمین مورد نیاز است"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// OperatorMiddleware میدلور بررسی دسترسی اپراتور
func OperatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
			c.Abort()
			return
		}

		// بررسی نقش اپراتور یا ادمین
		roleStr := role.(string)
		if roleStr != "operator" && roleStr != "admin" && roleStr != "developer" {
			c.JSON(http.StatusForbidden, gin.H{"error": "دسترسی اپراتور مورد نیاز است"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireCustomerOrOwner middleware برای customer ها - فقط دسترسی به منابع شخصی
// این middleware چک می‌کند که کاربر customer باشد و فقط به داده‌های خودش دسترسی داشته باشد
func RequireCustomerOrOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
			c.Abort()
			return
		}

		// دریافت اطلاعات کاربر
		var user models.User
		if err := database.GormDB.First(&user, userID.(uint)).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر یافت نشد"})
			c.Abort()
			return
		}

		// اگر کاربر customer است، فقط به منابع خودش دسترسی دارد
		if user.RoleID == 1 {
			// بررسی اینکه آیا کاربر در حال دسترسی به منابع خودش است
			resourceUserID := c.Param("user_id")
			resourceID := c.Param("id")

			// اگر user_id در URL هست، باید با ID کاربر فعلی برابر باشد
			if resourceUserID != "" && resourceUserID != fmt.Sprintf("%d", user.ID) {
				c.JSON(http.StatusNotFound, gin.H{"error": "منبع مورد نظر یافت نشد"})
				c.Abort()
				return
			}

			// اگر id در URL هست و این مربوط به کاربر است، باید با ID کاربر فعلی برابر باشد
			if resourceID != "" && (c.FullPath() == "/api/users/:id" || strings.Contains(c.FullPath(), "/users/")) {
				if resourceID != fmt.Sprintf("%d", user.ID) {
					c.JSON(http.StatusNotFound, gin.H{"error": "منبع مورد نظر یافت نشد"})
					c.Abort()
					return
				}
			}
		}
		// اگر کاربر customer نیست (admin, operator, etc.) همه دسترسی‌ها را دارد

		c.Next()
	}
}

// GetUserPermissions تمام دسترسی‌های کاربر را برمی‌گرداند
func GetUserPermissions(userID uint) ([]string, error) {
	var user models.User
	if err := database.GormDB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	// developer یا RoleID=100 → همه دسترسی‌ها
	if user.Username == "dev" || user.RoleID == 100 {
		var permissions []models.Permission
		if err := database.GormDB.Find(&permissions).Error; err != nil {
			return nil, err
		}
		var names []string
		for _, p := range permissions {
			names = append(names, p.Name)
		}
		return names, nil
	}

	// جمع‌آوری role_id ها (ستون users + جدول user_roles)
	roleIDs := []uint{}
	if user.RoleID > 0 {
		roleIDs = append(roleIDs, user.RoleID)
	}
	var extra []uint
	if err := database.GormDB.Table("user_roles").Select("role_id").Where("user_id = ? AND (valid_from IS NULL OR valid_from <= NOW()) AND (valid_to IS NULL OR valid_to >= NOW())", user.ID).Scan(&extra).Error; err != nil {
		return nil, err
	}
	roleIDs = append(roleIDs, extra...)

	// حذف تکراری‌ها
	uniqueMap := map[uint]struct{}{}
	unique := []uint{}
	for _, r := range roleIDs {
		if _, ok := uniqueMap[r]; !ok {
			uniqueMap[r] = struct{}{}
			unique = append(unique, r)
		}
	}

	if len(unique) == 0 {
		return []string{}, nil
	}

	var permissions []models.Permission
	err := database.GormDB.Table("permissions").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id IN ? AND permissions.is_active = ?", unique, true).
		Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	var names []string
	for _, p := range permissions {
		names = append(names, p.Name)
	}
	return names, nil
}
