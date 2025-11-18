package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	searchbootstrap "my-go-backend/internal/search/bootstrap"
)

func main() {
	var (
		indexDir string
		dataDir  string
		force    bool
	)

	flag.StringVar(&indexDir, "index-dir", getEnv("SEARCH_INDEX_DIR", "var/search/indexes"), "Directory to store Bleve indexes")
	flag.StringVar(&dataDir, "data-dir", getEnv("SEARCH_SAMPLE_DATA_DIR", "data/search"), "Directory containing sample NDJSON data")
	flag.BoolVar(&force, "force", getEnv("SEARCH_BOOTSTRAP_FORCE", "0") == "1", "Remove existing indexes before rebuilding")
	flag.Parse()

	indexes := parseIndexes(getEnv("SEARCH_ALLOWED_INDEXES", ""))

	if err := searchbootstrap.Run(searchbootstrap.Options{
		IndexDir: filepath.Clean(indexDir),
		DataDir:  filepath.Clean(dataDir),
		Force:    force,
		Indexes:  indexes,
	}); err != nil {
		fatal("search bootstrap", err)
	}

	fmt.Println("✅ local search indexes ready")
}

func getEnv(key, fallback string) string {
	if val := strings.TrimSpace(os.Getenv(key)); val != "" {
		return val
	}
	return fallback
}

func parseIndexes(raw string) []string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	indexes := make([]string, 0, len(parts))
	seen := make(map[string]struct{})
	for _, part := range parts {
		name := strings.TrimSpace(part)
		if name == "" {
			continue
		}
		lower := strings.ToLower(name)
		if _, ok := seen[lower]; ok {
			continue
		}
		seen[lower] = struct{}{}
		indexes = append(indexes, lower)
	}
	return indexes
}

func fatal(action string, err error) {
	fmt.Fprintf(os.Stderr, "❌ %s: %v\n", action, err)
	os.Exit(1)
}
