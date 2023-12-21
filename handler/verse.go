package handler

import (
	"scriptura/scriptura-api/db"
	m "scriptura/scriptura-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetVerse(c *fiber.Ctx) error {
	db := db.DB
	var verses []m.Verse
	var totalItems int64

	offset, limit := c.Locals("offset").(int), c.Locals("limit").(int)

	query := db.Table("verse").
		Select("verse.*, bible.text, book.name as book_name, COUNT(*) OVER () as totalItems").
		Joins("left join book on book.id = book_id").
		Joins("left join bible_en_kjv bible on bible.verse_id = verse.id").
		Where("book.id::varchar ilike $1 OR book.code ilike $1 OR book.short_name ilike $1", c.Params("book"))

	query.Count(&totalItems)

	query.Offset(offset).
		Limit(limit).
		Scan(&verses)

	response, _ := formatVerseResponse(verses, int(totalItems), offset, limit)

	return c.JSON(response)
}

func formatVerseResponse(verses []m.Verse, totalItems int, offset int, limit int) (m.SliceResponse, error) {

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

	response.Meta.Pagination.TotalItems = totalItems
	response.Meta.Pagination.CurrentPage = offset/limit + 1
	response.Meta.Pagination.PageSize = limit
	response.Meta.Pagination.TotalPages = (totalItems + limit - 1) / limit

	return response, nil
}
