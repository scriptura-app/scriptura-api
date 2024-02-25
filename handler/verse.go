package handler

import (
	"net/http"
	"scriptura/scriptura-api/repository"
	"scriptura/scriptura-api/utils"
	"strconv"
)

type VerseHandler interface {
	GetById(w http.ResponseWriter, r *http.Request)
}

type verseHandler struct {
	repository repository.VerseRepository
}

func NewVerseHandler(repo repository.VerseRepository) VerseHandler {
	return &verseHandler{repository: repo}
}

func (h *verseHandler) GetById(w http.ResponseWriter, r *http.Request) {
	verseId, _ := strconv.Atoi(r.PathValue("id"))
	verse, _ := h.repository.GetById(verseId)
	response := utils.FormatResponse(verse)
	w.Write(response)
}
