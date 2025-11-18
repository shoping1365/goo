package repository

import (
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// PaymentTransactionRepository repository برای مدیریت تراکنش‌های پرداخت
type PaymentTransactionRepository struct {
	db *gorm.DB
}

// NewPaymentTransactionRepository سازنده repository تراکنش‌های پرداخت
func NewPaymentTransactionRepository(db *gorm.DB) *PaymentTransactionRepository {
	return &PaymentTransactionRepository{
		db: db,
	}
}

// Create ایجاد تراکنش جدید
func (r *PaymentTransactionRepository) Create(transaction *models.PaymentTransaction) error {
	return r.db.Create(transaction).Error
}

// GetByID دریافت تراکنش بر اساس ID
func (r *PaymentTransactionRepository) GetByID(id uint) (*models.PaymentTransaction, error) {
	var transaction models.PaymentTransaction
	err := r.db.Preload("Gateway").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// GetByTransactionID دریافت تراکنش بر اساس TransactionID
func (r *PaymentTransactionRepository) GetByTransactionID(transactionID string) (*models.PaymentTransaction, error) {
	var transaction models.PaymentTransaction
	err := r.db.Preload("Gateway").Where("transaction_id = ?", transactionID).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// DB بازگرداندن اتصال GORM برای استفاده داخلی (موارد ضروری)
func (r *PaymentTransactionRepository) DB() *gorm.DB { return r.db }

// GetByOrderID دریافت تراکنش بر اساس OrderID
func (r *PaymentTransactionRepository) GetByOrderID(orderID string) (*models.PaymentTransaction, error) {
	var transaction models.PaymentTransaction
	err := r.db.Preload("Gateway").Where("order_id = ?", orderID).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// Update بروزرسانی تراکنش
func (r *PaymentTransactionRepository) Update(transaction *models.PaymentTransaction) error {
	return r.db.Save(transaction).Error
}

// Delete حذف تراکنش
func (r *PaymentTransactionRepository) Delete(id uint) error {
	return r.db.Delete(&models.PaymentTransaction{}, id).Error
}

// GetPaginated دریافت تراکنش‌ها با صفحه‌بندی
func (r *PaymentTransactionRepository) GetPaginated(page, limit int, status string, gatewayID uint) ([]models.PaymentTransaction, int64, error) {
	var transactions []models.PaymentTransaction
	var total int64

	query := r.db.Model(&models.PaymentTransaction{}).Preload("Gateway")

	// اعمال فیلترها
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if gatewayID > 0 {
		query = query.Where("gateway_id = ?", gatewayID)
	}

	// شمارش کل رکوردها
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// دریافت رکوردها با صفحه‌بندی
	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&transactions).Error
	if err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}

// GetByGateway دریافت تراکنش‌های یک درگاه خاص
func (r *PaymentTransactionRepository) GetByGateway(gatewayID uint, page, limit int) ([]models.PaymentTransaction, int64, error) {
	return r.GetPaginated(page, limit, "", gatewayID)
}

// GetByStatus دریافت تراکنش‌ها بر اساس وضعیت
func (r *PaymentTransactionRepository) GetByStatus(status string, page, limit int) ([]models.PaymentTransaction, int64, error) {
	return r.GetPaginated(page, limit, status, 0)
}

// GetStatistics دریافت آمار تراکنش‌ها
func (r *PaymentTransactionRepository) GetStatistics(gatewayID uint) (map[string]interface{}, error) {
	var stats struct {
		TotalTransactions int64 `json:"total_transactions"`
		TotalAmount       int64 `json:"total_amount"`
		SuccessCount      int64 `json:"success_count"`
		SuccessAmount     int64 `json:"success_amount"`
		FailedCount       int64 `json:"failed_count"`
		FailedAmount      int64 `json:"failed_amount"`
		PendingCount      int64 `json:"pending_count"`
		PendingAmount     int64 `json:"pending_amount"`
	}

	query := r.db.Model(&models.PaymentTransaction{})
	if gatewayID > 0 {
		query = query.Where("gateway_id = ?", gatewayID)
	}

	// کل تراکنش‌ها
	err := query.Count(&stats.TotalTransactions).Error
	if err != nil {
		return nil, err
	}

	// کل مبلغ
	err = query.Select("COALESCE(SUM(amount), 0)").Scan(&stats.TotalAmount).Error
	if err != nil {
		return nil, err
	}

	// تراکنش‌های موفق
	err = query.Where("status = ?", "success").Count(&stats.SuccessCount).Error
	if err != nil {
		return nil, err
	}

	err = query.Where("status = ?", "success").Select("COALESCE(SUM(amount), 0)").Scan(&stats.SuccessAmount).Error
	if err != nil {
		return nil, err
	}

	// تراکنش‌های ناموفق
	err = query.Where("status = ?", "failed").Count(&stats.FailedCount).Error
	if err != nil {
		return nil, err
	}

	err = query.Where("status = ?", "failed").Select("COALESCE(SUM(amount), 0)").Scan(&stats.FailedAmount).Error
	if err != nil {
		return nil, err
	}

	// تراکنش‌های در انتظار
	err = query.Where("status = ?", "pending").Count(&stats.PendingCount).Error
	if err != nil {
		return nil, err
	}

	err = query.Where("status = ?", "pending").Select("COALESCE(SUM(amount), 0)").Scan(&stats.PendingAmount).Error
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"total_transactions": stats.TotalTransactions,
		"total_amount":       stats.TotalAmount,
		"success_count":      stats.SuccessCount,
		"success_amount":     stats.SuccessAmount,
		"failed_count":       stats.FailedCount,
		"failed_amount":      stats.FailedAmount,
		"pending_count":      stats.PendingCount,
		"pending_amount":     stats.PendingAmount,
		"success_rate":       0.0,
	}

	// محاسبه نرخ موفقیت
	if stats.TotalTransactions > 0 {
		result["success_rate"] = float64(stats.SuccessCount) / float64(stats.TotalTransactions) * 100
	}

	return result, nil
}

// GetRecentTransactions دریافت تراکنش‌های اخیر
func (r *PaymentTransactionRepository) GetRecentTransactions(limit int) ([]models.PaymentTransaction, error) {
	var transactions []models.PaymentTransaction
	err := r.db.Preload("Gateway").Order("created_at DESC").Limit(limit).Find(&transactions).Error
	return transactions, err
}

// GetTransactionsByDateRange دریافت تراکنش‌ها در بازه زمانی
func (r *PaymentTransactionRepository) GetTransactionsByDateRange(startDate, endDate string, gatewayID uint) ([]models.PaymentTransaction, error) {
	var transactions []models.PaymentTransaction
	query := r.db.Preload("Gateway").Where("created_at BETWEEN ? AND ?", startDate, endDate)

	if gatewayID > 0 {
		query = query.Where("gateway_id = ?", gatewayID)
	}

	err := query.Order("created_at DESC").Find(&transactions).Error
	return transactions, err
}
