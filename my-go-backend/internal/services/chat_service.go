package services

import (
	"context"
	"errors"
	"fmt"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

// ChatService سرویس مدیریت چت (استفاده از UnitOfWork طبق Rule0)
type ChatService struct {
	uowFactory    unitofwork.UnitOfWorkFactory
	encryptionSvc *EncryptionService
	db            *gorm.DB // باقی‌مانده متدهای قدیمی هنوز از db استفاده می‌کنند؛ در نسخه بعد حذف می‌شود
}

// NewChatService ایجاد Chat Service جدید
func NewChatService(uowFactory unitofwork.UnitOfWorkFactory) (*ChatService, error) {
	encryptionSvc, err := NewEncryptionService()
	if err != nil {
		return nil, fmt.Errorf("خطا در ایجاد سرویس رمزنگاری: %v", err)
	}

	uow := uowFactory.Create()

	return &ChatService{
		uowFactory:    uowFactory,
		encryptionSvc: encryptionSvc,
		db:            uow.DB(),
	}, nil
}

// CreateChatSession ایجاد جلسه چت جدید (با UnitOfWork)
func (cs *ChatService) CreateChatSession(sessionData *models.ChatSession) (*models.ChatSession, error) {
	ctx := context.Background()

	// رمزنگاری نام و شماره تلفن مشتری
	if sessionData.CustomerName != "" {
		encryptedName, err := cs.encryptionSvc.EncryptUserName(sessionData.CustomerName)
		if err != nil {
			return nil, fmt.Errorf("خطا در رمزنگاری نام مشتری: %v", err)
		}
		sessionData.CustomerName = encryptedName
	}

	if sessionData.CustomerPhone != "" {
		encryptedPhone, err := cs.encryptionSvc.EncryptPhoneNumber(sessionData.CustomerPhone)
		if err != nil {
			return nil, fmt.Errorf("خطا در رمزنگاری شماره تلفن مشتری: %v", err)
		}
		sessionData.CustomerPhone = encryptedPhone
	}

	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return nil, err
	}

	defer uow.Rollback() // در صورت خطا

	chatRepo := uow.ChatSessionRepository()

	existing, err := chatRepo.GetWaitingSessions(ctx)
	if err == nil {
		for _, s := range existing {
			if s.CustomerPhone == sessionData.CustomerPhone {
				return &s, nil
			}
		}
	}

	sessionData.Status = "waiting"
	sessionData.StartedAt = time.Now()

	if err := chatRepo.Create(ctx, sessionData); err != nil {
		return nil, err
	}

	if err := uow.Commit(); err != nil {
		return nil, err
	}

	return sessionData, nil
}

// GetChatSession دریافت جلسه چت
func (cs *ChatService) GetChatSession(sessionID string) (*models.ChatSession, error) {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return nil, err
	}

	sessionRepo := uow.ChatSessionRepository()
	session, err := sessionRepo.GetBySessionID(ctx, sessionID)
	if err != nil {
		uow.Rollback()
		return nil, err
	}

	// TODO: preload related entities via separate queries if needed

	_ = uow.Commit()

	// رمزگشایی نام و شماره تلفن مشتری
	if session.CustomerName != "" {
		decryptedName, err := cs.encryptionSvc.DecryptUserName(session.CustomerName)
		if err != nil {
			return nil, fmt.Errorf("خطا در رمزگشایی نام مشتری: %v", err)
		}
		session.CustomerName = decryptedName
	}

	if session.CustomerPhone != "" {
		decryptedPhone, err := cs.encryptionSvc.DecryptPhoneNumber(session.CustomerPhone)
		if err != nil {
			return nil, fmt.Errorf("خطا در رمزگشایی شماره تلفن مشتری: %v", err)
		}
		session.CustomerPhone = decryptedPhone
	}

	// رمزگشایی پیام‌ها
	for i := range session.Messages {
		if session.Messages[i].Message != "" {
			decryptedMessage, err := cs.encryptionSvc.DecryptMessage(session.Messages[i].Message)
			if err != nil {
				return nil, fmt.Errorf("خطا در رمزگشایی پیام: %v", err)
			}
			session.Messages[i].Message = decryptedMessage
		}
	}

	return session, nil
}

// GetWaitingSessions دریافت جلسه‌های در انتظار
func (cs *ChatService) GetWaitingSessions() ([]models.ChatSession, error) {
	var sessions []models.ChatSession
	err := cs.db.Where("status = ?", "waiting").
		Preload("Operator").
		Preload("Widget").
		Order("created_at ASC").
		Find(&sessions).Error

	return sessions, err
}

// GetSessionsByStatus دریافت لیست جلسات بر اساس وضعیت (اگر status خالی باشد، همه)
// این متد برای پنل ادمین استفاده می‌شود و پیام آخر و شمارش خوانده‌نشده را می‌توان در لایه query اضافه کرد
func (cs *ChatService) GetSessionsByStatus(status string) ([]models.ChatSession, error) {
	var sessions []models.ChatSession
	q := cs.db.Model(&models.ChatSession{}).
		Preload("Operator").
		Preload("Widget").
		Order("updated_at DESC")

	if status != "" && status != "all" {
		q = q.Where("status = ?", status)
	}

	if err := q.Find(&sessions).Error; err != nil {
		return nil, err
	}

	// رمزگشایی نام/تلفن در صورت نیاز
	for i := range sessions {
		if sessions[i].CustomerName != "" {
			if name, err := cs.encryptionSvc.DecryptUserName(sessions[i].CustomerName); err == nil {
				sessions[i].CustomerName = name
			}
		}
		if sessions[i].CustomerPhone != "" {
			if phone, err := cs.encryptionSvc.DecryptPhoneNumber(sessions[i].CustomerPhone); err == nil {
				sessions[i].CustomerPhone = phone
			}
		}
	}

	return sessions, nil
}

// ListSessionsByStatus لیست جلسات با فیلتر وضعیت و صفحه‌بندی
// اگر status خالی یا "all" باشد، همه جلسات بازگردانده می‌شود. limit و offset اختیاری‌اند.
func (cs *ChatService) ListSessionsByStatus(status string, limit, offset int) ([]models.ChatSession, error) {
	var sessions []models.ChatSession
	qb := cs.db.Model(&models.ChatSession{}).
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

	if err := qb.Find(&sessions).Error; err != nil {
		return nil, err
	}

	// رمزگشایی نام/تلفن در صورت نیاز
	for i := range sessions {
		if sessions[i].CustomerName != "" {
			if name, err := cs.encryptionSvc.DecryptUserName(sessions[i].CustomerName); err == nil {
				sessions[i].CustomerName = name
			}
		}
		if sessions[i].CustomerPhone != "" {
			if phone, err := cs.encryptionSvc.DecryptPhoneNumber(sessions[i].CustomerPhone); err == nil {
				sessions[i].CustomerPhone = phone
			}
		}
	}

	return sessions, nil
}

// AssignOperatorToSession تخصیص اپراتور به جلسه
func (cs *ChatService) AssignOperatorToSession(sessionID string, operatorID uint) error {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}

	opRepo := uow.ChatOperatorRepository()
	sessionRepo := uow.ChatSessionRepository()

	operator, err := opRepo.GetByID(ctx, operatorID)
	if err != nil || !operator.IsActive {
		uow.Rollback()
		return errors.New("اپراتور یافت نشد")
	}
	if operator.CurrentChats >= operator.MaxConcurrentChats {
		uow.Rollback()
		return errors.New("اپراتور ظرفیت کافی ندارد")
	}

	// پیدا کردن جلسه
	s, err := sessionRepo.GetBySessionID(ctx, sessionID)
	if err != nil {
		uow.Rollback()
		return err
	}
	if err := sessionRepo.UpdateFields(ctx, s.ID, map[string]interface{}{"operator_id": operatorID, "status": "active"}); err != nil {
		uow.Rollback()
		return err
	}
	if err := opRepo.UpdateFields(ctx, operatorID, map[string]interface{}{"current_chats": operator.CurrentChats + 1}); err != nil {
		uow.Rollback()
		return err
	}
	return uow.Commit()
}

// SaveMessage ذخیره پیام جدید با UnitOfWork
func (cs *ChatService) SaveMessage(message *models.OnlineChatMessage) error {
	ctx := context.Background()

	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}

	sessionRepo := uow.ChatSessionRepository()
	msgRepo := uow.ChatMessageRepository()

	session, err := sessionRepo.GetByID(ctx, message.SessionID)
	if err != nil {
		uow.Rollback()
		return errors.New("جلسه چت یافت نشد")
	}

	// رمزنگاری محتوای پیام
	if message.Message != "" {
		encryptedMessage, err := cs.encryptionSvc.EncryptMessage(message.Message)
		if err != nil {
			return fmt.Errorf("خطا در رمزنگاری محتوای پیام: %v", err)
		}
		message.Message = encryptedMessage
	}

	if err := msgRepo.Create(ctx, message); err != nil {
		uow.Rollback()
		return err
	}

	// به‌روزرسانی آخرین فعالیت جلسه
	_ = sessionRepo.UpdateFields(ctx, session.ID, map[string]interface{}{"updated_at": time.Now()})

	return uow.Commit()
}

// GetSessionMessages دریافت پیام‌های جلسه
func (cs *ChatService) GetSessionMessages(sessionID uint, limit, offset int) ([]models.OnlineChatMessage, error) {
	// ابتدا بررسی کنیم که جلسه وجود دارد
	var session models.ChatSession
	if err := cs.db.First(&session, sessionID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("session not found: %d", sessionID)
		}
		return nil, fmt.Errorf("خطا در بررسی وجود جلسه: %v", err)
	}

	var messages []models.OnlineChatMessage
	err := cs.db.Where("session_id = ?", sessionID).
		Preload("Operator").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error

	if err != nil {
		return nil, fmt.Errorf("خطا در دریافت پیام‌ها: %v", err)
	}

	// رمزگشایی پیام‌ها
	for i := range messages {
		if messages[i].Message != "" {
			decryptedMessage, err := cs.encryptionSvc.DecryptMessage(messages[i].Message)
			if err != nil {
				return nil, fmt.Errorf("خطا در رمزگشایی پیام: %v", err)
			}
			messages[i].Message = decryptedMessage
		}
	}

	return messages, nil
}

// MarkMessageAsRead علامت‌گذاری پیام به عنوان خوانده شده
func (cs *ChatService) MarkMessageAsRead(messageID uint) error {
	now := time.Now()
	return cs.db.Model(&models.OnlineChatMessage{}).
		Where("id = ?", messageID).
		Updates(map[string]interface{}{
			"is_read": true,
			"read_at": &now,
		}).Error
}

// MarkSessionMessagesAsRead علامت‌گذاری همه پیام‌های یک جلسه به عنوان خوانده شده
func (cs *ChatService) MarkSessionMessagesAsRead(sessionID uint) error {
	now := time.Now()
	return cs.db.Model(&models.OnlineChatMessage{}).
		Where("session_id = ? AND is_read = false", sessionID).
		Updates(map[string]interface{}{
			"is_read": true,
			"read_at": &now,
		}).Error
}

// EndChatSession پایان دادن به جلسه چت
func (cs *ChatService) EndChatSession(sessionID string) error {
	var session models.ChatSession
	if err := cs.db.Where("session_id = ?", sessionID).First(&session).Error; err != nil {
		return err
	}

	now := time.Now()
	updates := map[string]interface{}{
		"status":     "ended",
		"ended_at":   &now,
		"updated_at": now,
	}

	// اگر اپراتور دارد، تعداد چت‌های فعالش را کاهش دهید
	if session.OperatorID != nil {
		var operator models.ChatOperator
		if err := cs.db.Where("id = ?", *session.OperatorID).First(&operator).Error; err == nil {
			if operator.CurrentChats > 0 {
				if err := cs.db.Model(&operator).Update("current_chats", operator.CurrentChats-1).Error; err != nil {
					// لاگ کردن خطا در صورت عدم موفقیت در کاهش تعداد چت‌های فعال
					fmt.Printf("Failed to update operator current chats: %v\n", err)
				}
			}
		}
	}

	return cs.db.Model(&session).Updates(updates).Error
}

// GetAvailableOperators دریافت اپراتورهای در دسترس
func (cs *ChatService) GetAvailableOperators() ([]models.ChatOperator, error) {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	// استفاده از DB خام برای فیلتر ترکیبی ساده‌تر است
	var ops []models.ChatOperator
	err := uow.DB().WithContext(ctx).Where("is_active = ? AND status = ? AND current_chats < max_concurrent_chats", true, "online").Preload("User").Find(&ops).Error
	return ops, err
}

// UpdateOperatorStatus به‌روزرسانی وضعیت اپراتور
func (cs *ChatService) UpdateOperatorStatus(operatorID uint, status string) error {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	return uow.ChatOperatorRepository().UpdateFields(ctx, operatorID, map[string]interface{}{"status": status})
}

// GetChatStatistics دریافت آمار چت
func (cs *ChatService) GetChatStatistics(startDate, endDate time.Time, operatorID *uint) (*models.ChatStatistics, error) {
	query := cs.db.Model(&models.ChatSession{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate)

	if operatorID != nil {
		query = query.Where("operator_id = ?", *operatorID)
	}

	var totalSessions int64
	if err := query.Count(&totalSessions).Error; err != nil {
		return nil, fmt.Errorf("failed to count total sessions: %w", err)
	}

	var completedSessions int64
	if err := cs.db.Model(&models.ChatSession{}).
		Where("created_at BETWEEN ? AND ? AND status = ?", startDate, endDate, "ended").
		Count(&completedSessions).Error; err != nil {
		return nil, fmt.Errorf("failed to count completed sessions: %w", err)
	}

	var totalMessages int64
	if err := cs.db.Model(&models.OnlineChatMessage{}).
		Joins("JOIN chat_sessions ON chat_messages.session_id = chat_sessions.id").
		Where("chat_sessions.created_at BETWEEN ? AND ?", startDate, endDate).
		Count(&totalMessages).Error; err != nil {
		return nil, fmt.Errorf("failed to count total messages: %w", err)
	}

	// محاسبه میانگین زمان پاسخ
	var avgResponseTime float64
	if err := cs.db.Raw(`
		SELECT AVG(EXTRACT(EPOCH FROM (cm.created_at - cs.created_at))) as avg_response_time
		FROM chat_messages cm
		JOIN chat_sessions cs ON cm.session_id = cs.id
		WHERE cs.created_at BETWEEN ? AND ? AND cm.sender_type = 'operator'
	`, startDate, endDate).Scan(&avgResponseTime).Error; err != nil {
		// در صورت خطا، میانگین زمان پاسخ صفر تنظیم می‌شود
		avgResponseTime = 0
		fmt.Printf("Failed to calculate average response time: %v\n", err)
	}

	statistics := &models.ChatStatistics{
		Date:              startDate,
		TotalSessions:     int(totalSessions),
		CompletedSessions: int(completedSessions),
		TotalMessages:     int(totalMessages),
		AvgResponseTime:   avgResponseTime,
		OperatorID:        operatorID,
	}

	return statistics, nil
}

// CreateChatWidget ایجاد ابزارک چت با UnitOfWork
func (cs *ChatService) CreateChatWidget(widget *models.ChatWidget) error {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	if err := uow.ChatWidgetRepository().Create(ctx, widget); err != nil {
		uow.Rollback()
		return err
	}
	return uow.Commit()
}

// GetActiveChatWidgets دریافت ابزارک‌های فعال
func (cs *ChatService) GetActiveChatWidgets() ([]models.ChatWidget, error) {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	return uow.ChatWidgetRepository().GetActive(ctx)
}

// UpdateChatWidget بروزرسانی ابزارک چت
func (cs *ChatService) UpdateChatWidget(id uint, input *models.ChatWidget) (*models.ChatWidget, error) {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return nil, err
	}
	repo := uow.ChatWidgetRepository()
	item, err := repo.GetByID(ctx, id)
	if err != nil {
		uow.Rollback()
		return nil, err
	}
	// apply updates
	if input.Name != "" {
		item.Name = input.Name
	}
	if input.Description != "" {
		item.Description = input.Description
	}
	if input.WelcomeMessage != "" {
		item.WelcomeMessage = input.WelcomeMessage
	}
	if input.OfflineMessage != "" {
		item.OfflineMessage = input.OfflineMessage
	}
	if input.Theme != "" {
		item.Theme = input.Theme
	}
	if input.Position != "" {
		item.Position = input.Position
	}
	item.IsActive = input.IsActive

	if err := repo.Update(ctx, item); err != nil {
		uow.Rollback()
		return nil, err
	}
	if err := uow.Commit(); err != nil {
		return nil, err
	}
	return item, nil
}

// DeleteChatWidget حذف ابزارک چت
func (cs *ChatService) DeleteChatWidget(id uint) error {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	if err := uow.ChatWidgetRepository().Delete(ctx, id); err != nil {
		uow.Rollback()
		return err
	}
	return uow.Commit()
}

// CreateAIBot ایجاد بات هوش مصنوعی
func (cs *ChatService) CreateAIBot(bot *models.ChatAIBot) error {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	if err := uow.ChatAIBotRepository().Create(ctx, bot); err != nil {
		uow.Rollback()
		return err
	}
	return uow.Commit()
}

// GetActiveAIBots دریافت بات‌های فعال
func (cs *ChatService) GetActiveAIBots() ([]models.ChatAIBot, error) {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	return uow.ChatAIBotRepository().GetActive(ctx)
}

// UpdateAIBot بروزرسانی بات هوش مصنوعی
func (cs *ChatService) UpdateAIBot(id uint, input *models.ChatAIBot) (*models.ChatAIBot, error) {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return nil, err
	}
	repo := uow.ChatAIBotRepository()
	bot, err := repo.GetByID(ctx, id)
	if err != nil {
		uow.Rollback()
		return nil, err
	}
	if input.Name != "" {
		bot.Name = input.Name
	}
	if input.Description != "" {
		bot.Description = input.Description
	}
	if input.Model != "" {
		bot.Model = input.Model
	}
	if input.MaxTokens > 0 {
		bot.MaxTokens = input.MaxTokens
	}
	if input.Temperature > 0 {
		bot.Temperature = input.Temperature
	}
	if input.SystemPrompt != "" {
		bot.SystemPrompt = input.SystemPrompt
	}
	bot.IsActive = input.IsActive

	if err := repo.Update(ctx, bot); err != nil {
		uow.Rollback()
		return nil, err
	}
	if err := uow.Commit(); err != nil {
		return nil, err
	}
	return bot, nil
}

// DeleteAIBot حذف بات هوش مصنوعی
func (cs *ChatService) DeleteAIBot(id uint) error {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	if err := uow.ChatAIBotRepository().Delete(ctx, id); err != nil {
		uow.Rollback()
		return err
	}
	return uow.Commit()
}

// CreateKnowledgeBase ایجاد پایگاه دانش
func (cs *ChatService) CreateKnowledgeBase(kb *models.ChatKnowledgeBase) error {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	if err := uow.ChatKnowledgeBaseRepository().Create(ctx, kb); err != nil {
		uow.Rollback()
		return err
	}
	return uow.Commit()
}

// SearchKnowledgeBase جستجو در پایگاه دانش
func (cs *ChatService) SearchKnowledgeBase(query string) ([]models.ChatKnowledgeBase, error) {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	return uow.ChatKnowledgeBaseRepository().Search(ctx, query)
}

// ListKnowledgeBase لیست سوالات پایگاه دانش با فیلتر و صفحه‌بندی
func (cs *ChatService) ListKnowledgeBase(q string, limit, offset int) ([]models.ChatKnowledgeBase, int64, error) {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	return uow.ChatKnowledgeBaseRepository().List(ctx, q, limit, offset)
}

// GetKnowledgeBaseByID دریافت یک آیتم پایگاه دانش
func (cs *ChatService) GetKnowledgeBaseByID(id uint) (*models.ChatKnowledgeBase, error) {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	return uow.ChatKnowledgeBaseRepository().GetByID(ctx, id)
}

// UpdateKnowledgeBase بروزرسانی آیتم پایگاه دانش
func (cs *ChatService) UpdateKnowledgeBase(id uint, input *models.ChatKnowledgeBase) (*models.ChatKnowledgeBase, error) {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return nil, err
	}
	repo := uow.ChatKnowledgeBaseRepository()
	item, err := repo.GetByID(ctx, id)
	if err != nil {
		uow.Rollback()
		return nil, err
	}
	item.Title = input.Title
	item.Content = input.Content
	item.Category = input.Category
	item.Tags = input.Tags
	item.IsActive = input.IsActive
	if err := repo.Update(ctx, item); err != nil {
		uow.Rollback()
		return nil, err
	}
	if err := uow.Commit(); err != nil {
		return nil, err
	}
	return item, nil
}

// DeleteKnowledgeBase حذف آیتم پایگاه دانش
func (cs *ChatService) DeleteKnowledgeBase(id uint) error {
	ctx := context.Background()
	uow := cs.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	if err := uow.ChatKnowledgeBaseRepository().Delete(ctx, id); err != nil {
		uow.Rollback()
		return err
	}
	return uow.Commit()
}
