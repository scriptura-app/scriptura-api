package handler

import (
	"net/http"
	"scriptura/scriptura-api/models"
	"scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"
)

type BibleHandler interface {
	GetByRef(w http.ResponseWriter, r *http.Request)
}

type bibleHandler struct {
	repository repository.BibleRepository
}

func NewBibleHandler(repo repository.BibleRepository) BibleHandler {
	return &bibleHandler{repository: repo}
}

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
func (h *bibleHandler) GetByRef(w http.ResponseWriter, r *http.Request) {
	var verses []models.Verse

	// offset, limit := c.Locals("offset").(int), c.Locals("limit").(int)

	i := repository.BibleTextInput{
		Bible:      "en_kj",
		Book:       r.PathValue("book"),
		Chapter:    r.PathValue("chapter"),
		StartVerse: r.PathValue("start"),
		EndVerse:   r.PathValue("end"),
		// Offset:     c.Locals("offset").(int),
		// Limit:      c.Locals("limit").(int),
	}

	verses, totalItems, _ := h.repository.GetBibleText(i)

	response := utils.FormatPaginationResponse(verses, totalItems, 0, 0)
	w.Write(response)
}
