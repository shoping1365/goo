package services

import (
	"fmt"
	"math/rand"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"strconv"
	"time"
)

type GiftCardService struct {
	repo *repository.GiftCardRepository
}

// CreateGiftCard ایجاد گیفت کارت جدید
func (s *GiftCardService) CreateGiftCard(giftCard *models.GiftCard) error {
	// تولید کد منحصر به فرد
	giftCard.Code = s.generateUniqueCode()

	// تنظیم مقادیر پیش‌فرض
	if giftCard.Status == "" {
		giftCard.Status = "active"
	}
	if giftCard.RemainingAmount == 0 {
		giftCard.RemainingAmount = giftCard.Amount
	}
	if giftCard.ExpiryDate.IsZero() {
		giftCard.ExpiryDate = time.Now().AddDate(1, 0, 0) // یک سال بعد
	}

	// ایجاد گیفت کارت
	err := s.repo.Create(giftCard)
	if err != nil {
		return err
	}

	// ثبت فعالیت
	activity := &models.GiftCardActivity{
		GiftCardID: giftCard.ID,
		Action:     "created",
		Amount:     giftCard.Amount,
		CreatedBy:  giftCard.CreatedBy,
		CreatedAt:  time.Now(),
	}
	if err := s.repo.CreateActivity(activity); err != nil {
		// لاگ کردن خطا در صورت عدم موفقیت در ثبت فعالیت
		fmt.Printf("Failed to create gift card activity: %v\n", err)
	}

	return nil
}

// GetGiftCard دریافت گیفت کارت بر اساس ID
func (s *GiftCardService) GetGiftCard(id uint) (*models.GiftCard, error) {
	return s.repo.GetByID(id)
}

// GetGiftCardByCode دریافت گیفت کارت بر اساس کد
func (s *GiftCardService) GetGiftCardByCode(code string) (*models.GiftCard, error) {
	return s.repo.GetByCode(code)
}

// GetAllGiftCards دریافت تمام گیفت کارت‌ها با فیلتر و صفحه‌بندی
func (s *GiftCardService) GetAllGiftCards(page, pageSize int, status, cardType, category string) ([]models.GiftCard, int64, error) {
	offset := (page - 1) * pageSize
	return s.repo.GetAll(pageSize, offset, status, cardType, category)
}

// UpdateGiftCard به‌روزرسانی گیفت کارت
func (s *GiftCardService) UpdateGiftCard(giftCard *models.GiftCard) error {
	// بررسی وجود گیفت کارت
	_, err := s.repo.GetByID(giftCard.ID)
	if err != nil {
		return err
	}

	// ثبت فعالیت تغییر
	activity := &models.GiftCardActivity{
		GiftCardID: giftCard.ID,
		Action:     "updated",
		Amount:     giftCard.Amount,
		CreatedBy:  giftCard.CreatedBy,
		CreatedAt:  time.Now(),
	}
	if err := s.repo.CreateActivity(activity); err != nil {
		// لاگ کردن خطا در صورت عدم موفقیت در ثبت فعالیت
		fmt.Printf("Failed to create gift card activity: %v\n", err)
	}

	return s.repo.Update(giftCard)
}

// DeleteGiftCard حذف گیفت کارت
func (s *GiftCardService) DeleteGiftCard(id uint, userID uint) error {
	// بررسی وجود گیفت کارت
	giftCard, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// ثبت فعالیت حذف
	activity := &models.GiftCardActivity{
		GiftCardID: id,
		Action:     "deleted",
		Amount:     giftCard.Amount,
		CreatedBy:  userID,
		CreatedAt:  time.Now(),
	}
	if err := s.repo.CreateActivity(activity); err != nil {
		// لاگ کردن خطا در صورت عدم موفقیت در ثبت فعالیت
		fmt.Printf("Failed to create gift card activity: %v\n", err)
	}

	return s.repo.Delete(id)
}

// GetDashboardStats دریافت آمار داشبورد
func (s *GiftCardService) GetDashboardStats() (*models.GiftCardStats, error) {
	return s.repo.GetStats()
}

// GetRecentActivities دریافت فعالیت‌های اخیر
func (s *GiftCardService) GetRecentActivities(limit int) ([]models.GiftCardActivity, error) {
	return s.repo.GetRecentActivities(limit)
}

// GetChartData دریافت داده‌های نمودار
func (s *GiftCardService) GetChartData() (map[string]interface{}, error) {
	return s.repo.GetChartData()
}

// UseGiftCard استفاده از گیفت کارت
func (s *GiftCardService) UseGiftCard(code string, amount int64, userID uint) error {
	giftCard, err := s.repo.GetByCode(code)
	if err != nil {
		return fmt.Errorf("گیفت کارت یافت نشد")
	}

	if giftCard.Status != "active" {
		return fmt.Errorf("گیفت کارت غیرفعال است")
	}

	if giftCard.RemainingAmount < amount {
		return fmt.Errorf("مبلغ باقی‌مانده کافی نیست")
	}

	// بررسی محدودیت‌های استفاده
	if giftCard.UsageLimit == "once" && giftCard.UsageCount > 0 {
		return fmt.Errorf("این گیفت کارت فقط یک بار قابل استفاده است")
	}

	if giftCard.MaxUsageCount > 0 && giftCard.UsageCount >= giftCard.MaxUsageCount {
		return fmt.Errorf("تعداد استفاده مجاز تمام شده است")
	}

	// به‌روزرسانی مبلغ باقی‌مانده
	giftCard.RemainingAmount -= amount
	giftCard.UsageCount++

	// بررسی انقضا
	if giftCard.ExpiryDate.Before(time.Now()) {
		giftCard.Status = "expired"
	}

	// بررسی استفاده کامل
	if giftCard.RemainingAmount == 0 {
		giftCard.Status = "used"
	}

	// به‌روزرسانی آمار استفاده
	if giftCard.UsageCount > 0 {
		usedAmount := giftCard.Amount - giftCard.RemainingAmount
		giftCard.AvgUsageAmount = usedAmount / int64(giftCard.UsageCount)
	}
	giftCard.MaxUsageAmount = giftCard.Amount - giftCard.RemainingAmount
	giftCard.LastUsageAmount = amount

	// ذخیره تغییرات
	err = s.repo.Update(giftCard)
	if err != nil {
		return err
	}

	// ثبت تراکنش
	transaction := &models.GiftCardTransaction{
		GiftCardID:  giftCard.ID,
		Amount:      amount,
		Type:        "usage",
		Description: fmt.Sprintf("استفاده از گیفت کارت %s", giftCard.Code),
		CreatedBy:   userID,
		CreatedAt:   time.Now(),
	}
	if err := s.repo.CreateTransaction(transaction); err != nil {
		// لاگ کردن خطا در صورت عدم موفقیت در ثبت تراکنش
		fmt.Printf("Failed to create gift card transaction: %v\n", err)
	}

	// ثبت فعالیت
	activity := &models.GiftCardActivity{
		GiftCardID: giftCard.ID,
		Action:     "used",
		Amount:     amount,
		CreatedBy:  userID,
		CreatedAt:  time.Now(),
	}
	if err := s.repo.CreateActivity(activity); err != nil {
		// لاگ کردن خطا در صورت عدم موفقیت در ثبت فعالیت
		fmt.Printf("Failed to create gift card activity: %v\n", err)
	}

	return nil
}

// ActivateGiftCard فعال کردن گیفت کارت
func (s *GiftCardService) ActivateGiftCard(id uint, userID uint) error {
	giftCard, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if giftCard.Status == "active" {
		return fmt.Errorf("گیفت کارت قبلاً فعال است")
	}

	giftCard.Status = "active"

	// ثبت فعالیت
	activity := &models.GiftCardActivity{
		GiftCardID: id,
		Action:     "activated",
		Amount:     giftCard.Amount,
		CreatedBy:  userID,
		CreatedAt:  time.Now(),
	}
	if err := s.repo.CreateActivity(activity); err != nil {
		// لاگ کردن خطا در صورت عدم موفقیت در ثبت فعالیت
		fmt.Printf("Failed to create gift card activity: %v\n", err)
	}

	return s.repo.Update(giftCard)
}

// DeactivateGiftCard غیرفعال کردن گیفت کارت
func (s *GiftCardService) DeactivateGiftCard(id uint, userID uint) error {
	giftCard, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if giftCard.Status != "active" {
		return fmt.Errorf("گیفت کارت فعال نیست")
	}

	giftCard.Status = "inactive"

	// ثبت فعالیت
	activity := &models.GiftCardActivity{
		GiftCardID: id,
		Action:     "deactivated",
		Amount:     giftCard.Amount,
		CreatedBy:  userID,
		CreatedAt:  time.Now(),
	}
	if err := s.repo.CreateActivity(activity); err != nil {
		// لاگ کردن خطا در صورت عدم موفقیت در ثبت فعالیت
		fmt.Printf("Failed to create gift card activity: %v\n", err)
	}

	return s.repo.Update(giftCard)
}

// SearchGiftCards جستجو در گیفت کارت‌ها
func (s *GiftCardService) SearchGiftCards(query string, page, pageSize int) ([]models.GiftCard, int64, error) {
	offset := (page - 1) * pageSize
	return s.repo.Search(query, pageSize, offset)
}

// GetTransactions دریافت تراکنش‌های گیفت کارت
func (s *GiftCardService) GetTransactions(giftCardID uint, page, pageSize int) ([]models.GiftCardTransaction, int64, error) {
	offset := (page - 1) * pageSize
	return s.repo.GetTransactions(giftCardID, pageSize, offset)
}

// generateUniqueCode تولید کد منحصر به فرد
func (s *GiftCardService) generateUniqueCode() string {
	// rand.Seed is deprecated in Go 1.20+, the global random generator is automatically seeded

	for {
		// تولید کد 8 رقمی
		code := "GC-" + strconv.Itoa(rand.Intn(90000000)+10000000)

		// بررسی یکتا بودن
		_, err := s.repo.GetByCode(code)
		if err != nil {
			// اگر خطا داد یعنی کد وجود ندارد
			return code
		}
	}
}

// ValidateGiftCard اعتبارسنجی گیفت کارت
func (s *GiftCardService) ValidateGiftCard(giftCard *models.GiftCard) error {
	if giftCard.Amount <= 0 {
		return fmt.Errorf("مبلغ باید بزرگتر از صفر باشد")
	}

	if giftCard.RecipientEmail != "" {
		// اعتبارسنجی ایمیل (ساده)
		if len(giftCard.RecipientEmail) < 5 || !contains(giftCard.RecipientEmail, "@") {
			return fmt.Errorf("ایمیل معتبر نیست")
		}
	}

	if giftCard.RecipientPhone != "" {
		// اعتبارسنجی شماره تلفن (ساده)
		if len(giftCard.RecipientPhone) < 10 {
			return fmt.Errorf("شماره تلفن معتبر نیست")
		}
	}

	if giftCard.MaxUsageCount < 0 {
		return fmt.Errorf("تعداد استفاده مجاز نمی‌تواند منفی باشد")
	}

	if giftCard.MinOrderAmount < 0 || giftCard.MaxOrderAmount < 0 {
		return fmt.Errorf("محدودیت مبلغ سفارش نمی‌تواند منفی باشد")
	}

	if giftCard.MinOrderAmount > giftCard.MaxOrderAmount {
		return fmt.Errorf("حداقل مبلغ سفارش نمی‌تواند بیشتر از حداکثر باشد")
	}

	return nil
}

// contains بررسی وجود رشته در رشته دیگر
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > len(substr) && (s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			contains(s[1:], substr))))
}
