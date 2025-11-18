package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterQARoutes defines endpoints for questions and QA categories.
// این تابع مسیرهای مربوط به مدیریت پرسش و پاسخ را با محافظت امنیتی تعریف می‌کند
func RegisterQARoutes(public *gin.RouterGroup, db *gorm.DB) {
	qHandler := handlers.NewQuestionHandler(db)
	catHandler := handlers.NewCategoryQAHandler(db)

	// Questions endpoints
	questions := public.Group("/questions")
	{
		questions.GET("", qHandler.ListQuestions)
		questions.POST("" /* middleware.Auth(), */, middleware.RequirePermission("qa_manage"), qHandler.CreateQuestion)
		questions.GET("/:id", qHandler.GetQuestion)
		questions.PUT("/:id" /* middleware.Auth(), */, middleware.RequirePermission("qa_manage"), qHandler.UpdateQuestion)
		questions.DELETE("/:id" /* middleware.Auth(), */, middleware.RequirePermission("qa_manage"), qHandler.DeleteQuestion)
	}

	// Product-specific questions endpoints
	productQuestions := public.Group("/questions/product/:productId")
	{
		productQuestions.GET("", qHandler.GetProductQuestions)
		productQuestions.POST("" /* middleware.Auth(), */, middleware.RequirePermission("qa_manage"), qHandler.CreateProductQuestion)
	}

	// Categories for QA endpoints
	categoriesQA := public.Group("/categories-qa")
	{
		categoriesQA.GET("", catHandler.GetCategories)
		categoriesQA.POST("" /* middleware.Auth(), */, middleware.RequirePermission("qa_manage"), catHandler.CreateCategory)
		categoriesQA.POST("/transfer-questions" /* middleware.Auth(), */, middleware.RequirePermission("qa_manage"), catHandler.TransferQuestions)
		categoriesQA.DELETE("/:key" /* middleware.Auth(), */, middleware.RequirePermission("qa_manage"), catHandler.DeleteCategory)
	}

	// Admin endpoint for all product questions
	public.GET("/questions/admin" /* middleware.Auth(), */, middleware.RequirePermission("qa_view"), qHandler.ListAllProductQuestions)
}
