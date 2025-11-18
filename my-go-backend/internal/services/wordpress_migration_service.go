package services

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"unicode/utf8"

	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/utils"

	"github.com/chai2010/webp"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"gorm.io/gorm"
)

// WordPressMigrationService handles importing entities from WordPress.
// Queries (fetch from WP) and Commands (persist) are separated to follow CQRS.
type WordPressMigrationService struct {
	// cache برای ذخیره دسته‌بندی‌های وردپرس
	WcCategoriesCache map[int]wcCategory
	categoriesCacheMu sync.RWMutex
}

// fixEncoding متن را از encoding های مختلف به UTF-8 تبدیل می‌کند
func fixEncoding(text string) string {
	if text == "" {
		return text
	}

	// ابتدا کاراکترهای نامعتبر را حذف کن
	text = strings.ToValidUTF8(text, "")

	// اگر متن قبلاً UTF-8 معتبر است، همان را برگردان
	if utf8.ValidString(text) {
		return text
	}

	// تلاش برای تبدیل از Windows-1256 (encoding رایج برای فارسی)
	decoder := charmap.Windows1256.NewDecoder()
	if result, _, err := transform.String(decoder, text); err == nil && utf8.ValidString(result) {
		return result
	}

	// تلاش برای تبدیل از ISO-8859-1
	decoder = charmap.ISO8859_1.NewDecoder()
	if result, _, err := transform.String(decoder, text); err == nil && utf8.ValidString(result) {
		return result
	}

	// تلاش برای تبدیل از Windows-1252
	decoder = charmap.Windows1252.NewDecoder()
	if result, _, err := transform.String(decoder, text); err == nil && utf8.ValidString(result) {
		return result
	}

	// اگر هیچ کدام کار نکرد، کاراکترهای نامعتبر را حذف کن
	return strings.ToValidUTF8(text, "")
}

// safeStringForDB متن را برای ذخیره در دیتابیس آماده می‌کند
func safeStringForDB(text string) string {
	if text == "" {
		return text
	}

	// ابتدا کاراکترهای مشکل‌ساز را حذف کن
	text = strings.ReplaceAll(text, "\x00", "") // null bytes
	text = strings.ReplaceAll(text, "\x1a", "") // substitute character
	text = strings.ReplaceAll(text, "\xda", "") // مشکل‌ساز byte
	text = strings.ReplaceAll(text, "\xd8", "") // مشکل‌ساز byte

	// سپس fixEncoding را اعمال کن
	text = fixEncoding(text)

	// اگر هنوز مشکل دارد، کاراکترهای نامعتبر را حذف کن
	if !utf8.ValidString(text) {
		text = strings.ToValidUTF8(text, "")
	}

	return text
}

// safeDBOperation عملیات دیتابیس را با error handling انجام می‌دهد
func safeDBOperation(operation func() error, operationName string) error {
	if err := operation(); err != nil {
		// اگر خطای encoding است، آن را log کن اما ignore کن
		if strings.Contains(err.Error(), "invalid byte sequence for encoding") ||
			strings.Contains(err.Error(), "SQLSTATE 22021") ||
			strings.Contains(err.Error(), "current transaction is aborted") ||
			strings.Contains(err.Error(), "SQLSTATE 25P02") {
			log.Printf("[WARNING] خطای encoding/transaction در %s (ادامه می‌دهیم): %v", operationName, err)
			return nil // خطا را ignore کن
		}
		// برای سایر خطاها، خطا را برگردان
		return err
	}
	return nil
}

// safeDBOperationWithRecovery عملیات دیتابیس را با transaction recovery انجام می‌دهد
func safeDBOperationWithRecovery(uow unitofwork.UnitOfWork, operation func() error, operationName string) error {
	if err := operation(); err != nil {
		// اگر خطای encoding یا transaction abort است
		if strings.Contains(err.Error(), "invalid byte sequence for encoding") ||
			strings.Contains(err.Error(), "SQLSTATE 22021") ||
			strings.Contains(err.Error(), "current transaction is aborted") ||
			strings.Contains(err.Error(), "SQLSTATE 25P02") {
			log.Printf("[WARNING] خطای encoding/transaction در %s (تلاش برای recovery): %v", operationName, err)

			// تلاش برای rollback و شروع transaction جدید
			if rollbackErr := uow.Rollback(); rollbackErr != nil {
				log.Printf("[WARNING] خطا در rollback: %v", rollbackErr)
			}

			// شروع transaction جدید
			if beginErr := uow.BeginTx(context.Background()); beginErr != nil {
				log.Printf("[ERROR] خطا در شروع transaction جدید: %v", beginErr)
				return beginErr
			}

			log.Printf("[INFO] Transaction جدید شروع شد برای %s", operationName)
			return nil // خطا را ignore کن
		}
		// برای سایر خطاها، خطا را برگردان
		return err
	}
	return nil
}

type WPConfig struct {
	SiteURL string
	APIKey  string
}

type AdvancedOptions struct {
	Mode          string
	BatchSize     int
	IncludeImages bool
}

func NewWordPressMigrationService() *WordPressMigrationService {
	return &WordPressMigrationService{
		WcCategoriesCache: make(map[int]wcCategory),
	}
}

// loadWCCategoriesCache بارگذاری دسته‌بندی‌های وردپرس در cache
func (s *WordPressMigrationService) loadWCCategoriesCache(siteURL, ck, cs string) error {
	s.categoriesCacheMu.Lock()
	defer s.categoriesCacheMu.Unlock()

	// اگر cache قبلاً بارگذاری شده، نیازی به بارگذاری مجدد نیست
	if len(s.WcCategoriesCache) > 0 {
		return nil
	}

	categories, err := s.FetchAllWCCategories(siteURL, ck, cs)
	if err != nil {
		return fmt.Errorf("خطا در بارگذاری دسته‌بندی‌های وردپرس: %v", err)
	}

	// تبدیل به map برای دسترسی سریع
	for _, cat := range categories {
		s.WcCategoriesCache[cat.ID] = cat
	}

	log.Printf("[loadWCCategoriesCache] %d دسته‌بندی در cache بارگذاری شد", len(categories))
	return nil
}

// getAllWCCategories دریافت همه دسته‌بندی‌ها از cache
func (s *WordPressMigrationService) getAllWCCategories() []wcCategory {
	s.categoriesCacheMu.RLock()
	defer s.categoriesCacheMu.RUnlock()

	var categories []wcCategory
	for _, cat := range s.WcCategoriesCache {
		categories = append(categories, cat)
	}
	return categories
}

// getWCCategory دریافت دسته‌بندی از cache
func (s *WordPressMigrationService) getWCCategory(id int) (wcCategory, bool) {
	s.categoriesCacheMu.RLock()
	defer s.categoriesCacheMu.RUnlock()
	cat, exists := s.WcCategoriesCache[id]
	return cat, exists
}

// MetaExtractionOptions تنظیمات استخراج متا تگ‌ها
type MetaExtractionOptions struct {
	ExtractFromYoast     bool
	ExtractFromRankMath  bool
	ExtractFromOpenGraph bool
	ExtractFromMetaTags  bool
}

// MetaData ساختار داده‌های متا استخراج شده
type MetaData struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Keywords      string `json:"keywords"`
	OgTitle       string `json:"ogTitle"`
	OgDescription string `json:"ogDescription"`
	OgImage       string `json:"ogImage"`
	Canonical     string `json:"canonical"`
	Robots        string `json:"robots"`
	YoastTitle    string `json:"yoastTitle"`
	YoastDesc     string `json:"yoastDesc"`
	RankMathTitle string `json:"rankMathTitle"`
	RankMathDesc  string `json:"rankMathDesc"`
}

// ExtractMetaFromProductPage استخراج متا تگ‌ها از صفحه محصول وردپرس
func (s *WordPressMigrationService) ExtractMetaFromProductPage(productURL string, options MetaExtractionOptions) (*MetaData, error) {
	// اعتبارسنجی URL برای جلوگیری از SSRF
	validatedURL, err := validateURLForNetworkRequest(productURL)
	if err != nil {
		return nil, fmt.Errorf("URL نامعتبر: %v", err)
	}

	// دریافت محتوای صفحه
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", validatedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("خطا در ایجاد درخواست: %v", err)
	}

	// اضافه کردن User-Agent برای جلوگیری از بلاک شدن
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("خطا در دریافت صفحه: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("خطا در دریافت صفحه: کد وضعیت %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("خطا در خواندن محتوا: %v", err)
	}

	html := string(body)
	metaData := &MetaData{}

	// استخراج متا تگ‌های عمومی
	if options.ExtractFromMetaTags {
		metaData.Title = s.extractMetaTag(html, "title", "", "", "")
		metaData.Description = s.extractMetaTag(html, "meta", "name", "description", "content")
		metaData.Keywords = s.extractMetaTag(html, "meta", "name", "keywords", "content")
		metaData.Canonical = s.extractMetaTag(html, "link", "rel", "canonical", "href")
		metaData.Robots = s.extractMetaTag(html, "meta", "name", "robots", "content")
	}

	// استخراج Open Graph
	if options.ExtractFromOpenGraph {
		metaData.OgTitle = s.extractMetaTag(html, "meta", "property", "og:title", "content")
		metaData.OgDescription = s.extractMetaTag(html, "meta", "property", "og:description", "content")
		metaData.OgImage = s.extractMetaTag(html, "meta", "property", "og:image", "content")
	}

	// استخراج Yoast SEO
	if options.ExtractFromYoast {
		metaData.YoastTitle = s.extractMetaTag(html, "meta", "name", "yoast_wpseo_title", "content")
		metaData.YoastDesc = s.extractMetaTag(html, "meta", "name", "yoast_wpseo_metadesc", "content")
	}

	// استخراج RankMath
	if options.ExtractFromRankMath {
		metaData.RankMathTitle = s.extractMetaTag(html, "meta", "name", "rank_math_title", "content")
		metaData.RankMathDesc = s.extractMetaTag(html, "meta", "name", "rank_math_description", "content")
	}

	return metaData, nil
}

// extractMetaTag تابع کمکی برای استخراج متا تگ‌ها
func (s *WordPressMigrationService) extractMetaTag(html, tagType, attr1Type, attr1Value, attr2Type string) string {
	var pattern string

	if tagType == "title" {
		pattern = `<title[^>]*>([^<]*)</title>`
	} else if tagType == "meta" {
		if attr1Type == "name" {
			pattern = fmt.Sprintf(`<meta[^>]*name=["']%s["'][^>]*%s=["']([^"']*)["'][^>]*>`, attr1Value, attr2Type)
		} else if attr1Type == "property" {
			pattern = fmt.Sprintf(`<meta[^>]*property=["']%s["'][^>]*%s=["']([^"']*)["'][^>]*>`, attr1Value, attr2Type)
		}
	} else if tagType == "link" {
		pattern = fmt.Sprintf(`<link[^>]*rel=["']%s["'][^>]*%s=["']([^"']*)["'][^>]*>`, attr1Value, attr2Type)
	}

	if pattern == "" {
		return ""
	}

	re := regexp.MustCompile(`(?i)` + pattern)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}

	return ""
}

// ImportAttributesOnly: ویژگی‌هایی که در سیستم مقصد وجود دارند را چک می‌کند
// این تابع فقط چک می‌کند که ویژگی‌ها موجود هستند و ساختار را تغییر نمی‌دهد
func (s *WordPressMigrationService) ImportAttributesOnly(uow unitofwork.UnitOfWork, wpAttributes []map[string]interface{}) (imported, skipped int, err error) {
	ctx := context.Background()
	for _, a := range wpAttributes {
		code, _ := a["slug"].(string)
		if code == "" {
			code, _ = a["name"].(string)
		}
		if code == "" {
			skipped++
			continue
		}
		// یافتن ویژگی مقصد بر اساس code - فقط چک کردن وجود
		dest, derr := uow.AttributeRepository().GetByCode(ctx, code)
		if derr != nil || dest == nil {
			skipped++
			continue
		}
		// فقط چک کردن وجود - هیچ تغییری در ساختار ایجاد نمی‌کنیم
		imported++
	}
	return imported, skipped, nil
}

func normalizeSiteURL(siteURL string) string {
	s := strings.TrimSpace(siteURL)
	s = strings.TrimLeft(s, "/")
	if !strings.HasPrefix(s, "http://") && !strings.HasPrefix(s, "https://") {
		s = "https://" + s
	}
	return strings.TrimRight(s, "/")
}

// validateURLForNetworkRequest اعتبارسنجی URL برای جلوگیری از SSRF
// این تابع بررسی می‌کند که URL فقط http/https باشد و به IPهای داخلی یا localhost اشاره نکند
func validateURLForNetworkRequest(rawURL string) (string, error) {
	if rawURL == "" {
		return "", fmt.Errorf("URL خالی است")
	}

	// Parse کردن URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("URL نامعتبر: %v", err)
	}

	// بررسی پروتکل - فقط http و https مجاز است
	scheme := strings.ToLower(parsedURL.Scheme)
	if scheme != "http" && scheme != "https" {
		return "", fmt.Errorf("پروتکل نامعتبر: فقط http و https مجاز است")
	}

	// بررسی hostname
	host := parsedURL.Hostname()
	if host == "" {
		return "", fmt.Errorf("hostname خالی است")
	}

	// بررسی localhost و variants
	hostLower := strings.ToLower(host)
	if hostLower == "localhost" || hostLower == "127.0.0.1" || hostLower == "::1" || hostLower == "0.0.0.0" {
		return "", fmt.Errorf("دسترسی به localhost مجاز نیست")
	}

	// بررسی IP address - اگر host یک IP است، بررسی کن که private/internal نباشد
	if ip := net.ParseIP(host); ip != nil {
		// بررسی IPهای private/internal
		if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
			return "", fmt.Errorf("دسترسی به IP داخلی مجاز نیست")
		}
		// بررسی IPهای private (RFC 1918)
		if ip.IsPrivate() {
			return "", fmt.Errorf("دسترسی به IP خصوصی مجاز نیست")
		}
		// بررسی IPهای multicast
		if ip.IsMulticast() {
			return "", fmt.Errorf("دسترسی به IP multicast مجاز نیست")
		}
		// بررسی IPهای unspecified
		if ip.IsUnspecified() {
			return "", fmt.Errorf("دسترسی به IP نامشخص مجاز نیست")
		}
	} else {
		// اگر hostname است، بررسی کن که به localhost اشاره نکند
		if strings.HasSuffix(hostLower, ".localhost") || strings.HasSuffix(hostLower, ".local") {
			return "", fmt.Errorf("دسترسی به دامنه محلی مجاز نیست")
		}
	}

	// بازگرداندن URL معتبر
	return parsedURL.String(), nil
}

// -----------------------------
// WooCommerce Products (Import)
// -----------------------------

// wcCategory ساختار کامل دسته‌بندی وردپرس با اطلاعات والد
type wcCategory struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Parent int    `json:"parent"`
	Count  int    `json:"count"`
	Image  struct {
		ID  int    `json:"id"`
		Src string `json:"src"`
	} `json:"image"`
}

type wcProduct struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	Permalink        string `json:"permalink"`
	SKU              string `json:"sku"`
	Price            string `json:"price"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	StockQuantity    int    `json:"stock_quantity"`
	StockStatus      string `json:"stock_status"`
	Images           []struct {
		ID    int    `json:"id"`
		Src   string `json:"src"`
		Name  string `json:"name"`
		Alt   string `json:"alt"`
		Index int    `json:"position"`
	} `json:"images"`
	Categories []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"categories"`
	Attributes []struct {
		ID       int      `json:"id"`
		Name     string   `json:"name"`
		Slug     string   `json:"slug"`
		Position int      `json:"position"`
		Visible  bool     `json:"visible"`
		Options  []string `json:"options"`
	} `json:"attributes"`
	MetaData []struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	} `json:"meta_data"`
}

// FetchWCCategoriesPage دریافت دسته‌بندی‌های کامل از WooCommerce API
func (s *WordPressMigrationService) FetchWCCategoriesPage(siteURL, ck, cs string, page, perPage int) ([]wcCategory, error) {
	base := normalizeSiteURL(siteURL)
	log.Printf("[FetchWCCategoriesPage] شروع واکشی دسته‌بندی‌ها: siteURL=%s, page=%d, perPage=%d", siteURL, page, perPage)

	client := &http.Client{Timeout: 30 * time.Second}
	versions := []string{"v3", "v2", "v1"}

	for _, ver := range versions {
		apiURL := fmt.Sprintf("%s/wp-json/wc/%s/products/categories?per_page=%d&page=%d", base, ver, perPage, page)
		if ck != "" && cs != "" {
			apiURL = apiURL + fmt.Sprintf("&consumer_key=%s&consumer_secret=%s", ck, cs)
		}

		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(apiURL)
		if err != nil {
			log.Printf("[FetchWCCategoriesPage] URL نامعتبر: %v", err)
			continue
		}

		log.Printf("[FetchWCCategoriesPage] تلاش برای دریافت دسته‌بندی‌ها از: %s", validatedURL)

		req, err := http.NewRequest("GET", validatedURL, nil)
		if err != nil {
			log.Printf("[FetchWCCategoriesPage] خطا در ایجاد درخواست: %v", err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("[FetchWCCategoriesPage] خطا در ارسال درخواست: %v", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("[FetchWCCategoriesPage] خطا در خواندن پاسخ: %v", err)
				continue
			}

			var categories []wcCategory
			if err := json.Unmarshal(bodyBytes, &categories); err != nil {
				log.Printf("[FetchWCCategoriesPage] خطا در parse کردن JSON: %v", err)
				continue
			}

			log.Printf("[FetchWCCategoriesPage] موفقیت: %d دسته‌بندی دریافت شد", len(categories))
			return categories, nil
		}

		log.Printf("[FetchWCCategoriesPage] خطا در دریافت دسته‌بندی‌ها: HTTP %d", resp.StatusCode)
	}

	return nil, fmt.Errorf("خطا در دریافت دسته‌بندی‌ها از WooCommerce API")
}

// FetchAllWCCategories دریافت همه دسته‌بندی‌ها از WooCommerce
func (s *WordPressMigrationService) FetchAllWCCategories(siteURL, ck, cs string) ([]wcCategory, error) {
	var allCategories []wcCategory
	page := 1
	perPage := 100

	for {
		categories, err := s.FetchWCCategoriesPage(siteURL, ck, cs, page, perPage)
		if err != nil {
			return nil, err
		}
		if len(categories) == 0 {
			break
		}
		allCategories = append(allCategories, categories...)
		page++
	}

	return allCategories, nil
}

func (s *WordPressMigrationService) FetchWCProductsPage(siteURL, ck, cs string, page, perPage int) ([]wcProduct, error) {
	base := normalizeSiteURL(siteURL)
	log.Printf("[FetchWCProductsPage] شروع واکشی محصولات: siteURL=%s, page=%d, perPage=%d", siteURL, page, perPage)

	client := &http.Client{Timeout: 30 * time.Second}
	versions := []string{"v3", "v2", "v1"}

	for _, ver := range versions {
		apiURL := fmt.Sprintf("%s/wp-json/wc/%s/products?per_page=%d&page=%d&status=publish", base, ver, perPage, page)
		if ck != "" && cs != "" {
			apiURL = apiURL + fmt.Sprintf("&consumer_key=%s&consumer_secret=%s", ck, cs)
		}

		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(apiURL)
		if err != nil {
			log.Printf("[FetchWCProductsPage] URL نامعتبر: %v", err)
			continue
		}

		log.Printf("[FetchWCProductsPage] درخواست به: %s", validatedURL)

		req, _ := http.NewRequest("GET", validatedURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("[FetchWCProductsPage] خطا در درخواست HTTP: %v", err)
			continue
		}

		if resp == nil {
			log.Printf("[FetchWCProductsPage] پاسخ nil دریافت شد")
			continue
		}

		body, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()

		if readErr != nil {
			log.Printf("[FetchWCProductsPage] خطا در خواندن body: %v", readErr)
			continue
		}

		log.Printf("[FetchWCProductsPage] وضعیت HTTP: %d، طول body: %d", resp.StatusCode, len(body))

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			var arr []wcProduct
			if err := json.Unmarshal(body, &arr); err != nil {
				log.Printf("[FetchWCProductsPage] خطا در JSON unmarshal: %v", err)
				log.Printf("[FetchWCProductsPage] محتوای body: %s", string(body)[:min(500, len(body))])
				continue
			}
			log.Printf("[FetchWCProductsPage] موفقیت: %d محصول واکشی شد", len(arr))
			return arr, nil
		} else {
			log.Printf("[FetchWCProductsPage] وضعیت HTTP ناموفق: %d", resp.StatusCode)
			log.Printf("[FetchWCProductsPage] محتوای خطا: %s", string(body)[:min(200, len(body))])
		}
	}

	log.Printf("[FetchWCProductsPage] هیچ نسخه‌ای کار نکرد، آرایه خالی برمی‌گردانیم")
	return []wcProduct{}, nil
}

// FetchWCProductsPageWithMeta مشابه FetchWCProductsPage ولی همراه برگرداندن تعداد کل محصولات (X-WP-Total)
func (s *WordPressMigrationService) FetchWCProductsPageWithMeta(siteURL, ck, cs string, page, perPage int) ([]wcProduct, int, error) {
	base := normalizeSiteURL(siteURL)
	client := &http.Client{Timeout: 30 * time.Second}
	versions := []string{"v3", "v2", "v1"}
	for _, ver := range versions {
		url := fmt.Sprintf("%s/wp-json/wc/%s/products?per_page=%d&page=%d&status=publish", base, ver, perPage, page)
		if ck != "" && cs != "" {
			url = url + fmt.Sprintf("&consumer_key=%s&consumer_secret=%s", ck, cs)
		}
		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(url)
		if err != nil {
			continue
		}
		req, _ := http.NewRequest("GET", validatedURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")
		resp, err := client.Do(req)
		if err != nil || resp == nil {
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			var arr []wcProduct
			if err := json.Unmarshal(body, &arr); err != nil {
				continue
			}
			// گرفتن هدر تعداد کل
			totalStr := resp.Header.Get("X-WP-Total")
			total := 0
			if totalStr != "" {
				fmt.Sscanf(totalStr, "%d", &total)
			}
			return arr, total, nil
		}
	}
	return []wcProduct{}, 0, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func slugifyTitle(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	var out []rune
	prevDash := false
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			out = append(out, r)
			prevDash = false
		} else if !prevDash {
			out = append(out, '-')
			prevDash = true
		}
	}
	res := strings.Trim(string(out), "-")
	if res == "" {
		// fallback to random identifier
		return "item-" + uuid.NewString()[:8]
	}

	// محدود کردن طول slug به 100 کاراکتر (مطابق با محدودیت دیتابیس)
	if len(res) > 100 {
		// بریدن از انتها و اضافه کردن شناسه منحصر به فرد
		truncated := res[:90] // 90 کاراکتر برای slug + 10 کاراکتر برای شناسه
		uniqueId := uuid.NewString()[:8]
		res = truncated + "-" + uniqueId
	}

	return res
}

// createCategoriesInOrder ایجاد همه دسته‌بندی‌ها به عنوان دسته‌بندی اصلی
func (s *WordPressMigrationService) createCategoriesInOrder(uow unitofwork.UnitOfWork, wcCategories []wcCategory, allowCreate bool) error {
	// ایجاد map برای دسترسی سریع
	wcCategoriesMap := make(map[int]wcCategory)
	for _, cat := range wcCategories {
		wcCategoriesMap[cat.ID] = cat
	}

	// همه دسته‌بندی‌ها را به عنوان دسته‌بندی اصلی ایجاد کن
	for _, wcCat := range wcCategories {
		_, err := ensureCategoryAsMain(uow, wcCat, wcCategoriesMap, allowCreate)
		if err != nil {
			log.Printf("[WARNING] خطا در ایجاد دسته‌بندی %s: %v", wcCat.Name, err)
		}
	}

	return nil
}

// ensureCategoryAsMain ایجاد دسته‌بندی با حفظ سلسله مراتب
func ensureCategoryAsMain(uow unitofwork.UnitOfWork, wcCat wcCategory, wcCategoriesMap map[int]wcCategory, allowCreate bool) (*models.Category, error) {
	ctx := context.Background()

	// بررسی وجود دسته‌بندی در دیتابیس
	if existing, err := uow.CategoryRepository().GetBySlug(ctx, wcCat.Slug); err == nil && existing != nil {
		return existing, nil
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// اگر رکورد یافت نشد و اجازه ایجاد نداریم
	if !allowCreate {
		return nil, nil
	}

	// ایجاد دسته‌بندی جدید
	c := &models.Category{
		Name:      wcCat.Name,
		Slug:      wcCat.Slug,
		Published: true,
	}

	// همه دسته‌بندی‌ها به عنوان "اصلی" ایجاد می‌شوند
	// سلسله مراتب فقط در parent_id و parent_slug حفظ می‌شود
	if wcCat.Parent > 0 {
		if parentWC, exists := wcCategoriesMap[wcCat.Parent]; exists {
			// ابتدا والد را ایجاد کن
			parent, err := ensureCategoryAsMain(uow, parentWC, wcCategoriesMap, allowCreate)
			if err != nil {
				log.Printf("[WARNING] خطا در ایجاد دسته‌بندی والد %s: %v", parentWC.Name, err)
			} else if parent != nil {
				c.ParentID = &parent.ID
				c.ParentSlug = &parentWC.Slug // تعیین parent_slug
				log.Printf("[ensureCategoryAsMain] دسته‌بندی با والد ایجاد شد: %s (ParentID: %d, ParentSlug: %s)", c.Name, parent.ID, parentWC.Slug)
			}
		}
	} else {
		// دسته‌بندی بدون والد
		c.ParentID = nil
		c.ParentSlug = nil
		log.Printf("[ensureCategoryAsMain] دسته‌بندی بدون والد ایجاد شد: %s (ParentID: %v)", c.Name, c.ParentID)
	}

	// ایجاد دسته‌بندی در دیتابیس
	if err := uow.CategoryRepository().Create(ctx, c); err != nil {
		return nil, err
	}

	return c, nil
}

// ensureCategory ایجاد دسته‌بندی ساده (برای سازگاری با کد موجود)
func ensureCategory(uow unitofwork.UnitOfWork, name, slug string, allowCreate bool) (*models.Category, error) {
	ctx := context.Background()
	if slug == "" {
		slug = slugifyTitle(name)
	}
	if existing, err := uow.CategoryRepository().GetBySlug(ctx, slug); err == nil && existing != nil {
		return existing, nil
	} else if err != nil && err != gorm.ErrRecordNotFound {
		// اگر خطای دیگری غیر از "رکورد یافت نشد" رخ داد
		return nil, err
	}
	// اگر رکورد یافت نشد (err == gorm.ErrRecordNotFound) یا existing == nil
	if !allowCreate {
		return nil, nil
	}
	c := &models.Category{Name: name, Slug: slug, Published: true}
	if err := uow.CategoryRepository().Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

// extractBrandFromProductName استخراج برند از نام محصول
func extractBrandFromProductName(productName string) string {
	// لیست برندهای شناخته شده
	knownBrands := []string{
		"اکووکس", "Ecovacs",
		"سامسونگ", "Samsung",
		"ال جی", "LG",
		"سونی", "Sony",
		"اپل", "Apple",
		"هواوی", "Huawei",
		"شیائومی", "Xiaomi",
		"وان پلاس", "OnePlus",
		"نوکیا", "Nokia",
		"موتورولا", "Motorola",
		"ایسوس", "ASUS",
		"ایسر", "Acer",
		"لنوو", "Lenovo",
		"دل", "Dell",
		"اچ پی", "HP",
		"اینتل", "Intel",
		"ای ام دی", "AMD",
		"انویدیا", "NVIDIA",
		"کینگستون", "Kingston",
		"کورسیر", "Corsair",
		"رازر", "Razer",
		"لوژیتک", "Logitech",
		"مایکروسافت", "Microsoft",
		"گوگل", "Google",
		"آمازون", "Amazon",
		"فیلیپس", "Philips",
		"پاناسونیک", "Panasonic",
		"شارپ", "Sharp",
		"توشیبا", "Toshiba",
		"کنن", "Canon",
		"نایکون", "Nikon",
		"سونی", "Sony",
		"جی بی ال", "JBL",
		"بوز", "Bose",
		"سونی", "Sony",
		"سامسونگ", "Samsung",
	}

	// جستجو در نام محصول
	for _, brand := range knownBrands {
		if strings.Contains(strings.ToLower(productName), strings.ToLower(brand)) {
			return brand
		}
	}

	return ""
}

func ensureBrand(uow unitofwork.UnitOfWork, name, slug string, allowCreate bool) (*models.Brand, error) {
	ctx := context.Background()
	if slug == "" {
		slug = slugifyTitle(name)
	}
	// BrandRepository has GetBySlug
	if repo, ok := interface{}(uow).(interface {
		BrandRepository() repository.BrandRepositoryInterface
	}); ok {
		if b, err := repo.BrandRepository().GetBySlug(ctx, slug); err == nil && b != nil {
			return b, nil
		} else if err != nil && err != gorm.ErrRecordNotFound {
			// اگر خطای دیگری غیر از "رکورد یافت نشد" رخ داد
			return nil, err
		}
		// اگر رکورد یافت نشد (err == gorm.ErrRecordNotFound) یا b == nil
		if !allowCreate {
			return nil, nil
		}
		nb := &models.Brand{Name: name, Slug: slug, Published: true}
		if err := repo.BrandRepository().Create(ctx, nb); err != nil {
			return nil, err
		}
		return nb, nil
	}
	return nil, nil
}

// UpsertProductFromWC: درج/به‌روزرسانی محصول با حفظ ترتیب تصاویر و ویژگی‌ها
func (s *WordPressMigrationService) UpsertProductFromWC(
	uow unitofwork.UnitOfWork,
	p wcProduct,
	allowCreateCategories bool,
	allowCreateBrands bool,
	_ bool,
	redirectProducts bool,
	redirectCode int,
	draftMode bool,
	includeSlugInURL bool,
	transferSEOfromYoast bool,
	transferSEOfromRankMath bool,
	transferPrices bool,
	transferStock bool,
	_ bool,
	transferDescriptions bool,
	transferShortDescription bool,
	transferImages bool,
	transferAttributes bool,
	redirectGroupName string,
	siteURL string,
	consumerKey string,
	consumerSecret string,
) (created bool, err error) {
	ctx := context.Background()

	// ایجاد یک savepoint برای هر محصول تا در صورت خطا در هر بخش بتوان به آن بازگشت
	txName := fmt.Sprintf("sp_product_%d", p.ID)
	savepointErr := uow.DB().SavePoint(txName).Error
	if savepointErr != nil {
		log.Printf("[ERROR] خطا در ایجاد savepoint برای محصول %d: %v", p.ID, savepointErr)
		return false, savepointErr
	}

	defer func() {
		// در صورت وجود خطا، به savepoint برگرد
		if err != nil {
			rollbackErr := uow.DB().RollbackTo(txName).Error
			if rollbackErr != nil {
				log.Printf("[ERROR] خطا در rollback به savepoint %s: %v", txName, rollbackErr)
			}
		}
	}()

	// برند را از attributes با نام/slug brand استخراج می‌کنیم
	var brandName, brandSlug string
	for _, a := range p.Attributes {
		ls := strings.ToLower(a.Slug)
		ln := strings.ToLower(a.Name)
		if ls == "brand" || ln == "brand" || ln == "برند" {
			if len(a.Options) > 0 {
				brandName = a.Options[0]
				brandSlug = slugifyTitle(brandName)
			}
			break
		}
	}

	// اگر برند از attributes پیدا نشد، از نام محصول استخراج کن
	if brandName == "" {
		brandName = extractBrandFromProductName(p.Name)
		if brandName != "" {
			brandSlug = slugifyTitle(brandName)
		}
	}

	// بارگذاری دسته‌بندی‌های وردپرس در cache اگر نیاز باشد
	if allowCreateCategories && len(p.Categories) > 0 {
		if err := s.loadWCCategoriesCache(siteURL, consumerKey, consumerSecret); err != nil {
			log.Printf("[WARNING] خطا در بارگذاری دسته‌بندی‌های وردپرس: %v", err)
		}

		// ایجاد دسته‌بندی‌ها به ترتیب صحیح (ابتدا والدها، سپس فرزندها)
		if err := s.createCategoriesInOrder(uow, s.getAllWCCategories(), allowCreateCategories); err != nil {
			log.Printf("[WARNING] خطا در ایجاد دسته‌بندی‌ها: %v", err)
		}
	}

	// دسته را اولین دسته از WC در نظر می‌گیریم
	var cat *models.Category
	if len(p.Categories) > 0 {
		first := p.Categories[0]

		// اگر دسته‌بندی کامل در cache موجود است، از آن استفاده کن
		if wcCat, exists := s.getWCCategory(first.ID); exists {
			c, err := ensureCategoryAsMain(uow, wcCat, s.WcCategoriesCache, allowCreateCategories)
			if err != nil {
				log.Printf("[WARNING] خطا در ایجاد دسته‌بندی %s: %v", wcCat.Name, err)
				// fallback به روش قدیمی
				c, _ = ensureCategory(uow, first.Name, first.Slug, allowCreateCategories)
			}
			cat = c
		} else {
			// fallback به روش قدیمی اگر دسته‌بندی کامل در cache نباشد
			c, _ := ensureCategory(uow, first.Name, first.Slug, allowCreateCategories)
			cat = c
		}
	}

	// یافتن یا ساخت برند
	var brandID *uint
	if brandName != "" {
		if b, _ := ensureBrand(uow, brandName, brandSlug, allowCreateBrands); b != nil {
			brandID = &b.ID
		}
	}

	// یافتن محصول مقصد بر اساس SKU یا Slug
	prodRepo := uow.ProductRepository()
	var dest *models.Product
	if p.SKU != "" {
		if found, err := prodRepo.GetBySKU(ctx, p.SKU); err == nil && found != nil {
			dest = found
		}
	}
	if dest == nil && p.Slug != "" {
		if found, err := prodRepo.GetBySlug(ctx, p.Slug); err == nil && found != nil {
			dest = found
		}
	}

	// تولید slug مقصد از نام
	slug := strings.TrimSpace(p.Slug)
	if slug == "" {
		slug = slugifyTitle(p.Name)
	}

	// تعیین slug نهایی برای محصول (از target path ریدایرکت)
	var finalProductSlug string
	if includeSlugInURL {
		finalProductSlug = slug // اگر slug در URL باشد، از همان استفاده کن
	} else {
		finalProductSlug = "" // اگر slug در URL نباشد، خالی بگذار
	}

	log.Printf("[INFO] محصول '%s' - Slug نهایی: '%s' (includeSlugInURL: %v)", p.Name, finalProductSlug, includeSlugInURL)
	// ensure uniqueness in DB
	if existing, err2 := prodRepo.GetBySlug(ctx, slug); err2 == nil && existing != nil {
		slug = slug + "-" + uuid.NewString()[:6]
	}

	price := 0.0
	if transferPrices {
		if p.Price != "" {
			fmt.Sscanf(p.Price, "%f", &price)
		}
	}

	// موجودی
	stockQty := 0
	stockStatus := ""
	callForPrice := false
	if transferStock {
		stockQty = p.StockQuantity
		s := strings.ToLower(strings.TrimSpace(p.StockStatus))
		switch s {
		case "instock", "in-stock", "in_stock":
			stockStatus = "in_stock"
		case "outofstock", "out-of-stock", "out_of_stock":
			stockStatus = "out_of_stock"
		case "onbackorder", "backorder":
			stockStatus = "backorder"
		default:
			if stockQty > 0 {
				stockStatus = "in_stock"
			} else {
				stockStatus = "out_of_stock"
			}
		}
	}

	// بررسی وضعیت "برای استعلام قیمت تماس بگیرید" از meta_data
	for _, md := range p.MetaData {
		k := strings.ToLower(md.Key)
		if k == "_call_for_price" || k == "call_for_price" || k == "_price_request" {
			if v, ok := md.Value.(string); ok {
				callForPrice = v == "yes" || v == "1" || v == "true"
			} else if v, ok := md.Value.(bool); ok {
				callForPrice = v
			}
			break
		}
	}

	// اگر قیمت خالی باشد، call_for_price را فعال کن
	if transferPrices && price == 0.0 && p.Price == "" {
		callForPrice = true
	}

	if dest == nil {
		// اگر SKU خالی باشد، از slug استفاده کن
		productSKU := p.SKU
		if productSKU == "" {
			productSKU = slug
		}

		np := models.Product{
			UUID: uuid.NewString(),
			Name: p.Name, Slug: finalProductSlug, SKU: productSKU, Price: price, Status: "active", IsActive: true,
		}

		// انتقال توضیحات با پردازش HTML
		if transferDescriptions && p.Description != "" {
			np.FullDescription = p.Description
			// پردازش HTML توضیحات کامل
			processedHTML, err := s.ProcessHTML(ctx, p.Description, normalizeSiteURL(p.Permalink), slug, true)
			if err != nil {
				log.Printf("خطا در پردازش توضیحات محصول %s: %v", p.Name, err)
				// در صورت خطا، از متن اصلی استفاده می‌کنیم
				np.FullDescription = p.Description
			} else {
				np.FullDescription = processedHTML
			}
		}
		if transferShortDescription && p.ShortDescription != "" {
			np.Description = p.ShortDescription
			// پردازش HTML توضیحات کوتاه
			processedShortHTML, err := s.ProcessHTML(ctx, p.ShortDescription, normalizeSiteURL(p.Permalink), slug, false)
			if err != nil {
				log.Printf("خطا در پردازش توضیحات کوتاه محصول %s: %v", p.Name, err)
				// در صورت خطا، از متن اصلی استفاده می‌کنیم
				np.Description = p.ShortDescription
			} else {
				np.Description = processedShortHTML
			}
		}

		np.CallForPrice = callForPrice

		if cat != nil {
			cid := uint(cat.ID)
			np.CategoryID = &cid
		}
		if brandID != nil {
			np.BrandID = brandID
		}
		if transferStock {
			np.StockQuantity = stockQty
			np.StockStatus = stockStatus
			np.TrackInventory = true
			np.ShowStockToCustomer = true
		}

		// انتقال فیلدهای سئو از Yoast/RankMath در صورت انتخاب
		if len(p.MetaData) > 0 {
			var seoTitle, seoDesc string
			if transferSEOfromYoast {
				for _, md := range p.MetaData {
					lk := strings.ToLower(md.Key)
					if lk == "_yoast_wpseo_title" {
						if v, ok := md.Value.(string); ok {
							seoTitle = safeStringForDB(v)
						}
					} else if lk == "_yoast_wpseo_metadesc" {
						if v, ok := md.Value.(string); ok {
							seoDesc = safeStringForDB(v)
						}
					}
				}
			}
			if transferSEOfromRankMath {
				for _, md := range p.MetaData {
					lk := strings.ToLower(md.Key)
					if lk == "rank_math_title" {
						if v, ok := md.Value.(string); ok {
							seoTitle = safeStringForDB(v)
						}
					} else if lk == "rank_math_description" {
						if v, ok := md.Value.(string); ok {
							seoDesc = safeStringForDB(v)
						}
					}
				}
			}
			if seoTitle != "" {
				np.SeoTitle = seoTitle
			}
			if seoDesc != "" {
				np.MetaDescription = seoDesc
			}
		}

		if err := prodRepo.Create(ctx, &np); err != nil {
			log.Printf("[ERROR] خطا در ایجاد محصول '%s': %v", p.Name, err)
			return false, fmt.Errorf("خطا در ایجاد محصول '%s': %w", p.Name, err)
		}
		dest = &np
		created = true
	} else {
		updates := map[string]interface{}{ // update existing product
			"name":      p.Name,
			"slug":      finalProductSlug, // ensure slug is saved from target path
			"status":    "active",
			"is_active": true,
		}

		// انتقال توضیحات با پردازش HTML
		if transferDescriptions && p.Description != "" {
			updates["full_description"] = p.Description
			// پردازش HTML توضیحات کامل
			processedHTML, err := s.ProcessHTML(ctx, p.Description, normalizeSiteURL(p.Permalink), slug, true)
			if err != nil {
				log.Printf("خطا در پردازش توضیحات محصول %s: %v", p.Name, err)
				// در صورت خطا، از متن اصلی استفاده می‌کنیم
				updates["full_description"] = p.Description
			} else {
				updates["full_description"] = processedHTML
			}
		}
		if transferShortDescription && p.ShortDescription != "" {
			updates["description"] = p.ShortDescription
			// پردازش HTML توضیحات کوتاه
			processedShortHTML, err := s.ProcessHTML(ctx, p.ShortDescription, normalizeSiteURL(p.Permalink), slug, false)
			if err != nil {
				log.Printf("خطا در پردازش توضیحات کوتاه محصول %s: %v", p.Name, err)
				// در صورت خطا، از متن اصلی استفاده می‌کنیم
				updates["description"] = p.ShortDescription
			} else {
				updates["description"] = processedShortHTML
			}
		}

		updates["call_for_price"] = callForPrice

		if transferPrices {
			updates["price"] = price
		}
		if cat != nil {
			updates["category_id"] = cat.ID
		}
		// Draft mode handling
		if draftMode {
			updates["status"] = "draft"
			updates["is_active"] = false
		}
		if brandID != nil {
			updates["brand_id"] = *brandID
		}
		if transferStock {
			updates["stock_quantity"] = stockQty
			updates["stock_status"] = stockStatus
			updates["track_inventory"] = true
			updates["show_stock_to_customer"] = true
		}
		// به‌روزرسانی فیلدها در محصول موجود
		if err := uow.DB().WithContext(ctx).Model(&models.Product{}).Where("id = ?", dest.ID).Updates(updates).Error; err != nil {
			return false, err
		}
	}

	// تصاویر: پاک‌سازی قبلی و درج بر اساس position
	if transferImages {
		if err := uow.DB().WithContext(ctx).Where("product_id = ?", dest.ID).Delete(&models.ProductImage{}).Error; err != nil {
			return false, err
		}
	}

	// thumbnail = کمترین position
	thumb := ""
	if transferImages && len(p.Images) > 0 {
		// sort by Index asc
		for i := 0; i < len(p.Images)-1; i++ {
			for j := i + 1; j < len(p.Images); j++ {
				if p.Images[j].Index < p.Images[i].Index {
					p.Images[i], p.Images[j] = p.Images[j], p.Images[i]
				}
			}
		}
		// مسیر ذخیره‌سازی: /public/uploads/media/products/images/YYYY/MM/ یا /YYYY/MM/batch_X/
		// برای هر تصویر جداگانه مسیر مناسب را تعیین می‌کنیم
		for idx, im := range p.Images {
			// تعیین مسیر مناسب برای این تصویر خاص
			prodDir, err := getProductImageDir()
			if err != nil {
				log.Printf("خطا در تعیین مسیر ذخیره عکس: %v", err)
				// تلاش برای ایجاد مسیر سال/ماه به صورت دستی
				now := time.Now()
				year := now.Format("2006")
				month := now.Format("01")
				fallbackDir := filepath.Join(productImagesBaseDir(), year, month)

				if ensureDirErr := ensureDir(fallbackDir); ensureDirErr != nil {
					log.Printf("خطا در ایجاد پوشه سال/ماه: %v", ensureDirErr)
					// آخرین fallback: مسیر پایه
					prodDir = productImagesBaseDir()
					if ensureDirErr := ensureDir(prodDir); ensureDirErr != nil {
						log.Printf("خطا در ایجاد پوشه پایه: %v", ensureDirErr)
						// اگر حتی پوشه پایه هم ایجاد نشد، از مسیر فعلی استفاده کن
						prodDir = "."
					}
				} else {
					prodDir = fallbackDir
				}
			}

			savedPath, err := downloadAndSaveImage(ctx, im.Src, prodDir, fmt.Sprintf("img-%d", idx), true)
			if err != nil {
				// اگر دانلود شکست خورد، فقط لاگ بزن و این عکس رو رد کن
				log.Printf("خطا در دانلود عکس محصول (رد شد): %v - URL: %s", err, im.Src)
				continue // این عکس رو رد کن و برو سراغ عکس بعدی
			}

			// فقط اگر دانلود موفق بود، عکس رو ذخیره کن
			pi := models.ProductImage{ProductID: dest.ID, ImageURL: savedPath, SortOrder: idx}
			if err := uow.DB().WithContext(ctx).Create(&pi).Error; err != nil {
				log.Println("خطا در ثبت عکس محصول در دیتابیس:", err)
				continue // عکس را رد کن، محصول را ادامه بده
			}
			if idx == 0 {
				thumb = savedPath
			}
		}
		if thumb != "" {
			if err := uow.DB().WithContext(ctx).Model(&models.Product{}).Where("id = ?", dest.ID).Update("image", thumb).Error; err != nil {
				return false, err
			}
		}
	}

	// ویژگی‌ها و تنوع‌ها - انتقال دقیق ساختار ویژگی‌ها از وردپرس
	if transferAttributes {
		// حذف ویژگی‌های قبلی محصول (فقط ProductSpecification و ProductAttributeValue)
		if err := uow.DB().WithContext(ctx).Where("product_id = ?", dest.ID).Delete(&models.ProductSpecification{}).Error; err != nil {
			return false, err
		}
		if err := uow.DB().WithContext(ctx).Where("product_id = ?", dest.ID).Delete(&models.ProductAttributeValue{}).Error; err != nil {
			return false, err
		}

		// انتقال دقیق ویژگی‌ها از وردپرس
		for sortOrder, a := range p.Attributes {
			if len(a.Options) == 0 {
				continue
			}

			// یافتن یا ایجاد ویژگی در سیستم مقصد - دقیقاً همان ساختار وردپرس
			attrSlug := strings.TrimSpace(a.Slug)
			if attrSlug == "" {
				attrSlug = slugifyTitle(a.Name)
			}

			var attr models.Attribute
			if err := uow.DB().WithContext(ctx).Where("name = ? OR code = ?", a.Name, attrSlug).First(&attr).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// ایجاد ویژگی جدید - دقیقاً همان ساختار وردپرس
					attr = models.Attribute{
						Name:         a.Name,
						DisplayName:  a.Name,
						Code:         attrSlug,
						DataType:     "string", // پیش‌فرض ساده
						IsFilterable: true,
						IsActive:     true,
						HasOptions:   true,
						SortOrder:    uint(sortOrder),
					}
					if createErr := uow.DB().WithContext(ctx).Create(&attr).Error; createErr != nil {
						// اگر ویژگی ایجاد نشد، فقط به عنوان specification ذخیره کن
						val := strings.Join(a.Options, ", ")
						// val = fixEncoding(val) // حذف شد - مقادیر دقیقاً همان وردپرس باقی می‌مانند
						spec := models.ProductSpecification{ProductID: dest.ID, Name: a.Name, Value: val}
						if err := safeDBOperation(func() error {
							return uow.DB().WithContext(ctx).Create(&spec).Error
						}, "create specification (fallback)"); err != nil {
							log.Printf("[WARNING] خطا در ایجاد specification برای ویژگی %s: %v", a.Name, err)
						}
						continue
					}
				} else {
					// خطای دیگر - فقط به عنوان specification ذخیره کن
					val := strings.Join(a.Options, ", ")
					// val = fixEncoding(val) // حذف شد - مقادیر دقیقاً همان وردپرس باقی می‌مانند
					spec := models.ProductSpecification{ProductID: dest.ID, Name: a.Name, Value: val}
					if err := safeDBOperation(func() error {
						return uow.DB().WithContext(ctx).Create(&spec).Error
					}, "create specification (error fallback)"); err != nil {
						log.Printf("[WARNING] خطا در ایجاد specification برای ویژگی %s: %v", a.Name, err)
					}
					continue
				}
			}

			// انتقال دقیق مقادیر ویژگی از وردپرس
			for optionIndex, rawOption := range a.Options {
				option := strings.TrimSpace(rawOption) // فقط trim، بدون safeStringForDB
				if option == "" {
					continue
				}

				if !utf8.ValidString(option) {
					log.Printf("[WARNING] مقدار ویژگی '%s' encoding مشکل دارد، skip می‌شود", option)
					continue
				}

				optionSlug := slugifyTitle(option) // فقط برای slug از safeStringForDB استفاده کن
				if !utf8.ValidString(optionSlug) {
					log.Printf("[WARNING] slug ویژگی '%s' encoding مشکل دارد، skip می‌شود", optionSlug)
					continue
				}

				// یافتن یا ایجاد مقدار ویژگی - دقیقاً همان مقادیر وردپرس
				var attrValue models.AttributeValue
				if err := safeDBOperationWithRecovery(uow, func() error {
					return uow.DB().WithContext(ctx).Where("attribute_id = ? AND (value = ? OR slug = ?)", attr.ID, option, optionSlug).First(&attrValue).Error
				}, "find attribute value"); err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						// ایجاد مقدار جدید - دقیقاً همان مقدار وردپرس (بدون تغییر)
						attrValue = models.AttributeValue{
							AttributeID: attr.ID,
							Value:       option, // مقدار کامل بدون تغییر
							Slug:        optionSlug,
							Meta:        "{}",
						}

						if createErr := uow.DB().WithContext(ctx).Create(&attrValue).Error; createErr != nil {
							log.Printf("[ERROR] خطا در ایجاد مقدار ویژگی '%s': %v", option, createErr)
							continue
						}
					} else {
						continue
					}
				}

				// ایجاد ارتباط محصول-ویژگی-مقدار - دقیقاً همان ساختار وردپرس
				pav := models.ProductAttributeValue{
					ProductID:        uint(dest.ID),
					AttributeID:      attr.ID,
					AttributeValueID: &attrValue.ID,
					SortOrder:        uint(sortOrder*100 + optionIndex),
				}

				// پر کردن فیلد مناسب بر اساس نوع مقدار
				if attrValue.NumericValue != nil {
					pav.ValueNumeric = attrValue.NumericValue
				} else if attrValue.BoolValue != nil {
					pav.ValueBool = attrValue.BoolValue
				} else if attrValue.DateValue != nil {
					pav.ValueDate = attrValue.DateValue
				} else if attrValue.TextValue != nil {
					pav.ValueText = attrValue.TextValue
				} else {
					pav.ValueText = &option
				}

				if createErr := uow.DB().WithContext(ctx).Create(&pav).Error; createErr != nil {
					log.Printf("[WARNING] خطا در ایجاد ارتباط محصول-ویژگی: %v", createErr)
					continue
				}
			}

			// همچنین به عنوان specification هم ذخیره کن (برای سازگاری)
			// استفاده از strings.Builder برای جلوگیری از مشکل encoding
			var valBuilder strings.Builder
			for i, option := range a.Options {
				if i > 0 {
					valBuilder.WriteString(", ")
				}
				valBuilder.WriteString(option) // بدون هیچ پردازش encoding
			}
			val := valBuilder.String()
			spec := models.ProductSpecification{ProductID: dest.ID, Name: a.Name, Value: val}
			if err := safeDBOperation(func() error {
				return uow.DB().WithContext(ctx).Create(&spec).Error
			}, "create specification (compatibility)"); err != nil {
				log.Printf("[WARNING] خطا در ایجاد specification: %v", err)
			}
		}
	}
	// انتقال فیلدهای سئو از Yoast/RankMath در صورت انتخاب
	if len(p.MetaData) > 0 {
		var seoTitle, seoDesc string
		if transferSEOfromYoast {
			for _, md := range p.MetaData {
				lk := strings.ToLower(md.Key)
				if lk == "_yoast_wpseo_title" {
					if v, ok := md.Value.(string); ok {
						seoTitle = v
					}
				} else if lk == "_yoast_wpseo_metadesc" {
					if v, ok := md.Value.(string); ok {
						seoDesc = v
					}
				}
			}
		}
		if transferSEOfromRankMath {
			for _, md := range p.MetaData {
				lk := strings.ToLower(md.Key)
				if lk == "rank_math_title" {
					if v, ok := md.Value.(string); ok {
						seoTitle = v
					}
				} else if lk == "rank_math_description" {
					if v, ok := md.Value.(string); ok {
						seoDesc = v
					}
				}
			}
		}
		upd := map[string]interface{}{}
		if seoTitle != "" {
			upd["seo_title"] = seoTitle
		}
		if seoDesc != "" {
			upd["meta_description"] = seoDesc
		}
		if len(upd) > 0 {
			if err := uow.DB().WithContext(ctx).Model(&models.Product{}).Where("id = ?", dest.ID).Updates(upd).Error; err != nil {
				return false, err
			}
		}
	}

	// ساخت ریدایرکت از permalink مبدا به مسیر مقصد در صورت انتخاب
	if redirectProducts {
		var target string
		if dest.SKU != "" {
			if includeSlugInURL {
				target = fmt.Sprintf("/product/sku-%s/%s", dest.SKU, slug)
			} else {
				target = fmt.Sprintf("/product/sku-%s", dest.SKU)
			}
		} else {
			// تضمین: همیشه SKU باید باشد؛ این شاخه عملاً رخ نمی‌دهد.
			target = fmt.Sprintf("/product/sku-%s", dest.SKU)
		}

		// ذخیره URL جدید در فیلد url محصول
		if err := safeDBOperation(func() error {
			return uow.DB().WithContext(ctx).Model(&models.Product{}).Where("id = ?", dest.ID).Update("url", target).Error
		}, "update product URL"); err != nil {
			log.Printf("[WARNING] خطا در ذخیره URL محصول (ادامه می‌دهیم): %v", err)
		} else {
			log.Printf("[INFO] URL جدید محصول ذخیره شد: %s", target)
		}

		// تلاش برای ساخت ریدایرکت؛ اگر جدول نباشد (42P01) خطا را نادیده بگیر
		if err := s.CreateRedirect(uow, p.Permalink, target, redirectGroupName, redirectCode); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "42p01") {
				// جدول ریدایرکت وجود ندارد؛ صرفاً هشدار
				log.Printf("[WARNING] جدول ریدایرکت وجود ندارد، ریدایرکت ایجاد نشد")
			} else if strings.Contains(err.Error(), "current transaction is aborted") ||
				strings.Contains(err.Error(), "SQLSTATE 25P02") ||
				strings.Contains(err.Error(), "invalid byte sequence for encoding") ||
				strings.Contains(err.Error(), "SQLSTATE 22021") {
				log.Printf("[WARNING] خطای transaction/encoding در ایجاد ریدایرکت (ادامه می‌دهیم): %v", err)
				// ادامه می‌دهیم، ریدایرکت مهم نیست
			} else {
				log.Printf("[WARNING] خطا در ایجاد ریدایرکت (ادامه می‌دهیم): %v", err)
				// ادامه می‌دهیم، ریدایرکت مهم نیست
			}
		}
	}

	// ویدیوها از meta_data (کلیدهای شناخته‌شده)
	if err := s.importProductVideosFromMeta(uow, ctx, dest.ID, slug, p); err != nil {
		log.Println("خطا در انتقال ویدیوها:", err)
		// ادامه می‌دهیم، ویدیو مهم نیست
	}

	return created, nil
}

// Query helpers (read side)
func (s *WordPressMigrationService) FetchJSON(url string, apiKey string, out interface{}) error {
	// اعتبارسنجی URL برای جلوگیری از SSRF
	validatedURL, err := validateURLForNetworkRequest(url)
	if err != nil {
		return fmt.Errorf("URL نامعتبر: %v", err)
	}
	req, _ := http.NewRequest("GET", validatedURL, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("wp http status: %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

// FetchTest یک درخواست سبک به وردپرس ارسال می‌کند تا اتصال بررسی شود
// این تابع یک پست را از WP می‌خواند تا صحت آدرس و کلید API سنجیده شود
func (s *WordPressMigrationService) FetchTest(siteURL, apiKey string, out interface{}) error {
	base := normalizeSiteURL(siteURL)
	url := fmt.Sprintf("%s/wp-json/wp/v2/posts?per_page=1", base)
	return s.FetchJSON(url, apiKey, out)
}

// TestWooCustomers یک تست سبک برای اتصال به WooCommerce انجام می‌دهد
// از endpoint: /wp-json/wc/v3/customers?per_page=1 استفاده می‌کند
func (s *WordPressMigrationService) TestWooCustomers(siteURL, consumerKey, consumerSecret string) (int, error) {
	client := &http.Client{}
	versions := []string{"v3", "v2", "v1"}
	base := normalizeSiteURL(siteURL)

	// 1) Try query-string auth (common in WC REST)
	for _, ver := range versions {
		url := base + fmt.Sprintf("/wp-json/wc/%s/customers?per_page=1&consumer_key=%s&consumer_secret=%s", ver, consumerKey, consumerSecret)
		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(url)
		if err != nil {
			continue
		}
		req, _ := http.NewRequest("GET", validatedURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")
		resp, err := client.Do(req)
		if err == nil && resp != nil {
			bodyBytes, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				var arr []map[string]interface{}
				if err := json.Unmarshal(bodyBytes, &arr); err != nil {
					return 0, err
				}
				return len(arr), nil
			}
			// parse WC error body if present
			if len(bodyBytes) > 0 {
				var wcErr struct{ Code, Message string }
				if err := json.Unmarshal(bodyBytes, &wcErr); err == nil && wcErr.Message != "" {
					return 0, fmt.Errorf("woocommerce %s: %s", wcErr.Code, wcErr.Message)
				}
				return 0, fmt.Errorf("woocommerce http status: %d: %s", resp.StatusCode, string(bodyBytes))
			}
		}
	}

	// 2) Fallback to Basic Auth (some hosts disable query credentials)
	for _, ver := range versions {
		url := base + fmt.Sprintf("/wp-json/wc/%s/customers?per_page=1", ver)
		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(url)
		if err != nil {
			continue
		}
		req, _ := http.NewRequest("GET", validatedURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")
		basic := base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret))
		req.Header.Set("Authorization", "Basic "+basic)
		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			var arr []map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &arr); err != nil {
				return 0, err
			}
			return len(arr), nil
		}
		if len(bodyBytes) > 0 {
			var wcErr struct{ Code, Message string }
			if err := json.Unmarshal(bodyBytes, &wcErr); err == nil && wcErr.Message != "" {
				return 0, fmt.Errorf("woocommerce %s: %s", wcErr.Code, wcErr.Message)
			}
			return 0, fmt.Errorf("woocommerce http status: %d: %s", resp.StatusCode, string(bodyBytes))
		}
	}

	return 0, fmt.Errorf("woocommerce connection failed: no successful auth method")
}

// -----------------------------
// WooCommerce Customers (Digits)
// -----------------------------

// wcCustomer مدل مختصر مشتری ووکامرس برای استخراج کاربران
type wcCustomer struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
	Billing      struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
	} `json:"billing"`
	Shipping struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
	} `json:"shipping"`
	MetaData []struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	} `json:"meta_data"`
}

// parseWooCommerceDate تابع کمکی برای parse کردن تاریخ‌های WooCommerce
func parseWooCommerceDate(dateStr string) time.Time {
	if dateStr == "" {
		return time.Time{}
	}
	// WooCommerce معمولاً از فرمت ISO 8601 استفاده می‌کند
	layouts := []string{
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05.000Z",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, dateStr); err == nil {
			return t
		}
	}

	// اگر هیچ فرمتی کار نکرد، زمان فعلی را برمی‌گردانیم
	return time.Now()
}

func (s *WordPressMigrationService) fetchWCCustomersPage(ctx context.Context, siteURL, ck, cs string, page, perPage int) ([]wcCustomer, error) {
	base := normalizeSiteURL(siteURL)
	client := &http.Client{}
	versions := []string{"v3", "v2", "v1"}

	// Try query-string auth
	for _, ver := range versions {
		url := fmt.Sprintf("%s/wp-json/wc/%s/customers?per_page=%d&page=%d&consumer_key=%s&consumer_secret=%s", base, ver, perPage, page, ck, cs)
		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(url)
		if err != nil {
			continue
		}
		req, _ := http.NewRequestWithContext(ctx, "GET", validatedURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")
		resp, err := client.Do(req)
		if err == nil && resp != nil {
			bodyBytes, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				var arr []wcCustomer
				if err := json.Unmarshal(bodyBytes, &arr); err != nil {
					return nil, err
				}
				return arr, nil
			}
			// parse WC error body if present
			if len(bodyBytes) > 0 {
				var wcErr struct{ Code, Message string }
				if err := json.Unmarshal(bodyBytes, &wcErr); err == nil && wcErr.Message != "" {
					return nil, fmt.Errorf("woocommerce %s: %s", wcErr.Code, wcErr.Message)
				}
				return nil, fmt.Errorf("woocommerce http status: %d: %s", resp.StatusCode, string(bodyBytes))
			}
		}
	}

	// Fallback to Basic Auth
	for _, ver := range versions {
		url := fmt.Sprintf("%s/wp-json/wc/%s/customers?per_page=%d&page=%d", base, ver, perPage, page)
		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(url)
		if err != nil {
			continue
		}
		req, _ := http.NewRequestWithContext(ctx, "GET", validatedURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")
		basic := base64.StdEncoding.EncodeToString([]byte(ck + ":" + cs))
		req.Header.Set("Authorization", "Basic "+basic)
		resp, err := client.Do(req)
		if err == nil && resp != nil {
			bodyBytes, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				var arr []wcCustomer
				if err := json.Unmarshal(bodyBytes, &arr); err != nil {
					return nil, err
				}
				return arr, nil
			}
			if len(bodyBytes) > 0 {
				var wcErr struct{ Code, Message string }
				if err := json.Unmarshal(bodyBytes, &wcErr); err == nil && wcErr.Message != "" {
					return nil, fmt.Errorf("woocommerce %s: %s", wcErr.Code, wcErr.Message)
				}
				return nil, fmt.Errorf("woocommerce http status: %d: %s", resp.StatusCode, string(bodyBytes))
			}
		}
	}
	return nil, fmt.Errorf("cannot fetch WooCommerce customers")
}

func normalizeIranMobile(input string) string {
	s := strings.TrimSpace(input)
	// remove spaces and non-digit except plus
	cleaned := make([]rune, 0, len(s))
	for _, r := range s {
		if r >= '0' && r <= '9' {
			cleaned = append(cleaned, r)
		}
		if r == '+' {
			cleaned = append(cleaned, r)
		}
	}
	t := string(cleaned)
	if strings.HasPrefix(t, "+98") {
		t = "0" + strings.TrimPrefix(t, "+98")
	} else if strings.HasPrefix(t, "0098") {
		t = "0" + strings.TrimPrefix(t, "0098")
	}
	return t
}

// ImportDigitsUsers استخراج کاربران با تاکید بر موبایل دیجیتس از ووکامرس و درج/به‌روزرسانی در سیستم
func (s *WordPressMigrationService) ImportDigitsUsers(
	uCtx context.Context,
	uow unitofwork.UnitOfWork,
	siteURL, ck, cs string,
	batchSize int,
	replaceUsernameWithMobile bool,
	skipEmails bool,
	includeAddresses bool,
	logs *[]string,
	onLog func(logType, message string),
) (imported, skipped, failed int, err error) {
	if batchSize <= 0 {
		batchSize = 20
	}
	page := 1
	db := uow.DB()

	// اطمینان از وجود نقش customer و دریافت ID آن (برای جلوگیری از خطای FK نقش)
	getOrCreateRoleID := func(tx *gorm.DB, name, displayName string) (uint, error) {
		var row struct{ ID uint }
		if err := tx.Table("roles").Select("id").Where("name = ?", name).Take(&row).Error; err != nil {
			// ایجاد نقش در صورت نبود
			_ = tx.Exec(`INSERT INTO roles (name, display_name, description, priority, is_active) VALUES (?, ?, ?, ?, ?) ON CONFLICT (name) DO NOTHING`, name, displayName, "", 1, true).Error
			// تلاش مجدد برای خواندن
			if err2 := tx.Table("roles").Select("id").Where("name = ?", name).Take(&row).Error; err2 != nil {
				return 0, err2
			}
		}
		return row.ID, nil
	}
	customerRoleID, roleErr := getOrCreateRoleID(db, "customer", "مشتری")
	if roleErr != nil || customerRoleID == 0 {
		if logs != nil {
			*logs = append(*logs, fmt.Sprintf("cannot ensure customer role: %v", roleErr))
		}
		if onLog != nil {
			onLog("warning", fmt.Sprintf("cannot ensure customer role: %v", roleErr))
		}
		// ادامه می‌دهیم، اما درج کاربران ممکن است به خطای FK بخورد
	}

	for {
		// بررسی توقف قبل از خواندن صفحه بعد
		select {
		case <-uCtx.Done():
			if logs != nil {
				*logs = append(*logs, "عملیات توسط کاربر متوقف شد (قبل از خواندن صفحه)")
			}
			if onLog != nil {
				onLog("warning", "عملیات توسط کاربر متوقف شد")
			}
			return imported, skipped, failed, uCtx.Err()
		default:
		}

		customers, fetchErr := s.fetchWCCustomersPage(uCtx, siteURL, ck, cs, page, batchSize)
		if fetchErr != nil {
			if logs != nil {
				*logs = append(*logs, fmt.Sprintf("خطا در دریافت کاربران از صفحه %d: %v", page, fetchErr))
			}
			if onLog != nil {
				onLog("error", fmt.Sprintf("خطا در دریافت کاربران از صفحه %d: %v", page, fetchErr))
			}
			return imported, skipped, failed, fetchErr
		}
		if len(customers) == 0 {
			if logs != nil {
				*logs = append(*logs, fmt.Sprintf("انتقال کاربران دیجیتس تکمیل شد. صفحه %d خالی است.", page))
			}
			if onLog != nil {
				onLog("success", "انتقال کاربران دیجیتس تکمیل شد")
			}
			break
		}
		if logs != nil {
			*logs = append(*logs, fmt.Sprintf("در حال پردازش صفحه %d - %d کاربر", page, len(customers)))
		}
		if onLog != nil {
			onLog("info", fmt.Sprintf("در حال پردازش صفحه %d - %d کاربر", page, len(customers)))
		}

		for _, cst := range customers {
			// بررسی توقف در هر کاربر
			select {
			case <-uCtx.Done():
				if logs != nil {
					*logs = append(*logs, "عملیات توسط کاربر متوقف شد (حین پردازش کاربر)")
				}
				if onLog != nil {
					onLog("warning", "عملیات توسط کاربر متوقف شد")
				}
				return imported, skipped, failed, uCtx.Err()
			default:
			}

			// حذف SavePoint - استفاده از روش ساده‌تر
			// اولویت: متادیتاهای digits سپس billing.phone
			phone := ""
			for _, md := range cst.MetaData {
				k := strings.ToLower(md.Key)
				if k == "digits_phone" || k == "digits_phone_no" || k == "digits_phone_with_code" {
					if v, ok := md.Value.(string); ok && v != "" {
						phone = v
						break
					}
				}
			}
			if phone == "" {
				phone = cst.Billing.Phone
			}
			phone = normalizeIranMobile(phone)
			if phone == "" {
				skipped++
				if logs != nil {
					*logs = append(*logs, fmt.Sprintf("کاربر رد شد: شماره موبایل خالی یا نامعتبر - %s", cst.Username))
				}
				if onLog != nil {
					onLog("info", fmt.Sprintf("کاربر رد شد: شماره موبایل خالی یا نامعتبر - %s", cst.Username))
				}
				continue
			}

			username := cst.Username
			if username == "" {
				username = phone
			}
			fullName := strings.TrimSpace(strings.TrimSpace(cst.FirstName + " " + cst.LastName))
			if fullName == "" {
				fullName = username
			}

			// مدیریت ایمیل بر اساس تنظیمات skipEmails
			var email string
			if skipEmails {
				// اگر skipEmails فعال است، از -- استفاده کنیم (توجه: تداخل uniqueIndex ایمیل ممکن است رخ دهد)
				email = "--"
			} else {
				// اگر skipEmails غیرفعال است، از ایمیل اصلی یا ایمیل پیش‌فرض استفاده کنیم
				email = strings.TrimSpace(cst.Email)
				if email == "" {
					email = username + "@import.local"
				}
			}

			// استفاده از نقش customer که بالاتر اطمینان یافتیم
			defaultRoleID := customerRoleID

			// UPSERT بر اساس موبایل یا ایمیل یا یوزرنیم
			var existing models.User
			var tx *gorm.DB
			if email != "" && email != "--" {
				tx = db.Where("mobile = ? OR email = ? OR username = ?", phone, email, username).First(&existing)
			} else {
				// وقتی ایمیل نداریم، شرط ایمیل را حذف می‌کنیم تا از برخورد با مقادیر خالی جلوگیری شود
				tx = db.Where("mobile = ? OR username = ?", phone, username).First(&existing)
			}
			if tx.Error == nil && existing.ID != 0 {
				updates := map[string]interface{}{
					"name":            fullName,
					"mobile":          phone,
					"mobile_verified": true,
				}
				// فقط اگر skipEmails غیرفعال است، ایمیل را به‌روزرسانی کنیم
				if !skipEmails && email != "" && email != "--" {
					updates["email"] = email
				}
				if replaceUsernameWithMobile {
					updates["username"] = phone
				}
				// به‌روزرسانی تاریخ ثبت‌نام اگر از دیجیتس قدیمی‌تر باشد
				dateCreated := parseWooCommerceDate(cst.DateCreated)
				if !dateCreated.IsZero() && (existing.RegisteredAt.IsZero() || dateCreated.Before(existing.RegisteredAt)) {
					updates["registered_at"] = dateCreated
				}
				if err := db.Model(&models.User{}).Where("id = ?", existing.ID).Updates(updates).Error; err != nil {
					failed++
					if logs != nil {
						*logs = append(*logs, fmt.Sprintf("خطا در به‌روزرسانی کاربر موجود: %s (%s) - %v", fullName, phone, err))
					}
					if onLog != nil {
						onLog("error", fmt.Sprintf("خطا در به‌روزرسانی کاربر موجود: %s (%s)", fullName, phone))
					}
					_ = db.RollbackTo("sp_import_user").Error // Rollback to savepoint
					continue
				}
				// آدرس را در صورت نیاز به‌روزرسانی/ایجاد کنیم
				if includeAddresses {
					addrFullName := fullName

					// اولویت: Billing، سپس Shipping
					var street, city, state, postcode string

					// بررسی Billing
					if cst.Billing.Address1 != "" || cst.Billing.City != "" || cst.Billing.State != "" || cst.Billing.Postcode != "" {
						street = strings.TrimSpace(cst.Billing.Address1 + " " + cst.Billing.Address2)
						city = cst.Billing.City
						state = cst.Billing.State
						postcode = cst.Billing.Postcode
					} else if cst.Shipping.Address1 != "" || cst.Shipping.City != "" || cst.Shipping.State != "" || cst.Shipping.Postcode != "" {
						// Fallback به Shipping
						street = strings.TrimSpace(cst.Shipping.Address1 + " " + cst.Shipping.Address2)
						city = cst.Shipping.City
						state = cst.Shipping.State
						postcode = cst.Shipping.Postcode
					}

					if street != "" || city != "" || state != "" || postcode != "" {
						var ua models.UserAddress
						if err := db.Where("user_id = ? AND is_default = ?", existing.ID, true).First(&ua).Error; err == nil && ua.ID != 0 {
							// به‌روزرسانی آدرس موجود
							ua.Label = "Default"
							ua.Street = street
							ua.PostalCode = postcode
							ua.RecipientName = addrFullName
							ua.RecipientMobile = phone
							ua.Phone = phone
							ua.Province = state
							ua.City = city
							ua.IsDefault = true
							if saveErr := db.Save(&ua).Error; saveErr != nil {
								if logs != nil {
									*logs = append(*logs, fmt.Sprintf("خطا در به‌روزرسانی آدرس: %s - %v", fullName, saveErr))
								}
								if onLog != nil {
									onLog("warning", fmt.Sprintf("خطا در به‌روزرسانی آدرس: %s", fullName))
								}
							} else {
								if logs != nil {
									*logs = append(*logs, fmt.Sprintf("آدرس به‌روزرسانی شد: %s - %s، %s، %s", fullName, street, city, state))
								}
								if onLog != nil {
									onLog("info", fmt.Sprintf("آدرس به‌روزرسانی شد: %s - %s، %s", fullName, city, state))
								}
							}
						} else {
							// ایجاد آدرس جدید
							nua := models.UserAddress{
								UserID:          existing.ID,
								Label:           "Default",
								Street:          street,
								PostalCode:      postcode,
								RecipientName:   addrFullName,
								RecipientMobile: phone,
								Phone:           phone,
								Province:        state,
								City:            city,
								IsDefault:       true,
							}
							if createErr := db.Create(&nua).Error; createErr != nil {
								if logs != nil {
									*logs = append(*logs, fmt.Sprintf("خطا در ایجاد آدرس: %s - %v", fullName, createErr))
								}
								if onLog != nil {
									onLog("warning", fmt.Sprintf("خطا در ایجاد آدرس: %s", fullName))
								}
							} else {
								if logs != nil {
									*logs = append(*logs, fmt.Sprintf("آدرس ایجاد شد: %s - %s، %s، %s", fullName, street, city, state))
								}
								if onLog != nil {
									onLog("info", fmt.Sprintf("آدرس ایجاد شد: %s - %s، %s", fullName, city, state))
								}
							}
						}
					}
				}
				// کاربر موجود به‌روزرسانی شد - این imported بشماریم
				imported++
				if logs != nil {
					*logs = append(*logs, fmt.Sprintf("کاربر موجود به‌روزرسانی شد: %s (%s)", fullName, phone))
				}
				if onLog != nil {
					onLog("info", fmt.Sprintf("کاربر موجود به‌روزرسانی شد: %s (%s)", fullName, phone))
				}
				continue
			}

			// ایجاد کاربر جدید
			user := models.User{
				Username: func() string {
					if replaceUsernameWithMobile {
						return phone
					}
					return username
				}(),
				Mobile:         phone,
				Name:           fullName,
				MobileVerified: true,
				Status:         "active",
				RoleID:         defaultRoleID,
				RegisteredAt:   parseWooCommerceDate(cst.DateCreated), // استفاده از تاریخ ثبت‌نام دیجیتس
			}
			// اگر کاربر از وردپرس وارد می‌شود، همان شناسه مبدا را برای ID تنظیم می‌کنیم تا نگاشت سفارشات به مشکل نخورد
			if cst.ID > 0 {
				user.ID = uint(cst.ID)
			}
			// ست‌کردن ایمیل فقط زمانی که مجاز است؛ در غیر این صورت، ستون ایمیل را در INSERT حذف می‌کنیم
			var createErr error
			if !skipEmails && email != "" && email != "--" {
				user.Email = email
				createErr = db.Create(&user).Error
			} else {
				createErr = db.Omit("email").Create(&user).Error
			}
			if createErr != nil {
				failed++
				if logs != nil {
					*logs = append(*logs, fmt.Sprintf("خطا در ایجاد کاربر جدید: %s (%s) - %v", fullName, phone, createErr))
				}
				if onLog != nil {
					onLog("error", fmt.Sprintf("خطا در ایجاد کاربر جدید: %s (%s)", fullName, phone))
				}
				// فقط ادامه به کاربر بعدی - بدون rollback
				continue
			}
			// هم‌تراز کردن sequence جدول users با بیشینه ID فعلی پس از درج دستی ID برای سازگاری با BIGSERIAL
			if cst.ID > 0 {
				var seqName string
				if err := db.Raw("SELECT pg_get_serial_sequence('users','id')").Scan(&seqName).Error; err == nil && strings.TrimSpace(seqName) != "" {
					_ = db.Exec(fmt.Sprintf("SELECT setval('%s', (SELECT GREATEST(COALESCE(MAX(id),0)+1, 1) FROM users))", seqName)).Error
				}
			}
			// آدرس پیش‌فرض کاربر جدید در صورت انتخاب
			if includeAddresses {
				addrFullName := fullName

				// اولویت: Billing، سپس Shipping
				var street, city, state, postcode string

				// بررسی Billing
				if cst.Billing.Address1 != "" || cst.Billing.City != "" || cst.Billing.State != "" || cst.Billing.Postcode != "" {
					street = strings.TrimSpace(cst.Billing.Address1 + " " + cst.Billing.Address2)
					city = cst.Billing.City
					state = cst.Billing.State
					postcode = cst.Billing.Postcode
				} else if cst.Shipping.Address1 != "" || cst.Shipping.City != "" || cst.Shipping.State != "" || cst.Shipping.Postcode != "" {
					// Fallback به Shipping
					street = strings.TrimSpace(cst.Shipping.Address1 + " " + cst.Shipping.Address2)
					city = cst.Shipping.City
					state = cst.Shipping.State
					postcode = cst.Shipping.Postcode
				}

				if street != "" || city != "" || state != "" || postcode != "" {
					nua := models.UserAddress{
						UserID:          user.ID,
						Label:           "Default",
						Street:          street,
						PostalCode:      postcode,
						RecipientName:   addrFullName,
						RecipientMobile: phone,
						Phone:           phone,
						Province:        state,
						City:            city,
						IsDefault:       true,
					}
					if createAddrErr := db.Create(&nua).Error; createAddrErr != nil {
						if logs != nil {
							*logs = append(*logs, fmt.Sprintf("خطا در ایجاد آدرس: %s - %v", fullName, createAddrErr))
						}
						if onLog != nil {
							onLog("warning", fmt.Sprintf("خطا در ایجاد آدرس: %s", fullName))
						}
					} else {
						if logs != nil {
							*logs = append(*logs, fmt.Sprintf("آدرس ایجاد شد: %s - %s، %s، %s", fullName, street, city, state))
						}
						if onLog != nil {
							onLog("info", fmt.Sprintf("آدرس ایجاد شد: %s - %s، %s", fullName, city, state))
						}
					}
				}
			}
			imported++
			if logs != nil {
				*logs = append(*logs, fmt.Sprintf("کاربر جدید ایجاد شد: %s (%s)", user.Username, user.Mobile))
			}
			if onLog != nil {
				onLog("success", fmt.Sprintf("کاربر جدید ایجاد شد: %s (%s)", user.Username, user.Mobile))
			}
		}
		page++
	}
	if logs != nil {
		*logs = append(*logs, fmt.Sprintf("انتقال کاربران دیجیتس تکمیل شد. مجموع: %d انتقال یافت، %d رد شد، %d ناموفق", imported, skipped, failed))
	}
	if onLog != nil {
		onLog("success", fmt.Sprintf("کاربران دیجیتس: %d انتقال یافت، %d رد شد، %d ناموفق", imported, skipped, failed))
	}
	return imported, skipped, failed, nil
}

// CreateRedirect اگر مسیر منبع از مبدا موجود باشد، یک ریدایرکت 301 به ساختار مقصد ثبت می‌کند
func (s *WordPressMigrationService) CreateRedirect(uow unitofwork.UnitOfWork, sourceURL, targetPath, group string, code int) error {
	// استخراج path از URL کامل مبدا
	src := sourceURL
	if strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
		// حذف scheme و hostname
		// مثال: https://domain.com/path -> /path
		idx := strings.Index(src, "://")
		if idx >= 0 {
			tail := src[idx+3:]
			slash := strings.Index(tail, "/")
			if slash >= 0 {
				src = "/" + strings.TrimLeft(tail[slash+0:], "/")
			} else {
				src = "/"
			}
		}
	}
	if !strings.HasPrefix(src, "/") {
		src = "/" + src
	}
	if code == 0 {
		code = 301
	}
	r := models.SEORedirect{SourcePath: src, TargetPath: targetPath, Code: code, GroupName: group}
	// upsert ساده بر اساس SourcePath
	db := uow.DB()

	// چک وجود جدول بدون ایجاد خطا در تراکنش (Postgres): اگر جدول نبود، صرفاً رد شو
	var tblName *string
	if err := safeDBOperation(func() error {
		return db.Raw("SELECT to_regclass('public.seo_redirects')").Scan(&tblName).Error
	}, "check seo_redirects table"); err != nil {
		return err
	}
	if tblName == nil || *tblName == "" {
		return nil
	}

	var existing models.SEORedirect
	if err := safeDBOperation(func() error {
		return db.Where("source_path = ?", src).First(&existing).Error
	}, "check existing redirect"); err == nil && existing.ID != 0 {
		existing.TargetPath = targetPath
		existing.GroupName = group
		if err := safeDBOperation(func() error {
			return db.Save(&existing).Error
		}, "update existing redirect"); err != nil {
			return err
		}
		// fmt.Printf("Redirect updated: %s -> %s (group: %s)\n", src, targetPath, group)
		return nil
	}
	if err := safeDBOperation(func() error {
		return db.Create(&r).Error
	}, "create new redirect"); err != nil {
		return err
	}
	// fmt.Printf("Redirect created: %s -> %s (group: %s)\n", src, targetPath, group)
	return nil
}

// Command: upsert category by slug
func (s *WordPressMigrationService) UpsertCategory(uow unitofwork.UnitOfWork, cat models.Category) error {
	ctx := context.Background()
	// Try find by slug
	if existing, err := uow.CategoryRepository().GetBySlug(ctx, cat.Slug); err == nil && existing != nil {
		return nil // already exists
	} else if err != nil && err != gorm.ErrRecordNotFound {
		// اگر خطای دیگری غیر از "رکورد یافت نشد" رخ داد
		return err
	}
	// اگر رکورد یافت نشد (err == gorm.ErrRecordNotFound) یا existing == nil
	return uow.CategoryRepository().Create(ctx, &cat)
}

// Command: create product with images
func (s *WordPressMigrationService) CreateProduct(uow unitofwork.UnitOfWork, p models.Product, images []models.ProductImage) error {
	ctx := context.Background()
	if err := uow.ProductRepository().Create(ctx, &p); err != nil {
		return err
	}
	if len(images) > 0 {
		for i := range images {
			images[i].ProductID = p.ID
			if err := uow.DB().WithContext(ctx).Create(&images[i]).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// Command: create order with items
func (s *WordPressMigrationService) CreateOrder(uow unitofwork.UnitOfWork, o models.Order, items []models.OrderItem) error {
	if err := uow.DB().Create(&o).Error; err != nil {
		return err
	}
	for i := range items {
		items[i].OrderID = o.ID
		if err := uow.DB().Create(&items[i]).Error; err != nil {
			return err
		}
	}
	return nil
}

// -----------------------------
// WooCommerce Orders (Import)
// -----------------------------

type wcOrder struct {
	ID          int         `json:"id"`
	CustomerID  int         `json:"customer_id"`
	Status      string      `json:"status"`
	Total       interface{} `json:"total"`
	DateCreated string      `json:"date_created"`
	// فیلدهای پرداخت و حمل‌ونقل از ووکامرس برای نگاشت دقیق‌تر
	PaymentMethod      string      `json:"payment_method"`
	PaymentMethodTitle string      `json:"payment_method_title"`
	CustomerIP         string      `json:"customer_ip_address"`
	ShippingTotal      interface{} `json:"shipping_total"`
	ShippingLines      []struct {
		MethodID    string      `json:"method_id"`
		MethodTitle string      `json:"method_title"`
		Total       interface{} `json:"total"`
	} `json:"shipping_lines"`
	// کد رهگیری ارسال در برخی درگاه‌ها/پلاگین‌ها داخل متادیتا یا فیلدهای افزوده ذخیره می‌شود
	// در این نسخه از meta_data استخراج خواهد شد
	MetaData []struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	} `json:"meta_data"`
	Billing struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
	} `json:"billing"`
	Shipping struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
	} `json:"shipping"`
	LineItems []struct {
		ProductID int         `json:"product_id"`
		SKU       string      `json:"sku"`
		Name      string      `json:"name"`
		Quantity  int         `json:"quantity"`
		Price     interface{} `json:"price"`
		Total     interface{} `json:"total"`
	} `json:"line_items"`
}

func (s *WordPressMigrationService) FetchWCOrdersPage(ctx context.Context, siteURL, ck, cs string, page, perPage int) ([]wcOrder, error) {
	base := normalizeSiteURL(siteURL)
	client := &http.Client{}
	versions := []string{"v3", "v2", "v1"}
	var lastErr error
	// 1) Query-string credentials (رایج در WC)
	for _, ver := range versions {
		url := fmt.Sprintf("%s/wp-json/wc/%s/orders?status=any&per_page=%d&page=%d&consumer_key=%s&consumer_secret=%s", base, ver, perPage, page, ck, cs)
		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(url)
		if err != nil {
			continue
		}
		req, _ := http.NewRequestWithContext(ctx, "GET", validatedURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")
		resp, err := client.Do(req)
		if err == nil && resp != nil {
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				var arr []wcOrder
				if err := json.Unmarshal(body, &arr); err != nil {
					return nil, err
				}
				return arr, nil
			} else {
				if len(body) > 0 {
					lastErr = fmt.Errorf("woocommerce orders %s http %d: %s", ver, resp.StatusCode, string(body))
				} else {
					lastErr = fmt.Errorf("woocommerce orders %s http %d", ver, resp.StatusCode)
				}
			}
		}
	}
	// 2) Fallback: Basic Auth (برخی هاست‌ها کوئری استرینگ را می‌بندند)
	for _, ver := range versions {
		url := fmt.Sprintf("%s/wp-json/wc/%s/orders?status=any&per_page=%d&page=%d", base, ver, perPage, page)
		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(url)
		if err != nil {
			continue
		}
		req, _ := http.NewRequestWithContext(ctx, "GET", validatedURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")
		if ck != "" && cs != "" {
			basic := base64.StdEncoding.EncodeToString([]byte(ck + ":" + cs))
			req.Header.Set("Authorization", "Basic "+basic)
		}
		resp, err := client.Do(req)
		if err == nil && resp != nil {
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				var arr []wcOrder
				if err := json.Unmarshal(body, &arr); err != nil {
					return nil, err
				}
				return arr, nil
			} else {
				if len(body) > 0 {
					lastErr = fmt.Errorf("woocommerce orders %s http %d: %s", ver, resp.StatusCode, string(body))
				} else {
					lastErr = fmt.Errorf("woocommerce orders %s http %d", ver, resp.StatusCode)
				}
			}
		}
	}
	// 3) Legacy Woo endpoints: /wc-api/v{ver}/orders
	for _, ver := range versions { // v3, v2, v1
		url := fmt.Sprintf("%s/wc-api/%s/orders?status=any&per_page=%d&page=%d&consumer_key=%s&consumer_secret=%s", base, ver, perPage, page, ck, cs)
		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(url)
		if err != nil {
			continue
		}
		req, _ := http.NewRequestWithContext(ctx, "GET", validatedURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")
		resp, err := client.Do(req)
		if err == nil && resp != nil {
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				// legacy ممکن است شکل پاسخ متفاوت باشد، تلاش برای استخراج فهرست سفارش‌ها
				var legacy struct {
					Orders []wcOrder `json:"orders"`
				}
				if err := json.Unmarshal(body, &legacy); err == nil && len(legacy.Orders) > 0 {
					return legacy.Orders, nil
				}
				// یا شاید آرایه مستقیم
				var arr []wcOrder
				if err := json.Unmarshal(body, &arr); err == nil {
					return arr, nil
				}
			} else {
				if len(body) > 0 {
					lastErr = fmt.Errorf("woocommerce legacy orders %s http %d: %s", ver, resp.StatusCode, string(body))
				} else {
					lastErr = fmt.Errorf("woocommerce legacy orders %s http %d", ver, resp.StatusCode)
				}
			}
		}
	}
	if lastErr != nil {
		return nil, lastErr
	}
	return []wcOrder{}, nil
}

func (s *WordPressMigrationService) ImportWooOrders(
	uCtx context.Context,
	uow unitofwork.UnitOfWork,
	siteURL, ck, cs string,
	batchSize int,
	_ bool,
	_ bool,
	logs *[]string,
	onLog func(logType, message string),
) (imported, skipped, failed int, err error) {
	if batchSize <= 0 {
		batchSize = 20
	}
	page := 1
	db := uow.DB()

	// اطمینان از وجود نقش customer و دریافت ID آن
	getOrCreateRoleID := func(tx *gorm.DB, name, displayName string) (uint, error) {
		var row struct{ ID uint }
		if err := tx.Table("roles").Select("id").Where("name = ?", name).Take(&row).Error; err != nil {
			// ایجاد نقش در صورت نبود
			_ = tx.Exec(`INSERT INTO roles (name, display_name, description, priority, is_active) VALUES (?, ?, ?, ?, ?) ON CONFLICT (name) DO NOTHING`, name, displayName, "", 1, true).Error
			// تلاش مجدد برای خواندن
			if err2 := tx.Table("roles").Select("id").Where("name = ?", name).Take(&row).Error; err2 != nil {
				return 0, err2
			}
		}
		return row.ID, nil
	}
	customerRoleID, roleErr := getOrCreateRoleID(db, "customer", "مشتری")
	if roleErr != nil || customerRoleID == 0 {
		if onLog != nil {
			onLog("warning", fmt.Sprintf("cannot ensure customer role: %v", roleErr))
		}
		customerRoleID = 1 // fallback
	}

	// شمارنده برای شناسه‌های تراکنش خالی
	emptyTransactionCounter := 1

	// لاگ شروع انتقال
	if onLog != nil {
		onLog("info", fmt.Sprintf("شروع انتقال سفارشات از %s (CK: %s, CS: %s)", siteURL, ck, cs))
	}
	for {
		select {
		case <-uCtx.Done():
			if onLog != nil {
				onLog("warning", "سفارشات: عملیات متوقف شد")
			}
			return imported, skipped, failed, uCtx.Err()
		default:
		}
		orders, ferr := s.FetchWCOrdersPage(uCtx, siteURL, ck, cs, page, batchSize)
		if ferr != nil {
			if onLog != nil {
				onLog("error", fmt.Sprintf("خطا در دریافت سفارشات از صفحه %d: %s", page, ferr.Error()))
			}
			return imported, skipped, failed, ferr
		}
		if len(orders) == 0 {
			if onLog != nil {
				onLog("info", fmt.Sprintf("صفحه %d: سفارش جدیدی یافت نشد", page))
			}
			break
		}
		if onLog != nil {
			onLog("info", fmt.Sprintf("در حال پردازش %d سفارش از صفحه %d", len(orders), page))
		}
		for _, wo := range orders {
			// ایجاد savepoint برای هر سفارش
			if err := db.SavePoint("sp_import_order").Error; err != nil {
				if onLog != nil {
					onLog("error", fmt.Sprintf("order %d: failed to create savepoint: %v", wo.ID, err))
				}
				skipped++
				continue
			}
			select {
			case <-uCtx.Done():
				if onLog != nil {
					onLog("warning", "سفارشات: عملیات متوقف شد")
				}
				return imported, skipped, failed, uCtx.Err()
			default:
			}
			// پیدا کردن کاربر بر اساس معیارهای مختلف (نام، wp_user_id، موبایل/ایمیل)
			// منطق استخراج موبایل مانند دیجیتس: اول متادیتاهای digits سپس billing.phone
			rawPhone := ""
			for _, md := range wo.MetaData {
				k := strings.ToLower(strings.TrimSpace(md.Key))
				if k == "digits_phone" || k == "digits_phone_no" || k == "digits_phone_with_code" {
					if v, ok := md.Value.(string); ok && strings.TrimSpace(v) != "" {
						rawPhone = v
						break
					}
				}
			}
			if rawPhone == "" {
				rawPhone = wo.Billing.Phone
			}
			phone := normalizeIranMobile(rawPhone)
			email := strings.TrimSpace(wo.Billing.Email)
			var user models.User
			foundUser := false
			// 1) بر اساس نام و نام خانوادگی (در فیلد name)
			fullName := strings.TrimSpace(strings.TrimSpace(wo.Billing.FirstName + " " + wo.Billing.LastName))
			if fullName != "" {
				if err := db.Where("name = ?", fullName).First(&user).Error; err == nil && user.ID != 0 {
					foundUser = true
					if onLog != nil {
						onLog("info", fmt.Sprintf("سفارش %d: کاربر با نام کامل یافت شد (%s) - ID: %d", wo.ID, fullName, user.ID))
					}
				}
			}
			// 2) بر اساس ID (مستقیم) - نگاشت مستقیم بر اساس کلید اصلی users.id برابر با شناسه وردپرس
			if !foundUser && wo.CustomerID != 0 {
				if err := db.Where("id = ?", wo.CustomerID).First(&user).Error; err == nil && user.ID != 0 {
					foundUser = true
					if onLog != nil {
						onLog("info", fmt.Sprintf("سفارش %d: کاربر یافت شد با ID %d", wo.ID, wo.CustomerID))
					}
				}
			}
			// 3) موبایل/ایمیل
			if !foundUser {
				qry := db
				if phone != "" && email != "" {
					qry = qry.Where("mobile = ? OR email = ?", phone, email)
				} else if phone != "" {
					qry = qry.Where("mobile = ?", phone)
				} else if email != "" {
					qry = qry.Where("email = ?", email)
				} else {
					skipped++
					if onLog != nil {
						onLog("warning", fmt.Sprintf("رد سفارش %d: اطلاعات تماس مشتری خالی است", wo.ID))
					}
					_ = db.RollbackTo("sp_import_order").Error
					continue
				}
				if err := qry.First(&user).Error; err == nil && user.ID != 0 {
					foundUser = true
				}
			}
			if !foundUser {
				// اگر کاربر یافت نشد، ابتدا سعی کنیم کاربر را ایجاد کنیم
				if onLog != nil {
					onLog("info", fmt.Sprintf("سفارش %d: کاربر یافت نشد، در حال ایجاد کاربر جدید (phone='%s', email='%s', name='%s')", wo.ID, phone, email, fullName))
				}

				// ایجاد کاربر جدید بر اساس اطلاعات سفارش
				newUser := models.User{
					Username:       phone, // استفاده از موبایل به عنوان username
					Mobile:         phone,
					Name:           fullName,
					Email:          email,
					MobileVerified: true,
					Status:         "active",
					RoleID:         customerRoleID, // استفاده از نقش customer که بالاتر اطمینان یافتیم
					RegisteredAt:   time.Now(),     // استفاده از زمان فعلی
				}

				if err := db.Create(&newUser).Error; err != nil {
					if onLog != nil {
						onLog("error", fmt.Sprintf("سفارش %d: خطا در ایجاد کاربر جدید: %v", wo.ID, err))
					}
					skipped++
					_ = db.RollbackTo("sp_import_order").Error
					continue
				}

				user = newUser
				foundUser = true
				if onLog != nil {
					onLog("success", fmt.Sprintf("سفارش %d: کاربر جدید ایجاد شد (ID: %d, Name: %s)", wo.ID, user.ID, user.Name))
				}
			} else {
				// کاربر موجود یافت شد - فقط لاگ کن و ادامه بده (بدون به‌روزرسانی)
				if onLog != nil {
					onLog("info", fmt.Sprintf("سفارش %d: کاربر موجود یافت شد (ID: %d, Name: %s) - بدون تغییر", wo.ID, user.ID, user.Name))
				}
			}
			// تضمین وجود آدرس ارسال و دریافت ID آن
			street := strings.TrimSpace(strings.TrimSpace(wo.Shipping.Address1 + " " + wo.Shipping.Address2))
			if street == "" && (wo.Shipping.City == "" && wo.Shipping.State == "" && wo.Shipping.Postcode == "") {
				// fallback به billing اگر shipping خالی بود
				street = strings.TrimSpace(strings.TrimSpace(wo.Billing.Address1 + " " + wo.Billing.Address2))
			}
			var shipAddrID uint
			if street != "" || wo.Shipping.City != "" || wo.Shipping.State != "" || wo.Shipping.Postcode != "" {
				var ua models.UserAddress
				// تلاش برای یافتن آدرسی با همین ویژگی‌ها
				if err := db.Where("user_id = ? AND street = ? AND city = ? AND province = ? AND postal_code = ?", user.ID, street, wo.Shipping.City, wo.Shipping.State, wo.Shipping.Postcode).First(&ua).Error; err == nil && ua.ID != 0 {
					// آدرس موجود یافت شد - استفاده از همان
					shipAddrID = ua.ID
					if onLog != nil {
						onLog("info", fmt.Sprintf("سفارش %d: آدرس ارسال موجود یافت شد (ID: %d) - بدون تغییر", wo.ID, ua.ID))
					}
				} else {
					// آدرس موجود نیست - ایجاد آدرس جدید
					nua := models.UserAddress{
						UserID:        user.ID,
						Label:         fmt.Sprintf("Order %d Shipping", wo.ID),
						Street:        street,
						PostalCode:    wo.Shipping.Postcode,
						RecipientName: strings.TrimSpace(wo.Shipping.FirstName + " " + wo.Shipping.LastName),
						Phone:         phone,
						Province:      wo.Shipping.State,
						City:          wo.Shipping.City,
						IsDefault:     false,
					}
					if err := db.Create(&nua).Error; err == nil {
						shipAddrID = nua.ID
						if onLog != nil {
							onLog("success", fmt.Sprintf("سفارش %d: آدرس ارسال جدید ایجاد شد (ID: %d)", wo.ID, nua.ID))
						}
					} else {
						if onLog != nil {
							onLog("error", fmt.Sprintf("سفارش %d: خطا در ایجاد آدرس ارسال: %v", wo.ID, err))
						}
					}
				}
			}
			if shipAddrID == 0 {
				// حداقل یک آدرس بسازیم تا محدودیت not null حفظ شود
				nua := models.UserAddress{UserID: user.ID, Label: fmt.Sprintf("Order %d", wo.ID), Street: street, City: wo.Shipping.City, Province: wo.Shipping.State, PostalCode: wo.Shipping.Postcode, Phone: phone, IsDefault: false}
				if err := db.Create(&nua).Error; err != nil {
					if onLog != nil {
						onLog("error", fmt.Sprintf("order %d: failed to create shipping address: %v", wo.ID, err))
					}
					// اگر آدرس ایجاد نشد، سفارش را رد کن
					skipped++
					_ = db.RollbackTo("sp_import_order").Error
					continue
				}
				shipAddrID = nua.ID
			}

			// بررسی صحت آدرس ارسال نهایی
			if shipAddrID == 0 {
				if onLog != nil {
					onLog("error", fmt.Sprintf("order %d: invalid shipping address ID", wo.ID))
				}
				skipped++
				_ = db.RollbackTo("sp_import_order").Error
				continue
			}
			// استخراج اطلاعات حمل و نقل و پرداخت برای نگاشت دقیق
			// روش ارسال و هزینه حمل
			shippingMethod := "" // عنوان روش ارسال (برای توضیحات)
			var shippingCost float64
			if len(wo.ShippingLines) > 0 {
				shippingMethod = strings.TrimSpace(wo.ShippingLines[0].MethodTitle)
				var tmp float64
				switch v := wo.ShippingLines[0].Total.(type) {
				case string:
					fmt.Sscanf(strings.TrimSpace(v), "%f", &tmp)
				case float64:
					tmp = v
				case int, int64:
					tmp = float64(reflect.ValueOf(v).Int())
				}
				shippingCost += tmp
				if len(wo.ShippingLines) > 1 {
					for i := 1; i < len(wo.ShippingLines); i++ {
						var t2 float64
						switch vv := wo.ShippingLines[i].Total.(type) {
						case string:
							fmt.Sscanf(strings.TrimSpace(vv), "%f", &t2)
						case float64:
							t2 = vv
						case int, int64:
							t2 = float64(reflect.ValueOf(vv).Int())
						}
						shippingCost += t2
					}
				}
			} else if wo.ShippingTotal != nil {
				switch v := wo.ShippingTotal.(type) {
				case string:
					fmt.Sscanf(strings.TrimSpace(v), "%f", &shippingCost)
				case float64:
					shippingCost = v
				case int, int64:
					shippingCost = float64(reflect.ValueOf(v).Int())
				}
			}

			if shippingMethod == "" {
				if onLog != nil {
					onLog("error", fmt.Sprintf("order %d: missing shipping method title — defaulting to 'post'", wo.ID))
				}
			}
			if wo.ShippingTotal == nil && len(wo.ShippingLines) == 0 {
				if onLog != nil {
					onLog("error", fmt.Sprintf("order %d: missing shipping cost data — defaulting to 0", wo.ID))
				}
			}

			// نگاشت عنوان روش ارسال ووکامرس به کدهای مجاز سیستم مقصد
			// مجاز: post, tipax, chapar, freight, pickup, courier
			mapShippingMethod := func(title string) string {
				t := strings.ToLower(strings.TrimSpace(title))
				switch {
				case strings.Contains(t, "تیپاکس") || strings.Contains(t, "tipax"):
					return "tipax"
				case strings.Contains(t, "چاپار") || strings.Contains(t, "chapar"):
					return "chapar"
				case strings.Contains(t, "باربری") || strings.Contains(t, "freight") || strings.Contains(t, "carrier"):
					return "freight"
				case strings.Contains(t, "پیک") || strings.Contains(t, "پیک موتوری") || strings.Contains(t, "courier") || strings.Contains(t, "motor"):
					return "courier"
				case strings.Contains(t, "تحویل حضوری") || strings.Contains(t, "pickup") || strings.Contains(t, "local pickup"):
					return "pickup"
				case strings.Contains(t, "post") || strings.Contains(t, "پست") || strings.Contains(t, "پیشتاز"):
					return "post"
				default:
					return "post"
				}
			}
			shippingMethodCode := mapShippingMethod(shippingMethod)

			// تعیین وضعیت پرداخت و روش پرداخت
			pmSlug := strings.ToLower(strings.TrimSpace(wo.PaymentMethod))
			pmTitle := strings.TrimSpace(wo.PaymentMethodTitle)
			// نگاشت وضعیت پرداخت - بهبود یافته
			paymentStatus := "awaiting_payment"
			isPaid := false
			switch strings.ToLower(strings.TrimSpace(wo.Status)) {
			case "completed", "shipped", "delivered", "processing", "on-hold":
				paymentStatus = "paid"
				isPaid = true
			case "failed", "cancelled":
				paymentStatus = "failed"
				isPaid = false
			case "refunded":
				paymentStatus = "refunded"
				isPaid = false
			case "returned":
				paymentStatus = "returned"
				isPaid = false
			default:
				paymentStatus = "awaiting_payment"
				isPaid = false
			}

			// درج سفارش اگر با tracking_code WO-<id> وجود ندارد
			var existing models.Order
			trk := fmt.Sprintf("WO-%d", wo.ID)
			if err := db.Where("tracking_code = ?", trk).First(&existing).Error; err == nil && existing.ID != 0 {
				// سفارش قبلاً وجود دارد
				_ = db.RollbackTo("sp_import_order").Error
				continue
			}
			createdAt := parseWooCommerceDate(wo.DateCreated)
			if createdAt.IsZero() {
				if onLog != nil {
					onLog("error", fmt.Sprintf("order %d: invalid date created '%v' — using current time", wo.ID, wo.DateCreated))
				}
				createdAt = time.Now()
			}
			var total float64
			if wo.Total != nil {
				switch v := wo.Total.(type) {
				case string:
					if _, err := fmt.Sscanf(strings.TrimSpace(v), "%f", &total); err != nil {
						if onLog != nil {
							onLog("error", fmt.Sprintf("order %d: invalid total amount '%v' — using 0", wo.ID, v))
						}
						total = 0
					}
				case float64:
					total = v
				case int, int64:
					total = float64(reflect.ValueOf(v).Int())
				default:
					if onLog != nil {
						onLog("error", fmt.Sprintf("order %d: unsupported total type %T — using 0", wo.ID, v))
					}
					total = 0
				}
			} else {
				if onLog != nil {
					onLog("error", fmt.Sprintf("order %d: missing total amount — using 0", wo.ID))
				}
				total = 0
			}
			// تعیین روش پرداخت سفارش (cod یا online)
			paymentMethod := "online"
			if pmSlug == "cod" || strings.Contains(strings.ToLower(pmTitle), "cod") || strings.Contains(strings.ToLower(pmTitle), "pay on delivery") {
				paymentMethod = "cod"
			}
			if pmSlug == "" && pmTitle == "" {
				if onLog != nil {
					onLog("error", fmt.Sprintf("order %d: missing payment method — defaulting to 'online'", wo.ID))
				}
			}
			// وضعیت سفارش - نگاشت کامل وضعیت‌های وردپرس
			statusMap := map[string]string{
				"pending":     "pending",
				"processing":  "processing",
				"on-hold":     "processing",
				"completed":   "shipped",
				"shipped":     "shipped",
				"delivered":   "shipped",
				"cancelled":   "cancelled",
				"refunded":    "refunded",
				"failed":      "cancelled",
				"returned":    "returned",
				"partial":     "processing",
				"backordered": "processing",
			}
			mappedStatus := statusMap[strings.ToLower(strings.TrimSpace(wo.Status))]
			if mappedStatus == "" {
				mappedStatus = "pending"
				if onLog != nil {
					onLog("warning", fmt.Sprintf("سفارش %d: وضعیت نامشخص '%s' - به pending تغییر یافت", wo.ID, wo.Status))
				}
			}
			// جمع کل: مقدار total ووکامرس شامل حمل و نقل است
			o := models.Order{UserID: &user.ID, ShippingAddressID: shipAddrID, PaymentMethod: paymentMethod, Status: mappedStatus, PaymentStatus: paymentStatus, IsPaid: isPaid, TotalAmount: total, FinalAmount: total, CreatedAt: createdAt, UpdatedAt: createdAt, TrackingCode: trk}
			if err := db.Create(&o).Error; err != nil {
				if s.handleTransactionError(db, wo.ID, err, onLog) {
					skipped++
					continue
				}
				failed++
				if onLog != nil {
					onLog("error", fmt.Sprintf("خطا در ایجاد سفارش %d: %v", wo.ID, err))
				}
				_ = db.RollbackTo("sp_import_order").Error
				continue
			}
			// آیتم‌ها (با savepoint، رد نرم هر آیتم مشکل‌دار)
			createdItems := 0
			for _, li := range wo.LineItems {
				var prod *models.Product
				prodRepo := uow.ProductRepository()
				if li.SKU != "" {
					if p, err := prodRepo.GetBySKU(context.Background(), li.SKU); err == nil && p != nil {
						prod = p
					}
				}

				// سپس بر اساس نام محصول جستجو می‌کنیم (نه ProductID وردپرس)
				if prod == nil && li.Name != "" {
					var pp models.Product
					if err := db.Where("name = ?", li.Name).First(&pp).Error; err == nil && pp.ID != 0 {
						prod = &pp
					}
				}
				// اگر محصول پیدا نشد، فقط نام را ذخیره می‌کنیم
				if prod == nil {
					productName := li.Name
					if productName == "" {
						productName = fmt.Sprintf("محصول با SKU: %s", li.SKU)
					}

					if onLog != nil {
						onLog("info", fmt.Sprintf("order %d: product not found for line item '%s' (sku='%s', pid=%d) — storing as text only", wo.ID, productName, li.SKU, li.ProductID))
					}
					// ایجاد یک محصول dummy با ID 0 برای OrderItem
					prod = &models.Product{
						ID:   0,
						Name: productName,
					}
				} else {
					// محصول پیدا شد - لاگ موفقیت
					if onLog != nil {
						onLog("info", fmt.Sprintf("order %d: product found for line item '%s' (sku='%s', pid=%d) — normal mapping", wo.ID, prod.Name, li.SKU, li.ProductID))
					}
				}
				var unit float64
				if li.Price != nil {
					switch pv := li.Price.(type) {
					case string:
						if _, err := fmt.Sscanf(strings.TrimSpace(pv), "%f", &unit); err != nil {
							if onLog != nil {
								onLog("error", fmt.Sprintf("order %d: invalid unit price '%v' for product %d — using 0", wo.ID, pv, prod.ID))
							}
							unit = 0
						}
					case float64:
						unit = pv
					case int, int64:
						unit = float64(reflect.ValueOf(pv).Int())
					default:
						if onLog != nil {
							onLog("error", fmt.Sprintf("order %d: unsupported unit price type %T for product %d — using 0", wo.ID, pv, prod.ID))
						}
						unit = 0
					}
				} else {
					if onLog != nil {
						onLog("error", fmt.Sprintf("order %d: missing unit price for product %d — using 0", wo.ID, prod.ID))
					}
					unit = 0
				}

				var totalItem float64
				if li.Total != nil {
					switch tv := li.Total.(type) {
					case string:
						if _, err := fmt.Sscanf(strings.TrimSpace(tv), "%f", &totalItem); err != nil {
							if onLog != nil {
								onLog("error", fmt.Sprintf("order %d: invalid total price '%v' for product %d — using 0", wo.ID, tv, prod.ID))
							}
							totalItem = 0
						}
					case float64:
						totalItem = tv
					case int, int64:
						totalItem = float64(reflect.ValueOf(tv).Int())
					default:
						if onLog != nil {
							onLog("error", fmt.Sprintf("order %d: unsupported total price type %T for product %d — using 0", wo.ID, tv, prod.ID))
						}
						totalItem = 0
					}
				} else {
					if onLog != nil {
						onLog("error", fmt.Sprintf("order %d: missing total price for product %d — using 0", wo.ID, prod.ID))
					}
					totalItem = 0
				}

				// بررسی صحت مقدار quantity
				quantity := li.Quantity
				if quantity <= 0 {
					if onLog != nil {
						onLog("error", fmt.Sprintf("order %d: invalid quantity %d for product %d — using 1", wo.ID, quantity, prod.ID))
					}
					quantity = 1
				}
				it := models.OrderItem{OrderID: o.ID, ProductID: uint(prod.ID), Quantity: quantity, UnitPrice: unit, FinalPrice: totalItem}
				if err := db.Create(&it).Error; err != nil {
					skipped++
					if onLog != nil {
						onLog("error", fmt.Sprintf("order %d: failed to insert order item for product %d — item skipped", wo.ID, prod.ID))
					}
				} else {
					createdItems++
				}
			}
			if createdItems == 0 {
				// هیچ آیتمی نگاشت نشد؛ اما سفارش را نگه می‌داریم
				if onLog != nil {
					onLog("warning", fmt.Sprintf("order %d: no order items mapped — keeping order without items", wo.ID))
				}
				// سفارش را نگه می‌داریم حتی اگر آیتمی نداشته باشد
			}

			// ایجاد تراکنش پرداخت (PaymentTransaction) برای نگهداری شماره تراکنش/درگاه
			// استخراج transaction_id از متادیتای سفارش
			txID := ""
			shippingTracking := ""
			for _, md := range wo.MetaData {
				k := strings.ToLower(strings.TrimSpace(md.Key))
				if k == "_transaction_id" || k == "transaction_id" {
					switch v := md.Value.(type) {
					case string:
						txID = strings.TrimSpace(v)
					default:
						txID = fmt.Sprintf("%v", v)
					}
					continue
				}
				// تلاش برای استخراج کد رهگیری ارسال از کلیدهای متداول
				if shippingTracking == "" {
					if k == "_tracking_number" || k == "tracking_number" || k == "shipping_tracking_number" || k == "tracking-code" || k == "yith_tracking_code" {
						switch v := md.Value.(type) {
						case string:
							shippingTracking = strings.TrimSpace(v)
						default:
							shippingTracking = fmt.Sprintf("%v", v)
						}
					}
				}
			}
			if txID == "" {
				if onLog != nil {
					onLog("info", fmt.Sprintf("order %d: no transaction_id found — using sequential ID: %d", wo.ID, emptyTransactionCounter))
				}
				// استفاده از شمارنده متوالی برای تضمین یکتا بودن
				txID = fmt.Sprintf("TXN-%d", emptyTransactionCounter)
				emptyTransactionCounter++
			}

			// یافتن/ایجاد درگاه پرداخت بر اساس pmSlug/pmTitle
			gatewayName := pmTitle
			gatewayEnglish := pmSlug
			if gatewayEnglish == "" {
				gatewayEnglish = "imported"
			}
			if gatewayName == "" {
				gatewayName = strings.ToUpper(gatewayEnglish)
			}
			var gateway models.PaymentGateway
			if err := db.Where("english_name = ?", gatewayEnglish).First(&gateway).Error; err != nil || gateway.ID == 0 {
				// ساخت درگاه حداقلی در صورت نبود
				gateway = models.PaymentGateway{
					Name:        gatewayName,
					EnglishName: gatewayEnglish,
					Type:        "other",
					Status:      "active",
					IsTestMode:  true,
					Priority:    1,
				}
				if err := db.Create(&gateway).Error; err != nil {
					if onLog != nil {
						onLog("error", fmt.Sprintf("order %d: failed to create payment gateway '%s': %v", wo.ID, gatewayEnglish, err))
					}
					// اگر درگاه ایجاد نشد، سفارش را رد کن
					skipped++
					_ = db.RollbackTo("sp_import_order").Error
					continue
				}
			}

			// بررسی صحت درگاه پرداخت نهایی
			if gateway.ID == 0 {
				if onLog != nil {
					onLog("error", fmt.Sprintf("order %d: invalid payment gateway ID", wo.ID))
				}
				skipped++
				_ = db.RollbackTo("sp_import_order").Error
				continue
			}

			ptStatus := "pending"
			if isPaid {
				ptStatus = "success"
			} else if strings.ToLower(strings.TrimSpace(wo.Status)) == "failed" {
				ptStatus = "failed"
			}
			pt := models.PaymentTransaction{
				GatewayID:     gateway.ID,
				OrderID:       trk,
				TransactionID: txID,
				Amount:        int64(total),
				Status:        ptStatus,
				PaymentMethod: func() string {
					if pmSlug != "" {
						return pmSlug
					}
					if pmTitle != "" {
						return pmTitle
					}
					return "imported"
				}(),
				Description: fmt.Sprintf("Imported from WooCommerce | Shipping: %s | ShipCost: %.2f | IP: %s", shippingMethod, shippingCost, wo.CustomerIP),
				GatewayType: "woocommerce",
				CreatedAt:   createdAt,
				UpdatedAt:   createdAt,
			}
			_ = db.Create(&pt).Error

			// ثبت لاگ مرتبط با درگاه (حاوی IP مشتری)
			if strings.TrimSpace(wo.CustomerIP) != "" {
				_ = db.Create(&models.PaymentGatewayLog{
					GatewayID: gateway.ID,
					Level:     "info",
					Message:   fmt.Sprintf("Imported order %s", trk),
					Details:   "Order imported from WooCommerce",
					IPAddress: wo.CustomerIP,
					UserAgent: "GoNuxt-WP-Migrator",
					CreatedAt: createdAt,
				}).Error
			}
			// به‌روزرسانی فیلدهای ساختاریافته ارسال در جدول orders
			upd := map[string]interface{}{
				"shipping_method": shippingMethodCode,
				"shipping_cost":   shippingCost,
			}
			if strings.TrimSpace(wo.CustomerIP) != "" {
				upd["customer_ip"] = wo.CustomerIP
			} else if onLog != nil {
				onLog("error", fmt.Sprintf("order %d: missing customer_ip — left empty", wo.ID))
			}
			if isPaid && strings.ToLower(strings.TrimSpace(mappedStatus)) == "completed" && strings.TrimSpace(shippingTracking) == "" {
				if onLog != nil {
					onLog("error", fmt.Sprintf("order %d: completed but missing shipping tracking code", wo.ID))
				}
			} else if strings.TrimSpace(shippingTracking) != "" {
				upd["shipping_tracking_code"] = shippingTracking
			}
			_ = db.Model(&models.Order{}).Where("id = ?", o.ID).Updates(upd).Error

			imported++
			if onLog != nil {
				onLog("success", fmt.Sprintf("سفارش %d ثبت شد برای کاربر %s", wo.ID, user.Username))
			}
			// موفقیت: نیازی به RollbackTo نیست
		}
		page++
	}
	return imported, skipped, failed, nil
}

// -----------------------------
// WooCommerce Attributes (Import)
// -----------------------------

type wcAttribute struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	OrderBy     string `json:"order_by"`
	HasArchives bool   `json:"has_archives"`
}

type wcAttributeTerm struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	MenuOrder   int    `json:"menu_order"`
	Count       int    `json:"count"`
}

func (s *WordPressMigrationService) FetchWCAttributes(ctx context.Context, siteURL, ck, cs string) ([]wcAttribute, error) {
	base := normalizeSiteURL(siteURL)
	client := &http.Client{}
	versions := []string{"v3", "v2", "v1"}
	all := make([]wcAttribute, 0, 64)
	for _, ver := range versions {
		// درخواست با per_page بسیار بزرگ تا همه ویژگی‌ها در یک پاسخ برگردند
		url := fmt.Sprintf("%s/wp-json/wc/%s/products/attributes?per_page=%d", base, ver, 10000000)
		if ck != "" && cs != "" {
			url += fmt.Sprintf("&consumer_key=%s&consumer_secret=%s", ck, cs)
		}
		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(url)
		if err != nil {
			continue
		}
		req, _ := http.NewRequestWithContext(ctx, "GET", validatedURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")
		resp, err := client.Do(req)
		if err != nil || resp == nil {
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			var arr []wcAttribute
			if err := json.Unmarshal(body, &arr); err != nil {
				return nil, err
			}
			all = append(all, arr...)
			if len(all) > 0 {
				return all, nil
			}
		}
	}
	return all, nil
}

func (s *WordPressMigrationService) FetchWCAttributeTerms(ctx context.Context, siteURL, ck, cs string, attrID int, page, perPage int) ([]wcAttributeTerm, error) {
	base := normalizeSiteURL(siteURL)
	client := &http.Client{}
	versions := []string{"v3", "v2", "v1"}
	for _, ver := range versions {
		url := fmt.Sprintf("%s/wp-json/wc/%s/products/attributes/%d/terms?per_page=%d&page=%d", base, ver, attrID, perPage, page)
		if ck != "" && cs != "" {
			url += fmt.Sprintf("&consumer_key=%s&consumer_secret=%s", ck, cs)
		}
		// اعتبارسنجی URL برای جلوگیری از SSRF
		validatedURL, err := validateURLForNetworkRequest(url)
		if err != nil {
			continue
		}
		req, _ := http.NewRequestWithContext(ctx, "GET", validatedURL, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "GoNuxt-WC-Migrator/1.0")
		resp, err := client.Do(req)
		if err == nil && resp != nil {
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				var arr []wcAttributeTerm
				if err := json.Unmarshal(body, &arr); err != nil {
					return nil, err
				}
				return arr, nil
			}
		}
	}
	return []wcAttributeTerm{}, nil
}

// ImportWCAttributes: ویژگی‌های ووکامرس را چک می‌کند (بدون ایجاد یا تغییر ساختار)
// این تابع فقط چک می‌کند که ویژگی‌ها موجود هستند و ساختار را تغییر نمی‌دهد
func (s *WordPressMigrationService) ImportWCAttributes(
	uCtx context.Context,
	uow unitofwork.UnitOfWork,
	siteURL, ck, cs string,
	logs *[]string,
	onLog func(logType, message string),
) (importedAttrs, importedValues, skipped, failed int, err error) {
	db := uow.DB()
	attrs, ferr := s.FetchWCAttributes(uCtx, siteURL, ck, cs)
	if ferr != nil {
		return 0, 0, 0, 0, ferr
	}
	if onLog != nil {
		onLog("info", fmt.Sprintf("یافت شد %d ویژگی از ووکامرس", len(attrs)))
	}
	for _, a := range attrs {
		select {
		case <-uCtx.Done():
			return importedAttrs, importedValues, skipped, failed, uCtx.Err()
		default:
		}
		code := strings.TrimSpace(a.Slug)
		if code == "" {
			code = slugifyTitle(a.Name)
		}
		// چک کردن وجود ویژگی - فقط چک کردن، بدون ایجاد یا تغییر
		var attr models.Attribute
		if err := db.Where("name = ? OR display_name = ? OR code = ?", a.Name, a.Name, code).First(&attr).Error; err == nil && attr.ID != 0 {
			// ویژگی قبلاً وجود دارد
			skipped++
			if onLog != nil {
				onLog("info", fmt.Sprintf("ویژگی %s قبلاً وجود دارد (اسکیپ)", a.Name))
			}
		} else {
			// ویژگی وجود ندارد - فقط شمارش می‌کنیم، ایجاد نمی‌کنیم
			failed++
			if onLog != nil {
				onLog("info", fmt.Sprintf("ویژگی %s در سیستم مقصد وجود ندارد (رد)", a.Name))
			}
		}
	}
	return importedAttrs, importedValues, skipped, failed, nil
}

// baseUploadsDir مسیر پایه‌ی ذخیره‌سازی فایل‌های آپلودی
func baseUploadsDir() string {
	if v := os.Getenv("UPLOADS_DIR"); strings.TrimSpace(v) != "" {
		return filepath.Clean(v)
	}
	// پیش‌فرض: ../public/uploads
	return filepath.Clean(filepath.Join("..", "public", "uploads"))
}

func ensureDir(dir string) error { return os.MkdirAll(dir, 0o755) }

// مسیرهای مطابق ساختار کاربر
func productImagesBaseDir() string {
	return filepath.Join(baseUploadsDir(), "media", "products", "images")
}

// getProductImageDir مسیر مناسب برای ذخیره عکس محصول بر اساس سال/ماه و تعداد فایل‌ها را برمی‌گرداند
func getProductImageDir() (string, error) {
	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")

	baseDir := productImagesBaseDir()

	// ابتدا پوشه پایه را ایجاد کن
	if err := ensureDir(baseDir); err != nil {
		return "", fmt.Errorf("خطا در ایجاد پوشه پایه: %v", err)
	}

	yearMonthDir := filepath.Join(baseDir, year, month)

	// اگر پوشه سال/ماه وجود ندارد، آن را ایجاد کن
	if err := ensureDir(yearMonthDir); err != nil {
		return "", fmt.Errorf("خطا در ایجاد پوشه سال/ماه: %v", err)
	}

	// شمارش فایل‌های موجود در پوشه سال/ماه
	files, err := os.ReadDir(yearMonthDir)
	if err != nil {
		// اگر خطا در خواندن پوشه رخ داد، حداقل پوشه سال/ماه را برگردان
		log.Printf("خطا در خواندن پوشه سال/ماه: %v", err)
		return yearMonthDir, nil
	}

	// شمارش کل تصاویر موجود در پوشه سال/ماه (شامل پوشه‌های batch)
	totalImageCount := 0

	// ابتدا تصاویر مستقیم در پوشه سال/ماه را بشمار
	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".webp" || ext == ".svg" || ext == ".avif" {
				totalImageCount++
			}
		}
	}

	// سپس تصاویر موجود در پوشه‌های batch را بشمار
	for _, file := range files {
		if file.IsDir() && strings.HasPrefix(file.Name(), "batch_") {
			batchDir := filepath.Join(yearMonthDir, file.Name())
			batchFiles, err := os.ReadDir(batchDir)
			if err != nil {
				log.Printf("خطا در خواندن پوشه %s: %v", file.Name(), err)
				continue
			}

			for _, batchFile := range batchFiles {
				if !batchFile.IsDir() {
					ext := strings.ToLower(filepath.Ext(batchFile.Name()))
					if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".webp" || ext == ".svg" || ext == ".avif" {
						totalImageCount++
					}
				}
			}
		}
	}

	// محاسبه شماره پوشه batch مناسب
	// تصاویر 1-500 در batch_1، 501-1000 در batch_2، و الی آخر
	batchNumber := (totalImageCount / 500) + 1
	folderName := fmt.Sprintf("batch_%d", batchNumber)
	batchDir := filepath.Join(yearMonthDir, folderName)

	// ایجاد پوشه batch اگر وجود ندارد
	if err := ensureDir(batchDir); err != nil {
		log.Printf("خطا در ایجاد پوشه batch_%d: %v", batchNumber, err)
		return yearMonthDir, nil
	}

	return batchDir, nil
}
func productVideosBaseDir() string {
	return filepath.Join(baseUploadsDir(), "media", "products", "videos")
}

func guessExtFromURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return ".jpg"
	}
	name := filepath.Base(u.Path)
	ext := filepath.Ext(name)
	if ext == "" {
		return ".jpg"
	}
	return strings.ToLower(ext)
}

func downloadAndSaveImage(ctx context.Context, src string, destDir string, baseName string, createVariants bool) (string, error) {
	if err := ensureDir(destDir); err != nil {
		return "", err
	}
	// نام فایل بر اساس hash از URL تا یکتا شود
	sum := md5.Sum([]byte(src))
	fn := fmt.Sprintf("%s-%x%s", baseName, sum[:4], guessExtFromURL(src))
	destPath := filepath.Join(destDir, fn)

	// اعتبارسنجی URL برای جلوگیری از SSRF
	validatedURL, err := validateURLForNetworkRequest(src)
	if err != nil {
		return "", fmt.Errorf("URL نامعتبر: %v", err)
	}

	// دانلود
	req, _ := http.NewRequestWithContext(ctx, "GET", validatedURL, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("http %d", resp.StatusCode)
	}
	// خواندن کامل در حافظه برای decode و فشرده‌سازی
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// تلاش برای decode تصویر
	var img image.Image
	var decErr error
	// تشخیص بر اساس پسوند و تلاش جایگزین
	switch strings.ToLower(filepath.Ext(fn)) {
	case ".webp":
		img, decErr = webp.Decode(bytes.NewReader(data))
	default:
		img, _, decErr = image.Decode(bytes.NewReader(data))
	}
	// اگر decode نشد، همان فایل خام ذخیره شود
	if decErr != nil {
		if err := os.WriteFile(destPath, data, 0o644); err != nil {
			return "", err
		}
	} else {
		// فشرده‌سازی هوشمند مانند کتابخانه رسانه
		// برای سازگاری، خروجی را webp/jpeg بر اساس SmartOptions انتخاب می‌کنیم
		// detect mime on first up to 512 bytes
		sniff := 512
		if len(data) < sniff {
			sniff = len(data)
		}
		filetype := http.DetectContentType(data[:sniff])
		res, err := utils.SmartCompress(img, filetype, utils.SmartOptions{
			OutputFormat: "", // auto decide
			OriginalSize: len(data),
			DesiredRatio: 0.6, // کاهش پیش‌فرض شبیه آپلودر
			MinQuality:   20,
		})
		if err != nil {
			// fallback: ذخیره raw
			if err := os.WriteFile(destPath, data, 0o644); err != nil {
				return "", err
			}
		} else {
			// اگر فرمت عوض شد، پسوند را هم هماهنگ کنیم
			outExt := map[string]string{"jpeg": ".jpg", "jpg": ".jpg", "png": ".png", "webp": ".webp"}[res.Format]
			if outExt == "" {
				outExt = strings.ToLower(filepath.Ext(fn))
			}
			base := strings.TrimSuffix(destPath, filepath.Ext(destPath))
			destPath = base + outExt
			// اطمینان از وجود دایرکتوری
			if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
				return "", err
			}
			if err := os.WriteFile(destPath, res.Data.Bytes(), 0o644); err != nil {
				return "", err
			}
		}

		// فقط در صورتی که createVariants=true باشد واریانت‌ها را ایجاد کن
		if createVariants {
			// تولید واریانت‌ها برای محصولات و دسته‌بندی‌ها (سایزبندی مشابه MediaHandler) — هم‌زمان (sync)
			f, err := os.Open(destPath)
			if err == nil {
				img2, _, err2 := image.Decode(f)
				_ = f.Close()
				if err2 == nil {
					sizes := map[string]struct{ w, h uint }{
						"thumbnail": {150, 150},
						"small":     {400, 400},
						"medium":    {800, 800},
						"large":     {1600, 1600},
					}
					for name, dim := range sizes {
						variant := resize.Resize(dim.w, dim.h, img2, resize.Lanczos3)
						vPath := strings.TrimSuffix(destPath, filepath.Ext(destPath)) + "_" + name + filepath.Ext(destPath)
						if vf, err := os.Create(vPath); err == nil {
							switch strings.ToLower(filepath.Ext(vPath)) {
							case ".png":
								_ = png.Encode(vf, variant)
							case ".webp":
								_ = webp.Encode(vf, variant, &webp.Options{Quality: 85})
							default:
								_ = jpeg.Encode(vf, variant, &jpeg.Options{Quality: 85})
							}
							_ = vf.Close()
						}
					}
				}
			}
		}
	}

	// مسیر وب برای ذخیره در DB (نسبت به public)
	rel := strings.TrimPrefix(destPath, filepath.Clean(filepath.Join("..", "public")))
	rel = filepath.ToSlash(rel)
	if !strings.HasPrefix(rel, "/") {
		rel = "/" + rel
	}
	return rel, nil
}

// importProductVideosFromMeta استخراج ویدیوهای محصول بر اساس meta_data
func (s *WordPressMigrationService) importProductVideosFromMeta(uow unitofwork.UnitOfWork, ctx context.Context, productID int, slug string, p wcProduct) error {
	// حذف ویدیوهای قبلی
	if err := uow.DB().WithContext(ctx).Where("product_id = ?", productID).Delete(&models.ProductVideo{}).Error; err != nil {
		return err
	}
	urls := make([]string, 0, 4)
	for _, md := range p.MetaData {
		k := strings.ToLower(md.Key)
		if strings.Contains(k, "video") {
			switch v := md.Value.(type) {
			case string:
				s := strings.TrimSpace(v)
				if s != "" {
					urls = append(urls, s)
				}
			case []interface{}:
				for _, it := range v {
					if ss, ok := it.(string); ok {
						ss = strings.TrimSpace(ss)
						if ss != "" {
							urls = append(urls, ss)
						}
					}
				}
			case map[string]interface{}:
				if u, ok := v["url"].(string); ok {
					u = strings.TrimSpace(u)
					if u != "" {
						urls = append(urls, u)
					}
				}
			}
		}
	}
	// ذخیره محلی ویدیوها در /public/uploads/media/products/videos/<slug>/
	vidDir := filepath.Join(productVideosBaseDir(), slug)
	for idx, u := range urls {
		saved := u
		if strings.HasPrefix(u, "http://") || strings.HasPrefix(u, "https://") {
			if pth, err := downloadAndSaveImage(ctx, u, vidDir, fmt.Sprintf("video-%d", idx), false); err == nil {
				saved = pth
			}
		}
		pv := models.ProductVideo{ProductID: productID, Title: fmt.Sprintf("Video %d", idx+1), VideoURL: saved, SortOrder: idx}
		if err := uow.DB().WithContext(ctx).Create(&pv).Error; err != nil {
			return err
		}
	}
	return nil
}

// ProcessHTML processes HTML content and downloads/replaces images and videos
func (s *WordPressMigrationService) ProcessHTML(ctx context.Context, html, siteURL, _ string, isFullDescription bool) (string, error) {
	if html == "" {
		return "", nil
	}

	// دامنه سایت وردپرس را نرمالایز می‌کنیم
	baseURL := normalizeSiteURL(siteURL)

	// با استفاده از regex تصاویر را پیدا می‌کنیم
	imgRegex := regexp.MustCompile(`<img[^>]*src=["']([^"']+)["'][^>]*>`)
	videoRegex := regexp.MustCompile(`<video[^>]*src=["']([^"']+)["'][^>]*>|<iframe[^>]*src=["']([^"']+)["'][^>]*>`)
	linkRegex := regexp.MustCompile(`<a[^>]*href=["']([^"']+)["'][^>]*>`)

	// دایرکتوری مقصد برای تصاویر توضیحات - استفاده از همان ساختار محصولات
	destImgDir, err := getProductImageDir()
	if err != nil {
		log.Printf("خطا در تعیین مسیر ذخیره عکس توضیحات: %v", err)
		// fallback به مسیر سال/ماه
		now := time.Now()
		year := now.Format("2006")
		month := now.Format("01")
		destImgDir = filepath.Join(productImagesBaseDir(), year, month)
		if ensureDirErr := ensureDir(destImgDir); ensureDirErr != nil {
			destImgDir = productImagesBaseDir()
		}
	}

	// پردازش تصاویر
	processedHTML := imgRegex.ReplaceAllStringFunc(html, func(imgTag string) string {
		matches := imgRegex.FindStringSubmatch(imgTag)
		if len(matches) < 2 {
			return imgTag
		}

		imgSrc := matches[1]
		// اگر آدرس نسبی است آن را به مطلق تبدیل می‌کنیم
		if !strings.HasPrefix(imgSrc, "http://") && !strings.HasPrefix(imgSrc, "https://") {
			imgSrc = baseURL + "/" + strings.TrimPrefix(imgSrc, "/")
		}

		// دانلود و ذخیره تصویر
		prefix := "desc"
		if !isFullDescription {
			prefix = "short-desc"
		}
		// پارامتر آخر را false قرار می‌دهیم تا واریانت ساخته نشود
		savedPath, err := downloadAndSaveImage(ctx, imgSrc, destImgDir, fmt.Sprintf("%s-img-%x", prefix, md5.Sum([]byte(imgSrc))), false)
		if err != nil {
			// در صورت خطا، از همان آدرس اصلی استفاده می‌کنیم
			return imgTag
		}

		// جایگزینی آدرس تصویر در تگ img
		return strings.Replace(imgTag, matches[1], savedPath, 1)
	})

	// پردازش ویدیوها
	processedHTML = videoRegex.ReplaceAllStringFunc(processedHTML, func(videoTag string) string {
		matches := videoRegex.FindStringSubmatch(videoTag)
		if len(matches) < 3 {
			return videoTag
		}

		videoSrc := ""
		if matches[1] != "" {
			videoSrc = matches[1] // از تگ video
		} else if matches[2] != "" {
			videoSrc = matches[2] // از تگ iframe
		}

		if videoSrc == "" {
			return videoTag
		}

		// اگر آدرس نسبی است آن را به مطلق تبدیل می‌کنیم
		if !strings.HasPrefix(videoSrc, "http://") && !strings.HasPrefix(videoSrc, "https://") {
			videoSrc = baseURL + "/" + strings.TrimPrefix(videoSrc, "/")
		}

		// برای ویدیوهای YouTube و Vimeo نیازی به دانلود نیست
		if strings.Contains(videoSrc, "youtube.com") || strings.Contains(videoSrc, "youtu.be") ||
			strings.Contains(videoSrc, "vimeo.com") {
			return videoTag
		}

		// دانلود و ذخیره ویدیو در پوشه ویدیو
		vidDir := productVideosBaseDir()
		prefix := "desc"
		if !isFullDescription {
			prefix = "short-desc"
		}
		// برای ویدیوها نیز واریانت ساخته نشود
		savedPath, err := downloadAndSaveImage(ctx, videoSrc, vidDir, fmt.Sprintf("%s-video-%x", prefix, md5.Sum([]byte(videoSrc))), false)
		if err != nil {
			// در صورت خطا، از همان آدرس اصلی استفاده می‌کنیم
			return videoTag
		}

		// جایگزینی آدرس ویدیو در تگ video یا iframe
		if matches[1] != "" {
			return strings.Replace(videoTag, matches[1], savedPath, 1)
		} else {
			return strings.Replace(videoTag, matches[2], savedPath, 1)
		}
	})

	// پردازش لینک‌ها
	processedHTML = linkRegex.ReplaceAllStringFunc(processedHTML, func(linkTag string) string {
		matches := linkRegex.FindStringSubmatch(linkTag)
		if len(matches) < 2 {
			return linkTag
		}

		href := matches[1]

		// اگر لینک مربوط به دامنه خود سایت است، آن را به یک لینک نسبی تبدیل می‌کنیم
		// یا می‌توانیم یک قانون ریدایرکت ایجاد کنیم (اختیاری)
		if strings.HasPrefix(href, baseURL) {
			// می‌توانیم اینجا منطق ریدایرکت اضافه کنیم
			return linkTag
		}

		return linkTag
	})

	return processedHTML, nil
}

// handleTransactionError مدیریت خطاهای transaction و rollback
func (s *WordPressMigrationService) handleTransactionError(db *gorm.DB, orderID int, err error, onLog func(string, string)) bool {
	if err == nil {
		return false
	}

	// بررسی نوع خطا
	if strings.Contains(err.Error(), "current transaction is aborted") ||
		strings.Contains(err.Error(), "SQLSTATE 25P02") ||
		strings.Contains(err.Error(), "duplicate key") ||
		strings.Contains(err.Error(), "SQLSTATE 23505") {

		if onLog != nil {
			onLog("error", fmt.Sprintf("order %d: transaction error detected: %v", orderID, err))
		}

		// تلاش برای rollback
		if rollbackErr := db.RollbackTo("sp_import_order").Error; rollbackErr != nil {
			if onLog != nil {
				onLog("error", fmt.Sprintf("order %d: failed to rollback: %v", orderID, rollbackErr))
			}
		}

		return true
	}

	return false
}
