package repository

import (
	"context"
	"errors"
	"sort"

	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// MenuRepository رابط برای عملیات دیتابیس منو
type MenuRepository struct {
	db *gorm.DB
}

// NewMenuRepository ایجاد نمونه جدید از repository منو
func NewMenuRepository(db *gorm.DB) *MenuRepository {
	// فعال‌سازی Statement Cache برای استفاده از Prepared Statements در کوئری‌های پرتردد
	return &MenuRepository{db: db.Session(&gorm.Session{PrepareStmt: true})}
}

// ==================== Menu Operations ====================

// CreateMenu ایجاد منوی جدید
func (r *MenuRepository) CreateMenu(ctx context.Context, menu *models.Menu) error {
	return r.db.WithContext(ctx).Create(menu).Error
}

// GetMenuByID دریافت منو بر اساس ID
func (r *MenuRepository) GetMenuByID(ctx context.Context, id uint) (*models.Menu, error) {
	var menu models.Menu
	err := r.db.WithContext(ctx).Preload("Items").Preload("Columns").Preload("Assignments").First(&menu, id).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

// GetMenuBySlug دریافت منو بر اساس slug
func (r *MenuRepository) GetMenuBySlug(ctx context.Context, slug string) (*models.Menu, error) {
	var menu models.Menu
	err := r.db.WithContext(ctx).Preload("Items").Preload("Columns").Preload("Assignments").Where("slug = ?", slug).First(&menu).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

// GetAllMenus دریافت تمام منوها
func (r *MenuRepository) GetAllMenus(ctx context.Context) ([]models.Menu, error) {
	// ابتدا فقط ID ها را بگیریم (کوئری سبک و قابل کش)
	var ids []uint
	if err := r.db.WithContext(ctx).Table("menus").Order("\"order\" ASC").Pluck("id", &ids).Error; err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return []models.Menu{}, nil
	}
	// سپس با یک کوئری جدا، منوها را با Preload ها بارگذاری می‌کنیم
	var menus []models.Menu
	if err := r.db.WithContext(ctx).
		Preload("Items").Preload("Columns").Preload("Assignments").
		Where("id IN ?", ids).Find(&menus).Error; err != nil {
		return nil, err
	}
	// مرتب‌سازی نتایج مطابق ترتیب ids
	index := make(map[uint]int, len(ids))
	for i, id := range ids {
		index[id] = i
	}
	sort.SliceStable(menus, func(i, j int) bool { return index[menus[i].ID] < index[menus[j].ID] })
	return menus, nil
}

// GetEnabledMenus دریافت منوهای فعال
func (r *MenuRepository) GetEnabledMenus(ctx context.Context) ([]models.Menu, error) {
	// رویکرد دو مرحله‌ای: ابتدا ID های منوهای فعال
	var ids []uint
	if err := r.db.WithContext(ctx).Table("menus").Where("enabled = ?", true).Order("\"order\" ASC").Pluck("id", &ids).Error; err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return []models.Menu{}, nil
	}
	// سپس بارگذاری کامل با Preload
	var menus []models.Menu
	if err := r.db.WithContext(ctx).
		Preload("Items").Preload("Columns").Preload("Assignments").
		Where("id IN ?", ids).Find(&menus).Error; err != nil {
		return nil, err
	}
	index := make(map[uint]int, len(ids))
	for i, id := range ids {
		index[id] = i
	}
	sort.SliceStable(menus, func(i, j int) bool { return index[menus[i].ID] < index[menus[j].ID] })
	return menus, nil
}

// UpdateMenu به‌روزرسانی منو
func (r *MenuRepository) UpdateMenu(ctx context.Context, menu *models.Menu) error {
	return r.db.WithContext(ctx).Save(menu).Error
}

// DeleteMenu حذف منو
func (r *MenuRepository) DeleteMenu(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Menu{}, id).Error
}

// ==================== MenuItem Operations ====================

// CreateMenuItem ایجاد آیتم منوی جدید
func (r *MenuRepository) CreateMenuItem(ctx context.Context, item *models.MenuItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

// GetMenuItemByID دریافت آیتم منو بر اساس ID
func (r *MenuRepository) GetMenuItemByID(ctx context.Context, id uint) (*models.MenuItem, error) {
	var item models.MenuItem
	err := r.db.WithContext(ctx).Preload("Children").First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetMenuItemsByMenuID دریافت آیتم‌های منو بر اساس ID منو
func (r *MenuRepository) GetMenuItemsByMenuID(ctx context.Context, menuID uint) ([]models.MenuItem, error) {
	var items []models.MenuItem
	err := r.db.WithContext(ctx).Where("menu_id = ? AND parent_id IS NULL", menuID).Order("\"order\" ASC").Find(&items).Error
	if err != nil {
		return nil, err
	}

	// بارگذاری زیرمنوها
	for i := range items {
		if err := r.loadChildren(ctx, &items[i]); err != nil {
			return nil, err
		}
	}

	return items, nil
}

// loadChildren بارگذاری زیرمنوها به صورت بازگشتی
func (r *MenuRepository) loadChildren(ctx context.Context, item *models.MenuItem) error {
	var children []models.MenuItem
	err := r.db.WithContext(ctx).Where("parent_id = ?", item.ID).Order("\"order\" ASC").Find(&children).Error
	if err != nil {
		return err
	}

	item.Children = children

	// بارگذاری زیرمنوهای زیرمنوها
	for i := range children {
		if err := r.loadChildren(ctx, &children[i]); err != nil {
			return err
		}
	}

	return nil
}

// UpdateMenuItem به‌روزرسانی آیتم منو
func (r *MenuRepository) UpdateMenuItem(ctx context.Context, item *models.MenuItem) error {
	return r.db.WithContext(ctx).Save(item).Error
}

// DeleteMenuItem حذف آیتم منو
func (r *MenuRepository) DeleteMenuItem(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.MenuItem{}, id).Error
}

// ReorderMenuItems تغییر ترتیب آیتم‌های منو
func (r *MenuRepository) ReorderMenuItems(ctx context.Context, menuID uint, itemOrders []struct {
	ID    uint `json:"id"`
	Order int  `json:"order"`
}) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, itemOrder := range itemOrders {
			if err := tx.Model(&models.MenuItem{}).Where("id = ? AND menu_id = ?", itemOrder.ID, menuID).Update("\"order\"", itemOrder.Order).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// ==================== MenuColumn Operations ====================

// CreateMenuColumn ایجاد ستون منوی جدید
func (r *MenuRepository) CreateMenuColumn(ctx context.Context, column *models.MenuColumn) error {
	return r.db.WithContext(ctx).Create(column).Error
}

// GetMenuColumnByID دریافت ستون منو بر اساس ID
func (r *MenuRepository) GetMenuColumnByID(ctx context.Context, id uint) (*models.MenuColumn, error) {
	var column models.MenuColumn
	err := r.db.WithContext(ctx).Preload("Items").First(&column, id).Error
	if err != nil {
		return nil, err
	}
	return &column, nil
}

// GetMenuColumnsByMenuID دریافت ستون‌های منو بر اساس ID منو
func (r *MenuRepository) GetMenuColumnsByMenuID(ctx context.Context, menuID uint) ([]models.MenuColumn, error) {
	var columns []models.MenuColumn
	err := r.db.WithContext(ctx).Preload("Items").Where("menu_id = ?", menuID).Order("\"order\" ASC").Find(&columns).Error
	return columns, err
}

// UpdateMenuColumn به‌روزرسانی ستون منو
func (r *MenuRepository) UpdateMenuColumn(ctx context.Context, column *models.MenuColumn) error {
	return r.db.WithContext(ctx).Save(column).Error
}

// DeleteMenuColumn حذف ستون منو
func (r *MenuRepository) DeleteMenuColumn(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.MenuColumn{}, id).Error
}

// ReorderMenuColumns تغییر ترتیب ستون‌های منو
func (r *MenuRepository) ReorderMenuColumns(ctx context.Context, menuID uint, columnOrders []struct {
	ID    uint `json:"id"`
	Order int  `json:"order"`
}) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, columnOrder := range columnOrders {
			if err := tx.Model(&models.MenuColumn{}).Where("id = ? AND menu_id = ?", columnOrder.ID, menuID).Update("\"order\"", columnOrder.Order).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// ==================== MenuLocation Operations ====================

// CreateMenuLocation ایجاد جایگاه منوی جدید
func (r *MenuRepository) CreateMenuLocation(ctx context.Context, location *models.MenuLocation) error {
	return r.db.WithContext(ctx).Create(location).Error
}

// GetMenuLocationByID دریافت جایگاه منو بر اساس ID
func (r *MenuRepository) GetMenuLocationByID(ctx context.Context, id uint) (*models.MenuLocation, error) {
	var location models.MenuLocation
	err := r.db.WithContext(ctx).Preload("Assignments.Menu").First(&location, id).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

// GetMenuLocationBySlug دریافت جایگاه منو بر اساس slug
func (r *MenuRepository) GetMenuLocationBySlug(ctx context.Context, slug string) (*models.MenuLocation, error) {
	var location models.MenuLocation
	err := r.db.WithContext(ctx).Preload("Assignments.Menu").Where("slug = ?", slug).First(&location).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

// GetAllMenuLocations دریافت تمام جایگاه‌های منو
func (r *MenuRepository) GetAllMenuLocations(ctx context.Context) ([]models.MenuLocation, error) {
	var locations []models.MenuLocation
	err := r.db.WithContext(ctx).Preload("Assignments.Menu").Order("\"order\" ASC").Find(&locations).Error
	return locations, err
}

// UpdateMenuLocation به‌روزرسانی جایگاه منو
func (r *MenuRepository) UpdateMenuLocation(ctx context.Context, location *models.MenuLocation) error {
	return r.db.WithContext(ctx).Save(location).Error
}

// DeleteMenuLocation حذف جایگاه منو
func (r *MenuRepository) DeleteMenuLocation(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.MenuLocation{}, id).Error
}

// ==================== MenuAssignment Operations ====================

// CreateMenuAssignment ایجاد تخصیص منو به جایگاه
func (r *MenuRepository) CreateMenuAssignment(ctx context.Context, assignment *models.MenuAssignment) error {
	return r.db.WithContext(ctx).Create(assignment).Error
}

// GetMenuAssignmentByID دریافت تخصیص منو بر اساس ID
func (r *MenuRepository) GetMenuAssignmentByID(ctx context.Context, id uint) (*models.MenuAssignment, error) {
	var assignment models.MenuAssignment
	err := r.db.WithContext(ctx).Preload("Menu").Preload("Location").First(&assignment, id).Error
	if err != nil {
		return nil, err
	}
	return &assignment, nil
}

// GetMenuAssignmentsByLocationID دریافت تخصیص‌های منو بر اساس ID جایگاه
func (r *MenuRepository) GetMenuAssignmentsByLocationID(ctx context.Context, locationID uint) ([]models.MenuAssignment, error) {
	var assignments []models.MenuAssignment
	err := r.db.WithContext(ctx).Preload("Menu").Preload("Location").Where("location_id = ? AND enabled = ?", locationID, true).Order("\"order\" ASC").Find(&assignments).Error
	return assignments, err
}

// GetMenuAssignmentsByMenuID دریافت تخصیص‌های منو بر اساس ID منو
func (r *MenuRepository) GetMenuAssignmentsByMenuID(ctx context.Context, menuID uint) ([]models.MenuAssignment, error) {
	var assignments []models.MenuAssignment
	err := r.db.WithContext(ctx).Preload("Menu").Preload("Location").Where("menu_id = ?", menuID).Order("\"order\" ASC").Find(&assignments).Error
	return assignments, err
}

// UpdateMenuAssignment به‌روزرسانی تخصیص منو
func (r *MenuRepository) UpdateMenuAssignment(ctx context.Context, assignment *models.MenuAssignment) error {
	return r.db.WithContext(ctx).Save(assignment).Error
}

// DeleteMenuAssignment حذف تخصیص منو
func (r *MenuRepository) DeleteMenuAssignment(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.MenuAssignment{}, id).Error
}

// AssignMenuToLocation تخصیص منو به جایگاه
func (r *MenuRepository) AssignMenuToLocation(ctx context.Context, menuID, locationID uint, order int) error {
	// بررسی اینکه آیا این منو قبلاً در این جایگاه تخصیص داده شده یا نه
	var existingAssignment models.MenuAssignment
	err := r.db.WithContext(ctx).Where("menu_id = ? AND location_id = ?", menuID, locationID).First(&existingAssignment).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// ایجاد تخصیص جدید
		assignment := &models.MenuAssignment{
			MenuID:     menuID,
			LocationID: locationID,
			Order:      order,
			Enabled:    true,
		}
		return r.db.WithContext(ctx).Create(assignment).Error
	} else if err != nil {
		return err
	} else {
		// به‌روزرسانی تخصیص موجود
		existingAssignment.Order = order
		existingAssignment.Enabled = true
		return r.db.WithContext(ctx).Save(&existingAssignment).Error
	}
}

// UnassignMenuFromLocation حذف تخصیص منو از جایگاه
func (r *MenuRepository) UnassignMenuFromLocation(ctx context.Context, menuID, locationID uint) error {
	return r.db.WithContext(ctx).Where("menu_id = ? AND location_id = ?", menuID, locationID).Delete(&models.MenuAssignment{}).Error
}

// GetMenusByLocation دریافت منوهای تخصیص داده شده به یک جایگاه
func (r *MenuRepository) GetMenusByLocation(ctx context.Context, locationSlug string) ([]models.Menu, error) {
	var menus []models.Menu
	err := r.db.WithContext(ctx).
		Joins("JOIN menu_assignments ON menus.id = menu_assignments.menu_id").
		Joins("JOIN menu_locations ON menu_assignments.location_id = menu_locations.id").
		Where("menu_locations.slug = ? AND menu_assignments.enabled = ? AND menus.enabled = ?", locationSlug, true, true).
		Preload("Items").
		Preload("Columns").
		Order("menu_assignments.\"order\" ASC").
		Find(&menus).Error

	if err != nil {
		return nil, err
	}

	// بارگذاری آیتم‌های منو به صورت سلسله‌مراتبی
	for i := range menus {
		items, err := r.GetMenuItemsByMenuID(ctx, menus[i].ID)
		if err != nil {
			return nil, err
		}
		menus[i].Items = items
	}

	return menus, nil
}

// ==================== Menu Content Operations ====================

// GetMenuContentPages دریافت لیست برگه‌ها برای منو
// فعلاً از posts استفاده می‌کنیم چون جدول pages وجود ندارد
func (r *MenuRepository) GetMenuContentPages(ctx context.Context) ([]struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}, error) {
	var pages []struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
		Slug  string `json:"slug"`
	}

	err := r.db.WithContext(ctx).Table("posts").
		Select("id, title, slug").
		Where("status = ?", "published").
		Order("created_at DESC").
		Find(&pages).Error

	return pages, err
}

// GetMenuContentPosts دریافت لیست نوشته‌ها برای منو
func (r *MenuRepository) GetMenuContentPosts(ctx context.Context) ([]struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}, error) {
	var posts []struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
		Slug  string `json:"slug"`
	}

	err := r.db.WithContext(ctx).Table("posts").
		Select("id, title, slug").
		Where("status = ?", "published").
		Order("created_at DESC").
		Find(&posts).Error

	return posts, err
}

// GetMenuContentCategories دریافت لیست دسته‌ها برای منو
func (r *MenuRepository) GetMenuContentCategories(ctx context.Context) ([]struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}, error) {
	var categories []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}

	err := r.db.WithContext(ctx).Table("categories").
		Select("id, name, slug").
		Where("published = ?", true).
		Order("name ASC").
		Find(&categories).Error

	return categories, err
}

// GetMenuContentProductCategories دریافت لیست دسته‌های محصولات برای منو
// فعلاً از categories استفاده می‌کنیم چون جدول product_categories وجود ندارد
func (r *MenuRepository) GetMenuContentProductCategories(ctx context.Context) ([]struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}, error) {
	var categories []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}

	err := r.db.WithContext(ctx).Table("categories").
		Select("id, name, slug").
		Where("published = ?", true).
		Order("name ASC").
		Find(&categories).Error

	return categories, err
}
