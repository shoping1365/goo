package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// ProductQARepository پیاده‌سازی repository برای پرسش و پاسخ محصولات
type ProductQARepository struct {
	db *gorm.DB
}

// NewProductQARepository ایجاد نمونه جدید از repository پرسش و پاسخ
func NewProductQARepository(db *gorm.DB) ProductQARepositoryInterface {
	return &ProductQARepository{db: db}
}

// Create ایجاد پرسش و پاسخ جدید
func (r *ProductQARepository) Create(ctx context.Context, qa *models.ProductQA) error {
	return r.db.WithContext(ctx).Create(qa).Error
}

// GetByID دریافت پرسش و پاسخ بر اساس ID
func (r *ProductQARepository) GetByID(ctx context.Context, id uint) (*models.ProductQA, error) {
	var qa models.ProductQA
	err := r.db.WithContext(ctx).Preload("Product").Preload("User").Preload("Parent").Preload("Replies").First(&qa, id).Error
	if err != nil {
		return nil, err
	}
	return &qa, nil
}

// GetAll دریافت تمام پرسش و پاسخ‌ها
func (r *ProductQARepository) GetAll(ctx context.Context) ([]models.ProductQA, error) {
	var qas []models.ProductQA
	err := r.db.WithContext(ctx).Preload("Product").Preload("User").Preload("Parent").Preload("Replies").Order("created_at DESC").Find(&qas).Error
	return qas, err
}

// GetByProduct دریافت پرسش و پاسخ‌های یک محصول
func (r *ProductQARepository) GetByProduct(ctx context.Context, productID uint, limit, offset int) ([]models.ProductQA, error) {
	var qas []models.ProductQA
	query := r.db.WithContext(ctx).Preload("Product").Preload("User").Preload("Replies").Where("product_id = ? AND parent_id IS NULL", productID)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&qas).Error
	return qas, err
}

// GetByUser دریافت پرسش و پاسخ‌های یک کاربر
func (r *ProductQARepository) GetByUser(ctx context.Context, userID uint, limit, offset int) ([]models.ProductQA, error) {
	var qas []models.ProductQA
	query := r.db.WithContext(ctx).Preload("Product").Preload("Parent").Preload("Replies").Where("user_id = ?", userID)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&qas).Error
	return qas, err
}

// GetQuestions دریافت سوالات
func (r *ProductQARepository) GetQuestions(ctx context.Context, productID uint, limit, offset int) ([]models.ProductQA, error) {
	var questions []models.ProductQA
	query := r.db.WithContext(ctx).Preload("User").Preload("Replies").Where("product_id = ? AND parent_id IS NULL", productID)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&questions).Error
	return questions, err
}

// GetAnswers دریافت پاسخ‌ها
func (r *ProductQARepository) GetAnswers(ctx context.Context, questionID uint, limit, offset int) ([]models.ProductQA, error) {
	var answers []models.ProductQA
	query := r.db.WithContext(ctx).Preload("User").Where("parent_id = ?", questionID)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at ASC").Find(&answers).Error
	return answers, err
}

// GetPublished دریافت پرسش و پاسخ‌های منتشر شده
func (r *ProductQARepository) GetPublished(ctx context.Context, productID uint, limit, offset int) ([]models.ProductQA, error) {
	var qas []models.ProductQA
	query := r.db.WithContext(ctx).Preload("User").Preload("Replies").Where("product_id = ? AND parent_id IS NULL AND status = ?", productID, "approved")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&qas).Error
	return qas, err
}

// GetPending دریافت پرسش و پاسخ‌های در انتظار تایید
func (r *ProductQARepository) GetPending(ctx context.Context, limit, offset int) ([]models.ProductQA, error) {
	var qas []models.ProductQA
	query := r.db.WithContext(ctx).Preload("Product").Preload("User").Where("status = ?", "pending")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&qas).Error
	return qas, err
}

// UpdateStatus به‌روزرسانی وضعیت
func (r *ProductQARepository) UpdateStatus(ctx context.Context, qaID uint, status string) error {
	return r.db.WithContext(ctx).Model(&models.ProductQA{}).Where("id = ?", qaID).Update("status", status).Error
}

// GetQuestionCount دریافت تعداد سوالات محصول
func (r *ProductQARepository) GetQuestionCount(ctx context.Context, productID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.ProductQA{}).
		Where("product_id = ? AND parent_id IS NULL AND status = ?", productID, "approved").
		Count(&count).Error
	return count, err
}

// GetAnswerCount دریافت تعداد پاسخ‌های سوال
func (r *ProductQARepository) GetAnswerCount(ctx context.Context, questionID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.ProductQA{}).
		Where("parent_id = ? AND status = ?", questionID, "approved").
		Count(&count).Error
	return count, err
}

// Update به‌روزرسانی پرسش و پاسخ
func (r *ProductQARepository) Update(ctx context.Context, qa *models.ProductQA) error {
	return r.db.WithContext(ctx).Save(qa).Error
}

// Delete حذف پرسش و پاسخ
func (r *ProductQARepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.ProductQA{}, id).Error
}

// Count تعداد پرسش و پاسخ‌ها
func (r *ProductQARepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.ProductQA{}).Count(&count).Error
	return count, err
}
