package models

import (
	"encoding/json"
	"scriptura/scriptura-api/utils"
)

type Chapter struct {
	Id         int      `json:"id,omitempty"`
	Type       string   `json:"type"`
	Self       string   `json:"self"`
	ChapterNum uint     `json:"chapterNum,omitempty"`
	VerseCount int      `json:"verseCount,omitempty"`
	YearRange  []int    `json:"yearWrittenRange,omitempty"`
	Book       *Book    `json:"book,omitempty"`
	Verses     *[]Verse `json:"verses,omitempty"`

	VersesJson string `json:"-"`
	BookJson   string `json:"-"`
}

func (c Chapter) MarshalJSON() ([]byte, error) {
	type Alias Chapter
	c.Type = "chapter"
	c.Self = utils.GetURI("chapter", c.Id)
	return json.Marshal(&struct{ Alias }{Alias: (Alias)(c)})
}
