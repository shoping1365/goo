package models

import (
	"time"

	"gorm.io/datatypes"
)

// UserVerification - مدل احراز هویت کاربران (KYC)
// این مدل برای ذخیره درخواست‌های احراز هویت و تایید اطلاعات شخصی کاربران استفاده می‌شود
type UserVerification struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"not null;index"`

	// نوع فیلد و مقدار آن
	FieldName  string `json:"field_name" gorm:"type:varchar(50);not null;index"` // 'national_code', 'sheba_number', 'birth_date', 'gender', etc.
	FieldValue string `json:"field_value" gorm:"type:text;not null"`             // مقدار فیلد

	// وضعیت تایید
	Status string `json:"status" gorm:"type:varchar(20);default:'pending';index"` // 'pending', 'verified', 'rejected'

	// اطلاعات تایید/رد
	VerifiedAt      *time.Time `json:"verified_at"`
	VerifiedBy      *uint      `json:"verified_by"` // ID ادمین
	RejectionReason string     `json:"rejection_reason" gorm:"type:text"`

	// مدارک پیوست (لیست URLها)
	Documents datatypes.JSON `json:"documents" gorm:"type:jsonb;default:'[]'"`

	// متاداده اضافی
	Metadata datatypes.JSON `json:"metadata" gorm:"type:jsonb;default:'{}'"`

	// روش احراز هویت
	VerificationMethod string `json:"verification_method" gorm:"type:varchar(20);default:'manual'"` // 'manual', 'api', 'ekyc'

	// اطلاعات API (در صورت استفاده)
	APIProvider   string         `json:"api_provider" gorm:"type:varchar(50)"`
	APIResponse   datatypes.JSON `json:"api_response" gorm:"type:jsonb"`
	APIVerifiedAt *time.Time     `json:"api_verified_at"`

	// زمان‌ها
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;index"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// روابط
	User           User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	VerifiedByUser *User `json:"verified_by_user,omitempty" gorm:"foreignKey:VerifiedBy"`
}

// TableName - نام جدول
func (UserVerification) TableName() string {
	return "user_verifications"
}

// VerificationStatus - ثابت‌های وضعیت
const (
	VerificationStatusPending  = "pending"
	VerificationStatusVerified = "verified"
	VerificationStatusRejected = "rejected"
)

// VerificationFieldName - ثابت‌های نام فیلدها
const (
	FieldNationalCode   = "national_code"   // کد ملی
	FieldShebaNumber    = "sheba_number"    // شماره شبا
	FieldBirthDate      = "birth_date"      // تاریخ تولد
	FieldGender         = "gender"          // جنسیت
	FieldJob            = "job"             // شغل
	FieldEconomicCode   = "economic_code"   // کد اقتصادی
	FieldDisabilityType = "disability_type" // نوع معلولیت
	FieldLegalInfo      = "legal_info"      // اطلاعات حقوقی
	FieldFullInfo       = "full_profile"    // پروفایل کامل
)

// VerificationMethod - ثابت‌های روش احراز هویت
const (
	MethodManual = "manual" // تایید دستی توسط ادمین
	MethodAPI    = "api"    // استفاده از API خارجی
	MethodEKYC   = "ekyc"   // احراز هویت چهره‌ای
)

// CreateUserVerificationRequest - درخواست ایجاد احراز هویت جدید
type CreateUserVerificationRequest struct {
	UserID     uint                   `json:"user_id" binding:"required"`
	FieldName  string                 `json:"field_name" binding:"required"`
	FieldValue string                 `json:"field_value" binding:"required"`
	Documents  []string               `json:"documents"`
	Metadata   map[string]interface{} `json:"metadata"`
}

// UpdateVerificationStatusRequest - درخواست تایید/رد احراز هویت
type UpdateVerificationStatusRequest struct {
	Status          string `json:"status" binding:"required,oneof=verified rejected"`
	RejectionReason string `json:"rejection_reason"`
	VerifiedBy      uint   `json:"verified_by" binding:"required"`
}

// UserVerificationResponse - پاسخ احراز هویت برای نمایش
type UserVerificationResponse struct {
	ID                 uint                   `json:"id"`
	UserID             uint                   `json:"user_id"`
	UserName           string                 `json:"user_name,omitempty"`
	UserMobile         string                 `json:"user_mobile,omitempty"`
	FieldName          string                 `json:"field_name"`
	FieldValue         string                 `json:"field_value"`
	Status             string                 `json:"status"`
	VerifiedAt         *time.Time             `json:"verified_at,omitempty"`
	VerifiedBy         *uint                  `json:"verified_by,omitempty"`
	VerifiedByName     string                 `json:"verified_by_name,omitempty"`
	RejectionReason    string                 `json:"rejection_reason,omitempty"`
	Documents          []string               `json:"documents"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
	VerificationMethod string                 `json:"verification_method"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
}
