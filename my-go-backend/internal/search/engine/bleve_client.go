package engine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"my-go-backend/internal/search/models"

	"github.com/blevesearch/bleve/v2"
	bleveSearch "github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/query"
)

// Client defines behaviour expected by the service layer.
type Client interface {
	Search(ctx context.Context, req models.SearchRequest) (models.SearchResponse, error)
	Suggest(ctx context.Context, req models.SuggestRequest) (models.SuggestResponse, error)
	Health(ctx context.Context) (models.HealthResponse, error)
	IndexDocument(ctx context.Context, index string, id string, doc map[string]any) error
	DeleteDocument(ctx context.Context, index string, id string) error
}

// Options configures the Bleve client.
type Options struct {
	Directory      string
	AllowedIndexes []string
	SuggestFields  map[string]string
}

// BleveClient implements the search Client using local Bleve indexes.
type BleveClient struct {
	mu            sync.RWMutex
	indexes       map[string]bleve.Index
	allowed       map[string]struct{}
	suggestFields map[string]string
}

// NewBleveClient opens the configured indexes from disk.
func NewBleveClient(opts Options) (*BleveClient, error) {
	if strings.TrimSpace(opts.Directory) == "" {
		return nil, errors.New("index directory is required")
	}

	allowed := make(map[string]struct{})
	for _, idx := range opts.AllowedIndexes {
		name := strings.TrimSpace(idx)
		if name == "" {
			continue
		}
		allowed[name] = struct{}{}
	}
	if len(allowed) == 0 {
		return nil, errors.New("no search indexes configured")
	}

	indexes := make(map[string]bleve.Index, len(allowed))
	for name := range allowed {
		idxPath := filepath.Join(opts.Directory, name+".bleve")
		handle, err := bleve.Open(idxPath)
		if err != nil {
			if errors.Is(err, bleve.ErrorIndexPathDoesNotExist) {
				return nil, fmt.Errorf("index %s not found at %s (run search_bootstrap)", name, idxPath)
			}
			return nil, fmt.Errorf("open index %s: %w", name, err)
		}
		indexes[name] = handle
	}

	suggest := make(map[string]string, len(allowed))
	for name := range allowed {
		if field := strings.TrimSpace(opts.SuggestFields[name]); field != "" {
			suggest[name] = field
			continue
		}
		if field := defaultSuggestField(name); field != "" {
			suggest[name] = field
		}
	}

	return &BleveClient{
		indexes:       indexes,
		allowed:       allowed,
		suggestFields: suggest,
	}, nil
}

// Close releases all index handles.
func (c *BleveClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var errs []error
	for name, idx := range c.indexes {
		if err := idx.Close(); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", name, err))
		}
	}
	c.indexes = map[string]bleve.Index{}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

// Search executes a text query with optional filters and sorting.
func (c *BleveClient) Search(ctx context.Context, req models.SearchRequest) (models.SearchResponse, error) {
	indexName := strings.TrimSpace(req.Index)
	if indexName == "" {
		return models.SearchResponse{}, errors.New("index is required")
	}

	idx, err := c.getIndex(indexName)
	if err != nil {
		return models.SearchResponse{}, err
	}

	size := req.Page.Size
	if size <= 0 {
		size = 20
	}
	if size > 500 {
		size = 500
	}

	offset := 0
	if req.Page.Cursor != "" {
		if val, convErr := strconv.Atoi(req.Page.Cursor); convErr == nil && val >= 0 {
			offset = val
		}
	}

	q, err := buildSearchQuery(req)
	if err != nil {
		return models.SearchResponse{}, err
	}

	searchReq := bleve.NewSearchRequestOptions(q, size, offset, false)
	searchReq.Fields = []string{"*"}
	searchReq.Highlight = bleve.NewHighlightWithStyle("html")
	if sorts := buildSortOrders(req.Sorts); len(sorts) > 0 {
		searchReq.SortBy(sorts)
	}

	res, err := idx.SearchInContext(ctx, searchReq)
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("bleve search failed: %w", err)
	}

	response := models.SearchResponse{
		Hits:        collectHits(res.Hits),
		Facets:      nil,
		Diagnostics: collectDiagnostics(res),
	}

	if res.Took > 0 {
		response.ProcessingTimeMs = float64(res.Took) / float64(time.Millisecond)
	}

	if total := int(res.Total); total > offset+size {
		response.NextCursor = strconv.Itoa(offset + size)
	}

	return response, nil
}

// Suggest returns lightweight autocomplete suggestions based on a prefix.
func (c *BleveClient) Suggest(ctx context.Context, req models.SuggestRequest) (models.SuggestResponse, error) {
	indexName := strings.TrimSpace(req.Index)
	if indexName == "" {
		return models.SuggestResponse{}, errors.New("index is required")
	}

	idx, err := c.getIndex(indexName)
	if err != nil {
		return models.SuggestResponse{}, err
	}

	prefix := strings.TrimSpace(req.Prefix)
	if prefix == "" {
		return models.SuggestResponse{}, nil
	}

	limit := req.Limit
	if limit <= 0 {
		limit = 5
	}
	if limit > 20 {
		limit = 20
	}

	field := strings.TrimSpace(c.suggestFields[indexName])
	query := bleve.NewPrefixQuery(strings.ToLower(prefix))
	if field != "" {
		query.SetField(field)
	}

	searchReq := bleve.NewSearchRequestOptions(query, limit, 0, false)
	if field != "" {
		searchReq.Fields = []string{field}
	} else {
		searchReq.Fields = []string{"*"}
	}
	searchReq.SortBy([]string{"-_score"})

	res, err := idx.SearchInContext(ctx, searchReq)
	if err != nil {
		return models.SuggestResponse{}, fmt.Errorf("bleve suggest failed: %w", err)
	}

	suggestions := make([]models.Suggestion, 0, len(res.Hits))
	seen := make(map[string]struct{})
	for _, hit := range res.Hits {
		text := ""
		if field != "" {
			text = stringify(hit.Fields[field])
		}
		if text == "" {
			text = stringify(hit.Fields["title"])
		}
		if text == "" {
			text = stringify(hit.Fields["question"])
		}
		if text == "" {
			text = hit.ID
		}
		if _, exists := seen[text]; exists {
			continue
		}
		seen[text] = struct{}{}
		suggestions = append(suggestions, models.Suggestion{
			Text:  text,
			Score: hit.Score,
			Type:  indexName,
		})
	}

	return models.SuggestResponse{Suggestions: suggestions}, nil
}

// Health checks index handles and returns simple diagnostics.
func (c *BleveClient) Health(ctx context.Context) (models.HealthResponse, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.indexes) == 0 {
		return models.HealthResponse{Status: models.HealthStatusDown, Details: "no indexes loaded"}, nil
	}

	for name, idx := range c.indexes {
		if ctx.Err() != nil {
			return models.HealthResponse{Status: models.HealthStatusUnknown, Details: ctx.Err().Error()}, nil
		}
		if _, err := idx.DocCount(); err != nil {
			return models.HealthResponse{Status: models.HealthStatusDegraded, Details: fmt.Sprintf("index %s doc count failed: %v", name, err)}, nil
		}
	}

	return models.HealthResponse{Status: models.HealthStatusHealthy, Details: "bleve ready"}, nil
}

// IndexDocument درج یا به‌روزرسانی سند در ایندکس مشخص
func (c *BleveClient) IndexDocument(ctx context.Context, indexName, id string, doc map[string]any) error {
	if ctx != nil && ctx.Err() != nil {
		return ctx.Err()
	}

	idx, err := c.getIndex(indexName)
	if err != nil {
		return err
	}

	if err := idx.Index(id, doc); err != nil {
		return fmt.Errorf("bleve index %s doc %s: %w", indexName, id, err)
	}
	return nil
}

// DeleteDocument حذف سند از ایندکس مشخص
func (c *BleveClient) DeleteDocument(ctx context.Context, indexName, id string) error {
	if ctx != nil && ctx.Err() != nil {
		return ctx.Err()
	}

	idx, err := c.getIndex(indexName)
	if err != nil {
		return err
	}

	if err := idx.Delete(id); err != nil {
		msg := strings.ToLower(err.Error())
		if strings.Contains(msg, "not found") || strings.Contains(msg, "does not exist") {
			return nil
		}
		return fmt.Errorf("bleve delete %s doc %s: %w", indexName, id, err)
	}
	return nil
}

func (c *BleveClient) getIndex(name string) (bleve.Index, error) {
	c.mu.RLock()
	idx, ok := c.indexes[name]
	c.mu.RUnlock()
	if !ok {
		if len(c.indexes) == 0 {
			return nil, errors.New("no search indexes loaded")
		}
		return nil, fmt.Errorf("index '%s' is not available", name)
	}
	return idx, nil
}

func buildSearchQuery(req models.SearchRequest) (query.Query, error) {
	queryString := strings.TrimSpace(req.Query)
	if queryString == "" {
		return nil, errors.New("query is required")
	}

	mainQuery := bleve.NewQueryStringQuery(queryString)

	clauses := []query.Query{mainQuery}
	for _, filter := range req.Filters {
		if q := buildFilterQuery(filter); q != nil {
			clauses = append(clauses, q)
		}
	}

	if len(clauses) == 1 {
		return clauses[0], nil
	}
	return bleve.NewConjunctionQuery(clauses...), nil
}

func buildFilterQuery(f models.Filter) query.Query {
	field := strings.TrimSpace(f.Field)
	if field == "" || len(f.Values) == 0 {
		return nil
	}

	operator := strings.ToLower(strings.TrimSpace(f.Operator))
	values := f.Values

	switch operator {
	case "", "eq", "in":
		var disj []query.Query
		for _, value := range values {
			term := strings.TrimSpace(value)
			if term == "" {
				continue
			}
			q := bleve.NewMatchQuery(term)
			q.SetField(field)
			disj = append(disj, q)
		}
		if len(disj) == 0 {
			return nil
		}
		if len(disj) == 1 {
			return disj[0]
		}
		return bleve.NewDisjunctionQuery(disj...)
	case "gt", "gte", "lt", "lte":
		value := strings.TrimSpace(values[0])
		if q := buildNumericRangeQuery(field, operator, value); q != nil {
			return q
		}
		if q := buildDateRangeQuery(field, operator, value); q != nil {
			return q
		}
		return nil
	default:
		return nil
	}
}

func buildNumericRangeQuery(field, operator, raw string) query.Query {
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return nil
	}
	boolFalse := func() *bool {
		v := false
		return &v
	}
	boolTrue := func() *bool {
		v := true
		return &v
	}

	switch operator {
	case "gt":
		min := val
		q := bleve.NewNumericRangeInclusiveQuery(&min, nil, boolFalse(), nil)
		q.SetField(field)
		return q
	case "gte":
		min := val
		q := bleve.NewNumericRangeInclusiveQuery(&min, nil, boolTrue(), nil)
		q.SetField(field)
		return q
	case "lt":
		max := val
		q := bleve.NewNumericRangeInclusiveQuery(nil, &max, nil, boolFalse())
		q.SetField(field)
		return q
	case "lte":
		max := val
		q := bleve.NewNumericRangeInclusiveQuery(nil, &max, nil, boolTrue())
		q.SetField(field)
		return q
	default:
		return nil
	}
}

func buildDateRangeQuery(field, operator, raw string) query.Query {
	tm, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return nil
	}
	boolFalse := func() *bool {
		v := false
		return &v
	}
	boolTrue := func() *bool {
		v := true
		return &v
	}

	zero := time.Time{}

	switch operator {
	case "gt":
		q := bleve.NewDateRangeInclusiveQuery(tm, zero, boolFalse(), nil)
		q.SetField(field)
		return q
	case "gte":
		q := bleve.NewDateRangeInclusiveQuery(tm, zero, boolTrue(), nil)
		q.SetField(field)
		return q
	case "lt":
		q := bleve.NewDateRangeInclusiveQuery(zero, tm, nil, boolFalse())
		q.SetField(field)
		return q
	case "lte":
		q := bleve.NewDateRangeInclusiveQuery(zero, tm, nil, boolTrue())
		q.SetField(field)
		return q
	default:
		return nil
	}
}

func buildSortOrders(sorts []models.Sort) []string {
	if len(sorts) == 0 {
		return nil
	}
	out := make([]string, 0, len(sorts))
	for _, s := range sorts {
		field := strings.TrimSpace(s.Field)
		if field == "" {
			continue
		}
		if s.Descending {
			out = append(out, "-"+field)
		} else {
			out = append(out, field)
		}
	}
	return out
}

func collectHits(hits []*bleveSearch.DocumentMatch) []models.Hit {
	if len(hits) == 0 {
		return nil
	}

	out := make([]models.Hit, 0, len(hits))
	for _, hit := range hits {
		fields := make(map[string]string, len(hit.Fields))
		for k, v := range hit.Fields {
			fields[k] = stringify(v)
		}

		highlight := make(map[string][]string, len(hit.Fragments))
		for field, frags := range hit.Fragments {
			clean := make([]string, 0, len(frags))
			for _, frag := range frags {
				frag = strings.TrimSpace(frag)
				if frag != "" {
					clean = append(clean, frag)
				}
			}
			if len(clean) > 0 {
				highlight[field] = clean
			}
		}
		if len(highlight) == 0 {
			highlight = nil
		}

		scores := map[string]float64{"bleve": hit.Score}

		out = append(out, models.Hit{
			ID:        hit.ID,
			Score:     hit.Score,
			Fields:    fields,
			Highlight: highlight,
			Scores:    scores,
		})
	}

	return out
}

func collectDiagnostics(res *bleve.SearchResult) []models.Diagnostic {
	if res == nil {
		return nil
	}
	diags := []models.Diagnostic{
		{Name: "total_hits", Value: strconv.FormatUint(res.Total, 10)},
	}
	if res.MaxScore > 0 {
		diags = append(diags, models.Diagnostic{Name: "max_score", Value: fmt.Sprintf("%.4f", res.MaxScore)})
	}
	if res.Took > 0 {
		diags = append(diags, models.Diagnostic{Name: "took_ms", Value: fmt.Sprintf("%.2f", float64(res.Took)/float64(time.Millisecond))})
	}
	return diags
}

func stringify(value interface{}) string {
	switch v := value.(type) {
	case nil:
		return ""
	case string:
		return v
	case json.Number:
		return v.String()
	case fmt.Stringer:
		return v.String()
	case []interface{}:
		if len(v) == 0 {
			return ""
		}
		parts := make([]string, 0, len(v))
		for _, item := range v {
			parts = append(parts, stringify(item))
		}
		return strings.Join(parts, ", ")
	default:
		return fmt.Sprint(v)
	}
}

func defaultSuggestField(index string) string {
	switch strings.ToLower(index) {
	case "products":
		return "title"
	case "content":
		return "title"
	case "support":
		return "question"
	default:
		return ""
	}
}
