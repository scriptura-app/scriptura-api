package router

import (
	"net/http"

	"scriptura/scriptura-api/graphql"
	"scriptura/scriptura-api/handler"
	"scriptura/scriptura-api/repository"
)

func NewAppRouter(repo *repository.AppRepository, handlers *handler.AppHandlers) *http.ServeMux {
	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()

	mainMux.Handle("GET /api/v1", apiMux)

	apiMux.HandleFunc("/book/{id}", handlers.Book.GetById)
	apiMux.HandleFunc("/chapter/{id}", handlers.Chapter.GetById)
	apiMux.HandleFunc("/verse/{id}", handlers.Verse.GetById)

	gqlServer, gqlPlayground := graphql.NewServer(repo)

	mainMux.Handle("/graphql", gqlServer)
	mainMux.Handle("/playground", gqlPlayground)

	return mainMux
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
