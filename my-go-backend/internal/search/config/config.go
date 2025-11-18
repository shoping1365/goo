package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config نگهداری تنظیمات سرویس جستجو
// بیشتر مقادیر از متغیرهای محیطی خوانده می‌شوند تا در دیپلوی منعطف باشد.
type Config struct {
	HTTPAddr       string
	MetricsAddr    string
	IndexDir       string
	EnableMetrics  bool
	AllowedIndexes []string
	SuggestFields  map[string]string
}

// Load خواندن تنظیمات و اعتبارسنجی آن‌ها.
func Load() (Config, error) {
	cfg := Config{
		HTTPAddr:      getEnv("SEARCH_HTTP_ADDR", ":8085"),
		MetricsAddr:   getEnv("SEARCH_METRICS_ADDR", ":9105"),
		IndexDir:      filepath.Clean(getEnv("SEARCH_INDEX_DIR", "var/search/indexes")),
		EnableMetrics: getEnv("SEARCH_ENABLE_METRICS", "1") != "0",
		SuggestFields: map[string]string{},
	}

	indexes := strings.TrimSpace(getEnv("SEARCH_ALLOWED_INDEXES", "products,content,support"))
	if indexes != "" {
		parts := strings.Split(indexes, ",")
		seen := make(map[string]struct{})
		for _, part := range parts {
			name := strings.TrimSpace(part)
			if name == "" {
				continue
			}
			if _, ok := seen[name]; ok {
				continue
			}
			seen[name] = struct{}{}
			cfg.AllowedIndexes = append(cfg.AllowedIndexes, name)
		}
	}

	if len(cfg.AllowedIndexes) == 0 {
		return cfg, fmt.Errorf("SEARCH_ALLOWED_INDEXES cannot be empty")
	}

	if cfg.IndexDir == "" {
		return cfg, fmt.Errorf("SEARCH_INDEX_DIR cannot be empty")
	}

	if _, err := os.Stat(cfg.IndexDir); err != nil {
		if os.IsNotExist(err) {
			return cfg, fmt.Errorf("search index dir %s does not exist; run search_bootstrap first", cfg.IndexDir)
		}
		return cfg, fmt.Errorf("check search index dir: %w", err)
	}

	for _, idx := range cfg.AllowedIndexes {
		key := fmt.Sprintf("SEARCH_SUGGEST_FIELD_%s", strings.ToUpper(idx))
		field := strings.TrimSpace(os.Getenv(key))
		if field == "" {
			field = defaultSuggestField(idx)
		}
		if field != "" {
			cfg.SuggestFields[idx] = field
		}
	}

	return cfg, nil
}

func getEnv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}

func defaultSuggestField(index string) string {
	switch strings.ToLower(index) {
	case "products":
		return "title"
	case "content":
		return "title"
	case "support":
		return "question"
	case "categories":
		return "title"
	default:
		return ""
	}
}
