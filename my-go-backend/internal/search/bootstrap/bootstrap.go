package bootstrap

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
)

type Options struct {
	IndexDir string
	DataDir  string
	Force    bool
	Indexes  []string
}

type indexSpec struct {
	name     string
	dataPath string
}

func Run(opts Options) error {
	indexDir := filepath.Clean(strings.TrimSpace(opts.IndexDir))
	if indexDir == "" {
		return fmt.Errorf("index directory is required")
	}

	dataDir := strings.TrimSpace(opts.DataDir)
	if dataDir == "" {
		dataDir = "data/search"
	}
	dataDir = filepath.Clean(dataDir)

	if err := os.MkdirAll(indexDir, 0o755); err != nil {
		return fmt.Errorf("creating index directory: %w", err)
	}

	available := availableSpecs(dataDir)
	order := normalizeIndexes(opts.Indexes)
	if len(order) == 0 {
		order = normalizeIndexes([]string{"products", "content", "support"})
	}

	specs := make([]indexSpec, 0, len(order))
	for _, idx := range order {
		spec, ok := available[idx]
		if !ok {
			fmt.Printf("⚠️  unsupported index %s requested; skipping\n", idx)
			continue
		}
		specs = append(specs, spec)
	}

	if len(specs) == 0 {
		return fmt.Errorf("no supported indexes selected for bootstrap")
	}

	for _, spec := range specs {
		if err := buildIndex(indexDir, spec, opts.Force); err != nil {
			return err
		}
	}

	return nil
}

func normalizeIndexes(list []string) []string {
	if len(list) == 0 {
		return nil
	}
	seen := make(map[string]struct{})
	cleaned := make([]string, 0, len(list))
	for _, raw := range list {
		name := strings.TrimSpace(strings.ToLower(raw))
		if name == "" {
			continue
		}
		if _, exists := seen[name]; exists {
			continue
		}
		seen[name] = struct{}{}
		cleaned = append(cleaned, name)
	}
	return cleaned
}

func availableSpecs(dataDir string) map[string]indexSpec {
	return map[string]indexSpec{
		"products":   {name: "products", dataPath: filepath.Join(dataDir, "products_sample.ndjson")},
		"content":    {name: "content", dataPath: filepath.Join(dataDir, "content_sample.ndjson")},
		"support":    {name: "support", dataPath: filepath.Join(dataDir, "support_sample.ndjson")},
		"categories": {name: "categories", dataPath: filepath.Join(dataDir, "categories_sample.ndjson")},
	}
}

func buildIndex(indexDir string, spec indexSpec, force bool) error {
	builder, ok := mappingBuilders()[spec.name]
	if !ok {
		return fmt.Errorf("no mapping registered for index %s", spec.name)
	}

	indexPath := filepath.Join(indexDir, spec.name+".bleve")
	if _, err := os.Stat(indexPath); err == nil {
		if !force {
			fmt.Printf("ℹ️  index %s already exists, skipping (use --force to rebuild)\n", spec.name)
			return nil
		}
		if err := os.RemoveAll(indexPath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("remove existing index %s: %w", spec.name, err)
		}
	}

	idx, err := createIndex(indexPath, builder)
	if err != nil {
		return err
	}
	defer idx.Close()

	file, err := os.Open(spec.dataPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("⚠️  sample data not found for %s at %s, skipping\n", spec.name, spec.dataPath)
			return nil
		}
		return fmt.Errorf("open sample data for %s: %w", spec.name, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	batch := idx.NewBatch()
	batchSize := 0
	total := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var doc map[string]any
		if err := json.Unmarshal([]byte(line), &doc); err != nil {
			return fmt.Errorf("parse document for %s: %w", spec.name, err)
		}
		docID, normalized, err := normalizeDocument(spec.name, doc)
		if err != nil {
			return err
		}
		if err := batch.Index(docID, normalized); err != nil {
			return fmt.Errorf("index document %s in %s: %w", docID, spec.name, err)
		}
		batchSize++
		total++
		if batchSize >= 100 {
			if err := idx.Batch(batch); err != nil {
				return fmt.Errorf("commit batch for %s: %w", spec.name, err)
			}
			batch = idx.NewBatch()
			batchSize = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read sample data for %s: %w", spec.name, err)
	}

	if batchSize > 0 {
		if err := idx.Batch(batch); err != nil {
			return fmt.Errorf("commit final batch for %s: %w", spec.name, err)
		}
	}

	fmt.Printf("✅ indexed %d documents for %s\n", total, spec.name)
	return nil
}

func createIndex(path string, builder func() (*mapping.IndexMappingImpl, error)) (bleve.Index, error) {
	mapping, err := builder()
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("prepare index directory %s: %w", filepath.Dir(path), err)
	}
	idx, err := bleve.New(path, mapping)
	if err != nil {
		return nil, fmt.Errorf("create index at %s: %w", path, err)
	}
	return idx, nil
}

func normalizeDocument(index string, doc map[string]any) (string, map[string]any, error) {
	idRaw, ok := doc["id"]
	if !ok {
		return "", nil, fmt.Errorf("document in %s missing id field", index)
	}
	id := strings.TrimSpace(fmt.Sprint(idRaw))
	if id == "" {
		return "", nil, fmt.Errorf("document in %s has empty id", index)
	}

	doc["id"] = id
	doc["_type"] = index
	delete(doc, "vector_embedding")

	for key, value := range doc {
		arr, ok := value.([]any)
		if !ok {
			continue
		}
		strValues := make([]string, 0, len(arr))
		convertible := true
		for _, item := range arr {
			if item == nil {
				continue
			}
			strVal := strings.TrimSpace(fmt.Sprint(item))
			if strVal == "" {
				continue
			}
			strValues = append(strValues, strVal)
			if _, ok := item.(string); !ok {
				convertible = false
			}
		}
		if convertible {
			doc[key] = strValues
		}
	}

	return id, doc, nil
}

func mappingBuilders() map[string]func() (*mapping.IndexMappingImpl, error) {
	return map[string]func() (*mapping.IndexMappingImpl, error){
		"products":   productsMapping,
		"content":    contentMapping,
		"support":    supportMapping,
		"categories": categoriesMapping,
	}
}

func productsMapping() (*mapping.IndexMappingImpl, error) {
	textField := mapping.NewTextFieldMapping()
	textField.Store = true
	textField.IncludeInAll = true
	textField.IncludeTermVectors = true

	keywordField := mapping.NewTextFieldMapping()
	keywordField.Store = true
	keywordField.Analyzer = "keyword"

	numericField := mapping.NewNumericFieldMapping()
	numericField.Store = true
	numericField.DocValues = true

	boolField := mapping.NewBooleanFieldMapping()
	boolField.Store = true

	dateField := mapping.NewDateTimeFieldMapping()
	dateField.Store = true
	dateField.DocValues = true

	doc := mapping.NewDocumentMapping()
	doc.AddFieldMappingsAt("id", keywordField)
	doc.AddFieldMappingsAt("sku", keywordField)
	doc.AddFieldMappingsAt("title", textField)
	doc.AddFieldMappingsAt("description", textField)
	doc.AddFieldMappingsAt("categories", textField)
	doc.AddFieldMappingsAt("brand", keywordField)
	doc.AddFieldMappingsAt("price", numericField)
	doc.AddFieldMappingsAt("discount_percent", numericField)
	doc.AddFieldMappingsAt("inventory", numericField)
	doc.AddFieldMappingsAt("is_available", boolField)
	doc.AddFieldMappingsAt("rating", numericField)
	doc.AddFieldMappingsAt("sales_rank", numericField)
	doc.AddFieldMappingsAt("created_at", dateField)
	doc.AddFieldMappingsAt("updated_at", dateField)
	doc.AddFieldMappingsAt("lang", keywordField)

	idxMapping := bleve.NewIndexMapping()
	idxMapping.TypeField = "_type"
	idxMapping.DefaultType = "doc"
	idxMapping.DefaultAnalyzer = "standard"
	idxMapping.StoreDynamic = true
	idxMapping.IndexDynamic = true
	idxMapping.DefaultMapping = doc
	return idxMapping, nil
}

func contentMapping() (*mapping.IndexMappingImpl, error) {
	textField := mapping.NewTextFieldMapping()
	textField.Store = true
	textField.IncludeInAll = true
	textField.IncludeTermVectors = true

	keywordField := mapping.NewTextFieldMapping()
	keywordField.Store = true
	keywordField.Analyzer = "keyword"

	numericField := mapping.NewNumericFieldMapping()
	numericField.Store = true
	numericField.DocValues = true

	dateField := mapping.NewDateTimeFieldMapping()
	dateField.Store = true
	dateField.DocValues = true

	doc := mapping.NewDocumentMapping()
	doc.AddFieldMappingsAt("id", keywordField)
	doc.AddFieldMappingsAt("slug", keywordField)
	doc.AddFieldMappingsAt("title", textField)
	doc.AddFieldMappingsAt("excerpt", textField)
	doc.AddFieldMappingsAt("body", textField)
	doc.AddFieldMappingsAt("tags", textField)
	doc.AddFieldMappingsAt("author", keywordField)
	doc.AddFieldMappingsAt("categories", textField)
	doc.AddFieldMappingsAt("published_at", dateField)
	doc.AddFieldMappingsAt("lang", keywordField)
	doc.AddFieldMappingsAt("reading_time", numericField)

	idxMapping := bleve.NewIndexMapping()
	idxMapping.TypeField = "_type"
	idxMapping.DefaultType = "doc"
	idxMapping.DefaultAnalyzer = "standard"
	idxMapping.StoreDynamic = true
	idxMapping.IndexDynamic = true
	idxMapping.DefaultMapping = doc
	return idxMapping, nil
}

func supportMapping() (*mapping.IndexMappingImpl, error) {
	textField := mapping.NewTextFieldMapping()
	textField.Store = true
	textField.IncludeInAll = true
	textField.IncludeTermVectors = true

	keywordField := mapping.NewTextFieldMapping()
	keywordField.Store = true
	keywordField.Analyzer = "keyword"

	numericField := mapping.NewNumericFieldMapping()
	numericField.Store = true
	numericField.DocValues = true

	dateField := mapping.NewDateTimeFieldMapping()
	dateField.Store = true
	dateField.DocValues = true

	doc := mapping.NewDocumentMapping()
	doc.AddFieldMappingsAt("id", keywordField)
	doc.AddFieldMappingsAt("question", textField)
	doc.AddFieldMappingsAt("answer", textField)
	doc.AddFieldMappingsAt("product_ids", keywordField)
	doc.AddFieldMappingsAt("category", keywordField)
	doc.AddFieldMappingsAt("tags", textField)
	doc.AddFieldMappingsAt("updated_at", dateField)
	doc.AddFieldMappingsAt("lang", keywordField)
	doc.AddFieldMappingsAt("confidence_score", numericField)

	idxMapping := bleve.NewIndexMapping()
	idxMapping.TypeField = "_type"
	idxMapping.DefaultType = "doc"
	idxMapping.DefaultAnalyzer = "standard"
	idxMapping.StoreDynamic = true
	idxMapping.IndexDynamic = true
	idxMapping.DefaultMapping = doc
	return idxMapping, nil
}

func categoriesMapping() (*mapping.IndexMappingImpl, error) {
	textField := mapping.NewTextFieldMapping()
	textField.Store = true
	textField.IncludeInAll = true
	textField.IncludeTermVectors = true

	keywordField := mapping.NewTextFieldMapping()
	keywordField.Store = true
	keywordField.Analyzer = "keyword"

	doc := mapping.NewDocumentMapping()
	doc.AddFieldMappingsAt("id", keywordField)
	doc.AddFieldMappingsAt("title", textField)
	doc.AddFieldMappingsAt("name", textField)
	doc.AddFieldMappingsAt("slug", keywordField)
	doc.AddFieldMappingsAt("description", textField)
	doc.AddFieldMappingsAt("breadcrumbs", textField)
	doc.AddFieldMappingsAt("parent_id", keywordField)
	doc.AddFieldMappingsAt("lang", keywordField)

	idxMapping := bleve.NewIndexMapping()
	idxMapping.TypeField = "_type"
	idxMapping.DefaultType = "doc"
	idxMapping.DefaultAnalyzer = "standard"
	idxMapping.StoreDynamic = true
	idxMapping.IndexDynamic = true
	idxMapping.DefaultMapping = doc
	return idxMapping, nil
}
