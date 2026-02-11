package domain

import (
	"time"
)

type URL[T any] struct {
	ID        T         `json:"id"`
	LongURL   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
}

func NewURL[T any](id T, longURL string) *URL[T] {
	return &URL[T]{
		ID:        id,
		LongURL:   longURL,
		CreatedAt: time.Now(),
	}
}
