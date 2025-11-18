package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// ChatSessionRepository پیاده‌سازی ChatSessionRepositoryInterface
// Rule 0: هیچ منطق تجاری اینجا قرار نمی‌گیرد؛ فقط دسترسی داده.
type ChatSessionRepository struct {
	DB *gorm.DB
}

func (r *ChatSessionRepository) Create(ctx context.Context, session *models.ChatSession) error {
	return r.DB.WithContext(ctx).Create(session).Error
}

func (r *ChatSessionRepository) GetByID(ctx context.Context, id uint) (*models.ChatSession, error) {
	var s models.ChatSession
	err := r.DB.WithContext(ctx).First(&s, id).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *ChatSessionRepository) GetBySessionID(ctx context.Context, sessionID string) (*models.ChatSession, error) {
	var s models.ChatSession
	err := r.DB.WithContext(ctx).Where("session_id = ?", sessionID).First(&s).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *ChatSessionRepository) GetWaitingSessions(ctx context.Context) ([]models.ChatSession, error) {
	var sessions []models.ChatSession
	err := r.DB.WithContext(ctx).Where("status = ?", "waiting").Order("created_at ASC").Find(&sessions).Error
	return sessions, err
}

func (r *ChatSessionRepository) ListSessionsByStatus(ctx context.Context, status string, limit, offset int) ([]models.ChatSession, error) {
	var sessions []models.ChatSession
	qb := r.DB.WithContext(ctx).Model(&models.ChatSession{}).
		Preload("Operator").
		Preload("Widget").
		Order("updated_at DESC")
	if status != "" && status != "all" {
		qb = qb.Where("status = ?", status)
	}
	if limit > 0 {
		qb = qb.Limit(limit)
	}
	if offset > 0 {
		qb = qb.Offset(offset)
	}
	err := qb.Find(&sessions).Error
	return sessions, err
}

func (r *ChatSessionRepository) UpdateFields(ctx context.Context, id uint, fields map[string]interface{}) error {
	return r.DB.WithContext(ctx).Model(&models.ChatSession{}).Where("id = ?", id).Updates(fields).Error
}

// ChatMessageRepository پیاده‌سازی ChatMessageRepositoryInterface

type ChatMessageRepository struct {
	DB *gorm.DB
}

func (r *ChatMessageRepository) Create(ctx context.Context, msg *models.OnlineChatMessage) error {
	return r.DB.WithContext(ctx).Create(msg).Error
}

func (r *ChatMessageRepository) GetBySessionID(ctx context.Context, sessionID uint, limit, offset int) ([]models.OnlineChatMessage, error) {
	var messages []models.OnlineChatMessage
	err := r.DB.WithContext(ctx).Where("session_id = ?", sessionID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&messages).Error
	return messages, err
}

func (r *ChatMessageRepository) MarkAsRead(ctx context.Context, messageID uint) error {
	return r.DB.WithContext(ctx).Model(&models.OnlineChatMessage{}).Where("id = ?", messageID).Updates(map[string]interface{}{"is_read": true, "read_at": gorm.Expr("NOW()")}).Error
}

func (r *ChatMessageRepository) MarkSessionMessagesAsRead(ctx context.Context, sessionID uint) error {
	return r.DB.WithContext(ctx).Model(&models.OnlineChatMessage{}).Where("session_id = ? AND is_read = ?", sessionID, false).Updates(map[string]interface{}{"is_read": true, "read_at": gorm.Expr("NOW()")}).Error
}

// ChatOperatorRepository پیاده‌سازی ChatOperatorRepositoryInterface
type ChatOperatorRepository struct {
	DB *gorm.DB
}

func (r *ChatOperatorRepository) GetByID(ctx context.Context, id uint) (*models.ChatOperator, error) {
	var op models.ChatOperator
	if err := r.DB.WithContext(ctx).First(&op, id).Error; err != nil {
		return nil, err
	}
	return &op, nil
}

func (r *ChatOperatorRepository) UpdateFields(ctx context.Context, id uint, fields map[string]interface{}) error {
	return r.DB.WithContext(ctx).Model(&models.ChatOperator{}).Where("id = ?", id).Updates(fields).Error
}

// ==================== Widget Repository ====================

type ChatWidgetRepository struct {
	DB *gorm.DB
}

func (r *ChatWidgetRepository) Create(ctx context.Context, widget *models.ChatWidget) error {
	return r.DB.WithContext(ctx).Create(widget).Error
}

func (r *ChatWidgetRepository) GetActive(ctx context.Context) ([]models.ChatWidget, error) {
	var items []models.ChatWidget
	err := r.DB.WithContext(ctx).Where("is_active = ?", true).Find(&items).Error
	return items, err
}

func (r *ChatWidgetRepository) GetByID(ctx context.Context, id uint) (*models.ChatWidget, error) {
	var w models.ChatWidget
	if err := r.DB.WithContext(ctx).First(&w, id).Error; err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *ChatWidgetRepository) Update(ctx context.Context, widget *models.ChatWidget) error {
	return r.DB.WithContext(ctx).Save(widget).Error
}

func (r *ChatWidgetRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.ChatWidget{}, id).Error
}

// ==================== AI Bot Repository ====================

type ChatAIBotRepository struct {
	DB *gorm.DB
}

func (r *ChatAIBotRepository) Create(ctx context.Context, bot *models.ChatAIBot) error {
	return r.DB.WithContext(ctx).Create(bot).Error
}

func (r *ChatAIBotRepository) GetActive(ctx context.Context) ([]models.ChatAIBot, error) {
	var bots []models.ChatAIBot
	err := r.DB.WithContext(ctx).Where("is_active = ?", true).Find(&bots).Error
	return bots, err
}

func (r *ChatAIBotRepository) GetByID(ctx context.Context, id uint) (*models.ChatAIBot, error) {
	var bot models.ChatAIBot
	if err := r.DB.WithContext(ctx).First(&bot, id).Error; err != nil {
		return nil, err
	}
	return &bot, nil
}

func (r *ChatAIBotRepository) Update(ctx context.Context, bot *models.ChatAIBot) error {
	return r.DB.WithContext(ctx).Save(bot).Error
}

func (r *ChatAIBotRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.ChatAIBot{}, id).Error
}

// ==================== Knowledge Base Repository ====================

type ChatKnowledgeBaseRepository struct {
	DB *gorm.DB
}

func (r *ChatKnowledgeBaseRepository) Create(ctx context.Context, kb *models.ChatKnowledgeBase) error {
	return r.DB.WithContext(ctx).Create(kb).Error
}

func (r *ChatKnowledgeBaseRepository) Search(ctx context.Context, q string) ([]models.ChatKnowledgeBase, error) {
	var res []models.ChatKnowledgeBase
	like := "%" + q + "%"
	err := r.DB.WithContext(ctx).Where("is_active = ? AND (title ILIKE ? OR content ILIKE ? OR tags ILIKE ?)", true, like, like, like).Find(&res).Error
	return res, err
}

func (r *ChatKnowledgeBaseRepository) List(ctx context.Context, q string, limit, offset int) ([]models.ChatKnowledgeBase, int64, error) {
	var total int64
	qb := r.DB.WithContext(ctx).Model(&models.ChatKnowledgeBase{}).Where("is_active = ?", true)
	if q != "" {
		like := "%" + q + "%"
		qb = qb.Where("title ILIKE ? OR content ILIKE ? OR tags ILIKE ?", like, like, like)
	}
	if err := qb.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var items []models.ChatKnowledgeBase
	if err := qb.Order("updated_at DESC").Limit(limit).Offset(offset).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *ChatKnowledgeBaseRepository) GetByID(ctx context.Context, id uint) (*models.ChatKnowledgeBase, error) {
	var item models.ChatKnowledgeBase
	if err := r.DB.WithContext(ctx).First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ChatKnowledgeBaseRepository) Update(ctx context.Context, kb *models.ChatKnowledgeBase) error {
	return r.DB.WithContext(ctx).Save(kb).Error
}

func (r *ChatKnowledgeBaseRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.ChatKnowledgeBase{}, id).Error
}
