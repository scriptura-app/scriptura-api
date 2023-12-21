package handler

import (
	"scriptura/scriptura-api/db"
	m "scriptura/scriptura-api/models"
	"scriptura/scriptura-api/utils"

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

	response, _ := utils.FormatResponse(verses, int(totalItems), offset, limit)

	return c.JSON(response)
}
