package services

import (
	"errors"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"regexp"
	"strings"
	"time"
)

// UserBankingInfoService سرویس مدیریت اطلاعات بانکی کاربران
type UserBankingInfoService struct {
	bankingInfoRepo repository.UserBankingInfoRepository
}

// NewUserBankingInfoService ایجاد نمونه جدید از UserBankingInfoService
func NewUserBankingInfoService(bankingInfoRepo repository.UserBankingInfoRepository) *UserBankingInfoService {
	return &UserBankingInfoService{
		bankingInfoRepo: bankingInfoRepo,
	}
}

// Create ایجاد اطلاعات بانکی جدید
func (s *UserBankingInfoService) Create(userID uint, req *models.BankingInfoRequest) (*models.BankingInfoResponse, error) {
	// اعتبارسنجی داده‌ها
	if err := s.validateBankingInfo(req); err != nil {
		return nil, err
	}

	// اگر این حساب پیشفرض است، ابتدا تمام حساب‌های دیگر را غیرپیشفرض کنیم
	if req.IsDefault {
		if err := s.setAllNonDefault(userID); err != nil {
			return nil, err
		}
	}

	// آماده‌سازی داده‌ها
	bankingInfo := &models.UserBankingInfo{
		UserID:            userID,
		BankName:          strings.TrimSpace(req.BankName),
		CardNumber:        s.cleanCardNumber(req.CardNumber),
		AccountNumber:     strings.TrimSpace(req.AccountNumber),
		ShebaNumber:       s.cleanShebaNumber(req.ShebaNumber),
		AccountHolderName: strings.TrimSpace(req.AccountHolderName),
		AccountType:       strings.TrimSpace(req.AccountType),
		IsDefault:         req.IsDefault,
	}

	// ایجاد اطلاعات جدید
	if err := s.bankingInfoRepo.Create(bankingInfo); err != nil {
		return nil, err
	}

	return s.toResponse(bankingInfo), nil
}

// Update به‌روزرسانی اطلاعات بانکی موجود
func (s *UserBankingInfoService) Update(userID uint, id uint, req *models.BankingInfoRequest) (*models.BankingInfoResponse, error) {
	// اعتبارسنجی داده‌ها
	if err := s.validateBankingInfo(req); err != nil {
		return nil, err
	}

	// دریافت اطلاعات موجود
	bankingInfo, err := s.bankingInfoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if bankingInfo == nil {
		return nil, errors.New("اطلاعات بانکی یافت نشد")
	}

	// بررسی مالکیت
	if bankingInfo.UserID != userID {
		return nil, errors.New("شما مجاز به ویرایش این اطلاعات نیستید")
	}

	// اگر این حساب پیشفرض است، ابتدا تمام حساب‌های دیگر را غیرپیشفرض کنیم
	if req.IsDefault && !bankingInfo.IsDefault {
		if err := s.setAllNonDefault(userID); err != nil {
			return nil, err
		}
	}

	// به‌روزرسانی اطلاعات
	bankingInfo.BankName = strings.TrimSpace(req.BankName)
	bankingInfo.CardNumber = s.cleanCardNumber(req.CardNumber)
	bankingInfo.AccountNumber = strings.TrimSpace(req.AccountNumber)
	bankingInfo.ShebaNumber = s.cleanShebaNumber(req.ShebaNumber)
	bankingInfo.AccountHolderName = strings.TrimSpace(req.AccountHolderName)
	bankingInfo.AccountType = strings.TrimSpace(req.AccountType)
	bankingInfo.IsDefault = req.IsDefault

	if err := s.bankingInfoRepo.Update(bankingInfo); err != nil {
		return nil, err
	}

	return s.toResponse(bankingInfo), nil
}

// GetAllByUserID دریافت تمام اطلاعات بانکی کاربر
func (s *UserBankingInfoService) GetAllByUserID(userID uint) ([]models.BankingInfoResponse, error) {
	bankingInfos, err := s.bankingInfoRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []models.BankingInfoResponse
	for _, info := range bankingInfos {
		responses = append(responses, *s.toResponse(&info))
	}

	return responses, nil
}

// GetDefaultByUserID دریافت حساب پیشفرض کاربر
func (s *UserBankingInfoService) GetDefaultByUserID(userID uint) (*models.BankingInfoResponse, error) {
	bankingInfo, err := s.bankingInfoRepo.GetDefaultByUserID(userID)
	if err != nil {
		return nil, err
	}
	if bankingInfo == nil {
		return nil, nil
	}
	return s.toResponse(bankingInfo), nil
}

// Delete حذف اطلاعات بانکی
func (s *UserBankingInfoService) Delete(userID uint, id uint) error {
	// دریافت اطلاعات موجود
	bankingInfo, err := s.bankingInfoRepo.GetByID(id)
	if err != nil {
		return err
	}
	if bankingInfo == nil {
		return errors.New("اطلاعات بانکی یافت نشد")
	}

	// بررسی مالکیت
	if bankingInfo.UserID != userID {
		return errors.New("شما مجاز به حذف این اطلاعات نیستید")
	}

	return s.bankingInfoRepo.Delete(id)
}

// SetDefault تنظیم حساب پیشفرض
func (s *UserBankingInfoService) SetDefault(userID uint, id uint) error {
	// بررسی مالکیت
	bankingInfo, err := s.bankingInfoRepo.GetByID(id)
	if err != nil {
		return err
	}
	if bankingInfo == nil {
		return errors.New("اطلاعات بانکی یافت نشد")
	}
	if bankingInfo.UserID != userID {
		return errors.New("شما مجاز به تغییر این اطلاعات نیستید")
	}

	return s.bankingInfoRepo.SetDefault(userID, id)
}

// Verify تایید اطلاعات بانکی توسط ادمین
func (s *UserBankingInfoService) Verify(bankingInfoID uint, verifiedBy uint, note string) error {
	bankingInfo, err := s.bankingInfoRepo.GetByID(bankingInfoID)
	if err != nil {
		return err
	}
	if bankingInfo == nil {
		return errors.New("اطلاعات بانکی یافت نشد")
	}

	now := time.Now()
	bankingInfo.IsVerified = true
	bankingInfo.VerifiedAt = &now
	bankingInfo.VerifiedBy = &verifiedBy
	bankingInfo.VerificationNote = note

	return s.bankingInfoRepo.Update(bankingInfo)
}

// Unverify لغو تایید اطلاعات بانکی
func (s *UserBankingInfoService) Unverify(bankingInfoID uint, verifiedBy uint, note string) error {
	bankingInfo, err := s.bankingInfoRepo.GetByID(bankingInfoID)
	if err != nil {
		return err
	}
	if bankingInfo == nil {
		return errors.New("اطلاعات بانکی یافت نشد")
	}

	bankingInfo.IsVerified = false
	bankingInfo.VerifiedAt = nil
	bankingInfo.VerifiedBy = &verifiedBy
	bankingInfo.VerificationNote = note

	return s.bankingInfoRepo.Update(bankingInfo)
}

// AutoCompleteFromCard تکمیل خودکار از شماره کارت
func (s *UserBankingInfoService) AutoCompleteFromCard(cardNumber string) (*models.BankingInfoRequest, error) {
	cardNumber = s.cleanCardNumber(cardNumber)
	if len(cardNumber) < 16 {
		return nil, errors.New("شماره کارت باید حداقل 16 رقم باشد")
	}

	// تشخیص بانک از BIN
	bankName := s.detectBankFromBIN(cardNumber[:6])
	if bankName == "" {
		return nil, errors.New("نمی‌توان بانک را از شماره کارت تشخیص داد")
	}

	return &models.BankingInfoRequest{
		BankName:    bankName,
		CardNumber:  cardNumber,
		ShebaNumber: "", // خالی بگذاریم تا کاربر خودش وارد کند
	}, nil
}

// AutoCompleteFromSheba تکمیل خودکار از شماره شبا
func (s *UserBankingInfoService) AutoCompleteFromSheba(shebaNumber string) (*models.BankingInfoRequest, error) {
	shebaNumber = s.cleanShebaNumber(shebaNumber)
	if len(shebaNumber) != 26 {
		return nil, errors.New("شماره شبا باید 26 کاراکتر باشد")
	}

	// تشخیص بانک از شبا
	bankName := s.detectBankFromSheba(shebaNumber)
	if bankName == "" {
		return nil, errors.New("نمی‌توان بانک را از شماره شبا تشخیص داد")
	}

	return &models.BankingInfoRequest{
		BankName:    bankName,
		CardNumber:  "", // خالی بگذاریم تا کاربر خودش وارد کند
		ShebaNumber: shebaNumber,
	}, nil
}

// generateSimpleAccountFromCard تولید ساده شماره حساب از کارت
func (s *UserBankingInfoService) generateSimpleAccountFromCard(cardNumber, bankName string) string {
	bankCode := s.getBankCodeFromName(bankName)
	if bankCode == "" {
		return ""
	}

	// حذف BIN (6 رقم اول) و گرفتن 8 رقم بعدی
	cardDigits := cardNumber[6:]
	if len(cardDigits) < 8 {
		return ""
	}

	// تولید شماره حساب ساده
	return bankCode + cardDigits[:8]
}

// generateSimpleCardFromAccount تولید ساده شماره کارت از حساب
func (s *UserBankingInfoService) generateSimpleCardFromAccount(accountNumber, bankName string) string {
	bankCode := s.getBankCodeFromName(bankName)
	if bankCode == "" {
		return ""
	}

	// دریافت BIN کد برای بانک
	binCode := s.getBINFromBankName(bankName)
	if binCode == "" {
		return ""
	}

	// حذف کد بانک از شماره حساب
	accountDigits := accountNumber[2:]
	if len(accountDigits) < 10 {
		return ""
	}

	// تولید شماره کارت ساده
	return binCode + accountDigits[:10]
}

// generateSimpleAccountFromSheba تولید ساده شماره حساب از شبا
func (s *UserBankingInfoService) generateSimpleAccountFromSheba(shebaNumber, bankName string) string {
	bankCode := s.getBankCodeFromName(bankName)
	if bankCode == "" {
		return ""
	}

	// استخراج ارقام حساب از شبا
	accountDigits := shebaNumber[7:26]
	if len(accountDigits) < 8 {
		return ""
	}

	// تولید شماره حساب ساده
	return bankCode + accountDigits[:8]
}

// setAllNonDefault تنظیم تمام حساب‌های کاربر به غیرپیشفرض
func (s *UserBankingInfoService) setAllNonDefault(userID uint) error {
	// این تابع در repository پیاده‌سازی می‌شود
	return s.bankingInfoRepo.SetDefault(userID, 0) // 0 به معنای غیرفعال کردن همه
}

// validateBankingInfo اعتبارسنجی اطلاعات بانکی
func (s *UserBankingInfoService) validateBankingInfo(req *models.BankingInfoRequest) error {
	if strings.TrimSpace(req.BankName) == "" {
		return errors.New("نام بانک الزامی است")
	}
	if strings.TrimSpace(req.CardNumber) == "" {
		return errors.New("شماره کارت الزامی است")
	}
	if strings.TrimSpace(req.AccountNumber) == "" {
		return errors.New("شماره حساب الزامی است")
	}
	if strings.TrimSpace(req.AccountHolderName) == "" {
		return errors.New("نام صاحب حساب الزامی است")
	}

	// اعتبارسنجی شماره کارت
	cardNumber := s.cleanCardNumber(req.CardNumber)
	if len(cardNumber) < 16 || len(cardNumber) > 19 {
		return errors.New("شماره کارت باید بین 16 تا 19 رقم باشد")
	}

	// اعتبارسنجی شماره شبا (اختیاری)
	shebaNumber := s.cleanShebaNumber(req.ShebaNumber)
	if shebaNumber != "" {
		if len(shebaNumber) != 26 {
			return errors.New("شماره شبا باید 26 کاراکتر باشد")
		}
		if !strings.HasPrefix(shebaNumber, "IR") {
			return errors.New("شماره شبا باید با IR شروع شود")
		}
	}

	return nil
}

// cleanCardNumber پاک‌سازی شماره کارت
func (s *UserBankingInfoService) cleanCardNumber(cardNumber string) string {
	// حذف فاصله و خط تیره
	cleaned := strings.ReplaceAll(cardNumber, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	return cleaned
}

// cleanShebaNumber پاک‌سازی شماره شبا
func (s *UserBankingInfoService) cleanShebaNumber(shebaNumber string) string {
	// حذف فاصله و تبدیل به حروف بزرگ
	cleaned := strings.ReplaceAll(shebaNumber, " ", "")
	cleaned = strings.ToUpper(cleaned)
	return cleaned
}

// detectBankFromBIN تشخیص بانک از BIN کد
func (s *UserBankingInfoService) detectBankFromBIN(bin string) string {
	binDatabase := map[string]string{
		"603799": "بانک ملی",
		"589210": "بانک سپه",
		"627353": "بانک تجارت",
		"627648": "بانک صادرات",
		"610433": "بانک ملت",
		"622106": "بانک پارسیان",
		"639347": "بانک پاسارگاد",
		"621986": "بانک سامان",
		"636214": "بانک آینده",
		"502938": "بانک دی",
		"504706": "بانک شهر",
		"627488": "بانک کارآفرین",
		"627381": "بانک اقتصاد نوین",
		"627961": "بانک انصار",
		"627760": "بانک پست بانک",
		"639599": "بانک قوامین",
		"639370": "بانک مهر اقتصاد",
		"505416": "بانک گردشگری",
		"603770": "بانک کشاورزی",
		"628023": "بانک مسکن",
		"627884": "بانک توسعه تعاون",
		"589463": "بانک رفاه کارگران",
		"505801": "بانک خاورمیانه",
		"639346": "بانک سینا",
	}

	return binDatabase[bin]
}

// detectBankFromAccount تشخیص بانک از شماره حساب
func (s *UserBankingInfoService) detectBankFromAccount(accountNumber string) string {
	// الگوهای شناخته شده برای شماره حساب‌های مختلف بانک‌ها
	accountPatterns := map[string]string{
		"^01": "بانک ملی",
		"^02": "بانک سپه",
		"^03": "بانک تجارت",
		"^04": "بانک صادرات",
		"^05": "بانک ملت",
		"^06": "بانک پارسیان",
		"^07": "بانک پاسارگاد",
		"^08": "بانک سامان",
		"^09": "بانک آینده",
		"^10": "بانک دی",
		"^11": "بانک شهر",
		"^12": "بانک کارآفرین",
		"^13": "بانک اقتصاد نوین",
		"^14": "بانک انصار",
		"^15": "بانک پست بانک",
		"^16": "بانک قوامین",
		"^17": "بانک مهر اقتصاد",
		"^18": "بانک گردشگری",
		"^19": "بانک صنعت و معدن",
		"^20": "بانک کشاورزی",
		"^21": "بانک مسکن",
		"^22": "بانک توسعه تعاون",
		"^23": "بانک رفاه کارگران",
		"^24": "بانک ایران زمین",
		"^25": "بانک خاورمیانه",
		"^26": "بانک سینا",
		"^27": "بانک توسعه",
		"^28": "بانک حکمت ایرانیان",
	}

	for pattern, bankName := range accountPatterns {
		matched, _ := regexp.MatchString(pattern, accountNumber)
		if matched {
			return bankName
		}
	}

	return ""
}

// detectBankFromSheba تشخیص بانک از شماره شبا
func (s *UserBankingInfoService) detectBankFromSheba(shebaNumber string) string {
	if len(shebaNumber) < 7 {
		return ""
	}

	bankCode := shebaNumber[2:7]
	shebaDatabase := map[string]string{
		"001": "بانک ملی",
		"002": "بانک سپه",
		"003": "بانک تجارت",
		"004": "بانک صادرات",
		"005": "بانک ملت",
		"006": "بانک پارسیان",
		"007": "بانک پاسارگاد",
		"008": "بانک سامان",
		"009": "بانک آینده",
		"010": "بانک دی",
		"011": "بانک شهر",
		"012": "بانک کارآفرین",
		"013": "بانک اقتصاد نوین",
		"014": "بانک انصار",
		"015": "بانک پست بانک",
		"016": "بانک قوامین",
		"017": "بانک مهر اقتصاد",
		"018": "بانک گردشگری",
		"019": "بانک صنعت و معدن",
		"020": "بانک کشاورزی",
		"021": "بانک مسکن",
		"022": "بانک توسعه تعاون",
		"023": "بانک رفاه کارگران",
		"024": "بانک ایران زمین",
		"025": "بانک خاورمیانه",
		"026": "بانک سینا",
		"027": "بانک توسعه",
		"028": "بانک حکمت ایرانیان",
	}

	return shebaDatabase[bankCode]
}

// getBankCodeFromName دریافت کد بانک از نام بانک
func (s *UserBankingInfoService) getBankCodeFromName(bankName string) string {
	bankCodes := map[string]string{
		"بانک ملی":           "01",
		"بانک سپه":           "02",
		"بانک تجارت":         "03",
		"بانک صادرات":        "04",
		"بانک ملت":           "05",
		"بانک پارسیان":       "06",
		"بانک پاسارگاد":      "07",
		"بانک سامان":         "08",
		"بانک آینده":         "09",
		"بانک دی":            "10",
		"بانک شهر":           "11",
		"بانک کارآفرین":      "12",
		"بانک اقتصاد نوین":   "13",
		"بانک انصار":         "14",
		"بانک پست بانک":      "15",
		"بانک قوامین":        "16",
		"بانک مهر اقتصاد":    "17",
		"بانک گردشگری":       "18",
		"بانک صنعت و معدن":   "19",
		"بانک کشاورزی":       "20",
		"بانک مسکن":          "21",
		"بانک توسعه تعاون":   "22",
		"بانک رفاه کارگران":  "23",
		"بانک ایران زمین":    "24",
		"بانک خاورمیانه":     "25",
		"بانک سینا":          "26",
		"بانک توسعه":         "27",
		"بانک حکمت ایرانیان": "28",
		"بانک مرکزی":         "00",
	}

	return bankCodes[bankName]
}

// getBINFromBankName دریافت BIN کد از نام بانک
func (s *UserBankingInfoService) getBINFromBankName(bankName string) string {
	binCodes := map[string]string{
		"بانک ملی":           "603799",
		"بانک سپه":           "589210",
		"بانک تجارت":         "627353",
		"بانک صادرات":        "627648",
		"بانک ملت":           "610433",
		"بانک پارسیان":       "622106",
		"بانک پاسارگاد":      "639347",
		"بانک سامان":         "621986",
		"بانک آینده":         "636214",
		"بانک دی":            "502938",
		"بانک شهر":           "504706",
		"بانک کارآفرین":      "627488",
		"بانک اقتصاد نوین":   "627381",
		"بانک انصار":         "627961",
		"بانک پست بانک":      "627760",
		"بانک قوامین":        "639599",
		"بانک مهر اقتصاد":    "639370",
		"بانک گردشگری":       "505416",
		"بانک صنعت و معدن":   "627961",
		"بانک کشاورزی":       "603770",
		"بانک مسکن":          "628023",
		"بانک توسعه تعاون":   "627884",
		"بانک رفاه کارگران":  "589463",
		"بانک ایران زمین":    "505416",
		"بانک خاورمیانه":     "505801",
		"بانک سینا":          "639346",
		"بانک توسعه":         "628023",
		"بانک حکمت ایرانیان": "639370",
	}

	return binCodes[bankName]
}

// toResponse تبدیل مدل به پاسخ
func (s *UserBankingInfoService) toResponse(bankingInfo *models.UserBankingInfo) *models.BankingInfoResponse {
	return &models.BankingInfoResponse{
		ID:                bankingInfo.ID,
		BankName:          bankingInfo.BankName,
		CardNumber:        bankingInfo.CardNumber,
		AccountNumber:     bankingInfo.AccountNumber,
		ShebaNumber:       bankingInfo.ShebaNumber,
		AccountHolderName: bankingInfo.AccountHolderName,
		AccountType:       bankingInfo.AccountType,
		IsDefault:         bankingInfo.IsDefault,
		IsVerified:        bankingInfo.IsVerified,
		VerifiedAt:        bankingInfo.VerifiedAt,
		CreatedAt:         bankingInfo.CreatedAt,
		UpdatedAt:         bankingInfo.UpdatedAt,
	}
}
