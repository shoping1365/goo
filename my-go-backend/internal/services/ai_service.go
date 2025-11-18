package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
	"net/http"
	"regexp"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AIService struct {
	APISettingsService *APISettingsService
	DB                 *gorm.DB
	sessions           map[string]*models.AISession // مدیریت جلسات در حافظه
}

func NewAIService(apiSettingsService *APISettingsService, db *gorm.DB) *AIService {
	return &AIService{
		APISettingsService: apiSettingsService,
		DB:                 db,
		sessions:           make(map[string]*models.AISession),
	}
}

// GenerateContent ارسال درخواست به OpenAI و دریافت پاسخ
func (s *AIService) GenerateContent(ctx context.Context, req *models.AIGenerateContentRequest, userId interface{}) (*models.AIGenerateContentResponse, error) {
	// بررسی و تنظیم مدل پیش‌فرض اگر مدل درخواستی موجود نباشد
	if req.Model == "" {
		req.Model = "gpt-4.1-nano-2025-04-14" // مدل پیش‌فرض جدید
	}

	fmt.Printf("🚀 شروع GenerateContent با مدل: %v\n", req.Model)
	config, err := s.APISettingsService.GetOpenAIConfig(ctx)
	if err != nil {
		fmt.Printf("❌ خطا در دریافت تنظیمات OpenAI: %v\n", err)
		errorResponse := utils.HandleDatabaseError("دریافت تنظیمات OpenAI", err)
		return nil, fmt.Errorf("%s: %w", errorResponse.UserMessage, err)
	}

	// لاگ کلید نهایی
	maskedKey := ""
	if len(config.APIKey) > 8 {
		maskedKey = config.APIKey[:4] + "..." + config.APIKey[len(config.APIKey)-4:]
	} else {
		maskedKey = config.APIKey
	}
	fmt.Printf("🔑 کلید نهایی ارسالی به OpenAI (GenerateContent): %s\n", maskedKey)

	if !config.Enabled || config.APIKey == "" {
		fmt.Printf("❌ OpenAI فعال نیست یا API Key تنظیم نشده است\n")
		errorResponse := utils.HandleBusinessError("تنظیمات OpenAI", "OpenAI فعال نیست یا API Key تنظیم نشده است")
		return nil, fmt.Errorf("%s", errorResponse.UserMessage)
	}

	openaiURL := config.APIUrl
	if openaiURL == "" {
		openaiURL = "https://api.openai.com/v1"
	}
	// اضافه کردن endpoint chat/completions
	if !strings.Contains(openaiURL, "/chat/completions") {
		openaiURL = strings.TrimRight(openaiURL, "/") + "/chat/completions"
	}
	fmt.Printf("🌐 ارسال درخواست به OpenAI: %s\n", openaiURL)

	openaiReq := map[string]interface{}{
		"model":       req.Model,
		"messages":    req.Messages,
		"max_tokens":  req.WordCount,
		"temperature": req.Temperature,
	}
	body, _ := json.Marshal(openaiReq)
	fmt.Printf("📦 بدنه ارسالی: %s\n", string(body))

	request, _ := http.NewRequestWithContext(ctx, "POST", openaiURL, bytes.NewBuffer(body))
	request.Header.Set("Authorization", "Bearer "+config.APIKey)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("❌ خطا در ارسال درخواست به OpenAI: %v\n", err)
		errorResponse := utils.HandleNetworkError("OpenAI", err)
		return nil, fmt.Errorf("%s", errorResponse.UserMessage)
	}
	defer resp.Body.Close()

	fmt.Printf("📥 کد وضعیت پاسخ OpenAI: %d\n", resp.StatusCode)
	if resp.StatusCode != 200 {
		respBody, _ := io.ReadAll(resp.Body)
		fmt.Printf("❌ پاسخ نامعتبر از OpenAI: %d | متن پاسخ: %s\n", resp.StatusCode, string(respBody))

		// تشخیص نوع خطا بر اساس کد وضعیت
		var errorResponse *utils.ErrorResponse
		switch resp.StatusCode {
		case 401:
			errorResponse = utils.HandleExternalAPIError("OpenAI", "/chat/completions", 401, string(respBody))
		case 403:
			errorResponse = utils.HandleExternalAPIError("OpenAI", "/chat/completions", 403, string(respBody))
		case 429:
			errorResponse = utils.HandleExternalAPIError("OpenAI", "/chat/completions", 429, string(respBody))
		case 500, 502, 503, 504:
			errorResponse = utils.HandleExternalAPIError("OpenAI", "/chat/completions", resp.StatusCode, string(respBody))
		default:
			errorResponse = utils.HandleExternalAPIError("OpenAI", "/chat/completions", resp.StatusCode, string(respBody))
		}

		return nil, fmt.Errorf("%s", errorResponse.UserMessage)
	}

	var openaiResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&openaiResp); err != nil {
		fmt.Printf("❌ خطا در خواندن پاسخ OpenAI: %v\n", err)
		errorResponse := utils.HandleExternalAPIError("OpenAI", "/chat/completions", 500, "خطا در پردازش پاسخ")
		return nil, fmt.Errorf("%s", errorResponse.UserMessage)
	}

	content := ""
	if len(openaiResp.Choices) > 0 {
		content = openaiResp.Choices[0].Message.Content
	}

	fmt.Printf("✅ پاسخ نهایی AI: %s\n", content)
	return &models.AIGenerateContentResponse{
		Content: content,
	}, nil
}

// HandleChat پردازش چت تعاملی
func (s *AIService) HandleChat(ctx context.Context, req *models.AIChatRequest, userId interface{}) (*models.AIChatResponse, error) {
	fmt.Printf("💬 شروع HandleChat با پیام: %s\n", req.Message)

	// دریافت تنظیمات OpenAI
	config, err := s.APISettingsService.GetOpenAIConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("خطا در دریافت تنظیمات OpenAI: %w", err)
	}

	if !config.Enabled || config.APIKey == "" {
		return nil, fmt.Errorf("OpenAI فعال نیست یا API Key تنظیم نشده است")
	}

	// مدیریت جلسه
	session := s.getOrCreateSession(req.SessionID)
	session.LastActivity = time.Now().Unix()

	// اضافه کردن پیام کاربر به تاریخچه
	session.ChatHistory = append(session.ChatHistory, models.ChatMessage{
		Role:    "user",
		Content: req.Message,
	})

	// ذخیره پیام کاربر در دیتابیس
	s.saveMessageToDatabase(session.SessionID, "user", req.Message, "user_message")

	// تشخیص نوع درخواست و پردازش
	messageType := s.detectMessageType(req.Message, session)

	var response *models.AIChatResponse
	switch messageType {
	case "title_request":
		response, err = s.handleTitleRequest(ctx, req, config, session)
	case "title_selection":
		response, err = s.handleTitleSelection(ctx, req, config, session)
	case "article_generation":
		response, err = s.handleArticleGeneration(ctx, req, config, session, userId)
	case "article_edit":
		response, err = s.handleArticleEdit(ctx, req, config, session)
	case "article_confirmation":
		response, err = s.handleArticleConfirmation(ctx, req, config, session, userId)
	case "suggestion_request":
		response, err = s.handleSuggestionRequest(ctx, req, config, session)
	case "confirmation":
		response, err = s.handleConfirmation(ctx, req, config, session)
	default:
		response, err = s.handleGeneralChat(ctx, req, config, session)
	}

	if err != nil {
		return nil, err
	}

	// اضافه کردن پاسخ AI به تاریخچه
	session.ChatHistory = append(session.ChatHistory, models.ChatMessage{
		Role:    "assistant",
		Content: response.Message,
	})

	// ذخیره پیام AI در دیتابیس
	s.saveMessageToDatabase(session.SessionID, "assistant", response.Message, response.Type)

	// به‌روزرسانی جلسه در دیتابیس
	s.saveSessionToDatabase(session, userId)

	return response, nil
}

// getOrCreateSession دریافت یا ایجاد جلسه جدید
func (s *AIService) getOrCreateSession(sessionID string) *models.AISession {
	if sessionID == "" {
		sessionID = fmt.Sprintf("session_%d", time.Now().Unix())
	}

	// ابتدا از حافظه بررسی کن
	if session, exists := s.sessions[sessionID]; exists {
		return session
	}

	// اگر در حافظه نبود، از دیتابیس بارگذاری کن
	var sessionDB models.AISessionDB
	if err := s.DB.Where("session_id = ?", sessionID).First(&sessionDB).Error; err == nil {
		// تبدیل از دیتابیس به حافظه
		session := &models.AISession{
			SessionID:     sessionDB.SessionID,
			CurrentState:  sessionDB.CurrentState,
			SelectedTitle: sessionDB.SelectedTitle,
			CreatedAt:     sessionDB.CreatedAt.Unix(),
			LastActivity:  sessionDB.UpdatedAt.Unix(),
		}

		// بارگذاری پیام‌ها از دیتابیس
		var messages []models.AIChatMessage
		if err := s.DB.Where("session_id = ?", sessionID).Order("created_at ASC").Find(&messages).Error; err == nil {
			for _, msg := range messages {
				session.ChatHistory = append(session.ChatHistory, models.ChatMessage{
					Role:    msg.Role,
					Content: msg.Content,
				})
			}
		}

		// بارگذاری DraftArticle از JSON
		if sessionDB.DraftArticle != nil {
			var article models.GeneratedArticle
			if err := json.Unmarshal(sessionDB.DraftArticle, &article); err == nil {
				session.DraftArticle = &article
			}
		}

		// بارگذاری EditHistory از JSON
		if sessionDB.EditHistory != nil {
			var editHistory []string
			if err := json.Unmarshal(sessionDB.EditHistory, &editHistory); err == nil {
				session.EditHistory = editHistory
			}
		}

		s.sessions[sessionID] = session
		return session
	}

	// ایجاد جلسه جدید
	session := &models.AISession{
		SessionID:    sessionID,
		CurrentState: "initial",
		CreatedAt:    time.Now().Unix(),
		LastActivity: time.Now().Unix(),
	}
	s.sessions[sessionID] = session

	// ذخیره در دیتابیس
	s.saveSessionToDatabase(session, nil)

	return session
}

// saveSessionToDatabase ذخیره جلسه در دیتابیس
func (s *AIService) saveSessionToDatabase(session *models.AISession, userId interface{}) error {
	var userID uint
	if userId != nil {
		if uid, ok := userId.(uint); ok {
			userID = uid
		}
	}

	// تبدیل DraftArticle به JSON
	var draftArticleJSON datatypes.JSON
	if session.DraftArticle != nil {
		draftBytes, _ := json.Marshal(session.DraftArticle)
		draftArticleJSON = datatypes.JSON(draftBytes)
	}

	// تبدیل EditHistory به JSON
	editHistoryBytes, _ := json.Marshal(session.EditHistory)
	editHistoryJSON := datatypes.JSON(editHistoryBytes)

	sessionDB := models.AISessionDB{
		SessionID:     session.SessionID,
		UserID:        userID,
		CurrentState:  session.CurrentState,
		SelectedTitle: session.SelectedTitle,
		DraftArticle:  draftArticleJSON,
		EditHistory:   editHistoryJSON,
	}

	// بررسی وجود جلسه
	var existingSession models.AISessionDB
	if err := s.DB.Where("session_id = ?", session.SessionID).First(&existingSession).Error; err != nil {
		// ایجاد جلسه جدید
		return s.DB.Create(&sessionDB).Error
	} else {
		// به‌روزرسانی جلسه موجود
		return s.DB.Model(&existingSession).Updates(sessionDB).Error
	}
}

// saveMessageToDatabase ذخیره پیام در دیتابیس
func (s *AIService) saveMessageToDatabase(sessionID string, role string, content string, messageType string) error {
	message := models.AIChatMessage{
		SessionID: sessionID,
		Role:      role,
		Content:   content,
		Type:      messageType,
	}

	return s.DB.Create(&message).Error
}

// detectMessageType تشخیص نوع پیام (به‌روزرسانی شده)
func (s *AIService) detectMessageType(message string, session *models.AISession) string {
	message = strings.ToLower(message)

	// درخواست عنوان
	titleKeywords := []string{
		"عنوان", "تیتر", "سرتیتر", "5 تا", "چند تا", "پیشنهاد عنوان",
	}
	for _, keyword := range titleKeywords {
		if strings.Contains(message, keyword) {
			return "title_request"
		}
	}

	// انتخاب عنوان (اگر در حالت انتخاب عنوان هستیم)
	if session.CurrentState == "title_selection" {
		// اگر عدد انتخاب کرده یا "این" یا "همین" گفته
		if regexp.MustCompile(`^\d+$`).MatchString(strings.TrimSpace(message)) ||
			strings.Contains(message, "این") || strings.Contains(message, "همین") {
			return "title_selection"
		}
	}

	// درخواست تولید مقاله
	articleKeywords := []string{
		"بنویس", "نویس", "تولید کن", "بساز", "ایجاد کن", "ساخت",
		"مقاله", "محتوا", "متن", "نوشته", "پست", "بلاگ",
	}
	for _, keyword := range articleKeywords {
		if strings.Contains(message, keyword) {
			return "article_generation"
		}
	}

	// ویرایش مقاله
	if session.CurrentState == "article_preview" || session.CurrentState == "editing" {
		editKeywords := []string{
			"ویرایش", "تغییر", "اصلاح", "این جا", "این قسمت", "این بخش",
		}
		for _, keyword := range editKeywords {
			if strings.Contains(message, keyword) {
				return "article_edit"
			}
		}
	}

	// تأیید نهایی مقاله
	if session.CurrentState == "article_preview" || session.CurrentState == "editing" {
		confirmationKeywords := []string{
			"همین خوبه", "تأیید", "عالی شده", "همین رو انجام بده", "باشه", "خب",
		}
		for _, keyword := range confirmationKeywords {
			if strings.Contains(message, keyword) {
				return "article_confirmation"
			}
		}
	}

	// درخواست پیشنهاد
	suggestionKeywords := []string{
		"چی", "پیشنهاد", "نوع", "کدام", "چه", "چطور",
	}
	for _, keyword := range suggestionKeywords {
		if strings.Contains(message, keyword) {
			return "suggestion_request"
		}
	}

	// تأیید عمومی
	confirmationKeywords := []string{
		"خب", "باشه", "تأیید", "بله", "درست", "اوکی", "ok",
	}
	for _, keyword := range confirmationKeywords {
		if strings.Contains(message, keyword) {
			return "confirmation"
		}
	}

	return "general"
}

// handleTitleRequest پردازش درخواست عنوان
func (s *AIService) handleTitleRequest(ctx context.Context, req *models.AIChatRequest, config *models.OpenAIConfig, session *models.AISession) (*models.AIChatResponse, error) {
	prompt := fmt.Sprintf(`
شما یک دستیار هوشمند برای تولید محتوای وب‌سایت هستید. بر اساس درخواست کاربر، 5 عنوان جذاب و SEO-friendly برای مقاله پیشنهاد دهید.

درخواست کاربر: %s

لطفاً 5 عنوان مختلف و جذاب پیشنهاد دهید. هر عنوان را در یک خط جداگانه بنویسید و با شماره مشخص کنید:

1. عنوان اول
2. عنوان دوم
3. عنوان سوم
4. عنوان چهارم
5. عنوان پنجم

عنوان‌ها باید جذاب، SEO-friendly و مرتبط با درخواست کاربر باشند.
`, req.Message)

	content, err := s.callOpenAI(ctx, prompt, config, session)
	if err != nil {
		return nil, err
	}

	// تجزیه عنوان‌ها
	lines := strings.Split(content, "\n")
	var titles []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && (strings.HasPrefix(line, "1.") || strings.HasPrefix(line, "2.") ||
			strings.HasPrefix(line, "3.") || strings.HasPrefix(line, "4.") || strings.HasPrefix(line, "5.")) {
			// حذف شماره از ابتدای خط
			title := strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(strings.TrimPrefix(
				strings.TrimPrefix(strings.TrimPrefix(line, "1."), "2."), "3."), "4."), "5."))
			if title != "" {
				titles = append(titles, title)
			}
		}
	}

	// اگر کمتر از 5 عنوان پیدا شد، از خطوط دیگر استفاده کن
	if len(titles) < 5 {
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" && !strings.HasPrefix(line, "1.") && !strings.HasPrefix(line, "2.") &&
				!strings.HasPrefix(line, "3.") && !strings.HasPrefix(line, "4.") && !strings.HasPrefix(line, "5.") {
				titles = append(titles, line)
				if len(titles) >= 5 {
					break
				}
			}
		}
	}

	session.CurrentState = "title_selection"
	session.EditHistory = append(session.EditHistory, fmt.Sprintf("درخواست عنوان: %s", req.Message))

	return &models.AIChatResponse{
		Message:              "بر اساس درخواست شما، این 5 عنوان را پیشنهاد می‌دهم:\n\n" + strings.Join(titles, "\n\n"),
		Type:                 "title_suggestions",
		Suggestions:          titles,
		RequiresConfirmation: true,
	}, nil
}

// handleTitleSelection پردازش انتخاب عنوان
func (s *AIService) handleTitleSelection(ctx context.Context, req *models.AIChatRequest, config *models.OpenAIConfig, session *models.AISession) (*models.AIChatResponse, error) {
	// تشخیص شماره انتخاب شده
	message := strings.TrimSpace(req.Message)
	var selectedIndex int

	if regexp.MustCompile(`^\d+$`).MatchString(message) {
		selectedIndex = int(message[0]-'0') - 1 // تبدیل "1" به 0, "2" به 1, ...
	} else {
		// اگر "این" یا "همین" گفته، اولین عنوان را انتخاب کن
		selectedIndex = 0
	}

	// پیدا کردن عنوان انتخاب شده از سابقه چت
	var selectedTitle string
	if len(session.ChatHistory) >= 2 {
		// آخرین پیام AI باید شامل عنوان‌ها باشد
		lastAIMessage := session.ChatHistory[len(session.ChatHistory)-2].Content

		// تجزیه عنوان‌ها از پیام AI
		lines := strings.Split(lastAIMessage, "\n")
		var titles []string
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" && (strings.HasPrefix(line, "1.") || strings.HasPrefix(line, "2.") ||
				strings.HasPrefix(line, "3.") || strings.HasPrefix(line, "4.") || strings.HasPrefix(line, "5.")) {
				// حذف شماره از ابتدای خط
				title := strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(strings.TrimPrefix(
					strings.TrimPrefix(strings.TrimPrefix(line, "1."), "2."), "3."), "4."), "5."))
				if title != "" {
					titles = append(titles, title)
				}
			}
		}

		// انتخاب عنوان بر اساس شماره
		if selectedIndex >= 0 && selectedIndex < len(titles) {
			selectedTitle = titles[selectedIndex]
		}
	}

	// اگر عنوان پیدا نشد، از پیام کاربر استفاده کن
	if selectedTitle == "" {
		selectedTitle = req.Message
	}

	session.CurrentState = "article_generation"
	session.SelectedTitle = selectedTitle
	session.EditHistory = append(session.EditHistory, fmt.Sprintf("انتخاب عنوان: %s", selectedTitle))

	return &models.AIChatResponse{
		Message:              fmt.Sprintf("عالی! عنوان انتخاب شده: «%s»\n\nحالا مقاله‌ای کامل برای این عنوان تولید می‌کنم. لطفاً کمی صبر کنید...", selectedTitle),
		Type:                 "title_selected",
		RequiresConfirmation: false,
	}, nil
}

// handleArticleGeneration تولید مقاله
func (s *AIService) handleArticleGeneration(ctx context.Context, req *models.AIChatRequest, config *models.OpenAIConfig, session *models.AISession, userId interface{}) (*models.AIChatResponse, error) {
	title := session.SelectedTitle
	if title == "" {
		title = req.Message // اگر عنوان انتخاب نشده، از پیام کاربر استفاده کن
	}

	prompt := fmt.Sprintf(`
شما یک دستیار هوشمند برای تولید محتوای وب‌سایت هستید. یک مقاله کامل برای عنوان زیر تولید کنید.

عنوان: %s

لطفاً مقاله‌ای کامل با این مشخصات تولید کنید:
- محتوای کامل و مفصل (حداقل 800 کلمه)
- خلاصه مقاله (excerpt) - حداکثر 200 کلمه
- توضیحات متا برای SEO - حداکثر 160 کاراکتر
- URL مناسب (slug) - فقط حروف انگلیسی، اعداد و خط تیره
- تگ‌های مرتبط (حداکثر 5 تگ)

پاسخ را به صورت JSON با این ساختار ارائه دهید:
{
  "title": "عنوان مقاله",
  "content": "محتوای کامل مقاله",
  "excerpt": "خلاصه مقاله",
  "meta_description": "توضیحات متا",
  "slug": "english-url-slug",
  "tags": "تگ1,تگ2,تگ3"
}
`, title)

	content, err := s.callOpenAI(ctx, prompt, config, session)
	if err != nil {
		return nil, err
	}

	var article models.GeneratedArticle
	if err := json.Unmarshal([]byte(content), &article); err != nil {
		return nil, fmt.Errorf("خطا در تجزیه پاسخ AI: %w", err)
	}

	// اطمینان از انگلیسی بودن slug
	article.Slug = s.generateEnglishSlug(article.Title)

	session.CurrentState = "article_preview"
	session.DraftArticle = &article
	session.EditHistory = append(session.EditHistory, fmt.Sprintf("تولید مقاله: %s", article.Title))

	return &models.AIChatResponse{
		Message: fmt.Sprintf(`مقاله تولید شد! لطفاً بررسی کنید:

**عنوان:** %s
**خلاصه:** %s
**تگ‌ها:** %s

**محتوای کامل:**
%s

اگر می‌خواهید تغییراتی اعمال کنم، بگویید "این قسمت را ویرایش کن" یا "این جا را تغییر بده".
اگر راضی هستید، بگویید "همین خوبه" یا "تأیید".`,
			article.Title, article.Excerpt, article.Tags, article.Content),
		Type:                 "article_preview",
		Article:              &article,
		RequiresConfirmation: true,
	}, nil
}

// handleArticleEdit ویرایش مقاله
func (s *AIService) handleArticleEdit(ctx context.Context, req *models.AIChatRequest, config *models.OpenAIConfig, session *models.AISession) (*models.AIChatResponse, error) {
	if session.DraftArticle == nil {
		return nil, fmt.Errorf("مقاله‌ای برای ویرایش موجود نیست")
	}

	// ساخت context از تاریخچه چت
	var contextInfo string
	if len(session.ChatHistory) > 0 {
		contextInfo = "**سابقه مکالمه:**\n"
		// فقط 5 پیام آخر برای context
		startIndex := 0
		if len(session.ChatHistory) > 10 {
			startIndex = len(session.ChatHistory) - 10
		}
		for i := startIndex; i < len(session.ChatHistory); i++ {
			msg := session.ChatHistory[i]
			contextInfo += fmt.Sprintf("%s: %s\n", msg.Role, msg.Content)
		}
	}

	prompt := fmt.Sprintf(`
شما یک دستیار هوشمند برای ویرایش محتوای وب‌سایت هستید. مقاله زیر را بر اساس درخواست کاربر ویرایش کنید.

%s

**مقاله فعلی:**
عنوان: %s
محتوا: %s
خلاصه: %s
توضیحات متا: %s
تگ‌ها: %s

**درخواست ویرایش کاربر:** %s

لطفاً مقاله را بر اساس درخواست کاربر ویرایش کنید. تغییرات را اعمال کنید اما ساختار کلی را حفظ کنید.
پاسخ را به صورت JSON با همان ساختار ارائه دهید:
{
  "title": "عنوان مقاله",
  "content": "محتوای کامل مقاله",
  "excerpt": "خلاصه مقاله",
  "meta_description": "توضیحات متا",
  "slug": "english-url-slug",
  "tags": "تگ1,تگ2,تگ3"
}
`, contextInfo, session.DraftArticle.Title, session.DraftArticle.Content, session.DraftArticle.Excerpt,
		session.DraftArticle.MetaDescription, session.DraftArticle.Tags, req.Message)

	content, err := s.callOpenAI(ctx, prompt, config, session)
	if err != nil {
		return nil, err
	}

	var article models.GeneratedArticle
	if err := json.Unmarshal([]byte(content), &article); err != nil {
		return nil, fmt.Errorf("خطا در تجزیه پاسخ AI: %w", err)
	}

	// اطمینان از انگلیسی بودن slug
	article.Slug = s.generateEnglishSlug(article.Title)

	session.CurrentState = "editing"
	session.DraftArticle = &article
	session.EditHistory = append(session.EditHistory, fmt.Sprintf("ویرایش مقاله: %s", req.Message))

	return &models.AIChatResponse{
		Message: fmt.Sprintf(`مقاله ویرایش شد! لطفاً بررسی کنید:

**عنوان:** %s
**خلاصه:** %s
**تگ‌ها:** %s

**محتوای کامل:**
%s

اگر می‌خواهید تغییرات بیشتری اعمال کنم، بگویید "این قسمت را ویرایش کن".
اگر راضی هستید، بگویید "همین خوبه" یا "تأیید".`,
			article.Title, article.Excerpt, article.Tags, article.Content),
		Type:                 "article_edit",
		Article:              &article,
		RequiresConfirmation: true,
	}, nil
}

// handleArticleConfirmation تأیید نهایی و ذخیره مقاله
func (s *AIService) handleArticleConfirmation(ctx context.Context, req *models.AIChatRequest, config *models.OpenAIConfig, session *models.AISession, userId interface{}) (*models.AIChatResponse, error) {
	if session.DraftArticle == nil {
		return nil, fmt.Errorf("مقاله‌ای برای ذخیره موجود نیست")
	}

	// ذخیره در دیتابیس
	postID, err := s.saveArticleToDatabase(session.DraftArticle, userId)
	if err != nil {
		return nil, err
	}

	session.CurrentState = "initial"
	session.DraftArticle = nil
	session.SelectedTitle = ""
	session.EditHistory = append(session.EditHistory, "مقاله تأیید و ذخیره شد")

	return &models.AIChatResponse{
		Message: fmt.Sprintf("🎉 مقاله با موفقیت ایجاد و به عنوان پیش‌نویس ذخیره شد!\n\n**شناسه مقاله:** %d\n\nمی‌توانید آن را در بخش مدیریت نوشته‌ها مشاهده و ویرایش کنید.", postID),
		Type:    "article_created",
		Article: session.DraftArticle,
		PostID:  &postID,
	}, nil
}

// generateEnglishSlug تولید slug انگلیسی
func (s *AIService) generateEnglishSlug(title string) string {
	// تبدیل به حروف کوچک
	slug := strings.ToLower(title)

	// حذف کاراکترهای غیرمجاز و جایگزینی با خط تیره
	slug = regexp.MustCompile(`[^a-z0-9\s-]`).ReplaceAllString(slug, "")

	// جایگزینی فاصله با خط تیره
	slug = regexp.MustCompile(`\s+`).ReplaceAllString(slug, "-")

	// حذف خط تیره‌های اضافی
	slug = regexp.MustCompile(`-+`).ReplaceAllString(slug, "-")

	// حذف خط تیره از ابتدا و انتها
	slug = strings.Trim(slug, "-")

	// اگر خالی شد، یک slug پیش‌فرض
	if slug == "" {
		slug = "article"
	}

	return slug
}

// handleSuggestionRequest پردازش درخواست پیشنهادات
func (s *AIService) handleSuggestionRequest(ctx context.Context, req *models.AIChatRequest, config *models.OpenAIConfig, session *models.AISession) (*models.AIChatResponse, error) {
	prompt := fmt.Sprintf(`
شما یک دستیار هوشمند برای تولید محتوای وب‌سایت هستید. بر اساس درخواست کاربر، پیشنهادات مفید ارائه دهید.

درخواست کاربر: %s

لطفاً پیشنهادات مفید و متنوع ارائه دهید. پاسخ را به صورت لیست ساده ارائه دهید.
`, req.Message)

	content, err := s.callOpenAI(ctx, prompt, config, session)
	if err != nil {
		return nil, err
	}

	// تجزیه پیشنهادات
	suggestions := strings.Split(content, "\n")
	var cleanSuggestions []string
	for _, s := range suggestions {
		s = strings.TrimSpace(s)
		if s != "" && !strings.HasPrefix(s, "-") && !strings.HasPrefix(s, "*") {
			cleanSuggestions = append(cleanSuggestions, s)
		}
	}

	return &models.AIChatResponse{
		Message:     "پیشنهادات من:",
		Type:        "suggestion",
		Suggestions: cleanSuggestions,
	}, nil
}

// handleConfirmation پردازش تأیید
func (s *AIService) handleConfirmation(ctx context.Context, req *models.AIChatRequest, config *models.OpenAIConfig, session *models.AISession) (*models.AIChatResponse, error) {
	prompt := fmt.Sprintf(`
شما یک دستیار هوشمند برای تولید محتوای وب‌سایت هستید. کاربر درخواست تأیید کرده است.

درخواست کاربر: %s

لطفاً تأیید کنید و آمادگی خود را برای تولید محتوا اعلام کنید.
`, req.Message)

	content, err := s.callOpenAI(ctx, prompt, config, session)
	if err != nil {
		return nil, err
	}

	return &models.AIChatResponse{
		Message: content,
		Type:    "confirmation",
	}, nil
}

// handleGeneralChat پردازش چت عمومی
func (s *AIService) handleGeneralChat(ctx context.Context, req *models.AIChatRequest, config *models.OpenAIConfig, session *models.AISession) (*models.AIChatResponse, error) {
	prompt := fmt.Sprintf(`
شما یک دستیار هوشمند برای تولید محتوای وب‌سایت هستید. به سوال کاربر پاسخ دهید.

درخواست کاربر: %s

لطفاً پاسخ مفید و مرتبط ارائه دهید.
`, req.Message)

	content, err := s.callOpenAI(ctx, prompt, config, session)
	if err != nil {
		return nil, err
	}

	return &models.AIChatResponse{
		Message: content,
		Type:    "general",
	}, nil
}

// callOpenAI فراخوانی OpenAI با تاریخچه چت
func (s *AIService) callOpenAI(ctx context.Context, prompt string, config *models.OpenAIConfig, session *models.AISession) (string, error) {
	openaiURL := config.APIUrl
	if openaiURL == "" {
		openaiURL = "https://api.openai.com/v1"
	}
	if !strings.Contains(openaiURL, "/chat/completions") {
		openaiURL = strings.TrimRight(openaiURL, "/") + "/chat/completions"
	}

	// ساخت پیام‌ها با تاریخچه چت
	messages := []interface{}{
		map[string]string{"role": "system", "content": "شما یک دستیار هوشمند فارسی برای تولید محتوای وب‌سایت هستید. همیشه context مکالمه را در نظر بگیرید و پاسخ‌های منسجم ارائه دهید."},
	}

	// اضافه کردن تاریخچه چت (حداکثر 10 پیام آخر)
	if session != nil && len(session.ChatHistory) > 0 {
		startIndex := 0
		if len(session.ChatHistory) > 10 {
			startIndex = len(session.ChatHistory) - 10
		}

		for i := startIndex; i < len(session.ChatHistory); i++ {
			msg := session.ChatHistory[i]
			messages = append(messages, map[string]string{
				"role":    msg.Role,
				"content": msg.Content,
			})
		}
	}

	// اضافه کردن پیام فعلی
	messages = append(messages, map[string]string{"role": "user", "content": prompt})

	openaiReq := map[string]interface{}{
		"model":       "gpt-4.1-nano-2025-04-14",
		"messages":    messages,
		"max_tokens":  2000,
		"temperature": 0.7,
	}

	body, _ := json.Marshal(openaiReq)
	request, _ := http.NewRequestWithContext(ctx, "POST", openaiURL, bytes.NewBuffer(body))
	request.Header.Set("Authorization", "Bearer "+config.APIKey)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("خطا در ارسال درخواست به OpenAI: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("خطا از OpenAI: %d - %s", resp.StatusCode, string(respBody))
	}

	var openaiResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&openaiResp); err != nil {
		return "", fmt.Errorf("خطا در خواندن پاسخ OpenAI: %w", err)
	}

	if len(openaiResp.Choices) == 0 {
		return "", fmt.Errorf("پاسخ خالی از OpenAI")
	}

	return openaiResp.Choices[0].Message.Content, nil
}

// saveArticleToDatabase ذخیره مقاله در جدول posts موجود
func (s *AIService) saveArticleToDatabase(article *models.GeneratedArticle, userId interface{}) (uint, error) {
	// تبدیل تگ‌ها به JSON
	var tagsJSON datatypes.JSON
	if article.Tags != "" {
		tags := strings.Split(article.Tags, ",")
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(tag)
		}
		tagsBytes, _ := json.Marshal(tags)
		tagsJSON = datatypes.JSON(tagsBytes)
	}

	// ایجاد رکورد جدید در جدول posts
	post := models.Post{
		Title:           article.Title,
		Content:         article.Content,
		Excerpt:         article.Excerpt,
		MetaDescription: article.MetaDescription,
		Slug:            article.Slug,
		Status:          "draft", // پیش‌نویس
		CategoryID:      article.CategoryID,
		Tags:            tagsJSON,
	}

	// تنظیم author_id اگر موجود باشد
	if userId != nil {
		if uid, ok := userId.(uint); ok {
			post.AuthorID = uid
		}
	}

	if err := s.DB.Create(&post).Error; err != nil {
		return 0, fmt.Errorf("خطا در ذخیره مقاله: %w", err)
	}

	return post.ID, nil
}

// clearSessionHistory پاک کردن تاریخچه جلسه
func (s *AIService) clearSessionHistory(sessionID string) {
	if session, exists := s.sessions[sessionID]; exists {
		session.ChatHistory = []models.ChatMessage{}
		session.EditHistory = []string{}
		session.CurrentState = "initial"
		session.SelectedTitle = ""
		session.DraftArticle = nil
		session.LastActivity = time.Now().Unix()
	}
}

// getSessionInfo دریافت اطلاعات جلسه
func (s *AIService) getSessionInfo(sessionID string) *models.AISession {
	if session, exists := s.sessions[sessionID]; exists {
		return session
	}
	return nil
}
