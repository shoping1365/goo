package database

import (
	"database/sql"
	"fmt"
	"log"
	"my-go-backend/internal/database/migrations"

	"os"
	"time"

	"gorm.io/gorm"
)

// RunMigrations اجرای همهٔ مایگریشن‌ها با رجیستری نسخه‌مند و قفل یکتا
func RunMigrations(db *gorm.DB) error {
	// اجرای پیش‌فرض مایگریشن‌ها؛ فقط در صورت تنظیم صریح MIGRATIONS_ENABLED=false غیرفعال می‌شود
	if os.Getenv("MIGRATIONS_ENABLED") == "false" {
		log.Println("[MIGRATIONS] Disabled (set MIGRATIONS_ENABLED=true or unset to run)")
		return nil
	}
	// قفل یکتا با advisory lock
	if err := withAdvisoryLock(db, 928374, func(tx *gorm.DB) error {
		// اطمینان از وجود جدول رهگیری نسخه‌ها
		if err := tx.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
            version INT PRIMARY KEY,
            name TEXT NOT NULL,
            applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
        );`).Error; err != nil {
			return err
		}

		// رجیستری داخل Runner (ثابت و نسخه‌مند)
		regs := []struct {
			Number int
			Name   string
			Up     func(*gorm.DB) error
		}{
			{1, "schema_ledger", migrations.Up001SchemaMigrationsLedger},
			{2, "core_schema", migrations.Up002CoreSchema},
			{3, "catalog", migrations.Up003Catalog},
			{4, "products_core", migrations.Up004ProductsCore},
			{5, "orders_core", migrations.Up005OrdersCore},
			{6, "user_addresses", migrations.Up006UserAddresses},
			{7, "attributes", migrations.Up007Attributes},
			{8, "media", migrations.Up008Media},
			{9, "cms", migrations.Up009CMS},
			{10, "auth_logs_wallet", migrations.Up010AuthLogsWallet},
			{11, "notifications", migrations.Up011Notifications},
			{12, "sms_payment", migrations.Up012SMSPayment},
			{13, "traffic", migrations.Up013Traffic},
			{14, "chat_ai_monitoring", migrations.Up014ChatAIMonitoring},
			{15, "schema_templates", migrations.Up015SchemaTemplates},
			{16, "api_settings_ai_sessions_header_deltas", migrations.Up016ApiSettingsAndDeltas},
			{17, "password_meta_pgcrypto", migrations.Up017PasswordMetaAndPgcrypto},
			{18, "user_bank_accounts", migrations.Up018UserBankAccounts},
			{19, "user_notifications_and_logs", migrations.Up019UserNotificationsAndLogs},
			{20, "reviews_and_questions_device", migrations.Up020ReviewsAndQuestionsDevice},
			{21, "user_collections_and_address_label", migrations.Up021UserCollectionsAndAddressLabel},
			{22, "media_advanced_indexes", migrations.Up022MediaAdvancedIndexes},
			{23, "shop_settings_and_product_features", migrations.Up023ShopSettingsAndProductFeatures},
			{24, "roles_permissions_core", migrations.Up024RolesPermissionsCore},
			{25, "fraud_core_and_indexes", migrations.Up025FraudCoreAndIndexes},
			{26, "auth_optimized_core", migrations.Up026AuthOptimizedCore},
			{27, "chat_indexes", migrations.Up027ChatIndexes},
			{28, "discounts_and_lookup", migrations.Up028DiscountsAndLookup},
			{29, "header_layers_fixes", migrations.Up029HeaderLayersFixes},
			{30, "products_content_extensions_core", migrations.Up030ProductsContentExtensionsCore},
			{31, "cart_order_enhancements", migrations.Up031CartOrderEnhancements},
			{32, "media_processing_and_backups", migrations.Up032MediaProcessingAndBackups},
			{33, "admin_operational_tables", migrations.Up033AdminOperationalTables},
			{34, "widgets_and_cms_indexes", migrations.Up034WidgetsAndCMSIndexes},
			{35, "geo_and_reference_data", migrations.Up035GeoAndReferenceData},
			{36, "api_settings_updates", migrations.Up036APISettingsUpdates},
			{37, "online_chat_system", migrations.Up037OnlineChatSystem},
			{38, "chat_performance_monitoring", migrations.Up038ChatPerformanceMonitoring},
			{41, "enable_pgcrypto_extension", migrations.Up041EnablePgcrypto},
			{42, "user_roles_m2m_and_security_indexes", migrations.Up042UserRolesM2MAndSecurityIndexes},
			{43, "users_password_meta", migrations.Up043UsersPasswordMeta},
			{44, "performance_critical_indexes_v2", migrations.Up044PerformanceCriticalIndexesV2},
			{45, "media_indexes_v2", migrations.Up045MediaIndexesV2},
			{47, "sms_patterns_updated_at", migrations.Up047SMSPatternsUpdatedAt},
			{48, "sms_patterns_extend", migrations.Up048SMSPatternsExtend},
			{49, "fix_chat_messages_session_id_v2", migrations.Up049FixChatMessagesSessionIDV2},
			{50, "user_addresses_label_v2", migrations.Up050UserAddressesLabelV2},
			{51, "user_bank_accounts_v2", migrations.Up051UserBankAccountsV2},
			{53, "user_notifications_v2", migrations.Up053UserNotificationsV2},
			{54, "user_collections_v2", migrations.Up054UserCollectionsV2},
			{57, "installments_core_v2", migrations.Up057InstallmentsCoreV2},
			{59, "reviews_enhancements_v2", migrations.Up059ReviewsEnhancementsV2},
			{60, "product_questions_device_v2", migrations.Up060ProductQuestionsDeviceV2},
			{61, "fraud_core_v2", migrations.Up061FraudCoreV2},
			{62, "fraud_indexes_v2", migrations.Up062FraudIndexesV2},
			{63, "traffic_logs_and_indexes_v2", migrations.Up063TrafficLogsAndIndexesV2},
			{64, "security_login_permissions_v2", migrations.Up064SecurityLoginPermissionsV2},
			{65, "ensure_dev_user", migrations.Up065EnsureDevUser},
			{66, "settings_soft_delete_fix", migrations.Up066SettingsSoftDeleteFix},
			{67, "attributes_soft_delete_fix", migrations.Up067AttributesSoftDeleteFix},
			{68, "category_qas_core", migrations.Up068CategoryQACore},
			{69, "sms_gateway_patterns_fix", migrations.Up069SMSGatewayPatternsFix},
			{70, "product_attribute_values_core", migrations.Up070ProductAttributeValuesCore},
			{71, "seed_core_settings", migrations.Up071SeedCoreSettings},
			{72, "seo_redirects", migrations.Up072SEORedirects},
			{73, "product_variant_image_settings", migrations.Up073ProductVariantImageSettings},
			{74, "orders_shipping_fields", migrations.Up074OrdersShippingFields},
			{75, "orders_shipping_tracking_code", migrations.Up075OrdersShippingTrackingCode},
			{76, "orders_shipping_extend_courier", migrations.Up076OrdersShippingExtendCourier},
			{77, "warehouses_core", migrations.Up077WarehousesCore},
			{78, "product_reservation_flag", migrations.Up078ProductReservationFlag},
			{79, "product_default_warehouse", migrations.Up079ProductDefaultWarehouse},
			{80, "warehouse_based_inventory", migrations.Up080WarehouseBasedInventory},
			{81, "product_sale_schedule", migrations.Up081ProductSaleSchedule},
			{82, "product_special_offers", migrations.Up082ProductSpecialOffers},
			{83, "drop_products_video_url", migrations.Up083DropProductsVideoURL},
			{84, "product_videos_lazy_load", migrations.Up084ProductVideosLazyLoad},
			{85, "image_seo_jobs", migrations.Up085ImageSEOJobs},
			{86, "product_complements", migrations.Up086ProductComplements},
			{87, "empty_placeholder", migrations.Up087EmptyPlaceholder},
			{88, "seed_auth_settings", migrations.Up088SeedAuthSettings},
			{89, "header_layers_padding_fields", migrations.Up089HeaderLayersPaddingFields},
			{91, "header_layers_padding_fields_v2", migrations.Up091HeaderLayersPaddingFields},
			{102, "footer_core", migrations.Up102FooterCore},
			{104, "header_footer_permissions", migrations.Up104HeaderFooterPermissions},
			{110, "fix_otp_codes_updated_at", migrations.Up110FixOtpCodesUpdatedAt},
			{111, "seed_base_roles", migrations.Up111SeedBaseRoles},
			{112, "media_library_permissions", migrations.Up112MediaLibraryPermissions},
			{113, "default_category", migrations.Up113DefaultCategory},
			{114, "mobile_app_header_footer", migrations.Up114MobileAppHeaderFooter},
			{115, "mobile_app_navigation", migrations.Up115MobileAppNavigation},
			{116, "mobile_app_header_simplify", migrations.Up116MobileAppHeaderSimplify},
			{117, "add_image_fields_to_mobile_app_headers", migrations.AddImageFieldsToMobileAppHeaders},
			{118, "add_navigation_fields_to_mobile_app_navigation", migrations.AddNavigationFieldsToMobileAppNavigation},
			{119, "seo_redirects_permissions", migrations.Up119SEORedirectsPermissions},
			{120, "add_status_to_seo_redirects", migrations.Up120AddStatusToSEORedirects},
			{121, "add_url_to_products", migrations.Up121AddURLToProducts},
			{123, "add_parent_slug_to_categories", migrations.Up123AddParentSlugToCategories},
			{124, "payment_gateways_missing_columns", migrations.Up124PaymentGatewaysMissingColumns},
			{126, "orders_user_agent", migrations.Up126OrdersUserAgent},
			{127, "inventory_management", migrations.Up025InventoryManagement},
			{128, "fix_zero_stock_issue", migrations.Up081FixZeroStockIssue},
			{129, "add_created_at_to_product_warehouse_stocks", migrations.Up129AddCreatedAtToProductWarehouseStocks},
			{130, "recent_views_table", migrations.Migration130AddRecentViewsTable},
			{131, "add_analytics_to_recent_views", migrations.Migration131AddAnalyticsToRecentViews},
			{132, "add_device_info_to_recent_views", migrations.Migration132AddDeviceInfoToRecentViews},
			{133, "fix_recent_views_defaults", migrations.Migration133FixRecentViewsDefaults},
			{134, "user_verifications_table", migrations.Migration134AddUserVerificationsTable},
			{135, "user_banking_infos_table", migrations.Migration135AddUserBankingInfosTable},
			{136, "update_user_banking_infos_table", migrations.Migration136UpdateUserBankingInfosTable},
			{137, "add_show_on_mobile_to_widgets", migrations.Migration137AddShowOnMobileToWidgets},
			{138, "footer_layers_border_shadow", migrations.Up138FooterLayersBorderShadow},
			{139, "header_layers_style_settings", migrations.Up139HeaderLayersStyleSettings},
			{140, "header_layers_complete_fields", migrations.Up140HeaderLayersCompleteFields},
			{141, "search_index_queue", migrations.Up141SearchIndexQueue},
		} // execute in version order
		for _, m := range regs {
			log.Printf("[RUNNER] 🔍 Checking migration %d (%s)...", m.Number, m.Name)

			var cnt int64
			if err := tx.Raw("SELECT COUNT(1) FROM schema_migrations WHERE version = ?", m.Number).Scan(&cnt).Error; err != nil {
				log.Printf("[RUNNER] ❌ Error when checking migration %d: %v", m.Number, err)
				return err
			}
			if cnt > 0 {
				log.Printf("[RUNNER] ⏭️  Migration %d already applied; skipping", m.Number)
				continue // already applied
			}

			log.Printf("[MIGRATION] ▶️  Up%03d_%s starting...", m.Number, m.Name)
			log.Printf("[MIGRATION] ⏰ Start time: %v", time.Now().Format("15:04:05"))

			startTime := time.Now()
			if err := m.Up(tx); err != nil {
				log.Printf("[MIGRATION] ❌ Up%03d_%s ERROR after %v: %v", m.Number, m.Name, time.Since(startTime), err)
				return err
			}

			log.Printf("[MIGRATION] 💾 Recording migration %d in schema_migrations...", m.Number)
			if err := tx.Exec("INSERT INTO schema_migrations(version, name) VALUES (?, ?)", m.Number, m.Name).Error; err != nil {
				log.Printf("[MIGRATION] ❌ Failed to record migration %d: %v", m.Number, err)
				return err
			}

			log.Printf("[MIGRATION] ✅ Up%03d_%s completed in %v.", m.Number, m.Name, time.Since(startTime))
		}

		// DEV_SEED removed; sample data is no longer required because Up065EnsureDevUser ensures it exists
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// withAdvisoryLock اجرای تابع f با استفاده از Postgres advisory lock
func withAdvisoryLock(db *gorm.DB, key int64, f func(*gorm.DB) error) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	var locked bool
	if err := sqlDB.QueryRow("SELECT pg_try_advisory_lock($1)", key).Scan(&locked); err != nil {
		return err
	}
	if !locked {
		return fmt.Errorf("could not acquire advisory lock")
	}
	defer func(c *sql.DB) {
		_, _ = c.Exec("SELECT pg_advisory_unlock($1)", key)
	}(sqlDB)
	return f(db)
}
