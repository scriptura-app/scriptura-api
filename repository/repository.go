package repository

import (
	"gorm.io/gorm"
)

type AppRepository struct {
	Book    BookRepository
	Verse   VerseRepository
	Chapter ChapterRepository
	Bible   BibleRepository
}

func NewAppRepository(db *gorm.DB) AppRepository {
	var r AppRepository
	r.Book = NewBookRepository(db, &r)
	r.Verse = NewVerseRepository(db, &r)
	r.Chapter = NewChapterRepository(db, &r)
	r.Bible = NewBibleRepository(db, &r)
	return r
}
