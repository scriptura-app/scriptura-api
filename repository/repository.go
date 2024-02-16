package repository

import (
	"gorm.io/gorm"
)

type AppRepository struct {
	BookRepository BookRepository
}

func NewAppRepository(db *gorm.DB) AppRepository {
	var r AppRepository
	r.BookRepository = NewBookRepository(db, &r)

	return r
}
