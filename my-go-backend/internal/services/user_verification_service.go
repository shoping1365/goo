package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"

	"gorm.io/datatypes"
)

// UserVerificationService - سرویس احراز هویت کاربران
type UserVerificationService struct {
	verificationRepo *repository.UserVerificationRepository
	userRepo         *repository.UserRepository
}

// NewUserVerificationService - سازنده سرویس
func NewUserVerificationService(
	verificationRepo *repository.UserVerificationRepository,
	userRepo *repository.UserRepository,
) *UserVerificationService {
	return &UserVerificationService{
		verificationRepo: verificationRepo,
		userRepo:         userRepo,
	}
}

// CreateVerificationRequest - ایجاد درخواست احراز هویت جدید
func (s *UserVerificationService) CreateVerificationRequest(req models.CreateUserVerificationRequest) (*models.UserVerification, error) {
	// بررسی وجود کاربر
	user, err := s.userRepo.FindByID(req.UserID)
	if err != nil {
		return nil, errors.New("کاربر یافت نشد")
	}
	if user == nil {
		return nil, errors.New("کاربر یافت نشد")
	}

	// تبدیل documents به JSON
	documentsJSON, err := json.Marshal(req.Documents)
	if err != nil {
		documentsJSON = []byte("[]")
	}

	// تبدیل metadata به JSON
	metadataJSON, err := json.Marshal(req.Metadata)
	if err != nil {
		metadataJSON = []byte("{}")
	}

	// ایجاد رکورد جدید
	verification := &models.UserVerification{
		UserID:             req.UserID,
		FieldName:          req.FieldName,
		FieldValue:         req.FieldValue,
		Status:             models.VerificationStatusPending,
		Documents:          datatypes.JSON(documentsJSON),
		Metadata:           datatypes.JSON(metadataJSON),
		VerificationMethod: models.MethodManual,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	if err := s.verificationRepo.Create(verification); err != nil {
		return nil, fmt.Errorf("خطا در ایجاد درخواست احراز هویت: %v", err)
	}

	return verification, nil
}

// GetVerificationByID - دریافت اطلاعات درخواست احراز هویت
func (s *UserVerificationService) GetVerificationByID(id uint) (*models.UserVerification, error) {
	return s.verificationRepo.FindByID(id)
}

// GetUserVerifications - دریافت لیست درخواست‌های یک کاربر
func (s *UserVerificationService) GetUserVerifications(userID uint) ([]models.UserVerification, error) {
	return s.verificationRepo.FindByUserID(userID)
}

// GetPendingVerifications - دریافت لیست درخواست‌های در انتظار تایید
func (s *UserVerificationService) GetPendingVerifications(page, pageSize int) ([]models.UserVerification, int64, error) {
	offset := (page - 1) * pageSize
	return s.verificationRepo.FindByStatus(models.VerificationStatusPending, pageSize, offset)
}

// GetAllVerifications - دریافت تمام درخواست‌ها با صفحه‌بندی
func (s *UserVerificationService) GetAllVerifications(page, pageSize int) ([]models.UserVerification, int64, error) {
	offset := (page - 1) * pageSize
	return s.verificationRepo.FindAll(pageSize, offset)
}

// VerifyRequest - تایید درخواست احراز هویت توسط ادمین (فاز 1)
func (s *UserVerificationService) VerifyRequest(id uint, verifiedBy uint) error {
	// دریافت درخواست
	verification, err := s.verificationRepo.FindByID(id)
	if err != nil {
		return errors.New("درخواست یافت نشد")
	}

	// بررسی وضعیت فعلی
	if verification.Status != models.VerificationStatusPending {
		return errors.New("این درخواست قبلاً پردازش شده است")
	}

	// به‌روزرسانی وضعیت
	return s.verificationRepo.UpdateStatus(id, models.VerificationStatusVerified, &verifiedBy, "")
}

// RejectRequest - رد درخواست احراز هویت توسط ادمین (فاز 1)
func (s *UserVerificationService) RejectRequest(id uint, verifiedBy uint, reason string) error {
	// دریافت درخواست
	verification, err := s.verificationRepo.FindByID(id)
	if err != nil {
		return errors.New("درخواست یافت نشد")
	}

	// بررسی وضعیت فعلی
	if verification.Status != models.VerificationStatusPending {
		return errors.New("این درخواست قبلاً پردازش شده است")
	}

	// بررسی وجود دلیل رد
	if reason == "" {
		return errors.New("دلیل رد درخواست الزامی است")
	}

	// به‌روزرسانی وضعیت
	return s.verificationRepo.UpdateStatus(id, models.VerificationStatusRejected, &verifiedBy, reason)
}

// GetVerifiedFields - دریافت فیلدهای تایید شده یک کاربر
func (s *UserVerificationService) GetVerifiedFields(userID uint) (map[string]bool, error) {
	return s.verificationRepo.GetVerifiedFieldsByUserID(userID)
}

// GetPendingCount - دریافت تعداد درخواست‌های در انتظار
func (s *UserVerificationService) GetPendingCount() (int64, error) {
	return s.verificationRepo.GetPendingCount()
}

// SearchVerifications - جستجوی درخواست‌ها
func (s *UserVerificationService) SearchVerifications(params map[string]interface{}, page, pageSize int) ([]models.UserVerification, int64, error) {
	offset := (page - 1) * pageSize
	return s.verificationRepo.SearchVerifications(params, pageSize, offset)
}

// VerifyWithAPI - احراز هویت از طریق API خارجی (فاز 2)
// این متد برای اتصال به سرویس‌های خارجی مثل Raypa، Fanap و... استفاده می‌شود
func (s *UserVerificationService) VerifyWithAPI(id uint, apiProvider string) error {
	// دریافت درخواست
	verification, err := s.verificationRepo.FindByID(id)
	if err != nil {
		return errors.New("درخواست یافت نشد")
	}

	// بررسی وضعیت
	if verification.Status != models.VerificationStatusPending {
		return errors.New("این درخواست قبلاً پردازش شده است")
	}

	// TODO: پیاده‌سازی ارتباط با API های مختلف
	// switch apiProvider {
	// case "raypa":
	//     return s.verifyWithRaypa(verification)
	// case "fanap":
	//     return s.verifyWithFanap(verification)
	// case "rahkareshna":
	//     return s.verifyWithRahkareshna(verification)
	// default:
	//     return errors.New("سرویس‌دهنده نامعتبر است")
	// }

	return errors.New("این قابلیت هنوز پیاده‌سازی نشده است (فاز 2)")
}

// VerifyWithEKYC - احراز هویت چهره‌ای (فاز 3)
// این متد برای احراز هویت چهره‌ای با سلفی و کارت ملی استفاده می‌شود
func (s *UserVerificationService) VerifyWithEKYC(id uint, selfieURL, idCardURL string) error {
	// دریافت درخواست
	verification, err := s.verificationRepo.FindByID(id)
	if err != nil {
		return errors.New("درخواست یافت نشد")
	}

	// بررسی وضعیت
	if verification.Status != models.VerificationStatusPending {
		return errors.New("این درخواست قبلاً پردازش شده است")
	}

	// TODO: پیاده‌سازی احراز هویت چهره‌ای
	// 1. ارسال تصاویر به سرویس تشخیص چهره
	// 2. استخراج اطلاعات از کارت ملی (OCR)
	// 3. مقایسه چهره سلفی با چهره کارت ملی
	// 4. تایید نهایی با سرویس ثبت احوال

	return errors.New("این قابلیت هنوز پیاده‌سازی نشده است (فاز 3)")
}

// DeleteVerification - حذف درخواست احراز هویت
func (s *UserVerificationService) DeleteVerification(id uint) error {
	return s.verificationRepo.Delete(id)
}

// ConvertToResponse - تبدیل مدل به Response
func (s *UserVerificationService) ConvertToResponse(verification *models.UserVerification) *models.UserVerificationResponse {
	var documents []string
	var metadata map[string]interface{}

	// Parse documents
	if len(verification.Documents) > 0 {
		_ = json.Unmarshal(verification.Documents, &documents)
	}

	// Parse metadata
	if len(verification.Metadata) > 0 {
		_ = json.Unmarshal(verification.Metadata, &metadata)
	}

	response := &models.UserVerificationResponse{
		ID:                 verification.ID,
		UserID:             verification.UserID,
		FieldName:          verification.FieldName,
		FieldValue:         verification.FieldValue,
		Status:             verification.Status,
		VerifiedAt:         verification.VerifiedAt,
		VerifiedBy:         verification.VerifiedBy,
		RejectionReason:    verification.RejectionReason,
		Documents:          documents,
		Metadata:           metadata,
		VerificationMethod: verification.VerificationMethod,
		CreatedAt:          verification.CreatedAt,
		UpdatedAt:          verification.UpdatedAt,
	}

	// اضافه کردن اطلاعات کاربر
	if verification.User.ID > 0 {
		response.UserName = verification.User.Name
		response.UserMobile = verification.User.Mobile
	}

	// اضافه کردن اطلاعات تایید کننده
	if verification.VerifiedByUser != nil && verification.VerifiedByUser.ID > 0 {
		response.VerifiedByName = verification.VerifiedByUser.Name
	}

	return response
}
