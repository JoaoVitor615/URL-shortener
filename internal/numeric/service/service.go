package service

import (
	"github.com/JoaoVitor615/URL-shortener/internal/core/encoder"
	"github.com/JoaoVitor615/URL-shortener/internal/core/idgenerator"
)

type NumericService struct {
}

func (s *NumericService) GetLongURL(numericID int) (longURL string, err error) {
	// TODO: get the longURL from the database
	// For now, return not found if ID is 0
	if numericID == 0 {
		return "", ErrURLNotFound
	}

	return longURL, nil
}

func (s *NumericService) CreateShortURL(longURL string) (shortURL string, err error) {
	numericID, err := s.generateNumericID()
	if err != nil {
		return "", err
	}

	// TODO: save the longURL and the numericID in the database

	shortURL = encoder.Encode(numericID)

	return shortURL, nil
}

func (s *NumericService) generateNumericID() (numericID int, err error) {
	numericID = idgenerator.GenerateID()

	if numericID == 0 {
		return 0, ErrIDGeneration
	}

	return numericID, nil
}
