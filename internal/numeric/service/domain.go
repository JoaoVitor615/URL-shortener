package service

import "github.com/JoaoVitor615/URL-shortener/internal/domain"

type INumericService interface {
	GetLongURL(shortURL string) (url *domain.URL[int], err error)
	CreateShortURL(url *domain.URL[int]) (shortURL string, err error)
}

func NewNumericService(repository domain.IRepositoryShortURL[int]) INumericService {
	return &NumericService{
		repository: repository,
	}
}
