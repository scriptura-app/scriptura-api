package handler

import (
	"scriptura/scriptura-api/repository"
)

type AppHandlers struct {
	Book    BookHandler
	Chapter ChapterHandler
	Bible   BibleHandler
	Verse   VerseHandler
}

func NewAppHandlers(repo repository.AppRepository) AppHandlers {
	h := AppHandlers{
		Book:    NewBookHandler(repo.Book),
		Chapter: NewChapterHandler(repo.Chapter),
		Verse:   NewVerseHandler(repo.Verse),
		Bible:   NewBibleHandler(repo.Bible),
	}

	return h
}
