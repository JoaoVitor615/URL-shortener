package urlformatter

import (
	"net/http"

	"github.com/JoaoVitor615/URL-shortener/internal/pkg/apperrors"
)

var (
	ErrInvalidURL = apperrors.New("Invalid URL", http.StatusBadRequest)
)

const (
	BaseURL = "https://short.joaovitor.com/"
)

func FormatURL(url string) (string, error) {
	if url == "" {
		return "", ErrInvalidURL
	}

	return BaseURL + url, nil
}
