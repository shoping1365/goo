package services_test

import (
	"context"
	"errors"
	"testing"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/tests/mocks"
)

func TestMenuService_CreateMenu(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := &mocks.MockMenuRepository{}
		service := services.NewMenuService(mockRepo)

		menu := &models.Menu{Name: "Main Menu", Slug: "main-menu"}

		mockRepo.CreateMenuFunc = func(ctx context.Context, m *models.Menu) error {
			if m != menu {
				return errors.New("unexpected menu")
			}
			return nil
		}

		err := service.CreateMenu(context.Background(), menu)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("ValidationError_Name", func(t *testing.T) {
		mockRepo := &mocks.MockMenuRepository{}
		service := services.NewMenuService(mockRepo)

		menu := &models.Menu{Slug: "main-menu"}

		err := service.CreateMenu(context.Background(), menu)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if err.Error() != "نام منو الزامی است" {
			t.Errorf("expected error 'نام منو الزامی است', got '%v'", err.Error())
		}
	})

	t.Run("ValidationError_Slug", func(t *testing.T) {
		mockRepo := &mocks.MockMenuRepository{}
		service := services.NewMenuService(mockRepo)

		menu := &models.Menu{Name: "Main Menu"}

		err := service.CreateMenu(context.Background(), menu)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if err.Error() != "شناسه منو الزامی است" {
			t.Errorf("expected error 'شناسه منو الزامی است', got '%v'", err.Error())
		}
	})
}

func TestMenuService_GetMenuByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := &mocks.MockMenuRepository{}
		service := services.NewMenuService(mockRepo)

		expectedMenu := &models.Menu{ID: 1, Name: "Main Menu"}

		mockRepo.GetMenuByIDFunc = func(ctx context.Context, id uint) (*models.Menu, error) {
			if id != 1 {
				return nil, errors.New("unexpected id")
			}
			return expectedMenu, nil
		}

		menu, err := service.GetMenuByID(context.Background(), 1)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if menu != expectedMenu {
			t.Errorf("expected menu %v, got %v", expectedMenu, menu)
		}
	})
}

func TestMenuService_CreateMenuItem(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := &mocks.MockMenuRepository{}
		service := services.NewMenuService(mockRepo)

		item := &models.MenuItem{Title: "Home", MenuID: 1}

		mockRepo.CreateMenuItemFunc = func(ctx context.Context, i *models.MenuItem) error {
			if i != item {
				return errors.New("unexpected item")
			}
			return nil
		}

		err := service.CreateMenuItem(context.Background(), item)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("ValidationError_Title", func(t *testing.T) {
		mockRepo := &mocks.MockMenuRepository{}
		service := services.NewMenuService(mockRepo)

		item := &models.MenuItem{MenuID: 1}

		err := service.CreateMenuItem(context.Background(), item)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if err.Error() != "عنوان آیتم الزامی است" {
			t.Errorf("expected error 'عنوان آیتم الزامی است', got '%v'", err.Error())
		}
	})

	t.Run("ValidationError_MenuID", func(t *testing.T) {
		mockRepo := &mocks.MockMenuRepository{}
		service := services.NewMenuService(mockRepo)

		item := &models.MenuItem{Title: "Home"}

		err := service.CreateMenuItem(context.Background(), item)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if err.Error() != "شناسه منو الزامی است" {
			t.Errorf("expected error 'شناسه منو الزامی است', got '%v'", err.Error())
		}
	})
}
