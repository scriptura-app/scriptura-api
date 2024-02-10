package handler

import (
	"net/http"
	"scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"

	"github.com/go-chi/chi/v5"
)

type BookHandler interface {
	GetBook(w http.ResponseWriter, r *http.Request)
}

type bookHandler struct {
	repository repository.BookRepository
}

func NewBookHandler(r repository.BookRepository) BookHandler {
	return &bookHandler{
		repository: r,
	}
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
func (h *bookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	bookInput := chi.URLParam(r, "book")
	book, _ := h.repository.GetBook(bookInput)

	//TODO
	//if book.Id == 0 {
	//	return c.Status(404).JSON("Book not found")
	//}

	response := utils.FormatResponse(book)
	w.Write(response)
}
