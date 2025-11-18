package migrations

import (
	"gorm.io/gorm"
)

// Migration130AddRecentViewsTable creates the recent_views table
func Migration130AddRecentViewsTable(db *gorm.DB) error {
	return db.Exec(`
		-- Create recent_views table
		CREATE TABLE IF NOT EXISTS recent_views (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			viewed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			
			-- Foreign key constraints
			CONSTRAINT fk_recent_views_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			CONSTRAINT fk_recent_views_product_id FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
			
			-- Unique constraint to prevent duplicate entries
			CONSTRAINT unique_user_product_recent_view UNIQUE(user_id, product_id)
		);
		
		-- Create indexes for better performance
		CREATE INDEX IF NOT EXISTS idx_recent_views_user_id ON recent_views(user_id);
		CREATE INDEX IF NOT EXISTS idx_recent_views_product_id ON recent_views(product_id);
		CREATE INDEX IF NOT EXISTS idx_recent_views_viewed_at ON recent_views(viewed_at);
		CREATE INDEX IF NOT EXISTS idx_recent_views_user_viewed_at ON recent_views(user_id, viewed_at DESC);
		
		-- Comment for documentation
		COMMENT ON TABLE recent_views IS 'Stores user recent product views for personalization';
		COMMENT ON COLUMN recent_views.user_id IS 'Reference to the user who viewed the product';
		COMMENT ON COLUMN recent_views.product_id IS 'Reference to the viewed product';
		COMMENT ON COLUMN recent_views.viewed_at IS 'Timestamp when the product was viewed';
	`).Error
}
