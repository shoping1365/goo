package repository

import (
	"context"
	"errors"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// WishlistRepository پیاده‌سازی repository برای لیست علاقمندی‌ها
type WishlistRepository struct {
	db *gorm.DB
}

// NewWishlistRepository ایجاد نمونه جدید از repository لیست علاقمندی
func NewWishlistRepository(db *gorm.DB) WishlistRepositoryInterface {
	return &WishlistRepository{db: db}
}

// Create ایجاد لیست علاقمندی جدید
func (r *WishlistRepository) Create(ctx context.Context, wishlist *models.Wishlist) error {
	return r.db.WithContext(ctx).Create(wishlist).Error
}

// GetByID دریافت لیست علاقمندی بر اساس ID
func (r *WishlistRepository) GetByID(ctx context.Context, id uint) (*models.Wishlist, error) {
	var wishlist models.Wishlist
	err := r.db.WithContext(ctx).Preload("User").Preload("Items.Product").First(&wishlist, id).Error
	if err != nil {
		return nil, err
	}
	return &wishlist, nil
}

// GetByUser دریافت لیست علاقمندی‌های کاربر
func (r *WishlistRepository) GetByUser(ctx context.Context, userID uint) (*models.Wishlist, error) {
	var wishlist models.Wishlist
	err := r.db.WithContext(ctx).Preload("User").Preload("Items.Product").Where("user_id = ?", userID).First(&wishlist).Error
	if err != nil {
		return nil, err
	}
	return &wishlist, nil
}

// GetAll دریافت تمام لیست‌های علاقمندی
func (r *WishlistRepository) GetAll(ctx context.Context) ([]models.Wishlist, error) {
	var wishlists []models.Wishlist
	err := r.db.WithContext(ctx).Preload("User").Preload("Items.Product").Order("created_at DESC").Find(&wishlists).Error
	return wishlists, err
}

// AddProduct افزودن محصول به لیست علاقمندی‌ها
func (r *WishlistRepository) AddProduct(ctx context.Context, userID, productID uint) error {
	// ابتدا بررسی می‌کنیم که آیا لیست علاقمندی برای کاربر وجود دارد یا نه
	var wishlist models.Wishlist
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&wishlist).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// ایجاد لیست علاقمندی جدید
		wishlist = models.Wishlist{UserID: userID}
		if err := r.db.WithContext(ctx).Create(&wishlist).Error; err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// بررسی اینکه آیا محصول قبلاً در لیست وجود دارد یا نه
	var existingItem models.WishlistItem
	err = r.db.WithContext(ctx).Where("wishlist_id = ? AND product_id = ?", wishlist.ID, productID).First(&existingItem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// افزودن محصول جدید
		item := models.WishlistItem{
			WishlistID: wishlist.ID,
			ProductID:  productID,
		}
		return r.db.WithContext(ctx).Create(&item).Error
	}
	return nil // محصول قبلاً وجود دارد
}

// RemoveProduct حذف محصول از لیست علاقمندی‌ها
func (r *WishlistRepository) RemoveProduct(ctx context.Context, userID, productID uint) error {
	return r.db.WithContext(ctx).
		Joins("JOIN wishlists ON wishlist_items.wishlist_id = wishlists.id").
		Where("wishlists.user_id = ? AND wishlist_items.product_id = ?", userID, productID).
		Delete(&models.WishlistItem{}).Error
}

// IsInWishlist بررسی وجود محصول در لیست علاقمندی‌ها
func (r *WishlistRepository) IsInWishlist(ctx context.Context, userID, productID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Joins("JOIN wishlists ON wishlist_items.wishlist_id = wishlists.id").
		Where("wishlists.user_id = ? AND wishlist_items.product_id = ?", userID, productID).
		Model(&models.WishlistItem{}).Count(&count).Error
	return count > 0, err
}

// GetWishlistProducts دریافت محصولات لیست علاقمندی‌ها
func (r *WishlistRepository) GetWishlistProducts(ctx context.Context, userID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(ctx).
		Joins("JOIN wishlist_items ON products.id = wishlist_items.product_id").
		Joins("JOIN wishlists ON wishlist_items.wishlist_id = wishlists.id").
		Where("wishlists.user_id = ?", userID).
		Preload("Category").
		Preload("Images").
		Order("wishlist_items.created_at DESC").
		Find(&products).Error
	return products, err
}

// GetWishlistCount دریافت تعداد محصولات در لیست علاقمندی‌ها
func (r *WishlistRepository) GetWishlistCount(ctx context.Context, userID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Joins("JOIN wishlists ON wishlist_items.wishlist_id = wishlists.id").
		Where("wishlists.user_id = ?", userID).
		Model(&models.WishlistItem{}).Count(&count).Error
	return count, err
}

// Update به‌روزرسانی لیست علاقمندی
func (r *WishlistRepository) Update(ctx context.Context, wishlist *models.Wishlist) error {
	return r.db.WithContext(ctx).Save(wishlist).Error
}

// Delete حذف لیست علاقمندی
func (r *WishlistRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Wishlist{}, id).Error
}

// Count تعداد لیست‌های علاقمندی
func (r *WishlistRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Wishlist{}).Count(&count).Error
	return count, err
}
