package routes

import (
	"backend/internal/handlers"

	"github.com/labstack/echo/v4"
)

func Routes(
	e *echo.Echo,
	userHandler *handlers.UserHandler,
	categoryHandler *handlers.CategoryHandler,
	productHandler *handlers.ProductHandler,
) {
	v1 := e.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/register", userHandler.Register)
	auth.POST("/login", userHandler.Login)
	auth.POST("/refresh", userHandler.Refresh)

	products := v1.Group("/products")
	products.POST("", productHandler.Create)
	products.GET("", productHandler.GetAll)
	products.GET("/:id", productHandler.GetByID)
	products.PUT("/:id", productHandler.Update)
	products.DELETE("/:id", productHandler.Delete)

	categories := v1.Group("/categories")
	categories.POST("", categoryHandler.Create)
	categories.GET("", categoryHandler.GetAll)
	categories.GET("/:id", categoryHandler.GetByID)
	categories.PUT("/:id", categoryHandler.Update)
	categories.DELETE("/:id", categoryHandler.Delete)
}
