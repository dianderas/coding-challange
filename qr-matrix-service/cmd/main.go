package main

import (
	"fmt"
	"log"
	"os"
	"qr-matrix-service/internal/handlers"
	"qr-matrix-service/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	initializeEnvVars()

	app := fiber.New()

	app.Use(logger.New())

	// ruta publica para el login
	app.Post("/auth/login", handlers.Login)

	// rutas protegidas
	api := app.Group("/api", middleware.AuthMiddleware())
	api.Post("/process-matrix", handlers.ProcessMatrix)

	log.Fatal(app.Listen(":8080"))
}

func initializeEnvVars() {
	secret := os.Getenv("JWT_SECRET")

	// Verificar si está definida
	if secret == "" {
		fmt.Println("❌ JWT_SECRET no está configurado")
	} else {
		fmt.Println("✅ JWT_SECRET está configurado:", secret)
	}
}
