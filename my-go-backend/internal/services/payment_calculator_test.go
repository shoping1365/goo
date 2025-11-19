package services

import (
	"testing"

	"my-go-backend/internal/models"
)

func TestCalculateFee(t *testing.T) {
	pc := NewPaymentCalculator()

	tests := []struct {
		name     string
		amount   int64
		fee      float64
		expected int64
	}{
		{"No Fee", 10000, 0, 0},
		{"10% Fee", 10000, 10, 1000},
		{"5.5% Fee", 10000, 5.5, 550},
		{"Zero Amount", 0, 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gateway := &models.PaymentGateway{Fee: tt.fee}
			result := pc.CalculateFee(tt.amount, gateway)
			if result != tt.expected {
				t.Errorf("CalculateFee() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetTotalAmount(t *testing.T) {
	pc := NewPaymentCalculator()
	gateway := &models.PaymentGateway{Fee: 10} // 10% fee
	amount := int64(1000)
	expected := int64(1100)

	result := pc.GetTotalAmount(amount, gateway)
	if result != expected {
		t.Errorf("GetTotalAmount() = %v, want %v", result, expected)
	}
}

func TestValidateAmount(t *testing.T) {
	pc := NewPaymentCalculator()

	tests := []struct {
		name      string
		amount    int64
		min       int64
		max       int64
		isValid   bool
	}{
		{"Valid Amount", 5000, 1000, 10000, true},
		{"Below Min", 500, 1000, 10000, false},
		{"Above Max", 15000, 1000, 10000, false},
		{"No Limits", 5000, 0, 0, true},
		{"Only Min Valid", 5000, 1000, 0, true},
		{"Only Min Invalid", 500, 1000, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gateway := &models.PaymentGateway{MinAmount: tt.min, MaxAmount: tt.max}
			valid, _ := pc.ValidateAmount(tt.amount, gateway)
			if valid != tt.isValid {
				t.Errorf("ValidateAmount() valid = %v, want %v", valid, tt.isValid)
			}
		})
	}
}

func TestGetAmountLimits(t *testing.T) {
	pc := NewPaymentCalculator()
	gateway := &models.PaymentGateway{MinAmount: 1000, MaxAmount: 5000}

	min, max, hasMin, hasMax := pc.GetAmountLimits(gateway)

	if min != 1000 || !hasMin {
		t.Error("GetAmountLimits failed for min amount")
	}
	if max != 5000 || !hasMax {
		t.Error("GetAmountLimits failed for max amount")
	}

	gatewayNoLimit := &models.PaymentGateway{}
	_, _, hasMin2, hasMax2 := pc.GetAmountLimits(gatewayNoLimit)
	if hasMin2 || hasMax2 {
		t.Error("GetAmountLimits should return false for zero limits")
	}
}

func TestGetFeeInfo(t *testing.T) {
	pc := NewPaymentCalculator()
	gateway := &models.PaymentGateway{Fee: 2.5}

	fee, hasFee := pc.GetFeeInfo(gateway)
	if !hasFee || fee != 2.5 {
		t.Error("GetFeeInfo failed")
	}

	gatewayNoFee := &models.PaymentGateway{Fee: 0}
	_, hasFee2 := pc.GetFeeInfo(gatewayNoFee)
	if hasFee2 {
		t.Error("GetFeeInfo should return false for zero fee")
	}
}

func TestCalculateBreakdown(t *testing.T) {
	pc := NewPaymentCalculator()
	gateway := &models.PaymentGateway{Fee: 10}
	amount := int64(1000)

	breakdown := pc.CalculateBreakdown(amount, gateway)

	if breakdown.OriginalAmount != 1000 {
		t.Errorf("Breakdown OriginalAmount = %v, want 1000", breakdown.OriginalAmount)
	}
	if breakdown.FeeAmount != 100 {
		t.Errorf("Breakdown FeeAmount = %v, want 100", breakdown.FeeAmount)
	}
	if breakdown.TotalAmount != 1100 {
		t.Errorf("Breakdown TotalAmount = %v, want 1100", breakdown.TotalAmount)
	}
	if !breakdown.HasFee {
		t.Error("Breakdown HasFee should be true")
	}
}
