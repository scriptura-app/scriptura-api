package models

import (
	"encoding/json"
	"scriptura/scriptura-api/utils"
)

type Book struct {
	Id            int    `json:"id"`
	TheographicId string `json:"-"`
	Type          string `json:"type"`
	Self          string `json:"self"`
	Name          string `json:"name,omitempty"`
	Slug          string `json:"slug"`
	Testament     string `json:"testament,omitempty"`
	Division      string `json:"division,omitempty"`
	ShortName     string `json:"shortName,omitempty"`

	ChapterCount int       `json:"chapterCount,omitempty"`
	VerseCount   int       `json:"verseCount,omitempty"`
	Chapters     []Chapter `json:"chapters"`

	ChaptersJson string `json:"-"`
}

func (b Book) MarshalJSON() ([]byte, error) {
	type Alias Book
	b.Type = "book"
	b.Self = utils.GetURI("book", b.Id)
	return json.Marshal(&struct{ Alias }{Alias: (Alias)(b)})
}
