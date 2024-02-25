package handler

import (
	"net/http"
	"scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"
	"strconv"
)

type ChapterHandler interface {
	GetById(w http.ResponseWriter, r *http.Request)
}

type chapterHandler struct {
	repository repository.ChapterRepository
}

func NewChapterHandler(repo repository.ChapterRepository) ChapterHandler {
	return &chapterHandler{repository: repo}
}

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
func (h *chapterHandler) GetById(w http.ResponseWriter, r *http.Request) {
	chapterId := r.PathValue("id")
	id, err := strconv.Atoi(chapterId)
	if err != nil {
		http.Error(w, "Chapter ID must be a number", http.StatusBadRequest)
		return
	}

	chapter, err := h.repository.GetById(id)
	if err != nil {
		http.Error(w, "Unknown error", http.StatusInternalServerError)
		return
	}

	if chapter.Id == 0 {
		http.Error(w, "Chapter not found", http.StatusNotFound)
		return
	}

	response := utils.FormatResponse(chapter)
	w.Write(response)
}
