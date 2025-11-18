package repository

import (
	"context"
	"errors"
	"my-go-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

type TrafficSessionRepositoryInterface interface {
	CountOnline(ctx context.Context, since time.Time) (int64, error)
	ListOnline(ctx context.Context, since time.Time, page, limit int) ([]models.TrafficSession, int64, error)
	DeactivateByIP(ctx context.Context, ip string) error
	CleanupExpired(ctx context.Context, before time.Time) error
	GetBySessionID(ctx context.Context, sessionID string) (*models.TrafficSession, error)
	Create(ctx context.Context, sess *models.TrafficSession) error
	Save(ctx context.Context, sess *models.TrafficSession) error
}

type TrafficSessionRepository struct{ DB *gorm.DB }

func (r *TrafficSessionRepository) CountOnline(ctx context.Context, since time.Time) (int64, error) {
	var total int64
	err := r.DB.WithContext(ctx).Model(&models.TrafficSession{}).
		Where("is_active = ? AND last_seen > ?", true, since).Count(&total).Error
	return total, err
}

func (r *TrafficSessionRepository) ListOnline(ctx context.Context, since time.Time, page, limit int) ([]models.TrafficSession, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 1000 {
		limit = 50
	}
	q := r.DB.WithContext(ctx).Model(&models.TrafficSession{}).
		Where("is_active = ? AND last_seen > ?", true, since)
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var sessions []models.TrafficSession
	if err := q.Preload("User").Order("last_seen DESC").Offset((page - 1) * limit).Limit(limit).Find(&sessions).Error; err != nil {
		return nil, 0, err
	}
	return sessions, total, nil
}

func (r *TrafficSessionRepository) DeactivateByIP(ctx context.Context, ip string) error {
	return r.DB.WithContext(ctx).Model(&models.TrafficSession{}).Where("ip_address = ?", ip).Update("is_active", false).Error
}

func (r *TrafficSessionRepository) CleanupExpired(ctx context.Context, before time.Time) error {
	return r.DB.WithContext(ctx).Model(&models.TrafficSession{}).Where("last_seen < ?", before).Update("is_active", false).Error
}

func (r *TrafficSessionRepository) GetBySessionID(ctx context.Context, sessionID string) (*models.TrafficSession, error) {
	var s models.TrafficSession
	if err := r.DB.WithContext(ctx).Where("session_id = ?", sessionID).First(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &s, nil
}

func (r *TrafficSessionRepository) Create(ctx context.Context, sess *models.TrafficSession) error {
	return r.DB.WithContext(ctx).Create(sess).Error
}

func (r *TrafficSessionRepository) Save(ctx context.Context, sess *models.TrafficSession) error {
	return r.DB.WithContext(ctx).Save(sess).Error
}
