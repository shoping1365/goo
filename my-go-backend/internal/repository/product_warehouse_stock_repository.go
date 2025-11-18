package repository

import (
	"context"
	"errors"
	"my-go-backend/internal/models"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ProductWarehouseStockRepositoryInterface اینترفیس موجودی محصول-انبار
type ProductWarehouseStockRepositoryInterface interface {
	// خواندن
	GetByProduct(ctx context.Context, productID uint) ([]models.ProductWarehouseStock, error)
	GetByProductAndWarehouseForUpdate(ctx context.Context, productID uint, warehouseID uint) (*models.ProductWarehouseStock, error)
	SumByProduct(ctx context.Context, productID uint) (int, int, error)
	// نوشتن
	Upsert(ctx context.Context, s *models.ProductWarehouseStock) error
	AdjustQuantity(ctx context.Context, productID, warehouseID uint, delta int) error
}

type ProductWarehouseStockRepository struct {
	DB *gorm.DB
}

func NewProductWarehouseStockRepository(db *gorm.DB) ProductWarehouseStockRepositoryInterface {
	return &ProductWarehouseStockRepository{DB: db}
}

func (r *ProductWarehouseStockRepository) GetByProduct(ctx context.Context, productID uint) ([]models.ProductWarehouseStock, error) {
	var res []models.ProductWarehouseStock
	if err := r.DB.WithContext(ctx).Where("product_id = ?", productID).Order("warehouse_id ASC").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ProductWarehouseStockRepository) GetByProductAndWarehouseForUpdate(ctx context.Context, productID uint, warehouseID uint) (*models.ProductWarehouseStock, error) {
	var m models.ProductWarehouseStock
	if err := r.DB.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			m = models.ProductWarehouseStock{ProductID: productID, WarehouseID: warehouseID, Quantity: 0, Reserved: 0}
			if err := r.DB.WithContext(ctx).Create(&m).Error; err != nil {
				return nil, err
			}
			return &m, nil
		}
		return nil, err
	}
	return &m, nil
}

func (r *ProductWarehouseStockRepository) SumByProduct(ctx context.Context, productID uint) (int, int, error) {
	type agg struct {
		Qty int
		Res int
	}
	var a agg
	if err := r.DB.WithContext(ctx).Model(&models.ProductWarehouseStock{}).
		Select("COALESCE(SUM(quantity),0) as qty, COALESCE(SUM(reserved),0) as res").
		Where("product_id = ?", productID).Scan(&a).Error; err != nil {
		return 0, 0, err
	}
	return a.Qty, a.Res, nil
}

func (r *ProductWarehouseStockRepository) Upsert(ctx context.Context, s *models.ProductWarehouseStock) error {
	// اگر رکورد موجود بود، به‌روزرسانی؛ در غیر این صورت ایجاد
	var existing models.ProductWarehouseStock
	err := r.DB.WithContext(ctx).Where("product_id = ? AND warehouse_id = ?", s.ProductID, s.WarehouseID).First(&existing).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// تنظیم زمان ایجاد برای رکورد جدید
			s.CreatedAt = time.Now()
			s.UpdatedAt = time.Now()
			return r.DB.WithContext(ctx).Create(s).Error
		}
		return err
	}
	// به‌روزرسانی فیلدها
	existing.Quantity = s.Quantity
	existing.Reserved = s.Reserved
	existing.MinQty = s.MinQty
	existing.MaxQty = s.MaxQty
	existing.UpdatedAt = time.Now()
	return r.DB.WithContext(ctx).Save(&existing).Error
}

func (r *ProductWarehouseStockRepository) AdjustQuantity(ctx context.Context, productID, warehouseID uint, delta int) error {
	return r.DB.WithContext(ctx).Exec(
		`UPDATE product_warehouse_stocks SET quantity = quantity + ?, updated_at = NOW() WHERE product_id = ? AND warehouse_id = ?`,
		delta, productID, warehouseID,
	).Error
}
