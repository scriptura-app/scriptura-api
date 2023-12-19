package router

import (
	"scriptura/scriptura-api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app fiber.Router) {
	api := app.Group("/api", logger.New())

	// Auth
	verse := api.Group("/verse")
	verse.Get("/", handler.GetVerse)
}
