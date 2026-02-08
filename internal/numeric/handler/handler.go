package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/JoaoVitor615/URL-shortener/internal/numeric/service"
	"github.com/JoaoVitor615/URL-shortener/internal/pkg/apperrors"
)

type NumericHandler struct {
	service service.INumericService
}

func (h *NumericHandler) GetLongURL(w http.ResponseWriter, r *http.Request) {
	numericID := chi.URLParam(r, "numericID")

	numericIDInt, err := strconv.Atoi(numericID)
	if err != nil {
		apperrors.WriteError(w, ErrInvalidNumericID)
		return
	}

	longURL, err := h.service.GetLongURL(numericIDInt)
	if err != nil {
		apperrors.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(longURL))
}

func (h *NumericHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("longURL")

	if longURL == "" {
		apperrors.WriteError(w, ErrURLRequired)
		return
	}

	shortURL, err := h.service.CreateShortURL(longURL)
	if err != nil {
		apperrors.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(shortURL))
}
