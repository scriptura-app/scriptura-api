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
//	@Summary		Return a range of biblical text
//	@Description	Get biblical text
//	@Tags			Bible
//	@Accept			json
//	@Produce		json
//	@Param			version		query		string		false	"Bible Version"
//	@Param			book		query		string		false	"Book id or shortname"
//	@Param			chapter		query		int			false	"Chapter number"
//	@Param			startVerse	query		int			false	"Verse number"
//	@Param			endVerse	query		int			false	"End verse number"
//	@Param			limit		query		int			false	"Limit"
//	@Param			offset		query		int			false	"Offset"
//	@Param			cursor		query		int			false	"For cursor based pagination (id)"
//	@Success		200			{array}		models.Verse
//	@Failure		400			{boolean}	bool
//	@Failure		404			{boolean}	bool
//	@Failure		500			{boolean}	bool
//	@Router			/bible 	[get]
func (h *bibleHandler) GetByRef(w http.ResponseWriter, r *http.Request) {
	var verses []models.Verse

	limit := utils.GetIntParam(r, "limit", 10)
	offset := utils.GetIntParam(r, "offset", 0)
	cursor := utils.GetIntParam(r, "cursor", 0)
	chapter := utils.GetIntParam(r, "chapter", 0)
	startVerse := utils.GetIntParam(r, "startVerse", 0)
	endVerse := utils.GetIntParam(r, "startVerse", 0)

	input := repository.BibleTextInput{
		Version:    utils.GetParam(r, "version", "en_kj"),
		Book:       utils.GetParam(r, "book", ""),
		Chapter:    chapter,
		StartVerse: startVerse,
		EndVerse:   endVerse,
		Cursor:     cursor,
		Limit:      limit,
		Offset:     offset,
	}

	if input.Book == "" && input.Cursor != 0 {
		input.Book = "1"
	}

	if input.Book != "" && input.Chapter == 0 && input.StartVerse != 0 {
		input.Chapter = 1
	}

	if input.Cursor != 0 && input.Offset != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Can't use offset and cursor at same time"))
		return
	}

	verses, totalItems, _ := h.repository.GetBibleText(input)

	response := utils.FormatPaginationResponse(verses, totalItems, input.Offset, input.Limit, input.Cursor)
	w.Write(response)
}
