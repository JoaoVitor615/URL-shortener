package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewURL(t *testing.T) {
	t.Run("with int ID", func(t *testing.T) {
		before := time.Now()
		url := NewURL(123, "https://google.com")
		after := time.Now()

		assert.Equal(t, 123, url.ID)
		assert.Equal(t, "https://google.com", url.LongURL)
		assert.True(t, url.CreatedAt.After(before) || url.CreatedAt.Equal(before))
		assert.True(t, url.CreatedAt.Before(after) || url.CreatedAt.Equal(after))
	})

	t.Run("with string ID", func(t *testing.T) {
		url := NewURL("abc123", "https://github.com")

		assert.Equal(t, "abc123", url.ID)
		assert.Equal(t, "https://github.com", url.LongURL)
		assert.NotZero(t, url.CreatedAt)
	})

	t.Run("with zero ID", func(t *testing.T) {
		url := NewURL(0, "https://example.com")

		assert.Equal(t, 0, url.ID)
		assert.Equal(t, "https://example.com", url.LongURL)
	})

	t.Run("with empty long URL", func(t *testing.T) {
		url := NewURL(1, "")

		assert.Equal(t, 1, url.ID)
		assert.Empty(t, url.LongURL)
	})
}

func TestURL_ValidateLongURL(t *testing.T) {
	t.Run("valid URL", func(t *testing.T) {
		url := NewURL(1, "https://google.com")

		err := url.ValidateLongURL()

		assert.NoError(t, err)
	})

	t.Run("empty URL returns error", func(t *testing.T) {
		url := NewURL(1, "")

		err := url.ValidateLongURL()

		assert.Error(t, err)
		assert.Equal(t, ErrLongURLRequired, err)
	})

	t.Run("whitespace only is considered empty", func(t *testing.T) {
		url := &URL[int]{ID: 1, LongURL: ""}

		err := url.ValidateLongURL()

		assert.Error(t, err)
		assert.Equal(t, ErrLongURLRequired, err)
	})
}

func TestErrLongURLRequired(t *testing.T) {
	assert.NotNil(t, ErrLongURLRequired)
	assert.Equal(t, "long url is required", ErrLongURLRequired.Message)
	assert.Equal(t, 400, ErrLongURLRequired.StatusCode)
}
