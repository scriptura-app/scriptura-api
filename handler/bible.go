package handler

import (
	"scriptura/scriptura-api/models"
	"scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"

	"github.com/gofiber/fiber/v2"
)

// ListAccounts lists all existing accounts
//
//	@Summary		Return a range of biblical
//	@Description	get accounts
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			q	query		string	false	"name search by q"	Format(email)
//	@Success		200	{array}		models.Verse
//	@Failure		400	{boolean}	bool
//	@Failure		404	{boolean}	bool
//	@Failure		500	{boolean}	bool
//	@Router			/accounts [get]
func GetBible(c *fiber.Ctx) error {
	var verses []models.Verse

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
