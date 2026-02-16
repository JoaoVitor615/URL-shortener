package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/JoaoVitor615/URL-shortener/internal/domain"
)

// MockNumericService is a mock implementation of INumericService
type MockNumericService struct {
	mock.Mock
}

func (m *MockNumericService) GetLongURL(shortURL string) (*domain.URL[int], error) {
	args := m.Called(shortURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.URL[int]), args.Error(1)
}

func (m *MockNumericService) CreateShortURL(url *domain.URL[int]) (string, error) {
	args := m.Called(url)
	return args.String(0), args.Error(1)
}

func TestNumericHandler_GetLongURL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockService := new(MockNumericService)
		handler := &NumericHandler{service: mockService}

		expectedURL := &domain.URL[int]{
			ID:      12345,
			LongURL: "https://google.com",
		}

		mockService.On("GetLongURL", "abc123").Return(expectedURL, nil)

		req := httptest.NewRequest(http.MethodGet, "/numeric/abc123", nil)
		rec := httptest.NewRecorder()

		// Add chi context with URL param
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("shortURL", "abc123")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		handler.GetLongURL(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "https://google.com", rec.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("service error", func(t *testing.T) {
		mockService := new(MockNumericService)
		handler := &NumericHandler{service: mockService}

		serviceErr := errors.New("not found")
		mockService.On("GetLongURL", "invalid").Return(nil, serviceErr)

		req := httptest.NewRequest(http.MethodGet, "/numeric/invalid", nil)
		rec := httptest.NewRecorder()

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("shortURL", "invalid")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		handler.GetLongURL(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockService.AssertExpectations(t)
	})
}

func TestNumericHandler_CreateShortURL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockService := new(MockNumericService)
		handler := &NumericHandler{service: mockService}

		mockService.On("CreateShortURL", mock.AnythingOfType("*domain.URL[int]")).Return("https://short.joaovitor.com/abc123", nil)

		req := httptest.NewRequest(http.MethodPost, "/numeric?url=https://google.com", nil)
		rec := httptest.NewRecorder()

		handler.CreateShortURL(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "https://short.joaovitor.com/abc123", rec.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("service error", func(t *testing.T) {
		mockService := new(MockNumericService)
		handler := &NumericHandler{service: mockService}

		serviceErr := errors.New("failed to create")
		mockService.On("CreateShortURL", mock.AnythingOfType("*domain.URL[int]")).Return("", serviceErr)

		req := httptest.NewRequest(http.MethodPost, "/numeric?url=https://google.com", nil)
		rec := httptest.NewRecorder()

		handler.CreateShortURL(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("empty URL param", func(t *testing.T) {
		mockService := new(MockNumericService)
		handler := &NumericHandler{service: mockService}

		// When URL is empty, service will be called with empty LongURL
		// and should return an error
		serviceErr := domain.ErrLongURLRequired
		mockService.On("CreateShortURL", mock.AnythingOfType("*domain.URL[int]")).Return("", serviceErr)

		req := httptest.NewRequest(http.MethodPost, "/numeric?url=", nil)
		rec := httptest.NewRecorder()

		handler.CreateShortURL(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		mockService.AssertExpectations(t)
	})
}

func TestErrInvalidNumericID(t *testing.T) {
	assert.NotNil(t, ErrInvalidNumericID)
	assert.Equal(t, "Invalid numeric ID", ErrInvalidNumericID.Message)
	assert.Equal(t, http.StatusBadRequest, ErrInvalidNumericID.StatusCode)
}

func TestNewNumericHandler(t *testing.T) {
	mockService := new(MockNumericService)

	handler := NewNumericHandler(mockService)

	assert.NotNil(t, handler)

	// Verify it implements INumericHandler interface
	var _ INumericHandler = handler
}
