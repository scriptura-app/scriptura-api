package router

import (
	"net/http"

	"scriptura/scriptura-api/graphql"
	"scriptura/scriptura-api/handler"
	"scriptura/scriptura-api/middleware"
	"scriptura/scriptura-api/repository"

	_ "scriptura/scriptura-api/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewAppRouter(repo *repository.AppRepository, handlers *handler.AppHandlers) *http.ServeMux {
	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()

	mainMux.Handle("/api/v1/", middleware.JsonMiddleware(http.StripPrefix("/api/v1", apiMux)))

	apiMux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	apiMux.HandleFunc("/book/{id}", handlers.Book.GetById)
	apiMux.HandleFunc("/chapter/{id}", handlers.Chapter.GetById)
	apiMux.HandleFunc("/verse/{id}", handlers.Verse.GetById)
	apiMux.HandleFunc("/bible", handlers.Bible.GetByRef)

	gqlServer, gqlPlayground := graphql.NewServer(repo)

	mainMux.Handle("/graphql", gqlServer)
	mainMux.Handle("/playground", gqlPlayground)

	return mainMux
}
