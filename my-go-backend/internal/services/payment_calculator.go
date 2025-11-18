package services

import (
	"my-go-backend/internal/models"
)

// PaymentCalculator سرویس محاسبه کارمزد و محدودیت‌های پرداخت
// شامل منطق صفر برای نادیده گرفتن محدودیت‌ها

type PaymentCalculator struct{}

// NewPaymentCalculator سازنده سرویس محاسبه‌گر
func NewPaymentCalculator() *PaymentCalculator {
	return &PaymentCalculator{}
}

// CalculateFee محاسبه کارمزد بر اساس درصد درگاه
// اگر کارمزد صفر باشد، کارمزدی محاسبه نمی‌شود
func (pc *PaymentCalculator) CalculateFee(amount int64, gateway *models.PaymentGateway) int64 {
	// اگر کارمزد صفر باشد، کارمزدی محاسبه نمی‌شود
	if gateway.Fee <= 0 {
		return 0
	}

	// محاسبه کارمزد
	feeAmount := int64(float64(amount) * gateway.Fee / 100)
	return feeAmount
}

// GetTotalAmount محاسبه مبلغ کل شامل کارمزد
func (pc *PaymentCalculator) GetTotalAmount(amount int64, gateway *models.PaymentGateway) int64 {
	fee := pc.CalculateFee(amount, gateway)
	return amount + fee
}

// ValidateAmount بررسی اعتبار مبلغ بر اساس محدودیت‌های درگاه
// اگر حداقل یا حداکثر صفر باشد، نادیده گرفته می‌شود
func (pc *PaymentCalculator) ValidateAmount(amount int64, gateway *models.PaymentGateway) (bool, string) {
	// بررسی حداقل مبلغ - اگر صفر باشد نادیده گرفته می‌شود
	if gateway.MinAmount > 0 && amount < gateway.MinAmount {
		return false, "مبلغ کمتر از حداقل مجاز است"
	}

	// بررسی حداکثر مبلغ - اگر صفر باشد نادیده گرفته می‌شود
	if gateway.MaxAmount > 0 && amount > gateway.MaxAmount {
		return false, "مبلغ بیشتر از حداکثر مجاز است"
	}

	return true, ""
}

// GetAmountLimits دریافت محدودیت‌های مبلغ درگاه
// اگر صفر باشد، محدودیتی اعمال نمی‌شود
func (pc *PaymentCalculator) GetAmountLimits(gateway *models.PaymentGateway) (minAmount, maxAmount int64, hasMinLimit, hasMaxLimit bool) {
	hasMinLimit = gateway.MinAmount > 0
	hasMaxLimit = gateway.MaxAmount > 0

	if hasMinLimit {
		minAmount = gateway.MinAmount
	}
	if hasMaxLimit {
		maxAmount = gateway.MaxAmount
	}

	return
}

// GetFeeInfo دریافت اطلاعات کارمزد
// اگر کارمزد صفر باشد، کارمزدی اعمال نمی‌شود
func (pc *PaymentCalculator) GetFeeInfo(gateway *models.PaymentGateway) (fee float64, hasFee bool) {
	hasFee = gateway.Fee > 0
	if hasFee {
		fee = gateway.Fee
	}
	return
}

// CalculateBreakdown محاسبه تفصیلی مبلغ و کارمزد
type PaymentBreakdown struct {
	OriginalAmount int64   `json:"original_amount"` // مبلغ اصلی
	FeeAmount      int64   `json:"fee_amount"`      // مبلغ کارمزد
	TotalAmount    int64   `json:"total_amount"`    // مبلغ کل
	FeePercentage  float64 `json:"fee_percentage"`  // درصد کارمزد
	HasFee         bool    `json:"has_fee"`         // آیا کارمزد دارد
}

func (pc *PaymentCalculator) CalculateBreakdown(amount int64, gateway *models.PaymentGateway) *PaymentBreakdown {
	feeAmount := pc.CalculateFee(amount, gateway)
	hasFee := gateway.Fee > 0

	return &PaymentBreakdown{
		OriginalAmount: amount,
		FeeAmount:      feeAmount,
		TotalAmount:    amount + feeAmount,
		FeePercentage:  gateway.Fee,
		HasFee:         hasFee,
	}
}
