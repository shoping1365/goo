package models

import "time"

// BackupEntry maps each original file to the month-period zip/folder where a raw copy is stored.
// It allows constant-time lookup when deleting media or restoring.
// Table name: backups
//
// rel_path     → uploads/media/library/images/x.jpg  (without leading slash)
// period       → YYYY-MM (same as folder or zip name)
// archived     → false = still folder; true = already zipped
// updated_at   → when archived flipped to true
//
// NOTE: this table is populated in UploadMediaHandler (AddOriginal) and
//       updated by the monthly archive job.
type BackupEntry struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	RelPath   string    `gorm:"type:varchar(255);uniqueIndex" json:"rel_path"`
	Period    string    `gorm:"type:char(7);index" json:"period"` // YYYY-MM
	Archived  bool      `gorm:"default:false" json:"archived"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
