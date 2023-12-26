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

	api.Get("/book/:book", handler.GetBook)
	api.Get("/book/:book/chapter/:chapter", handler.GetChapter)
	api.Get("/book/:book/chapter/:chapter/verse/:from-:to", m.PaginationMiddleware, handler.GetVerseRange)
	api.Get("/book/:book/chapter/:chapter/verse/:verse", handler.GetSingleVerse)
}
