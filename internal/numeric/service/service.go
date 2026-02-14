package service

import (
	"context"
	"net/http"

	"github.com/JoaoVitor615/URL-shortener/internal/core/encoder"
	"github.com/JoaoVitor615/URL-shortener/internal/core/idgenerator"
	"github.com/JoaoVitor615/URL-shortener/internal/domain"
	"github.com/JoaoVitor615/URL-shortener/internal/pkg/apperrors"
	urlformatter "github.com/JoaoVitor615/URL-shortener/internal/pkg/url_formatter"
)

var (
	ErrInvalidNumericID = apperrors.New("Invalid numeric ID", http.StatusBadRequest)
	ErrIDGeneration     = apperrors.New("Failed to generate ID", http.StatusInternalServerError)
	ErrDatabaseError    = apperrors.New("Database error", http.StatusInternalServerError)
)

type NumericService struct {
	repository domain.IRepositoryShortURL[int]
}

func (s *NumericService) GetLongURL(shortURL string) (url *domain.URL[int], err error) {
	numericID, err := encoder.Decode(shortURL)
	if err != nil {
		return nil, ErrInvalidNumericID
	}

	return s.repository.GetURL(context.Background(), numericID)
}

func (s *NumericService) CreateShortURL(url *domain.URL[int]) (shortURL string, err error) {
	err = url.ValidateLongURL()
	if err != nil {
		return "", err
	}

	url.ID = idgenerator.GenerateID()

	err = s.repository.SaveURL(context.Background(), url)
	if err != nil {
		return "", err
	}

	return urlformatter.FormatURL(encoder.Encode(url.ID))
}
