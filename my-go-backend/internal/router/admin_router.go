package router

import (
	"os"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"
)

// RegisterAdminRoutes adds /api/admin routes such as settings, SMS gateways, orders, etc.
func RegisterAdminRoutes(
	r *gin.Engine,
	db *gorm.DB,
	redisClient *redis.Client,
	settingHandler *handlers.SettingHandler,
	adminSettingHandler *handlers.AdminSettingHandler,
) {
	// --- SMS Gateway ---
	smsGatewayRepo := repository.NewSMSGatewayRepository(db)
	smsGatewayHandler := handlers.NewSMSGatewayHandler(smsGatewayRepo)
	r.POST("/api/sms-gateways", smsGatewayHandler.Create)
	r.GET("/api/sms-gateways", smsGatewayHandler.List)
	r.GET("/api/sms-gateways/:id", smsGatewayHandler.GetByID)
	r.PUT("/api/sms-gateways/:id", smsGatewayHandler.Update)
	r.DELETE("/api/sms-gateways/:id", smsGatewayHandler.Delete)
	r.POST("/api/sms-gateways/:id/test-connection", smsGatewayHandler.TestConnection)
	r.POST("/api/sms-gateways/:id/send-test", smsGatewayHandler.SendTestSMS)              // ارسال پیامک تستی
	r.GET("/api/sms-gateways/:id/inbox", smsGatewayHandler.GetInbox)                      // دریافت پیام‌های دریافتی (IPPanel)
	r.GET("/api/sms-gateways/:id/inbox/test", smsGatewayHandler.TestInbox)                // تست inbox (برای توسعه)
	r.GET("/api/sms-gateways/:id/debug", smsGatewayHandler.GetGatewayDebugInfo)           // اطلاعات دیباگ درگاه
	r.GET("/api/sms-gateways/:id/balance", smsGatewayHandler.GetGatewayBalance)           // دریافت موجودی درگاه
	r.GET("/api/sms-gateways/:id/melipayamak-info", smsGatewayHandler.GetMeliPayamakInfo) // دریافت اطلاعات کامل ملی پیامک
	r.PUT("/api/sms-gateways/priorities", smsGatewayHandler.UpdatePriorities)             // به‌روزرسانی اولویت‌های درگاه‌ها
	// مسیر dev بدون پیشوند /api برای فراخوانی‌های Nuxt
	r.PUT("/sms-gateways/priorities", smsGatewayHandler.UpdatePriorities)
	r.POST("/api/sms-gateways/:id/test-melipayamak", smsGatewayHandler.TestMeliPayamakDirect) // تست مستقیم ملی پیامک

	// --- فراز اس‌ام‌اس Outbox ---
	r.GET("/api/farazsms-outbox", handlers.GetFarazSMSOutbox) // دریافت پیامک‌های خروجی فراز اس‌ام‌اس

	// ---------------- Admin Settings ----------------
	// فعال‌سازی مسیرهای تنظیمات ادمین که قبلاً غیرفعال شده بودند
	admin := r.Group("/api/admin")

	admin.Use(middleware.Auth())  // احراز هویت فعال شد
	admin.Use(middleware.Admin()) // بررسی نقش ادمین (admin یا developer)

	settingsGroup := admin.Group("/settings")

	{
		settingsGroup.GET("", settingHandler.GetSettings)
		settingsGroup.GET("/auth", settingHandler.GetAuthSettings)
		settingsGroup.PUT("/auth", settingHandler.UpdateAuthSettings)
		// تنظیمات نظرات
		settingsGroup.PUT("/reviews", settingHandler.UpdateReviewSettings)
		settingsGroup.GET("/category/:category", settingHandler.GetSettingsByCategory)
		settingsGroup.GET("/:key", settingHandler.GetSetting)
		settingsGroup.PUT("/:key", settingHandler.UpdateSetting)
		settingsGroup.PUT("", settingHandler.UpdateSettings)
		settingsGroup.DELETE("/cache", settingHandler.ClearCache)

	}

	// مسیرهای تنظیمات فاکتور و چاپ
	admin.GET("/invoice-print-settings", settingHandler.GetInvoicePrintSettings)
	admin.PUT("/invoice-print-settings", settingHandler.UpdateInvoicePrintSettings)

	adminSettings := admin.Group("/admin-settings", middleware.Auth())
	{
		adminSettings.GET("/:key", adminSettingHandler.GetAdminSetting)
		adminSettings.PUT("/:key", adminSettingHandler.UpdateAdminSetting)
		adminSettings.GET("/bulk-edit-columns", adminSettingHandler.GetBulkEditColumns)
		adminSettings.PUT("/bulk-edit-columns", adminSettingHandler.UpdateBulkEditColumns)
	}

	// مسیر بدون /api برای سازگاری با فرانت
	r.GET("/admin/admin-settings/bulk-edit-columns", middleware.Auth(), adminSettingHandler.GetBulkEditColumns)
	r.PUT("/admin/admin-settings/bulk-edit-columns", middleware.Auth(), adminSettingHandler.UpdateBulkEditColumns)

	userHandler := handlers.NewUserHandler(db)
	adminUsers := admin.Group("/users")
	{
		adminUsers.POST("/admin-register", userHandler.AdminRegisterUser)
	}
	// admin.POST("/upload-video", handlers.UploadVideoHandler)
	// --- پایان تنظیمات و هندلرهای ادمین ---

	// بخش تنظیمات و هندلرهای ادمین رو دوباره فعال می کنم
	orderHandler := handlers.NewOrderHandler(db)
	ordersAnalyticsHandler := handlers.NewOrdersAnalyticsHandler(db)

	adminOrders := admin.Group("/orders", middleware.Auth(), middleware.Admin())
	{
		adminOrders.GET("", orderHandler.ListAllOrders)
		adminOrders.GET("/stats", orderHandler.AdminOrderStats)
		adminOrders.GET("/reports", orderHandler.AdminOrderReports)
		adminOrders.GET("/analytics", ordersAnalyticsHandler.GetAllOrdersAnalytics)
		adminOrders.GET("/advanced-analytics", ordersAnalyticsHandler.GetAdvancedAnalytics)
		adminOrders.GET("/in-progress", orderHandler.ListInProgressOrders)
		adminOrders.GET("/in-progress/stats", orderHandler.InProgressOrderStats)
		adminOrders.GET("/processing", orderHandler.ListProcessingOrders)
		adminOrders.GET("/processing/stats", orderHandler.ProcessingOrderStats)
		adminOrders.GET("/shipped", orderHandler.ListShippedOrders)
		adminOrders.GET("/shipped/stats", orderHandler.ShippedOrderStats)
		adminOrders.GET("/returned", orderHandler.ListReturnedOrders)
		adminOrders.GET("/returned/stats", orderHandler.ReturnedOrderStats)
		adminOrders.GET("/refunded", orderHandler.ListRefundedOrders)
		adminOrders.GET("/refunded/stats", orderHandler.RefundedOrderStats)
		adminOrders.GET("/cancelled", orderHandler.ListCancelledOrders)
		adminOrders.GET("/cancelled/stats", orderHandler.CancelledOrderStats)
		adminOrders.GET("/:id/items", orderHandler.GetOrderItems)
		adminOrders.PUT("/:id/status", orderHandler.UpdateOrderStatus)
		adminOrders.GET("/:id", orderHandler.GetOrderByID)
	}

	// محصولات برای پنل ادمین (فراخوانی فرانت: /api/admin/products)
	productHandler := handlers.NewProductHandler(db)
	admin.GET("/products", productHandler.ListProducts)
	// مسیر بدون /api برای سازگاری با فرانت
	r.GET("/admin/products", middleware.Auth(), middleware.Admin(), productHandler.ListProducts)

	// ---------------- Product Categories ----------------

	// --- Admin Menu Routes ---
	RegisterAdminMenuRoutes(admin, db)

	// --- API Settings Routes ---
	// فعال‌سازی مسیرهای تنظیمات API
	apiSettingsService := services.NewAPISettingsService(db)
	apiSettingsHandler := handlers.NewAPISettingsHandler(apiSettingsService)
	SetupAPISettingsRoutes(admin, apiSettingsHandler)

	// --- Fraud Routes ---
	uowFactory := unitofwork.NewUnitOfWorkFactory(db)
	salt := os.Getenv("FRAUD_SALT")
	if salt == "" {
		salt = "internal_salt"
	}
	fraudService := services.NewFraudService(uowFactory, salt)
	fraudHandler := handlers.NewFraudHandler(fraudService, uowFactory)

	adminFraud := admin.Group("/fraud", middleware.Auth(), middleware.RequirePermission("fraud.manage"), middleware.RateLimitWith(120, 1*time.Minute, "admin"))
	{
		adminFraud.GET("/cases", fraudHandler.ListCasesAdvanced)
		adminFraud.GET("/cases/:id", fraudHandler.GetCaseDetail)
		adminFraud.POST("/cases/:id/decision", fraudHandler.Decide)
		adminFraud.GET("/rules", fraudHandler.ListRules)
		adminFraud.PUT("/rules/:id", fraudHandler.UpdateRule)
		adminFraud.GET("/stats", fraudHandler.Stats)
		adminFraud.POST("/re-evaluate-open", fraudHandler.ReEvaluateOpenCases)
		adminFraud.POST("/list/whitelist", fraudHandler.UpsertWhitelist)
		adminFraud.POST("/list/blacklist", fraudHandler.UpsertBlacklist)
	}
	// system hooks
	r.POST("/api/fraud/evaluate/order/:orderId", middleware.RateLimitWith(60, 1*time.Minute, "evaluate"), fraudHandler.EvaluateOrder)

	// --- Wallet Admin Routes ---
	walletService := services.NewWalletService(uowFactory)
	walletAdminHandler := handlers.NewWalletAdminHandler(walletService)
	adminWallet := admin.Group("/wallet", middleware.Auth(), middleware.RequireAnyPermission("wallet.manage", "finance.manage"))
	{
		adminWallet.POST("/credit", walletAdminHandler.Credit)
		adminWallet.POST("/debit", walletAdminHandler.Debit)
		adminWallet.POST("/transfer", walletAdminHandler.Transfer)
		adminWallet.POST("/block/:userId", walletAdminHandler.Block)
		adminWallet.POST("/unblock/:userId", walletAdminHandler.Unblock)
		adminWallet.GET("/transactions", walletAdminHandler.AdminListTransactions)
		adminWallet.POST("/transactions/:id/status", walletAdminHandler.UpdateTxStatus)
		adminWallet.GET("/summary", walletAdminHandler.Summary)
		adminWallet.GET("/trend", walletAdminHandler.Trend)
	}

	// --- Bank Accounts Admin Routes ---
	bankAccountService := services.NewBankAccountService(uowFactory)
	bankAccountHandler := handlers.NewBankAccountAdminHandler(bankAccountService)
	adminBank := admin.Group("/bank-accounts", middleware.Auth(), middleware.RequireAnyPermission("wallet.manage", "finance.manage"))
	{
		adminBank.GET("", bankAccountHandler.List)
		adminBank.POST(":id/verify", bankAccountHandler.Verify)
		adminBank.POST(":id/block", bankAccountHandler.Block)
		adminBank.POST(":id/unblock", bankAccountHandler.Unblock)
		adminBank.POST(":id/reject", bankAccountHandler.Reject)
	}

	// --- Header Settings Admin Routes ---
	// Note: Header admin routes are now handled by RegisterHeaderRoutes in header_router.go

	// --- Footer Settings Admin Routes ---
	// Note: Footer admin routes are now handled by RegisterFooterRoutes in footer_router.go

	// --- System Routes ---
	systemHandler := handlers.NewSystemHandler(db)
	adminSystem := admin.Group("/check-table")
	{
		adminSystem.GET("/:tableName", systemHandler.CheckTable)
	}

	// --- Product Categories Admin Routes ---
	productCategoryHandler := handlers.NewProductCategoryHandler(db)
	adminCategories := admin.Group("/product-categories")
	{
		adminCategories.GET("", productCategoryHandler.ListCategories)
		adminCategories.POST("", productCategoryHandler.CreateCategory)
		adminCategories.GET("/:id", productCategoryHandler.GetCategory)
		adminCategories.PUT("/:id", productCategoryHandler.UpdateCategory)
		adminCategories.DELETE("/:id", productCategoryHandler.DeleteCategory)
		adminCategories.POST("/reorder", productCategoryHandler.ReorderCategories)
	}

	// Also register routes without admin prefix for dev compatibility
	r.GET("/api/check-table/:tableName", systemHandler.CheckTable)

	// --- Digikala Import Routes ---
	categoryRepo := repository.NewCategoryRepository(db)
	digikalaImportService := services.NewDigikalaImportService(db, categoryRepo)
	digikalaImportHandler := handlers.NewDigikalaImportHandler(db, digikalaImportService)

	digikalaGroup := admin.Group("/digikala")
	{
		digikalaGroup.GET("/stats", digikalaImportHandler.GetStats)
		digikalaGroup.GET("/import-history", digikalaImportHandler.GetImportHistory)
		digikalaGroup.POST("/validate-url", digikalaImportHandler.ValidateURL)
		digikalaGroup.POST("/start-import", digikalaImportHandler.StartImport)
		digikalaGroup.GET("/import-progress/:id", digikalaImportHandler.GetImportProgress)
		digikalaGroup.POST("/cancel-import/:id", digikalaImportHandler.CancelImport)
		digikalaGroup.POST("/retry-import/:id", digikalaImportHandler.RetryImport)
		digikalaGroup.GET("/categories", digikalaImportHandler.GetAvailableCategories)
		digikalaGroup.GET("/categories/search", digikalaImportHandler.SearchCategories)
		digikalaGroup.POST("/categories/import", digikalaImportHandler.ImportCategories)
		digikalaGroup.GET("/categories/:id", digikalaImportHandler.GetCategoryDetails)
		digikalaGroup.POST("/categories/import/:type", digikalaImportHandler.ImportSpecificCategories)
		digikalaGroup.GET("/status", digikalaImportHandler.GetImportStatus)
	}

}
