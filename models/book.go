package models

import (
	"encoding/json"
)

type Book struct {
	Id            string `json:"id"`
	TheographicId string `json:"-"`
	Type          string `json:"type"`
	Code          string `json:"code"`
	DivisionId    string `json:"divisionId,omitempty"`
	Testament     string `json:"testament,omitempty"`
	Name          string `json:"name,omitempty"`
	ShortName     string `json:"shortName,omitempty"`
	ChapterCount  uint   `json:"chapterCount,omitempty"`
	VerseCount    uint   `json:"verseCount,omitempty"`
	PeopleCount   uint   `json:"peopleCount",omitempty`
	PlaceCount    uint   `json:"placeCount,omitempty"`
}

func (b Book) MarshalJSON() ([]byte, error) {
	type Alias Book
	b.Type = "book"
	return json.Marshal(&struct{ Alias }{Alias: (Alias)(b)})
}
