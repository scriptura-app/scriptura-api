package repository

import (
	"fmt"
	"scriptura/scriptura-api/models"

	"gorm.io/gorm"
)

type VerseRepository interface {
	GetById(input int) (models.Verse, error)
}

type verseRepository struct {
	db      *gorm.DB
	appRepo *AppRepository
}

func NewVerseRepository(db *gorm.DB, appRepo *AppRepository) VerseRepository {
	return &verseRepository{db: db, appRepo: appRepo}
}

func (r *verseRepository) GetById(id int) (models.Verse, error) {
	db := r.db
	var verse models.Verse

	db.Table("verse v").
		Select("v.*, b.text").
		Joins(fmt.Sprintf("left join bible_%s b on b.verse_id = v.id", "en_kj")).
		Where("verse.id = ?", id).
		Scan(&verse)

	return verse, nil
}

// func GetVerseRangeByRef(bible string, book string, chapter string, from string, to string, offset int, limit int) ([]m.Verse, int, error) {
// 	db := db.DB
// 	var versesResponse []m.Verse
// 	var totalItems int64

// 	query := db.Table("verse").
// 		Select("verse.*, bible.text, book.name as book_name").
// 		Joins("left join book on book.id = book_id").
// 		Joins(fmt.Sprintf("left join bible_%s bible on bible.verse_id = verse.id", bible)).
// 		Where("book.id::varchar ilike ? OR book.code ilike ? OR book.short_name ilike ?", book, book, book).
// 		Where("verse.chapter_num = ?", chapter).
// 		Where("verse.verse_num >= ? AND verse.verse_num <= ?", from, to)

// 	query.Count(&totalItems)

// 	query.Offset(offset).
// 		Limit(limit).Scan(&versesResponse)

// 	return versesResponse, int(totalItems), nil
// }
