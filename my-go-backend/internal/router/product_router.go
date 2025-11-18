package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterProductRoutes defines all product-related endpoints including child resources
// such as images, variants, specifications and videos.
func RegisterProductRoutes(public *gin.RouterGroup, db *gorm.DB) {
	// Base product CRUD
	productHandler := handlers.NewProductHandler(db)

	products := public.Group("/products")
	{
		products.GET("", middleware.Auth(), productHandler.ListProducts)
		products.POST("", middleware.Auth(), middleware.RequirePermission("products_create"), productHandler.CreateProduct)
		products.GET("/public", productHandler.ListPublicProducts) // API عمومی برای صفحات عمومی
		products.GET("/:id", productHandler.GetProduct)
		products.PUT("/:id", middleware.Auth(), middleware.RequirePermission("products_edit"), productHandler.UpdateProduct)
		products.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("products_delete"), productHandler.DeleteProduct)
		// Helpers
		products.GET("/check-slug", productHandler.CheckSlug)
		products.GET("/generate-slug", productHandler.GenerateSlug)
		// انتقال گروهی محصولات
		products.POST("/bulk-transfer", middleware.Auth(), middleware.RequirePermission("products_edit"), productHandler.BulkTransferProducts)
		// لیست تصاویر محصولات
		products.GET("/images", productHandler.ListProductImages)
	}

	// Stand-alone child resources to avoid wildcard conflicts
	productImageHandler := handlers.NewProductImageHandler(db)
	public.POST("/product-images/:productId", middleware.Auth(), middleware.RequirePermission("products_edit"), productImageHandler.AddImage)
	public.DELETE("/product-images/:imageId", middleware.Auth(), middleware.RequirePermission("products_edit"), productImageHandler.DeleteImage)

	productVariantHandler := &handlers.ProductVariantHandler{DB: db}
	public.GET("/product-variants/:productId", productVariantHandler.GetVariants)
	public.POST("/product-variants/:productId", middleware.Auth(), middleware.RequirePermission("products_edit"), productVariantHandler.AddVariant)
	public.PUT("/product-variants/:variantId", middleware.Auth(), middleware.RequirePermission("products_edit"), productVariantHandler.UpdateVariant)
	public.DELETE("/product-variants/:variantId", middleware.Auth(), middleware.RequirePermission("products_edit"), productVariantHandler.DeleteVariant)

	// Complementary products
	complements := &handlers.ProductComplementHandler{DB: db}
	public.GET("/product-complements/:productId", complements.GetComplements)
	public.POST("/product-complements/:productId", middleware.Auth(), middleware.RequirePermission("products_edit"), complements.AddComplement)
	public.DELETE("/product-complements/:productId/:complementId", middleware.Auth(), middleware.RequirePermission("products_edit"), complements.RemoveComplement)

	productSpecHandler := &handlers.ProductSpecificationHandler{DB: db}
	public.POST("/product-specifications/:productId", middleware.Auth(), middleware.RequirePermission("products_edit"), productSpecHandler.AddSpecification)
	public.PUT("/product-specifications/:specId", middleware.Auth(), middleware.RequirePermission("products_edit"), productSpecHandler.UpdateSpecification)
	public.DELETE("/product-specifications/:specId", middleware.Auth(), middleware.RequirePermission("products_edit"), productSpecHandler.DeleteSpecification)

	productVideoHandler := &handlers.ProductVideoHandler{DB: db}
	public.GET("/product-videos/:productId", productVideoHandler.GetVideos)
	public.POST("/product-videos/:productId", middleware.Auth(), middleware.RequirePermission("products_edit"), productVideoHandler.AddVideo)
	public.PUT("/product-videos/:videoId", middleware.Auth(), middleware.RequirePermission("products_edit"), productVideoHandler.UpdateVideo)
	public.DELETE("/product-videos/:videoId", middleware.Auth(), middleware.RequirePermission("products_edit"), productVideoHandler.DeleteVideo)
}
