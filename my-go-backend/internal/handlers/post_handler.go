package handlers

import (
	"net/http"
	"time"

	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostHandler struct {
	DB      *gorm.DB
	Service *services.PostService
}

func NewPostHandler(db *gorm.DB) *PostHandler {
	return &PostHandler{DB: db, Service: services.NewPostService(db)}
}

// ListPosts دریافت همه نوشته‌ها
func (h *PostHandler) ListPosts(c *gin.Context) {
	role, _ := c.Get("role")
	showAll := (c.Query("all") == "1" || c.Query("all") == "true") && role == "developer"
	posts, err := h.Service.ListPosts(c.Request.Context(), showAll)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetPost دریافت یک نوشته خاص
func (h *PostHandler) GetPost(c *gin.Context) {
	id := c.Param("id")
	post, err := h.Service.GetPostByIDOrSlug(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	role, _ := c.Get("role")
	isPreview := c.Query("preview") == "1" || c.Query("preview") == "true"
	if !h.Service.CanViewPost(post, role, isPreview) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// CreatePost ایجاد نوشته جدید
func (h *PostHandler) CreatePost(c *gin.Context) {
	var input services.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	post, err := h.Service.CreatePost(c.Request.Context(), input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, post)
}

// UpdatePost ویرایش نوشته موجود
func (h *PostHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var input services.UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	post, err := h.Service.UpdatePost(c.Request.Context(), id, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, post)
}

// DeletePost حذف نوشته
func (h *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeletePost(c.Request.Context(), id); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("POST_NOT_FOUND", utils.GetErrorMessage("POST_NOT_FOUND"), nil))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "نوشته با موفقیت حذف شد"})
}

// PublishScheduledPosts انتشار خودکار پست‌های زمان‌بندی شده
// این تابع باید به صورت دوره‌ای (مثلاً هر دقیقه) اجرا شود و پست‌هایی که وضعیت scheduled دارند و زمان scheduled_at آن‌ها گذشته است را منتشر کند.
func PublishScheduledPosts(db *gorm.DB) {
	for {
		// جستجوی پست‌های زمان‌بندی شده که زمان انتشارشان رسیده یا گذشته است
		type Post struct{ ID uint }
		var posts []Post
		err := db.Where("status = ? AND scheduled_at IS NOT NULL AND scheduled_at <= ?", "scheduled", time.Now()).Find(&posts).Error
		if err == nil && len(posts) > 0 {
			for _, post := range posts {
				// به‌روزرسانی وضعیت به published و مقداردهی published_at
				db.Model(&post).Updates(map[string]interface{}{
					"status":       "published",
					"published_at": time.Now(),
				})
			}
		}
		// هر دقیقه یکبار بررسی شود
		time.Sleep(1 * time.Minute)
	}
}

// CheckSlugUnique بررسی تکراری نبودن slug
func (h *PostHandler) CheckSlugUnique(c *gin.Context) {
	slug := c.Query("slug")
	if slug == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "slug الزامی است", nil))
		return
	}
	exists, err := h.Service.CheckSlugUnique(c.Request.Context(), slug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"exists": !exists, "slug": slug})
}

// GenerateUniqueSlug تولید slug یکتا
func (h *PostHandler) GenerateUniqueSlug(c *gin.Context) {
	baseSlug := c.Query("base_slug")
	if baseSlug == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "base_slug الزامی است", nil))
		return
	}
	uniqueSlug, err := h.Service.GenerateUniqueSlug(c.Request.Context(), baseSlug)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"slug": uniqueSlug})
}
