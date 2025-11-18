package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"
)

// RegisterMenuRoutes ثبت مسیرهای مربوط به مدیریت منوها
func RegisterMenuRoutes(public *gin.RouterGroup, db *gorm.DB) {
	// ایجاد repository، service و handler
	menuRepo := repository.NewMenuRepository(db)
	menuService := services.NewMenuService(menuRepo)
	menuHandler := handlers.NewMenuHandler(menuService)

	// گروه مسیرهای منو
	menus := public.Group("/menus")
	{
		// مسیرهای عمومی (بدون احراز هویت)
		menus.GET("", menuHandler.GetAllMenus)
		menus.GET("/slug/:slug", menuHandler.GetMenuBySlug)
		menus.GET("/location/:location", menuHandler.GetMenusByLocation)
		menus.GET("/enabled", menuHandler.GetEnabledMenus)
		menus.GET("/:id", menuHandler.GetMenuByID)

		// مسیرهای مدیریتی (نیاز به احراز هویت)
		menus.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.CreateMenu)
		menus.PUT("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UpdateMenu)
		menus.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.DeleteMenu)
	}

	// گروه مسیرهای آیتم‌های منو
	menuItems := public.Group("/menu-items")
	{
		// مسیرهای عمومی
		menuItems.GET("/:id", menuHandler.GetMenuItemByID)
		menuItems.GET("/menu/:menuId", menuHandler.GetMenuItemsByMenuID)

		// مسیرهای مدیریتی (نیاز به احراز هویت)
		menuItems.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.CreateMenuItem)
		menuItems.PUT("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UpdateMenuItem)
		menuItems.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.DeleteMenuItem)
	}

	// گروه مسیرهای ترتیب آیتم‌ها
	menuReorder := public.Group("/menus/:menuId/reorder")
	{
		menuReorder.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.ReorderMenuItems)
	}

	// گروه مسیرهای جایگاه‌های منو
	menuLocations := public.Group("/menu-locations")
	{
		// مسیرهای عمومی
		menuLocations.GET("", menuHandler.GetAllMenuLocations)
		menuLocations.GET("/:id", menuHandler.GetMenuLocationByID)
		menuLocations.GET("/slug/:slug", menuHandler.GetMenuLocationBySlug)

		// مسیرهای مدیریتی (نیاز به احراز هویت)
		menuLocations.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.CreateMenuLocation)
		menuLocations.PUT("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UpdateMenuLocation)
		menuLocations.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.DeleteMenuLocation)
	}

	// گروه مسیرهای ستون‌های منو
	menuColumns := public.Group("/menu-columns")
	{
		// مسیرهای عمومی
		menuColumns.GET("/menu/:menuId", menuHandler.GetMenuColumnsByMenuID)
		menuColumns.GET("/:id", menuHandler.GetMenuColumnByID)

		// مسیرهای مدیریتی (نیاز به احراز هویت)
		menuColumns.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.CreateMenuColumn)
		menuColumns.PUT("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UpdateMenuColumn)
		menuColumns.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.DeleteMenuColumn)
	}

	// گروه مسیرهای ترتیب ستون‌ها
	menuColumnReorder := public.Group("/menus/:menuId/reorder-columns")
	{
		menuColumnReorder.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.ReorderMenuColumns)
	}

	// گروه مسیرهای تخصیص منو
	menuAssignments := public.Group("/menu-assignments")
	{
		// مسیرهای عمومی
		menuAssignments.GET("/location/:locationId", menuHandler.GetMenuAssignmentsByLocationID)
		menuAssignments.GET("/menu/:menuId", menuHandler.GetMenuAssignmentsByMenuID)

		// مسیرهای مدیریتی (نیاز به احراز هویت)
		menuAssignments.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.CreateMenuAssignment)
		menuAssignments.PUT("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UpdateMenuAssignment)
		menuAssignments.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.DeleteMenuAssignment)
	}

	// گروه مسیرهای تخصیص منو به جایگاه (با مسیر متفاوت برای جلوگیری از تداخل)
	menuAssignment := public.Group("/menu-assign")
	{
		menuAssignment.POST("/menu/:menuId/location/:locationId", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.AssignMenuToLocation)
		menuAssignment.DELETE("/menu/:menuId/location/:locationId", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UnassignMenuFromLocation)
	}
}

// RegisterAdminMenuRoutes ثبت مسیرهای admin برای مدیریت منوها
func RegisterAdminMenuRoutes(admin *gin.RouterGroup, db *gorm.DB) {
	// ایجاد repository، service و handler
	menuRepo := repository.NewMenuRepository(db)
	menuService := services.NewMenuService(menuRepo)
	menuHandler := handlers.NewMenuHandler(menuService)

	// گروه مسیرهای admin منو
	menus := admin.Group("/menus")
	{
		// مسیرهای admin (نیاز به احراز هویت)
		menus.GET("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetAllMenus)
		menus.GET("/enabled", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetEnabledMenus)
		menus.GET("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuByID)
		menus.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.CreateMenu)
		menus.PUT("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UpdateMenu)
		menus.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.DeleteMenu)
	}

	// گروه مسیرهای admin آیتم‌های منو
	menuItems := admin.Group("/menu-items")
	{
		// مسیرهای admin (نیاز به احراز هویت)
		menuItems.GET("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuItemByID)
		menuItems.GET("/menu/:menuId", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuItemsByMenuID)
		menuItems.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.CreateMenuItem)
		menuItems.PUT("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UpdateMenuItem)
		menuItems.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.DeleteMenuItem)
	}

	// گروه مسیرهای admin ترتیب آیتم‌ها
	menuReorder := admin.Group("/menus/:menuId/reorder")
	{
		menuReorder.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.ReorderMenuItems)
	}

	// گروه مسیرهای محتوای منو
	menuContent := admin.Group("/menu-content")
	{
		// مسیرهای محتوای منو (نیاز به احراز هویت)
		menuContent.GET("/pages", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuContentPages)
		menuContent.GET("/posts", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuContentPosts)
		menuContent.GET("/categories", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuContentCategories)
		menuContent.GET("/product-categories", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuContentProductCategories)
	}

	// گروه مسیرهای admin جایگاه‌های منو
	menuLocations := admin.Group("/menu-locations")
	{
		// مسیرهای admin (نیاز به احراز هویت)
		menuLocations.GET("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetAllMenuLocations)
		menuLocations.GET("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuLocationByID)
		menuLocations.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.CreateMenuLocation)
		menuLocations.PUT("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UpdateMenuLocation)
		menuLocations.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.DeleteMenuLocation)
	}

	// گروه مسیرهای admin ستون‌های منو
	menuColumns := admin.Group("/menu-columns")
	{
		// مسیرهای admin (نیاز به احراز هویت)
		menuColumns.GET("/menu/:menuId", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuColumnsByMenuID)
		menuColumns.GET("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuColumnByID)
		menuColumns.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.CreateMenuColumn)
		menuColumns.PUT("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UpdateMenuColumn)
		menuColumns.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.DeleteMenuColumn)
	}

	// گروه مسیرهای admin ترتیب ستون‌ها
	menuColumnReorder := admin.Group("/menus/:menuId/reorder-columns")
	{
		menuColumnReorder.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.ReorderMenuColumns)
	}

	// گروه مسیرهای admin تخصیص منو
	menuAssignments := admin.Group("/menu-assignments")
	{
		// مسیرهای admin (نیاز به احراز هویت)
		menuAssignments.GET("/location/:locationId", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuAssignmentsByLocationID)
		menuAssignments.GET("/menu/:menuId", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.GetMenuAssignmentsByMenuID)
		menuAssignments.POST("", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.CreateMenuAssignment)
		menuAssignments.PUT("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UpdateMenuAssignment)
		menuAssignments.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.DeleteMenuAssignment)
	}

	// گروه مسیرهای admin تخصیص منو به جایگاه (با مسیر متفاوت برای جلوگیری از تداخل)
	menuAssignment := admin.Group("/menu-assign")
	{
		menuAssignment.POST("/menu/:menuId/location/:locationId", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.AssignMenuToLocation)
		menuAssignment.DELETE("/menu/:menuId/location/:locationId", middleware.Auth(), middleware.RequirePermission("menus_manage"), menuHandler.UnassignMenuFromLocation)
	}
}
