package handler

import (
	"scriptura/scriptura-api/repository"
)

type AppHandlers struct {
	BookHandler BookHandler
}

func NewAppHandlers(repo repository.AppRepository) AppHandlers {
	h := AppHandlers{
		BookHandler: NewBookHandler(repo.BookRepository),
	}

	return h
}
