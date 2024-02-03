package handler

import (
	"scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"

	"github.com/gofiber/fiber/v2"
)

// GetChapter
//
//	@Summary		Get chapter details
//	@Description	Retrieve details of a chapter including book and verses based on input criteria.
//	@Tags			Chapter
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int				true	"Chapter id"
//	@Success		200	{object}	models.Chapter	"Success"
//	@Failure		400	{object}	interface{}		"Bad Request"
//	@Failure		404	{object}	interface{}		"Not Found"
//	@Failure		500	{object}	interface{}		"Internal Server Error"
//	@Router			/chapter/{input} [get]
func GetChapter(c *fiber.Ctx) error {
	ch := c.Params("chapter")
	chapter, _ := repository.GetChapter(ch)
	response := utils.FormatResponse(chapter)
	return c.JSON(response)
}
