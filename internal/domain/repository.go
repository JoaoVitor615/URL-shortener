package repository

type IRepositoryShortURL[T any] interface {
	SaveId(id T, longURL string) (err error)
	GetId(id T) (longURL string, err error)
}
