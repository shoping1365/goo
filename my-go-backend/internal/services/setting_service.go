package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
)

/*
سرویس تنظیمات

این سرویس برای مدیریت تنظیمات سیستم استفاده می‌شود و شامل قابلیت‌های زیر است:
- دریافت تنظیمات
- بروزرسانی تنظیمات
- کش کردن تنظیمات در Redis
- مدیریت دسته‌بندی‌های مختلف تنظیمات
*/

type SettingService struct {
	db    *gorm.DB
	redis *redis.Client
}

// NewSettingService یک نمونه جدید از سرویس تنظیمات ایجاد می‌کند
func NewSettingService(db *gorm.DB, redis *redis.Client) *SettingService {
	return &SettingService{
		db:    db,
		redis: redis,
	}
}

// GetSetting تنظیم مورد نظر را بر اساس کلید دریافت می‌کند
func (s *SettingService) GetSetting(ctx context.Context, key string) (*models.Setting, error) {
	// ابتدا از کش بررسی می‌کنیم
	cacheKey := fmt.Sprintf("setting:%s", key)
	cachedSetting, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var setting models.Setting
		if err := json.Unmarshal([]byte(cachedSetting), &setting); err == nil {
			return &setting, nil
		}
	}

	// اگر در کش نبود، از دیتابیس می‌خوانیم
	var setting models.Setting
	if err := s.db.Where("key = ?", key).First(&setting).Error; err != nil {
		return nil, err
	}

	// در کش ذخیره می‌کنیم
	if settingJSON, err := json.Marshal(setting); err == nil {
		s.redis.Set(ctx, cacheKey, settingJSON, 24*time.Hour)
	}

	return &setting, nil
}

// IsTrafficLoggingEnabled بررسی می‌کند که آیا لاگ‌برداری ترافیک فعال است یا خیر.
// منبع حقیقت: جدول settings با کلید "traffic.logging.enabled" (true/false یا 1/0)
// در صورت نبود کلید، مقدار پیش‌فرض true است تا قابلیت فعال باشد مگر اینکه ادمین خاموش کند.
func IsTrafficLoggingEnabled(ctx context.Context) bool {
	// تلاش سریع: اگر gin.Context در کانتکست وجود دارد، DB را از آن بگیریم
	if gc, ok := ctx.Value("gin_context").(interface {
		Get(string) (interface{}, bool)
	}); ok {
		if v, exists := gc.Get("db"); exists {
			if db, ok2 := v.(*gorm.DB); ok2 && db != nil {
				var st models.Setting
				if err := db.WithContext(ctx).Where("key = ?", "traffic.logging.enabled").First(&st).Error; err == nil {
					val := strings.ToLower(strings.TrimSpace(st.Value))
					return val == "true" || val == "1" || val == "yes"
				}
				// در صورت خطا یا نبود مقدار، پیش‌فرض: فعال
				return true
			}
		}
	}
	//Fallback: فعال
	return true
}

// IsTrafficLoggingEnabledWithDB نسخه‌ای که مستقیماً از *gorm.DB استفاده می‌کند
func IsTrafficLoggingEnabledWithDB(ctx context.Context, db *gorm.DB) bool {
	if db != nil {
		var st models.Setting
		if err := db.WithContext(ctx).Where("key = ?", "traffic.logging.enabled").First(&st).Error; err == nil {
			val := strings.ToLower(strings.TrimSpace(st.Value))
			return val == "true" || val == "1" || val == "yes"
		}
		return true
	}
	return true
}

// GetSettingsByKeyPrefix تنظیمات با پیشوند کلید مشخص را دریافت می‌کند
func (s *SettingService) GetSettingsByKeyPrefix(ctx context.Context, prefix string) ([]models.Setting, error) {
	var settings []models.Setting
	if err := s.db.WithContext(ctx).Where("key LIKE ?", prefix+"%").Find(&settings).Error; err != nil {
		return nil, err
	}
	return settings, nil
}

// GetSettingsByCategory تنظیمات یک دسته‌بندی خاص را دریافت می‌کند
func (s *SettingService) GetSettingsByCategory(ctx context.Context, category string) ([]models.Setting, error) {
	// ابتدا از کش بررسی می‌کنیم
	cacheKey := fmt.Sprintf("settings:category:%s", category)
	cachedSettings, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var settings []models.Setting
		if err := json.Unmarshal([]byte(cachedSettings), &settings); err == nil {
			return settings, nil
		}
	}

	// اگر در کش نبود، از دیتابیس می‌خوانیم
	var settings []models.Setting
	if err := s.db.Where("category = ?", category).Find(&settings).Error; err != nil {
		return nil, err
	}

	// در کش ذخیره می‌کنیم
	if settingsJSON, err := json.Marshal(settings); err == nil {
		s.redis.Set(ctx, cacheKey, settingsJSON, 24*time.Hour)
	}

	return settings, nil
}

// UpdateSetting تنظیم مورد نظر را بروزرسانی می‌کند (با پشتیبانی UPSERT)
func (s *SettingService) UpdateSetting(ctx context.Context, setting *models.Setting) error {
	// جستجو برای تنظیم موجود
	var existing models.Setting
	err := s.db.WithContext(ctx).Where("key = ?", setting.Key).First(&existing).Error

	if err == nil {
		// رکورد موجود است، فقط بروزرسانی کن
		setting.ID = existing.ID
		setting.CreatedAt = existing.CreatedAt
		setting.CreatedBy = existing.CreatedBy
		setting.UpdatedAt = time.Now()

		if err := s.db.WithContext(ctx).Omit("DeletedAt").Save(setting).Error; err != nil {
			return err
		}
	} else if err == gorm.ErrRecordNotFound {
		// رکورد جدید است، ایجاد کن
		setting.CreatedAt = time.Now()
		setting.UpdatedAt = time.Now()
		if err := s.db.WithContext(ctx).Create(setting).Error; err != nil {
			return err
		}
	} else {
		// خطای دیگری رخ داده
		return err
	}

	// حذف از کش
	cacheKey := fmt.Sprintf("setting:%s", setting.Key)
	s.redis.Del(ctx, cacheKey)

	// حذف کش دسته‌بندی
	categoryCacheKey := fmt.Sprintf("settings:category:%s", setting.Category)
	s.redis.Del(ctx, categoryCacheKey)

	// پاک کردن کش‌های کلی و عمومی تا تغییرات بلافاصله قابل مشاهده باشند
	s.redis.Del(ctx, "settings:all")
	s.redis.Del(ctx, "settings:public")

	return nil
}

// UpdateSettings بروزرسانی چندین تنظیم به صورت یکجا با استفاده از تراکنش
func (s *SettingService) UpdateSettings(ctx context.Context, settings []models.Setting) error {
	// شروع تراکنش
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// جمع‌آوری کلیدهای تنظیمات
	var keys []string
	for _, setting := range settings {
		keys = append(keys, setting.Key)
	}

	// دریافت همه تنظیمات موجود در یک query
	var existingSettings []models.Setting
	if err := tx.Where("key IN ?", keys).Find(&existingSettings).Error; err != nil {
		tx.Rollback()
		return err
	}

	// ایجاد map برای دسترسی سریع
	existingMap := make(map[string]models.Setting)
	for _, setting := range existingSettings {
		existingMap[setting.Key] = setting
	}

	// بروزرسانی تنظیمات به صورت batch
	for _, setting := range settings {
		if existingSetting, exists := existingMap[setting.Key]; exists {
			// اگر تنظیم وجود دارد، آن را بروزرسانی کن
			setting.ID = existingSetting.ID
			setting.CreatedAt = existingSetting.CreatedAt
			setting.CreatedBy = existingSetting.CreatedBy
			setting.UpdatedAt = time.Now()

			if err := tx.Omit("DeletedAt").Save(&setting).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			// اگر تنظیم وجود ندارد، آن را ایجاد کن
			setting.CreatedAt = time.Now()
			setting.UpdatedAt = time.Now()
			if err := tx.Omit("DeletedAt").Create(&setting).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// تایید تراکنش
	if err := tx.Commit().Error; err != nil {
		return err
	}

	// پاک کردن کش های مرتبط تا پس از رفرش مقدار جدید بارگذاری شود
	if s.redis != nil {
		// حذف کش تک‌کلیدها
		var delKeys []string
		for _, setting := range settings {
			delKeys = append(delKeys, fmt.Sprintf("setting:%s", setting.Key))
		}

		// حذف کش دسته‌بندی‌ها
		categorySet := map[string]struct{}{}
		for _, setting := range settings {
			if setting.Category != "" {
				categorySet[setting.Category] = struct{}{}
			}
		}
		for cat := range categorySet {
			delKeys = append(delKeys, fmt.Sprintf("settings:category:%s", cat))
		}

		// حذف کش‌های کلی
		delKeys = append(delKeys, "settings:all", "settings:public", "settings")
		s.redis.Del(ctx, delKeys...)
	}

	return nil
}

// GetAllSettings تمام تنظیمات را دریافت می‌کند
func (s *SettingService) GetAllSettings(ctx context.Context) ([]models.Setting, error) {
	// ابتدا از کش بررسی می‌کنیم
	cacheKey := "settings:all"
	cachedSettings, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var settings []models.Setting
		if err := json.Unmarshal([]byte(cachedSettings), &settings); err == nil {
			return settings, nil
		}
	}

	// اگر در کش نبود، از دیتابیس می‌خوانیم
	var settings []models.Setting
	if err := s.db.Find(&settings).Error; err != nil {
		return nil, err
	}

	// در کش ذخیره می‌کنیم
	if settingsJSON, err := json.Marshal(settings); err == nil {
		s.redis.Set(ctx, cacheKey, settingsJSON, 24*time.Hour)
	}

	return settings, nil
}

// GetPublicSettings تنظیمات عمومی را دریافت می‌کند
func (s *SettingService) GetPublicSettings(ctx context.Context) ([]models.Setting, error) {
	// ابتدا از کش بررسی می‌کنیم
	cacheKey := "settings:public"
	cachedSettings, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var settings []models.Setting
		if err := json.Unmarshal([]byte(cachedSettings), &settings); err == nil {
			return settings, nil
		}
	}

	// اگر در کش نبود، از دیتابیس می‌خوانیم
	var settings []models.Setting
	if err := s.db.Where("is_public = ?", true).Find(&settings).Error; err != nil {
		return nil, err
	}

	// در کش ذخیره می‌کنیم
	if settingsJSON, err := json.Marshal(settings); err == nil {
		s.redis.Set(ctx, cacheKey, settingsJSON, 24*time.Hour)
	}

	return settings, nil
}

// ClearCache کش تنظیمات را پاک می‌کند
func (s *SettingService) ClearCache(ctx context.Context) error {
	// حذف همه کلیدهای مربوط به تنظیمات
	pattern := "setting:*"
	iter := s.redis.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		s.redis.Del(ctx, iter.Val())
	}
	return iter.Err()
}

// GetVideoCompressionStartHour ساعت شروع فشرده‌سازی ویدیو را برمی‌گرداند
func (s *SettingService) GetVideoCompressionStartHour(ctx context.Context) (int, error) {
	setting, err := s.GetSetting(ctx, "video_compression.start_hour")
	if err != nil {
		return 1, err // مقدار پیش‌فرض
	}

	var hour int
	if _, err := fmt.Sscanf(setting.Value, "%d", &hour); err != nil {
		return 1, err
	}

	// بررسی محدوده مجاز
	if hour < 0 || hour > 23 {
		return 1, fmt.Errorf("invalid hour: %d", hour)
	}

	return hour, nil
}

// GetVideoCompressionEndHour ساعت پایان فشرده‌سازی ویدیو را برمی‌گرداند
func (s *SettingService) GetVideoCompressionEndHour(ctx context.Context) (int, error) {
	setting, err := s.GetSetting(ctx, "video_compression.end_hour")
	if err != nil {
		return 13, err // مقدار پیش‌فرض
	}

	var hour int
	if _, err := fmt.Sscanf(setting.Value, "%d", &hour); err != nil {
		return 13, err
	}

	// بررسی محدوده مجاز
	if hour < 0 || hour > 23 {
		return 13, fmt.Errorf("invalid hour: %d", hour)
	}

	return hour, nil
}

// GetVideoCompressionEnabled وضعیت فعال بودن فشرده‌سازی ویدیو را برمی‌گرداند
func (s *SettingService) GetVideoCompressionEnabled(ctx context.Context) (bool, error) {
	setting, err := s.GetSetting(ctx, "video_compression.enabled")
	if err != nil {
		return false, err // مقدار پیش‌فرض - غیرفعال
	}

	// تبدیل به boolean
	enabled := setting.Value == "1" || strings.ToLower(setting.Value) == "true"
	return enabled, nil
}

// دسترسی مستقیم به *gorm.DB برای عملیات CRUD هدرها
func (s *SettingService) DB() *gorm.DB {
	return s.db
}
