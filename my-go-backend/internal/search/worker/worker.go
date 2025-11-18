package worker

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"my-go-backend/internal/models"
	searchservice "my-go-backend/internal/search/service"
)

// Config تنظیمات Worker
// PollInterval: فاصله زمانی بررسی صف
// MaxBatchSize: تعداد رکوردهایی که در هر نوبت پردازش می‌شوند
// MaxAttempts: حداکثر دفعات تلاش جهت پردازش یک رکورد قبل از علامت‌گذاری به عنوان failed
// ContextTimeout: تایم‌اوت هر درخواست دسته‌ای برای جلوگیری از گیرکردن طولانی
// IndexResolver: تابع تبدیل resourceType به نام ایندکس Bleve
// DocumentBuilder: تابعی که داده لازم برای درج در ایندکس را تولید می‌کند.
type Config struct {
	PollInterval    time.Duration
	MaxBatchSize    int
	MaxAttempts     int
	ContextTimeout  time.Duration
	IndexResolver   func(resourceType string) (string, error)
	DocumentBuilder func(ctx context.Context, tx *gorm.DB, evt models.SearchIndexQueue) (map[string]any, error)
}

// Worker اجرای وظایف همگام‌سازی ایندکس جستجو از پایگاه داده است.
type Worker struct {
	db     *gorm.DB
	svc    *searchservice.Service
	cfg    Config
	logger *log.Logger
	stopCh chan struct{}
}

// NewWorker ایجاد نمونه جدید Worker
func NewWorker(db *gorm.DB, svc *searchservice.Service, cfg Config, logger *log.Logger) (*Worker, error) {
	if db == nil {
		return nil, errors.New("db is required")
	}
	if svc == nil {
		return nil, errors.New("search service is required")
	}
	if cfg.PollInterval <= 0 {
		cfg.PollInterval = 2 * time.Second
	}
	if cfg.MaxBatchSize <= 0 {
		cfg.MaxBatchSize = 25
	}
	if cfg.MaxAttempts <= 0 {
		cfg.MaxAttempts = 5
	}
	if cfg.ContextTimeout <= 0 {
		cfg.ContextTimeout = 5 * time.Second
	}
	if cfg.IndexResolver == nil {
		return nil, errors.New("index resolver is required")
	}
	if cfg.DocumentBuilder == nil {
		return nil, errors.New("document builder is required")
	}
	if logger == nil {
		logger = log.Default()
	}

	return &Worker{
		db:     db,
		svc:    svc,
		cfg:    cfg,
		logger: logger,
		stopCh: make(chan struct{}),
	}, nil
}

// Start اجرای Worker در گوروتین جدید است.
func (w *Worker) Start() {
	go w.run()
}

// Stop خاتمه اجرای Worker است.
func (w *Worker) Stop() {
	close(w.stopCh)
}

func (w *Worker) run() {
	w.logger.Printf("ℹ️  search worker bootstrap starting")
	if err := w.bootstrapExisting(); err != nil {
		w.logger.Printf("⚠️  search worker bootstrap error: %v", err)
	}
	w.logger.Printf("ℹ️  search worker loop started (poll interval=%s, batch=%d)", w.cfg.PollInterval, w.cfg.MaxBatchSize)
	w.loop()
}

func (w *Worker) loop() {
	ticker := time.NewTicker(w.cfg.PollInterval)
	defer ticker.Stop()

	for {
		if err := w.drainQueue(); err != nil {
			w.logger.Printf("⚠️  search worker drain error: %v", err)
		}

		select {
		case <-w.stopCh:
			return
		case <-ticker.C:
			continue
		}
	}
}

func (w *Worker) drainQueue() error {
	for {
		processed, err := w.processBatch()
		if err != nil {
			return err
		}
		if processed == 0 {
			return nil
		}
		// اگر هنوز رویداد وجود دارد، بدون انتظار قبل از برداشت بعدی ادامه بده
		select {
		case <-w.stopCh:
			return nil
		default:
		}
	}
}

func (w *Worker) processBatch() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), w.cfg.ContextTimeout)
	defer cancel()

	var events []models.SearchIndexQueue
	err := w.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("status = ? AND available_at <= NOW()", models.SearchStatusPending).
			Order("id ASC").
			Limit(w.cfg.MaxBatchSize).
			Find(&events).Error; err != nil {
			return err
		}

		if len(events) == 0 {
			return nil
		}

		for i := range events {
			events[i].Status = models.SearchStatusPending
			events[i].Attempts++
			events[i].UpdatedAt = time.Now()
			if err := tx.Model(&models.SearchIndexQueue{}).
				Where("id = ?", events[i].ID).
				Updates(map[string]any{
					"attempts":   events[i].Attempts,
					"updated_at": events[i].UpdatedAt,
					"status":     models.SearchStatusPending,
				}).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	if len(events) == 0 {
		return 0, nil
	}

	for _, evt := range events {
		err := w.handleEvent(ctx, evt)
		if err != nil {
			w.logger.Printf("⚠️  search worker handle error id=%d: %v", evt.ID, err)
			w.retryEvent(evt, err)
			continue
		}
		w.completeEvent(evt)
	}

	w.logger.Printf("ℹ️  search worker processed %d event(s) (up to id=%d)", len(events), events[len(events)-1].ID)
	return len(events), nil
}

func (w *Worker) handleEvent(ctx context.Context, evt models.SearchIndexQueue) error {
	indexName, err := w.cfg.IndexResolver(evt.ResourceType)
	if err != nil {
		return err
	}

	docID := fmt.Sprint(evt.ResourceID)
	switch evt.Operation {
	case models.SearchOpDelete:
		return w.svc.DeleteDocument(ctx, indexName, docID)
	case models.SearchOpUpsert:
		doc, err := w.cfg.DocumentBuilder(ctx, w.db, evt)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return w.svc.DeleteDocument(ctx, indexName, docID)
			}
			return err
		}
		return w.svc.IndexDocument(ctx, indexName, docID, doc)
	default:
		return fmt.Errorf("unknown search operation: %s", evt.Operation)
	}
}

func (w *Worker) retryEvent(evt models.SearchIndexQueue, cause error) {
	now := time.Now()
	nextAvailable := now.Add(time.Second * time.Duration(evt.Attempts*evt.Attempts))
	status := models.SearchStatusPending
	if evt.Attempts >= w.cfg.MaxAttempts {
		status = models.SearchStatusFailed
	}

	_ = w.db.Model(&models.SearchIndexQueue{}).
		Where("id = ?", evt.ID).
		Updates(map[string]any{
			"status":        status,
			"available_at":  nextAvailable,
			"error_message": cause.Error(),
			"updated_at":    now,
		})
}

func (w *Worker) completeEvent(evt models.SearchIndexQueue) {
	now := time.Now()
	_ = w.db.Model(&models.SearchIndexQueue{}).
		Where("id = ?", evt.ID).
		Updates(map[string]any{
			"status":     models.SearchStatusCompleted,
			"updated_at": now,
		})
}

func (w *Worker) bootstrapExisting() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// اگر قبلاً رکورد pending وجود داشته باشد، فرض می‌کنیم بوت‌استرپ قبلاً اجرا شده است تا از صف‌های تکراری جلوگیری شود
	var pendingCount int64
	if err := w.db.WithContext(ctx).Model(&models.SearchIndexQueue{}).
		Where("status = ?", models.SearchStatusPending).
		Count(&pendingCount).Error; err != nil {
		return fmt.Errorf("count pending: %w", err)
	}
	if pendingCount > 0 {
		w.logger.Printf("ℹ️  search bootstrap skipped (pending records already present: %d)", pendingCount)
		return nil
	}

	totalQueued := int64(0)

	if _, err := w.cfg.IndexResolver(models.SearchResourceProduct); err == nil {
		if queued, err := w.enqueueMissingResources(ctx, models.SearchResourceProduct, "products"); err != nil {
			return fmt.Errorf("bootstrap products: %w", err)
		} else if queued > 0 {
			totalQueued += queued
			w.logger.Printf("ℹ️  search bootstrap queued %d product record(s)", queued)
		}
	} else {
		w.logger.Printf("ℹ️  search bootstrap skipping products: %v", err)
	}

	if _, err := w.cfg.IndexResolver(models.SearchResourcePost); err == nil {
		if queued, err := w.enqueueMissingResources(ctx, models.SearchResourcePost, "posts"); err != nil {
			return fmt.Errorf("bootstrap posts: %w", err)
		} else if queued > 0 {
			totalQueued += queued
			w.logger.Printf("ℹ️  search bootstrap queued %d post record(s)", queued)
		}
	} else {
		w.logger.Printf("ℹ️  search bootstrap skipping posts: %v", err)
	}

	if totalQueued > 0 {
		w.logger.Printf("ℹ️  search bootstrap completed, queued total %d record(s)", totalQueued)
	} else {
		w.logger.Printf("ℹ️  search bootstrap found no missing records")
	}

	return nil
}

func (w *Worker) enqueueMissingResources(ctx context.Context, resourceType, table string) (int64, error) {
	stmt := fmt.Sprintf(`
		INSERT INTO search_index_queue (resource_type, resource_id, operation)
		SELECT ?, CAST(src.id AS BIGINT), ?
		FROM %s AS src
		WHERE NOT EXISTS (
			SELECT 1 FROM search_index_queue q
			WHERE q.resource_type = ?
			  AND q.resource_id = CAST(src.id AS BIGINT)
			  AND q.operation = ?
			  AND q.status = ?
		)
		AND NOT EXISTS (
			SELECT 1 FROM search_index_queue q
			WHERE q.resource_type = ?
			  AND q.resource_id = CAST(src.id AS BIGINT)
			  AND q.operation = ?
			  AND q.status = ?
			  AND q.updated_at >= COALESCE(src.updated_at, src.created_at, NOW())
		);
	`, table)

	result := w.db.WithContext(ctx).Exec(stmt,
		resourceType, models.SearchOpUpsert,
		resourceType, models.SearchOpUpsert, models.SearchStatusPending,
		resourceType, models.SearchOpUpsert, models.SearchStatusCompleted,
	)
	return result.RowsAffected, result.Error
}

// BuildDocumentFromDB نمونه‌ی پیش‌فرض DocumentBuilder را پیاده‌سازی می‌کند
// و بر اساس نوع منبع داده را از پایگاه داده واکشی می‌کند.
func BuildDocumentFromDB(db *gorm.DB) func(ctx context.Context, tx *gorm.DB, evt models.SearchIndexQueue) (map[string]any, error) {
	return func(ctx context.Context, tx *gorm.DB, evt models.SearchIndexQueue) (map[string]any, error) {
		base := tx
		if base == nil {
			base = db
		}
		switch evt.ResourceType {
		case models.SearchResourceProduct:
			return buildProductDocument(ctx, base, evt.ResourceID)
		case models.SearchResourcePost:
			return buildPostDocument(ctx, base, evt.ResourceID)
		default:
			return nil, fmt.Errorf("unknown resource type: %s", evt.ResourceType)
		}
	}
}

func buildProductDocument(ctx context.Context, db *gorm.DB, id uint64) (map[string]any, error) {
	var product models.Product
	if err := db.WithContext(ctx).
		Preload("Category").
		Preload("Brand").
		First(&product, id).Error; err != nil {
		return nil, err
	}

	var fallbackImage, fallbackThumb string
	if strings.TrimSpace(product.ImageURL) == "" || strings.TrimSpace(product.Image) == "" {
		type imageRow struct {
			ImageURL     string `gorm:"column:image_url"`
			ThumbnailURL string `gorm:"column:thumbnail_url"`
		}
		var primary imageRow
		err := db.WithContext(ctx).
			Table("product_images").
			Select("image_url, thumbnail_url").
			Where("product_id = ?", id).
			Order("sort_order ASC, id ASC").
			Limit(1).
			Scan(&primary).Error
		if err == nil {
			fallbackImage = strings.TrimSpace(primary.ImageURL)
			fallbackThumb = strings.TrimSpace(primary.ThumbnailURL)
		}
	}

	if strings.TrimSpace(product.ImageURL) == "" {
		product.ImageURL = firstString(fallbackThumb, fallbackImage, product.Image)
	}
	if strings.TrimSpace(product.Image) == "" {
		product.Image = firstString(fallbackThumb, fallbackImage, product.ImageURL)
	}

	categories := []string{}
	if product.Category != nil && product.Category.Name != "" {
		categories = append(categories, product.Category.Name)
	}

	brand := ""
	if product.Brand != nil {
		brand = product.Brand.Name
	}

	imageURL := strings.TrimSpace(product.ImageURL)
	thumbnail := strings.TrimSpace(product.Image)
	if thumbnail == "" {
		thumbnail = imageURL
	}

	doc := map[string]any{
		"id":           product.ID,
		"title":        product.Name,
		"description":  product.Description,
		"categories":   categories,
		"brand":        brand,
		"price":        product.Price,
		"sale_price":   product.SalePrice,
		"is_available": product.StockQuantity > 0,
		"image":        imageURL,
		"thumbnail":    thumbnail,
		"main_image":   imageURL,
		"slug":         product.Slug,
		"sku":          product.SKU,
		"lang":         "fa",
		"updated_at":   product.UpdatedAt,
	}

	if imageURL != "" {
		doc["image_url"] = imageURL
	}

	return doc, nil
}

func firstString(values ...string) string {
	for _, v := range values {
		if s := strings.TrimSpace(v); s != "" {
			return s
		}
	}
	return ""
}

func buildPostDocument(ctx context.Context, db *gorm.DB, id uint64) (map[string]any, error) {
	var post models.Post
	if err := db.WithContext(ctx).
		Preload("Category").
		First(&post, id).Error; err != nil {
		return nil, err
	}

	category := ""
	if post.Category.Name != "" {
		category = post.Category.Name
	}

	return map[string]any{
		"id":           post.ID,
		"title":        post.Title,
		"description":  post.Excerpt,
		"body":         post.Content,
		"slug":         post.Slug,
		"category":     category,
		"lang":         post.Language,
		"published_at": post.PublishedAt,
		"updated_at":   post.UpdatedAt,
	}, nil
}

// ResolverFromConfig تابعی ساده برای نگاشت نوع منابع به نام ایندکس.
func ResolverFromConfig(productIndex, postIndex string) func(string) (string, error) {
	return func(resourceType string) (string, error) {
		resourceType = strings.TrimSpace(strings.ToLower(resourceType))
		switch resourceType {
		case models.SearchResourceProduct:
			if productIndex == "" {
				return "", errors.New("product index not configured")
			}
			return productIndex, nil
		case models.SearchResourcePost:
			if postIndex == "" {
				return "", errors.New("content index not configured")
			}
			return postIndex, nil
		default:
			return "", fmt.Errorf("unknown resource type: %s", resourceType)
		}
	}
}
