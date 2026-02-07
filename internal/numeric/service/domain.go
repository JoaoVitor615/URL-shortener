package service

type INumericService interface {
	GetLongURL(numericID int) (longURL string, err error)
	CreateShortURL(longURL string) (shortURL string, err error)
}

func NewNumericService() INumericService {
	return &NumericService{}
}
