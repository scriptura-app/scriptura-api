package gql

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

type Input struct {
	Query         string                 `query:"query"`
	OperationName string                 `query:"operationName"`
	Variables     map[string]interface{} `query:"variables"`
}

var schema graphql.Schema

func CreateSchema() error {
	var err error
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"foo": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "bar", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return err
}

func Handler(ctx *fiber.Ctx) error {
	var input Input
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			SendString("Cannot parse body: " + err.Error())
	}

	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  input.Query,
		OperationName:  input.OperationName,
		VariableValues: input.Variables,
	})

	ctx.Set("Content-Type", "application/graphql-response+json")
	return ctx.JSON(result)
}
