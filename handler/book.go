//go:build exclude

package handler

import (
	r "scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetBook(c *fiber.Ctx) error {
	bk := c.Params("book")
	book, _ := r.GetBook(bk)
	response := utils.FormatResponse(book)
	return c.JSON(response)
}
