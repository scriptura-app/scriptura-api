package models

import (
	"encoding/json"
)

type Verse struct {
	Id            string `json:"id"`
	Type          string `json:"type"`
	BookName      string `json:"bookName,omitempty"`
	ChapterNum    uint   `json:"chapterNum,omitempty"`
	VerseNum      uint   `json:"verseNum,omitempty"`
	YearWritten   int    `json:"yearWritten,omitempty"`
	Text          string `json:"text,omitempty"`
	BookId        uint   `json:"bookId,omitempty"`
	TheographicId string `json:"-"`
	VerseCode     uint   `json:"verseCode,omitempty"`
	People        string `json:"people,omitempty"`
	Places        string `json:"places,omitempty"`
	Events        string `json:"events,omitempty"`
}

func (v Verse) MarshalJSON() ([]byte, error) {
	type Alias Verse
	v.Type = "verse"
	return json.Marshal(&struct{ Alias }{Alias: (Alias)(v)})
}
