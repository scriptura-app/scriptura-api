package router

import (
	"scriptura/scriptura-api/graphql"
	"scriptura/scriptura-api/handler"
	"scriptura/scriptura-api/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	bookRepository := repository.NewBookRepository()
	bookHandler := handler.NewBookHandler(bookRepository)

	r.Get("/book/{book}", bookHandler.GetBook)

	gqlServer, gqlPlayground := graphql.NewServer(bookRepository)

	r.Handle("/graphql", gqlServer)
	r.Handle("/playground", gqlPlayground)

	return r
}

//OLD------------V
/*
func SetupRoutes(app fiber.Router) {
	api := app.Group("/api/v1", logger.New())

	app.Use(swagger.New(swagger.Config{
		FilePath: "docs/swagger.json",
		Path:     "swagger",
	}))

	api.Use(m.JsonMiddleware)

	api.Get("/book/:book", handler.GetBook)
	api.Get("/chapter/:chapter", handler.GetChapter)

	bible := api.Group("/bible", m.PaginationMiddleware)

	bible.Get("/book/:book", handler.GetBible)

	bible.Get("/book/:book/chapter/:chapter", handler.GetBible)
	bible.Get("/book/:book/chapter/:chapter/verse/:start-:end", handler.GetBible)
	bible.Get("/book/:book/chapter/:chapter/verse/:start", handler.GetBible)

	app.Post("/graphql", gql.Handler)
}
*/
