package handler

import (
	"net/http"
	"scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"
)

type BookHandler interface {
	GetById(w http.ResponseWriter, r *http.Request)
}

type bookHandler struct {
	repository repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) BookHandler {
	return &bookHandler{repository: repo}
}

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
func (h *bookHandler) GetById(w http.ResponseWriter, r *http.Request) {
	bookId := r.PathValue("id")
	book, _ := h.repository.GetById(bookId)

	//TODO
	//if book.Id == 0 {
	//	return c.Status(404).JSON("Book not found")
	//}

	response := utils.FormatResponse(book)
	w.Write(response)
}
