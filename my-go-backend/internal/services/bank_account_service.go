package services

import (
	"context"

	"my-go-backend/internal/database/unitofwork"
)

// BankAccountService سرویس مدیریت کارت/حساب بانکی کاربران
type BankAccountService struct {
	uowFactory unitofwork.UnitOfWorkFactory
}

func NewBankAccountService(uowFactory unitofwork.UnitOfWorkFactory) *BankAccountService {
	return &BankAccountService{uowFactory: uowFactory}
}

func (s *BankAccountService) List(ctx context.Context, page, pageSize int, status, bank string) (items any, total int64, err error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return s.uowFactory.Create().BankAccountRepository().List(ctx, pageSize, offset, status, bank)
}

func (s *BankAccountService) Verify(ctx context.Context, id uint) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() { _ = uow.Rollback() }()
	if err := uow.BankAccountRepository().Verify(ctx, id); err != nil {
		return err
	}
	return uow.Commit()
}

func (s *BankAccountService) Block(ctx context.Context, id uint) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() { _ = uow.Rollback() }()
	if err := uow.BankAccountRepository().Block(ctx, id); err != nil {
		return err
	}
	return uow.Commit()
}

func (s *BankAccountService) Unblock(ctx context.Context, id uint) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() { _ = uow.Rollback() }()
	if err := uow.BankAccountRepository().Unblock(ctx, id); err != nil {
		return err
	}
	return uow.Commit()
}

func (s *BankAccountService) Reject(ctx context.Context, id uint) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() { _ = uow.Rollback() }()
	if err := uow.BankAccountRepository().Reject(ctx, id); err != nil {
		return err
	}
	return uow.Commit()
}
