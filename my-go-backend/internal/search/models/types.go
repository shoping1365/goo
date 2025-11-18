package models

// Filter تعریف فیلترهای اعمال شده روی جستجو
// operator با مقدارهایی مثل eq, gt, lt, in پشتیبانی می‌شود.
type Filter struct {
	Field    string   `json:"field"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

// Sort مشخص‌کننده‌ی نحوه‌ی مرتب‌سازی نتایج
// اگر descending true باشد، ترتیب نزولی خواهد بود.
type Sort struct {
	Field      string `json:"field"`
	Descending bool   `json:"descending"`
}

// PageCursor برای صفحه‌بندی با cursor-base
// اگر cursor خالی باشد، از ابتدای نتایج شروع می‌شود.
type PageCursor struct {
	Size   int    `json:"size"`
	Cursor string `json:"cursor"`
}

// VectorQuery پارامترهای جستجوی برداری را نگه می‌دارد.
type VectorQuery struct {
	Enabled   bool      `json:"enabled"`
	Embedding []float32 `json:"embedding"`
	TopK      int       `json:"top_k"`
	Weight    float64   `json:"weight"`
}

// ClientContext اطلاعات تکمیلی درباره‌ی کاربر یا نشست را حمل می‌کند.
type ClientContext struct {
	UserID    string            `json:"user_id"`
	SessionID string            `json:"session_id"`
	Locale    string            `json:"locale"`
	Tags      map[string]string `json:"tags"`
}

// SearchRequest ساختار ورودی برای endpoint جستجو
// index باید در لیست اندیس‌های مجاز قرار داشته باشد.
type SearchRequest struct {
	Index   string        `json:"index"`
	Query   string        `json:"query"`
	Filters []Filter      `json:"filters"`
	Sorts   []Sort        `json:"sorts"`
	Page    PageCursor    `json:"page"`
	Vector  VectorQuery   `json:"vector"`
	Client  ClientContext `json:"client"`
}

// Hit نتیجه‌ی تک‌ آیتم.
// fields برای نگهداری فیلدهای برگشتی (stringified) استفاده می‌شود.
type Hit struct {
	ID        string              `json:"id"`
	Score     float64             `json:"score"`
	Fields    map[string]string   `json:"fields"`
	Highlight map[string][]string `json:"highlight"`
	Scores    map[string]float64  `json:"scores"`
}

// FacetBucket شامل مقدار و تعداد هر bucket است.
type FacetBucket struct {
	Value string `json:"value"`
	Count int64  `json:"count"`
}

// Facet لیست bucketها را برای یک فیلد نگه می‌دارد.
type Facet struct {
	Field   string        `json:"field"`
	Buckets []FacetBucket `json:"buckets"`
}

// Diagnostic اطلاعات دیباگ/تحلیلی از موتور جستجو است.
type Diagnostic struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// SearchResponse پاسخ اصلی endpoint جستجو.
type SearchResponse struct {
	Hits             []Hit        `json:"hits"`
	NextCursor       string       `json:"next_cursor"`
	Facets           []Facet      `json:"facets"`
	Diagnostics      []Diagnostic `json:"diagnostics"`
	ProcessingTimeMs float64      `json:"processing_time_ms"`
}

// SuggestRequest پارامترهای endpoint پیشنهاد را نگه می‌دارد.
type SuggestRequest struct {
	Index  string `json:"index"`
	Prefix string `json:"prefix"`
	Limit  int    `json:"limit"`
	Locale string `json:"locale"`
}

// Suggestion خروجی پیشنهاد.
type Suggestion struct {
	Text  string  `json:"text"`
	Score float64 `json:"score"`
	Type  string  `json:"type"`
}

// SuggestResponse جواب endpoint پیشنهاد است.
type SuggestResponse struct {
	Suggestions []Suggestion `json:"suggestions"`
}

// HealthStatus وضعیت سرویس یا موتور جستجو را نمایش می‌دهد.
type HealthStatus string

const (
	// HealthStatusUnknown وقتی است که وضعیت مشخص نیست یا ارتباط شکست خورده.
	HealthStatusUnknown HealthStatus = "unknown"
	// HealthStatusHealthy وضعیت سالم.
	HealthStatusHealthy HealthStatus = "healthy"
	// HealthStatusDegraded وضعیت نیمه پایدار.
	HealthStatusDegraded HealthStatus = "degraded"
	// HealthStatusDown وضعیت از کار افتاده.
	HealthStatusDown HealthStatus = "down"
)

// HealthResponse ساختار وضعیت سلامت.
type HealthResponse struct {
	Status  HealthStatus `json:"status"`
	Details string       `json:"details"`
}
