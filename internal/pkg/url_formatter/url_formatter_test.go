package urlformatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatURL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result, err := FormatURL("abc123")

		assert.NoError(t, err)
		assert.Equal(t, "https://short.joaovitor.com/abc123", result)
	})

	t.Run("empty URL returns error", func(t *testing.T) {
		result, err := FormatURL("")

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Equal(t, ErrInvalidURL, err)
	})
}

func TestErrInvalidURL(t *testing.T) {
	assert.NotNil(t, ErrInvalidURL)
	assert.Equal(t, "Invalid URL", ErrInvalidURL.Message)
	assert.Equal(t, 400, ErrInvalidURL.StatusCode)
}

func TestBaseURL(t *testing.T) {
	assert.Equal(t, "https://short.joaovitor.com/", BaseURL)
}
