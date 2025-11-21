package services_test

import (
	"context"
	"errors"
	"testing"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"
	"my-go-backend/tests/mocks"
)

func TestWalletService_CreditWallet(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		userID := uint(1)
		amount := 1000.0

		mockWalletRepo := &mocks.MockWalletRepository{
			GetByUserIDFunc: func(ctx context.Context, uid uint) (*models.UserWallet, error) {
				return &models.UserWallet{ID: 1, UserID: userID, Balance: 500.0}, nil
			},
			UpdateBalanceFunc: func(ctx context.Context, walletID uint, newBalance float64) error {
				if newBalance != 1500.0 {
					t.Errorf("expected newBalance 1500.0, got %f", newBalance)
				}
				return nil
			},
		}

		mockTxRepo := &mocks.MockWalletTransactionRepository{
			CreateFunc: func(ctx context.Context, tx *models.WalletTransaction) error {
				if tx.Amount != amount {
					t.Errorf("expected tx amount %f, got %f", amount, tx.Amount)
				}
				return nil
			},
		}

		mockUoW := &mocks.MockUnitOfWork{
			WalletRepoFunc: func() repository.WalletRepositoryInterface {
				return mockWalletRepo
			},
			WalletTxRepoFunc: func() repository.WalletTransactionRepositoryInterface {
				return mockTxRepo
			},
			CommitFunc: func() error { return nil },
		}

		mockUoWFactory := &mocks.MockUnitOfWorkFactory{
			CreateFunc: func() unitofwork.UnitOfWork {
				return mockUoW
			},
		}

		service := services.NewWalletService(mockUoWFactory)

		err := service.CreditWallet(context.Background(), 0, userID, amount, "manual", "admin", nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("WalletNotFound", func(t *testing.T) {
		userID := uint(1)
		mockWalletRepo := &mocks.MockWalletRepository{
			GetByUserIDFunc: func(ctx context.Context, uid uint) (*models.UserWallet, error) {
				return nil, errors.New("record not found")
			},
		}

		mockUoW := &mocks.MockUnitOfWork{
			WalletRepoFunc: func() repository.WalletRepositoryInterface {
				return mockWalletRepo
			},
			RollbackFunc: func() error { return nil },
		}

		mockUoWFactory := &mocks.MockUnitOfWorkFactory{
			CreateFunc: func() unitofwork.UnitOfWork {
				return mockUoW
			},
		}

		service := services.NewWalletService(mockUoWFactory)

		err := service.CreditWallet(context.Background(), 0, userID, 100.0, "manual", "admin", nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
