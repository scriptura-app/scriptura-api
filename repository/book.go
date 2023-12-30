package repository

import (
	"scriptura/scriptura-api/db"
	m "scriptura/scriptura-api/models"
)

func GetBook(book string) (m.Book, error) {
	db := db.DB
	var bookResponse m.Book

	db.Table("book").
		Select("book.*").
		Where("book.id::varchar ilike ? OR book.code ilike ? OR book.short_name ilike ?", book, book, book).
		Scan(&bookResponse)

	return bookResponse, nil
}
