package domain

import "context"

type IRepositoryShortURL[T any] interface {
	SaveURL(ctx context.Context, url *URL[T]) (err error)
	GetURL(ctx context.Context, id T) (url *URL[T], err error)
	GetLongURL(ctx context.Context, longURL string) (url *URL[T], err error)
}
