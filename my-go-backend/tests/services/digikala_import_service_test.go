package services_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/tests/mocks"
)

func TestDigikalaImportService_ImportCategories(t *testing.T) {
	// Mock Digikala API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := services.DigikalaCategoryResponse{
			Status: "success",
			Data: []services.DigikalaCategory{
				{
					ID:    123,
					Title: "Test Category",
					Slug:  "test-category",
					Image: "http://example.com/image.jpg",
				},
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	t.Run("Success", func(t *testing.T) {
		mockCategoryRepo := &mocks.MockCategoryRepository{
			GetBySlugFunc: func(ctx context.Context, slug string) (*models.Category, error) {
				return nil, nil // Not found, so create
			},
			CreateFunc: func(ctx context.Context, cat *models.Category) error {
				if cat.Name != "Test Category" {
					t.Errorf("expected name 'Test Category', got '%s'", cat.Name)
				}
				if cat.Slug != "test-category" {
					t.Errorf("expected slug 'test-category', got '%s'", cat.Slug)
				}
				return nil
			},
		}

		service := services.NewDigikalaImportService(nil, mockCategoryRepo)
		service.SetBaseURL(ts.URL)
		service.SetHTTPClient(ts.Client())

		imported, failed, err := service.ImportCategories(context.Background(), []string{"123"}, nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if imported != 1 {
			t.Errorf("expected 1 imported, got %d", imported)
		}
		if failed != 0 {
			t.Errorf("expected 0 failed, got %d", failed)
		}
	})
}
