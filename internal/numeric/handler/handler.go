package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/JoaoVitor615/URL-shortener/internal/domain"
	"github.com/JoaoVitor615/URL-shortener/internal/numeric/service"
	"github.com/JoaoVitor615/URL-shortener/internal/pkg/apperrors"
)

type NumericHandler struct {
	service service.INumericService
}

func (h *NumericHandler) GetLongURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")

	longURL, err := h.service.GetLongURL(shortURL)
	if err != nil {
		apperrors.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(longURL.LongURL))
}

func (h *NumericHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("longURL")

	if longURL == "" {
		apperrors.WriteError(w, ErrURLRequired)
		return
	}

	url := domain.NewURL(0, longURL)

	shortURL, err := h.service.CreateShortURL(url)
	if err != nil {
		apperrors.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(shortURL))
}
