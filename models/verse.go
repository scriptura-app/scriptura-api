package models

type Verse struct {
	Id            string `json:"id"`
	Id            string `json:"-"`
	BookName      string `json:"bookName,omitempty"`
	ChapterNum    uint   `json:"chapterNum,omitempty"`
	VerseNum      uint   `json:"verseNum,omitempty"`
	YearWritten   int    `json:"yearWritten,omitempty"`
	Text          string `json:"text,omitempty"`
	BookId        uint   `json:"bookId,omitempty"`
	TheographicId string `json:"-"`
	VerseCode     uint   `json:"verseCode,omitempty"`
}
