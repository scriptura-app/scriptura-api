package repository

import (
	"scriptura/scriptura-api/db"
	m "scriptura/scriptura-api/models"
)

type BibleTextInput struct {
	Bible      string
	Book       string
	Chapter    string
	StartVerse string
	EndVerse   string
	Offset     int
	Limit      int
}

func GetBibleText(i BibleTextInput) ([]m.Verse, int, error) {

	var verses []m.Verse
	var totalItems int64

	err := db.Select(&verses, "verses_get")

	if err != nil {
		return nil, 0, err
	}

	// if i.Chapter != "" {
	// 	query = query.Where("verse.chapter_num = ?", i.Chapter)
	// }

	// if i.StartVerse != "" {
	// 	query = query.Where("verse.verse_num >= ?", i.StartVerse)

	// 	if i.EndVerse == "" {
	// 		query = query.Where("verse.verse_num <= ?", i.StartVerse)
	// 	}
	// }

	// if i.EndVerse != "" {
	// 	query = query.Where("verse.verse_num <= ?", i.EndVerse)
	// }

	// query.Count(&totalItems)

	// query.Offset(i.Offset).
	// 	Limit(i.Limit).
	// 	Scan(&response)

	// if query.Error != nil {
	// 	return response, 0, query.Error
	// }

	return verses, int(totalItems), nil
}
