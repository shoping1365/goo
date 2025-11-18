package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
)

// DigikalaProduct ساختار محصول دیجی‌کالا
type DigikalaProduct struct {
	ID             int               `json:"id"`
	Title          string            `json:"title"`
	Slug           string            `json:"slug"`
	Price          DigikalaPrice     `json:"price"`
	Images         []DigikalaImage   `json:"images"`
	Specifications []DigikalaSpec    `json:"specifications"`
	Category       DigikalaCategory  `json:"category"`
	Brand          DigikalaBrand     `json:"brand"`
	Description    string            `json:"description"`
	Rating         float64           `json:"rating"`
	Colors         []DigikalaColor   `json:"colors"`
	Variants       []DigikalaVariant `json:"variants"`
	Status         string            `json:"status"`
}

type DigikalaPrice struct {
	Selling         int `json:"selling_price"`
	RRPPrice        int `json:"rrp_price"`
	DiscountPercent int `json:"discount_percent"`
}

type DigikalaImage struct {
	URL       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}

type DigikalaSpec struct {
	Title  string              `json:"title"`
	Values []DigikalaSpecValue `json:"values"`
}

type DigikalaSpecValue struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type DigikalaBrand struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type DigikalaColor struct {
	Title string `json:"title"`
	Hex   string `json:"hex"`
}

type DigikalaVariant struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

// DigikalaProductAPIResponse پاسخ API محصول دیجی‌کالا
type DigikalaProductAPIResponse struct {
	Status string `json:"status"`
	Data   struct {
		Product DigikalaProduct `json:"product"`
	} `json:"data"`
}

// DigikalaCategoryProductsResponse پاسخ لیست محصولات یک دسته‌بندی
type DigikalaCategoryProductsResponse struct {
	Status string `json:"status"`
	Data   struct {
		Products []struct {
			ID    int    `json:"id"`
			Title string `json:"title_fa"`
			URL   string `json:"url"`
		} `json:"products"`
		Pagination struct {
			Total int `json:"total"`
			Page  int `json:"page"`
		} `json:"pagination"`
	} `json:"data"`
}

// DigikalaProductScraperService سرویس اسکرپ محصولات از دیجی‌کالا
type DigikalaProductScraperService struct {
	DB               *gorm.DB
	ProductRepo      repository.ProductRepositoryInterface
	CategoryRepo     repository.CategoryRepositoryInterface
	ImportLogRepo    repository.DigikalaImportLogRepositoryInterface
	httpClient       *http.Client
	baseURL          string
	apiURL           string
	importID         uint
	targetCategoryID *uint
	settings         ImportSettings
}

// ImportSettings تنظیمات import
type ImportSettings struct {
	ItemsPerMinute int
	MaxProducts    int
	SkipExisting   bool
	ImportImages   bool
	ImportSpecs    bool
	Delay          time.Duration // تاخیر بین درخواست‌ها برای جلوگیری از block شدن
}

// NewDigikalaProductScraperService ایجاد instance جدید
func NewDigikalaProductScraperService(
	db *gorm.DB,
	productRepo repository.ProductRepositoryInterface,
	categoryRepo repository.CategoryRepositoryInterface,
	importLogRepo repository.DigikalaImportLogRepositoryInterface,
) *DigikalaProductScraperService {
	return &DigikalaProductScraperService{
		DB:            db,
		ProductRepo:   productRepo,
		CategoryRepo:  categoryRepo,
		ImportLogRepo: importLogRepo,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				IdleConnTimeout:     90 * time.Second,
			},
		},
		baseURL: "https://www.digikala.com",
		apiURL:  "https://api.digikala.com/v2",
		settings: ImportSettings{
			ItemsPerMinute: 10,
			MaxProducts:    100,
			SkipExisting:   true,
			ImportImages:   true,
			ImportSpecs:    true,
			Delay:          3 * time.Second, // 3 ثانیه تاخیر بین درخواست‌ها
		},
	}
}

// SetImportID تنظیم شناسه import
func (s *DigikalaProductScraperService) SetImportID(importID uint) {
	s.importID = importID
}

// SetTargetCategory تنظیم دسته‌بندی مقصد
func (s *DigikalaProductScraperService) SetTargetCategory(categoryID *uint) {
	s.targetCategoryID = categoryID
}

// SetSettings تنظیم پارامترهای import
func (s *DigikalaProductScraperService) SetSettings(settings ImportSettings) {
	s.settings = settings
	if s.settings.Delay == 0 {
		s.settings.Delay = 3 * time.Second
	}
}

// addLog اضافه کردن لاگ
func (s *DigikalaProductScraperService) addLog(level, message string) {
	if s.importID > 0 && s.ImportLogRepo != nil {
		s.ImportLogRepo.Create(&models.DigikalaImportLog{
			ImportID: s.importID,
			Level:    level,
			Message:  message,
		})
	}
	log.Printf("[DigikalaProductScraper] [%s] %s", level, message)
}

// FetchCategoryProducts دریافت لیست محصولات یک دسته‌بندی
func (s *DigikalaProductScraperService) FetchCategoryProducts(categoryURL string, maxProducts int) ([]string, error) {
	s.addLog("INFO", fmt.Sprintf("شروع دریافت لیست محصولات از: %s", categoryURL))

	// استخراج slug دسته‌بندی از URL
	// مثال: https://www.digikala.com/search/category-mobile-phone/
	re := regexp.MustCompile(`/search/category-([^/]+)/?`)
	matches := re.FindStringSubmatch(categoryURL)
	if len(matches) < 2 {
		return nil, fmt.Errorf("فرمت URL نامعتبر است")
	}

	categorySlug := matches[1]
	var productURLs []string
	page := 1

	for len(productURLs) < maxProducts {
		// ساخت URL صفحه فهرست محصولات
		listURL := fmt.Sprintf("%s/search/category-%s/?page=%d", s.baseURL, categorySlug, page)
		s.addLog("INFO", fmt.Sprintf("دریافت صفحه %d از فهرست محصولات", page))

		// ارسال درخواست با header های واقعی
		req, err := http.NewRequest("GET", listURL, nil)
		if err != nil {
			return nil, fmt.Errorf("خطا در ایجاد درخواست: %v", err)
		}

		s.setHumanLikeHeaders(req)

		resp, err := s.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("خطا در ارسال درخواست: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("خطای HTTP: %d", resp.StatusCode)
		}

		// Parse HTML با goquery
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		resp.Body.Close()

		if err != nil {
			return nil, fmt.Errorf("خطا در parse کردن HTML: %v", err)
		}

		// یافتن لینک محصولات
		foundInThisPage := 0
		doc.Find("a[href*='/product/']").Each(func(i int, sel *goquery.Selection) {
			if len(productURLs) >= maxProducts {
				return
			}
			href, exists := sel.Attr("href")
			if exists && strings.Contains(href, "/product/") {
				// ساخت URL کامل محصول
				if !strings.HasPrefix(href, "http") {
					href = s.baseURL + href
				}
				// جلوگیری از تکرار
				if !containsURL(productURLs, href) {
					productURLs = append(productURLs, href)
					foundInThisPage++
				}
			}
		})

		s.addLog("INFO", fmt.Sprintf("تعداد %d محصول در صفحه %d یافت شد", foundInThisPage, page))

		// اگر محصولی پیدا نشد، از loop خارج شو
		if foundInThisPage == 0 {
			break
		}

		// تاخیر بین صفحات
		time.Sleep(s.settings.Delay)
		page++

		// محدودیت تعداد صفحات (برای جلوگیری از infinite loop)
		if page > 50 {
			break
		}
	}

	// محدود کردن به maxProducts
	if len(productURLs) > maxProducts {
		productURLs = productURLs[:maxProducts]
	}

	s.addLog("INFO", fmt.Sprintf("در مجموع %d محصول یافت شد", len(productURLs)))
	return productURLs, nil
}

// ScrapeProduct اسکرپ یک محصول از دیجی‌کالا
func (s *DigikalaProductScraperService) ScrapeProduct(productURL string) (*DigikalaProduct, error) {
	s.addLog("INFO", fmt.Sprintf("شروع اسکرپ محصول: %s", productURL))

	// استخراج DK کد محصول از URL
	// مثال: https://www.digikala.com/product/dkp-123456/
	re := regexp.MustCompile(`/product/dkp-(\d+)`)
	matches := re.FindStringSubmatch(productURL)
	if len(matches) < 2 {
		return nil, fmt.Errorf("نتوانستیم کد محصول را از URL استخراج کنیم")
	}

	productID := matches[1]

	// ارسال درخواست به صفحه محصول
	req, err := http.NewRequest("GET", productURL, nil)
	if err != nil {
		return nil, fmt.Errorf("خطا در ایجاد درخواست: %v", err)
	}

	s.setHumanLikeHeaders(req)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("خطای HTTP: %d", resp.StatusCode)
	}

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("خطا در parse کردن HTML: %v", err)
	}

	product := &DigikalaProduct{
		ID: mustParseInt(productID),
	}

	// استخراج اطلاعات محصول از HTML
	s.extractProductInfo(doc, product)

	// تلاش برای دریافت اطلاعات از API (اگر موجود باشد)
	s.enrichProductFromAPI(product)

	s.addLog("INFO", fmt.Sprintf("محصول '%s' با موفقیت اسکرپ شد", product.Title))
	return product, nil
}

// extractProductInfo استخراج اطلاعات محصول از HTML
func (s *DigikalaProductScraperService) extractProductInfo(doc *goquery.Document, product *DigikalaProduct) {
	// عنوان محصول
	product.Title = strings.TrimSpace(doc.Find("h1.product-title, h1[data-cro-id='pdp-title']").First().Text())

	// قیمت (با پاکسازی کاراکترهای اضافی)
	priceText := doc.Find(".product-price, [data-testid='price-final']").First().Text()
	priceText = strings.ReplaceAll(priceText, "تومان", "")
	priceText = strings.ReplaceAll(priceText, "٬", "")
	priceText = strings.ReplaceAll(priceText, ",", "")
	priceText = strings.TrimSpace(priceText)
	product.Price.Selling = mustParseInt(priceText)

	// قیمت قبلی (در صورت تخفیف)
	oldPriceText := doc.Find(".product-price-before, [data-testid='price-no-discount']").First().Text()
	oldPriceText = strings.ReplaceAll(oldPriceText, "تومان", "")
	oldPriceText = strings.ReplaceAll(oldPriceText, "٬", "")
	oldPriceText = strings.ReplaceAll(oldPriceText, ",", "")
	oldPriceText = strings.TrimSpace(oldPriceText)
	product.Price.RRPPrice = mustParseInt(oldPriceText)

	// درصد تخفیف
	if product.Price.RRPPrice > 0 && product.Price.Selling > 0 {
		product.Price.DiscountPercent = int(float64(product.Price.RRPPrice-product.Price.Selling) / float64(product.Price.RRPPrice) * 100)
	}

	// تصاویر محصول
	doc.Find("img[data-testid='product-image'], .product-gallery img").Each(func(i int, sel *goquery.Selection) {
		imgURL, exists := sel.Attr("src")
		if !exists {
			imgURL, exists = sel.Attr("data-src")
		}
		if exists && imgURL != "" {
			// حذف query parameters اضافی
			imgURL = strings.Split(imgURL, "?")[0]
			product.Images = append(product.Images, DigikalaImage{
				URL:       imgURL,
				Thumbnail: imgURL,
			})
		}
	})

	// مشخصات فنی
	doc.Find(".product-specifications, [data-testid='specifications'] table tr").Each(func(i int, sel *goquery.Selection) {
		title := strings.TrimSpace(sel.Find("th, td:first-child").Text())
		value := strings.TrimSpace(sel.Find("td:last-child").Text())

		if title != "" && value != "" {
			// جستجو برای spec موجود با همین عنوان
			found := false
			for i := range product.Specifications {
				if product.Specifications[i].Title == title {
					product.Specifications[i].Values = append(product.Specifications[i].Values, DigikalaSpecValue{
						Title: title,
						Value: value,
					})
					found = true
					break
				}
			}

			if !found {
				product.Specifications = append(product.Specifications, DigikalaSpec{
					Title: title,
					Values: []DigikalaSpecValue{{
						Title: title,
						Value: value,
					}},
				})
			}
		}
	})

	// توضیحات محصول
	product.Description = strings.TrimSpace(doc.Find(".product-description, [data-testid='product-description']").First().Text())

	// برند
	brandText := doc.Find(".product-brand, [data-testid='brand-name']").First().Text()
	product.Brand.Title = strings.TrimSpace(brandText)
}

// enrichProductFromAPI تکمیل اطلاعات محصول از API
func (s *DigikalaProductScraperService) enrichProductFromAPI(product *DigikalaProduct) {
	// تلاش برای دریافت داده از API دیجی‌کالا
	apiURL := fmt.Sprintf("%s/product/%d/", s.apiURL, product.ID)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return
	}

	s.setHumanLikeHeaders(req)
	req.Header.Set("Accept", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var apiResp DigikalaProductAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return
	}

	// تکمیل داده‌های ناقص از API
	if product.Title == "" {
		product.Title = apiResp.Data.Product.Title
	}
	if product.Price.Selling == 0 {
		product.Price = apiResp.Data.Product.Price
	}
	if len(product.Images) == 0 {
		product.Images = apiResp.Data.Product.Images
	}
	if len(product.Specifications) == 0 {
		product.Specifications = apiResp.Data.Product.Specifications
	}
}

// setHumanLikeHeaders تنظیم header های شبیه مرورگر واقعی
func (s *DigikalaProductScraperService) setHumanLikeHeaders(req *http.Request) {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0",
	}

	req.Header.Set("User-Agent", userAgents[rand.Intn(len(userAgents))])
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "fa,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Referer", "https://www.digikala.com/")
}

// downloadImage دانلود تصویر و ذخیره در public/uploads/media
func (s *DigikalaProductScraperService) downloadImage(imageURL string) (string, error) {
	// ساخت hash از URL برای نام فایل یکتا
	// استفاده از SHA-256 به جای MD5 برای امنیت بیشتر
	hasher := sha256.New()
	hasher.Write([]byte(imageURL))
	hash := hex.EncodeToString(hasher.Sum(nil))

	// گرفتن پسوند فایل
	ext := filepath.Ext(imageURL)
	if ext == "" || len(ext) > 5 {
		ext = ".jpg" // پیش‌فرض
	}

	// نام فایل جدید
	filename := fmt.Sprintf("%s%s", hash, ext)
	uploadDir := "public/uploads/media"
	localPath := filepath.Join(uploadDir, filename)

	// ایجاد پوشه در صورت عدم وجود
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("خطا در ایجاد پوشه: %v", err)
	}

	// بررسی وجود فایل
	if _, err := os.Stat(localPath); err == nil {
		// فایل از قبل وجود دارد
		return "/uploads/media/" + filename, nil
	}

	// دانلود تصویر
	req, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		return "", fmt.Errorf("خطا در ایجاد request: %v", err)
	}

	s.setHumanLikeHeaders(req)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("خطا در دانلود تصویر: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("خطا در دانلود: status %d", resp.StatusCode)
	}

	// ذخیره فایل
	outFile, err := os.Create(localPath)
	if err != nil {
		return "", fmt.Errorf("خطا در ایجاد فایل: %v", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return "", fmt.Errorf("خطا در ذخیره فایل: %v", err)
	}

	return "/uploads/media/" + filename, nil
}

// SaveProduct ذخیره محصول در دیتابیس
func (s *DigikalaProductScraperService) SaveProduct(ctx context.Context, dkProduct *DigikalaProduct) (*models.Product, error) {
	s.addLog("INFO", fmt.Sprintf("شروع ذخیره محصول: %s", dkProduct.Title))

	// بررسی وجود محصول با همین slug
	slug := generateSlug(dkProduct.Title)
	existingProduct, _ := s.ProductRepo.GetBySlug(ctx, slug)

	if existingProduct != nil && s.settings.SkipExisting {
		s.addLog("INFO", fmt.Sprintf("محصول '%s' قبلاً موجود است، رد شد", dkProduct.Title))
		return nil, fmt.Errorf("محصول قبلاً موجود است")
	}

	// ساخت محصول جدید
	product := &models.Product{
		Name:            dkProduct.Title,
		Slug:            slug,
		Price:           float64(dkProduct.Price.Selling),
		OldPrice:        float64(dkProduct.Price.RRPPrice),
		DiscountPercent: float64(dkProduct.Price.DiscountPercent),
		Description:     dkProduct.Description,
		FullDescription: dkProduct.Description,
		Status:          "active",
		IsActive:        true,
		StockStatus:     "in_stock",
		StockQuantity:   100, // مقدار پیش‌فرض
		CategoryID:      s.targetCategoryID,
		Tags:            "digikala-import",
		URL:             fmt.Sprintf("%s/product/dkp-%d/", s.baseURL, dkProduct.ID),
	}

	// محاسبه قیمت فروش
	if product.OldPrice > 0 && product.Price < product.OldPrice {
		product.SalePrice = product.Price
		product.IsOnSale = true
		product.DiscountAmount = product.OldPrice - product.Price
	}

	// ذخیره محصول
	if err := s.ProductRepo.Create(ctx, product); err != nil {
		s.addLog("ERROR", fmt.Sprintf("خطا در ذخیره محصول: %v", err))
		return nil, fmt.Errorf("خطا در ذخیره محصول: %v", err)
	}

	// ذخیره تصاویر
	if s.settings.ImportImages {
		for i, img := range dkProduct.Images {
			if i >= 10 { // محدودیت 10 تصویر
				break
			}

			// دانلود تصویر به صورت local
			localURL, err := s.downloadImage(img.URL)
			if err != nil {
				s.addLog("WARNING", fmt.Sprintf("خطا در دانلود تصویر %d: %v", i+1, err))
				// در صورت خطا، از URL اصلی استفاده کن
				localURL = img.URL
			}

			productImage := &models.ProductImage{
				ProductID: product.ID,
				ImageURL:  localURL, // استفاده از local path
				SortOrder: i,
			}
			s.DB.Create(productImage)

			// یک تاخیر کوچک بین دانلود عکس‌ها
			time.Sleep(200 * time.Millisecond)
		}
	}

	// ذخیره مشخصات فنی
	if s.settings.ImportSpecs {
		for _, spec := range dkProduct.Specifications {
			for _, val := range spec.Values {
				productSpec := &models.ProductSpecification{
					ProductID: product.ID,
					Name:      val.Title,
					Value:     val.Value,
				}
				s.DB.Create(productSpec)
			}
		}
	}

	s.addLog("SUCCESS", fmt.Sprintf("محصول '%s' با موفقیت ذخیره شد (ID: %d)", product.Name, product.ID))
	return product, nil
}

// ImportProductsFromCategory import محصولات از یک دسته‌بندی
func (s *DigikalaProductScraperService) ImportProductsFromCategory(
	ctx context.Context,
	categoryURL string,
	maxProducts int,
	progressCallback func(imported, failed, skipped, total int),
) (imported int, failed int, skipped int, err error) {

	s.addLog("INFO", fmt.Sprintf("شروع import محصولات از دسته‌بندی (حداکثر %d محصول)", maxProducts))

	// دریافت لیست URL محصولات
	productURLs, err := s.FetchCategoryProducts(categoryURL, maxProducts)
	if err != nil {
		s.addLog("ERROR", fmt.Sprintf("خطا در دریافت لیست محصولات: %v", err))
		return 0, 0, 0, err
	}

	totalProducts := len(productURLs)
	s.addLog("INFO", fmt.Sprintf("شروع اسکرپ %d محصول", totalProducts))

	// اسکرپ و ذخیره هر محصول
	for i, productURL := range productURLs {
		s.addLog("INFO", fmt.Sprintf("پردازش محصول %d از %d", i+1, totalProducts))

		// اسکرپ محصول
		dkProduct, err := s.ScrapeProduct(productURL)
		if err != nil {
			s.addLog("ERROR", fmt.Sprintf("خطا در اسکرپ محصول: %v", err))
			failed++
			if progressCallback != nil {
				progressCallback(imported, failed, skipped, totalProducts)
			}
			continue
		}

		// ذخیره محصول
		_, err = s.SaveProduct(ctx, dkProduct)
		if err != nil {
			if strings.Contains(err.Error(), "قبلاً موجود است") {
				skipped++
			} else {
				failed++
			}
			if progressCallback != nil {
				progressCallback(imported, failed, skipped, totalProducts)
			}
			continue
		}

		imported++

		// به‌روزرسانی progress
		if progressCallback != nil {
			progressCallback(imported, failed, skipped, totalProducts)
		}

		// تاخیر بین محصولات (مثل انسان!)
		sleepDuration := s.settings.Delay + time.Duration(rand.Intn(2000))*time.Millisecond
		s.addLog("INFO", fmt.Sprintf("صبر %v قبل از محصول بعدی...", sleepDuration))
		time.Sleep(sleepDuration)
	}

	s.addLog("SUCCESS", fmt.Sprintf("Import تکمیل شد: %d موفق, %d ناموفق, %d رد شده", imported, failed, skipped))
	return imported, failed, skipped, nil
}

// containsURL بررسی وجود URL در لیست
func containsURL(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func mustParseInt(s string) int {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "٬", "")
	s = strings.ReplaceAll(s, ",", "")

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func generateSlug(title string) string {
	// تبدیل به lowercase
	slug := strings.ToLower(title)

	// جایگزینی فاصله با -
	slug = strings.ReplaceAll(slug, " ", "-")

	// حذف کاراکترهای غیرمجاز
	reg := regexp.MustCompile(`[^a-z0-9\-\p{Arabic}]+`)
	slug = reg.ReplaceAllString(slug, "")

	// حذف - های متوالی
	slug = regexp.MustCompile(`-+`).ReplaceAllString(slug, "-")

	// حذف - از ابتدا و انتها
	slug = strings.Trim(slug, "-")

	// اضافه کردن timestamp برای unique بودن
	slug = fmt.Sprintf("%s-%d", slug, time.Now().Unix()%10000)

	return slug
}
