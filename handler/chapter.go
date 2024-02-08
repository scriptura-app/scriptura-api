package handler

import (
	"scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetChapter
//
//	@Summary		Get chapter details
//	@Description	Retrieve details of a chapter including book and verses based on input criteria.
//	@Tags			Chapter
//	@Accept			json
//	@Produce		json
//	@Param			input	path		int	 	 true	 "Chapter id"
//	@Success		200	{object}	models.Chapter	"Success"
//	@Failure		400	{object}	interface{}		"Bad Request"
//	@Failure		404	{object}	interface{}		"Not Found"
//	@Failure		500	{object}	interface{}		"Internal Server Error"
//	@Router			/chapter/{input} [get]
func GetChapter(c *fiber.Ctx) error {
	ch, err := strconv.Atoi(c.Params("chapter"))
	if err != nil {
		return c.Status(400).JSON("Chapter ID must be a number")
	}

	chapter, err := repository.GetChapter(ch)
	if err != nil {
		return c.Status(500).JSON("Unknown error")
	}

	if chapter.Id == 0 {
		return c.Status(404).JSON("Book not found")
	}

	response := utils.FormatResponse(chapter)
	return c.JSON(response)
}
