//go:build exclude

package repository

import (
	"fmt"
	"scriptura/scriptura-api/db"
	m "scriptura/scriptura-api/models"
)

func GetVerseByRef(bible string, book string, chapter string, verse string) (m.Verse, error) {
	db := db.DB
	var verseResponse m.Verse

	db.Table("verse").
		Select("verse.*, bible.text, book.name as book_name").
		Joins("left join book on book.id = book_id").
		Joins(fmt.Sprintf("left join bible_%s bible on bible.verse_id = verse.id", bible)).
		Where("book.id::varchar ilike ? OR book.code ilike ? OR book.short_name ilike ?", book, book, book).
		Where("verse.chapter_num = ?", chapter).
		Where("verse.verse_num = ?", verse).
		Scan(&verseResponse)

	return verseResponse, nil
}

func GetVerseRangeByRef(bible string, book string, chapter string, from string, to string, offset int, limit int) ([]m.Verse, int, error) {
	db := db.DB
	var versesResponse []m.Verse
	var totalItems int64

	query := db.Table("verse").
		Select("verse.*, bible.text, book.name as book_name").
		Joins("left join book on book.id = book_id").
		Joins(fmt.Sprintf("left join bible_%s bible on bible.verse_id = verse.id", bible)).
		Where("book.id::varchar ilike ? OR book.code ilike ? OR book.short_name ilike ?", book, book, book).
		Where("verse.chapter_num = ?", chapter).
		Where("verse.verse_num >= ? AND verse.verse_num <= ?", from, to)

	query.Count(&totalItems)

	query.Offset(offset).
		Limit(limit).Scan(&versesResponse)

	return versesResponse, int(totalItems), nil
}
