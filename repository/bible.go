package repository

import (
	"fmt"
	"scriptura/scriptura-api/models"

	"gorm.io/gorm"
)

type BibleRepository interface {
	GetBibleText(i BibleTextInput) ([]models.Verse, int, error)
}

type bibleRepository struct {
	db *gorm.DB
}

func NewBibleRepository(db *gorm.DB) BibleRepository {
	return &bibleRepository{db: db}
}

type BibleTextInput struct {
	Version    string
	Book       string
	Chapter    int
	StartVerse int
	EndVerse   int
	Cursor     int
	Offset     int
	Limit      int
}

func (r *bibleRepository) GetBibleText(i BibleTextInput) ([]models.Verse, int, error) {
	db := r.db
	var response []models.Verse
	var totalItems int64

	query := db.Table("verses v").
		Select("v.*, bv.text").
		Joins("left join books b on b.id = v.book_id").
		Joins(fmt.Sprintf("left join bible_%s bv on bv.verse_id = v.id", i.Version))

	if i.Cursor != 0 {
		query = query.Where("v.id >= ?", i.Cursor)
	} else {
		if i.Book != "" {
			query = query.Where("b.id::varchar ilike ? OR b.slug ilike ? OR b.short_name ilike ?", i.Book, i.Book, i.Book)
		}
		if i.Chapter != 0 {
			query = query.Where("v.chapter_num = ?", i.Chapter)
		}
		if i.StartVerse != 0 {
			query = query.Where("v.verse_num >= ?", i.StartVerse)

			if i.EndVerse == 0 {
				query = query.Where("v.verse_num <= ?", i.StartVerse)
			}
		}
		if i.EndVerse != 0 {
			query = query.Where("v.verse_num <= ?", i.EndVerse)
		}
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
