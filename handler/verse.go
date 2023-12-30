package handler

import (
	m "scriptura/scriptura-api/models"
	r "scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetSingleVerse(c *fiber.Ctx) error {
	bk, ch, vr := c.Params("book"), c.Params("chapter"), c.Params("verse")

	verse, _ := r.GetVerseByRef("en_kjv", bk, ch, vr)

	response := utils.FormatResponse(verse)
	return c.JSON(response)
}

func GetVerseRange(c *fiber.Ctx) error {
	var verses []m.Verse

	offset, limit := c.Locals("offset").(int), c.Locals("limit").(int)
	bk, ch, fr, to := c.Params("book"), c.Params("chapter"), c.Params("from"), c.Params("to")

	verses, totalItems, _ := r.GetVerseRangeByRef("en_kjv", bk, ch, fr, to, offset, limit)

	response, _ := utils.FormatPaginationResponse(verses, int(totalItems), offset, limit)

	return c.JSON(response)
}
