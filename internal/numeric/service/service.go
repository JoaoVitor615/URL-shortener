package service

import (
	"github.com/JoaoVitor615/URL-shortener/internal/core/encoder"
	"github.com/JoaoVitor615/URL-shortener/internal/core/idgenerator"
	"github.com/JoaoVitor615/URL-shortener/internal/domain"
)

type NumericService struct {
	repository domain.IRepositoryShortURL[int]
}

func (s *NumericService) GetLongURL(shortURL string) (url *domain.URL[int], err error) {
	numericID := encoder.Decode(shortURL)

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
