package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/JoaoVitor615/URL-shortener/internal/numeric/service"
)

func TestNewURLNumericHandler(t *testing.T) {
	svc := service.NewURLNumericService()
	handler := NewURLNumericHandler(svc)

	assert.NotNil(t, handler)
	assert.NotNil(t, handler.service)
}

func TestURLNumericHandler_Test(t *testing.T) {
	svc := service.NewURLNumericService()
	handler := NewURLNumericHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/numeric/test", nil)
	rec := httptest.NewRecorder()

	handler.Test(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "test", rec.Body.String())
}
