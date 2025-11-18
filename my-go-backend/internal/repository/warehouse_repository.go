package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// WarehouseRepositoryInterface اینترفیس عملیات انبار
type WarehouseRepositoryInterface interface {
	BaseRepository[models.Warehouse]
	GetDefault(ctx context.Context) (*models.Warehouse, error)
	SetDefault(ctx context.Context, id uint) error
}

type WarehouseRepository struct {
	DB *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepositoryInterface {
	return &WarehouseRepository{DB: db}
}

func (r *WarehouseRepository) Create(ctx context.Context, e *models.Warehouse) error {
	return r.DB.WithContext(ctx).Create(e).Error
}
func (r *WarehouseRepository) GetByID(ctx context.Context, id uint) (*models.Warehouse, error) {
	var m models.Warehouse
	if err := r.DB.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
func (r *WarehouseRepository) GetAll(ctx context.Context) ([]models.Warehouse, error) {
	var res []models.Warehouse
	if err := r.DB.WithContext(ctx).Order("priority ASC, id ASC").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func (r *WarehouseRepository) Update(ctx context.Context, e *models.Warehouse) error {
	return r.DB.WithContext(ctx).Save(e).Error
}
func (r *WarehouseRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Warehouse{}, id).Error
}
func (r *WarehouseRepository) Count(ctx context.Context) (int64, error) {
	var c int64
	if err := r.DB.WithContext(ctx).Model(&models.Warehouse{}).Count(&c).Error; err != nil {
		return 0, err
	}
	return c, nil
}
func (r *WarehouseRepository) GetDefault(ctx context.Context) (*models.Warehouse, error) {
	var m models.Warehouse
	if err := r.DB.WithContext(ctx).Where("is_default = true").First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
func (r *WarehouseRepository) SetDefault(ctx context.Context, id uint) error {
	if err := r.DB.WithContext(ctx).Model(&models.Warehouse{}).Where("is_default = true").Update("is_default", false).Error; err != nil {
		return err
	}
	return r.DB.WithContext(ctx).Model(&models.Warehouse{}).Where("id = ?", id).Update("is_default", true).Error
}
