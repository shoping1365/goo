package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

// ImageSEOResult خروجی استاندارد تولید متادیتا
type ImageSEOResult struct {
	AltText     string `json:"alt_text"`
	Caption     string `json:"caption"`
	Description string `json:"description"`
	// Title اختیاری است؛ بر اساس تنظیمات تولید می‌شود
	Title string `json:"title,omitempty"`
}

// ImageSEOService سرویس مرکزی تولید متادیتای تصاویر با OpenAI
type ImageSEOService struct {
	DB                 *gorm.DB
	SettingService     *SettingService
	APISettingsService *APISettingsService
	HTTPClient         *http.Client
}

func NewImageSEOService(db *gorm.DB, ss *SettingService, api *APISettingsService) *ImageSEOService {
	return &ImageSEOService{
		DB:                 db,
		SettingService:     ss,
		APISettingsService: api,
		HTTPClient:         &http.Client{Timeout: 35 * time.Second},
	}
}

// BuildPrompt متن پرامپت را با توجه به تنظیمات می‌سازد
func (s *ImageSEOService) BuildPrompt(lang, pageTitle string, generateTitle bool) string {
	// در آینده می‌توان از image_seo.prompt_template خواند؛ الان پرامپت استاندارد
	titleLine := ""
	if generateTitle {
		titleLine = "\n- title: حداکثر 60 کاراکتر و واقعی؛ اگر ضروری نبود، خالی بگذار."
	}
	ctxLine := ""
	if strings.TrimSpace(pageTitle) != "" {
		ctxLine = fmt.Sprintf("\n- در نظر بگیر عنوان صفحه: %s", pageTitle)
	}
	return fmt.Sprintf("بساز خروجی JSON سخت‌گیرانه برای توصیف تصویر. زبان: %s.%s\n- alt_text: 80 تا 120 کاراکتر، دقیق و بدون کیورد استافینگ\n- caption: 1 تا 2 جمله بافت‌دار برای UI\n- description: 2 تا 4 جمله مفید برای صفحه%s\nفقط JSON برگردان.", lang, ctxLine, titleLine)
}

// CallOpenAI تماس با OpenAI (chat/completions) با ورودی تصویر base64
func (s *ImageSEOService) CallOpenAI(ctx context.Context, imgBase64, pageTitle, lang string, generateTitle bool) (*ImageSEOResult, error) {
	cfg, err := s.APISettingsService.GetOpenAIConfig(ctx)
	if err != nil {
		return nil, err
	}
	if !cfg.Enabled || cfg.APIKey == "" {
		return nil, fmt.Errorf("openai disabled or api key missing")
	}
	apiURL := cfg.APIUrl
	if apiURL == "" {
		apiURL = "https://api.openai.com/v1"
	}
	if !strings.Contains(apiURL, "/chat/completions") {
		apiURL = strings.TrimRight(apiURL, "/") + "/chat/completions"
	}

	model := cfg.DefaultModel
	if model == "" {
		// تنظیم اختصاصی این بخش
		if st, _ := s.SettingService.GetSetting(ctx, "image_seo.model"); st != nil && st.Value != "" {
			model = st.Value
		} else {
			model = "gpt-4o-mini"
		}
	}

	// پیام چندرسانه‌ای: متن + تصویر
	prompt := s.BuildPrompt(lang, pageTitle, generateTitle)
	reqBody := map[string]interface{}{
		"model": model,
		"messages": []interface{}{
			map[string]interface{}{"role": "system", "content": fmt.Sprintf("You are an SEO assistant. Answer in %s.", lang)},
			map[string]interface{}{"role": "user", "content": []interface{}{
				map[string]interface{}{"type": "text", "text": prompt},
				map[string]interface{}{"type": "image_url", "image_url": map[string]interface{}{"url": "data:image/jpeg;base64," + imgBase64}},
			}},
		},
		"temperature": 0.4,
		"max_tokens":  400,
	}

	b, _ := json.Marshal(reqBody)
	req, _ := http.NewRequestWithContext(ctx, "POST", apiURL, strings.NewReader(string(b)))
	req.Header.Set("Authorization", "Bearer "+cfg.APIKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("openai error %d: %s", resp.StatusCode, string(body))
	}
	var openaiResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&openaiResp); err != nil {
		return nil, err
	}
	if len(openaiResp.Choices) == 0 {
		return nil, fmt.Errorf("empty response")
	}
	content := strings.TrimSpace(openaiResp.Choices[0].Message.Content)
	var out ImageSEOResult
	if err := json.Unmarshal([]byte(content), &out); err != nil {
		// اگر JSON نبود، fallback حداقلی
		out = ImageSEOResult{AltText: "تصویر مربوط به صفحه", Caption: "", Description: ""}
	}
	return &out, nil
}

// ReadFileAsBase64 خواندن فایل از public و تبدیل به base64
func (s *ImageSEOService) ReadFileAsBase64(publicPath string) (string, error) {
	// publicPath شبیه "/uploads/media/..."
	pr := ""
	if wd, err := os.Getwd(); err == nil {
		pr = filepath.Dir(wd)
	}
	abs := filepath.Join(pr, "public", strings.TrimPrefix(publicPath, "/"))
	data, err := os.ReadFile(abs)
	if err != nil {
		return "", err
	}
	enc := base64.StdEncoding.EncodeToString(data)
	return enc, nil
}
