package handler

import (
	m "scriptura/scriptura-api/models"
	"scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetBible(c *fiber.Ctx) error {
	var verses []m.Verse

	offset, limit := c.Locals("offset").(int), c.Locals("limit").(int)

	i := repository.BibleTextInput{
		Bible:      "en_kj",
		Book:       c.Params("book"),
		Chapter:    c.Params("chapter"),
		StartVerse: c.Params("start"),
		EndVerse:   c.Params("end"),
		Offset:     c.Locals("offset").(int),
		Limit:      c.Locals("limit").(int),
	}

	verses, totalItems, _ := repository.GetBibleText(i)

	response, _ := utils.FormatPaginationResponse(verses, totalItems, offset, limit)

	return c.JSON(response)
}
