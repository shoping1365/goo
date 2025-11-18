package repository

import (
	"my-go-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

type GiftCardRepository struct {
	db *gorm.DB
}

func NewGiftCardRepository(db *gorm.DB) *GiftCardRepository {
	return &GiftCardRepository{db: db}
}

// Create ایجاد گیفت کارت جدید
func (r *GiftCardRepository) Create(giftCard *models.GiftCard) error {
	return r.db.Create(giftCard).Error
}

// GetByID دریافت گیفت کارت بر اساس ID
func (r *GiftCardRepository) GetByID(id uint) (*models.GiftCard, error) {
	var giftCard models.GiftCard
	err := r.db.Preload("User").Preload("Transactions").First(&giftCard, id).Error
	return &giftCard, err
}

// GetByCode دریافت گیفت کارت بر اساس کد
func (r *GiftCardRepository) GetByCode(code string) (*models.GiftCard, error) {
	var giftCard models.GiftCard
	err := r.db.Preload("User").Preload("Transactions").Where("code = ?", code).First(&giftCard).Error
	return &giftCard, err
}

// GetAll دریافت تمام گیفت کارت‌ها با فیلتر
func (r *GiftCardRepository) GetAll(limit, offset int, status, cardType, category string) ([]models.GiftCard, int64, error) {
	var giftCards []models.GiftCard
	var total int64

	query := r.db.Model(&models.GiftCard{}).Preload("User")

	// اعمال فیلترها
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if cardType != "" {
		query = query.Where("type = ?", cardType)
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// شمارش کل رکوردها
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// دریافت رکوردها
	err = query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&giftCards).Error
	return giftCards, total, err
}

// Update به‌روزرسانی گیفت کارت
func (r *GiftCardRepository) Update(giftCard *models.GiftCard) error {
	return r.db.Save(giftCard).Error
}

// Delete حذف گیفت کارت
func (r *GiftCardRepository) Delete(id uint) error {
	return r.db.Delete(&models.GiftCard{}, id).Error
}

// GetStats دریافت آمار کلی (محاسبات real-time)
func (r *GiftCardRepository) GetStats() (*models.GiftCardStats, error) {
	var stats models.GiftCardStats

	// کل کارت‌ها
	err := r.db.Model(&models.GiftCard{}).Count(&stats.TotalCards).Error
	if err != nil {
		return nil, err
	}

	// کارت‌های فعال
	err = r.db.Model(&models.GiftCard{}).Where("status = ?", "active").Count(&stats.ActiveCards).Error
	if err != nil {
		return nil, err
	}

	// کارت‌های در حال انقضا (7 روز آینده)
	expiryDate := time.Now().AddDate(0, 0, 7)
	err = r.db.Model(&models.GiftCard{}).Where("expiry_date <= ? AND status = ?", expiryDate, "active").Count(&stats.ExpiringSoon).Error
	if err != nil {
		return nil, err
	}

	// ارزش کل
	err = r.db.Model(&models.GiftCard{}).Where("status = ?", "active").Select("COALESCE(SUM(remaining_amount), 0)").Scan(&stats.TotalValue).Error
	if err != nil {
		return nil, err
	}

	// کارت‌های جدید امروز
	today := time.Now().Truncate(24 * time.Hour)
	err = r.db.Model(&models.GiftCard{}).Where("created_at >= ?", today).Count(&stats.NewCardsToday).Error
	if err != nil {
		return nil, err
	}

	// محاسبه درصد کارت‌های فعال
	if stats.TotalCards > 0 {
		stats.ActivePercentage = float64(stats.ActiveCards) / float64(stats.TotalCards) * 100
	}

	// کارت‌های استفاده شده
	err = r.db.Model(&models.GiftCard{}).Where("status = ?", "used").Count(&stats.UsedCards).Error
	if err != nil {
		return nil, err
	}

	// محاسبه نرخ استفاده
	if stats.TotalCards > 0 {
		stats.UsageRate = float64(stats.UsedCards) / float64(stats.TotalCards) * 100
	}

	// میانگین مبلغ استفاده
	err = r.db.Model(&models.GiftCard{}).Where("status = ?", "used").Select("COALESCE(AVG(amount - remaining_amount), 0)").Scan(&stats.AvgUsageAmount).Error
	if err != nil {
		return nil, err
	}

	// کارت‌های تحویل شده
	err = r.db.Model(&models.GiftCard{}).Where("delivery_method != ? AND delivery_date IS NOT NULL", "manual").Count(&stats.DeliveredCards).Error
	if err != nil {
		return nil, err
	}

	// محاسبه نرخ تحویل
	if stats.TotalCards > 0 {
		stats.DeliveryRate = float64(stats.DeliveredCards) / float64(stats.TotalCards) * 100
	}

	// آمار تحویل ایمیل و پیامک (نمونه - نیاز به محاسبه دقیق‌تر)
	stats.EmailDelivery = 60.0 // درصد نمونه
	stats.SmsDelivery = 40.0   // درصد نمونه

	// رشد ارزش (نمونه - نیاز به محاسبه با داده‌های تاریخی)
	stats.ValueGrowth = 12.5 // درصد نمونه

	return &stats, nil
}

// GetRecentActivities دریافت فعالیت‌های اخیر
func (r *GiftCardRepository) GetRecentActivities(limit int) ([]models.GiftCardActivity, error) {
	var activities []models.GiftCardActivity
	err := r.db.Preload("GiftCard").Preload("User").Order("created_at DESC").Limit(limit).Find(&activities).Error
	return activities, err
}

// GetChartData دریافت داده‌های نمودار
func (r *GiftCardRepository) GetChartData() (map[string]interface{}, error) {
	data := make(map[string]interface{})

	// داده‌های وضعیت کارت‌ها
	var statusData []map[string]interface{}
	err := r.db.Model(&models.GiftCard{}).
		Select("status, COUNT(*) as count").
		Group("status").
		Scan(&statusData).Error
	if err != nil {
		return nil, err
	}
	data["statuses"] = statusData

	// داده‌های روند ایجاد کارت‌ها (30 روز گذشته)
	var trendData []map[string]interface{}
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	err = r.db.Model(&models.GiftCard{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("created_at >= ?", thirtyDaysAgo).
		Group("DATE(created_at)").
		Order("date").
		Scan(&trendData).Error
	if err != nil {
		return nil, err
	}
	data["trend"] = trendData

	// داده‌های توزیع مبلغ
	var amountData []map[string]interface{}
	err = r.db.Model(&models.GiftCard{}).
		Select("CASE " +
			"WHEN amount < 100000 THEN 'کمتر از ۱۰۰ هزار' " +
			"WHEN amount < 500000 THEN '۱۰۰ تا ۵۰۰ هزار' " +
			"WHEN amount < 1000000 THEN '۵۰۰ هزار تا ۱ میلیون' " +
			"ELSE 'بیش از ۱ میلیون' " +
			"END as range, COUNT(*) as count").
		Group("range").
		Scan(&amountData).Error
	if err != nil {
		return nil, err
	}
	data["amounts"] = amountData

	return data, nil
}

// CreateTransaction ایجاد تراکنش جدید
func (r *GiftCardRepository) CreateTransaction(transaction *models.GiftCardTransaction) error {
	return r.db.Create(transaction).Error
}

// GetTransactions دریافت تراکنش‌های گیفت کارت
func (r *GiftCardRepository) GetTransactions(giftCardID uint, limit, offset int) ([]models.GiftCardTransaction, int64, error) {
	var transactions []models.GiftCardTransaction
	var total int64

	query := r.db.Model(&models.GiftCardTransaction{}).Where("gift_card_id = ?", giftCardID)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Preload("User").Limit(limit).Offset(offset).Order("created_at DESC").Find(&transactions).Error
	return transactions, total, err
}

// CreateActivity ایجاد فعالیت جدید
func (r *GiftCardRepository) CreateActivity(activity *models.GiftCardActivity) error {
	return r.db.Create(activity).Error
}

// Search جستجو در گیفت کارت‌ها
func (r *GiftCardRepository) Search(query string, limit, offset int) ([]models.GiftCard, int64, error) {
	var giftCards []models.GiftCard
	var total int64

	searchQuery := r.db.Model(&models.GiftCard{}).Preload("User").
		Where("code LIKE ? OR recipient_name LIKE ? OR recipient_email LIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%")

	err := searchQuery.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = searchQuery.Limit(limit).Offset(offset).Order("created_at DESC").Find(&giftCards).Error
	return giftCards, total, err
}
