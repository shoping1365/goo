package services

import (
	"context"
	"fmt"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/repository"
)

// FirewallService سرویس فایروال (CQRS + UnitOfWork)
// توضیحات: تمام تراکنش‌ها از طریق UoW مدیریت شده و از ریپازیتوری OS-specific استفاده می‌شود
type FirewallService struct {
	uowFactory unitofwork.UnitOfWorkFactory
	repo       repository.FirewallRepositoryInterface
}

// NewFirewallService سازنده سرویس فایروال
func NewFirewallService(uowFactory unitofwork.UnitOfWorkFactory, repo repository.FirewallRepositoryInterface) *FirewallService {
	return &FirewallService{uowFactory: uowFactory, repo: repo}
}

// Query APIs
func (s *FirewallService) GetStatus(ctx context.Context) (*repository.FirewallStatus, error) {
	return s.repo.GetStatus(ctx)
}

func (s *FirewallService) ListRules(ctx context.Context) ([]repository.FirewallRule, error) {
	return s.repo.ListRules(ctx)
}

func (s *FirewallService) ReadLogs(ctx context.Context, limit int) ([]repository.FirewallLog, error) {
	return s.repo.ReadLogs(ctx, limit)
}

// Command APIs (managed via UoW)
func (s *FirewallService) Toggle(ctx context.Context, enabled bool) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer uow.Rollback()
	if err := s.repo.SetEnabled(ctx, enabled); err != nil {
		return err
	}
	return uow.Commit()
}

func (s *FirewallService) SetMode(ctx context.Context, mode repository.FirewallMode) error {
	if mode != repository.FirewallModeLow && mode != repository.FirewallModeMedium && mode != repository.FirewallModeHigh && mode != repository.FirewallModeStrict {
		return fmt.Errorf("invalid firewall mode: %s", mode)
	}
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer uow.Rollback()
	if err := s.repo.SetMode(ctx, mode); err != nil {
		return err
	}
	return uow.Commit()
}

func (s *FirewallService) CreateRule(ctx context.Context, rule repository.CreateFirewallRule) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer uow.Rollback()
	if err := s.repo.CreateRule(ctx, rule); err != nil {
		return err
	}
	return uow.Commit()
}

func (s *FirewallService) UpdateRule(ctx context.Context, name string, upd repository.UpdateFirewallRule) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer uow.Rollback()
	if err := s.repo.UpdateRule(ctx, name, upd); err != nil {
		return err
	}
	return uow.Commit()
}

func (s *FirewallService) ToggleRule(ctx context.Context, name string, enabled bool) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer uow.Rollback()
	if err := s.repo.ToggleRule(ctx, name, enabled); err != nil {
		return err
	}
	return uow.Commit()
}

func (s *FirewallService) DeleteRule(ctx context.Context, name string) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer uow.Rollback()
	if err := s.repo.DeleteRule(ctx, name); err != nil {
		return err
	}
	return uow.Commit()
}

func (s *FirewallService) ClearLogs(ctx context.Context) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer uow.Rollback()
	if err := s.repo.ClearLogs(ctx); err != nil {
		return err
	}
	return uow.Commit()
}

func (s *FirewallService) Backup(ctx context.Context, dest string) (string, error) {
	// Backup نیازی به تراکنش دیتابیس ندارد اما برای یکنواختی استفاده می‌کنیم
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return "", err
	}
	defer uow.Rollback()
	p, err := s.repo.Backup(ctx, dest)
	if err != nil {
		return "", err
	}
	if err := uow.Commit(); err != nil {
		return "", err
	}
	return p, nil
}

// Logging config
func (s *FirewallService) GetLogging(ctx context.Context) (*repository.FirewallLoggingConfig, error) {
	return s.repo.GetLogging(ctx)
}

func (s *FirewallService) SetLogging(ctx context.Context, cfg repository.FirewallLoggingConfig) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer uow.Rollback()
	if err := s.repo.SetLogging(ctx, cfg); err != nil {
		return err
	}
	return uow.Commit()
}

func (s *FirewallService) Restore(ctx context.Context, source string) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer uow.Rollback()
	if err := s.repo.Restore(ctx, source); err != nil {
		return err
	}
	return uow.Commit()
}
