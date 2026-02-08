package service

import (
	"net/http"

	"github.com/JoaoVitor615/URL-shortener/internal/pkg/apperrors"
)

var (
	ErrURLNotFound   = apperrors.New("URL not found", http.StatusNotFound)
	ErrIDGeneration  = apperrors.New("Failed to generate ID", http.StatusInternalServerError)
	ErrDatabaseError = apperrors.New("Database error", http.StatusInternalServerError)
)
