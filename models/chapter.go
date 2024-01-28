package models

import (
	"encoding/json"
	"scriptura/scriptura-api/utils"
)

type Chapter struct {
	Id         int     `json:"id,omitempty"`
	Type       string  `json:"type"`
	Self       string  `json:"self"`
	BookName   string  `json:"bookName,omitempty"`
	BookId     uint    `json:"bookId,omitempty"`
	ChapterNum uint    `json:"chapterNum,omitempty"`
	VerseCount int     `json:"verseCount,omitempty"`
	YearRange  []int   `json:"yearWrittenRange,omitempty"`
	Verses     []Verse `json:"verses,omitempty"`
}

func (c Chapter) MarshalJSON() ([]byte, error) {
	type Alias Chapter
	c.Type = "chapter"
	c.Self = utils.GetURI("chapter", c.Id)
	return json.Marshal(&struct{ Alias }{Alias: (Alias)(c)})
}
