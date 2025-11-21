package services_test

import (
	"context"
	"testing"
	"time"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/services"
	"my-go-backend/internal/repository"
	"my-go-backend/tests/mocks"
)

func TestTrafficService_BlockIP(t *testing.T) {
	t.Run("BlockIP", func(t *testing.T) {
		ip := "192.168.1.1"
		reason := "Test Block"
		blockedBy := "admin"
		durationHours := 24

		mockBlockedIPRepo := &mocks.MockBlockedIPRepository{
			CreateOrReactivateFunc: func(ctx context.Context, i, r, b string, e *time.Time) error {
				if i != ip {
					t.Errorf("expected ip %s, got %s", ip, i)
				}
				if r != reason {
					t.Errorf("expected reason %s, got %s", reason, r)
				}
				return nil
			},
		}

		mockTrafficSessionRepo := &mocks.MockTrafficSessionRepository{
			DeactivateByIPFunc: func(ctx context.Context, i string) error {
				if i != ip {
					t.Errorf("expected ip %s, got %s", ip, i)
				}
				return nil
			},
		}

		mockUoW := &mocks.MockUnitOfWork{
			BlockedIPRepoFunc: func() repository.BlockedIPRepositoryInterface {
				return mockBlockedIPRepo
			},
			TrafficSessionRepoFunc: func() repository.TrafficSessionRepositoryInterface {
				return mockTrafficSessionRepo
			},
			CommitFunc: func() error { return nil },
		}

		mockUoWFactory := &mocks.MockUnitOfWorkFactory{
			CreateFunc: func() unitofwork.UnitOfWork {
				return mockUoW
			},
		}

		service := services.NewTrafficService(mockUoWFactory)

		err := service.BlockIP(context.Background(), ip, reason, blockedBy, durationHours)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("UnblockIP", func(t *testing.T) {
		ip := "192.168.1.1"

		mockBlockedIPRepo := &mocks.MockBlockedIPRepository{
			DeactivateFunc: func(ctx context.Context, i string) error {
				if i != ip {
					t.Errorf("expected ip %s, got %s", ip, i)
				}
				return nil
			},
		}

		mockUoW := &mocks.MockUnitOfWork{
			BlockedIPRepoFunc: func() repository.BlockedIPRepositoryInterface {
				return mockBlockedIPRepo
			},
			CommitFunc: func() error { return nil },
		}

		mockUoWFactory := &mocks.MockUnitOfWorkFactory{
			CreateFunc: func() unitofwork.UnitOfWork {
				return mockUoW
			},
		}

		service := services.NewTrafficService(mockUoWFactory)

		err := service.UnblockIP(context.Background(), ip)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
