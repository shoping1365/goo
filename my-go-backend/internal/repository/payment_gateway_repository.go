package repository

import (
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// PaymentGatewayRepository repository برای مدیریت درگاه‌های پرداخت
// شامل تمام عملیات CRUD و جستجوهای پیشرفته

type PaymentGatewayRepository struct {
	db *gorm.DB
}

// NewPaymentGatewayRepository سازنده repository
func NewPaymentGatewayRepository(db *gorm.DB) *PaymentGatewayRepository {
	return &PaymentGatewayRepository{db: db}
}

// Create ثبت درگاه پرداخت جدید
func (r *PaymentGatewayRepository) Create(gateway *models.PaymentGateway) error {
	return r.db.Create(gateway).Error
}

// GetByID دریافت درگاه پرداخت بر اساس ID
func (r *PaymentGatewayRepository) GetByID(id uint) (*models.PaymentGateway, error) {
	var gateway models.PaymentGateway
	err := r.db.First(&gateway, id).Error
	if err != nil {
		return nil, err
	}
	return &gateway, nil
}

// GetAll دریافت تمام درگاه‌های پرداخت
func (r *PaymentGatewayRepository) GetAll() ([]models.PaymentGateway, error) {
	var gateways []models.PaymentGateway
	err := r.db.Order("priority ASC, created_at DESC").Find(&gateways).Error
	return gateways, err
}

// GetActive دریافت درگاه‌های فعال
func (r *PaymentGatewayRepository) GetActive() ([]models.PaymentGateway, error) {
	var gateways []models.PaymentGateway
	err := r.db.Where("status = ?", "active").Order("priority ASC").Find(&gateways).Error
	return gateways, err
}

// GetByType دریافت درگاه‌ها بر اساس نوع
func (r *PaymentGatewayRepository) GetByType(gatewayType string) ([]models.PaymentGateway, error) {
	var gateways []models.PaymentGateway
	err := r.db.Where("type = ?", gatewayType).Order("priority ASC").Find(&gateways).Error
	return gateways, err
}

// GetFirstByType دریافت اولین درگاه بر اساس نوع
func (r *PaymentGatewayRepository) GetFirstByType(gatewayType string) (*models.PaymentGateway, error) {
	var gateway models.PaymentGateway
	err := r.db.Where("type = ?", gatewayType).First(&gateway).Error
	if err != nil {
		return nil, err
	}
	return &gateway, nil
}

// GetByStatus دریافت درگاه‌ها بر اساس وضعیت
func (r *PaymentGatewayRepository) GetByStatus(status string) ([]models.PaymentGateway, error) {
	var gateways []models.PaymentGateway
	err := r.db.Where("status = ?", status).Order("priority ASC").Find(&gateways).Error
	return gateways, err
}

// Search جستجو در درگاه‌های پرداخت
func (r *PaymentGatewayRepository) Search(query string) ([]models.PaymentGateway, error) {
	var gateways []models.PaymentGateway
	err := r.db.Where("name LIKE ? OR english_name LIKE ?", "%"+query+"%", "%"+query+"%").Order("priority ASC").Find(&gateways).Error
	return gateways, err
}

// Update به‌روزرسانی درگاه پرداخت
func (r *PaymentGatewayRepository) Update(gateway *models.PaymentGateway) error {
	return r.db.Save(gateway).Error
}

// Delete حذف درگاه پرداخت
func (r *PaymentGatewayRepository) Delete(id uint) error {
	return r.db.Delete(&models.PaymentGateway{}, id).Error
}

// UpdateStatus تغییر وضعیت درگاه پرداخت
func (r *PaymentGatewayRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.PaymentGateway{}).Where("id = ?", id).Update("status", status).Error
}

// UpdatePriority تغییر اولویت درگاه پرداخت
func (r *PaymentGatewayRepository) UpdatePriority(id uint, priority int) error {
	return r.db.Model(&models.PaymentGateway{}).Where("id = ?", id).Update("priority", priority).Error
}

// UpdateTestMode تغییر حالت تست درگاه پرداخت
func (r *PaymentGatewayRepository) UpdateTestMode(id uint, isTestMode bool) error {
	return r.db.Model(&models.PaymentGateway{}).Where("id = ?", id).Update("is_test_mode", isTestMode).Error
}

// GetStats دریافت آمار کلی درگاه‌های پرداخت
func (r *PaymentGatewayRepository) GetStats() (map[string]interface{}, error) {
	var stats struct {
		TotalGateways       int64
		ActiveGateways      int64
		InactiveGateways    int64
		MaintenanceGateways int64
		TodayTransactions   int64
		TodayRevenue        int64
	}

	// تعداد کل درگاه‌ها
	r.db.Model(&models.PaymentGateway{}).Count(&stats.TotalGateways)

	// تعداد درگاه‌های فعال
	r.db.Model(&models.PaymentGateway{}).Where("status = ?", "active").Count(&stats.ActiveGateways)

	// تعداد درگاه‌های غیرفعال
	r.db.Model(&models.PaymentGateway{}).Where("status = ?", "inactive").Count(&stats.InactiveGateways)

	// تعداد درگاه‌های در حال نگهداری
	r.db.Model(&models.PaymentGateway{}).Where("status = ?", "maintenance").Count(&stats.MaintenanceGateways)

	// مجموع تراکنش‌های امروز
	r.db.Model(&models.PaymentGateway{}).Select("COALESCE(SUM(today_transactions), 0)").Scan(&stats.TodayTransactions)

	// مجموع درآمد امروز
	r.db.Model(&models.PaymentGateway{}).Select("COALESCE(SUM(today_revenue), 0)").Scan(&stats.TodayRevenue)

	result := map[string]interface{}{
		"total_gateways":       stats.TotalGateways,
		"active_gateways":      stats.ActiveGateways,
		"inactive_gateways":    stats.InactiveGateways,
		"maintenance_gateways": stats.MaintenanceGateways,
		"today_transactions":   stats.TodayTransactions,
		"today_revenue":        stats.TodayRevenue,
	}

	return result, nil
}

// GetTopGateways دریافت درگاه‌های برتر بر اساس تراکنش‌های امروز
func (r *PaymentGatewayRepository) GetTopGateways(limit int) ([]models.PaymentGateway, error) {
	var gateways []models.PaymentGateway
	err := r.db.Where("today_transactions > 0").Order("today_transactions DESC").Limit(limit).Find(&gateways).Error
	return gateways, err
}

// UpdateTodayStats به‌روزرسانی آمار امروز درگاه
func (r *PaymentGatewayRepository) UpdateTodayStats(id uint, transactions int64, revenue int64) error {
	return r.db.Model(&models.PaymentGateway{}).Where("id = ?", id).Updates(map[string]interface{}{
		"today_transactions": transactions,
		"today_revenue":      revenue,
	}).Error
}

// ResetTodayStats ریست کردن آمار امروز تمام درگاه‌ها
func (r *PaymentGatewayRepository) ResetTodayStats() error {
	return r.db.Model(&models.PaymentGateway{}).Updates(map[string]interface{}{
		"today_transactions": 0,
		"today_revenue":      0,
	}).Error
}

// GetByFilters دریافت درگاه‌ها با فیلترهای پیشرفته
func (r *PaymentGatewayRepository) GetByFilters(filters map[string]interface{}) ([]models.PaymentGateway, error) {
	query := r.db

	// فیلتر بر اساس نوع
	if gatewayType, exists := filters["type"]; exists && gatewayType != "" {
		query = query.Where("type = ?", gatewayType)
	}

	// فیلتر بر اساس وضعیت
	if status, exists := filters["status"]; exists && status != "" {
		query = query.Where("status = ?", status)
	}

	// فیلتر بر اساس حالت تست
	if isTestMode, exists := filters["is_test_mode"]; exists {
		query = query.Where("is_test_mode = ?", isTestMode)
	}

	// جستجو در نام
	if search, exists := filters["search"]; exists && search != "" {
		searchStr := "%" + search.(string) + "%"
		query = query.Where("name LIKE ? OR english_name LIKE ?", searchStr, searchStr)
	}

	// مرتب‌سازی
	if orderBy, exists := filters["order_by"]; exists {
		query = query.Order(orderBy.(string))
	} else {
		query = query.Order("priority ASC, created_at DESC")
	}

	// محدودیت تعداد نتایج
	if limit, exists := filters["limit"]; exists {
		query = query.Limit(limit.(int))
	}

	var gateways []models.PaymentGateway
	err := query.Find(&gateways).Error
	return gateways, err
}

// Count تعداد کل درگاه‌ها
func (r *PaymentGatewayRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.PaymentGateway{}).Count(&count).Error
	return count, err
}

// Exists بررسی وجود درگاه بر اساس نام انگلیسی
func (r *PaymentGatewayRepository) Exists(englishName string) (bool, error) {
	var count int64
	err := r.db.Model(&models.PaymentGateway{}).Where("english_name = ?", englishName).Count(&count).Error
	return count > 0, err
}

// GetNextPriority دریافت اولویت بعدی برای درگاه جدید
func (r *PaymentGatewayRepository) GetNextPriority() (int, error) {
	var maxPriority int
	err := r.db.Model(&models.PaymentGateway{}).Select("COALESCE(MAX(priority), 0)").Scan(&maxPriority).Error
	return maxPriority + 1, err
}
