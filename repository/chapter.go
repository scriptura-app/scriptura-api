//go:build exclude

package repository

import (
	"scriptura/scriptura-api/db"
	m "scriptura/scriptura-api/models"
)

func GetChapter(book string, chapter string) (m.Chapter, error) {
	db := db.DB
	var chapterResponse m.Chapter

	query := db.Table("verse").
		Select("count(*) as verse_count, book.name as book_name, book_id, chapter_num, array_agg(year_written) as year_range").
		Joins("left join book on book.id = book_id").
		Where("book.id::varchar ilike ? OR book.code ilike ? OR book.short_name ilike ?", book, book, book).
		Where("verse.chapter_num = ?", chapter).
		Group("verse, book.name, book_id, chapter_num").
		Scan(&chapterResponse)

	if query.Error != nil {
		return chapterResponse, query.Error
	}

	verseQuery := db.Table("verse").
		Select("verse.id, verse.verse_num").
		Joins("left join book on book.id = book_id").
		Where("book.id::varchar ilike ? OR book.code ilike ? OR book.short_name ilike ?", book, book, book).
		Where("verse.chapter_num = ?", chapter).
		Scan(&chapterResponse.Verses)

	if verseQuery.Error != nil {
		return chapterResponse, verseQuery.Error
	}

	return chapterResponse, nil
}
