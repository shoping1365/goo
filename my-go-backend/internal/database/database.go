package database

import (
	"fmt"
	"log"
	"os"

	"my-go-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

// ConnectGorm اتصال به دیتابیس با GORM و مقداردهی GormDB
func ConnectGorm() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL env not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		return nil, err
	}
	GormDB = db
	return db, nil
}

// CloseGorm بستن اتصال دیتابیس (در صورت نیاز)
func CloseGorm() error {
	if GormDB == nil {
		return nil
	}
	sqlDB, err := GormDB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// MigrateAll اجرای همه مایگریشن‌های اصلی پروژه
func MigrateAll(db *gorm.DB) error {
	if os.Getenv("AUTO_MIGRATE_DEV") != "true" {
		log.Println("*************************************************************************************************")
		log.Println("⚠️  WARNING: 'MigrateAll' (GORM AutoMigrate) is DEPRECATED and should only be used in DEV.")
		log.Println("⚠️  Use explicit, versioned migrations under internal/database/migrations and the migrate scripts.")
		log.Println("⚠️  To silence this warning temporarily, set AUTO_MIGRATE_DEV=true (not recommended for CI/Prod).")
		log.Println("*************************************************************************************************")
	}
	return db.AutoMigrate(
		// مدل تنظیمات سیستم (برای ذخیره تنظیمات کل پروژه)
		&models.Setting{},

		// مدل کاربر (اطلاعات کاربران سامانه)
		&models.User{},
		// مدل آدرس‌های کاربر (آدرس‌های چندگانه برای هر کاربر)
		&models.UserAddress{},

		// مدل محصول (اطلاعات محصولات فروشگاه)
		&models.Product{},
		// مدل تصویر محصول (تصاویر مرتبط با هر محصول)
		&models.ProductImage{},
		// مدل تنوع محصول (رنگ، سایز و ... برای هر محصول)
		&models.ProductVariant{},
		// مدل مشخصات محصول (ویژگی‌های فنی و توضیحات تکمیلی)
		&models.ProductSpecification{},
		// مدل ویدیوی محصول (ویدیوهای مرتبط با هر محصول)
		&models.ProductVideo{},

		// مدل دسته‌بندی Q&A (دسته‌بندی سوالات و پاسخ‌ها)
		&models.CategoryQA{},

		// مدل دسته‌بندی محصولات (دسته‌بندی محصولات فروشگاه)
		&models.Category{},
		&models.CategoryBrandPage{},

		// مدل رسانه (فایل‌های آپلود شده: عکس، ویدیو، ...)
		&models.MediaFile{},
		// مدل واریانت رسانه (سایزهای مختلف یک رسانه)
		&models.MediaVariant{},

		// مدل مشتری (اطلاعات مشتریان و ارتباط با نظرات)
		&models.Customer{},
		// مدل نظر (نظرات کاربران روی محصولات)
		&models.Review{},
		// مدل تصویر نظر (تصاویر ضمیمه شده به هر نظر)
		&models.ReviewImage{},

		// Backup & audit tables
		&models.BackupEntry{},
		&models.MediaDeleteLog{},

		// مدل نوشته‌ها و دسته‌بندی‌های نوشته‌ها
		&models.PostCategory{},
		&models.Post{},

		// مدل تنظیمات API (برای مدیریت API های خارجی مثل OpenAI)
		&models.APISettings{},

		// مدل‌های احراز هویت یکپارچه و بهینه
		&models.OTPCode{},
		&models.AuthEvent{},
		&models.UserSession{},

		// مدل‌های ترافیک و مانیتورینگ
		&models.TrafficSession{},
		&models.BlockedIP{},

		// مدل‌های نقش و مجوز (RBAC System)
		&models.Role{},
		&models.Permission{},
		&models.RolePermission{},

		// مدل‌های چت و پشتیبانی (Chat System) - جداول توسط مایگریشن‌های دستی ایجاد می‌شوند
		// مایگریشن‌های 037_online_chat_system و 038_chat_performance_monitoring جداول چت را ایجاد می‌کنند
		// AutoMigrate برای چت غیرفعال شده تا از تداخل با مایگریشن‌های دستی جلوگیری شود
		// &models.ChatOperator{},
		// &models.ChatWidget{},
		// &models.ChatAIBot{},
		// &models.ChatKnowledgeBase{},
		// &models.ChatStatistics{},
		// &models.ChatSession{},
		// &models.OnlineChatMessage{},
	)
}
