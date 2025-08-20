package main

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/routes"
	"backend/internal/services"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, relying on system env vars")
	}

	config.ConnectDatabase()
	db := config.DB

	userRepository := repositories.NewUserRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)
	productRepository := repositories.NewProductRepository(db)

	secret := os.Getenv("JWT_SECRET")
	userService := services.NewUserService(userRepository, secret)
	categoryService := services.NewCategoryService(categoryRepository)
	productService := services.NewProductService(productRepository, categoryRepository)

	userHandler := handlers.NewUserHandler(userService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	productHandler := handlers.NewProductHandler(productService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.Routes(e, secret, userHandler, categoryHandler, productHandler)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is ON!")
	})

	port := os.Getenv("APP_PORT")
	app_address := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(app_address))
}
