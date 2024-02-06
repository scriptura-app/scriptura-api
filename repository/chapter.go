package repository

import (
	"encoding/json"
	"scriptura/scriptura-api/db"
	"scriptura/scriptura-api/models"
)

func GetChapter(input int) (models.Chapter, error) {
	db := db.DB
	var chapter models.Chapter

	versSubq := db.Table("verses v").
		Select("json_agg(to_json((SELECT ve FROM (SELECT v.id, v.verse_num as verseNum) ve)))").
		Where("c.id = v.chapter_id")

	bookSubq := db.Table("books b").
		Select("to_json((SELECT bk FROM (SELECT b.id, b.slug) bk))").
		Where("b.id = c.book_id")

	versCountSubq := db.Table("verses v").Select("count(*)").Where("c.id = v.chapter_id")

	db.Table("chapters c").
		Select("c.*, (?) as verses_json, (?) as verse_count, (?) as book_json", versSubq, versCountSubq, bookSubq).
		Where("c.id::varchar ilike ?", input).
		First(&chapter)

	err := json.Unmarshal([]byte(chapter.VersesJson), &chapter.Verses)

	if err != nil {
		return chapter, err
	}

	err2 := json.Unmarshal([]byte(chapter.BookJson), &chapter.Book)

	if err2 != nil {
		return chapter, err2
	}

	return chapter, nil
}
