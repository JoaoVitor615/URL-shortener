package service

import (
	"errors"
	"fmt"

	"github.com/JoaoVitor615/URL-shortener/internal/core/encoder"
	"github.com/JoaoVitor615/URL-shortener/internal/core/idgenerator"
)

type NumericService struct {
}

func (s *NumericService) GetLongURL(numericID int) (longURL string, err error) {
	// get the longURL from the database

	return longURL, nil
}

func (s *NumericService) CreateShortURL(longURL string) (shortURL string, err error) {
	numericID, err := s.generateNumericID()
	fmt.Println(numericID, err)
	if err != nil {
		return "", err
	}

	// save the longURL and the numericID in the database

	shortURL = encoder.Encode(numericID)

	return shortURL, nil
}

func (s *NumericService) generateNumericID() (numericID int, err error) {
	numericID = idgenerator.GenerateID()

	if numericID == 0 {
		return 0, errors.New("failed to generate numeric ID")
	}

	return numericID, nil
}
