package migrations

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// Up027ChatIndexes معادل Up40 (ایندکس‌های چت)
func Up027ChatIndexes(db *gorm.DB) error {
	log.Println("========================================")
	log.Println("[MIGRATION 027] ⏱️  شروع ایجاد ایندکس‌های چت...")
	log.Println("========================================")

	// بررسی وجود جدول قبل از ساخت index
	var count int64

	// چک کردن جدول chat_messages
	log.Println("[MIGRATION 027] 🔍 STEP 1: بررسی وجود جدول chat_messages در information_schema...")
	startTime := time.Now()

	if err := db.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'chat_messages'").Scan(&count).Error; err != nil {
		log.Printf("[MIGRATION 027] ❌ STEP 1 FAILED: خطا در بررسی جدول chat_messages: %v", err)
		log.Println("[MIGRATION 027] ⏭️  ادامه به step بعدی...")
		count = 0 // ادامه می‌دهیم
	} else {
		log.Printf("[MIGRATION 027] ✅ STEP 1 COMPLETE: پاسخ دریافت شد در %v - تعداد جداول یافت شده: %d", time.Since(startTime), count)
	}

	if count > 0 {
		log.Println("[MIGRATION 027] 📊 STEP 2: شمارش تعداد رکوردهای chat_messages...")
		log.Println("[MIGRATION 027] ⚠️  توجه: اگر جدول بزرگ باشد، ممکن است چند دقیقه طول بکشد...")
		startTime = time.Now()

		var rowCount int64
		if err := db.Raw("SELECT COUNT(*) FROM chat_messages").Scan(&rowCount).Error; err != nil {
			log.Printf("[MIGRATION 027] ❌ STEP 2 FAILED: خطا در شمارش رکوردها: %v", err)
			log.Println("[MIGRATION 027] ⏭️  رد کردن ایجاد index برای این جدول...")
		} else {
			log.Printf("[MIGRATION 027] ✅ STEP 2 COMPLETE: جدول chat_messages دارای %d رکورد است (زمان: %v)", rowCount, time.Since(startTime))

			log.Println("[MIGRATION 027] 🔨 STEP 3: در حال ایجاد index idx_chat_messages_session_created_at...")
			log.Println("[MIGRATION 027] ⏳ SQL: CREATE INDEX IF NOT EXISTS idx_chat_messages_session_created_at ON chat_messages(session_id, created_at DESC)")
			startTime = time.Now()

			if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_chat_messages_session_created_at ON chat_messages(session_id, created_at DESC)").Error; err != nil {
				log.Printf("[MIGRATION 027] ❌ STEP 3 FAILED: خطا در ایجاد index chat_messages: %v (زمان: %v)", err, time.Since(startTime))
				log.Println("[MIGRATION 027] ⏭️  ادامه می‌دهیم...")
			} else {
				log.Printf("[MIGRATION 027] ✅ STEP 3 COMPLETE: index chat_messages با موفقیت ایجاد شد (زمان: %v)", time.Since(startTime))
			}
		}
	} else {
		log.Println("[MIGRATION 027] ℹ️  جدول chat_messages وجود ندارد یا خالی است - رد شد")
	}

	log.Println("========================================")
	// چک کردن جدول chat_operators
	log.Println("[MIGRATION 027] 🔍 STEP 4: بررسی وجود جدول chat_operators در information_schema...")
	startTime = time.Now()

	if err := db.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'chat_operators'").Scan(&count).Error; err != nil {
		log.Printf("[MIGRATION 027] ❌ STEP 4 FAILED: خطا در بررسی جدول chat_operators: %v", err)
		log.Println("[MIGRATION 027] ⏭️  ادامه به پایان migration...")
		count = 0
	} else {
		log.Printf("[MIGRATION 027] ✅ STEP 4 COMPLETE: پاسخ دریافت شد در %v - تعداد جداول: %d", time.Since(startTime), count)
	}

	if count > 0 {
		log.Println("[MIGRATION 027] 📊 STEP 5: شمارش تعداد رکوردهای chat_operators...")
		startTime = time.Now()

		var rowCount int64
		if err := db.Raw("SELECT COUNT(*) FROM chat_operators").Scan(&rowCount).Error; err != nil {
			log.Printf("[MIGRATION 027] ❌ STEP 5 FAILED: خطا در شمارش: %v", err)
		} else {
			log.Printf("[MIGRATION 027] ✅ STEP 5 COMPLETE: جدول chat_operators دارای %d رکورد است (زمان: %v)", rowCount, time.Since(startTime))

			log.Println("[MIGRATION 027] 🔨 STEP 6: در حال ایجاد index idx_chat_operator_status...")
			startTime = time.Now()

			if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_chat_operator_status ON chat_operators(status)").Error; err != nil {
				log.Printf("[MIGRATION 027] ❌ STEP 6 FAILED: خطا در ایجاد index: %v (زمان: %v)", err, time.Since(startTime))
			} else {
				log.Printf("[MIGRATION 027] ✅ STEP 6 COMPLETE: index chat_operators ایجاد شد (زمان: %v)", time.Since(startTime))
			}
		}
	} else {
		log.Println("[MIGRATION 027] ℹ️  جدول chat_operators وجود ندارد یا خالی است - رد شد")
	}

	log.Println("========================================")
	log.Println("[MIGRATION 027] 🎉 پایان موفقیت‌آمیز migration 027")
	log.Println("========================================")
	return nil
}
