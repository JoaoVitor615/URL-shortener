package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	deps := NewDependencies()
	router := NewRouter(deps)

	assert.NotNil(t, router)
}

func TestPingEndpoint(t *testing.T) {
	deps := NewDependencies()
	router := NewRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "pong", rec.Body.String())
}

func TestNumericTestEndpoint(t *testing.T) {
	deps := NewDependencies()
	router := NewRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/numeric/test", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "test", rec.Body.String())
}

func TestRandomTestEndpoint(t *testing.T) {
	deps := NewDependencies()
	router := NewRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/random/test", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "test", rec.Body.String())
}

func TestNotFoundEndpoint(t *testing.T) {
	deps := NewDependencies()
	router := NewRouter(deps)

	req := httptest.NewRequest(http.MethodGet, "/notfound", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
