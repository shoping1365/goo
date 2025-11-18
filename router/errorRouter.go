package router

import (
	"your_project/middleware"

	"github.com/labstack/echo/v4"
)

func SetupErrorRoutes(e *echo.Echo) {
	e.Use(middleware.ErrorHandler)
	// سایر روت‌های مربوط به ارور
}
