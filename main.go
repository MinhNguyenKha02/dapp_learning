package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"dapp_learning/internal/handlers"
	"dapp_learning/internal/middleware"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 25 * 1024 * 1024, // 25MB
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// File upload routes with validation middleware
	api.Post("/upload", middleware.ValidateFileUpload(), handlers.UploadFiles)
}
