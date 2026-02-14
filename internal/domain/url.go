package domain

import (
	"net/http"
	"time"

	"github.com/JoaoVitor615/URL-shortener/internal/pkg/apperrors"
)

type URL[T any] struct {
	ID        T         `json:"id" dynamodbav:"id"`
	LongURL   string    `json:"long_url" dynamodbav:"long_url"`
	CreatedAt time.Time `json:"created_at" dynamodbav:"created_at"`
}

var (
	ErrLongURLRequired = apperrors.New("long url is required", http.StatusBadRequest)
)

func NewURL[T any](id T, longURL string) *URL[T] {
	return &URL[T]{
		ID:        id,
		LongURL:   longURL,
		CreatedAt: time.Now(),
	}
}

func (u *URL[T]) ValidateLongURL() error {
	if u.LongURL == "" {
		return ErrLongURLRequired
	}
	return nil
}
