package router

import (
	"scriptura/scriptura-api/handler"
	m "scriptura/scriptura-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app fiber.Router) {
	api := app.Group("/api", logger.New())

	api.Use(m.JsonMiddleware)
	api.Use(m.PaginationMiddleware)

	verse := api.Group("/bible/:book?/:chapter?/:verse?")
	verse.Get("/", handler.GetVerse)
}
