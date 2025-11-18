package migrations

import "gorm.io/gorm"

// Up035GeoAndReferenceData جداول مرجع (حداقلی برای استان/شهر در صورت نیاز)
func Up035GeoAndReferenceData(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS provinces (id BIGSERIAL PRIMARY KEY, name VARCHAR(100) UNIQUE NOT NULL)`)
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS cities (id BIGSERIAL PRIMARY KEY, province_id BIGINT NOT NULL, name VARCHAR(100) NOT NULL, CONSTRAINT fk_cities_province FOREIGN KEY(province_id) REFERENCES provinces(id) ON DELETE CASCADE)`)
		_ = tx.Exec(`CREATE INDEX IF NOT EXISTS idx_cities_province ON cities(province_id)`)
		return nil
	})
}
