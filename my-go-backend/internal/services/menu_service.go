package services

import (
	"context"
	"errors"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
)

// MenuService سرویس مدیریت منوها
type MenuService struct {
	menuRepo *repository.MenuRepository
}

// NewMenuService ایجاد نمونه جدید از سرویس منو
func NewMenuService(menuRepo *repository.MenuRepository) *MenuService {
	return &MenuService{
		menuRepo: menuRepo,
	}
}

// ==================== Menu Operations ====================

// CreateMenu ایجاد منوی جدید
func (s *MenuService) CreateMenu(ctx context.Context, menu *models.Menu) error {
	if menu.Name == "" {
		return errors.New("نام منو الزامی است")
	}
	if menu.Slug == "" {
		return errors.New("شناسه منو الزامی است")
	}
	return s.menuRepo.CreateMenu(ctx, menu)
}

// GetMenuByID دریافت منو بر اساس ID
func (s *MenuService) GetMenuByID(ctx context.Context, id uint) (*models.Menu, error) {
	return s.menuRepo.GetMenuByID(ctx, id)
}

// GetMenuBySlug دریافت منو بر اساس slug
func (s *MenuService) GetMenuBySlug(ctx context.Context, slug string) (*models.Menu, error) {
	return s.menuRepo.GetMenuBySlug(ctx, slug)
}

// GetAllMenus دریافت تمام منوها
func (s *MenuService) GetAllMenus(ctx context.Context) ([]models.Menu, error) {
	return s.menuRepo.GetAllMenus(ctx)
}

// GetEnabledMenus دریافت منوهای فعال
func (s *MenuService) GetEnabledMenus(ctx context.Context) ([]models.Menu, error) {
	return s.menuRepo.GetEnabledMenus(ctx)
}

// UpdateMenu به‌روزرسانی منو
func (s *MenuService) UpdateMenu(ctx context.Context, menu *models.Menu) error {
	if menu.Name == "" {
		return errors.New("نام منو الزامی است")
	}
	return s.menuRepo.UpdateMenu(ctx, menu)
}

// DeleteMenu حذف منو
func (s *MenuService) DeleteMenu(ctx context.Context, id uint) error {
	return s.menuRepo.DeleteMenu(ctx, id)
}

// ==================== MenuItem Operations ====================

// CreateMenuItem ایجاد آیتم منوی جدید
func (s *MenuService) CreateMenuItem(ctx context.Context, item *models.MenuItem) error {
	if item.Title == "" {
		return errors.New("عنوان آیتم الزامی است")
	}
	if item.MenuID == 0 {
		return errors.New("شناسه منو الزامی است")
	}
	return s.menuRepo.CreateMenuItem(ctx, item)
}

// GetMenuItemByID دریافت آیتم منو بر اساس ID
func (s *MenuService) GetMenuItemByID(ctx context.Context, id uint) (*models.MenuItem, error) {
	return s.menuRepo.GetMenuItemByID(ctx, id)
}

// GetMenuItemsByMenuID دریافت آیتم‌های منو بر اساس ID منو
func (s *MenuService) GetMenuItemsByMenuID(ctx context.Context, menuID uint) ([]models.MenuItem, error) {
	return s.menuRepo.GetMenuItemsByMenuID(ctx, menuID)
}

// UpdateMenuItem به‌روزرسانی آیتم منو
func (s *MenuService) UpdateMenuItem(ctx context.Context, item *models.MenuItem) error {
	if item.Title == "" {
		return errors.New("عنوان آیتم الزامی است")
	}
	return s.menuRepo.UpdateMenuItem(ctx, item)
}

// DeleteMenuItem حذف آیتم منو
func (s *MenuService) DeleteMenuItem(ctx context.Context, id uint) error {
	return s.menuRepo.DeleteMenuItem(ctx, id)
}

// ReorderMenuItems تغییر ترتیب آیتم‌های منو
func (s *MenuService) ReorderMenuItems(ctx context.Context, menuID uint, itemOrders []struct {
	ID    uint `json:"id"`
	Order int  `json:"order"`
}) error {
	return s.menuRepo.ReorderMenuItems(ctx, menuID, itemOrders)
}

// ==================== MenuColumn Operations ====================

// CreateMenuColumn ایجاد ستون منوی جدید
func (s *MenuService) CreateMenuColumn(ctx context.Context, column *models.MenuColumn) error {
	if column.Title == "" {
		return errors.New("عنوان ستون الزامی است")
	}
	if column.MenuID == 0 {
		return errors.New("شناسه منو الزامی است")
	}
	return s.menuRepo.CreateMenuColumn(ctx, column)
}

// GetMenuColumnByID دریافت ستون منو بر اساس ID
func (s *MenuService) GetMenuColumnByID(ctx context.Context, id uint) (*models.MenuColumn, error) {
	return s.menuRepo.GetMenuColumnByID(ctx, id)
}

// GetMenuColumnsByMenuID دریافت ستون‌های منو بر اساس ID منو
func (s *MenuService) GetMenuColumnsByMenuID(ctx context.Context, menuID uint) ([]models.MenuColumn, error) {
	return s.menuRepo.GetMenuColumnsByMenuID(ctx, menuID)
}

// UpdateMenuColumn به‌روزرسانی ستون منو
func (s *MenuService) UpdateMenuColumn(ctx context.Context, column *models.MenuColumn) error {
	if column.Title == "" {
		return errors.New("عنوان ستون الزامی است")
	}
	return s.menuRepo.UpdateMenuColumn(ctx, column)
}

// DeleteMenuColumn حذف ستون منو
func (s *MenuService) DeleteMenuColumn(ctx context.Context, id uint) error {
	return s.menuRepo.DeleteMenuColumn(ctx, id)
}

// ReorderMenuColumns تغییر ترتیب ستون‌های منو
func (s *MenuService) ReorderMenuColumns(ctx context.Context, menuID uint, columnOrders []struct {
	ID    uint `json:"id"`
	Order int  `json:"order"`
}) error {
	return s.menuRepo.ReorderMenuColumns(ctx, menuID, columnOrders)
}

// ==================== MenuLocation Operations ====================

// CreateMenuLocation ایجاد جایگاه منوی جدید
func (s *MenuService) CreateMenuLocation(ctx context.Context, location *models.MenuLocation) error {
	if location.Name == "" {
		return errors.New("نام جایگاه الزامی است")
	}
	if location.Slug == "" {
		return errors.New("شناسه جایگاه الزامی است")
	}
	return s.menuRepo.CreateMenuLocation(ctx, location)
}

// GetMenuLocationByID دریافت جایگاه منو بر اساس ID
func (s *MenuService) GetMenuLocationByID(ctx context.Context, id uint) (*models.MenuLocation, error) {
	return s.menuRepo.GetMenuLocationByID(ctx, id)
}

// GetMenuLocationBySlug دریافت جایگاه منو بر اساس slug
func (s *MenuService) GetMenuLocationBySlug(ctx context.Context, slug string) (*models.MenuLocation, error) {
	return s.menuRepo.GetMenuLocationBySlug(ctx, slug)
}

// GetAllMenuLocations دریافت تمام جایگاه‌های منو
func (s *MenuService) GetAllMenuLocations(ctx context.Context) ([]models.MenuLocation, error) {
	return s.menuRepo.GetAllMenuLocations(ctx)
}

// UpdateMenuLocation به‌روزرسانی جایگاه منو
func (s *MenuService) UpdateMenuLocation(ctx context.Context, location *models.MenuLocation) error {
	if location.Name == "" {
		return errors.New("نام جایگاه الزامی است")
	}
	return s.menuRepo.UpdateMenuLocation(ctx, location)
}

// DeleteMenuLocation حذف جایگاه منو
func (s *MenuService) DeleteMenuLocation(ctx context.Context, id uint) error {
	return s.menuRepo.DeleteMenuLocation(ctx, id)
}

// ==================== MenuAssignment Operations ====================

// CreateMenuAssignment ایجاد تخصیص منو به جایگاه
func (s *MenuService) CreateMenuAssignment(ctx context.Context, assignment *models.MenuAssignment) error {
	if assignment.MenuID == 0 {
		return errors.New("شناسه منو الزامی است")
	}
	if assignment.LocationID == 0 {
		return errors.New("شناسه جایگاه الزامی است")
	}
	return s.menuRepo.CreateMenuAssignment(ctx, assignment)
}

// GetMenuAssignmentByID دریافت تخصیص منو بر اساس ID
func (s *MenuService) GetMenuAssignmentByID(ctx context.Context, id uint) (*models.MenuAssignment, error) {
	return s.menuRepo.GetMenuAssignmentByID(ctx, id)
}

// GetMenuAssignmentsByLocationID دریافت تخصیص‌های منو بر اساس ID جایگاه
func (s *MenuService) GetMenuAssignmentsByLocationID(ctx context.Context, locationID uint) ([]models.MenuAssignment, error) {
	return s.menuRepo.GetMenuAssignmentsByLocationID(ctx, locationID)
}

// GetMenuAssignmentsByMenuID دریافت تخصیص‌های منو بر اساس ID منو
func (s *MenuService) GetMenuAssignmentsByMenuID(ctx context.Context, menuID uint) ([]models.MenuAssignment, error) {
	return s.menuRepo.GetMenuAssignmentsByMenuID(ctx, menuID)
}

// UpdateMenuAssignment به‌روزرسانی تخصیص منو
func (s *MenuService) UpdateMenuAssignment(ctx context.Context, assignment *models.MenuAssignment) error {
	return s.menuRepo.UpdateMenuAssignment(ctx, assignment)
}

// DeleteMenuAssignment حذف تخصیص منو
func (s *MenuService) DeleteMenuAssignment(ctx context.Context, id uint) error {
	return s.menuRepo.DeleteMenuAssignment(ctx, id)
}

// AssignMenuToLocation تخصیص منو به جایگاه
func (s *MenuService) AssignMenuToLocation(ctx context.Context, menuID, locationID uint, order int) error {
	return s.menuRepo.AssignMenuToLocation(ctx, menuID, locationID, order)
}

// UnassignMenuFromLocation حذف تخصیص منو از جایگاه
func (s *MenuService) UnassignMenuFromLocation(ctx context.Context, menuID, locationID uint) error {
	return s.menuRepo.UnassignMenuFromLocation(ctx, menuID, locationID)
}

// GetMenusByLocation دریافت منوهای تخصیص داده شده به یک جایگاه
func (s *MenuService) GetMenusByLocation(ctx context.Context, locationSlug string) ([]models.Menu, error) {
	return s.menuRepo.GetMenusByLocation(ctx, locationSlug)
}

// ==================== Menu Content Operations ====================

// GetMenuContentPages دریافت لیست برگه‌ها برای منو
func (s *MenuService) GetMenuContentPages(ctx context.Context) ([]struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}, error) {
	return s.menuRepo.GetMenuContentPages(ctx)
}

// GetMenuContentPosts دریافت لیست نوشته‌ها برای منو
func (s *MenuService) GetMenuContentPosts(ctx context.Context) ([]struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}, error) {
	return s.menuRepo.GetMenuContentPosts(ctx)
}

// GetMenuContentCategories دریافت لیست دسته‌ها برای منو
func (s *MenuService) GetMenuContentCategories(ctx context.Context) ([]struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}, error) {
	return s.menuRepo.GetMenuContentCategories(ctx)
}

// GetMenuContentProductCategories دریافت لیست دسته‌های محصولات برای منو
func (s *MenuService) GetMenuContentProductCategories(ctx context.Context) ([]struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}, error) {
	return s.menuRepo.GetMenuContentProductCategories(ctx)
}
