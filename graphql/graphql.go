package graphql

import (
	"net/http"
	gql "scriptura/scriptura-api/graphql/.generated"
	resolver "scriptura/scriptura-api/graphql/resolver"
	"scriptura/scriptura-api/repository"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// temp
func NewServer(appRepo *repository.AppRepository) (*handler.Server, http.HandlerFunc) {
	resolver := &resolver.Resolver{
		AppRepository: appRepo,
	}
	server := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: resolver}))
	playground := playground.Handler("GraphQL playground", "/graphql")
	return server, playground
}
