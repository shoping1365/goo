package repository

import (
	"context"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// DTO/Entity mirror of migration structs (local to repo)
type FraudCase struct {
	ID        uint   `gorm:"primaryKey"`
	OrderID   *uint  `gorm:"index"`
	UserID    *uint  `gorm:"index"`
	RiskScore int    `gorm:"index"`
	RiskLevel string `gorm:"index"`
	Status    string `gorm:"index"`
	Summary   string
	Metadata  datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ClosedAt  *time.Time
}

func (FraudCase) TableName() string { return "fraud_cases" }

type FraudEvent struct {
	ID        uint   `gorm:"primaryKey"`
	CaseID    uint   `gorm:"index"`
	EventType string `gorm:"index"`
	Weight    int
	Impact    int
	Details   datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt time.Time      `gorm:"index"`
}

func (FraudEvent) TableName() string { return "fraud_events" }

type FraudScoreLog struct {
	ID          uint `gorm:"primaryKey"`
	CaseID      uint `gorm:"index"`
	ScoreBefore int
	ScoreAfter  int
	Reason      datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt   time.Time      `gorm:"index"`
}

func (FraudScoreLog) TableName() string { return "fraud_scores" }

type FraudRule struct {
	ID          uint   `gorm:"primaryKey"`
	Key         string `gorm:"uniqueIndex"`
	Description string
	Weight      int
	Enabled     bool           `gorm:"index"`
	Params      datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (FraudRule) TableName() string { return "fraud_rules" }

type FraudListItem struct {
	ID        uint   `gorm:"primaryKey"`
	Kind      string `gorm:"index"`
	ValueHash string `gorm:"index"`
	Note      string
	ExpiresAt *time.Time `gorm:"index"`
	CreatedAt time.Time
}

func (FraudListItem) TableName() string { return "fraud_list" }

type FraudActionLog struct {
	ID        uint   `gorm:"primaryKey"`
	CaseID    uint   `gorm:"index"`
	Action    string `gorm:"index"`
	ActorID   *uint  `gorm:"index"`
	Note      string
	CreatedAt time.Time `gorm:"index"`
}

func (FraudActionLog) TableName() string { return "fraud_actions_log" }

// ================= Interfaces =================

type FraudCaseFilter struct {
	RiskLevel     *string
	Status        *string
	UserID        *uint
	OrderID       *uint
	Search        *string // by tracking code, phone, etc. (on joined tables)
	PaymentMethod *string
	AmountMin     *float64
	AmountMax     *float64
	From          *time.Time
	To            *time.Time
	Page          int
	PageSize      int
}

type FraudCaseRepositoryInterface interface {
	Create(ctx context.Context, c *FraudCase) error
	Update(ctx context.Context, c *FraudCase) error
	GetByID(ctx context.Context, id uint) (*FraudCase, error)
	GetByOrderID(ctx context.Context, orderID uint) (*FraudCase, error)
	List(ctx context.Context, f FraudCaseFilter) ([]FraudCase, int64, error)
	ListEnriched(ctx context.Context, f FraudCaseFilter) ([]AdminFraudCaseRow, int64, error)
	GetDetail(ctx context.Context, id uint) (*FraudCaseDetail, error)
}

type FraudEventRepositoryInterface interface {
	CreateMany(ctx context.Context, events []FraudEvent) error
	ListByCase(ctx context.Context, caseID uint) ([]FraudEvent, error)
}

type FraudRuleRepositoryInterface interface {
	ListActive(ctx context.Context) ([]FraudRule, error)
	GetByKey(ctx context.Context, key string) (*FraudRule, error)
	GetByID(ctx context.Context, id uint) (*FraudRule, error)
	ListAll(ctx context.Context) ([]FraudRule, error)
	Update(ctx context.Context, r *FraudRule) error
}

type FraudListRepositoryInterface interface {
	IsWhitelisted(ctx context.Context, kind, valueHash string) (bool, error)
	IsBlacklisted(ctx context.Context, kind, valueHash string) (bool, error)
	Upsert(ctx context.Context, item *FraudListItem) error
	Delete(ctx context.Context, kind, valueHash string) error
}

type FraudActionLogRepositoryInterface interface {
	Create(ctx context.Context, l *FraudActionLog) error
	ListByCase(ctx context.Context, caseID uint) ([]FraudActionLog, error)
}

// ================= Implementations =================
// تمام ریپازیتوری‌ها فقط منطق دسترسی به داده را شامل می‌شوند و هیچ منطق تجاری
// در این لایه قرار نمی‌گیرد. تمامی کوئری‌ها با پارامتر و ایندکس‌های مناسب انجام می‌شوند.

type FraudCaseRepository struct{ DB *gorm.DB }

func (r *FraudCaseRepository) Create(ctx context.Context, c *FraudCase) error {
	return r.DB.WithContext(ctx).Create(c).Error
}

func (r *FraudCaseRepository) Update(ctx context.Context, c *FraudCase) error {
	return r.DB.WithContext(ctx).Save(c).Error
}

func (r *FraudCaseRepository) GetByID(ctx context.Context, id uint) (*FraudCase, error) {
	var entity FraudCase
	err := r.DB.WithContext(ctx).First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *FraudCaseRepository) GetByOrderID(ctx context.Context, orderID uint) (*FraudCase, error) {
	var entity FraudCase
	err := r.DB.WithContext(ctx).Where("order_id = ?", orderID).Order("id DESC").First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *FraudCaseRepository) List(ctx context.Context, f FraudCaseFilter) ([]FraudCase, int64, error) {
	db := r.DB.WithContext(ctx).Model(&FraudCase{})
	if f.RiskLevel != nil && *f.RiskLevel != "" {
		db = db.Where("risk_level = ?", *f.RiskLevel)
	}
	if f.Status != nil && *f.Status != "" {
		db = db.Where("status = ?", *f.Status)
	}
	if f.UserID != nil {
		db = db.Where("user_id = ?", *f.UserID)
	}
	if f.OrderID != nil {
		db = db.Where("order_id = ?", *f.OrderID)
	}
	if f.From != nil {
		db = db.Where("created_at >= ?", *f.From)
	}
	if f.To != nil {
		db = db.Where("created_at <= ?", *f.To)
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	page := f.Page
	size := f.PageSize
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}
	var list []FraudCase
	err := db.Order("created_at DESC").Offset((page - 1) * size).Limit(size).Find(&list).Error
	return list, total, err
}

// AdminFraudCaseRow ردیف خروجی برای لیست ادمین با اطلاعات سفارش/کاربر
type AdminFraudCaseRow struct {
	ID            uint      `json:"id"`
	OrderID       *uint     `json:"orderId"`
	OrderNumber   *string   `json:"orderNumber"`
	CustomerName  *string   `json:"customerName"`
	CustomerPhone *string   `json:"customerPhone"`
	TotalAmount   *float64  `json:"totalAmount"`
	PaymentMethod *string   `json:"paymentMethod"`
	Status        *string   `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	RiskLevel     string    `json:"riskLevel"`
	FraudScore    int       `json:"fraudScore"`
}

// FraudCaseDetail ساختار دیتیل پرونده برای API /cases/:id
type FraudCaseDetail struct {
	Case    FraudCase        `json:"case"`
	Events  []FraudEvent     `json:"events"`
	Actions []FraudActionLog `json:"actions"`
}

func (r *FraudCaseRepository) ListEnriched(ctx context.Context, f FraudCaseFilter) ([]AdminFraudCaseRow, int64, error) {
	// ساخت شرط‌ها و پارامترها
	where := "WHERE 1=1"
	args := []any{}
	if f.RiskLevel != nil && *f.RiskLevel != "" {
		where += " AND fc.risk_level = ?"
		args = append(args, *f.RiskLevel)
	}
	if f.Status != nil && *f.Status != "" {
		where += " AND fc.status = ?"
		args = append(args, *f.Status)
	}
	if f.UserID != nil {
		where += " AND fc.user_id = ?"
		args = append(args, *f.UserID)
	}
	if f.OrderID != nil {
		where += " AND fc.order_id = ?"
		args = append(args, *f.OrderID)
	}
	if f.From != nil {
		where += " AND fc.created_at >= ?"
		args = append(args, *f.From)
	}
	if f.To != nil {
		where += " AND fc.created_at <= ?"
		args = append(args, *f.To)
	}
	if f.PaymentMethod != nil && *f.PaymentMethod != "" {
		where += " AND COALESCE(o.payment_method, '') = ?"
		args = append(args, *f.PaymentMethod)
	}
	if f.AmountMin != nil {
		where += " AND COALESCE(o.total_amount, 0) >= ?"
		args = append(args, *f.AmountMin)
	}
	if f.AmountMax != nil {
		where += " AND COALESCE(o.total_amount, 0) <= ?"
		args = append(args, *f.AmountMax)
	}
	if f.Search != nil && *f.Search != "" {
		where += " AND (o.tracking_code ILIKE ? OR u.mobile ILIKE ? OR u.name ILIKE ?)"
		like := "%" + *f.Search + "%"
		args = append(args, like, like, like)
	}

	// شمارش کل
	var total int64
	countSQL := "SELECT COUNT(*) FROM fraud_cases fc " + where
	if err := r.DB.WithContext(ctx).Raw(countSQL, args...).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	page := f.Page
	size := f.PageSize
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}
	offset := (page - 1) * size

	// داده صفحه‌شده با join سفارش و کاربر
	dataSQL := `
        SELECT
            fc.id,
            fc.order_id,
            o.tracking_code     AS order_number,
            u.name              AS customer_name,
            u.mobile            AS customer_phone,
            o.total_amount      AS total_amount,
            o.payment_method    AS payment_method,
            o.status            AS status,
            COALESCE(o.created_at, fc.created_at) AS created_at,
            fc.risk_level,
            fc.risk_score
        FROM fraud_cases fc
        LEFT JOIN orders o ON o.id = fc.order_id
        LEFT JOIN users u ON u.id = o.user_id
        ` + where + `
        ORDER BY fc.created_at DESC
        LIMIT ? OFFSET ?`

	args2 := append([]any{}, args...)
	args2 = append(args2, size, offset)

	var rows []AdminFraudCaseRow
	if err := r.DB.WithContext(ctx).Raw(dataSQL, args2...).Scan(&rows).Error; err != nil {
		return nil, 0, err
	}
	// نرمال‌سازی nullها به فیلدهای اشاره‌گر
	return rows, total, nil
}

// GetDetail برگرداندن کیس + رویدادها + اکشن‌ها
func (r *FraudCaseRepository) GetDetail(ctx context.Context, id uint) (*FraudCaseDetail, error) {
	var fc FraudCase
	if err := r.DB.WithContext(ctx).First(&fc, id).Error; err != nil {
		return nil, err
	}
	var events []FraudEvent
	if err := r.DB.WithContext(ctx).Where("case_id = ?", id).Order("created_at").Find(&events).Error; err != nil {
		return nil, err
	}
	var actions []FraudActionLog
	if err := r.DB.WithContext(ctx).Where("case_id = ?", id).Order("created_at").Find(&actions).Error; err != nil {
		return nil, err
	}
	return &FraudCaseDetail{Case: fc, Events: events, Actions: actions}, nil
}

type FraudEventRepository struct{ DB *gorm.DB }

func (r *FraudEventRepository) CreateMany(ctx context.Context, events []FraudEvent) error {
	if len(events) == 0 {
		return nil
	}
	return r.DB.WithContext(ctx).Create(&events).Error
}

func (r *FraudEventRepository) ListByCase(ctx context.Context, caseID uint) ([]FraudEvent, error) {
	var out []FraudEvent
	err := r.DB.WithContext(ctx).Where("case_id = ?", caseID).Order("created_at").Find(&out).Error
	return out, err
}

type FraudRuleRepository struct{ DB *gorm.DB }

func (r *FraudRuleRepository) ListActive(ctx context.Context) ([]FraudRule, error) {
	var rules []FraudRule
	err := r.DB.WithContext(ctx).Where("enabled = ?", true).Order("id").Find(&rules).Error
	return rules, err
}

func (r *FraudRuleRepository) GetByKey(ctx context.Context, key string) (*FraudRule, error) {
	var rule FraudRule
	err := r.DB.WithContext(ctx).Where("key = ?", key).First(&rule).Error
	if err != nil {
		return nil, err
	}
	return &rule, nil
}

func (r *FraudRuleRepository) GetByID(ctx context.Context, id uint) (*FraudRule, error) {
	var rule FraudRule
	err := r.DB.WithContext(ctx).First(&rule, id).Error
	if err != nil {
		return nil, err
	}
	return &rule, nil
}

func (r *FraudRuleRepository) ListAll(ctx context.Context) ([]FraudRule, error) {
	var rules []FraudRule
	err := r.DB.WithContext(ctx).Order("id").Find(&rules).Error
	return rules, err
}

func (r *FraudRuleRepository) Update(ctx context.Context, ru *FraudRule) error {
	return r.DB.WithContext(ctx).Save(ru).Error
}

type FraudListRepository struct{ DB *gorm.DB }

func (r *FraudListRepository) IsWhitelisted(ctx context.Context, kind, valueHash string) (bool, error) {
	// convention: whitelist kind prefix "wl:" vs blacklist "bl:"? Instead keep single table and handle upstream
	var count int64
	err := r.DB.WithContext(ctx).Model(&FraudListItem{}).Where("kind = ? AND value_hash = ?", kind, valueHash).Count(&count).Error
	return count > 0, err
}

func (r *FraudListRepository) IsBlacklisted(ctx context.Context, kind, valueHash string) (bool, error) {
	var count int64
	// use separate kinds e.g., ip_blacklist vs ip_whitelist decided by caller
	err := r.DB.WithContext(ctx).Model(&FraudListItem{}).Where("kind = ? AND value_hash = ?", kind, valueHash).Count(&count).Error
	return count > 0, err
}

func (r *FraudListRepository) Upsert(ctx context.Context, item *FraudListItem) error {
	// upsert روی (kind, value_hash)
	return r.DB.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "kind"}, {Name: "value_hash"}},
		DoUpdates: clause.AssignmentColumns([]string{"note", "expires_at"}),
	}).Create(item).Error
}

func (r *FraudListRepository) Delete(ctx context.Context, kind, valueHash string) error {
	return r.DB.WithContext(ctx).Where("kind = ? AND value_hash = ?", kind, valueHash).Delete(&FraudListItem{}).Error
}

type FraudActionLogRepository struct{ DB *gorm.DB }

func (r *FraudActionLogRepository) Create(ctx context.Context, l *FraudActionLog) error {
	return r.DB.WithContext(ctx).Create(l).Error
}

func (r *FraudActionLogRepository) ListByCase(ctx context.Context, caseID uint) ([]FraudActionLog, error) {
	var out []FraudActionLog
	err := r.DB.WithContext(ctx).Where("case_id = ?", caseID).Order("created_at").Find(&out).Error
	return out, err
}
