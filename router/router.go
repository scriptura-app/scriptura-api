package router

import (
	"scriptura/scriptura-api/handler"
	m "scriptura/scriptura-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app fiber.Router) {
	api := app.Group("/api/v1", logger.New())

	api.Use(m.JsonMiddleware)

	api.Get("/bible/:book", handler.GetBook)
	api.Get("/bible/:book/:chapter", handler.GetChapter)
	api.Get("/bible/:book/:chapter/:from-:to", m.PaginationMiddleware, handler.GetVerseRange)
	api.Get("/bible/:book/:chapter/:verse", handler.GetSingleVerse)
}
