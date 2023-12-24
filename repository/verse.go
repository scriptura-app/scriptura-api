package repository

import (
	"scriptura/scriptura-api/db"
	m "scriptura/scriptura-api/models"
)

func GetVerseByRef(bible string, book string, chapter string, verse string) (m.Verse, error) {
	db := db.DB
	var verseResponse m.Verse

	db.Table("verse").
		Select("verse.*, bible.text, book.name as book_name").
		Joins("left join book on book.id = book_id").
		Joins("left join bible_en_kjv bible on bible.verse_id = verse.id").
		Where("book.id::varchar ilike $1 OR book.code ilike $1 OR book.short_name ilike $1", book).
		Where("verse.chapter_num = ?", chapter).
		Where("verse.verse_num = ?", verse).
		Scan(&verseResponse)

	return verseResponse, nil
}

// func getVerseRangeByRef(bible string, book string, chapter string, verse string) (m.Verse, error) {
// 	db := db.DB
// 	var verseResponse m.Verse

// }
