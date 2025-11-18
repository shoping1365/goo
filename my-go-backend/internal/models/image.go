package models

type Image struct {
	// ... سایر فیلدها ...
	Versions []MediaVersion `gorm:"foreignKey:ImageID"`
}
