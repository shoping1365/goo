package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	searchmodels "my-go-backend/internal/search/models"
	searchservice "my-go-backend/internal/search/service"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// SearchHandler مدیریت درخواست‌های جستجوی عمومی مبتنی بر موتور Bleve است.
type SearchHandler struct {
	svc          *searchservice.Service
	allowed      []string
	allowedSet   map[string]struct{}
	indexAliases map[string][]string
}

// NewSearchHandler ایجاد نمونه جدید از handler.
func NewSearchHandler(svc *searchservice.Service, allowedIndexes []string) *SearchHandler {
	set := make(map[string]struct{}, len(allowedIndexes))
	cleaned := make([]string, 0, len(allowedIndexes))
	for _, idx := range allowedIndexes {
		name := strings.TrimSpace(strings.ToLower(idx))
		if name == "" {
			continue
		}
		if _, exists := set[name]; exists {
			continue
		}
		set[name] = struct{}{}
		cleaned = append(cleaned, name)
	}

	aliases := map[string][]string{
		"all": cleaned,
	}
	aliases["products"] = filterExisting(set, "products")
	aliases["product"] = aliases["products"]
	aliases["posts"] = filterExisting(set, "content")
	aliases["post"] = aliases["posts"]
	aliases["content"] = filterExisting(set, "content")
	aliases["support"] = filterExisting(set, "support")
	aliases["faqs"] = aliases["support"]
	aliases["help"] = aliases["support"]
	aliases["brands"] = filterExisting(set, "brands")
	aliases["categories"] = filterExisting(set, "categories")

	return &SearchHandler{
		svc:          svc,
		allowed:      cleaned,
		allowedSet:   set,
		indexAliases: aliases,
	}
}

// SearchResult نتیجه جستجو
type SearchResult struct {
	Type        string      `json:"type"`
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	Image       string      `json:"image"`
	Price       *float64    `json:"price"`
	Score       float64     `json:"score"`
	Data        interface{} `json:"data"`
}

// SearchResponse پاسخ API جستجو
type SearchResponse struct {
	Success bool           `json:"success"`
	Data    []SearchResult `json:"data"`
	Total   int64          `json:"total"`
	Query   string         `json:"query"`
}

// GlobalSearch جستجوی عمومی در همه بخش‌های سایت با موتور Bleve.
func (h *SearchHandler) GlobalSearch(c *gin.Context) {
	query := strings.TrimSpace(c.Query("q"))
	if query == "" {
		c.JSON(http.StatusBadRequest, utils.New("INVALID_QUERY", "عبارت جستجو الزامی است", ""))
		return
	}

	limit := 20
	if limitStr := strings.TrimSpace(c.Query("limit")); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	indexes := h.resolveIndexes(c.Query("type"))
	if len(indexes) == 0 {
		c.JSON(http.StatusOK, SearchResponse{Success: true, Data: []SearchResult{}, Total: 0, Query: query})
		return
	}

	results := make([]SearchResult, 0, limit*len(indexes))
	for _, idx := range indexes {
		reqCtx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		resp, err := h.svc.Search(reqCtx, searchmodels.SearchRequest{
			Index: idx,
			Query: query,
			Page: searchmodels.PageCursor{
				Size:   limit,
				Cursor: c.Query("cursor"),
			},
		})
		cancel()
		if err != nil {
			h.handleSearchError(c, err)
			return
		}

		for _, hit := range resp.Hits {
			results = append(results, h.mapHitToResult(idx, hit))
		}
	}

	sort.SliceStable(results, func(i, j int) bool { return results[i].Score > results[j].Score })
	if len(results) > limit {
		results = results[:limit]
	}

	c.JSON(http.StatusOK, SearchResponse{
		Success: true,
		Data:    results,
		Total:   int64(len(results)),
		Query:   query,
	})
}

// SearchSuggestions پیشنهادات جستجو
func (h *SearchHandler) SearchSuggestions(c *gin.Context) {
	query := strings.TrimSpace(c.Query("q"))
	if len(query) < 1 {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": []string{}})
		return
	}

	limit := 10
	if limitStr := strings.TrimSpace(c.Query("limit")); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 20 {
			limit = l
		}
	}

	indexes := h.resolveIndexes(c.Query("type"))
	if len(indexes) == 0 {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": []string{}})
		return
	}

	unique := make(map[string]struct{})
	suggestions := make([]string, 0, limit)

	for _, idx := range indexes {
		reqCtx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
		resp, err := h.svc.Suggest(reqCtx, searchmodels.SuggestRequest{
			Index:  idx,
			Prefix: query,
			Limit:  limit,
		})
		cancel()
		if err != nil {
			h.handleSearchError(c, err)
			return
		}

		for _, suggestion := range resp.Suggestions {
			text := strings.TrimSpace(suggestion.Text)
			if text == "" {
				continue
			}
			if _, exists := unique[text]; exists {
				continue
			}
			unique[text] = struct{}{}
			suggestions = append(suggestions, text)
			if len(suggestions) >= limit {
				break
			}
		}
		if len(suggestions) >= limit {
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    suggestions,
	})
}

func (h *SearchHandler) handleSearchError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, searchservice.ErrEmptyQuery()):
		c.JSON(http.StatusBadRequest, utils.New("INVALID_QUERY", "عبارت جستجو الزامی است", err.Error()))
	case errors.Is(err, searchservice.ErrInvalidLimit()):
		c.JSON(http.StatusBadRequest, utils.New("INVALID_LIMIT", "محدوده نتایج نامعتبر است", err.Error()))
	default:
		c.JSON(http.StatusInternalServerError, utils.New("SEARCH_ERROR", "خطا در اجرای جستجو", err.Error()))
	}
}

func (h *SearchHandler) resolveIndexes(kind string) []string {
	kind = strings.TrimSpace(strings.ToLower(kind))
	if kind == "" || kind == "all" {
		return h.allowed
	}

	if indexes, ok := h.indexAliases[kind]; ok && len(indexes) > 0 {
		return indexes
	}

	if _, ok := h.allowedSet[kind]; ok {
		return []string{kind}
	}

	return nil
}

func (h *SearchHandler) mapHitToResult(index string, hit searchmodels.Hit) SearchResult {
	fields := make(map[string]string, len(hit.Fields))
	for k, v := range hit.Fields {
		fields[k] = v
	}

	result := SearchResult{
		Type:  h.mapIndexToType(index),
		Score: hit.Score,
		Data:  fields,
	}

	if id, err := strconv.ParseUint(hit.ID, 10, 64); err == nil {
		result.ID = uint(id)
	}

	// Debug: چاپ تمام فیلدها برای دیباگ
	if index == "products" {
		log.Printf("Product %s fields: %+v", hit.ID, fields)
	}

	result.Title = firstNonEmpty(fields["title"], fields["name"], fields["question"], hit.ID)
	result.Description = firstNonEmpty(fields["description"], fields["excerpt"], fields["answer"], fields["body"])
	result.Image = firstNonEmpty(fields["image"], fields["image_url"], fields["thumbnail"], fields["featured_image"], fields["cover"], fields["main_image"], fields["media"], fields["imageUrl"])
	result.URL = h.resolveURL(result.Type, fields, hit.ID)

	// چک کردن فیلدهای مختلف قیمت
	priceFields := []string{
		"price",
		"final_price",
		"sale_price",
		"price_current",
		"price_sale",
		"current_price",
		"selling_price",
		"discounted_price",
	}

	for _, field := range priceFields {
		if priceStr := fields[field]; priceStr != "" {
			if price, ok := parsePrice(priceStr); ok && price > 0 {
				result.Price = &price
				break
			}
		}
	}

	return result
}

func (h *SearchHandler) mapIndexToType(index string) string {
	switch strings.ToLower(index) {
	case "products":
		return "product"
	case "content":
		return "post"
	case "support":
		return "support"
	case "brands":
		return "brand"
	case "categories":
		return "category"
	default:
		return index
	}
}

func (h *SearchHandler) resolveURL(resultType string, fields map[string]string, fallbackID string) string {
	if url := firstNonEmpty(fields["url"], fields["link"], fields["permalink"]); url != "" {
		return url
	}

	slug := fields["slug"]
	sku := fields["sku"]
	switch resultType {
	case "product":
		trimmedSKU := strings.TrimSpace(sku)
		trimmedSlug := strings.TrimSpace(slug)
		switch {
		case trimmedSKU != "" && trimmedSlug != "":
			return "/product/sku-" + trimmedSKU + "/" + trimmedSlug
		case trimmedSKU != "":
			return "/product/sku-" + trimmedSKU
		case trimmedSlug != "":
			return "/product/" + trimmedSlug
		case fallbackID != "":
			return "/product?id=" + fallbackID
		}
	case "post":
		if slug != "" {
			return "/blog/" + slug
		}
	case "category":
		if slug != "" {
			return "/category/" + slug
		}
	case "brand":
		if slug != "" {
			return "/brand/" + slug
		}
	case "support":
		if slug != "" {
			return "/support/" + slug
		}
	}

	return "/search/" + fallbackID
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

func parsePrice(raw string) (float64, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return 0, false
	}

	filtered := make([]rune, 0, len(raw))
	for _, r := range raw {
		if (r >= '0' && r <= '9') || r == '.' {
			filtered = append(filtered, r)
		}
	}

	if len(filtered) == 0 {
		return 0, false
	}

	val, err := strconv.ParseFloat(string(filtered), 64)
	if err != nil {
		return 0, false
	}
	return val, true
}

func filterExisting(set map[string]struct{}, idx string) []string {
	key := strings.ToLower(strings.TrimSpace(idx))
	if key == "" {
		return nil
	}
	if _, ok := set[key]; ok {
		return []string{key}
	}
	return nil
}

// PopularSearchesHandler برگرداندن جستجوهای پرطرفدار
func (h *SearchHandler) PopularSearchesHandler(c *gin.Context) {
	// TODO: این داده‌ها باید از دیتابیس یا cache واقعی بیاد
	// می‌تونی یه جدول search_logs داشته باشی که هر جستجو رو لاگ کنه
	// و بعد با GROUP BY و COUNT محبوب‌ترین‌ها رو برگردونی

	popularSearches := []map[string]interface{}{
		{"query": "گوشی سامسونگ", "count": 1250, "context": "موبایل"},
		{"query": "لپ تاپ ایسوس", "count": 980, "context": "کامپیوتر"},
		{"query": "هدفون بلوتوثی", "count": 850, "context": "صوتی"},
		{"query": "ساعت هوشمند", "count": 720, "context": "پوشیدنی"},
		{"query": "کیبورد گیمینگ", "count": 650, "context": "جانبی"},
		{"query": "مانیتور", "count": 580, "context": "نمایشگر"},
		{"query": "موس بی سیم", "count": 520, "context": "جانبی"},
		{"query": "شارژر فست", "count": 480, "context": "لوازم جانبی"},
		{"query": "پاوربانک", "count": 450, "context": "لوازم جانبی"},
		{"query": "کابل تایپ سی", "count": 420, "context": "لوازم جانبی"},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    popularSearches,
	})
}
