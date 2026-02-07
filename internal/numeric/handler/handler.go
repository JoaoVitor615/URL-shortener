package handler

import (
	"net/http"
	"strconv"

	"github.com/JoaoVitor615/URL-shortener/internal/numeric/service"
	"github.com/go-chi/chi/v5"
)

type NumericHandler struct {
	service service.INumericService
}

func (h *NumericHandler) GetLongURL(w http.ResponseWriter, r *http.Request) {
	numericID := chi.URLParam(r, "numericID")

	numericIDInt, err := strconv.Atoi(numericID)
	if err != nil {
		http.Error(w, "Invalid numeric ID", http.StatusBadRequest)
		return
	}

	longURL, err := h.service.GetLongURL(numericIDInt)
	if err != nil {
		http.Error(w, "Failed to get long URL", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(longURL))
}

func (h *NumericHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("longURL")

	if longURL == "" {
		http.Error(w, "Long URL is required", http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.CreateShortURL(longURL)
	if err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(shortURL))
}
