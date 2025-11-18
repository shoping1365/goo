package router

import (
	"os"
	"path/filepath"
	"time"

	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	searchservice "my-go-backend/internal/search/service"

	"github.com/gin-gonic/gin"

	"my-go-backend/internal/services"
	trafficSvc "my-go-backend/internal/services"

	"my-go-backend/internal/database/unitofwork"
	repoPkg "my-go-backend/internal/repository"

	"github.com/redis/go-redis/v9"
)

/*
روتر اصلی برنامه

این فایل برای تعریف مسیرهای API استفاده می‌شود و شامل موارد زیر است:
- مسیرهای عمومی
- مسیرهای احراز هویت
- مسیرهای ادمین
- میدلورهای مورد نیاز
*/

// ----------------------------------------------------------------------------------
// ⚠️  IMPORTANT FOR DEVELOPERS & CODE-GENERATION TOOLS (INCLUDING AI ASSISTANTS) ⚠️
// ----------------------------------------------------------------------------------
// این فایل فقط نقش "تجمیع‌کننده" را دارد. هیچ Group یا Endpoint جدیدی نباید مستقیماً
// در اینجا تعریف شود. برای افزودن مسیر:
//   1) یک فایل router جداگانه در مسیر internal/router/ ایجاد کنید (در صورت نبود).
//   2) یک تابع RegisterXYZRoutes تعریف کرده و مسیرهای جدید را در آن ثبت کنید.
//   3) در اینجا تنها همان تابع RegisterXYZRoutes را فراخوانی کنید.
// به این ترتیب ساختار روتر ماژولار و قابل نگهداری باقی می‌ماند.
// ----------------------------------------------------------------------------------
// sentinel: ROUTER_AGGREGATOR_DO_NOT_DEFINE_ROUTES_HERE
// ----------------------------------------------------------------------------------

// SetupRouter روتر اصلی برنامه را تنظیم می‌کند
func SetupRouter(
	settingHandler *handlers.SettingHandler,
	settingService *services.SettingService,
	gormDB *gorm.DB,
	redisClient *redis.Client,
	adminSettingHandler *handlers.AdminSettingHandler,
	paymentHandler *handlers.PaymentHandler,
	mellatHandler *handlers.MellatHandler,
	chatHandler *handlers.ChatHandler,
	wsManager *services.WebSocketManager,
	monitoringHandler *handlers.MonitoringHandler,
	searchSvc *searchservice.Service,
	searchIndexes []string,
) *gin.Engine {
	// اطمینان از چاپ لاگ‌ها در کنسول
	gin.SetMode(gin.DebugMode)
	gin.ForceConsoleColor()
	gin.DefaultWriter = os.Stdout
	gin.DefaultErrorWriter = os.Stderr

	r := gin.New()

	// میدلورهای اصلی (بدون TrafficLogger)
	r.Use(gin.Recovery(), middleware.CustomLogger(), middleware.BlockBackupAccess(), middleware.SecurityHeaders(), middleware.CORS())

	// میدلور برای قرار دادن دیتابیس در context (قبل از TrafficLogger تا بتواند تنظیمات را بخواند)
	r.Use(func(c *gin.Context) {
		c.Set("db", gormDB)
		c.Next()
	})

	// افزودن TrafficLogger پس از ست شدن DB در context
	r.Use(middleware.TrafficLogger())

	// Initialize handlers
	userHandler := handlers.NewUserHandler(gormDB)
	userManagementHandler := handlers.NewUserManagementHandler(gormDB)

	// settingHandler باید از main.go پاس داده شود
	// چون نیاز به Redis دارد

	// تنظیم احراز هویت یکپارچه
	authService := services.NewAuthService(gormDB)
	smsService := services.NewSMSService(gormDB)
	unifiedAuthService := services.NewUnifiedAuthService(gormDB, authService, smsService)
	unifiedAuthHandler := handlers.NewUnifiedAuthHandler(gormDB, unifiedAuthService, authService)

	// تنظیم سرویس احراز هویت کاربران (KYC)
	verificationRepo := repoPkg.NewUserVerificationRepository(gormDB)
	userRepo := repoPkg.NewUserRepository(gormDB)
	verificationService := services.NewUserVerificationService(verificationRepo, userRepo)
	verificationHandler := handlers.NewUserVerificationHandler(verificationService)

	// مسیرهای عمومی
	public := r.Group("/api")

	// Debug handler
	debugHandler := handlers.NewDebugHandler(gormDB)

	// Debug routes
	public.GET("/debug/dev-user", debugHandler.GetDevUser)
	// مسیرهای قدیمی بدون پیشوند /api برای سازگاری با پروکسی dev
	publicNoPrefix := r.Group("")

	// API root endpoint
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Server",
			"version": "1.0",
			"status":  "running",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// Unified Auth endpoints (سیستم احراز هویت یکپارچه)
	SetupUnifiedAuthRoutes(public, unifiedAuthHandler)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	SetupUnifiedAuthRoutes(publicNoPrefix, unifiedAuthHandler)

	// حذف سیستم قدیمی - فقط سیستم یکپارچه فعال است

	// مسیرهای تنظیمات عمومی
	public.GET("/settings", settingHandler.GetPublicSettings)
	public.GET("/settings/compression", settingHandler.GetCompressionSettings)
	public.GET("/settings/video-compression", settingHandler.GetVideoCompressionSettings)
	public.PUT("/settings", settingHandler.UpdateSettings)

	// مسیرهای تنظیمات بدون پیشوند /api (برای dev)
	publicNoPrefix.GET("/settings", settingHandler.GetPublicSettings)
	publicNoPrefix.GET("/settings/compression", settingHandler.GetCompressionSettings)
	publicNoPrefix.GET("/settings/video-compression", settingHandler.GetVideoCompressionSettings)
	publicNoPrefix.PUT("/settings", settingHandler.UpdateSettings)

	// User Management endpoints
	SetupUserManagementRoutes(r, userManagementHandler)

	// Device Management Routes (مدیریت دستگاه‌های لاگین شده)
	SetupDeviceManagementRoutes(r, gormDB)

	// User endpoints
	RegisterUserRoutes(public, userHandler)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	RegisterUserRoutes(publicNoPrefix, userHandler)

	// User Verification endpoints (KYC - احراز هویت کاربران)
	RegisterUserVerificationRoutes(public, verificationHandler)
	RegisterUserVerificationRoutes(publicNoPrefix, verificationHandler)

	// User Banking Info endpoints (اطلاعات بانکی کاربران)
	bankingInfoRepo := repoPkg.NewUserBankingInfoRepository(gormDB)
	bankingInfoService := services.NewUserBankingInfoService(bankingInfoRepo)
	bankingInfoHandler := handlers.NewUserBankingInfoHandler(bankingInfoService)
	SetupUserBankingInfoRoutes(public.Group("/api"), bankingInfoHandler)
	SetupUserBankingInfoRoutes(publicNoPrefix.Group("/api"), bankingInfoHandler)

	// Recent Views endpoints - فقط اعلان handler (ثبت route ها بعداً انجام می‌شود)
	recentViewHandler := handlers.NewRecentViewHandler(gormDB)

	// User address endpoints
	RegisterUserAddressRoutes(public, gormDB)

	// Product endpoints (CRUD, images, variants, specs, videos)
	RegisterProductRoutes(public, gormDB)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	RegisterProductRoutes(publicNoPrefix, gormDB)

	// Product categories endpoints
	RegisterProductCategoryRoutes(public, gormDB)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	RegisterProductCategoryRoutes(publicNoPrefix, gormDB)

	// Post categories endpoints
	RegisterPostCategoryRoutes(public, gormDB)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	RegisterPostCategoryRoutes(publicNoPrefix, gormDB)

	// Post endpoints
	RegisterPostRoutes(public, gormDB)

	// Category-Brand pages endpoints
	RegisterCategoryBrandPageRoutes(public, gormDB)

	// QA endpoints (questions & categories)
	RegisterQARoutes(public, gormDB)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	RegisterQARoutes(publicNoPrefix, gormDB)

	// Review endpoints
	RegisterReviewRoutes(public, gormDB)

	// Geo endpoints (provinces & cities)
	RegisterGeoRoutes(public, gormDB)

	// Media endpoints
	RegisterMediaRoutes(public, gormDB, settingService)
	// در محیط توسعه بدون پیشوند /api
	RegisterMediaRoutes(publicNoPrefix, gormDB, settingService)

	// Image SEO admin aliases (/api/admin/image-seo/* and /admin/image-seo/*)
	RegisterImageSEORoutes(r, gormDB)

	// Video endpoints (separate from general media)
	RegisterVideoRoutes(public, gormDB, settingService)

	// Brand endpoints
	RegisterBrandRoutes(public, gormDB)
	// در محیط توسعه بدون پیشوند /api
	RegisterBrandRoutes(publicNoPrefix, gormDB)

	// Register attribute and attribute-value endpoints
	RegisterAttributeRoutes(public, gormDB)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	RegisterAttributeRoutes(publicNoPrefix, gormDB)

	// Register attribute-group endpoints
	RegisterAttributeGroupRoutes(public, gormDB)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	RegisterAttributeGroupRoutes(publicNoPrefix, gormDB)

	// Category attribute schema & filters
	RegisterAttributeSchemaRoutes(public, gormDB)

	// Schema template endpoints
	SetupSchemaRoutes(r, gormDB)

	// SEO Generator endpoints
	RegisterSeoGeneratorRoutes(r, gormDB)

	// SEO Redirects endpoints (admin)
	RegisterSEORedirectRoutes(r, gormDB)

	// Developer maintenance endpoints
	RegisterMaintenanceRoutes(public, gormDB)

	// Product specs endpoints
	RegisterProductSpecRoutes(public, gormDB)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	RegisterProductSpecRoutes(publicNoPrefix, gormDB)

	// Product pricing endpoints
	RegisterProductPriceRoutes(public, gormDB)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	RegisterProductPriceRoutes(publicNoPrefix, gormDB)

	// Product inventory endpoints
	RegisterProductInventoryRoutes(public, gormDB)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	RegisterProductInventoryRoutes(publicNoPrefix, gormDB)

	// Product warehouse stock endpoints (multi-warehouse)
	RegisterProductWarehouseStockRoutes(public, gormDB)
	// Dev routes without /api prefix (for local proxy compatibility)
	RegisterProductWarehouseStockRoutes(publicNoPrefix, gormDB)

	// Notify (stock/discount) endpoints
	SetupNotificationRoutes(r, public, gormDB)
	// Product shipping endpoints
	RegisterProductShippingRoutes(public, gormDB)

	// Cart endpoints
	cartHandler := handlers.NewCartHandler(gormDB)
	SetupCartRoutes(public, cartHandler)
	// Dev routes without /api prefix (for local proxy compatibility)
	SetupCartRoutes(publicNoPrefix, cartHandler)

	// Widget endpoints
	widgetHandler := handlers.NewWidgetHandler(gormDB)
	SetupWidgetRoutes(public, widgetHandler)

	// Menu endpoints
	RegisterMenuRoutes(public, gormDB)

	// Search endpoints
	SetupSearchRoutes(public, searchSvc, searchIndexes)

	// Role endpoints
	SetupRoleRoutes(r, gormDB)

	// Permission endpoints
	SetupPermissionRoutes(r, gormDB)

	// Role permission endpoints
	SetupRolePermissionRoutes(r, gormDB)

	// Order endpoints
	RegisterOrderRoutes(public, gormDB)

	// Admin routes (settings, admin-register)
	RegisterAdminRoutes(r, gormDB, redisClient, settingHandler, adminSettingHandler)

	// Shop settings routes (both with and without /api prefix)
	shopSettingsHandler := handlers.NewShopSettingsHandler(settingService)
	adminAPI := r.Group("/api/admin")
	adminAPI.Use(middleware.Auth())  // احراز هویت فعال شد
	adminAPI.Use(middleware.Admin()) // بررسی نقش ادمین (admin یا developer)
	adminDev := r.Group("/admin")    // بدون پیشوند برای dev

	// Recent Views routes (with admin support)
	RegisterRecentViewRoutes(public, adminAPI, recentViewHandler)
	// برای درخواست‌های بدون پیشوند /api در محیط توسعه
	RegisterRecentViewRoutes(publicNoPrefix, nil, recentViewHandler)

	// Header routes
	RegisterHeaderRoutes(public, adminAPI, gormDB, redisClient)
	// Header routes for dev (without /api prefix)
	RegisterHeaderRoutes(publicNoPrefix, adminDev, gormDB, redisClient)

	// Footer routes
	RegisterFooterRoutes(public, adminAPI, gormDB, redisClient)
	// Footer routes for dev (without /api prefix)
	RegisterFooterRoutes(publicNoPrefix, adminDev, gormDB, redisClient)

	// Mobile App Header routes
	RegisterMobileAppHeaderRoutes(public, adminAPI, gormDB, redisClient)
	// Mobile App Header routes for dev (without /api prefix)
	RegisterMobileAppHeaderRoutes(publicNoPrefix, adminDev, gormDB, redisClient)

	// Mobile App Footer routes
	RegisterMobileAppFooterRoutes(public, adminAPI, gormDB, redisClient)
	// Mobile App Footer routes for dev (without /api prefix)
	RegisterMobileAppFooterRoutes(publicNoPrefix, adminDev, gormDB, redisClient)

	// Mobile App Navigation routes
	RegisterMobileAppNavigationRoutes(public, adminAPI, gormDB, redisClient)
	// Mobile App Navigation routes for dev (without /api prefix)
	RegisterMobileAppNavigationRoutes(publicNoPrefix, adminDev, gormDB, redisClient)

	// Widget routes for admin (with authentication)
	SetupWidgetRoutes(adminAPI, widgetHandler)
	SetupShopSettingsRoutes(adminAPI, public, shopSettingsHandler)
	SetupShopSettingsRoutes(adminDev, publicNoPrefix, shopSettingsHandler)

	// Installment (admin) routes
	RegisterInstallmentRoutes(r, gormDB)

	// SMS Gateway routes
	smsGatewayRepo := repoPkg.NewSMSGatewayRepository(gormDB)
	smsGatewayHandler := handlers.NewSMSGatewayHandler(smsGatewayRepo)
	RegisterSMSGatewayRoutes(r, smsGatewayHandler)

	// SMS Pattern routes
	SetupSMSPatternRoutes(r, gormDB)

	// Payment Gateway routes
	paymentGatewayRepo := repoPkg.NewPaymentGatewayRepository(gormDB)
	paymentGatewayHandler := handlers.NewPaymentGatewayHandler(paymentGatewayRepo)
	RegisterPaymentGatewayRoutes(r, paymentGatewayHandler)

	// Payment routes (create, verify, manage transactions)
	SetupPaymentRoutes(r, paymentHandler, mellatHandler)

	// Melli payment routes
	paymentTransactionRepo := repoPkg.NewPaymentTransactionRepository(gormDB)
	melliHandler := handlers.NewMelliHandler(paymentGatewayRepo, paymentTransactionRepo)
	SetupMelliRoutes(r, melliHandler)

	// Saman payment routes
	samanHandler := handlers.NewSamanHandler(paymentGatewayRepo, paymentTransactionRepo)
	SetupSamanRoutes(r, samanHandler)

	// Parsian payment routes
	parsianHandler := handlers.NewParsianHandler(paymentGatewayRepo, paymentTransactionRepo)
	SetupParsianRoutes(r, parsianHandler)

	// Stripe payment routes
	stripeHandler := handlers.NewStripeHandler(paymentGatewayRepo, paymentTransactionRepo)
	SetupStripeRoutes(r, stripeHandler)

	// Chat routes
	RegisterChatRoutes(r, chatHandler, wsManager)

	// Warehouse admin routes
	RegisterWarehouseRoutes(r, gormDB)

	// Monitoring routes
	RegisterMonitoringRoutes(r, monitoringHandler)
	// Traffic monitoring routes
	RegisterTrafficMonitoringRoutes(r, gormDB)

	// System Security › Login routes
	RegisterSystemSecurityLoginRoutes(r, gormDB)

	// Firewall routes (System-level)
	// ایجاد سرویس و هندلر فایروال با رعایت CQRS + UoW
	firewallUoWFactory := unitofwork.NewUnitOfWorkFactory(gormDB)
	firewallRepo := repoPkg.NewFirewallRepositoryWindows()
	firewallSvc := services.NewFirewallService(firewallUoWFactory, firewallRepo)
	firewallHandler := handlers.NewFirewallHandler(firewallSvc)
	RegisterFirewallRoutes(r, firewallHandler)

	// Bot Shield middleware (global)
	r.Use(middleware.BotShield(settingService, gormDB))

	// Bot management routes
	RegisterBotRoutes(r, gormDB, settingService)

	// راه‌اندازی worker درج لاگ‌های ترافیک
	trafficUoWFactory := unitofwork.NewUnitOfWorkFactory(gormDB)
	trafficSvc.StartTrafficLogWorker(trafficUoWFactory)

	// System routes for admin developer pages
	RegisterSystemRoutes(r)
	RegisterDBPoolSettingsRoutes(r, gormDB, settingService)

	// Performance routes for admin developer pages
	RegisterPerformanceRoutes(r)

	// WordPress Migration routes (admin)
	RegisterWordPressMigrationRoutes(r, gormDB)

	// Static file serving - استفاده از مسیر پروژه اصلی (یک سطح بالاتر)
	uploadsPath := os.Getenv("UPLOADS_PATH")
	if uploadsPath == "" {
		// مسیر به پوشه public در ریشه پروژه (یک دایرکتوری بالاتر از my-go-backend)
		uploadsPath = "../public/uploads"
	}
	publicPath := os.Getenv("PUBLIC_PATH")
	if publicPath == "" {
		publicPath = "../public"
	}

	// تبدیل به absolute path برای اطمینان
	uploadsAbsPath, _ := filepath.Abs(uploadsPath)
	publicAbsPath, _ := filepath.Abs(publicPath)

	r.Static("/uploads", uploadsAbsPath)
	r.Static("/public", publicAbsPath)

	// Serve favicon
	faviconPath := filepath.Join(publicAbsPath, "favicon.ico")
	r.StaticFile("/favicon.ico", faviconPath) // Root path endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go Backend API Server",
			"status":  "running",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	return r
}
