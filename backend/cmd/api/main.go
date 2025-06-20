package main

import (
	"backend/config"
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
	addressRepository := repositories.NewAddressRepository(db)
	storeRepository := repositories.NewStoreRepository(db)
	establishmentRepository := repositories.NewEstablishmentRepository(db)

	userService := services.NewUserService(userRepository)
	establishmentService := services.NewEstablishmentService(addressRepository, establishmentRepository, storeRepository)
	storeService := services.NewStoreService(addressRepository, storeRepository)

	userHandler := handlers.NewUserHandler(userService)
	establishmentHandler := handlers.NewEstablishmentHandler(establishmentService)
	storeHandler := handlers.NewStoreHandler(storeService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.Routes(e, userHandler, establishmentHandler, storeHandler)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is ON!")
	})

	port := os.Getenv("APP_PORT")
	app_address := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(app_address))
}
