package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := New("test message", http.StatusBadRequest)

	assert.Equal(t, "test message", err.Message)
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Empty(t, err.Err)
}

func TestWrap(t *testing.T) {
	originalErr := errors.New("original error")
	err := Wrap(originalErr, "wrapped message", http.StatusInternalServerError)

	assert.Equal(t, "wrapped message", err.Message)
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "original error", err.Err)
}

func TestNewWithErr(t *testing.T) {
	errFactory := NewWithErr("database error", http.StatusInternalServerError)
	originalErr := errors.New("connection failed")
	
	err := errFactory(originalErr)

	assert.Equal(t, "database error", err.Message)
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "connection failed", err.Err)
}

func TestAppError_Error(t *testing.T) {
	t.Run("without underlying error", func(t *testing.T) {
		err := New("test message", http.StatusBadRequest)
		assert.Equal(t, "test message", err.Error())
	})

	t.Run("with underlying error", func(t *testing.T) {
		originalErr := errors.New("original error")
		err := Wrap(originalErr, "wrapped message", http.StatusInternalServerError)
		assert.Equal(t, "wrapped message: original error", err.Error())
	})
}

func TestAppError_Unwrap(t *testing.T) {
	originalErr := errors.New("original error")
	err := Wrap(originalErr, "wrapped message", http.StatusInternalServerError)

	unwrapped := err.Unwrap()
	assert.Equal(t, "original error", unwrapped.Error())
}

func TestWriteError_AppError(t *testing.T) {
	rec := httptest.NewRecorder()
	err := New("URL not found", http.StatusNotFound)

	WriteError(rec, err)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var response ErrorResponse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "URL not found", response.Error.Message)
	assert.Equal(t, http.StatusNotFound, response.Error.StatusCode)
}

func TestWriteError_GenericError(t *testing.T) {
	rec := httptest.NewRecorder()
	err := errors.New("some generic error")

	WriteError(rec, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var response ErrorResponse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "An unexpected error occurred", response.Error.Message)
	assert.Equal(t, http.StatusInternalServerError, response.Error.StatusCode)
}
