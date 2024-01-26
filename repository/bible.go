package repository

import (
	"fmt"
	"scriptura/scriptura-api/db"
	"scriptura/scriptura-api/models"
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

func GetBibleText(i BibleTextInput) ([]models.Verse, int, error) {
	db := db.DB
	var response []models.Verse
	var totalItems int64

	query := db.Table("verses v").
		Select("v.*, bv.text, b.name as book_name").
		Joins("left join books b on b.id = v.book_id").
		Joins(fmt.Sprintf("left join bible_%s bv on bv.verse_id = v.id", i.Bible)).
		Where("b.id::varchar ilike ? OR b.slug ilike ? OR b.short_name ilike ?", i.Book, i.Book, i.Book)

	if i.Chapter != "" {
		query = query.Where("v.chapter_num = ?", i.Chapter)
	}

	if i.StartVerse != "" {
		query = query.Where("v.verse_num >= ?", i.StartVerse)

		if i.EndVerse == "" {
			query = query.Where("v.verse_num <= ?", i.StartVerse)
		}
	}

	if i.EndVerse != "" {
		query = query.Where("v.verse_num <= ?", i.EndVerse)
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
