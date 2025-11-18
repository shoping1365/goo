package services

import (
	"encoding/json"
	"fmt"
	"html"
	"strings"
	"time"
)

// SeoGeneratorService سرویس تولید پویای فیلدهای SEO
type SeoGeneratorService struct{}

// NewSeoGeneratorService ایجاد نمونه جدید از سرویس
func NewSeoGeneratorService() *SeoGeneratorService {
	return &SeoGeneratorService{}
}

// PostSeoData داده‌های مورد نیاز برای تولید SEO
type PostSeoData struct {
	Title           string
	Excerpt         string
	Content         string
	Slug            string
	MetaTitle       string
	MetaDescription string
	MetaKeywords    string
	FeaturedImage   string
	RobotsMeta      string
	AuthorName      string
	PublishedAt     *time.Time
	UpdatedAt       time.Time
	CategoryName    string
	Tags            []string
	Language        string
	SiteURL         string

	// فیلدهای اضافی برای اسکیمای خاص
	Location      string
	EventDate     *time.Time
	EstimatedTime int
	ProductPrice  float64
	Rating        float64
}

// محاسبه تعداد کلمات
func (s *SeoGeneratorService) calculateWordCount(content string) int {
	// حذف HTML tags
	content = s.stripHTML(content)
	// تقسیم بر فاصله و شمارش
	words := strings.Fields(content)
	return len(words)
}

// محاسبه زمان مطالعه (به دقیقه)
func (s *SeoGeneratorService) calculateReadingTime(wordCount int) int {
	// میانگین سرعت مطالعه: 200 کلمه در دقیقه
	readingTime := wordCount / 200
	if readingTime < 1 {
		readingTime = 1
	}
	return readingTime
}

// GenerateOgTitle تولید عنوان Open Graph به صورت پویا
func (s *SeoGeneratorService) GenerateOgTitle(data PostSeoData) string {
	// اگر meta_title وجود داشت، از آن استفاده کن
	if data.MetaTitle != "" {
		return data.MetaTitle
	}

	// در غیر این صورت از عنوان اصلی استفاده کن
	return data.Title
}

// GenerateOgDescription تولید توضیحات Open Graph به صورت پویا
func (s *SeoGeneratorService) GenerateOgDescription(data PostSeoData) string {
	// اگر meta_description وجود داشت، از آن استفاده کن
	if data.MetaDescription != "" {
		return data.MetaDescription
	}

	// در غیر این صورت از خلاصه استفاده کن
	if data.Excerpt != "" {
		// حذف HTML tags و محدود کردن طول
		cleanExcerpt := s.stripHTML(data.Excerpt)
		if len(cleanExcerpt) > 160 {
			return cleanExcerpt[:157] + "..."
		}
		return cleanExcerpt
	}

	// در نهایت از محتوای اصلی استفاده کن
	if data.Content != "" {
		cleanContent := s.stripHTML(data.Content)
		if len(cleanContent) > 160 {
			return cleanContent[:157] + "..."
		}
		return cleanContent
	}

	return ""
}

// GenerateOgSiteName تولید نام سایت Open Graph به صورت پویا
func (s *SeoGeneratorService) GenerateOgSiteName(data PostSeoData) string {
	// اگر نام سایت تنظیم نشده، از URL استخراج کن
	if data.SiteURL != "" {
		// استخراج نام دامنه از URL
		parts := strings.Split(data.SiteURL, "//")
		if len(parts) > 1 {
			domain := strings.Split(parts[1], "/")[0]
			return domain
		}
	}

	return "وب‌سایت من"
}

// GenerateOgImage تولید تصویر Open Graph به صورت پویا
func (s *SeoGeneratorService) GenerateOgImage(data PostSeoData) string {
	// اگر تصویر شاخص وجود داشت، از آن استفاده کن
	if data.FeaturedImage != "" {
		return data.FeaturedImage
	}

	// در غیر این صورت تصویر پیش‌فرض
	return data.SiteURL + "/default-og-image.jpg"
}

// GenerateOgType تولید نوع Open Graph به صورت پویا
func (s *SeoGeneratorService) GenerateOgType(data PostSeoData) string {
	// برای مقالات همیشه "article"
	return "article"
}

// determineSchemaType تعیین نوع اسکیما بر اساس محتوا
func (s *SeoGeneratorService) determineSchemaType(data PostSeoData) string {
	// بر اساس محتوا و فیلدهای موجود، نوع اسکیما را تعیین کن
	// این منطق می‌تواند بر اساس category_name، tags، یا سایر فیلدها باشد

	// اگر دسته‌بندی خبری باشد
	if data.CategoryName != "" && strings.Contains(strings.ToLower(data.CategoryName), "خبر") {
		return "NewsArticle"
	}

	// اگر دسته‌بندی آموزشی باشد
	if data.CategoryName != "" && strings.Contains(strings.ToLower(data.CategoryName), "آموزش") {
		return "TechArticle"
	}

	// اگر دسته‌بندی بررسی محصول باشد
	if data.CategoryName != "" && strings.Contains(strings.ToLower(data.CategoryName), "بررسی") {
		return "Review"
	}

	// پیش‌فرض
	return "Article"
}

// addSchemaSpecificFields اضافه کردن فیلدهای خاص بر اساس نوع اسکیما
func (s *SeoGeneratorService) addSchemaSpecificFields(schema map[string]interface{}, data PostSeoData, schemaType string) {
	switch schemaType {
	case "NewsArticle":
		// فیلدهای خاص مقالات خبری
		if data.Location != "" {
			schema["dateline"] = data.Location
		}
		if data.EventDate != nil {
			schema["dateCreated"] = data.EventDate.Format(time.RFC3339)
		}

	case "TechArticle":
		// فیلدهای خاص مقالات آموزشی
		if data.EstimatedTime > 0 {
			schema["timeRequired"] = fmt.Sprintf("PT%dM", data.EstimatedTime)
		}

	case "Review":
		// فیلدهای خاص بررسی محصول
		if data.ProductPrice > 0 {
			schema["offers"] = map[string]interface{}{
				"@type":         "Offer",
				"price":         data.ProductPrice,
				"priceCurrency": "IRR",
			}
		}
		if data.Rating > 0 {
			schema["reviewRating"] = map[string]interface{}{
				"@type":       "Rating",
				"ratingValue": data.Rating,
				"bestRating":  5,
			}
		}
	}
}

// GenerateHTMLMetaTags تولید تگ‌های HTML Meta
func (s *SeoGeneratorService) GenerateHTMLMetaTags(data PostSeoData) string {
	var tags []string

	// Meta Title
	if data.MetaTitle != "" {
		tags = append(tags, fmt.Sprintf(`<meta name="title" content="%s">`, html.EscapeString(data.MetaTitle)))
	}

	// Meta Description
	if data.MetaDescription != "" {
		tags = append(tags, fmt.Sprintf(`<meta name="description" content="%s">`, html.EscapeString(data.MetaDescription)))
	}

	// Meta Keywords
	if data.MetaKeywords != "" {
		tags = append(tags, fmt.Sprintf(`<meta name="keywords" content="%s">`, html.EscapeString(data.MetaKeywords)))
	}

	// Robots
	if data.RobotsMeta != "" {
		tags = append(tags, fmt.Sprintf(`<meta name="robots" content="%s">`, data.RobotsMeta))
	}

	// Canonical URL
	if data.Slug != "" {
		canonicalURL := fmt.Sprintf("%s/blog/%s", data.SiteURL, data.Slug)
		tags = append(tags, fmt.Sprintf(`<link rel="canonical" href="%s">`, canonicalURL))
	}

	// Language
	if data.Language != "" {
		tags = append(tags, fmt.Sprintf(`<meta http-equiv="content-language" content="%s">`, data.Language))
	}

	return strings.Join(tags, "\n")
}

// GenerateOpenGraphTags تولید تگ‌های Open Graph
func (s *SeoGeneratorService) GenerateOpenGraphTags(data PostSeoData) string {
	var tags []string

	// OG Title
	ogTitle := s.GenerateOgTitle(data)
	if ogTitle != "" {
		tags = append(tags, fmt.Sprintf(`<meta property="og:title" content="%s">`, html.EscapeString(ogTitle)))
	}

	// OG Description
	ogDescription := s.GenerateOgDescription(data)
	if ogDescription != "" {
		tags = append(tags, fmt.Sprintf(`<meta property="og:description" content="%s">`, html.EscapeString(ogDescription)))
	}

	// OG Type
	ogType := s.GenerateOgType(data)
	tags = append(tags, fmt.Sprintf(`<meta property="og:type" content="%s">`, ogType))

	// OG Image
	ogImage := s.GenerateOgImage(data)
	tags = append(tags, fmt.Sprintf(`<meta property="og:image" content="%s">`, ogImage))

	// OG Site Name
	ogSiteName := s.GenerateOgSiteName(data)
	if ogSiteName != "" {
		tags = append(tags, fmt.Sprintf(`<meta property="og:site_name" content="%s">`, html.EscapeString(ogSiteName)))
	}

	// OG URL
	if data.Slug != "" {
		ogURL := fmt.Sprintf("%s/blog/%s", data.SiteURL, data.Slug)
		tags = append(tags, fmt.Sprintf(`<meta property="og:url" content="%s">`, ogURL))
	}

	// OG Locale
	if data.Language == "fa" {
		tags = append(tags, `<meta property="og:locale" content="fa_IR">`)
	}

	return strings.Join(tags, "\n")
}

// GenerateTwitterCardTags تولید تگ‌های Twitter Card
func (s *SeoGeneratorService) GenerateTwitterCardTags(data PostSeoData) string {
	var tags []string

	// Twitter Card Type
	tags = append(tags, `<meta name="twitter:card" content="summary_large_image">`)

	// Twitter Title
	ogTitle := s.GenerateOgTitle(data)
	if ogTitle != "" {
		tags = append(tags, fmt.Sprintf(`<meta name="twitter:title" content="%s">`, html.EscapeString(ogTitle)))
	}

	// Twitter Description
	ogDescription := s.GenerateOgDescription(data)
	if ogDescription != "" {
		tags = append(tags, fmt.Sprintf(`<meta name="twitter:description" content="%s">`, html.EscapeString(ogDescription)))
	}

	// Twitter Image
	ogImage := s.GenerateOgImage(data)
	tags = append(tags, fmt.Sprintf(`<meta name="twitter:image" content="%s">`, ogImage))

	return strings.Join(tags, "\n")
}

// GenerateJSONLD تولید JSON-LD Schema به صورت کاملاً پویا
func (s *SeoGeneratorService) GenerateJSONLD(data PostSeoData) string {
	// تعیین نوع اسکیما بر اساس محتوا
	schemaType := s.determineSchemaType(data)

	schema := map[string]interface{}{
		"@context": "https://schema.org",
		"@type":    schemaType,
	}

	// عنوان
	if data.Title != "" {
		schema["headline"] = data.Title
	}

	// توضیحات
	if data.MetaDescription != "" {
		schema["description"] = data.MetaDescription
	} else if data.Excerpt != "" {
		schema["description"] = s.stripHTML(data.Excerpt)
	}

	// نویسنده
	if data.AuthorName != "" {
		schema["author"] = map[string]interface{}{
			"@type": "Person",
			"name":  data.AuthorName,
		}
	}

	// تاریخ انتشار
	if data.PublishedAt != nil {
		schema["datePublished"] = data.PublishedAt.Format(time.RFC3339)
	}

	// تاریخ ویرایش
	schema["dateModified"] = data.UpdatedAt.Format(time.RFC3339)

	// ناشر
	ogSiteName := s.GenerateOgSiteName(data)
	if ogSiteName != "" {
		schema["publisher"] = map[string]interface{}{
			"@type": "Organization",
			"name":  ogSiteName,
		}
	}

	// URL اصلی
	if data.Slug != "" {
		mainEntityURL := fmt.Sprintf("%s/blog/%s", data.SiteURL, data.Slug)
		schema["mainEntityOfPage"] = map[string]interface{}{
			"@type": "WebPage",
			"@id":   mainEntityURL,
		}
	}

	// تصویر
	ogImage := s.GenerateOgImage(data)
	if ogImage != "" {
		schema["image"] = ogImage
	}

	// دسته‌بندی
	if data.CategoryName != "" {
		schema["articleSection"] = data.CategoryName
	}

	// کلمات کلیدی
	if data.MetaKeywords != "" {
		schema["keywords"] = data.MetaKeywords
	}

	// تعداد کلمات
	wordCount := s.calculateWordCount(data.Content)
	if wordCount > 0 {
		schema["wordCount"] = wordCount
	}

	// زمان مطالعه
	readingTime := s.calculateReadingTime(wordCount)
	if readingTime > 0 {
		schema["timeRequired"] = fmt.Sprintf("PT%dM", readingTime)
	}

	// زبان
	if data.Language != "" {
		schema["inLanguage"] = data.Language
	}

	// تگ‌ها
	if len(data.Tags) > 0 {
		schema["keywords"] = strings.Join(data.Tags, ", ")
	}

	// فیلدهای خاص بر اساس نوع اسکیما
	s.addSchemaSpecificFields(schema, data, schemaType)

	// تبدیل به JSON
	jsonData, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return ""
	}

	return string(jsonData)
}

// GenerateAllSeoTags تولید تمام تگ‌های SEO
func (s *SeoGeneratorService) GenerateAllSeoTags(data PostSeoData) map[string]string {
	return map[string]string{
		"html_meta_tags":    s.GenerateHTMLMetaTags(data),
		"open_graph_tags":   s.GenerateOpenGraphTags(data),
		"twitter_card_tags": s.GenerateTwitterCardTags(data),
		"json_ld_schema":    s.GenerateJSONLD(data),
		"og_title":          s.GenerateOgTitle(data),
		"og_description":    s.GenerateOgDescription(data),
		"og_site_name":      s.GenerateOgSiteName(data),
	}
}

// stripHTML حذف تگ‌های HTML از متن
func (s *SeoGeneratorService) stripHTML(html string) string {
	// حذف تگ‌های HTML ساده
	html = strings.ReplaceAll(html, "<br>", " ")
	html = strings.ReplaceAll(html, "<br/>", " ")
	html = strings.ReplaceAll(html, "<br />", " ")
	html = strings.ReplaceAll(html, "<p>", " ")
	html = strings.ReplaceAll(html, "</p>", " ")
	html = strings.ReplaceAll(html, "<div>", " ")
	html = strings.ReplaceAll(html, "</div>", " ")
	html = strings.ReplaceAll(html, "<span>", "")
	html = strings.ReplaceAll(html, "</span>", "")

	// حذف سایر تگ‌های HTML با regex
	// این یک روش ساده است، برای تولید بهتر می‌توان از کتابخانه‌های HTML parser استفاده کرد
	html = strings.ReplaceAll(html, "<[^>]*>", "")

	// حذف فاصله‌های اضافی
	html = strings.Join(strings.Fields(html), " ")

	return strings.TrimSpace(html)
}
