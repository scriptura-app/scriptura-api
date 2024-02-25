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
