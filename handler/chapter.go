package handler

import (
	"github.com/gofiber/fiber/v2"
)

func GetChapter(c *fiber.Ctx) error {
	return c.JSON(c.Params("chapter"))
}
