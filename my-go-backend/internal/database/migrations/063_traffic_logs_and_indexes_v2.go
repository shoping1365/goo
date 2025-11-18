package migrations

import "gorm.io/gorm"

// Up063TrafficLogsAndIndexesV2 ایندکس‌های ضروری ترافیک
func Up063TrafficLogsAndIndexesV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_traffic_logs_created ON traffic_logs(created_at)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_traffic_logs_path ON traffic_logs(request_path)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_blocked_ips_active ON blocked_ips(is_active)")
		return nil
	})
}
