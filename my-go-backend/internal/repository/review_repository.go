package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// ReviewRepository پیاده‌سازی repository برای نظرات
type ReviewRepository struct {
	db *gorm.DB
}

// NewReviewRepository ایجاد نمونه جدید از repository نظر
func NewReviewRepository(db *gorm.DB) ReviewRepositoryInterface {
	return &ReviewRepository{db: db}
}

// Create ایجاد نظر جدید
func (r *ReviewRepository) Create(ctx context.Context, review *models.Review) error {
	return r.db.WithContext(ctx).Create(review).Error
}

// GetByID دریافت نظر بر اساس ID
func (r *ReviewRepository) GetByID(ctx context.Context, id uint) (*models.Review, error) {
	var review models.Review
	err := r.db.WithContext(ctx).Preload("Customer").Preload("Product").First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// GetAll دریافت تمام نظرات
func (r *ReviewRepository) GetAll(ctx context.Context) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.WithContext(ctx).Preload("Customer").Preload("Product").Order("created_at DESC").Find(&reviews).Error
	return reviews, err
}

// GetByProduct دریافت نظرات یک محصول
func (r *ReviewRepository) GetByProduct(ctx context.Context, productID uint, limit, offset int) ([]models.Review, error) {
	var reviews []models.Review
	query := r.db.WithContext(ctx).Preload("Customer").Where("product_id = ?", productID)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&reviews).Error
	return reviews, err
}

// GetByCustomer دریافت نظرات یک مشتری
func (r *ReviewRepository) GetByCustomer(ctx context.Context, customerID uint, limit, offset int) ([]models.Review, error) {
	var reviews []models.Review
	query := r.db.WithContext(ctx).Preload("Product").Where("customer_id = ?", customerID)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&reviews).Error
	return reviews, err
}

// GetPublished دریافت نظرات منتشر شده
func (r *ReviewRepository) GetPublished(ctx context.Context, limit, offset int) ([]models.Review, error) {
	var reviews []models.Review
	query := r.db.WithContext(ctx).Preload("Customer").Preload("Product").Where("status = ?", "approved")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&reviews).Error
	return reviews, err
}

// GetPending دریافت نظرات در انتظار تایید
func (r *ReviewRepository) GetPending(ctx context.Context, limit, offset int) ([]models.Review, error) {
	var reviews []models.Review
	query := r.db.WithContext(ctx).Preload("Customer").Preload("Product").Where("status = ?", "pending")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&reviews).Error
	return reviews, err
}

// GetByRating دریافت نظرات بر اساس امتیاز
func (r *ReviewRepository) GetByRating(ctx context.Context, productID uint, rating int, limit, offset int) ([]models.Review, error) {
	var reviews []models.Review
	query := r.db.WithContext(ctx).Preload("Customer").Where("product_id = ? AND rating = ? AND status = ?", productID, rating, "approved")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&reviews).Error
	return reviews, err
}

// GetAverageRating دریافت میانگین امتیاز محصول
func (r *ReviewRepository) GetAverageRating(ctx context.Context, productID uint) (float64, error) {
	var avgRating float64
	err := r.db.WithContext(ctx).Model(&models.Review{}).
		Where("product_id = ? AND status = ?", productID, "approved").
		Select("COALESCE(AVG(rating), 0)").
		Scan(&avgRating).Error
	return avgRating, err
}

// GetReviewCount دریافت تعداد نظرات محصول
func (r *ReviewRepository) GetReviewCount(ctx context.Context, productID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Review{}).
		Where("product_id = ? AND status = ?", productID, "approved").
		Count(&count).Error
	return count, err
}

// UpdateStatus به‌روزرسانی وضعیت نظر
func (r *ReviewRepository) UpdateStatus(ctx context.Context, reviewID uint, status string) error {
	return r.db.WithContext(ctx).Model(&models.Review{}).Where("id = ?", reviewID).Update("status", status).Error
}

// IncrementHelpfulCount افزایش تعداد مفید بودن
func (r *ReviewRepository) IncrementHelpfulCount(ctx context.Context, reviewID uint) error {
	return r.db.WithContext(ctx).Model(&models.Review{}).Where("id = ?", reviewID).Update("helpful_count", gorm.Expr("helpful_count + 1")).Error
}

// HasHelpfulVote بررسی رأی تکراری کاربر برای یک نظر
func (r *ReviewRepository) HasHelpfulVote(ctx context.Context, reviewID uint, userID uint) (bool, error) {
	type ReviewHelpfulVote struct {
		ID       uint
		ReviewID uint
		UserID   uint
	}
	var count int64
	err := r.db.WithContext(ctx).Model(&ReviewHelpfulVote{}).Where("review_id = ? AND user_id = ?", reviewID, userID).Count(&count).Error
	return count > 0, err
}

// CreateHelpfulVote ثبت رأی مفید کاربر برای یک نظر
func (r *ReviewRepository) CreateHelpfulVote(ctx context.Context, reviewID uint, userID uint) error {
	type ReviewHelpfulVote struct {
		ID       uint
		ReviewID uint
		UserID   uint
	}
	return r.db.WithContext(ctx).Create(&ReviewHelpfulVote{ReviewID: reviewID, UserID: userID}).Error
}

// Update به‌روزرسانی نظر
func (r *ReviewRepository) Update(ctx context.Context, review *models.Review) error {
	return r.db.WithContext(ctx).Save(review).Error
}

// Delete حذف نظر
func (r *ReviewRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Review{}, id).Error
}

// Count تعداد نظرات
func (r *ReviewRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Review{}).Count(&count).Error
	return count, err
}
