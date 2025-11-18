package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// MenuHandler کنترلر مدیریت منوها
type MenuHandler struct {
	menuService *services.MenuService
}

// NewMenuHandler ایجاد نمونه جدید از کنترلر منو
func NewMenuHandler(menuService *services.MenuService) *MenuHandler {
	return &MenuHandler{
		menuService: menuService,
	}
}

// ==================== Menu Operations ====================

// CreateMenu ایجاد منوی جدید
func (h *MenuHandler) CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.CreateMenu(c.Request.Context(), &menu); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد منو", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    menu,
		"message": "منو با موفقیت ایجاد شد",
	})
}

// GetMenuByID دریافت منو بر اساس ID
func (h *MenuHandler) GetMenuByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه منو نامعتبر است", err.Error()))
		return
	}

	menu, err := h.menuService.GetMenuByID(c.Request.Context(), uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "منو یافت نشد", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    menu,
		"message": "منو با موفقیت دریافت شد",
	})
}

// GetMenuBySlug دریافت منو بر اساس slug
func (h *MenuHandler) GetMenuBySlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه منو الزامی است", ""))
		return
	}

	menu, err := h.menuService.GetMenuBySlug(c.Request.Context(), slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "منو یافت نشد", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    menu,
		"message": "منو با موفقیت دریافت شد",
	})
}

// GetAllMenus دریافت تمام منوها
func (h *MenuHandler) GetAllMenus(c *gin.Context) {
	menus, err := h.menuService.GetAllMenus(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت منوها", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    menus,
		"message": "منوها با موفقیت دریافت شدند",
	})
}

// GetEnabledMenus دریافت منوهای فعال
func (h *MenuHandler) GetEnabledMenus(c *gin.Context) {
	menus, err := h.menuService.GetEnabledMenus(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت منوهای فعال", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    menus,
		"message": "منوهای فعال با موفقیت دریافت شدند",
	})
}

// UpdateMenu به‌روزرسانی منو
func (h *MenuHandler) UpdateMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه منو نامعتبر است", err.Error()))
		return
	}

	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	menu.ID = uint(id)
	if err := h.menuService.UpdateMenu(c.Request.Context(), &menu); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    menu,
		"message": "منو با موفقیت به‌روزرسانی شد",
	})
}

// DeleteMenu حذف منو
func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه منو نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.DeleteMenu(c.Request.Context(), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "منو با موفقیت حذف شد",
	})
}

// ==================== MenuItem Operations ====================

// CreateMenuItem ایجاد آیتم منوی جدید
func (h *MenuHandler) CreateMenuItem(c *gin.Context) {
	var item models.MenuItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.CreateMenuItem(c.Request.Context(), &item); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد آیتم منو", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    item,
		"message": "آیتم منو با موفقیت ایجاد شد",
	})
}

// GetMenuItemByID دریافت آیتم منو بر اساس ID
func (h *MenuHandler) GetMenuItemByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه آیتم منو نامعتبر است", err.Error()))
		return
	}

	item, err := h.menuService.GetMenuItemByID(c.Request.Context(), uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "آیتم منو یافت نشد", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    item,
		"message": "آیتم منو با موفقیت دریافت شد",
	})
}

// GetMenuItemsByMenuID دریافت آیتم‌های منو بر اساس ID منو
func (h *MenuHandler) GetMenuItemsByMenuID(c *gin.Context) {
	menuIDStr := c.Param("menuId")
	menuID, err := strconv.ParseUint(menuIDStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه منو نامعتبر است", err.Error()))
		return
	}

	items, err := h.menuService.GetMenuItemsByMenuID(c.Request.Context(), uint(menuID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت آیتم‌های منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    items,
		"message": "آیتم‌های منو با موفقیت دریافت شدند",
	})
}

// UpdateMenuItem به‌روزرسانی آیتم منو
func (h *MenuHandler) UpdateMenuItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه آیتم منو نامعتبر است", err.Error()))
		return
	}

	var item models.MenuItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	item.ID = uint(id)
	if err := h.menuService.UpdateMenuItem(c.Request.Context(), &item); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی آیتم منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    item,
		"message": "آیتم منو با موفقیت به‌روزرسانی شد",
	})
}

// DeleteMenuItem حذف آیتم منو
func (h *MenuHandler) DeleteMenuItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه آیتم منو نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.DeleteMenuItem(c.Request.Context(), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف آیتم منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "آیتم منو با موفقیت حذف شد",
	})
}

// ReorderMenuItems تغییر ترتیب آیتم‌های منو
func (h *MenuHandler) ReorderMenuItems(c *gin.Context) {
	menuIDStr := c.Param("menuId")
	menuID, err := strconv.ParseUint(menuIDStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه منو نامعتبر است", err.Error()))
		return
	}

	var itemOrders []struct {
		ID    uint `json:"id"`
		Order int  `json:"order"`
	}
	if err := c.ShouldBindJSON(&itemOrders); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.ReorderMenuItems(c.Request.Context(), uint(menuID), itemOrders); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در تغییر ترتیب آیتم‌های منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ترتیب آیتم‌های منو با موفقیت تغییر یافت",
	})
}

// ==================== MenuColumn Operations ====================

// CreateMenuColumn ایجاد ستون منوی جدید
func (h *MenuHandler) CreateMenuColumn(c *gin.Context) {
	var column models.MenuColumn
	if err := c.ShouldBindJSON(&column); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.CreateMenuColumn(c.Request.Context(), &column); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد ستون منو", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    column,
		"message": "ستون منو با موفقیت ایجاد شد",
	})
}

// GetMenuColumnByID دریافت ستون منو بر اساس ID
func (h *MenuHandler) GetMenuColumnByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه ستون منو نامعتبر است", err.Error()))
		return
	}

	column, err := h.menuService.GetMenuColumnByID(c.Request.Context(), uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "ستون منو یافت نشد", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    column,
		"message": "ستون منو با موفقیت دریافت شد",
	})
}

// GetMenuColumnsByMenuID دریافت ستون‌های منو بر اساس ID منو
func (h *MenuHandler) GetMenuColumnsByMenuID(c *gin.Context) {
	menuIDStr := c.Param("menuId")
	menuID, err := strconv.ParseUint(menuIDStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه منو نامعتبر است", err.Error()))
		return
	}

	columns, err := h.menuService.GetMenuColumnsByMenuID(c.Request.Context(), uint(menuID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت ستون‌های منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    columns,
		"message": "ستون‌های منو با موفقیت دریافت شدند",
	})
}

// UpdateMenuColumn به‌روزرسانی ستون منو
func (h *MenuHandler) UpdateMenuColumn(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه ستون منو نامعتبر است", err.Error()))
		return
	}

	var column models.MenuColumn
	if err := c.ShouldBindJSON(&column); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	column.ID = uint(id)
	if err := h.menuService.UpdateMenuColumn(c.Request.Context(), &column); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی ستون منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    column,
		"message": "ستون منو با موفقیت به‌روزرسانی شد",
	})
}

// DeleteMenuColumn حذف ستون منو
func (h *MenuHandler) DeleteMenuColumn(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه ستون منو نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.DeleteMenuColumn(c.Request.Context(), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف ستون منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ستون منو با موفقیت حذف شد",
	})
}

// ReorderMenuColumns تغییر ترتیب ستون‌های منو
func (h *MenuHandler) ReorderMenuColumns(c *gin.Context) {
	menuIDStr := c.Param("menuId")
	menuID, err := strconv.ParseUint(menuIDStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه منو نامعتبر است", err.Error()))
		return
	}

	var columnOrders []struct {
		ID    uint `json:"id"`
		Order int  `json:"order"`
	}
	if err := c.ShouldBindJSON(&columnOrders); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.ReorderMenuColumns(c.Request.Context(), uint(menuID), columnOrders); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در تغییر ترتیب ستون‌های منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ترتیب ستون‌های منو با موفقیت تغییر یافت",
	})
}

// ==================== MenuLocation Operations ====================

// CreateMenuLocation ایجاد جایگاه منوی جدید
func (h *MenuHandler) CreateMenuLocation(c *gin.Context) {
	var location models.MenuLocation
	if err := c.ShouldBindJSON(&location); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.CreateMenuLocation(c.Request.Context(), &location); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد جایگاه منو", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    location,
		"message": "جایگاه منو با موفقیت ایجاد شد",
	})
}

// GetMenuLocationByID دریافت جایگاه منو بر اساس ID
func (h *MenuHandler) GetMenuLocationByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه جایگاه منو نامعتبر است", err.Error()))
		return
	}

	location, err := h.menuService.GetMenuLocationByID(c.Request.Context(), uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "جایگاه منو یافت نشد", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    location,
		"message": "جایگاه منو با موفقیت دریافت شد",
	})
}

// GetMenuLocationBySlug دریافت جایگاه منو بر اساس slug
func (h *MenuHandler) GetMenuLocationBySlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه جایگاه منو الزامی است", ""))
		return
	}

	location, err := h.menuService.GetMenuLocationBySlug(c.Request.Context(), slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "جایگاه منو یافت نشد", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    location,
		"message": "جایگاه منو با موفقیت دریافت شد",
	})
}

// GetAllMenuLocations دریافت تمام جایگاه‌های منو
func (h *MenuHandler) GetAllMenuLocations(c *gin.Context) {
	locations, err := h.menuService.GetAllMenuLocations(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت جایگاه‌های منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    locations,
		"message": "جایگاه‌های منو با موفقیت دریافت شدند",
	})
}

// UpdateMenuLocation به‌روزرسانی جایگاه منو
func (h *MenuHandler) UpdateMenuLocation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه جایگاه منو نامعتبر است", err.Error()))
		return
	}

	var location models.MenuLocation
	if err := c.ShouldBindJSON(&location); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	location.ID = uint(id)
	if err := h.menuService.UpdateMenuLocation(c.Request.Context(), &location); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی جایگاه منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    location,
		"message": "جایگاه منو با موفقیت به‌روزرسانی شد",
	})
}

// DeleteMenuLocation حذف جایگاه منو
func (h *MenuHandler) DeleteMenuLocation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه جایگاه منو نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.DeleteMenuLocation(c.Request.Context(), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف جایگاه منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "جایگاه منو با موفقیت حذف شد",
	})
}

// ==================== MenuAssignment Operations ====================

// CreateMenuAssignment ایجاد تخصیص منو به جایگاه
func (h *MenuHandler) CreateMenuAssignment(c *gin.Context) {
	var assignment models.MenuAssignment
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.CreateMenuAssignment(c.Request.Context(), &assignment); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد تخصیص منو", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    assignment,
		"message": "تخصیص منو با موفقیت ایجاد شد",
	})
}

// GetMenuAssignmentByID دریافت تخصیص منو بر اساس ID
func (h *MenuHandler) GetMenuAssignmentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه تخصیص منو نامعتبر است", err.Error()))
		return
	}

	assignment, err := h.menuService.GetMenuAssignmentByID(c.Request.Context(), uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "تخصیص منو یافت نشد", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    assignment,
		"message": "تخصیص منو با موفقیت دریافت شد",
	})
}

// GetMenuAssignmentsByLocationID دریافت تخصیص‌های منو بر اساس ID جایگاه
func (h *MenuHandler) GetMenuAssignmentsByLocationID(c *gin.Context) {
	locationIDStr := c.Param("locationId")
	locationID, err := strconv.ParseUint(locationIDStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه جایگاه منو نامعتبر است", err.Error()))
		return
	}

	assignments, err := h.menuService.GetMenuAssignmentsByLocationID(c.Request.Context(), uint(locationID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تخصیص‌های منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    assignments,
		"message": "تخصیص‌های منو با موفقیت دریافت شدند",
	})
}

// GetMenuAssignmentsByMenuID دریافت تخصیص‌های منو بر اساس ID منو
func (h *MenuHandler) GetMenuAssignmentsByMenuID(c *gin.Context) {
	menuIDStr := c.Param("menuId")
	menuID, err := strconv.ParseUint(menuIDStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه منو نامعتبر است", err.Error()))
		return
	}

	assignments, err := h.menuService.GetMenuAssignmentsByMenuID(c.Request.Context(), uint(menuID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تخصیص‌های منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    assignments,
		"message": "تخصیص‌های منو با موفقیت دریافت شدند",
	})
}

// UpdateMenuAssignment به‌روزرسانی تخصیص منو
func (h *MenuHandler) UpdateMenuAssignment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه تخصیص منو نامعتبر است", err.Error()))
		return
	}

	var assignment models.MenuAssignment
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	assignment.ID = uint(id)
	if err := h.menuService.UpdateMenuAssignment(c.Request.Context(), &assignment); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی تخصیص منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    assignment,
		"message": "تخصیص منو با موفقیت به‌روزرسانی شد",
	})
}

// DeleteMenuAssignment حذف تخصیص منو
func (h *MenuHandler) DeleteMenuAssignment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه تخصیص منو نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.DeleteMenuAssignment(c.Request.Context(), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف تخصیص منو", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "تخصیص منو با موفقیت حذف شد",
	})
}

// AssignMenuToLocation تخصیص منو به جایگاه
func (h *MenuHandler) AssignMenuToLocation(c *gin.Context) {
	var request struct {
		MenuID     uint `json:"menu_id" binding:"required"`
		LocationID uint `json:"location_id" binding:"required"`
		Order      int  `json:"order"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.AssignMenuToLocation(c.Request.Context(), request.MenuID, request.LocationID, request.Order); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در تخصیص منو به جایگاه", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "منو با موفقیت به جایگاه تخصیص داده شد",
	})
}

// UnassignMenuFromLocation حذف تخصیص منو از جایگاه
func (h *MenuHandler) UnassignMenuFromLocation(c *gin.Context) {
	var request struct {
		MenuID     uint `json:"menu_id" binding:"required"`
		LocationID uint `json:"location_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	if err := h.menuService.UnassignMenuFromLocation(c.Request.Context(), request.MenuID, request.LocationID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف تخصیص منو از جایگاه", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "تخصیص منو از جایگاه با موفقیت حذف شد",
	})
}

// GetMenusByLocation دریافت منوهای تخصیص داده شده به یک جایگاه
func (h *MenuHandler) GetMenusByLocation(c *gin.Context) {
	locationSlug := c.Param("locationSlug")
	if locationSlug == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه جایگاه منو الزامی است", ""))
		return
	}

	menus, err := h.menuService.GetMenusByLocation(c.Request.Context(), locationSlug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت منوهای جایگاه", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    menus,
		"message": "منوهای جایگاه با موفقیت دریافت شدند",
	})
}

// GetMenuContentPages دریافت لیست برگه‌ها برای منو
func (h *MenuHandler) GetMenuContentPages(c *gin.Context) {
	// دریافت برگه‌های فعال
	pages, err := h.menuService.GetMenuContentPages(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در دریافت برگه‌ها",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "برگه‌ها با موفقیت دریافت شدند",
		"data":    pages,
	})
}

// GetMenuContentPosts دریافت لیست نوشته‌ها برای منو
func (h *MenuHandler) GetMenuContentPosts(c *gin.Context) {
	// دریافت نوشته‌های فعال
	posts, err := h.menuService.GetMenuContentPosts(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در دریافت نوشته‌ها",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "نوشته‌ها با موفقیت دریافت شدند",
		"data":    posts,
	})
}

// GetMenuContentCategories دریافت لیست دسته‌ها برای منو
func (h *MenuHandler) GetMenuContentCategories(c *gin.Context) {
	// دریافت دسته‌های فعال
	categories, err := h.menuService.GetMenuContentCategories(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در دریافت دسته‌ها",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "دسته‌ها با موفقیت دریافت شدند",
		"data":    categories,
	})
}

// GetMenuContentProductCategories دریافت لیست دسته‌های محصولات برای منو
func (h *MenuHandler) GetMenuContentProductCategories(c *gin.Context) {
	// دریافت دسته‌های محصولات فعال
	categories, err := h.menuService.GetMenuContentProductCategories(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در دریافت دسته‌های محصولات",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "دسته‌های محصولات با موفقیت دریافت شدند",
		"data":    categories,
	})
}
