package graphql

import "scriptura/scriptura-api/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AppRepository *repository.AppRepository
}
