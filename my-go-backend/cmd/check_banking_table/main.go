package main

import (
	"fmt"
	"my-go-backend/internal/config"
	"my-go-backend/internal/database"
)

func main() {
	// بارگذاری تنظیمات
	_ = config.Load()

	// اتصال به دیتابیس
	db, err := database.ConnectGorm()
	if err != nil {
		fmt.Printf("خطا در اتصال به دیتابیس: %v\n", err)
		return
	}

	// بررسی وجود جدول user_banking_infos
	var tableExists bool
	err = db.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'user_banking_infos')").Scan(&tableExists).Error
	if err != nil {
		fmt.Printf("خطا در بررسی وجود جدول: %v\n", err)
		return
	}

	if tableExists {
		fmt.Println("✅ جدول user_banking_infos وجود دارد")
		
		// بررسی ساختار جدول
		var columns []struct {
			ColumnName string `gorm:"column:column_name"`
			DataType   string `gorm:"column:data_type"`
			IsNullable string `gorm:"column:is_nullable"`
		}
		
		err = db.Raw(`
			SELECT column_name, data_type, is_nullable 
			FROM information_schema.columns 
			WHERE table_name = 'user_banking_infos' 
			ORDER BY ordinal_position
		`).Scan(&columns).Error
		
		if err != nil {
			fmt.Printf("خطا در دریافت ساختار جدول: %v\n", err)
			return
		}
		
		fmt.Println("📋 ساختار جدول:")
		for _, col := range columns {
			fmt.Printf("  - %s: %s (nullable: %s)\n", col.ColumnName, col.DataType, col.IsNullable)
		}
	} else {
		fmt.Println("❌ جدول user_banking_infos وجود ندارد")
		fmt.Println("🔄 اجرای migration ها...")
		
		// اجرای migration ها
		err = database.RunMigrations(db)
		if err != nil {
			fmt.Printf("خطا در اجرای migration ها: %v\n", err)
			return
		}
		
		fmt.Println("✅ Migration ها با موفقیت اجرا شدند")
	}
}
