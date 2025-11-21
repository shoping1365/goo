package services_test

import (
	"context"
	"testing"

	"my-go-backend/internal/cqrs/commands"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"
	"my-go-backend/tests/mocks"
)

func TestProductService_CreateProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		cmd := commands.CreateProductCommand{
			Name:       "Test Product",
			Slug:       "test-product",
			SKU:        "TEST-SKU",
			CategoryID: 1,
			Price:      1000.0,
		}

		mockProductRepo := &mocks.MockProductRepository{
			GetBySKUFunc: func(ctx context.Context, sku string) (*models.Product, error) {
				return nil, nil // Not found
			},
			GetBySlugFunc: func(ctx context.Context, slug string) (*models.Product, error) {
				return nil, nil // Not found
			},
			CreateFunc: func(ctx context.Context, p *models.Product) error {
				if p.Name != cmd.Name {
					t.Errorf("expected name %s, got %s", cmd.Name, p.Name)
				}
				return nil
			},
		}

		mockCategoryRepo := &mocks.MockCategoryRepository{
			GetByIDFunc: func(ctx context.Context, id uint) (*models.Category, error) {
				return &models.Category{ID: id, Name: "Test Category"}, nil
			},
		}

		mockUoW := &mocks.MockUnitOfWork{
			ProductRepoFunc: func() repository.ProductRepositoryInterface {
				return mockProductRepo
			},
			CategoryRepoFunc: func() repository.CategoryRepositoryInterface {
				return mockCategoryRepo
			},
			BrandRepoFunc: func() repository.BrandRepositoryInterface {
				return &mocks.MockBrandRepository{}
			},
			CommitFunc: func() error { return nil },
		}

		mockUoWFactory := &mocks.MockUnitOfWorkFactory{
			CreateFunc: func() unitofwork.UnitOfWork {
				return mockUoW
			},
		}

		service := services.NewProductService(mockUoWFactory)

		product, err := service.CreateProduct(context.Background(), cmd)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if product == nil {
			t.Error("expected product, got nil")
		}
	})
}
