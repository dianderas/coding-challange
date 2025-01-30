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
	api.Post("/qrmatrix-process", handlers.ProcessMatrix)
	api.Post("/fullmatrix-process", handlers.FullProcessMatrix)

	log.Println("🚀 Iniciando backend-go en el puerto 8080...")

	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		log.Fatal("❌ Error al iniciar el servidor:", err)
	}
}

func initializeEnvVars() {
	secret := os.Getenv("JWT_SECRET")
	extUrl := os.Getenv("STATISTICS_SERVICE_URL")

	// Verificar si está definida
	if secret == "" {
		fmt.Println("❌ JWT_SECRET no está configurado")
		fmt.Println("❌ STATISTICS_SERVICE_URL no está configurado")
	} else {
		fmt.Println("✅ JWT_SECRET está configurado:", secret)
		fmt.Println("✅ STATISTICS_SERVICE_URL está configurado:", extUrl)
	}
}
