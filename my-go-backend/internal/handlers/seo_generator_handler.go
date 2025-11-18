package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
)

// SeoGeneratorHandler handler برای تولید پویای فیلدهای SEO
type SeoGeneratorHandler struct {
	DB *gorm.DB
}

// NewSeoGeneratorHandler ایجاد نمونه جدید از handler
func NewSeoGeneratorHandler(db *gorm.DB) *SeoGeneratorHandler {
	return &SeoGeneratorHandler{DB: db}
}

// GenerateSeoTags تولید تگ‌های SEO برای یک نوشته
func (h *SeoGeneratorHandler) GenerateSeoTags(c *gin.Context) {
	postID := c.Param("id")
	if postID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه نوشته الزامی است", nil))
		return
	}

	// تبدیل ID به عدد
	id, err := strconv.ParseUint(postID, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه نوشته نامعتبر است", nil))
		return
	}

	// دریافت نوشته از دیتابیس
	var post models.Post
	if err := h.DB.Preload("Category").Preload("Author").First(&post, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "نوشته یافت نشد", nil))
		return
	}

	// دریافت نام نویسنده
	authorName := ""
	if post.Author.Name != "" {
		authorName = post.Author.Name
	}

	// دریافت نام دسته‌بندی
	categoryName := ""
	if post.Category.Name != "" {
		categoryName = post.Category.Name
	}

	// تبدیل تگ‌ها از JSON به آرایه
	var tags []string
	if post.Tags != nil {
		// اینجا باید JSON را parse کنیم
		// برای سادگی فعلاً خالی می‌گذاریم
	}

	// دریافت URL سایت از تنظیمات (در آینده می‌توان از config خواند)
	siteURL := "https://example.com" // این مقدار باید از تنظیمات خوانده شود

	// آماده‌سازی داده‌ها برای تولید SEO
	seoData := services.PostSeoData{
		Title:           post.Title,
		Excerpt:         post.Excerpt,
		Content:         post.Content,
		Slug:            post.Slug,
		MetaTitle:       post.MetaTitle,
		MetaDescription: post.MetaDescription,
		MetaKeywords:    post.MetaKeywords,
		FeaturedImage:   post.FeaturedImage,
		RobotsMeta:      post.RobotsMeta,
		AuthorName:      authorName,
		PublishedAt:     post.PublishedAt,
		UpdatedAt:       post.UpdatedAt,
		CategoryName:    categoryName,
		Tags:            tags,
		Language:        post.Language,
		SiteURL:         siteURL,

		// فیلدهای اضافی برای اسکیمای خاص
		Location:      post.Location,
		EventDate:     post.EventDate,
		EstimatedTime: post.EstimatedTime,
		// ProductPrice و Rating از schema templates استفاده می‌کنند
	}

	// تولید تگ‌های SEO
	seoService := services.NewSeoGeneratorService()
	seoTags := seoService.GenerateAllSeoTags(seoData)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    seoTags,
	})
}

// GenerateSeoTagsFromData تولید تگ‌های SEO از داده‌های ارسالی
func (h *SeoGeneratorHandler) GenerateSeoTagsFromData(c *gin.Context) {
	type SeoDataRequest struct {
		Title           string   `json:"title"`
		Excerpt         string   `json:"excerpt"`
		Content         string   `json:"content"`
		Slug            string   `json:"slug"`
		MetaTitle       string   `json:"meta_title"`
		MetaDescription string   `json:"meta_description"`
		MetaKeywords    string   `json:"meta_keywords"`
		FeaturedImage   string   `json:"featured_image"`
		RobotsMeta      string   `json:"robots_meta"`
		AuthorName      string   `json:"author_name"`
		CategoryName    string   `json:"category_name"`
		Tags            []string `json:"tags"`
		Language        string   `json:"language"`
		SiteURL         string   `json:"site_url"`

		// فیلدهای اضافی برای اسکیمای خاص
		Location      string  `json:"location"`
		EventDate     string  `json:"event_date"`
		EstimatedTime int     `json:"estimated_time"`
		ProductPrice  float64 `json:"product_price"`
		Rating        float64 `json:"rating"`
	}

	var request SeoDataRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	// تنظیم مقادیر پیش‌فرض
	if request.RobotsMeta == "" {
		request.RobotsMeta = "index,follow"
	}
	if request.Language == "" {
		request.Language = "fa"
	}
	if request.SiteURL == "" {
		request.SiteURL = "https://example.com"
	}

	// آماده‌سازی داده‌ها برای تولید SEO
	seoData := services.PostSeoData{
		Title:           request.Title,
		Excerpt:         request.Excerpt,
		Content:         request.Content,
		Slug:            request.Slug,
		MetaTitle:       request.MetaTitle,
		MetaDescription: request.MetaDescription,
		MetaKeywords:    request.MetaKeywords,
		FeaturedImage:   request.FeaturedImage,
		RobotsMeta:      request.RobotsMeta,
		AuthorName:      request.AuthorName,
		CategoryName:    request.CategoryName,
		Tags:            request.Tags,
		Language:        request.Language,
		SiteURL:         request.SiteURL,

		// فیلدهای اضافی برای اسکیمای خاص
		Location:      request.Location,
		EventDate:     nil, // در آینده می‌توان parse کرد
		EstimatedTime: request.EstimatedTime,
		// ProductPrice و Rating از schema templates استفاده می‌کنند
	}

	// تولید تگ‌های SEO
	seoService := services.NewSeoGeneratorService()
	seoTags := seoService.GenerateAllSeoTags(seoData)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    seoTags,
	})
}

// PreviewSeoTags پیش‌نمایش تگ‌های SEO
func (h *SeoGeneratorHandler) PreviewSeoTags(c *gin.Context) {
	type PreviewRequest struct {
		Title           string `json:"title"`
		Excerpt         string `json:"excerpt"`
		Content         string `json:"content"`
		Slug            string `json:"slug"`
		MetaTitle       string `json:"meta_title"`
		MetaDescription string `json:"meta_description"`
		MetaKeywords    string `json:"meta_keywords"`
		FeaturedImage   string `json:"featured_image"`
		RobotsMeta      string `json:"robots_meta"`
		AuthorName      string `json:"author_name"`
		CategoryName    string `json:"category_name"`
		Tags            string `json:"tags"` // به صورت string دریافت می‌شود
		Language        string `json:"language"`
		SiteURL         string `json:"site_url"`

		// فیلدهای اضافی برای اسکیمای خاص
		Location      string  `json:"location"`
		EventDate     string  `json:"event_date"`
		EstimatedTime int     `json:"estimated_time"`
		ProductPrice  float64 `json:"product_price"`
		Rating        float64 `json:"rating"`
	}

	var request PreviewRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	// تبدیل تگ‌ها از string به آرایه
	var tags []string
	if request.Tags != "" {
		tags = strings.Split(request.Tags, ",")
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(tag)
		}
	}

	// تنظیم مقادیر پیش‌فرض
	if request.RobotsMeta == "" {
		request.RobotsMeta = "index,follow"
	}
	if request.Language == "" {
		request.Language = "fa"
	}
	if request.SiteURL == "" {
		request.SiteURL = "https://example.com"
	}

	// آماده‌سازی داده‌ها برای تولید SEO
	seoData := services.PostSeoData{
		Title:           request.Title,
		Excerpt:         request.Excerpt,
		Content:         request.Content,
		Slug:            request.Slug,
		MetaTitle:       request.MetaTitle,
		MetaDescription: request.MetaDescription,
		MetaKeywords:    request.MetaKeywords,
		FeaturedImage:   request.FeaturedImage,
		RobotsMeta:      request.RobotsMeta,
		AuthorName:      request.AuthorName,
		CategoryName:    request.CategoryName,
		Tags:            tags,
		Language:        request.Language,
		SiteURL:         request.SiteURL,

		// فیلدهای اضافی برای اسکیمای خاص
		Location:      request.Location,
		EventDate:     nil, // در آینده می‌توان parse کرد
		EstimatedTime: request.EstimatedTime,
		// ProductPrice و Rating از schema templates استفاده می‌کنند
	}

	// تولید تگ‌های SEO
	seoService := services.NewSeoGeneratorService()
	seoTags := seoService.GenerateAllSeoTags(seoData)

	// اضافه کردن اطلاعات اضافی برای پیش‌نمایش
	previewData := map[string]interface{}{
		"seo_tags": seoTags,
		"preview": map[string]interface{}{
			"og_title":         seoTags["og_title"],
			"og_description":   seoTags["og_description"],
			"og_site_name":     seoTags["og_site_name"],
			"meta_title":       request.MetaTitle,
			"meta_description": request.MetaDescription,
			"canonical_url":    request.SiteURL + "/blog/" + request.Slug,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    previewData,
	})
}
