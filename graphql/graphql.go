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
func NewServer(bookRepository repository.BookRepository) (*handler.Server, http.HandlerFunc) {
	resolver := &resolver.Resolver{
		BookRepository: bookRepository,
	}
	server := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: resolver}))
	playground := playground.Handler("GraphQL playground", "/graphql")
	return server, playground
}
