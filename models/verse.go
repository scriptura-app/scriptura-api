package models

import (
	"encoding/json"
	"scriptura/scriptura-api/utils"
)

type Verse struct {
	Id            int    `json:"id"`
	Type          string `json:"type"`
	Self          string `json:"self"`
	BookName      string `json:"bookName,omitempty"`
	ChapterNum    uint   `json:"chapterNum,omitempty"`
	VerseNum      uint   `json:"verseNum,omitempty"`
	YearWritten   int    `json:"yearWritten,omitempty"`
	Text          string `json:"text,omitempty"`
	BookId        uint   `json:"bookId,omitempty"`
	TheographicId string `json:"-"`
	VerseCode     uint   `json:"verseCode,omitempty"`
}

func (v Verse) MarshalJSON() ([]byte, error) {
	type Alias Verse
	v.Type = "verse"
	v.Self = utils.GetURI("verse", v.Id)
	return json.Marshal(&struct{ Alias }{Alias: (Alias)(v)})
}
