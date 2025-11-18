package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	searchservice "my-go-backend/internal/search/service"

	"github.com/gin-gonic/gin"
)

// SetupSearchRoutes تنظیم مسیرهای جستجو
func SetupSearchRoutes(r *gin.RouterGroup, searchSvc *searchservice.Service, allowedIndexes []string) {
	searchHandler := handlers.NewSearchHandler(searchSvc, allowedIndexes)

	// گروه مسیرهای جستجو
	searchGroup := r.Group("/search")
	{
		// جستجوی عمومی (نیاز به احراز هویت ندارد)
		searchGroup.GET("/", searchHandler.GlobalSearch)

		// پیشنهادات جستجو (نیاز به احراز هویت ندارد)
		searchGroup.GET("/suggestions", searchHandler.SearchSuggestions)

		// جستجوهای پرطرفدار (نیاز به احراز هویت ندارد)
		searchGroup.GET("/popular", searchHandler.PopularSearchesHandler)

		// مسیرهای جستجوی پیشرفته (نیاز به احراز هویت دارد)
		advancedGroup := searchGroup.Group("/advanced")
		advancedGroup.Use(middleware.AuthMiddleware())
		{
			// جستجوی پیشرفته با فیلترهای بیشتر
			advancedGroup.GET("/", searchHandler.GlobalSearch)

			// جستجو در بخش خاص
			advancedGroup.GET("/products", searchHandler.GlobalSearch)
			advancedGroup.GET("/posts", searchHandler.GlobalSearch)
			advancedGroup.GET("/categories", searchHandler.GlobalSearch)
			advancedGroup.GET("/brands", searchHandler.GlobalSearch)
		}
	}
}
