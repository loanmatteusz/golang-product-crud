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
	auth := e.Group("/auth")
	auth.POST("/register", userHandler.Register)
	auth.POST("/login", userHandler.Login)

	products := e.Group("/products")
	products.POST("", productHandler.Create)
	products.GET("", productHandler.GetAll)
	products.GET("/:id", productHandler.GetByID)
	products.PUT("/:id", productHandler.Update)
	products.DELETE("/:id", productHandler.Delete)

	categories := e.Group("/categories")
	categories.POST("", categoryHandler.Create)
	categories.GET("", categoryHandler.GetAll)
	categories.GET("/:id", categoryHandler.GetByID)
	categories.PUT("/:id", categoryHandler.Update)
	categories.DELETE("/:id", categoryHandler.Delete)
}
