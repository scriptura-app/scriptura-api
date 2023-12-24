package handler

import (
	"github.com/gofiber/fiber/v2"
)

func GetBook(c *fiber.Ctx) error {
	return c.JSON(c.Params("book"))
}
