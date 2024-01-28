package handler

import (
	r "scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"

	"github.com/gofiber/fiber/v2"
)

// GetBook
//
//	@Summary		Get book details
//	@Description	Retrieve details of a book including chapters and verses count based on input criteria.
//	@Tags			Book
//	@Accept			json
//	@Produce		json
//	@Param			input	path		string		true	"Input for Book Search (ID, Slug, or Short Name)"
//	@Success		200		{object}	models.Book	"Success"
//	@Failure		400		{object}	interface{}	"Bad Request"
//	@Failure		404		{object}	interface{}	"Not Found"
//	@Failure		500		{object}	interface{}	"Internal Server Error"
//	@Router			/book/{input} [get]
func GetBook(c *fiber.Ctx) error {
	bk := c.Params("book")
	book, _ := r.GetBook(bk)

	if book.Id == 0 {
		return c.Status(404).JSON("Book not found")
	}
	response := utils.FormatResponse(book)
	return c.JSON(response)
}
