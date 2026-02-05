package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/JoaoVitor615/URL-shortener/internal/random/service"
)

func TestNewURLRandomHandler(t *testing.T) {
	svc := service.NewURLRandomService()
	handler := NewURLRandomHandler(svc)

	assert.NotNil(t, handler)
	assert.NotNil(t, handler.service)
}

func TestURLRandomHandler_Test(t *testing.T) {
	svc := service.NewURLRandomService()
	handler := NewURLRandomHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/random/test", nil)
	rec := httptest.NewRecorder()

	handler.Test(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "test", rec.Body.String())
}
