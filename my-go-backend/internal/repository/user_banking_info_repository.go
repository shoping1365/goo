package repository

import (
	"errors"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// UserBankingInfoRepository رابط برای عملیات اطلاعات بانکی کاربران
type UserBankingInfoRepository interface {
	Create(bankingInfo *models.UserBankingInfo) error
	GetByUserID(userID uint) ([]models.UserBankingInfo, error)
	GetDefaultByUserID(userID uint) (*models.UserBankingInfo, error)
	GetByID(id uint) (*models.UserBankingInfo, error)
	Update(bankingInfo *models.UserBankingInfo) error
	Delete(id uint) error
	DeleteByUserID(userID uint) error
	SetDefault(userID uint, id uint) error
	GetAllVerified() ([]models.UserBankingInfo, error)
	GetByBankName(bankName string) ([]models.UserBankingInfo, error)
}

// userBankingInfoRepository پیاده‌سازی رابط UserBankingInfoRepository
type userBankingInfoRepository struct {
	db *gorm.DB
}

// NewUserBankingInfoRepository ایجاد نمونه جدید از UserBankingInfoRepository
func NewUserBankingInfoRepository(db *gorm.DB) UserBankingInfoRepository {
	return &userBankingInfoRepository{db: db}
}

// Create ایجاد اطلاعات بانکی جدید
func (r *userBankingInfoRepository) Create(bankingInfo *models.UserBankingInfo) error {
	return r.db.Create(bankingInfo).Error
}

// GetByUserID دریافت تمام اطلاعات بانکی بر اساس شناسه کاربر
func (r *userBankingInfoRepository) GetByUserID(userID uint) ([]models.UserBankingInfo, error) {
	var bankingInfos []models.UserBankingInfo
	err := r.db.Where("user_id = ?", userID).Order("is_default DESC, created_at DESC").Find(&bankingInfos).Error
	return bankingInfos, err
}

// GetDefaultByUserID دریافت حساب پیشفرض کاربر
func (r *userBankingInfoRepository) GetDefaultByUserID(userID uint) (*models.UserBankingInfo, error) {
	var bankingInfo models.UserBankingInfo
	err := r.db.Where("user_id = ? AND is_default = ?", userID, true).First(&bankingInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // حساب پیشفرض یافت نشد
		}
		return nil, err
	}
	return &bankingInfo, nil
}

// GetByID دریافت اطلاعات بانکی بر اساس شناسه
func (r *userBankingInfoRepository) GetByID(id uint) (*models.UserBankingInfo, error) {
	var bankingInfo models.UserBankingInfo
	err := r.db.First(&bankingInfo, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &bankingInfo, nil
}

// Update به‌روزرسانی اطلاعات بانکی
func (r *userBankingInfoRepository) Update(bankingInfo *models.UserBankingInfo) error {
	return r.db.Save(bankingInfo).Error
}

// Delete حذف اطلاعات بانکی بر اساس شناسه
func (r *userBankingInfoRepository) Delete(id uint) error {
	return r.db.Delete(&models.UserBankingInfo{}, id).Error
}

// DeleteByUserID حذف تمام اطلاعات بانکی کاربر
func (r *userBankingInfoRepository) DeleteByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.UserBankingInfo{}).Error
}

// SetDefault تنظیم حساب پیشفرض
func (r *userBankingInfoRepository) SetDefault(userID uint, id uint) error {
	// ابتدا تمام حساب‌های کاربر را غیرپیشفرض کنیم
	if err := r.db.Model(&models.UserBankingInfo{}).Where("user_id = ?", userID).Update("is_default", false).Error; err != nil {
		return err
	}

	// اگر id برابر با 0 است، فقط غیرپیشفرض کردن انجام شده
	if id == 0 {
		return nil
	}

	// سپس حساب مورد نظر را پیشفرض کنیم
	return r.db.Model(&models.UserBankingInfo{}).Where("id = ? AND user_id = ?", id, userID).Update("is_default", true).Error
}

// GetAllVerified دریافت تمام اطلاعات بانکی تایید شده
func (r *userBankingInfoRepository) GetAllVerified() ([]models.UserBankingInfo, error) {
	var bankingInfos []models.UserBankingInfo
	err := r.db.Where("is_verified = ?", true).Find(&bankingInfos).Error
	return bankingInfos, err
}

// GetByBankName دریافت اطلاعات بانکی بر اساس نام بانک
func (r *userBankingInfoRepository) GetByBankName(bankName string) ([]models.UserBankingInfo, error) {
	var bankingInfos []models.UserBankingInfo
	err := r.db.Where("bank_name = ?", bankName).Find(&bankingInfos).Error
	return bankingInfos, err
}
