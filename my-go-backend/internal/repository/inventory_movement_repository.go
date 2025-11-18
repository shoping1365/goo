package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// InventoryMovementRepositoryInterface اینترفیس ثبت/خواندن حرکات موجودی
type InventoryMovementRepositoryInterface interface {
	Create(ctx context.Context, m *models.InventoryMovement) error
	ListByProduct(ctx context.Context, productID uint, limit, offset int) ([]models.InventoryMovement, error)
}

type InventoryMovementRepository struct{ DB *gorm.DB }

func NewInventoryMovementRepository(db *gorm.DB) InventoryMovementRepositoryInterface {
	return &InventoryMovementRepository{DB: db}
}

func (r *InventoryMovementRepository) Create(ctx context.Context, m *models.InventoryMovement) error {
	return r.DB.WithContext(ctx).Create(m).Error
}

func (r *InventoryMovementRepository) ListByProduct(ctx context.Context, productID uint, limit, offset int) ([]models.InventoryMovement, error) {
	var res []models.InventoryMovement
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}
	if err := r.DB.WithContext(ctx).Where("product_id = ?", productID).
		Order("created_at DESC").Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
