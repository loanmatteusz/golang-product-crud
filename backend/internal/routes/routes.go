package routes

import (
	"backend/internal/handlers"

	"github.com/labstack/echo/v4"
)

func Routes(
	e *echo.Echo,
	userHandler *handlers.UserHandler,
	establishmentHandler *handlers.EstablishmentHandler,
	storeHandler *handlers.StoreHandler,
) {
	auth := e.Group("/auth")
	auth.POST("/register", userHandler.Register)
	auth.POST("/login", userHandler.Login)

	establishment := e.Group("/establishments")
	establishment.POST("", establishmentHandler.Create)
	establishment.GET("", establishmentHandler.GetAll)
	establishment.GET("/:id", establishmentHandler.GetByID)
	establishment.PUT("/:id", establishmentHandler.Update)
	establishment.DELETE("/:id", establishmentHandler.Delete)

	store := e.Group("/stores")
	store.POST("", storeHandler.Create)
	store.GET("", storeHandler.GetAll)
	store.GET("/:id", storeHandler.GetByID)
	store.GET("/establishment/:id", storeHandler.GetByEstablishmentID)
	store.PUT("/:id", storeHandler.Update)
	store.DELETE("/:id", storeHandler.Delete)
}
