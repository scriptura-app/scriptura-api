package models

type Verse struct {
	ID          uint   `json:"id"`
	BookId      uint   `json:"bookId"`
	ChapterNum  uint   `json:"chapterNum"`
	VerseNum    uint   `json:"verseNum"`
	VerseCode   uint   `json:"verseCode"`
	YearWritten int    `json:"year"`
	Text        string `json:"text"`
}

type Tabler interface {
	TableName() string
}

func (Verse) TableName() string {
	return "verse"
}
