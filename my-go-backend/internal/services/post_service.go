package services

import (
	"context"
	"fmt"
	"time"

	"my-go-backend/internal/models"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// PostService سرویس مدیریت نوشته‌ها (Blog Posts)
// وظیفه: شامل منطق تجاری مربوط به ایجاد/به‌روزرسانی/حذف و دریافت پست‌ها
// و اعمال قواعد دسترسی (preview, role) و تولید slug یکتا
type PostService struct {
	db *gorm.DB
}

// NewPostService ایجاد نمونه جدید سرویس پست
func NewPostService(db *gorm.DB) *PostService { return &PostService{db: db} }

// ListPosts فهرست پست‌ها با درنظر گرفتن نقش کاربر و نمایش همه یا فقط منتشرشده‌ها
func (s *PostService) ListPosts(ctx context.Context, showAll bool) ([]models.Post, error) {
	var posts []models.Post
	query := s.db.WithContext(ctx).Preload("Category").Preload("Author").Order("created_at DESC")
	if !showAll {
		query = query.Where("status = ?", "published")
	}
	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// GetPostByIDOrSlug دریافت پست با ID عددی یا slug متنی + قواعد preview/role در لایه بالاتر اعمال شود
func (s *PostService) GetPostByIDOrSlug(ctx context.Context, idOrSlug string) (*models.Post, error) {
	var post models.Post
	// تلاش ابتدا با ID
	if numID, err := strconvAtoi(idOrSlug); err == nil {
		if err := s.db.WithContext(ctx).Preload("Category").Preload("Author").First(&post, numID).Error; err != nil {
			return nil, err
		}
		return &post, nil
	}
	// سپس با slug
	if err := s.db.WithContext(ctx).Preload("Category").Preload("Author").Where("slug = ?", idOrSlug).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// CanViewPost آیا کاربر با نقش داده‌شده اجازه مشاهده پست را دارد
func (s *PostService) CanViewPost(post *models.Post, role any, isPreview bool) bool {
	if post.Status == "published" {
		return true
	}
	if role == nil {
		return false
	}
	// فقط developer برای پیش‌نویس/منتشرنشده
	if roleStr, ok := role.(string); ok {
		if roleStr == "developer" {
			return true
		}
	}
	return false
}

// CreatePost ایجاد پست جدید با تولید slug یکتا و اعتبارسنجی حداقلی
func (s *PostService) CreatePost(ctx context.Context, input CreatePostInput) (*models.Post, error) {
	if input.Title == "" || input.Slug == "" {
		return nil, fmt.Errorf("title/slug الزامی است")
	}
	newSlug, err := s.generateUniqueSlug(ctx, input.Slug)
	if err != nil {
		return nil, err
	}

	var eventDate *time.Time
	if input.EventDate != "" {
		if parsed, err := time.Parse("2006-01-02", input.EventDate); err == nil {
			eventDate = &parsed
		}
	}

	post := models.Post{
		Title:               input.Title,
		Slug:                newSlug,
		Excerpt:             input.Excerpt,
		Content:             input.Content,
		Status:              valueOrDefault(input.Status, "draft"),
		CategoryID:          input.CategoryID,
		FeaturedImage:       input.FeaturedImage,
		AuthorID:            valueOrDefaultUint(input.AuthorID, 1),
		MetaTitle:           input.MetaTitle,
		MetaKeywords:        input.MetaKeywords,
		MetaDescription:     input.MetaDescription,
		RobotsMeta:          input.RobotsMeta,
		Tags:                toJSONOrEmptyArray(input.Tags),
		Pros:                input.Pros,
		Cons:                input.Cons,
		EstimatedTime:       input.EstimatedTime,
		Location:            input.Location,
		EventDate:           eventDate,
		Language:            input.Language,
		IsAccessibleForFree: input.IsAccessibleForFree,
	}
	if post.Status == "published" {
		now := time.Now()
		post.PublishedAt = &now
	}
	if err := s.db.WithContext(ctx).Create(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// UpdatePost ویرایش پست موجود با کنترل یکتایی slug
func (s *PostService) UpdatePost(ctx context.Context, id string, input UpdatePostInput) (*models.Post, error) {
	var post models.Post
	if err := s.db.WithContext(ctx).First(&post, id).Error; err != nil {
		return nil, err
	}

	if input.Title != nil {
		post.Title = *input.Title
	}
	if input.Slug != nil && *input.Slug != "" {
		var cnt int64
		if err := s.db.WithContext(ctx).Model(&models.Post{}).Where("slug = ? AND id != ?", *input.Slug, post.ID).Count(&cnt).Error; err != nil {
			return nil, err
		}
		if cnt > 0 {
			return nil, fmt.Errorf("duplicate slug")
		}
		post.Slug = *input.Slug
	}
	if input.Excerpt != nil {
		post.Excerpt = *input.Excerpt
	}
	if input.Content != nil {
		post.Content = *input.Content
	}
	if input.Status != nil {
		post.Status = *input.Status
	}
	if input.CategoryID != nil {
		post.CategoryID = input.CategoryID
	}
	if input.FeaturedImage != nil {
		post.FeaturedImage = *input.FeaturedImage
	}
	if input.MetaTitle != nil {
		post.MetaTitle = *input.MetaTitle
	}
	if input.MetaKeywords != nil {
		post.MetaKeywords = *input.MetaKeywords
	}
	if input.MetaDescription != nil {
		post.MetaDescription = *input.MetaDescription
	}
	if input.RobotsMeta != nil {
		post.RobotsMeta = *input.RobotsMeta
	}
	if input.Tags != nil {
		post.Tags = toJSONOrEmptyArray(*input.Tags)
	}
	if input.Pros != nil {
		post.Pros = *input.Pros
	}
	if input.Cons != nil {
		post.Cons = *input.Cons
	}
	if input.EstimatedTime != nil {
		post.EstimatedTime = *input.EstimatedTime
	}
	if input.Location != nil {
		post.Location = *input.Location
	}
	if input.EventDate != nil && *input.EventDate != "" {
		if parsed, err := time.Parse("2006-01-02", *input.EventDate); err == nil {
			post.EventDate = &parsed
		}
	}
	if input.Language != nil {
		post.Language = *input.Language
	}
	if input.IsAccessibleForFree != nil {
		post.IsAccessibleForFree = *input.IsAccessibleForFree
	}

	if err := s.db.WithContext(ctx).Save(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// DeletePost حذف پست
func (s *PostService) DeletePost(ctx context.Context, id string) error {
	return s.db.WithContext(ctx).Unscoped().Delete(&models.Post{}, id).Error
}

// CheckSlugUnique بررسی یکتایی slug
func (s *PostService) CheckSlugUnique(ctx context.Context, slug string) (bool, error) {
	var count int64
	if err := s.db.WithContext(ctx).Model(&models.Post{}).Where("slug = ?", slug).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}

// GenerateUniqueSlug تولید slug یکتا بر اساس base
func (s *PostService) GenerateUniqueSlug(ctx context.Context, base string) (string, error) {
	return s.generateUniqueSlug(ctx, base)
}

// -------------------- انواع ورودی‌ها --------------------

type CreatePostInput struct {
	Title               string      `json:"title"`
	Slug                string      `json:"slug"`
	Excerpt             string      `json:"excerpt"`
	Content             string      `json:"content"`
	Status              string      `json:"status"`
	CategoryID          *uint       `json:"category_id"`
	FeaturedImage       string      `json:"featured_image"`
	AuthorID            uint        `json:"author_id"`
	MetaTitle           string      `json:"meta_title"`
	MetaKeywords        string      `json:"meta_keywords"`
	MetaDescription     string      `json:"meta_description"`
	RobotsMeta          string      `json:"robots_meta"`
	Tags                interface{} `json:"tags"`
	Pros                string      `json:"pros"`
	Cons                string      `json:"cons"`
	EstimatedTime       int         `json:"estimated_time"`
	Location            string      `json:"location"`
	EventDate           string      `json:"event_date"`
	Language            string      `json:"language"`
	IsAccessibleForFree bool        `json:"is_accessible_for_free"`
}

type UpdatePostInput struct {
	Title               *string      `json:"title"`
	Slug                *string      `json:"slug"`
	Excerpt             *string      `json:"excerpt"`
	Content             *string      `json:"content"`
	Status              *string      `json:"status"`
	CategoryID          *uint        `json:"category_id"`
	FeaturedImage       *string      `json:"featured_image"`
	MetaTitle           *string      `json:"meta_title"`
	MetaKeywords        *string      `json:"meta_keywords"`
	MetaDescription     *string      `json:"meta_description"`
	RobotsMeta          *string      `json:"robots_meta"`
	Tags                *interface{} `json:"tags"`
	Pros                *string      `json:"pros"`
	Cons                *string      `json:"cons"`
	EstimatedTime       *int         `json:"estimated_time"`
	Location            *string      `json:"location"`
	EventDate           *string      `json:"event_date"`
	Language            *string      `json:"language"`
	IsAccessibleForFree *bool        `json:"is_accessible_for_free"`
}

// -------------------- توابع کمکی داخلی --------------------

func (s *PostService) generateUniqueSlug(ctx context.Context, baseSlug string) (string, error) {
	slug := baseSlug
	counter := 1
	for {
		var count int64
		if err := s.db.WithContext(ctx).Model(&models.Post{}).Where("slug = ?", slug).Count(&count).Error; err != nil {
			return "", err
		}
		if count == 0 {
			return slug, nil
		}
		slug = fmt.Sprintf("%s-%d", baseSlug, counter)
		counter++
	}
}

func valueOrDefault[T comparable](v T, d T) T {
	var zero T
	if v == zero {
		return d
	}
	return v
}
func valueOrDefaultUint(v uint, d uint) uint {
	if v == 0 {
		return d
	}
	return v
}

func toJSONOrEmptyArray(val interface{}) datatypes.JSON {
	if val != nil {
		if b, err := models.MarshalJSON(val); err == nil {
			return b
		}
	}
	return datatypes.JSON([]byte("[]"))
}

func strconvAtoi(s string) (int, error) {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		return 0, fmt.Errorf("not int")
	}
	return n, nil
}
