package handler

import (
	"scriptura/scriptura-api/db"
	m "scriptura/scriptura-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetVerse(c *fiber.Ctx) error {
	db := db.DB
	var verses []m.Verse
	var totalCount int64

	query := db.Table("verse").
		Select("verse.*, bible.text, book.name as book_name, COUNT(*) OVER () as TotalCount").
		Joins("left join book on book.id = book_id").
		Joins("left join bible_en_kjv bible on bible.verse_id = verse.id").
		Where("book.id::varchar ilike $1 OR book.code ilike $1 OR book.short_name ilike $1", c.Params("book"))

	query.Count(&totalCount)

	query.Offset(c.Locals("offset").(int)).
		Limit(c.Locals("limit").(int)).
		Scan(&verses)

	response, _ := formatVerseResponse(verses, totalCount)

	return c.JSON(response)
}

func formatVerseResponse(verses []m.Verse, totalCount int64) (m.SliceResponse, error) {

	responseItems := make([]m.ResponseItem, len(verses))
	for i, verse := range verses {
		responseItems[i] = m.ResponseItem{
			Id:         verse.ID,
			Type:       "verse",
			Attributes: verse,
		}
	}

	var response m.SliceResponse
	response.Data = responseItems
	response.Meta.Pagination.TotalCount = totalCount

	return response, nil
}
