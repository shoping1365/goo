package migrations

import (
	"gorm.io/gorm"
)

// Up123AddParentSlugToCategories اضافه کردن فیلد parent_slug به جدول categories
func Up123AddParentSlugToCategories(db *gorm.DB) error {
	return db.Exec(`
		DO $$ 
		BEGIN
			-- بررسی وجود فیلد parent_slug
			IF NOT EXISTS (
				SELECT 1 FROM information_schema.columns 
				WHERE table_name = 'categories' 
				AND column_name = 'parent_slug'
			) THEN
				-- اضافه کردن فیلد parent_slug
				ALTER TABLE categories 
				ADD COLUMN parent_slug VARCHAR(120);
				
				-- ایجاد ایندکس
				CREATE INDEX IF NOT EXISTS idx_categories_parent_slug ON categories(parent_slug);
			END IF;
		END $$;
	`).Error
}

// Down123AddParentSlugToCategories حذف فیلد parent_slug از جدول categories
func Down123AddParentSlugToCategories(db *gorm.DB) error {
	return db.Exec(`
		DROP INDEX IF EXISTS idx_categories_parent_slug;
		ALTER TABLE categories 
		DROP COLUMN parent_slug;
	`).Error
}
