package models

import (
	"time"
)

// Menu مدل اصلی منو - کاملاً مستقل و قابل استفاده مجدد
type Menu struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"size:100;not null"`             // نام منو (مثل "منوی اصلی"، "منوی فوتر")
	Slug        string `json:"slug" gorm:"size:100;uniqueIndex;not null"` // شناسه یکتا (مثل "main-menu", "footer-menu")
	Description string `json:"description" gorm:"size:500"`               // توضیحات منو
	Type        string `json:"type" gorm:"size:50;default:'standard'"`    // نوع منو: standard, mega, vertical, horizontal
	Theme       string `json:"theme" gorm:"size:50;default:'default'"`    // تم منو: default, modern, minimal, classic
	MaxDepth    int    `json:"max_depth" gorm:"default:3"`                // حداکثر عمق منو
	Enabled     bool   `json:"enabled" gorm:"default:true;index"`
	ShowIcons   bool   `json:"show_icons" gorm:"default:true"`  // نمایش آیکون‌ها
	ShowBadges  bool   `json:"show_badges" gorm:"default:true"` // نمایش badge ها
	Order       int    `json:"order" gorm:"default:0;index"`

	// تنظیمات مگا منو
	IsMegaMenu  bool   `json:"is_mega_menu" gorm:"default:false"`         // آیا منوی مگا است
	MegaWidth   string `json:"mega_width" gorm:"size:20;default:'full'"`  // عرض مگا منو: full, container, custom
	MegaHeight  string `json:"mega_height" gorm:"size:20;default:'auto'"` // ارتفاع مگا منو
	MegaColumns int    `json:"mega_columns" gorm:"default:2"`             // تعداد ستون‌های مگا منو

	// تنظیمات نمایش
	ShowOnMobile  bool `json:"show_on_mobile" gorm:"default:true"`  // نمایش در موبایل
	ShowOnDesktop bool `json:"show_on_desktop" gorm:"default:true"` // نمایش در دسکتاپ
	ShowOnTablet  bool `json:"show_on_tablet" gorm:"default:true"`  // نمایش در تبلت

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"` // حذف شد

	// روابط
	Items       []MenuItem       `json:"items" gorm:"foreignKey:MenuID;constraint:OnDelete:CASCADE"`
	Columns     []MenuColumn     `json:"columns" gorm:"foreignKey:MenuID;constraint:OnDelete:CASCADE"`
	Assignments []MenuAssignment `json:"assignments" gorm:"foreignKey:MenuID;constraint:OnDelete:CASCADE"`
}

// MenuItem آیتم‌های منو - کاملاً انعطاف‌پذیر
type MenuItem struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	MenuID       uint   `json:"menu_id" gorm:"not null;index"`
	ParentID     *uint  `json:"parent_id" gorm:"index"`                   // برای زیرمنوها
	ColumnID     *uint  `json:"column_id" gorm:"index"`                   // برای مگا منو - ستون مربوطه
	Title        string `json:"title" gorm:"size:100;not null"`           // عنوان آیتم
	Path         string `json:"path" gorm:"size:200"`                     // مسیر یا URL
	Icon         string `json:"icon" gorm:"size:100"`                     // کلاس آیکون
	Badge        string `json:"badge" gorm:"size:50"`                     // متن badge
	BadgeColor   string `json:"badge_color" gorm:"size:20;default:'red'"` // رنگ badge
	Enabled      bool   `json:"enabled" gorm:"default:true;index"`
	Order        int    `json:"order" gorm:"default:0;index"`
	OpenInNewTab bool   `json:"open_in_new_tab" gorm:"default:false"` // باز شدن در تب جدید

	// نوع آیتم و محتوای آن
	ItemType   string `json:"item_type" gorm:"size:20;default:'link'"` // link, category, page, product, custom, separator
	TargetID   *uint  `json:"target_id" gorm:"index"`                  // ID آیتم هدف (مثل ID دسته‌بندی یا صفحه)
	TargetType string `json:"target_type" gorm:"size:50"`              // نوع هدف: category, page, product, post

	// محتوای اضافی برای مگا منو
	ImageURL    string `json:"image_url" gorm:"size:500"`     // تصویر آیتم
	Description string `json:"description" gorm:"size:300"`   // توضیحات آیتم
	Featured    bool   `json:"featured" gorm:"default:false"` // آیا آیتم ویژه است

	// تنظیمات نمایش
	ShowOnMobile  bool `json:"show_on_mobile" gorm:"default:true"`  // نمایش در موبایل
	ShowOnDesktop bool `json:"show_on_desktop" gorm:"default:true"` // نمایش در دسکتاپ
	ShowOnTablet  bool `json:"show_on_tablet" gorm:"default:true"`  // نمایش در تبلت

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"` // حذف شد

	// روابط
	Menu     Menu        `json:"menu" gorm:"foreignKey:MenuID"`
	Parent   *MenuItem   `json:"parent" gorm:"foreignKey:ParentID"`
	Children []MenuItem  `json:"children" gorm:"foreignKey:ParentID"`
	Column   *MenuColumn `json:"column" gorm:"foreignKey:ColumnID"`
}

// MenuColumn ستون‌های مگا منو
type MenuColumn struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	MenuID      uint   `json:"menu_id" gorm:"not null;index"`
	Title       string `json:"title" gorm:"size:100"`               // عنوان ستون
	Description string `json:"description" gorm:"size:300"`         // توضیحات ستون
	Width       string `json:"width" gorm:"size:20;default:'auto'"` // عرض ستون: auto, 1/4, 1/3, 1/2, full
	Order       int    `json:"order" gorm:"default:0;index"`
	Enabled     bool   `json:"enabled" gorm:"default:true"`
	ColumnType  string `json:"column_type" gorm:"size:20;default:'items'"` // items, categories, pages, products, custom

	// تنظیمات نمایش
	ShowOnMobile  bool `json:"show_on_mobile" gorm:"default:true"`  // نمایش در موبایل
	ShowOnDesktop bool `json:"show_on_desktop" gorm:"default:true"` // نمایش در دسکتاپ
	ShowOnTablet  bool `json:"show_on_tablet" gorm:"default:true"`  // نمایش در تبلت

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"` // حذف شد

	// روابط
	Menu  Menu       `json:"menu" gorm:"foreignKey:MenuID"`
	Items []MenuItem `json:"items" gorm:"foreignKey:ColumnID"`
}

// MenuLocation جایگاه‌های منو - مکان‌هایی که می‌توان منو را در آن‌ها قرار داد
type MenuLocation struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:100;not null;uniqueIndex"` // نام جایگاه (مثل "هدر اصلی")
	Slug        string    `json:"slug" gorm:"size:100;not null;uniqueIndex"` // شناسه جایگاه (مثل "header-main")
	Description string    `json:"description" gorm:"size:500"`               // توضیحات جایگاه
	Icon        string    `json:"icon" gorm:"size:100"`                      // آیکون جایگاه
	Enabled     bool      `json:"enabled" gorm:"default:true"`
	Order       int       `json:"order" gorm:"default:0"`
	MaxMenus    int       `json:"max_menus" gorm:"default:1"` // حداکثر تعداد منو در این جایگاه
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"` // حذف شد

	// روابط
	Assignments []MenuAssignment `json:"assignments" gorm:"foreignKey:LocationID"`
}

// MenuAssignment تخصیص منو به جایگاه - مثل وردپرس
type MenuAssignment struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	MenuID     uint      `json:"menu_id" gorm:"not null;index"`
	LocationID uint      `json:"location_id" gorm:"not null;index"`
	Order      int       `json:"order" gorm:"default:0;index"` // ترتیب نمایش در جایگاه
	Enabled    bool      `json:"enabled" gorm:"default:true"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	// DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"` // حذف شد

	// روابط
	Menu     Menu         `json:"menu" gorm:"foreignKey:MenuID"`
	Location MenuLocation `json:"location" gorm:"foreignKey:LocationID"`
}

// TableName نام جدول برای مدل‌ها
func (Menu) TableName() string {
	return "menus"
}

func (MenuItem) TableName() string {
	return "menu_items"
}

func (MenuColumn) TableName() string {
	return "menu_columns"
}

func (MenuLocation) TableName() string {
	return "menu_locations"
}

func (MenuAssignment) TableName() string {
	return "menu_assignments"
}
