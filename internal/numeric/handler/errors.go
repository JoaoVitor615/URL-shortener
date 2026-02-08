package handler

import (
	"net/http"

	"github.com/JoaoVitor615/URL-shortener/internal/pkg/apperrors"
)

var (
	ErrInvalidNumericID = apperrors.New("Invalid numeric ID", http.StatusBadRequest)
	ErrURLRequired      = apperrors.New("URL is required", http.StatusBadRequest)
)
