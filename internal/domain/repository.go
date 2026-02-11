package domain

type IRepositoryShortURL[T any] interface {
	SaveURL(url *URL[T]) (err error)
	GetURL(id T) (url *URL[T], err error)
}
