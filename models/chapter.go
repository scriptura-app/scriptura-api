package models

import (
	"encoding/json"
)

type Chapter struct {
	Id         string  `json:"id,omitempty"`
	Type       string  `json:"type"`
	BookName   string  `json:"bookName,omitempty"`
	BookId     uint    `json:"bookId,omitempty"`
	ChapterNum uint    `json:"chapterNum,omitempty"`
	VerseCount int     `json:"verseCount,omitempty"`
	YearRange  []int   `json:"yearWrittenRange,omitempty"`
	Verses     []Verse `json:"verses,omitempty"`
}

func (v Chapter) MarshalJSON() ([]byte, error) {
	type Alias Chapter
	v.Type = "chapter"
	return json.Marshal(&struct{ Alias }{Alias: (Alias)(v)})
}
