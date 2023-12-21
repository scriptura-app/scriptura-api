package handler

import (
	"scriptura/scriptura-api/db"
	m "scriptura/scriptura-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetVerse(c *fiber.Ctx) error {
	db := db.DB
	var response m.SliceResponse
	var verses []m.Verse

	db.Table("verse").
		Select("verse.*, bible.text, book.name as book_name").
		Joins("left join book on book.id = book_id").
		Joins("left join bible_en_kjv bible on bible.verse_id = verse.id").
		Where("book.id::varchar ilike $1 OR book.code ilike $1 OR book.short_name ilike $1", c.Params("book")).
		Offset(10).
		Limit(10).
		Scan(&verses)

	// response.MapSlice(verses, "verse")
	response.Data = mapVersesToResponseItems(verses)

	return c.JSON(response)
}

func mapVersesToResponseItems(verses []m.Verse) []m.ResponseItem {
	responseItems := make([]m.ResponseItem, len(verses))
	for i, verse := range verses {
		responseItems[i] = m.ResponseItem{
			Id:         verse.ID,
			Type:       "verse",
			Attributes: verse.ClearAtributes(),
		}
	}
	return responseItems
}
