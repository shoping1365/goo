package repository

import (
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// UserVerificationRepository - ریپوزیتوری برای عملیات احراز هویت
type UserVerificationRepository struct {
	db *gorm.DB
}

// NewUserVerificationRepository - سازنده ریپوزیتوری
func NewUserVerificationRepository(db *gorm.DB) *UserVerificationRepository {
	return &UserVerificationRepository{
		db: db.Session(&gorm.Session{PrepareStmt: true}),
	}
}

// Create - ایجاد درخواست احراز هویت جدید
func (r *UserVerificationRepository) Create(verification *models.UserVerification) error {
	return r.db.Create(verification).Error
}

// FindByID - جستجوی درخواست بر اساس ID
func (r *UserVerificationRepository) FindByID(id uint) (*models.UserVerification, error) {
	var verification models.UserVerification
	err := r.db.Preload("User").Preload("VerifiedByUser").First(&verification, id).Error
	if err != nil {
		return nil, err
	}
	return &verification, nil
}

// FindByUserID - جستجوی درخواست‌های یک کاربر
func (r *UserVerificationRepository) FindByUserID(userID uint) ([]models.UserVerification, error) {
	var verifications []models.UserVerification
	err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Preload("VerifiedByUser").
		Find(&verifications).Error
	if err != nil {
		return nil, err
	}
	return verifications, nil
}

// FindByUserIDAndField - جستجوی درخواست بر اساس کاربر و نام فیلد
func (r *UserVerificationRepository) FindByUserIDAndField(userID uint, fieldName string) (*models.UserVerification, error) {
	var verification models.UserVerification
	err := r.db.Where("user_id = ? AND field_name = ?", userID, fieldName).
		Order("created_at DESC").
		First(&verification).Error
	if err != nil {
		return nil, err
	}
	return &verification, nil
}

// FindByStatus - جستجوی درخواست‌ها بر اساس وضعیت
func (r *UserVerificationRepository) FindByStatus(status string, limit, offset int) ([]models.UserVerification, int64, error) {
	var verifications []models.UserVerification
	var total int64

	// شمارش کل
	if err := r.db.Model(&models.UserVerification{}).Where("status = ?", status).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// دریافت داده‌ها
	err := r.db.Where("status = ?", status).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Preload("User").
		Preload("VerifiedByUser").
		Find(&verifications).Error

	if err != nil {
		return nil, 0, err
	}

	return verifications, total, nil
}

// FindAll - دریافت تمام درخواست‌ها با صفحه‌بندی
func (r *UserVerificationRepository) FindAll(limit, offset int) ([]models.UserVerification, int64, error) {
	var verifications []models.UserVerification
	var total int64

	// شمارش کل
	if err := r.db.Model(&models.UserVerification{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// دریافت داده‌ها
	err := r.db.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Preload("User").
		Preload("VerifiedByUser").
		Find(&verifications).Error

	if err != nil {
		return nil, 0, err
	}

	return verifications, total, nil
}

// Update - به‌روزرسانی درخواست احراز هویت
func (r *UserVerificationRepository) Update(verification *models.UserVerification) error {
	return r.db.Save(verification).Error
}

// UpdateStatus - به‌روزرسانی وضعیت احراز هویت
func (r *UserVerificationRepository) UpdateStatus(id uint, status string, verifiedBy *uint, rejectionReason string) error {
	updates := map[string]interface{}{
		"status": status,
	}

	if status == models.VerificationStatusVerified {
		updates["verified_at"] = gorm.Expr("NOW()")
		if verifiedBy != nil {
			updates["verified_by"] = *verifiedBy
		}
	} else if status == models.VerificationStatusRejected {
		updates["rejection_reason"] = rejectionReason
		if verifiedBy != nil {
			updates["verified_by"] = *verifiedBy
		}
	}

	return r.db.Model(&models.UserVerification{}).Where("id = ?", id).Updates(updates).Error
}

// Delete - حذف درخواست احراز هویت
func (r *UserVerificationRepository) Delete(id uint) error {
	return r.db.Delete(&models.UserVerification{}, id).Error
}

// GetPendingCount - شمارش درخواست‌های در انتظار
func (r *UserVerificationRepository) GetPendingCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.UserVerification{}).
		Where("status = ?", models.VerificationStatusPending).
		Count(&count).Error
	return count, err
}

// GetVerifiedFieldsByUserID - دریافت فیلدهای تایید شده یک کاربر
func (r *UserVerificationRepository) GetVerifiedFieldsByUserID(userID uint) (map[string]bool, error) {
	var verifications []models.UserVerification
	err := r.db.Where("user_id = ? AND status = ?", userID, models.VerificationStatusVerified).
		Select("field_name").
		Find(&verifications).Error

	if err != nil {
		return nil, err
	}

	verifiedFields := make(map[string]bool)
	for _, v := range verifications {
		verifiedFields[v.FieldName] = true
	}

	return verifiedFields, nil
}

// SearchVerifications - جستجوی پیشرفته درخواست‌ها
func (r *UserVerificationRepository) SearchVerifications(params map[string]interface{}, limit, offset int) ([]models.UserVerification, int64, error) {
	var verifications []models.UserVerification
	var total int64

	query := r.db.Model(&models.UserVerification{})

	// فیلترها
	if userID, ok := params["user_id"].(uint); ok && userID > 0 {
		query = query.Where("user_id = ?", userID)
	}

	if status, ok := params["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}

	if fieldName, ok := params["field_name"].(string); ok && fieldName != "" {
		query = query.Where("field_name = ?", fieldName)
	}

	if verificationMethod, ok := params["verification_method"].(string); ok && verificationMethod != "" {
		query = query.Where("verification_method = ?", verificationMethod)
	}

	// شمارش کل
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// دریافت داده‌ها
	err := query.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Preload("User").
		Preload("VerifiedByUser").
		Find(&verifications).Error

	if err != nil {
		return nil, 0, err
	}

	return verifications, total, nil
}
