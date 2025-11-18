package unitofwork

import (
	"gorm.io/gorm"
)

// gormUnitOfWorkFactory پیاده‌سازی UnitOfWorkFactory برای GORM
type gormUnitOfWorkFactory struct {
	db *gorm.DB
}

// NewUnitOfWorkFactory ایجاد یک factory جدید برای ساخت UnitOfWork
func NewUnitOfWorkFactory(db *gorm.DB) UnitOfWorkFactory {
	return &gormUnitOfWorkFactory{db: db}
}

// Create ایجاد یک نمونه جدید از UnitOfWork
// این متد UnitOfWorkFactory interface را پیاده‌سازی می‌کند
func (f *gormUnitOfWorkFactory) Create() UnitOfWork {
	return NewGormUnitOfWork(f.db)
}
