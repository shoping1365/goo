package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"

	"gorm.io/gorm"
)

// DigikalaCategory نمایانگر دسته‌بندی دیجی‌کالا
type DigikalaCategory struct {
	ID       int                `json:"id"`
	Title    string             `json:"title"`
	Slug     string             `json:"slug"`
	Image    string             `json:"image"`
	ParentID *int               `json:"parent_id"`
	Children []DigikalaCategory `json:"children"`
}

// DigikalaCategoryResponse پاسخ API دیجی‌کالا
type DigikalaCategoryResponse struct {
	Status string             `json:"status"`
	Data   []DigikalaCategory `json:"data"`
}

// DigikalaImportService سرویس import از دیجی‌کالا
type DigikalaImportService struct {
	DB           *gorm.DB
	CategoryRepo repository.CategoryRepositoryInterface
	httpClient   *http.Client
	baseURL      string
}

// NewDigikalaImportService ایجاد instance جدید
func NewDigikalaImportService(
	db *gorm.DB,
	categoryRepo repository.CategoryRepositoryInterface,
) *DigikalaImportService {
	return &DigikalaImportService{
		DB:           db,
		CategoryRepo: categoryRepo,
		httpClient:   &http.Client{Timeout: 30 * time.Second},
		baseURL:      "https://api.digikala.com/v1",
	}
}

// FetchCategories دریافت دسته‌بندی‌ها از دیجی‌کالا
func (s *DigikalaImportService) FetchCategories(categoryID string) ([]DigikalaCategory, error) {
	url := fmt.Sprintf("%s/categories/", s.baseURL)
	if categoryID != "" {
		url = fmt.Sprintf("%s/categories/%s/", s.baseURL, categoryID)
	}

	log.Printf("[DigikalaImport] Fetching categories from: %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("خطا در ایجاد درخواست: %v", err)
	}

	req.Header.Set("User-Agent", "GoNuxt-Digikala-Importer/1.0")
	req.Header.Set("Accept", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("خطای HTTP: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("خطا در خواندن پاسخ: %v", err)
	}

	var response DigikalaCategoryResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, fmt.Errorf("خطا در parse کردن JSON: %v", err)
	}

	return response.Data, nil
}

// ImportCategories import دسته‌بندی‌های انتخابی از دیجی‌کالا
func (s *DigikalaImportService) ImportCategories(
	ctx context.Context,
	categoryIDs []string,
	parentCategoryID *uint,
) (imported int, failed int, err error) {
	log.Printf("[DigikalaImport] شروع import %d دسته‌بندی", len(categoryIDs))

	for _, categoryID := range categoryIDs {
		categories, err := s.FetchCategories(categoryID)
		if err != nil {
			log.Printf("[DigikalaImport] خطا در دریافت دسته‌بندی %s: %v", categoryID, err)
			failed++
			continue
		}

		for _, dkCat := range categories {
			if err := s.importSingleCategory(ctx, dkCat, parentCategoryID); err != nil {
				log.Printf("[DigikalaImport] خطا در import دسته‌بندی %s: %v", dkCat.Title, err)
				failed++
			} else {
				imported++
				log.Printf("[DigikalaImport] دسته‌بندی %s با موفقیت import شد", dkCat.Title)
			}
		}
	}

	return imported, failed, nil
}

// importSingleCategory import یک دسته‌بندی منفرد
func (s *DigikalaImportService) importSingleCategory(
	ctx context.Context,
	dkCat DigikalaCategory,
	parentID *uint,
) error {
	// بررسی وجود دسته‌بندی با همین slug
	existing, err := s.CategoryRepo.GetBySlug(ctx, dkCat.Slug)
	if err == nil && existing != nil {
		log.Printf("[DigikalaImport] دسته‌بندی %s قبلاً موجود است", dkCat.Slug)
		return nil
	}

	// ایجاد دسته‌بندی جدید
	category := models.Category{
		Name:            dkCat.Title,
		Slug:            dkCat.Slug,
		Description:     fmt.Sprintf("Import شده از دیجی‌کالا - ID: %d", dkCat.ID),
		ParentID:        parentID,
		ImageURL:        dkCat.Image,
		Published:       true,
		Order:           0,
		MetaTitle:       dkCat.Title,
		MetaDescription: fmt.Sprintf("دسته‌بندی %s import شده از دیجی‌کالا", dkCat.Title),
	}

	if err := s.CategoryRepo.Create(ctx, &category); err != nil {
		return fmt.Errorf("خطا در ایجاد دسته‌بندی: %v", err)
	}

	// Import کردن زیردسته‌ها
	for _, child := range dkCat.Children {
		if err := s.importSingleCategory(ctx, child, &category.ID); err != nil {
			log.Printf("[DigikalaImport] خطا در import زیردسته %s: %v", child.Title, err)
		}
	}

	return nil
}

// GetAvailableCategories دریافت لیست دسته‌بندی‌های موجود در دیجی‌کالا
func (s *DigikalaImportService) GetAvailableCategories() ([]DigikalaCategory, error) {
	return s.FetchCategories("")
}

// SearchCategories جستجو در دسته‌بندی‌های دیجی‌کالا
func (s *DigikalaImportService) SearchCategories(query string) ([]DigikalaCategory, error) {
	url := fmt.Sprintf("%s/search/categories/?q=%s", s.baseURL, query)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("خطا در ایجاد درخواست جستجو: %v", err)
	}

	req.Header.Set("User-Agent", "GoNuxt-Digikala-Importer/1.0")
	req.Header.Set("Accept", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست جستجو: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("خطای HTTP در جستجو: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("خطا در خواندن پاسخ جستجو: %v", err)
	}

	var response DigikalaCategoryResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, fmt.Errorf("خطا در parse کردن JSON جستجو: %v", err)
	}

	return response.Data, nil
}

// ValidateDigikalaURL اعتبارسنجی URL دیجی‌کالا و استخراج اطلاعات دسته‌بندی
func (s *DigikalaImportService) ValidateDigikalaURL(url string) (isValid bool, categoryTitle string, categoryID string) {
	// بررسی فرمت کلی URL - قبول هر دو فرمت:
	// 1. https://www.digikala.com/search/category-SLUG/
	// 2. https://www.digikala.com/category-SLUG/

	if !strings.Contains(url, "digikala.com") {
		return false, "", ""
	}

	// استخراج شناسه دسته‌بندی از URL - با دو pattern مختلف
	var categorySlug string

	// Pattern 1: /search/category-SLUG/
	re1 := regexp.MustCompile(`/search/category-([^/]+)/?`)
	matches1 := re1.FindStringSubmatch(url)
	if len(matches1) >= 2 {
		categorySlug = matches1[1]
	} else {
		// Pattern 2: /category-SLUG/ (بدون search)
		re2 := regexp.MustCompile(`/category-([^/]+)/?`)
		matches2 := re2.FindStringSubmatch(url)
		if len(matches2) >= 2 {
			categorySlug = matches2[1]
		}
	}

	// اگر هیچ pattern ای match نشد
	if categorySlug == "" {
		// آخرین تلاش: اگر category- در URL هست
		if strings.Contains(url, "category-") {
			parts := strings.Split(url, "category-")
			if len(parts) >= 2 {
				slug := strings.TrimSuffix(strings.Split(parts[1], "/")[0], "/")
				if slug != "" {
					categorySlug = slug
				}
			}
		}
	}

	// اگر هنوز slug پیدا نشد، URL نامعتبر است
	if categorySlug == "" {
		return false, "", ""
	}

	// تلاش برای دریافت اطلاعات دسته‌بندی از API دیجی‌کالا
	categories, err := s.FetchCategories("")
	if err != nil {
		// در صورت خطا در API، حداقل فرمت URL را بررسی می‌کنیم
		return true, categorySlug, categorySlug
	}

	// جستجو در دسته‌بندی‌ها برای یافتن عنوان
	categoryTitle = s.findCategoryTitle(categories, categorySlug)
	if categoryTitle == "" {
		categoryTitle = categorySlug
	}

	return true, categoryTitle, categorySlug
}

// findCategoryTitle جستجوی عنوان دسته‌بندی در لیست دسته‌بندی‌ها
func (s *DigikalaImportService) findCategoryTitle(categories []DigikalaCategory, slug string) string {
	for _, cat := range categories {
		if cat.Slug == slug {
			return cat.Title
		}
		// جستجو در زیردسته‌ها
		if title := s.findCategoryTitle(cat.Children, slug); title != "" {
			return title
		}
	}
	return ""
}

// SetBaseURL تنظیم آدرس پایه API (برای تست)
func (s *DigikalaImportService) SetBaseURL(url string) {
	s.baseURL = url
}

// SetHTTPClient تنظیم کلاینت HTTP (برای تست)
func (s *DigikalaImportService) SetHTTPClient(client *http.Client) {
	s.httpClient = client
}
