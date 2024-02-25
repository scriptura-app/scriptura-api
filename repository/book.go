package repository

import (
	"encoding/json"
	"scriptura/scriptura-api/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetById(input string) (models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetById(input string) (models.Book, error) {
	var err error
	db := r.db
	var book models.Book

	chapSubq := r.db.Table("chapters c").
		Select("json_agg(to_json((SELECT d FROM (SELECT c.id, c.chapter_num as chapterNum) d)))").
		Where("b.id = c.book_id")

	chapCountSubq := db.Table("chapters c").Select("count(*)").Where("b.id = c.book_id")
	verseCountSubq := db.Table("verses v").Select("count(*)").Where("b.id = v.book_id")

	db.Table("books b").
		Select("b.*, bd.name as division, t.name as testament, (?) as chapters_json, (?) as chapter_count, (?) as verse_count", chapSubq, chapCountSubq, verseCountSubq).
		Joins("join book_divisions bd on bd.id = b.division_id").
		Joins("join testaments t on t.id = bd.testament_id").
		Where("b.id::varchar ilike ? OR b.slug ilike ? OR b.short_name ilike ?", input, input, input).
		First(&book)

	err = json.Unmarshal([]byte(book.ChaptersJson), &book.Chapters)

	return book, err
}
