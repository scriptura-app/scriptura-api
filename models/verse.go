package models

type Verse struct {
	ID          uint   `json:"id"`
	BookName    string `json:"book"`
	ChapterNum  uint   `json:"chapterNum"`
	VerseNum    uint   `json:"verseNum"`
	YearWritten int    `json:"year"`
	Text        string `json:"text"`
	// BookId        uint   `json:"bookId"`
	// TheographicId string `json:"thegraphicId"`
	// VerseCode     uint   `json:"verseCode"`
}

type Tabler interface {
	TableName() string
}

func (Verse) TableName() string {
	return "verse"
}
