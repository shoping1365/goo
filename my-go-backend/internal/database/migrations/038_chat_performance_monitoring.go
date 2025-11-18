package migrations

import "gorm.io/gorm"

// Up038ChatPerformanceMonitoring ایندکس/جداول مانیتورینگ چت در صورت نبود
func Up038ChatPerformanceMonitoring(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("CREATE TABLE IF NOT EXISTS chat_performance_metrics (id BIGSERIAL PRIMARY KEY, active_connections INT, avg_response_time DOUBLE PRECISION, messages_per_second DOUBLE PRECISION, error_rate DOUBLE PRECISION, cpu_usage DOUBLE PRECISION, memory_usage DOUBLE PRECISION, disk_usage DOUBLE PRECISION, network_usage DOUBLE PRECISION, performance_score DOUBLE PRECISION, created_at TIMESTAMPTZ DEFAULT NOW(), updated_at TIMESTAMPTZ DEFAULT NOW(), deleted_at TIMESTAMPTZ)")
		_ = tx.Exec("CREATE TABLE IF NOT EXISTS chat_alerts (id BIGSERIAL PRIMARY KEY, level VARCHAR(20), message TEXT, category VARCHAR(50), is_read BOOLEAN DEFAULT FALSE, created_at TIMESTAMPTZ DEFAULT NOW(), updated_at TIMESTAMPTZ DEFAULT NOW(), deleted_at TIMESTAMPTZ)")
		_ = tx.Exec("CREATE TABLE IF NOT EXISTS chat_system_health (id BIGSERIAL PRIMARY KEY, status VARCHAR(20), uptime VARCHAR(50), last_check TIMESTAMPTZ, total_connections INT, active_sessions INT, total_messages INT, created_at TIMESTAMPTZ DEFAULT NOW(), updated_at TIMESTAMPTZ DEFAULT NOW(), deleted_at TIMESTAMPTZ)")
		_ = tx.Exec("ALTER TABLE chat_system_health ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ")
		return nil
	})
}
