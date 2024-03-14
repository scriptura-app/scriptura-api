package models

import (
	"encoding/json"
	"scriptura/scriptura-api/utils"
)

type Verse struct {
	Id            int    `json:"id"`
	Type          string `json:"type"`
	Self          string `json:"self"`
	BookId        int    `json:"bookId,omitempty"`
	BookName      string `json:"bookName,omitempty"`
	ChapterNum    int    `json:"chapterNum,omitempty"`
	VerseNum      int    `json:"verseNum,omitempty"`
	YearWritten   int    `json:"yearWritten,omitempty"`
	Text          string `json:"text,omitempty"`
	TheographicId string `json:"-"`
	VerseCode     int    `json:"verseCode,omitempty"`
}

func (v Verse) MarshalJSON() ([]byte, error) {
	type Alias Verse
	v.Type = "verse"
	v.Self = utils.GetURI("verse", v.Id)
	return json.Marshal(&struct{ Alias }{Alias: (Alias)(v)})
}
