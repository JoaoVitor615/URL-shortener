package service

import (
	"net/http"

	"github.com/JoaoVitor615/URL-shortener/internal/core/encoder"
	"github.com/JoaoVitor615/URL-shortener/internal/core/idgenerator"
	"github.com/JoaoVitor615/URL-shortener/internal/domain"
	"github.com/JoaoVitor615/URL-shortener/internal/pkg/apperrors"
)

var (
	ErrURLNotFound   = apperrors.New("URL not found", http.StatusNotFound)
	ErrIDGeneration  = apperrors.New("Failed to generate ID", http.StatusInternalServerError)
	ErrDatabaseError = apperrors.New("Database error", http.StatusInternalServerError)
)

type NumericService struct {
	repository domain.IRepositoryShortURL[int]
}

func (s *NumericService) GetLongURL(shortURL string) (url *domain.URL[int], err error) {
	numericID, err := encoder.Decode(shortURL)
	if err != nil {
		return nil, err
	}

	return s.repository.GetURL(numericID)
}

func (s *NumericService) CreateShortURL(url *domain.URL[int]) (shortURL string, err error) {
	url.ID = idgenerator.GenerateID()

	err = s.repository.SaveURL(url)
	if err != nil {
		return "", err
	}

	return encoder.Encode(url.ID), nil
}
