package router

import (
	"scriptura/scriptura-api/gql"
	"scriptura/scriptura-api/handler"
	m "scriptura/scriptura-api/middleware"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app fiber.Router) {
	api := app.Group("/api/v1", logger.New())

	app.Use(swagger.New(swagger.Config{
		FilePath: "docs/swagger.json",
		Path:     "swagger",
	}))

	api.Use(m.JsonMiddleware)

	bible := api.Group("/bible", m.PaginationMiddleware)

	book := bible.Group("/book")
	book.Get("/:book", handler.GetBook)

	chapter := book.Group("/chapter")
	chapter.Get("/:chapter", handler.GetChapter)

	verse := chapter.Group("/verse")
	verse.Get("/:start", handler.GetSingleVerse)
	verse.Get("/:start-:end", handler.GetVerseRange)

	app.Post("/graphql", gql.Handler)
}
