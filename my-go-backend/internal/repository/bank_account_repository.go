package repository

import (
	"context"
	"fmt"

	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// BankAccountRepositoryInterface اینترفیس عملیات دیتابیس برای کارت/حساب بانکی کاربر
type BankAccountRepositoryInterface interface {
	List(ctx context.Context, limit, offset int, status, bank string) ([]models.UserBankAccount, int64, error)
	Verify(ctx context.Context, id uint) error
	Block(ctx context.Context, id uint) error
	Unblock(ctx context.Context, id uint) error
	Reject(ctx context.Context, id uint) error
}

type BankAccountRepository struct{ DB *gorm.DB }

func NewBankAccountRepository(db *gorm.DB) BankAccountRepositoryInterface {
	return &BankAccountRepository{DB: db}
}

func (r *BankAccountRepository) List(ctx context.Context, limit, offset int, status, bank string) ([]models.UserBankAccount, int64, error) {
	var items []models.UserBankAccount
	var total int64
	q := r.DB.WithContext(ctx).Model(&models.UserBankAccount{})
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if bank != "" {
		q = q.Where("bank_name = ?", bank)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order("created_at DESC").Limit(limit).Offset(offset).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *BankAccountRepository) Verify(ctx context.Context, id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid id")
	}
	return r.DB.WithContext(ctx).Model(&models.UserBankAccount{}).Where("id = ?", id).Updates(map[string]any{"status": "active", "verified_at": gorm.Expr("NOW()")}).Error
}

func (r *BankAccountRepository) Block(ctx context.Context, id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid id")
	}
	return r.DB.WithContext(ctx).Model(&models.UserBankAccount{}).Where("id = ?", id).Update("status", "blocked").Error
}

func (r *BankAccountRepository) Unblock(ctx context.Context, id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid id")
	}
	return r.DB.WithContext(ctx).Model(&models.UserBankAccount{}).Where("id = ?", id).Update("status", "active").Error
}

func (r *BankAccountRepository) Reject(ctx context.Context, id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid id")
	}
	return r.DB.WithContext(ctx).Model(&models.UserBankAccount{}).Where("id = ?", id).Update("status", "rejected").Error
}
