package handler

import (
	r "scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetChapter(c *fiber.Ctx) error {
	bk, ch := c.Params("book"), c.Params("chapter")
	chapter, _ := r.GetChapter(bk, ch)
	response := utils.FormatResponse(chapter)
	return c.JSON(response)
}
