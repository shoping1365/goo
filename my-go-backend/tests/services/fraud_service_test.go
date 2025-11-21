package services_test

import (
	"context"
	"testing"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"
	"my-go-backend/tests/mocks"
)

func TestFraudService_EvaluateOrder(t *testing.T) {
	t.Run("HighRisk", func(t *testing.T) {
		input := services.EvaluateOrderInput{
			OrderID:          1,
			Amount:           1000000,
			ShippingCity:     "Tehran",
			GeoIPCity:        "Shiraz", // Mismatch
			ShippingProvince: "Tehran",
			GeoIPCountry:     "IR",
		}

		mockRuleRepo := &mocks.MockFraudRuleRepository{
			ListActiveFunc: func(ctx context.Context) ([]repository.FraudRule, error) {
				return []repository.FraudRule{
					{Key: "geoip_mismatch_city", Weight: 50, Enabled: true},
				}, nil
			},
		}

		mockCaseRepo := &mocks.MockFraudCaseRepository{
			CreateFunc: func(ctx context.Context, c *repository.FraudCase) error {
				if c.RiskScore != 50 {
					t.Errorf("expected risk score 50, got %d", c.RiskScore)
				}
				return nil
			},
		}

		mockEventRepo := &mocks.MockFraudEventRepository{
			CreateManyFunc: func(ctx context.Context, events []repository.FraudEvent) error {
				if len(events) != 1 {
					t.Errorf("expected 1 event, got %d", len(events))
				}
				if events[0].EventType != "geoip_mismatch_city" {
					t.Errorf("expected event type geoip_mismatch_city, got %s", events[0].EventType)
				}
				return nil
			},
		}

		mockUoW := &mocks.MockUnitOfWork{
			FraudRuleRepoFunc: func() repository.FraudRuleRepositoryInterface {
				return mockRuleRepo
			},
			FraudCaseRepoFunc: func() repository.FraudCaseRepositoryInterface {
				return mockCaseRepo
			},
			FraudEventRepoFunc: func() repository.FraudEventRepositoryInterface {
				return mockEventRepo
			},
			CommitFunc: func() error { return nil },
		}

		mockUoWFactory := &mocks.MockUnitOfWorkFactory{
			CreateFunc: func() unitofwork.UnitOfWork {
				return mockUoW
			},
		}

		service := services.NewFraudService(mockUoWFactory, "salt")

		result, err := service.EvaluateOrder(context.Background(), input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result.RiskScore != 50 {
			t.Errorf("expected result risk score 50, got %d", result.RiskScore)
		}
	})
}
