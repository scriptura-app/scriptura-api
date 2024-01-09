package repository

import (
	"fmt"
	"scriptura/scriptura-api/db"
	m "scriptura/scriptura-api/models"
)

type BibleTextInput struct {
	Bible      string
	Book       string
	Chapter    string
	StartVerse string
	EndVerse   string
	Offset     int
	Limit      int
}

func GetBibleText(i BibleTextInput) ([]m.Verse, int, error) {
	db := db.DB
	var response []m.Verse
	var totalItems int64

	query := db.Table("verse").
		Select("verse.*, bible.text, book.name as book_name").
		Joins("left join book on book.id = book_id").
		Joins(fmt.Sprintf("left join bible_%s bible on bible.verse_id = verse.id", i.Bible)).
		Where("book.id::varchar ilike ? OR book.code ilike ? OR book.short_name ilike ?", i.Book, i.Book, i.Book)

	if i.Chapter != "" {
		query = query.Where("verse.chapter_num = ?", i.Chapter)
	}

	if i.StartVerse != "" {
		query = query.Where("verse.verse_num >= ?", i.StartVerse)

		if i.EndVerse == "" {
			query = query.Where("verse.verse_num <= ?", i.StartVerse)
		}
	}

	if i.EndVerse != "" {
		query = query.Where("verse.verse_num <= ?", i.EndVerse)
	}

	query.Count(&totalItems)

	query.Offset(i.Offset).
		Limit(i.Limit).
		Scan(&response)

	if query.Error != nil {
		return response, 0, query.Error
	}

	return response, int(totalItems), nil
}
